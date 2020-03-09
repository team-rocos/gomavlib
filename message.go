package gomavlib

import (
	"errors"
	"reflect"
)

// Message is the interface that all Mavlink messages must implements.
// Furthermore, message structs must be labeled MessageNameOfMessage.
type Message interface {
	GetId() uint32
	SetField(field string, value interface{}) error
}

// MessageRaw is a special struct that contains a byte-encoded message,
// available in Content. It is used:
//
// * as intermediate step in the encoding/decoding process
//
// * when the parser receives an unknown message
//
type MessageRaw struct {
	Id      uint32
	Content []byte
}

// GetId implements the message interface.
func (m *MessageRaw) GetId() uint32 {
	return m.Id
}

func (m *MessageRaw) SetField(field string, value interface{}) error {
	return errors.New("cannot set fields in raw message")
}

func SetMessageField(m Message, field string, value interface{}) error {
	switch reflect.ValueOf(value).Kind() {
	case reflect.Int:
		reflect.ValueOf(m).Elem().FieldByName(field).SetInt(int64(value.(int)))
	case reflect.Int8:
		reflect.ValueOf(m).Elem().FieldByName(field).SetInt(int64(value.(int8)))
	case reflect.Int16:
		reflect.ValueOf(m).Elem().FieldByName(field).SetInt(int64(value.(int16)))
	case reflect.Int32:
		reflect.ValueOf(m).Elem().FieldByName(field).SetInt(int64(value.(int32)))
	case reflect.Int64:
		reflect.ValueOf(m).Elem().FieldByName(field).SetInt(int64(value.(int64)))
	case reflect.Uint:
		reflect.ValueOf(m).Elem().FieldByName(field).SetUint(uint64(value.(uint)))
	case reflect.Uint8:
		reflect.ValueOf(m).Elem().FieldByName(field).SetUint(uint64(value.(uint8)))
	case reflect.Uint16:
		reflect.ValueOf(m).Elem().FieldByName(field).SetUint(uint64(value.(uint16)))
	case reflect.Uint32:
		reflect.ValueOf(m).Elem().FieldByName(field).SetUint(uint64(value.(uint32)))
	case reflect.Uint64:
		reflect.ValueOf(m).Elem().FieldByName(field).SetUint(uint64(value.(uint64)))
	case reflect.Float32:
		reflect.ValueOf(m).Elem().FieldByName(field).SetFloat(float64(value.(float32)))
	case reflect.Float64:
		reflect.ValueOf(m).Elem().FieldByName(field).SetFloat(float64(value.(float64)))
	default:
		reflect.ValueOf(m).Elem().FieldByName(field).Set(reflect.ValueOf(value))
	}
	return nil
}
