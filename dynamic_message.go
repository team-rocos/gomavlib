package gomavlib

import (
	"encoding/base64"
	"encoding/json"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
	libgen "github.com/team-rocos/gomavlib/commands/dialgen/libgen"
)

// DynamicMessage : Used for RT message generation
type DynamicMessage struct {
	T      *DialectMessageRT
	Fields map[string]interface{}
}

// Assert to check we're implementing the interfaces we expect to be.
var _ Message = &DynamicMessage{}

// DynamicMessage :: Message

// GetId returns the MAVLink message ID (mID) of the DynamicMessage.
func (d DynamicMessage) GetId() uint32 {
	// Just look up the Id and return it.
	return uint32(d.T.Msg.Id)
}

// SetField sets DynamicMessage field matching field string, and based on its Type
func (d DynamicMessage) SetField(field string, value interface{}) error {
	// Search through the list of fields that the message is supposed to have.
	var fieldInfo *libgen.OutField
	for _, v := range d.T.Msg.Fields {
		if v.OriginalName == field {
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
		// Try to convert the value into a JsonFloat64.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]JsonFloat64); ok {
				d.Fields[field] = v
			} else if v, ok := value.([]float64); ok {
				jsonFloatSlice := make([]JsonFloat64, len(v))
				for i, val := range v {
					jsonFloatSlice[i] = JsonFloat64{F: val}
				}
				d.Fields[field] = jsonFloatSlice
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []JsonFloat64")
			}
		} else {
			if v, ok := value.(JsonFloat64); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else if v, ok := value.(float64); ok {
				d.Fields[field] = JsonFloat64{F: v}
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected JsonFloat64")
			}
		}
	case "float32":
		// Try to convert the value into a JsonFloat32.
		if fieldInfo.ArrayLength != 0 {
			if v, ok := value.([]JsonFloat32); ok {
				d.Fields[field] = v
			} else if v, ok := value.([]float32); ok {
				jsonFloatSlice := make([]JsonFloat32, len(v))
				for i, val := range v {
					jsonFloatSlice[i] = JsonFloat32{F: val}
				}
				d.Fields[field] = jsonFloatSlice
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected []JsonFloat32")
			}
		} else {
			if v, ok := value.(JsonFloat32); ok {
				// This is the correct type, so save it into our message.
				d.Fields[field] = v
			} else if v, ok := value.(float32); ok {
				d.Fields[field] = JsonFloat32{F: v}
			} else {
				// The value was the wrong type.
				return errors.New("incorrect type for field: " + field + " - expected JsonFloat32")
			}
		}
	case "string":
		// Try to convert the value into a string.
		if v, ok := value.(string); ok {
			// This is the correct type, so save it into our message.
			d.Fields[field] = v
		} else {
			// The value was the wrong type.
			return errors.New("incorrect type for field: " + field + " - expected string")
		}
	default:
		return errors.New("unsupported field type in dynamic MAVLink message")
	}

	// If we make it here, should mean everything went ok.
	return nil
}

// GetName returns OriginalName (in mavlink format)
func (d DynamicMessage) GetName() string {
	return d.T.Msg.OriginalName
}

// GenerateJSONSchema DynamicMessage exported function
func (d DynamicMessage) GenerateJSONSchema(prefix string, topic string) ([]byte, error) {
	return d.T.GenerateJSONSchema(prefix, topic)
}

func (d DynamicMessage) generateJSONSchemaProperties(topic string) (map[string]interface{}, error) {
	return d.T.generateJSONSchemaProperties(topic)
}

// MarshalJSON provides a custom implementation of JSON marshalling, so that when the DynamicMessage is recursively
// marshalled using the standard Go json.marshal() mechanism, the resulting JSON representation is a compact representation
// of just the important message payload (and not the message definition).  It's important that this representation matches
// the schema generated by GenerateJSONSchema().
func (d *DynamicMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Fields)
}

//UnmarshalJSON provides a custom implementation of JSON unmarshalling. Using the DynamicMessage provided, d.t.msg is used to
//determine the individual parsing of each JSON encoded payload item into the correct Type. It is important each type is
//correct so that the message serializes correctly and is understood by the MAVlink system
func (d *DynamicMessage) UnmarshalJSON(buf []byte) error {
	//Declaring temp variables to be used across the unmarshaller
	var err error
	var field *libgen.OutField
	var keyName []byte
	var data interface{}
	var fieldExists bool

	//Declaring jsonparser unmarshalling functions
	var arrayHandler func([]byte, jsonparser.ValueType, int, error)
	var objectHandler func([]byte, []byte, jsonparser.ValueType, int) error

	//JSON key is an array
	arrayHandler = func(key []byte, dataType jsonparser.ValueType, offset int, err error) {
		switch dataType.String() {
		//We have a string array
		case "string":
			// For special cases where floats are marshalled as a string (when they have a value of nan, +inf, or -inf)
			if field.Type == "float32" || field.Type == "float64" {
				data, err = strconv.ParseFloat(string(key), 64)
				if err != nil {
					errors.Wrap(err, "Field: "+field.OriginalName)
				}
				if field.Type == "float32" {
					d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]JsonFloat32), JsonFloat32{F: float32((data.(float64)))})
				} else {
					d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]JsonFloat64), JsonFloat64{F: data.(float64)})
				}
			} else {
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]string), string(key))
			}
		//We have a number or int array.
		case "number":
			//We have a float to parse
			if field.Type == "float64" || field.Type == "float32" {
				data, err = strconv.ParseFloat(string(key), 64)
				if err != nil {
					errors.Wrap(err, "Field: "+field.OriginalName)
				}
			} else {
				data, err = strconv.ParseInt(string(key), 0, 64)
				if err != nil {
					errors.Wrap(err, "Field: "+field.OriginalName)
				}
			}
			//Append field to data array
			switch field.Type {
			case "int8":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]int8), int8((data.(int64))))
			case "int16":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]int16), int16((data.(int64))))
			case "int32":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]int32), int32((data.(int64))))
			case "int64":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]int64), int64((data.(int64))))
			case "uint8":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]uint8), uint8((data.(int64))))
			case "uint16":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]uint16), uint16((data.(int64))))
			case "uint32":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]uint32), uint32((data.(int64))))
			case "uint64":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]uint64), uint64((data.(int64))))
			case "float32":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]JsonFloat32), JsonFloat32{F: float32((data.(float64)))})
			case "float64":
				d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]JsonFloat64), JsonFloat64{F: data.(float64)})
			}
		//We have a bool array
		case "boolean":
			data, err := jsonparser.GetBoolean(buf, string(key))
			_ = err
			d.Fields[field.OriginalName] = append(d.Fields[field.OriginalName].([]bool), data)
		}

		//Null error as it is not returned in ArrayEach, requires package modification
		_ = err
		//Null keyName to prevent repeat scenarios of same key usage
		_ = keyName

	}

	//JSON key handler
	objectHandler = func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		//Store keyName for usage in ArrayEach function
		keyName = key
		fieldExists = false
		//Find message spec field that matches JSON key
		for _, f := range d.T.Msg.Fields {
			if string(key) == f.OriginalName {
				field = f
				fieldExists = true
				break
			}
		}
		if fieldExists == true {
			//Scalars First
			switch dataType.String() {
			//We have a JSON string
			case "string":
				//Special case where we have a byte array encoded as JSON string
				if field.Type == "uint8" {
					data, err := base64.StdEncoding.DecodeString(string(value))
					if err != nil {
						return errors.Wrap(err, "Byte Array Field: "+field.OriginalName)
					}
					d.Fields[field.OriginalName] = data
				} else if field.Type == "float32" || field.Type == "float64" { //Case where we have marshalled a special float as a string
					data, err = strconv.ParseFloat(string(value), 64)
					if err != nil {
						return errors.Wrap(err, "Field: "+field.OriginalName)
					}
					if field.Type == "float32" {
						d.Fields[field.OriginalName] = JsonFloat32{F: float32(data.(float64))}
					} else {
						d.Fields[field.OriginalName] = JsonFloat64{F: data.(float64)}
					}
				} else {
					d.Fields[field.OriginalName] = string(value)
				}
			//We have a JSON number or int
			case "number":
				//We have a float to parse
				if field.Type == "float64" || field.Type == "float32" {
					data, err = jsonparser.GetFloat(buf, string(key))
					if err != nil {
						return errors.Wrap(err, "Field: "+field.OriginalName)
					}
					//We have an int to parse
				} else {
					data, err = jsonparser.GetInt(buf, string(key))
					if err != nil {
						return errors.Wrap(err, "Field: "+field.OriginalName)
					}
				}
				//Copy number value to message field
				switch field.Type {
				case "int8":
					d.Fields[field.OriginalName] = int8(data.(int64))
				case "int16":
					d.Fields[field.OriginalName] = int16(data.(int64))
				case "int32":
					d.Fields[field.OriginalName] = int32(data.(int64))
				case "int64":
					d.Fields[field.OriginalName] = int64(data.(int64))
				case "uint8":
					d.Fields[field.OriginalName] = uint8(data.(int64))
				case "uint16":
					d.Fields[field.OriginalName] = uint16(data.(int64))
				case "uint32":
					d.Fields[field.OriginalName] = uint32(data.(int64))
				case "uint64":
					d.Fields[field.OriginalName] = uint64(data.(int64))
				case "float32":
					d.Fields[field.OriginalName] = JsonFloat32{F: float32(data.(float64))}
				case "float64":
					d.Fields[field.OriginalName] = JsonFloat64{F: data.(float64)}
				}
			//We have a JSON bool
			case "boolean":
				data, err := jsonparser.GetBoolean(buf, string(key))
				if err != nil {
					return errors.Wrap(err, "Field: "+field.OriginalName)
				}
				d.Fields[field.OriginalName] = data
			//We have a JSON array
			case "array":
				//Redeclare message array fields incase they do not exist
				switch field.Type {
				case "bool":
					d.Fields[field.OriginalName] = make([]bool, 0)
				case "int8":
					d.Fields[field.OriginalName] = make([]int8, 0)
				case "int16":
					d.Fields[field.OriginalName] = make([]int16, 0)
				case "int32":
					d.Fields[field.OriginalName] = make([]int32, 0)
				case "int64":
					d.Fields[field.OriginalName] = make([]int64, 0)
				case "uint8":
					d.Fields[field.OriginalName] = make([]uint8, 0)
				case "uint16":
					d.Fields[field.OriginalName] = make([]uint16, 0)
				case "uint32":
					d.Fields[field.OriginalName] = make([]uint32, 0)
				case "uint64":
					d.Fields[field.OriginalName] = make([]uint64, 0)
				case "float32":
					d.Fields[field.OriginalName] = make([]JsonFloat32, 0)
				case "float64":
					d.Fields[field.OriginalName] = make([]JsonFloat64, 0)
				case "string":
					d.Fields[field.OriginalName] = make([]string, 0)
				default:
					//goType is a nested Message array
					d.Fields[field.OriginalName] = make([]Message, 0)
				}
				//Parse JSON array
				jsonparser.ArrayEach(value, arrayHandler)
			default:
				//We do nothing here as blank fields may return value type NotExist or Null
				err = errors.Wrap(err, "Null field: "+string(key))
			}
		} else {
			return errors.New("Field Unknown: " + string(key))
		}
		return err
	}
	//Perform JSON object handler function
	err = jsonparser.ObjectEach(buf, objectHandler)
	return err
}
