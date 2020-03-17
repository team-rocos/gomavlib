package gomavlib

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"

	"github.com/pkg/errors"
	libgen "github.com/team-rocos/gomavlib/commands/dialgen/libgen"
)

// DEFINE PUBLIC TYPES AND STRUCTURES.

// DialectRT contains available messages and the configuration needed to encode and decode them.
type DialectRT struct {
	version  uint
	defs     []*libgen.OutDefinition
	messages map[uint32]*dialectMessageRT
}

// Assert to check we're implementing the interfaces we expect to be.
var _ Dialect = &DialectRT{}

type DynamicMessage struct {
	t      *dialectMessageRT
	Fields map[string]interface{}
}

// Assert to check we're implementing the interfaces we expect to be.
var _ Message = &DynamicMessage{}

// DEFINE PRIVATE TYPES AND STRUCTURES.

type dialectMessageRT struct {
	dialect      *DialectRT
	msg          *libgen.OutMessage
	crcExtra     byte
	sizeNormal   uint
	sizeExtended uint
}

// Assert to check we're implementing the interfaces we expect to be.
var _ dialectMessage = &dialectMessageRT{}

// DEFINE PUBLIC STATIC FUNCTIONS.

// NewDialectRT allocates a Dialect.
func NewDialectRT(version uint, outDefs []*libgen.OutDefinition) (*DialectRT, error) {
	// Create the new dialect object.
	d := &DialectRT{
		version:  version,
		defs:     outDefs,
		messages: make(map[uint32]*dialectMessageRT),
	}

	// Populate the set of messages in the dialect by instantiating a dialectMessage from each message entry in the definitions.
	for _, def := range d.defs {
		for _, msg := range def.Messages {

			// Reorder fields as described in https://mavlink.io/en/guide/serialization.html#field_reordering
			sort.Slice(msg.Fields, func(i, j int) bool {
				// sort by weight if not extension
				if msg.Fields[i].IsExtension == false && msg.Fields[j].IsExtension == false {
					if w1, w2 := dialectFieldTypeSizes[dialectFieldTypeFromGo[msg.Fields[i].Type]], dialectFieldTypeSizes[dialectFieldTypeFromGo[msg.Fields[j].Type]]; w1 != w2 {
						return w1 > w2
					}
				}
				// sort by original index
				return msg.Fields[i].Index < msg.Fields[j].Index
			})

			// Work out what the CRC-extra value should be: https://mavlink.io/en/guide/serialization.html#crc_extra
			var crcExtra byte = func() byte {
				h := NewX25()
				h.Write([]byte(msg.Name + " ")) // TODO - This might be broken, because the "name" field might have been converted into a 'Go-like' name, whereas the CRC calc probably needs to use the original MAVLink-like form.

				for _, f := range msg.Fields {
					// skip extensions
					if f.IsExtension == true {
						continue
					}

					h.Write([]byte(dialectFieldTypeString[dialectFieldTypeFromGo[f.Type]] + " "))
					h.Write([]byte(f.Name + " ")) // TODO - This might be broken, because the "name" field might have been converted into a 'Go-like' name, whereas the CRC calc probably needs to use the original MAVLink-like form.

					// TODO - Need to support arrays, but not sure how that happens: libgen.OutField doesn't seem to contain info about array types?

					//if f.arrayLength > 0 {
					//	h.Write([]byte{f.arrayLength})
					//}
				}
				sum := h.Sum16()
				return byte((sum & 0xFF) ^ (sum >> 8))
			}()

			// Work out how large the message is expected to be.
			sizeNormal := uint(0)
			sizeExtended := uint(0)
			for _, f := range msg.Fields {
				// Work out how big this field will be.
				var size uint = uint(dialectFieldTypeSizes[dialectFieldTypeFromGo[f.Type]])

				// Extension fields count towards towards the extended size, but not the normal size.
				sizeExtended += size
				if f.IsExtension == false {
					sizeNormal += size
				}
			}

			// Create a new dialectMessage capturing the vital statistics.
			d.messages[uint32(msg.Id)] = &dialectMessageRT{
				dialect:      d,
				msg:          msg,
				crcExtra:     crcExtra,
				sizeNormal:   sizeNormal,
				sizeExtended: sizeExtended,
			}
		}
	}

	return d, nil
}

// DEFINE PUBLIC RECEIVER FUNCTIONS.

// DialectRT :: Dialect

func (d *DialectRT) getVersion() uint {
	return d.version
}

func (d *DialectRT) getMsgById(id uint32) (*dialectMessage, bool) {
	var msg dialectMessage
	var ok bool
	msg, ok = d.messages[id]
	return &msg, ok
}

// DynamicMessage :: Message

func (d DynamicMessage) GetId() uint32 {
	// Just look up the Id and return it.
	return uint32(d.t.msg.Id)
}

func (d DynamicMessage) SetField(field string, value interface{}) error {
	// Search through the list of fields that the message is supposed to have.
	var fieldInfo *libgen.OutField
	for _, v := range d.t.msg.Fields {
		if v.Name == field {
			// This is the field we are after, so remember it.
			fieldInfo = v
		}
	}

	// If we didn't find anything, that means this type of message isn't supposed to have a field with that name.
	if fieldInfo == nil {
		return errors.New("invalid field name: " + field)
	}

	// Else, need to check that the object we've been passed is the right type for the matching field.
	switch fieldInfo.Type {
	case "int8":
		// Try to convert the value into an int.
		if v, ok := value.(int8); ok {
			// This is the correct type, so save it into our message.
			d.Fields[field] = v
		} else {
			// The value was the wrong type.
			return errors.New("incorrect type for field: " + field + " - expected int")
		}
	case "uint8":
		// Try to convert the value into an int.
		if v, ok := value.(uint8); ok {
			// This is the correct type, so save it into our message.
			d.Fields[field] = v
		} else {
			// The value was the wrong type.
			return errors.New("incorrect type for field: " + field + " - expected int")
		}
	case "int16":
		// Try to convert the value into an int.
		if v, ok := value.(int16); ok {
			// This is the correct type, so save it into our message.
			d.Fields[field] = v
		} else {
			// The value was the wrong type.
			return errors.New("incorrect type for field: " + field + " - expected int")
		}
	case "uint16":
		// Try to convert the value into an int.
		if v, ok := value.(uint16); ok {
			// This is the correct type, so save it into our message.
			d.Fields[field] = v
		} else {
			// The value was the wrong type.
			return errors.New("incorrect type for field: " + field + " - expected int")
		}
	case "int32":
		// Try to convert the value into an int.
		if v, ok := value.(int32); ok {
			// This is the correct type, so save it into our message.
			d.Fields[field] = v
		} else {
			// The value was the wrong type.
			return errors.New("incorrect type for field: " + field + " - expected int")
		}
	case "uint32":
		// Try to convert the value into an int.
		if v, ok := value.(uint32); ok {
			// This is the correct type, so save it into our message.
			d.Fields[field] = v
		} else {
			// The value was the wrong type.
			return errors.New("incorrect type for field: " + field + " - expected int")
		}
	case "int64":
		// Try to convert the value into an int.
		if v, ok := value.(int64); ok {
			// This is the correct type, so save it into our message.
			d.Fields[field] = v
		} else {
			// The value was the wrong type.
			return errors.New("incorrect type for field: " + field + " - expected int")
		}
	case "uint64":
		// Try to convert the value into an int.
		if v, ok := value.(uint64); ok {
			// This is the correct type, so save it into our message.
			d.Fields[field] = v
		} else {
			// The value was the wrong type.
			return errors.New("incorrect type for field: " + field + " - expected int")
		}
	// TODO - All the other cases.
	default:
		return errors.New("unsupported field type in dynamic MAVLink message")
	}

	// If we make it here, should mean everything went ok.
	return nil
}

// DEFINE PRIVATE STATIC FUNCTIONS.

// DEFINE PRIVATE RECEIVER FUNCTIONS.

func (mp *dialectMessageRT) newMsg() Message {
	// Just make an empty dynamic message which points at this dialectMessage as its parent.
	msg := &DynamicMessage{
		t:      mp,
		Fields: make(map[string]interface{}),
	}
	return msg
}

func (mp *dialectMessageRT) getFields() []*dialectMessageField {
	// Iterate over each of the field definitions and construct a new dialectMessageField representation for each.
	fields := make([]*dialectMessageField, 0)
	for _, f := range mp.msg.Fields {
		ftype := dialectFieldTypeFromGo[f.Type]
		fields = append(fields, &dialectMessageField{
			isEnum:      f.IsEnum,
			ftype:       ftype,
			name:        f.Name,
			arrayLength: dialectFieldTypeSizes[ftype],
			index:       f.Index,
			isExtension: f.IsExtension,
		})
	}
	return fields
}

func (mp *dialectMessageRT) getCRCExtra() byte {
	return mp.crcExtra
}

func (mp *dialectMessageRT) decode(buf []byte, isFrameV2 bool) (Message, error) {
	// Insert any required padding.
	if isFrameV2 == true {
		// In V2 buffer length can be > message or < message.  In this latter case it must be filled with zeros to support empty-byte de-truncation and extension fields.
		if len(buf) < int(mp.sizeExtended) {
			buf = append(buf, bytes.Repeat([]byte{0x00}, int(mp.sizeExtended)-len(buf))...)
		}
	} else {
		// But in V1 buffer must fit message perfectly.
		if len(buf) != int(mp.sizeNormal) {
			return nil, fmt.Errorf("unexpected size (%d vs %d)", len(buf), mp.sizeNormal)
		}
	}

	// Convert the bytes into a consumable buffer to read from.
	b := bytes.NewBuffer(buf)

	// Create the dynamic message which we're gonna fill up.
	dm := mp.newMsg()

	// Decode field by field.
	for _, fieldDef := range mp.msg.Fields { // TODO - Hmm, this implies that mp.msg.Fields is in order?  In that case, why do we need orderedFields?
		// Skip extensions in V1 frames.
		if isFrameV2 == false && fieldDef.IsExtension == true {
			continue
		}

		// Need to handle each type of field separately.
		switch fieldDef.Type {
		case "int8":
			var val int8
			if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
				return nil, errors.Wrap(err, "failed to read field: "+fieldDef.Name+" : ")
			}
			if err := dm.SetField(fieldDef.Name, val); err != nil {
				return nil, errors.Wrap(err, "failed to set field: "+fieldDef.Name+" : ")
			}
		case "uint8":
			var val uint8
			if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
				return nil, errors.Wrap(err, "failed to read field: "+fieldDef.Name+" : ")
			}
			if err := dm.SetField(fieldDef.Name, val); err != nil {
				return nil, errors.Wrap(err, "failed to set field: "+fieldDef.Name+" : ")
			}
		case "int16":
			var val int16
			if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
				return nil, errors.Wrap(err, "failed to read field: "+fieldDef.Name+" : ")
			}
			if err := dm.SetField(fieldDef.Name, val); err != nil {
				return nil, errors.Wrap(err, "failed to set field: "+fieldDef.Name+" : ")
			}
		case "uint16":
			var val uint16
			if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
				return nil, errors.Wrap(err, "failed to read field: "+fieldDef.Name+" : ")
			}
			if err := dm.SetField(fieldDef.Name, val); err != nil {
				return nil, errors.Wrap(err, "failed to set field: "+fieldDef.Name+" : ")
			}
		case "int32":
			var val int32
			if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
				return nil, errors.Wrap(err, "failed to read field: "+fieldDef.Name+" : ")
			}
			if err := dm.SetField(fieldDef.Name, val); err != nil {
				return nil, errors.Wrap(err, "failed to set field: "+fieldDef.Name+" : ")
			}
		case "uint32":
			var val uint32
			if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
				return nil, errors.Wrap(err, "failed to read field: "+fieldDef.Name+" : ")
			}
			if err := dm.SetField(fieldDef.Name, val); err != nil {
				return nil, errors.Wrap(err, "failed to set field: "+fieldDef.Name+" : ")
			}
		case "int64":
			var val int64
			if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
				return nil, errors.Wrap(err, "failed to read field: "+fieldDef.Name+" : ")
			}
			if err := dm.SetField(fieldDef.Name, val); err != nil {
				return nil, errors.Wrap(err, "failed to set field: "+fieldDef.Name+" : ")
			}
		case "uint64":
			var val uint64
			if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
				return nil, errors.Wrap(err, "failed to read field: "+fieldDef.Name+" : ")
			}
			if err := dm.SetField(fieldDef.Name, val); err != nil {
				return nil, errors.Wrap(err, "failed to set field: "+fieldDef.Name+" : ")
			}
		// TODO - Support the other types.
		default:
			// We don't know what to do with this type.
			return nil, errors.New("unsupported field type: " + fieldDef.Name)
		}
	}

	// All done.
	return dm, nil
}

func (mp *dialectMessageRT) encode(msg Message, isFrameV2 bool) ([]byte, error) {
	// Make sure the message we're encoding matches the type of the dialectMessage being used to do the encoding.
	var dm DynamicMessage
	var ok bool
	if dm, ok = msg.(DynamicMessage); !ok {
		return nil, errors.New("message was not a DynamicMessage")
	}
	if dm.t != mp {
		return nil, errors.New("wrong DynamicMessage type")
	}

	// We're filling a buffer byte by byte.
	buf := &bytes.Buffer{}

	// Iterate over the definitions in the wire order.
	for _, fieldDef := range mp.msg.Fields {
		// Need to handle each type of field separately.
		switch fieldDef.Type {
		case "int8":
			// Look up the actual value for this field.
			var val int8
			if v, ok := dm.Fields[fieldDef.Name]; ok {
				if val, ok = v.(int8); !ok {
					// The value stored for this field wasn't the right type.
					return nil, errors.New("invalid value for field: " + fieldDef.Name)
				}
			} // Else just use the default value.
			binary.Write(buf, binary.LittleEndian, val)
		case "uint8":
			// Look up the actual value for this field.
			var val uint8
			if v, ok := dm.Fields[fieldDef.Name]; ok {
				if val, ok = v.(uint8); !ok {
					// The value stored for this field wasn't the right type.
					return nil, errors.New("invalid value for field: " + fieldDef.Name)
				}
			} // Else just use the default value.
			binary.Write(buf, binary.LittleEndian, val)
		case "int16":
			// Look up the actual value for this field.
			var val int16
			if v, ok := dm.Fields[fieldDef.Name]; ok {
				if val, ok = v.(int16); !ok {
					// The value stored for this field wasn't the right type.
					return nil, errors.New("invalid value for field: " + fieldDef.Name)
				}
			} // Else just use the default value.
			binary.Write(buf, binary.LittleEndian, val)
		case "uint16":
			// Look up the actual value for this field.
			var val uint16
			if v, ok := dm.Fields[fieldDef.Name]; ok {
				if val, ok = v.(uint16); !ok {
					// The value stored for this field wasn't the right type.
					return nil, errors.New("invalid value for field: " + fieldDef.Name)
				}
			} // Else just use the default value.
			binary.Write(buf, binary.LittleEndian, val)
		case "int32":
			// Look up the actual value for this field.
			var val int32
			if v, ok := dm.Fields[fieldDef.Name]; ok {
				if val, ok = v.(int32); !ok {
					// The value stored for this field wasn't the right type.
					return nil, errors.New("invalid value for field: " + fieldDef.Name)
				}
			} // Else just use the default value.
			binary.Write(buf, binary.LittleEndian, val)
		case "uint32":
			// Look up the actual value for this field.
			var val uint32
			if v, ok := dm.Fields[fieldDef.Name]; ok {
				if val, ok = v.(uint32); !ok {
					// The value stored for this field wasn't the right type.
					return nil, errors.New("invalid value for field: " + fieldDef.Name)
				}
			} // Else just use the default value.
			binary.Write(buf, binary.LittleEndian, val)
		case "int64":
			// Look up the actual value for this field.
			var val int64
			if v, ok := dm.Fields[fieldDef.Name]; ok {
				if val, ok = v.(int64); !ok {
					// The value stored for this field wasn't the right type.
					return nil, errors.New("invalid value for field: " + fieldDef.Name)
				}
			} // Else just use the default value.
			binary.Write(buf, binary.LittleEndian, val)
		case "uint64":
			// Look up the actual value for this field.
			var val uint64
			if v, ok := dm.Fields[fieldDef.Name]; ok {
				if val, ok = v.(uint64); !ok {
					// The value stored for this field wasn't the right type.
					return nil, errors.New("invalid value for field: " + fieldDef.Name)
				}
			} // Else just use the default value.
			binary.Write(buf, binary.LittleEndian, val)
		// TODO - Support the other types.
		default:
			// We don't know what to do with this type.
			return nil, errors.New("unsupported field type: " + fieldDef.Name)
		}
	}

	// All done.
	return buf.Bytes(), nil
}

// ALL DONE.
