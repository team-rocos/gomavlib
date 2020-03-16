package gomavlib

import (
	"encoding/binary"
	"math"
	"reflect"
	"regexp"
	"strings"
)

type dialectFieldType int

const (
	typeDouble dialectFieldType = iota + 1
	typeUint64
	typeInt64
	typeFloat
	typeUint32
	typeInt32
	typeUint16
	typeInt16
	typeUint8
	typeInt8
	typeChar
)

var dialectFieldTypeFromGo = map[string]dialectFieldType{
	"float64": typeDouble,
	"uint64":  typeUint64,
	"int64":   typeInt64,
	"float32": typeFloat,
	"uint32":  typeUint32,
	"int32":   typeInt32,
	"uint16":  typeUint16,
	"int16":   typeInt16,
	"uint8":   typeUint8,
	"int8":    typeInt8,
	"string":  typeChar,
}

var dialectFieldTypeString = map[dialectFieldType]string{
	typeDouble: "double",
	typeUint64: "uint64_t",
	typeInt64:  "int64_t",
	typeFloat:  "float",
	typeUint32: "uint32_t",
	typeInt32:  "int32_t",
	typeUint16: "uint16_t",
	typeInt16:  "int16_t",
	typeUint8:  "uint8_t",
	typeInt8:   "int8_t",
	typeChar:   "char",
}

var dialectFieldTypeSizes = map[dialectFieldType]byte{
	typeDouble: 8,
	typeUint64: 8,
	typeInt64:  8,
	typeFloat:  4,
	typeUint32: 4,
	typeInt32:  4,
	typeUint16: 2,
	typeInt16:  2,
	typeUint8:  1,
	typeInt8:   1,
	typeChar:   1,
}

func dialectFieldGoToDef(in string) string {
	re := regexp.MustCompile("([A-Z])")
	in = re.ReplaceAllString(in, "_${1}")
	return strings.ToLower(in[1:])
}

func dialectMsgGoToDef(in string) string {
	re := regexp.MustCompile("([A-Z])")
	in = re.ReplaceAllString(in, "_${1}")
	return strings.ToUpper(in[1:])
}

// Dialect : Interface
type Dialect interface {
	getVersion() uint
	getMsgById(id uint32) (*dialectMessage, bool)
}

type dialectMessageField struct {
	isEnum      bool
	ftype       dialectFieldType
	name        string
	arrayLength byte
	index       int
	isExtension bool
}

type dialectMessage interface {
	newMsg() *Message
	getFields() []*dialectMessageField
	getCRCExtra() byte
	decode(buf []byte, isFrameV2 bool) (Message, error)
	encode(msg Message, isFrameV2 bool) ([]byte, error)
}

func decodeValue(target reflect.Value, buf []byte, f *dialectMessageField) int {
	if f.isEnum == true {
		switch f.ftype {
		case typeUint8:
			target.SetInt(int64(buf[0]))
			return 1

		case typeUint16:
			target.SetInt(int64(binary.LittleEndian.Uint16(buf)))
			return 2

		case typeUint32:
			target.SetInt(int64(binary.LittleEndian.Uint32(buf)))
			return 4

		case typeInt32:
			target.SetInt(int64(binary.LittleEndian.Uint32(buf)))
			return 4

		case typeUint64:
			target.SetInt(int64(binary.LittleEndian.Uint64(buf)))
			return 8

		default:
			panic("unexpected type")
		}
	}

	switch tt := target.Addr().Interface().(type) {
	case *string:
		// find nil character or string end
		end := 0
		for buf[end] != 0 && end < int(f.arrayLength) {
			end++
		}
		*tt = string(buf[:end])
		return int(f.arrayLength) // return length including zeros

	case *int8:
		*tt = int8(buf[0])
		return 1

	case *uint8:
		*tt = buf[0]
		return 1

	case *int16:
		*tt = int16(binary.LittleEndian.Uint16(buf))
		return 2

	case *uint16:
		*tt = binary.LittleEndian.Uint16(buf)
		return 2

	case *int32:
		*tt = int32(binary.LittleEndian.Uint32(buf))
		return 4

	case *uint32:
		*tt = binary.LittleEndian.Uint32(buf)
		return 4

	case *int64:
		*tt = int64(binary.LittleEndian.Uint64(buf))
		return 8

	case *uint64:
		*tt = binary.LittleEndian.Uint64(buf)
		return 8

	case *float32:
		*tt = math.Float32frombits(binary.LittleEndian.Uint32(buf))
		return 4

	case *float64:
		*tt = math.Float64frombits(binary.LittleEndian.Uint64(buf))
		return 8

	default:
		panic("unexpected type")
	}
}

func encodeValue(buf []byte, target reflect.Value, f *dialectMessageField) int {
	if f.isEnum == true {
		switch f.ftype {
		case typeUint8:
			buf[0] = byte(target.Int())
			return 1

		case typeUint16:
			binary.LittleEndian.PutUint16(buf, uint16(target.Int()))
			return 2

		case typeUint32:
			binary.LittleEndian.PutUint32(buf, uint32(target.Int()))
			return 4

		case typeInt32:
			binary.LittleEndian.PutUint32(buf, uint32(target.Int()))
			return 4

		case typeUint64:
			binary.LittleEndian.PutUint64(buf, uint64(target.Int()))
			return 8

		default:
			panic("unexpected type")
		}
	}

	switch tt := target.Addr().Interface().(type) {
	case *string:
		copy(buf[:f.arrayLength], *tt)
		return int(f.arrayLength) // return length including zeros

	case *int8:
		buf[0] = uint8(*tt)
		return 1

	case *uint8:
		buf[0] = *tt
		return 1

	case *int16:
		binary.LittleEndian.PutUint16(buf, uint16(*tt))
		return 2

	case *uint16:
		binary.LittleEndian.PutUint16(buf, *tt)
		return 2

	case *int32:
		binary.LittleEndian.PutUint32(buf, uint32(*tt))
		return 4

	case *uint32:
		binary.LittleEndian.PutUint32(buf, *tt)
		return 4

	case *int64:
		binary.LittleEndian.PutUint64(buf, uint64(*tt))
		return 8

	case *uint64:
		binary.LittleEndian.PutUint64(buf, *tt)
		return 8

	case *float32:
		binary.LittleEndian.PutUint32(buf, math.Float32bits(*tt))
		return 4

	case *float64:
		binary.LittleEndian.PutUint64(buf, math.Float64bits(*tt))
		return 8

	default:
		panic("unexpected type")
	}
}
