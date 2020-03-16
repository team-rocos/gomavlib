package gomavlib

import (
	"fmt"
	"reflect"

	libgen "github.com/team-rocos/gomavlib/commands/dialgen/libgen"
)

// type dialectMessage interface {
// 	newMsg() *Message
// 	getFields() []*dialectMessageField
// 	getCRCExtra() byte
// 	decode(buf []byte, isFrameV2 bool) (Message, error)
// 	encode(msg Message, isFrameV2 bool) ([]byte, error)
// }

// TestingFunctions :
func TestingFunctions(d *DialectRT) {
	version := d.getVersion()
	fmt.Println("Version: ", version)

	message, ok := d.getMsgById(300)
	if !ok {
		fmt.Println("Error: message by this ID not found!")
	} else {
		fmt.Println(*message)
	}

	// var dm *dialectMessageRT
	// message.
	fields := (*message).getFields()
	fmt.Println(fields)
}

// DEFINE PUBLIC TYPES AND STRUCTURES.

// DialectRT contains available messages and the configuration needed to encode and decode them.
type DialectRT struct {
	version uint
	//outDefs []*libgen.OutDefinition
	messages map[uint32]*dialectMessageRT
}

// DynamicMessage :
type DynamicMessage struct {
	t      *dialectMessageRT
	fields map[string]interface{}
}

// GetId() : Required to meet requirements for Message Interface
func (d *DynamicMessage) GetId() uint32 {
	return uint32(d.t.msg.Id)
}

// SetField : Required to meet requirements for Message Interface
func (d *DynamicMessage) SetField(field string, value interface{}) error {
	return nil ////////////////////////////////////////////////////////////////////////////////// Ask about this???????????????????????/
}

// Assert to check we're implementing the interfaces we expect to be.
var _ Dialect = &DialectRT{}

// DEFINE PRIVATE TYPES AND STRUCTURES.
type dialectMessageRT struct {
	msg libgen.OutMessage
}

// Assert to check we're implementing the interfaces we expect to be.
var _ dialectMessage = &dialectMessageRT{}

// DEFINE PUBLIC STATUC FUNCTIONS.

// NewDialectRT allocates a Dialect.
func NewDialectRT(version uint, outDefs []*libgen.OutDefinition) (*DialectRT, error) {
	// TODO - This.
	d := &DialectRT{
		version:  version,
		messages: make(map[uint32]*dialectMessageRT),
	}

	for _, def := range outDefs {
		for _, msg := range def.Messages {
			d.messages[uint32(msg.Id)] = &dialectMessageRT{
				msg: *msg,
			}
		}
	}

	return d, nil
}

// DEFINE PUBLIC RECEIVER FUNCTIONS.
//
func (d *DialectRT) getVersion() uint {
	// TODO - This.
	return d.version
}

func (d *DialectRT) getMsgById(id uint32) (*dialectMessage, bool) {
	var msg dialectMessage
	var ok bool
	msg, ok = d.messages[id]
	return &msg, ok
}

// DEFINE PRIVATE STATIC FUNCTIONS.
// DEFINE PRIVATE RECEIVER FUNCTIONS.
func (mp *dialectMessageRT) newMsg() *Message {

	// msg := &DynamicMessage{
	// 	t:      mp,
	// 	fields: make(map[string]interface{}),
	// }
	// return msg

	// TODO - This.
	var fakeAnswer *Message
	return fakeAnswer
	// msg := mp.Interface().(Message)
	// return &msg
}

func (mp *dialectMessageRT) getFields() []*dialectMessageField {
	// TODO - This.
	//fieldInfo := mp.d.msgs[mp.id].fields[]
	var output []*dialectMessageField
	fields := mp.msg.Fields
	for _, f := range fields {
		ftype := dialectFieldTypeFromGo[f.Type]
		output = append(output, &dialectMessageField{
			isEnum:      f.IsEnum,
			ftype:       ftype,
			name:        f.Name,
			arrayLength: dialectFieldTypeSizes[ftype],
			index:       f.Index,
			isExtension: f.IsExtension,
		})
	}
	return output
}

func (mp *dialectMessageRT) getCRCExtra() byte {
	// TODO - This.
	var fakeAnswer byte
	return fakeAnswer
}

func (mp *dialectMessageRT) decode(buf []byte, isFrameV2 bool) (Message, error) {
	// TODO - This.
	// var fakeAnswer Message
	// return fakeAnswer, nil

	//msg := reflect.New(mp.elemType)

	// if isFrameV2 == true {
	// 	// in V2 buffer length can be > message or < message
	// 	// in this latter case it must be filled with zeros to support empty-byte de-truncation
	// 	// and extension fields
	// 	if len(buf) < int(mp.sizeExtended) {
	// 		buf = append(buf, bytes.Repeat([]byte{0x00}, int(mp.sizeExtended)-len(buf))...)
	// 	}

	// } else {
	// in V1 buffer must fit message perfectly
	// if len(buf) != int(mp.sizeNormal) {
	// 	return nil, fmt.Errorf("unexpected size (%d vs %d)", len(buf), mp.sizeNormal)
	// }
	//}

	// decode field by field
	for _, f := range mp.msg.Fields {
		// skip extensions in V1 frames
		if isFrameV2 == false && f.IsExtension == true {
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

func (mp *dialectMessageRT) encode(msg Message, isFrameV2 bool) ([]byte, error) {
	// TODO - This.
	// if dm, ok := msg.(DynamicMessage); !ok {
	// 	return nil, errors.New("oh noes")
	// }
	// if dm.t != mp {
	// 	panic()
	// }
	// for i := 0; i < len(mp.d.msgs[mp.id].fields)
	// 	typeOfField := mp.d.msgs[mp.id].fields[i]
	// 	case tyepof Fieldconst
	// 	byte:
	// 		valueOfField := dm.fields[i].(byte)
	// 		buf = append(buf, valucof field)
	var fakeAnswer []byte
	return fakeAnswer, nil
}

// ALL DONE.
