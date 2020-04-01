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
	Messages map[uint32]*DialectMessageRT
}

// Assert to check we're implementing the interfaces we expect to be.
var _ Dialect = &DialectRT{}

// DEFINE PRIVATE TYPES AND STRUCTURES.

// DialectMessageRT is a run time dialect message which contains the message information
// and a pointer to the DialectRT to which it belongs.
type DialectMessageRT struct {
	dialect      *DialectRT
	Msg          *libgen.OutMessage
	crcExtra     byte
	sizeNormal   uint
	sizeExtended uint
}

// Assert to check we're implementing the interfaces we expect to be.
var _ dialectMessage = &DialectMessageRT{}

// DEFINE PUBLIC STATIC FUNCTIONS.

// NewDialectRT allocates a Dialect.
func NewDialectRT(version uint, outDefs []*libgen.OutDefinition) (*DialectRT, error) {
	// Create the new dialect object.
	d := &DialectRT{
		version:  version,
		defs:     outDefs,
		Messages: make(map[uint32]*DialectMessageRT),
	}

	// Populate the set of messages in the dialect by instantiating a dialectMessage from each message entry in the definitions.
	for _, def := range d.defs {
		for _, msg := range def.Messages {

			// Reorder fields as described in https://mavlink.io/en/guide/serialization.html#field_reordering
			sort.Slice(msg.Fields, func(i, j int) bool {
				// sort by weight if not extension
				if msg.Fields[i].IsExtension == false && msg.Fields[j].IsExtension == false {
					if w1, w2 := DialectFieldTypeSizes[DialectFieldTypeFromGo[msg.Fields[i].Type]], DialectFieldTypeSizes[DialectFieldTypeFromGo[msg.Fields[j].Type]]; w1 != w2 {
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

					h.Write([]byte(DialectFieldTypeString[DialectFieldTypeFromGo[f.Type]] + " "))

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
				var size uint = uint(DialectFieldTypeSizes[DialectFieldTypeFromGo[f.Type]])

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
			d.Messages[uint32(msg.Id)] = &DialectMessageRT{
				dialect:      d,
				Msg:          msg,
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

// GetVersion returns the uint version of DialectRT
func (d *DialectRT) GetVersion() uint {
	return d.version
}

func (d *DialectRT) getVersion() uint {
	return d.version
}

func (d *DialectRT) getMsgById(id uint32) (*dialectMessage, bool) {
	var msg dialectMessage
	var ok bool
	msg, ok = d.Messages[id]
	return &msg, ok
}

// GetMessages is a public get function to return the messages variable of DialectRT
// func (d *DialectRT) GetMessages() map[uint32]*dialectMessageRT {
// 	return d.messages
// }

// CreateMessageById returns a DynamicMessage created from finding the message in the DialectRT given corresponding to id.
// The Fields map in the DynamicMessage created is empty.
func (d *DialectRT) CreateMessageById(id uint32) (*DynamicMessage, error) {
	dm := &DynamicMessage{}
	fields := make(map[string]interface{})
	msg, ok := d.Messages[id]
	if !ok {
		errorString := fmt.Sprintf("message with id=%d does not exist in this dialectRT", id)
		return dm, errors.New(errorString)
	}
	dm.T = msg
	dm.Fields = fields
	return dm, nil
}

// CreateMessageByName returns a DynamicMessage created from finding the message in the DialectRT given corresponding to name.
// The Fields map in the DynamicMessage created is empty.
func (d *DialectRT) CreateMessageByName(name string) (*DynamicMessage, error) {
	dm := &DynamicMessage{}
	fields := make(map[string]interface{})
	var msg *DialectMessageRT
	foundMessage := false
	for _, m := range d.Messages {
		if m.Msg.OriginalName == name {
			msg = m
			foundMessage = true
			break
		}
	}
	if !foundMessage {
		errorString := "message with name " + name + " does not exist in this dialectRT"
		return dm, errors.New(errorString)
	}
	dm.T = msg
	dm.Fields = fields
	return dm, nil
}

// JSON Variables
const (
	//Sep is a namespace separator string
	Sep = "/"
	//GlobalNS is the global namespace initial separator string
	GlobalNS = "/"
	//PrivateNS is private namespace initial separator string
	PrivateNS = "~"
)

// GenerateJSONSchema generates a (primitive) JSON schema for the associated DynamicMessage; however note that since
// we are mostly interested in making schema's for particular _topics_, the function takes a string prefix, and string topic name, which are
// used to id the resulting schema.
func (mp *DialectMessageRT) GenerateJSONSchema(prefix string, topic string) ([]byte, error) {
	// The JSON schema for a message consist of the (recursive) properties names/types:
	schemaItems, err := mp.generateJSONSchemaProperties(prefix + Sep + topic)
	if err != nil {
		return nil, err
	}

	// Plus some extra keywords:
	schemaItems["$schema"] = "https://json-schema.org/draft-07/schema#"
	schemaItems["$id"] = prefix + Sep + topic

	// The schema itself is created from the map of properties.
	schemaString, err := json.Marshal(schemaItems)
	if err != nil {
		return nil, err
	}

	// All done.
	return schemaString, nil
}

func (mp *DialectMessageRT) generateJSONSchemaProperties(topic string) (map[string]interface{}, error) {
	// // Each message's schema indicates that it is an 'object' with some nested properties: those properties are the fields and their types.
	properties := make(map[string]interface{})
	schemaItems := make(map[string]interface{})
	schemaItems["type"] = "object"
	schemaItems["title"] = topic
	schemaItems["properties"] = properties

	// Iterate over each of the fields in the message.
	for _, field := range mp.Msg.Fields {
		if field.ArrayLength != 0 && field.Type != "string" {
			// It's an array.
			propertyContent := make(map[string]interface{})
			properties[field.Name] = propertyContent

			if field.Type == "uint8" {
				propertyContent["title"] = topic + Sep + field.Name
				propertyContent["type"] = "string"
			} else {

				// Arrays all have a type of 'array', regardless of that the hold, then the 'item' keyword determines what type goes in the array.
				propertyContent["type"] = "array"
				propertyContent["title"] = topic + Sep + field.Name
				arrayItems := make(map[string]interface{})
				propertyContent["items"] = arrayItems

				// Need to handle each type appropriately.
				if field.Type == "string" {
					arrayItems["type"] = "string"
				} else {
					// It's a primitive.
					var jsonType string
					if field.Type == "int8" || field.Type == "int16" || field.Type == "uint16" ||
						field.Type == "int32" || field.Type == "uint32" || field.Type == "int64" || field.Type == "uint64" {
						jsonType = "integer"
					} else if field.Type == "float32" || field.Type == "float64" {
						jsonType = "number"
					} else if field.Type == "bool" {
						jsonType = "bool"
					} else {
						// Something went wrong.
						return nil, errors.New("we haven't implemented this primitive yet")
					}
					arrayItems["type"] = jsonType
				}
			}
		} else {
			// It's a scalar.
			propertyContent := make(map[string]interface{})
			properties[field.Name] = propertyContent
			propertyContent["title"] = topic + Sep + field.Name

			if field.Type == "string" {
				propertyContent["type"] = "string"
			} else {
				// It's a primitive.
				var jsonType string
				if field.Type == "int8" || field.Type == "uint8" || field.Type == "int16" || field.Type == "uint16" ||
					field.Type == "int32" || field.Type == "uint32" || field.Type == "int64" || field.Type == "uint64" {
					jsonType = "integer"
					jsonType = "integer"
					jsonType = "integer"
				} else if field.Type == "float32" || field.Type == "float64" {
					jsonType = "number"
				} else if field.Type == "bool" {
					jsonType = "bool"
				} else {
					// Something went wrong.
					return nil, errors.New("we haven't implemented this primitive yet")
				}
				propertyContent["type"] = jsonType
			}
		}
	}

	// All done.
	return schemaItems, nil
}

// GetMessageId returns the id of the dialectMessageRT
func (mp *DialectMessageRT) GetMessageId() int {
	return mp.Msg.Id
}

// GetSizeNormal returns sizeNormal of DialectMessageRT
func (mp *DialectMessageRT) GetSizeNormal() uint {
	return mp.sizeNormal
}

// GetSizeExtended returns sizeExtended of DialectMessageRT
func (mp *DialectMessageRT) GetSizeExtended() uint {
	return mp.sizeExtended
}

// GetCRCExtra returns crcExtra of DialectMessageRT
func (mp *DialectMessageRT) GetCRCExtra() byte {
	return mp.crcExtra
}

// DEFINE PRIVATE STATIC FUNCTIONS.

// DEFINE PRIVATE RECEIVER FUNCTIONS.

func (mp *DialectMessageRT) newMsg() Message {
	// Just make an empty dynamic message which points at this dialectMessage as its parent.
	msg := &DynamicMessage{
		T:      mp,
		Fields: make(map[string]interface{}),
	}
	return msg
}

func (mp *DialectMessageRT) getFields() []*DialectMessageField {
	// Iterate over each of the field definitions and construct a new dialectMessageField representation for each.
	fields := make([]*DialectMessageField, 0)
	for _, f := range mp.Msg.Fields {
		ftype := DialectFieldTypeFromGo[f.Type]
		fields = append(fields, &DialectMessageField{
			isEnum:      f.IsEnum,
			ftype:       ftype,
			name:        f.Name,
			arrayLength: DialectFieldTypeSizes[ftype],
			index:       f.Index,
			isExtension: f.IsExtension,
		})
	}
	return fields
}

func (mp *DialectMessageRT) getCRCExtra() byte {
	return mp.crcExtra
}

func (mp *DialectMessageRT) decode(buf []byte, isFrameV2 bool) (Message, error) {
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

	for _, fieldDef := range mp.Msg.Fields { // TODO - Hmm, this implies that mp.msg.Fields is in order?  In that case, why do we need orderedFields?
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
			var allVals string
			var val uint8 // ASCII represented chars
			for i := 0; i < fieldDef.ArrayLength; i++ {
				if err := binary.Read(b, binary.LittleEndian, &val); err != nil {
					return nil, errors.Wrap(err, "failed to read field : "+fieldDef.Name+" : ")
				}
				if val == 0 {
					continue // Continue reading until end of string as determined by fieldDef.ArrayLength
				}
				allVals += string(val)
			}
			if err := dm.SetField(fieldDef.Name, allVals); err != nil {
				return nil, errors.Wrap(err, "failed to set field : "+fieldDef.Name+" : ")
			}
		default:
			// We don't know what to do with this type.
			return nil, errors.New("unsupported field type: " + fieldDef.Name)
		}
	}

	// All done.
	return dm, nil
}

func (mp *DialectMessageRT) encode(msg Message, isFrameV2 bool) ([]byte, error) {
	// Make sure the message we're encoding matches the type of the dialectMessage being used to do the encoding.
	var dm *DynamicMessage
	var ok bool
	if dm, ok = msg.(*DynamicMessage); !ok {
		return nil, errors.New("message was not a *DynamicMessage")
	}
	if dm.T != mp {
		return nil, errors.New("wrong *DynamicMessage type")
	}

	// We're filling a buffer byte by byte.
	buf := &bytes.Buffer{}

	// Iterate over the definitions in the wire order.
	for _, fieldDef := range mp.Msg.Fields {
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
			var val string
			if v, ok := dm.Fields[fieldDef.Name]; ok {
				if val, ok = v.(string); !ok {
					// The value stored for this field wasn't the right type.
					return nil, errors.New("invalid value for field: " + fieldDef.Name)
				}
			} // Else just use the default value.
			numberWritten := 0
			for _, c := range val {
				binary.Write(buf, binary.LittleEndian, uint8(c))
				numberWritten++
			}
			// Write remaining zeros if not yet filled up to ArrayLength
			for numberWritten < fieldDef.ArrayLength {
				binary.Write(buf, binary.LittleEndian, uint8(0x0))
				numberWritten++
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

// Decode is a public function necessary for testing
func (mp *DialectMessageRT) Decode(buf []byte, isFrameV2 bool) (Message, error) {
	return mp.decode(buf, isFrameV2)
}

// Encode is a public function necessary for testing
func (mp *DialectMessageRT) Encode(msg Message, isFrameV2 bool) ([]byte, error) {
	return mp.encode(msg, isFrameV2)
}

// ALL DONE.
