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

// DialectCT : compile time dialect struct
type DialectCT struct {
	version  uint
	Messages map[uint32]*DialectMessageCT
}

// Assert to check we're implementing the interfaces we expect to be.
var _ Dialect = &DialectCT{}

// DEFINE PRIVATE TYPES AND STRUCTURES.

type DialectMessageCT struct {
	elemType     reflect.Type
	Fields       []*DialectMessageField
	sizeNormal   byte
	sizeExtended byte
	crcExtra     byte
}

// Assert to check we're implementing the interfaces we expect to be.
var _ dialectMessage = &DialectMessageCT{}

// DEFINE PUBLIC STATUC FUNCTIONS.

// NewDialectCT allocates a Dialect.
func NewDialectCT(version uint, messages []Message) (*DialectCT, error) {
	d := &DialectCT{
		version:  version,
		Messages: make(map[uint32]*DialectMessageCT),
	}

	for _, msg := range messages {
		mp, err := newDialectMessage(msg)
		if err != nil {
			return nil, fmt.Errorf("message %T: %s", msg, err)
		}
		d.Messages[msg.GetId()] = mp
	}

	return d, nil
}

// MustDialectCT is like NewDialect but panics in case of error.
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
	msg, ok = d.Messages[id]
	return &msg, ok
}

// DialectMessageCT

// GetSizeNormal returns sizeNormal of DialectMessageCT
func (mp *DialectMessageCT) GetSizeNormal() byte {
	return mp.sizeNormal
}

// GetSizeExtended returns sizeExtended of DialectMessageCT
func (mp *DialectMessageCT) GetSizeExtended() byte {
	return mp.sizeExtended
}

// GetCRCExtra returns crcExtra of DialectMessageCT
func (mp *DialectMessageCT) GetCRCExtra() byte {
	return mp.crcExtra
}

// DEFINE PRIVATE STATIC FUNCTIONS.

func newDialectMessage(msg Message) (*DialectMessageCT, error) {
	mp := &DialectMessageCT{}
	mp.elemType = reflect.TypeOf(msg).Elem()

	mp.Fields = make([]*DialectMessageField, mp.elemType.NumField())

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
		var dialectType DialectFieldType

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

			dialectType = DialectFieldTypeFromGo[tagEnum]
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
			dialectType = DialectFieldTypeFromGo[goType.Name()]
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
			size = DialectFieldTypeSizes[dialectType] * arrayLength
		} else {
			size = DialectFieldTypeSizes[dialectType]
		}

		mp.Fields[i] = &DialectMessageField{
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
	sort.Slice(mp.Fields, func(i, j int) bool {
		// sort by weight if not extension
		if mp.Fields[i].isExtension == false && mp.Fields[j].isExtension == false {
			if w1, w2 := DialectFieldTypeSizes[mp.Fields[i].ftype], DialectFieldTypeSizes[mp.Fields[j].ftype]; w1 != w2 {
				return w1 > w2
			}
		}
		// sort by original index
		return mp.Fields[i].index < mp.Fields[j].index
	})

	// generate CRC extra
	// https://mavlink.io/en/guide/serialization.html#crc_extra
	mp.crcExtra = func() byte {
		h := NewX25()
		h.Write([]byte(msgName + " "))

		for _, f := range mp.Fields {
			// skip extensions
			if f.isExtension == true {
				continue
			}

			h.Write([]byte(DialectFieldTypeString[f.ftype] + " "))
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

func (mp *DialectMessageCT) newMsg() Message {
	ref := reflect.New(mp.elemType)
	msg := ref.Interface().(Message)
	return msg
}

func (mp *DialectMessageCT) getFields() []*DialectMessageField {
	return mp.Fields
}

func (mp *DialectMessageCT) getCRCExtra() byte {
	return mp.crcExtra
}

func (mp *DialectMessageCT) decode(buf []byte, isFrameV2 bool) (Message, error) {
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
	for _, f := range mp.Fields {
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

func (mp *DialectMessageCT) encode(msg Message, isFrameV2 bool) ([]byte, error) {
	var buf []byte

	if isFrameV2 == true {
		buf = make([]byte, mp.sizeExtended)
	} else {
		buf = make([]byte, mp.sizeNormal)
	}

	start := buf

	// encode field by field
	for _, f := range mp.Fields {
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

// Decode is a public function necessary for testing
func (mp *DialectMessageCT) Decode(buf []byte, isFrameV2 bool) (Message, error) {
	return mp.decode(buf, isFrameV2)
}

// Encode is a public function necessary for testing
func (mp *DialectMessageCT) Encode(msg Message, isFrameV2 bool) ([]byte, error) {
	return mp.encode(msg, isFrameV2)
}

// ALL DONE.
