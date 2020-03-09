package gomavlib

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// DEFINE PUBLIC TYPES AND STRUCTURES.

// Dialect contains available messages and the configuration needed to encode and decode them.
type DialectCT struct {
	version  uint
	messages map[uint32]*dialectMessageCT
}

// Assert to check we're implementing the interfaces we expect to be.
var _ Dialect = &DialectCT{}

// DEFINE PRIVATE TYPES AND STRUCTURES.

type dialectMessageCT struct {
	elemType     reflect.Type
	fields       []*dialectMessageField
	sizeNormal   byte
	sizeExtended byte
	crcExtra     byte
}

// Assert to check we're implementing the interfaces we expect to be.
var _ dialectMessage = &dialectMessageCT{}

// DEFINE PUBLIC STATUC FUNCTIONS.

// NewDialect allocates a Dialect.
func NewDialectCT(version uint, messages []Message) (*DialectCT, error) {
	d := &DialectCT{
		version:  version,
		messages: make(map[uint32]*dialectMessageCT),
	}

	for _, msg := range messages {
		mp, err := newDialectMessage(msg)
		if err != nil {
			return nil, fmt.Errorf("message %T: %s", msg, err)
		}
		d.messages[msg.GetId()] = mp
	}

	return d, nil
}

// MustDialect is like NewDialect but panics in case of error.
func MustDialectCT(version uint, messages []Message) *DialectCT {
	d, err := NewDialectCT(version, messages)
	if err != nil {
		panic(err)
	}
	return d
}

// DEFINE PUBLIC RECEIVER FUNCTIONS.

// DialectCT

func (d *DialectCT) getVersion() uint {
	return d.version
}

func (d *DialectCT) getMsgById(id uint32) (*dialectMessage, bool) {
	var msg dialectMessage
	var ok bool
	msg, ok = d.messages[id]
	return &msg, ok
}

// DEFINE PRIVATE STATIC FUNCTIONS.

func newDialectMessage(msg Message) (*dialectMessageCT, error) {
	mp := &dialectMessageCT{}
	mp.elemType = reflect.TypeOf(msg).Elem()

	mp.fields = make([]*dialectMessageField, mp.elemType.NumField())

	// get name
	if strings.HasPrefix(mp.elemType.Name(), "Message") == false {
		return nil, fmt.Errorf("message struct name must begin with 'Message'")
	}
	msgName := dialectMsgGoToDef(mp.elemType.Name()[len("Message"):])

	// collect message fields
	for i := 0; i < mp.elemType.NumField(); i++ {
		field := mp.elemType.Field(i)
		arrayLength := byte(0)
		goType := field.Type

		// array
		if goType.Kind() == reflect.Array {
			arrayLength = byte(goType.Len())
			goType = goType.Elem()
		}

		isEnum := false
		var dialectType dialectFieldType

		// enum
		if field.Tag.Get("mavenum") != "" {
			isEnum = true

			if goType.Kind() != reflect.Int {
				return nil, fmt.Errorf("an enum must be an int")
			}

			tagEnum := field.Tag.Get("mavenum")

			if len(tagEnum) == 0 {
				return nil, fmt.Errorf("enum but tag not specified")
			}

			dialectType = dialectFieldTypeFromGo[tagEnum]
			if dialectType == 0 {
				return nil, fmt.Errorf("invalid go type: %v", tagEnum)
			}

			switch dialectType {
			case typeUint8:
			case typeUint16:
			case typeUint32:
			case typeInt32:
			case typeUint64:
				break

			default:
				return nil, fmt.Errorf("type %v cannot be used as enum", dialectType)
			}

		} else {
			dialectType = dialectFieldTypeFromGo[goType.Name()]
			if dialectType == 0 {
				return nil, fmt.Errorf("invalid go type: %v", goType.Name())
			}

			// string or char
			if goType.Kind() == reflect.String {
				tagLen := field.Tag.Get("mavlen")

				// char
				if len(tagLen) == 0 {
					arrayLength = 1

					// string
				} else {
					slen, err := strconv.Atoi(tagLen)
					if err != nil {
						return nil, fmt.Errorf("string has invalid length: %v", tagLen)
					}
					arrayLength = byte(slen)
				}
			}
		}

		// extension
		isExtension := (field.Tag.Get("mavext") == "true")

		// size
		var size byte
		if arrayLength > 0 {
			size = dialectFieldTypeSizes[dialectType] * arrayLength
		} else {
			size = dialectFieldTypeSizes[dialectType]
		}

		mp.fields[i] = &dialectMessageField{
			isEnum: isEnum,
			ftype:  dialectType,
			name: func() string {
				if mavname := field.Tag.Get("mavname"); mavname != "" {
					return mavname
				}
				return dialectFieldGoToDef(field.Name)
			}(),
			arrayLength: arrayLength,
			index:       i,
			isExtension: isExtension,
		}

		mp.sizeExtended += size
		if isExtension == false {
			mp.sizeNormal += size
		}
	}

	// reorder fields as described in
	// https://mavlink.io/en/guide/serialization.html#field_reordering
	sort.Slice(mp.fields, func(i, j int) bool {
		// sort by weight if not extension
		if mp.fields[i].isExtension == false && mp.fields[j].isExtension == false {
			if w1, w2 := dialectFieldTypeSizes[mp.fields[i].ftype], dialectFieldTypeSizes[mp.fields[j].ftype]; w1 != w2 {
				return w1 > w2
			}
		}
		// sort by original index
		return mp.fields[i].index < mp.fields[j].index
	})

	// generate CRC extra
	// https://mavlink.io/en/guide/serialization.html#crc_extra
	mp.crcExtra = func() byte {
		h := NewX25()
		h.Write([]byte(msgName + " "))

		for _, f := range mp.fields {
			// skip extensions
			if f.isExtension == true {
				continue
			}

			h.Write([]byte(dialectFieldTypeString[f.ftype] + " "))
			h.Write([]byte(f.name + " "))

			if f.arrayLength > 0 {
				h.Write([]byte{f.arrayLength})
			}
		}
		sum := h.Sum16()
		return byte((sum & 0xFF) ^ (sum >> 8))
	}()

	return mp, nil
}

// DEFINE PRIVATE RECEIVER FUNCTIONS.

func (mp *dialectMessageCT) newMsg() *Message {
	ref := reflect.New(mp.elemType)
	msg := ref.Interface().(Message)
	return &msg
}

func (mp *dialectMessageCT) getFields() []*dialectMessageField {
	return mp.fields
}

func (mp *dialectMessageCT) getCRCExtra() byte {
	return mp.crcExtra
}

func (mp *dialectMessageCT) decode(buf []byte, isFrameV2 bool) (Message, error) {
	msg := reflect.New(mp.elemType)

	if isFrameV2 == true {
		// in V2 buffer length can be > message or < message
		// in this latter case it must be filled with zeros to support empty-byte de-truncation
		// and extension fields
		if len(buf) < int(mp.sizeExtended) {
			buf = append(buf, bytes.Repeat([]byte{0x00}, int(mp.sizeExtended)-len(buf))...)
		}

	} else {
		// in V1 buffer must fit message perfectly
		if len(buf) != int(mp.sizeNormal) {
			return nil, fmt.Errorf("unexpected size (%d vs %d)", len(buf), mp.sizeNormal)
		}
	}

	// decode field by field
	for _, f := range mp.fields {
		// skip extensions in V1 frames
		if isFrameV2 == false && f.isExtension == true {
			continue
		}

		target := msg.Elem().Field(f.index)

		switch target.Kind() {
		case reflect.Array:
			length := target.Len()
			for i := 0; i < length; i++ {
				n := decodeValue(target.Index(i), buf, f)
				buf = buf[n:]
			}

		default:
			n := decodeValue(target, buf, f)
			buf = buf[n:]
		}
	}

	return msg.Interface().(Message), nil
}

func (mp *dialectMessageCT) encode(msg Message, isFrameV2 bool) ([]byte, error) {
	var buf []byte

	if isFrameV2 == true {
		buf = make([]byte, mp.sizeExtended)
	} else {
		buf = make([]byte, mp.sizeNormal)
	}

	start := buf

	// encode field by field
	for _, f := range mp.fields {
		// skip extensions in V1 frames
		if isFrameV2 == false && f.isExtension == true {
			continue
		}

		target := reflect.ValueOf(msg).Elem().Field(f.index)

		switch target.Kind() {
		case reflect.Array:
			length := target.Len()
			for i := 0; i < length; i++ {
				n := encodeValue(buf, target.Index(i), f)
				buf = buf[n:]
			}

		default:
			n := encodeValue(buf, target, f)
			buf = buf[n:]
		}
	}

	buf = start

	// empty-byte truncation
	// even with truncation, message length must be at least 1 byte
	// https://github.com/mavlink/c_library_v2/blob/master/mavlink_helpers.h#L103
	if isFrameV2 == true {
		end := len(buf)
		for end > 1 && buf[end-1] == 0x00 {
			end--
		}
		buf = buf[:end]
	}

	return buf, nil
}

// ALL DONE.
