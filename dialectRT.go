package gomavlib

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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

// DynamicMessage : Used for RT message generation
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
				h.Write([]byte(msg.OriginalName + " ")) // Using 'OriginalName' to ensure original Mavlink style name used

				for _, f := range msg.Fields {
					// skip extensions
					if f.IsExtension == true {
						continue
					}

					h.Write([]byte(dialectFieldTypeString[dialectFieldTypeFromGo[f.Type]] + " "))

					h.Write([]byte(f.OriginalName + " ")) // Using 'OriginalName' to ensure original Mavlink style name used
					if f.ArrayLength > 0 {
						h.Write([]byte{byte(f.ArrayLength)})
					}
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

				if f.ArrayLength > 0 {
					size = size * uint(f.ArrayLength)
				}

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

// GetId : DynamicMessage function to meet Message interface requirements
func (d DynamicMessage) GetId() uint32 {
	// Just look up the Id and return it.
	return uint32(d.t.msg.Id)
}

// SetField : DynamicMessage function to meet Message interface requirements
func (d DynamicMessage) SetField(field string, value interface{}) error {
	// Search through the list of fields that the message is supposed to have.
	var fieldInfo *libgen.OutField
	for _, v := range d.t.msg.Fields {
		if v.Name == field {
			// This is the field we are after, so remember it.
			fieldInfo = v
			break
		}
	}

	// If we didn't find anything, that means this type of message isn't supposed to have a field with that name.
	if fieldInfo == nil {
		return errors.New("invalid field name: " + field)
	}

	// Else, need to check that the object we've been passed is the right type for the matching field.
	switch fieldInfo.Type {
	case "int8":
		// Try to convert the value into an int8.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]int8); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []int8")
			}
		} else {
			if v, ok := value.(int8); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected int8")
			}
		}
	case "uint8":
		// Try to convert the value into an uint8.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]uint8); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []uint8")
			}
		} else {
			if v, ok := value.(uint8); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected uint8")
			}
		}
	case "int16":
		// Try to convert the value into an int16.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]int16); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []int16")
			}
		} else {
			if v, ok := value.(int16); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected int16")
			}
		}
	case "uint16":
		// Try to convert the value into an uint16.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]uint16); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []uint16")
			}
		} else {
			if v, ok := value.(uint16); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected uint16")
			}
		}
	case "int32":
		// Try to convert the value into an int32.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]int32); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []int32")
			}
		} else {
			if v, ok := value.(int32); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected int32")
			}
		}
	case "uint32":
		// Try to convert the value into an uint32.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]uint32); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []uint32")
			}
		} else {
			if v, ok := value.(uint32); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected uint32")
			}
		}
	case "int64":
		// Try to convert the value into an int64.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]int64); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []int64")
			}
		} else {
			if v, ok := value.(int64); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected int64")
			}
		}
	case "uint64":
		// Try to convert the value into an uint64.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]uint64); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []uint64")
			}
		} else {
			if v, ok := value.(uint64); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected uint64")
			}
		}
	case "float64":
		// Try to convert the value into an float64.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]float64); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []float64")
			}
		} else {
			if v, ok := value.(float64); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected float64")
			}
		}
	case "float32":
		// Try to convert the value into an float32.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]float32); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []float32")
			}
		} else {
			if v, ok := value.(float32); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected float32")
			}
		}
	case "string":
		// Try to convert the value into a string.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]string); ok {
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []string")
			}
		} else {
			if v, ok := value.(string); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected string")
			}
		}
	default:
		return errors.New("unsupported field type in dynamic MAVLink message")
	}

	// If we make it here, should mean everything went ok.
	return nil
}

// GetName : returns OriginalName (in mavlink format)
func (d DynamicMessage) GetName() string {
	return d.t.msg.OriginalName
}

// JSON Methods
const (
	//Sep is a namespace separator string
	Sep = "/"
	//GlobalNS is the global namespace initial separator string
	GlobalNS = "/"
	//PrivateNS is private namespace initial separator string
	PrivateNS = "~"
)

// GenerateJSONSchema generates a (primitive) JSON schema for the associated DynamicMessageType; however note that since
// we are mostly interested in making schema's for particular _topics_, the function takes a string prefix, and string topic name, which are
// used to id the resulting schema.
func (mp *dialectMessageRT) GenerateJSONSchema(prefix string, topic string) ([]byte, error) {
	// The JSON schema for a message consist of the (recursive) properties names/types:
	schemaItems, err := mp.generateJSONSchemaProperties(prefix + topic)
	if err != nil {
		return nil, err
	}

	// Plus some extra keywords:
	schemaItems["$schema"] = "https://json-schema.org/draft-07/schema#"
	schemaItems["$id"] = prefix + topic

	// The schema itself is created from the map of properties.
	schemaString, err := json.Marshal(schemaItems)
	if err != nil {
		return nil, err
	}

	// All done.
	return schemaString, errors.New("GenerateJSONSchema not yet written")
}

func (mp *dialectMessageRT) generateJSONSchemaProperties(topic string) (map[string]interface{}, error) {
	// // Each message's schema indicates that it is an 'object' with some nested properties: those properties are the fields and their types.
	// properties := make(map[string]interface{})
	schemaItems := make(map[string]interface{})
	// schemaItems["type"] = "object"
	// schemaItems["title"] = topic
	// schemaItems["properties"] = properties

	// // Iterate over each of the fields in the message.
	// for _, field := range mp.msg.Fields {
	// 	if field.ArrayLength != 0 {
	// 		// It's an array.
	// 		propertyContent := make(map[string]interface{})
	// 		properties[field.Name] = propertyContent

	// 		if field.Type == "uint8" {
	// 			propertyContent["title"] = topic + Sep + field.Name
	// 			propertyContent["type"] = "string"
	// 		} else {
	// 			// Arrays all have a type of 'array', regardless of that the hold, then the 'item' keyword determines what type goes in the array.
	// 			propertyContent["type"] = "array"
	// 			propertyContent["title"] = topic + Sep + field.Name
	// 			arrayItems := make(map[string]interface{})
	// 			propertyContent["items"] = arrayItems

	// 			// Need to handle each type appropriately.
	// 			if field.Type == "string" {
	// 				arrayItems["type"] = "string"
	// 			} else if field.Type == "time" {
	// 				timeItems := make(map[string]interface{})
	// 				timeItems["sec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "sec"}
	// 				timeItems["nsec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "nsec"}
	// 				arrayItems["type"] = "object"
	// 				arrayItems["properties"] = timeItems
	// 			} else if field.Type == "duration" {
	// 				timeItems := make(map[string]interface{})
	// 				timeItems["sec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "sec"}
	// 				timeItems["nsec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "nsec"}
	// 				arrayItems["type"] = "object"
	// 				arrayItems["properties"] = timeItems
	// 			} else {
	// 				// It's a primitive.
	// 				var jsonType string
	// 				if field.Type == "int8" || field.Type == "int16" || field.Type == "uint16" ||
	// 					field.Type == "int32" || field.Type == "uint32" || field.Type == "int64" || field.Type == "uint64" {
	// 					jsonType = "integer"
	// 				} else if field.Type == "float32" || field.Type == "float64" {
	// 					jsonType = "number"
	// 				} else if field.Type == "bool" {
	// 					jsonType = "bool"
	// 				} else {
	// 					// Something went wrong.
	// 					return nil, errors.New("we haven't implemented this primitive yet")
	// 				}
	// 				arrayItems["type"] = jsonType
	// 			}
	// 		}
	// 	} else {
	// 		// It's a scalar.
	// 		if field.IsBuiltin {
	// 			propertyContent := make(map[string]interface{})
	// 			properties[field.Name] = propertyContent
	// 			propertyContent["title"] = topic + Sep + field.Name

	// 			if field.Type == "string" {
	// 				propertyContent["type"] = "string"
	// 			} else if field.Type == "time" {
	// 				timeItems := make(map[string]interface{})
	// 				timeItems["sec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "sec"}
	// 				timeItems["nsec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "nsec"}
	// 				propertyContent["type"] = "object"
	// 				propertyContent["properties"] = timeItems
	// 			} else if field.Type == "duration" {
	// 				timeItems := make(map[string]interface{})
	// 				timeItems["sec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "sec"}
	// 				timeItems["nsec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "nsec"}
	// 				propertyContent["type"] = "object"
	// 				propertyContent["properties"] = timeItems
	// 			} else {
	// 				// It's a primitive.
	// 				var jsonType string
	// 				if field.GoType == "int8" || field.GoType == "uint8" || field.GoType == "int16" || field.GoType == "uint16" ||
	// 					field.GoType == "int32" || field.GoType == "uint32" || field.GoType == "int64" || field.GoType == "uint64" {
	// 					jsonType = "integer"
	// 					jsonType = "integer"
	// 					jsonType = "integer"
	// 				} else if field.GoType == "float32" || field.GoType == "float64" {
	// 					jsonType = "number"
	// 				} else if field.GoType == "bool" {
	// 					jsonType = "bool"
	// 				} else {
	// 					// Something went wrong.
	// 					return nil, errors.New("we haven't implemented this primitive yet")
	// 				}
	// 				propertyContent["type"] = jsonType
	// 			}
	// 		} else {
	// 			// It's another nested message.

	// 			// Generate the nested type.
	// 			msgType, err := newDynamicMessageTypeNested(field.Type, field.Package)
	// 			if err != nil {
	// 				return nil, errors.Wrap(err, "Schema Field: "+field.Name)
	// 			}

	// 			// Recursively generate schema information for the nested type.
	// 			schemaElement, err := msgType.generateJSONSchemaProperties(topic + Sep + field.Name)
	// 			if err != nil {
	// 				return nil, errors.Wrap(err, "Schema Field:"+field.Name)
	// 			}
	// 			properties[field.Name] = schemaElement
	// 		}
	// 	}
	// }

	// All done.
	return schemaItems, nil
}

// MarshalJSON provides a custom implementation of JSON marshalling, so that when the DynamicMessage is recursively
// marshalled using the standard Go json.marshal() mechanism, the resulting JSON representation is a compact representation
// of just the important message payload (and not the message definition).  It's important that this representation matches
// the schema generated by GenerateJSONSchema().
func (d *DynamicMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Fields)
}

//UnmarshalJSON provides a custom implementation of JSON unmarshalling. Using the dynamicMessage provided, Msgspec is used to
//determine the individual parsing of each JSON encoded payload item into the correct Go type. It is important each type is
//correct so that the message serializes correctly and is understood by the ROS system
func (d *DynamicMessage) UnmarshalJSON(buf []byte) error {
	return errors.New("UnmarshalJSON function not yet written")
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

		// Need to handle each type of field separately, and handle each type based on whether or not it is an array.
		switch fieldDef.Type {
		case "int8":
			if fieldDef.ArrayLength != 0 {
				var allVals []int8
				var val int8
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val int8
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "uint8":
			if fieldDef.ArrayLength != 0 {
				var allVals []uint8
				var val uint8
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val uint8
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "int16":
			if fieldDef.ArrayLength != 0 {
				var allVals []int16
				var val int16
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val int16
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "uint16":
			if fieldDef.ArrayLength != 0 {
				var allVals []uint16
				var val uint16
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val uint16
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "int32":
			if fieldDef.ArrayLength != 0 {
				var allVals []int32
				var val int32
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val int32
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "uint32":
			if fieldDef.ArrayLength != 0 {
				var allVals []uint32
				var val uint32
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val uint32
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "int64":
			if fieldDef.ArrayLength != 0 {
				var allVals []int64
				var val int64
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val int64
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "uint64":
			if fieldDef.ArrayLength != 0 {
				var allVals []uint64
				var val uint64
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val uint64
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "float64":
			if fieldDef.ArrayLength != 0 {
				var allVals []float64
				var val float64
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val float64
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "float32":
			if fieldDef.ArrayLength != 0 {
				var allVals []float32
				var val float32
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val float32
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
		case "string":
			if fieldDef.ArrayLength != 0 {
				var allVals []string
				var val string
				for i := 0; i < fieldDef.ArrayLength; i++ {
					if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
						return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
					}
					allVals = append(allVals, val)
				}
				if err := dm.SetField(fieldDef.Name, allVals); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			} else {
				var val string
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if err := dm.SetField(fieldDef.Name, val); err != nil {
					return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
				}
			}
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
	var dm *DynamicMessage
	var ok bool
	if dm, ok = msg.(*DynamicMessage); !ok {
		return nil, errors.New("message was not a *DynamicMessage")
	}
	if dm.t != mp {
		return nil, errors.New("wrong *DynamicMessage type")
	}

	// We're filling a buffer byte by byte.
	buf := &bytes.Buffer{}

	// Iterate over the definitions in the wire order.
	for _, fieldDef := range mp.msg.Fields {
		// Need to handle each type of field separately.
		switch fieldDef.Type {
		case "int8":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []int8
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]int8); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val int8
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(int8); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "uint8":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []uint8
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]uint8); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val uint8
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(uint8); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "int16":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []int16
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]int16); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val int16
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(int16); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "uint16":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []uint16
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]uint16); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val uint16
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(uint16); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "int32":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []int32
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]int32); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val int32
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(int32); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "uint32":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []uint32
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]uint32); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val uint32
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(uint32); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "int64":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []int64
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]int64); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val int64
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(int64); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "uint64":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []uint64
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]uint64); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val uint64
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(uint64); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "float64":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []float64
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]float64); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val float64
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(float64); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "float32":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []float32
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]float32); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val float32
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(float32); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		case "string":
			// Look up the actual value for this field.
			if fieldDef.ArrayLength != 0 {
				var val []string
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.([]string); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			} else {
				var val string
				if v, ok := dm.Fields[fieldDef.Name]; ok {
					if val, ok = v.(string); !ok {
						// The value stored for this field wasn't the right type.
						return nil, errors.New("invalid value for field: " + fieldDef.Name)
					}
				} // Else just use the default value.
				binary.Write(buf, binary.LittleEndian, val)
			}
		default:
			// We don't know what to do with this type.
			return nil, errors.New("unsupported field type: " + fieldDef.Name)
		}
	}

	newBuf := buf.Bytes()
	// empty-byte truncation
	// even with truncation, message length must be at least 1 byte
	// https://github.com/mavlink/c_library_v2/blob/master/mavlink_helpers.h#L103
	if isFrameV2 == true {
		end := len(newBuf)
		for end > 1 && newBuf[end-1] == 0x00 {
			end--
		}
		newBuf = newBuf[:end]
	}

	// All done.
	//fmt.Println(buf.Bytes())
	return newBuf, nil
}

// ALL DONE.
