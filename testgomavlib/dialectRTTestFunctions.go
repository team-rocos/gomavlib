package testgomavlib

import (
	"bytes"
	"errors"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/team-rocos/gomavlib"
	libgen "github.com/team-rocos/gomavlib/commands/dialgen/libgen"
	"github.com/xeipuuv/gojsonschema"
)

// DEFINE PUBLIC TYPES AND STRUCTURES.

// DEFINE PRIVATE TYPES AND STRUCTURES.
// DEFINE PUBLIC STATIC FUNCTIONS.

// CreateMessageByIdTest creates a dynamic message based on the input id and checks that the values within it are valid.
func CreateMessageByIdTest(t *testing.T, xmlPath string, includeDirs []string) {
	defs, version, err := libgen.XMLToFields(xmlPath, includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := gomavlib.NewDialectRT(version, defs)
	require.NoError(t, err)

	// Create dynamic message using id of each message in dRT
	for _, mRT := range dRT.Messages {
		dm, err := dRT.CreateMessageById(uint32(mRT.Msg.Id))
		require.NoError(t, err)
		require.Equal(t, mRT, dm.T)
	}

	// CreateMessageById using invalid id. Assert that error is returned
	_, err = dRT.CreateMessageById(40000000)
	assert.Error(t, err)
}

// CreateMessageByNameTest creates a dynamic message based on the input name and checks that the values within it are valid.
func CreateMessageByNameTest(t *testing.T, xmlPath string, includeDirs []string) {
	defs, version, err := libgen.XMLToFields(xmlPath, includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := gomavlib.NewDialectRT(version, defs)
	require.NoError(t, err)

	// Create dynamic message by name using name from each mRT in dRT
	for _, mRT := range dRT.Messages {
		dm, err := dRT.CreateMessageByName(mRT.Msg.OriginalName)
		require.NoError(t, err)
		require.Equal(t, mRT, dm.T)
	}

	// Create dynamic message using invalid name. Assert that error is returned
	_, err = dRT.CreateMessageByName("abcdefghijklmnop***")
	assert.Error(t, err)
}

// JSONMarshalAndUnmarshalTest tests JSON generation, schema generation, and JSON unmarshal code.
func JSONMarshalAndUnmarshalTest(t *testing.T, xmlPath string, includeDirs []string) {
	for i, c := range casesMsgsTest {
		dCT, err := gomavlib.NewDialectCT(3, ctMessages)
		require.NoError(t, err)
		dMsgCT, ok := dCT.Messages[c.id]
		require.Equal(t, true, ok)
		bytesEncoded, err := dMsgCT.Encode(c.parsed, c.isV2)
		require.NoError(t, err)
		require.Equal(t, c.raw, bytesEncoded)

		// Decode bytes using RT
		defs, version, err := libgen.XMLToFields(xmlPath, includeDirs)
		require.NoError(t, err)

		// Create dialect from the parsed defs.
		dRT, err := gomavlib.NewDialectRT(version, defs)
		require.NoError(t, err)
		dMsgRT := dRT.Messages[c.id]
		require.Equal(t, uint(3), dRT.GetVersion())

		// Decode bytes using RT
		msgDecoded, err := dMsgRT.Decode(c.raw, c.isV2)
		require.NoError(t, err)

		// Marshal JSON
		bytesCreated, err := msgDecoded.(*gomavlib.DynamicMessage).MarshalJSON()
		require.NoError(t, err)
		if i == 7 || i == 8 { // Test cases with altered JSON
			require.NotEqual(t, jsonTest[i], string(bytesCreated))
		} else {
			require.Equal(t, jsonTest[i], string(bytesCreated))
		}

		// Generate JSON Schema
		schemaBytes, err := msgDecoded.(*gomavlib.DynamicMessage).GenerateJSONSchema("/mavlink", "topic")
		require.NoError(t, err)
		if i == 7 { // Test case with altered schema example
			require.NotEqual(t, schemasTest[i], string(schemaBytes))
		} else {
			require.Equal(t, schemasTest[i], string(schemaBytes))
		}

		// Validate JSON document against schema
		schemaLoader := gojsonschema.NewStringLoader(schemasTest[i])
		documentLoader := gojsonschema.NewStringLoader(jsonTest[i])
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if i == 8 { // JSONTest[8] has a string entry where it should be float32 - should not validate against schemasTest[8]
			require.NoError(t, err)
			require.Equal(t, false, result.Valid())
		} else {
			require.NoError(t, err)
			require.Equal(t, true, result.Valid())
		}

		// Test Unmarshal
		// Create new DynamicMessage with empty fields for testing unmarshal
		dm, err := dRT.CreateMessageById(uint32(dRT.Messages[c.id].Msg.Id))
		require.NoError(t, err)
		err = dm.UnmarshalJSON(bytesCreated)
		require.NoError(t, err)
		require.Equal(t, msgDecoded.(*gomavlib.DynamicMessage).Fields, dm.Fields)
	}
}

// DialectRTCommonXMLTest tests the XMLToFields and RT dialect generation functionality added to gomavlib.
func DialectRTCommonXMLTest(t *testing.T, xmlPath string, includeDirs []string) {
	// Ensure that XMLToFields works with no include files, if xml file has no includes
	_, _, err := libgen.XMLToFields(xmlPath, make([]string, 0))
	require.NoError(t, err)

	// Parse the XML file.
	defs, version, err := libgen.XMLToFields(xmlPath, includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := gomavlib.NewDialectRT(version, defs)
	require.NoError(t, err)
	require.Equal(t, uint(3), dRT.GetVersion())

	// Check Individual Messages for RT
	msg := dRT.Messages[5].Msg
	require.Equal(t, "ChangeOperatorControl", msg.Name)
	require.Equal(t, 5, msg.Id)
	field := msg.Fields[0]
	require.Equal(t, "TargetSystem", field.Name)
	require.Equal(t, "uint8", field.Type)
	field = msg.Fields[3]
	require.Equal(t, "Passkey", field.Name)
	require.Equal(t, "string", field.Type)

	// Checking Message 82 - Has float[4] array as a field
	msg = dRT.Messages[82].Msg
	require.Equal(t, "SetAttitudeTarget", msg.Name)
	require.Equal(t, 82, msg.Id)
	field = msg.Fields[1]
	require.Equal(t, "Q", field.Name)
	require.Equal(t, "float32", field.Type)

	// Compare with DialectCT
	dCT, err := gomavlib.NewDialectCT(3, ctMessages)
	require.NoError(t, err)

	require.Equal(t, len(dCT.Messages), len(dRT.Messages))
	// Compare RT and CT for all messages
	for _, m := range ctMessages {
		index := m.GetId()
		// Compare dCT with dRT
		mCT := dCT.Messages[index]
		mRT := dRT.Messages[index]
		require.Equal(t, mCT.GetSizeNormal(), byte(mRT.GetSizeNormal()))
		require.Equal(t, mCT.GetSizeExtended(), byte(mRT.GetSizeExtended()))
		require.Equal(t, mCT.GetCRCExtra(), mRT.GetCRCExtra())

		// Compare all fields of all RT and CT Messages
		for i := 0; i < len(mCT.Fields); i++ {
			fCT := mCT.Fields[i]
			fRT := mRT.Msg.Fields[i]
			require.Equal(t, fCT.GetIsEnum(), fRT.IsEnum)
			require.Equal(t, fCT.GetFType(), gomavlib.DialectFieldTypeFromGo[fRT.Type])
			require.Equal(t, fCT.GetName(), fRT.OriginalName)
			require.Equal(t, fCT.GetArrayLength(), byte(fRT.ArrayLength))
			require.Equal(t, fCT.GetIndex(), fRT.Index)
			require.Equal(t, fCT.GetIsExtension(), fRT.IsExtension)
		}
	}
}

// DecodeAndEncodeRTTest tests run time (RT) encoding and decoding of messages
func DecodeAndEncodeRTTest(t *testing.T, xmlPath string, includeDirs []string) {
	for _, c := range casesMsgsTest {
		// Encode using CT
		dCT, err := gomavlib.NewDialectCT(3, ctMessages)
		require.NoError(t, err)
		dMsgCT, ok := dCT.Messages[c.id]
		require.Equal(t, true, ok)
		bytesEncoded, err := dMsgCT.Encode(c.parsed, c.isV2)
		require.NoError(t, err)
		require.Equal(t, c.raw, bytesEncoded)

		// Decode bytes using CT method for RT vs CT comparison later
		msgDecodedCT, err := dMsgCT.Decode(c.raw, c.isV2)
		require.NoError(t, err)
		require.Equal(t, c.parsed, msgDecodedCT)

		// Decode bytes using RT
		defs, version, err := libgen.XMLToFields(xmlPath, includeDirs)
		require.NoError(t, err)

		// Create dialect from the parsed defs.
		dRT, err := gomavlib.NewDialectRT(version, defs)
		dMsgRT := dRT.Messages[c.id]
		require.NoError(t, err)
		require.Equal(t, uint(3), dRT.GetVersion())

		// Decode bytes using RT
		msgDecoded, err := dMsgRT.Decode(bytesEncoded, c.isV2)
		require.NoError(t, err)

		//Make sure all fields of dMsgCT match equivalent values of RT msgDecoded
		//Compare all fields of all RT and CT Messages
		v := reflect.ValueOf(msgDecodedCT).Elem()
		for j := 0; j < len(dMsgCT.Fields); j++ {
			fCT := dMsgCT.Fields[j]
			originalName := fCT.GetName()
			name := dialectMsgDefToGo(originalName)
			fRT := msgDecoded.(*gomavlib.DynamicMessage).Fields[originalName]
			fCTVal := v.FieldByName(name)
			fieldType, arrayLength, err := findFieldType(msgDecoded.(*gomavlib.DynamicMessage), originalName)
			require.NoError(t, err)
			switch fieldType {
			case "int8":
				if arrayLength != 0 {
					rtResult := fRT.([]int8)
					ctResult := make([]int8, arrayLength)
					for i := 0; i < arrayLength; i++ {
						if fCTVal.Index(i).Kind() == reflect.Int || fCTVal.Index(i).Kind() == reflect.Int8 {
							ctResult[i] = int8(fCTVal.Index(i).Int())
						} else {
							ctResult[i] = int8(fCTVal.Index(i).Uint())
						}
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					if fCTVal.Kind() == reflect.Int || fCTVal.Kind() == reflect.Int8 {
						require.Equal(t, int8(fCTVal.Int()), fRT.(int8))
					} else {
						require.Equal(t, int8(fCTVal.Uint()), fRT.(int8))
					}
				}
			case "uint8":
				if arrayLength != 0 {
					rtResult := fRT.([]uint8)
					ctResult := make([]uint8, arrayLength)
					for i := 0; i < arrayLength; i++ {
						if fCTVal.Index(i).Kind() == reflect.Int || fCTVal.Index(i).Kind() == reflect.Int8 {
							ctResult[i] = uint8(fCTVal.Index(i).Int())
						} else {
							ctResult[i] = uint8(fCTVal.Index(i).Uint())
						}
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					if fCTVal.Kind() == reflect.Int || fCTVal.Kind() == reflect.Int8 {
						require.Equal(t, uint8(fCTVal.Int()), fRT.(uint8))
					} else {
						require.Equal(t, uint8(fCTVal.Uint()), fRT.(uint8))
					}
				}
			case "int16":
				if arrayLength != 0 {
					rtResult := fRT.([]int16)
					ctResult := make([]int16, arrayLength)
					for i := 0; i < arrayLength; i++ {
						if fCTVal.Index(i).Kind() == reflect.Int || fCTVal.Index(i).Kind() == reflect.Int16 {
							ctResult[i] = int16(fCTVal.Index(i).Int())
						} else {
							ctResult[i] = int16(fCTVal.Index(i).Uint())
						}
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					if fCTVal.Kind() == reflect.Int || fCTVal.Kind() == reflect.Int16 {
						require.Equal(t, int16(fCTVal.Int()), fRT.(int16))
					} else {
						require.Equal(t, int16(fCTVal.Uint()), fRT.(int16))
					}
				}

			case "uint16":
				if arrayLength != 0 {
					rtResult := fRT.([]uint16)
					ctResult := make([]uint16, arrayLength)
					for i := 0; i < arrayLength; i++ {
						if fCTVal.Index(i).Kind() == reflect.Int || fCTVal.Index(i).Kind() == reflect.Int16 {
							ctResult[i] = uint16(fCTVal.Index(i).Int())
						} else {
							ctResult[i] = uint16(fCTVal.Index(i).Uint())
						}
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					if fCTVal.Kind() == reflect.Int || fCTVal.Kind() == reflect.Int16 {
						require.Equal(t, uint16(fCTVal.Int()), fRT.(uint16))
					} else {
						require.Equal(t, uint16(fCTVal.Uint()), fRT.(uint16))
					}
				}

			case "int32":
				if arrayLength != 0 {
					rtResult := fRT.([]int32)
					ctResult := make([]int32, arrayLength)
					for i := 0; i < arrayLength; i++ {
						if fCTVal.Index(i).Kind() == reflect.Int || fCTVal.Index(i).Kind() == reflect.Int32 {
							ctResult[i] = int32(fCTVal.Index(i).Int())
						} else {
							ctResult[i] = int32(fCTVal.Index(i).Uint())
						}
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					if fCTVal.Kind() == reflect.Int || fCTVal.Kind() == reflect.Int32 {
						require.Equal(t, int32(fCTVal.Int()), fRT.(int32))
					} else {
						require.Equal(t, int32(fCTVal.Uint()), fRT.(int32))
					}
				}
			case "uint32":
				if arrayLength != 0 {
					rtResult := fRT.([]uint32)
					ctResult := make([]uint32, arrayLength)
					for i := 0; i < arrayLength; i++ {
						if fCTVal.Index(i).Kind() == reflect.Int || fCTVal.Index(i).Kind() == reflect.Int32 {
							ctResult[i] = uint32(fCTVal.Index(i).Int())
						} else {
							ctResult[i] = uint32(fCTVal.Index(i).Uint())
						}
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					if fCTVal.Kind() == reflect.Int || fCTVal.Kind() == reflect.Int32 {
						require.Equal(t, uint32(fCTVal.Int()), fRT.(uint32))
					} else {
						require.Equal(t, uint32(fCTVal.Uint()), fRT.(uint32))
					}
				}

			case "int64":
				if arrayLength != 0 {
					rtResult := fRT.([]int64)
					ctResult := make([]int64, arrayLength)
					for i := 0; i < arrayLength; i++ {
						if fCTVal.Index(i).Kind() == reflect.Int || fCTVal.Index(i).Kind() == reflect.Int64 {
							ctResult[i] = int64(fCTVal.Index(i).Int())
						} else {
							ctResult[i] = int64(fCTVal.Index(i).Uint())
						}
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					if fCTVal.Kind() == reflect.Int || fCTVal.Kind() == reflect.Int64 {
						require.Equal(t, int64(fCTVal.Int()), fRT.(int64))
					} else {
						require.Equal(t, int64(fCTVal.Uint()), fRT.(int64))
					}
				}

			case "uint64":
				if arrayLength != 0 {
					rtResult := fRT.([]uint64)
					ctResult := make([]uint64, arrayLength)
					for i := 0; i < arrayLength; i++ {
						if fCTVal.Index(i).Kind() == reflect.Int || fCTVal.Index(i).Kind() == reflect.Int64 {
							ctResult[i] = uint64(fCTVal.Index(i).Int())
						} else {
							ctResult[i] = uint64(fCTVal.Index(i).Uint())
						}
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					if fCTVal.Kind() == reflect.Int || fCTVal.Kind() == reflect.Int64 {
						require.Equal(t, uint64(fCTVal.Int()), fRT.(uint64))
					} else {
						require.Equal(t, uint64(fCTVal.Uint()), fRT.(uint64))
					}
				}
			case "float64":
				if arrayLength != 0 {
					rtResult := fRT.([]float64)
					ctResult := make([]float64, arrayLength)
					for i := 0; i < arrayLength; i++ {
						ctResult[i] = float64(fCTVal.Index(i).Float())
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					require.Equal(t, float64(fCTVal.Float()), fRT.(float64))
				}

			case "float32":
				if arrayLength != 0 {
					rtResult := fRT.([]float32)
					ctResult := make([]float32, arrayLength)
					for i := 0; i < arrayLength; i++ {
						ctResult[i] = float32(fCTVal.Index(i).Float())
					}
					require.Equal(t, ctResult, rtResult)
				} else {
					require.Equal(t, float32(fCTVal.Float()), fRT.(float32))
				}

			case "string":
				require.Equal(t, fCTVal.String(), fRT.(string))

			default:
				err = errors.New("invalid type so unable to convert interface values")
				panic(err)
			}
		}

		// Encode using RT
		bytesEncodedByRT, err := dMsgRT.Encode(msgDecoded, c.isV2)
		require.NoError(t, err)
		require.Equal(t, c.raw, bytesEncodedByRT)
	}
}

func findFieldType(dm *gomavlib.DynamicMessage, originalName string) (string, int, error) {
	for _, f := range dm.T.Msg.Fields {
		if f.OriginalName == originalName {
			return f.Type, f.ArrayLength, nil
		}
	}
	err := errors.New("field with given OriginalName does not exist")
	return "", 0, err
}

var schemasTest = []string{
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"acc_x\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/acc_x\",\"type\":\"array\"},\"acc_y\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/acc_y\",\"type\":\"array\"},\"acc_z\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/acc_z\",\"type\":\"array\"},\"command\":{\"items\":{\"type\":\"integer\"},\"title\":\"/mavlink/topic/command\",\"type\":\"array\"},\"pos_x\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/pos_x\",\"type\":\"array\"},\"pos_y\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/pos_y\",\"type\":\"array\"},\"pos_yaw\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/pos_yaw\",\"type\":\"array\"},\"pos_z\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/pos_z\",\"type\":\"array\"},\"time_usec\":{\"title\":\"/mavlink/topic/time_usec\",\"type\":\"integer\"},\"valid_points\":{\"title\":\"/mavlink/topic/valid_points\",\"type\":\"integer\"},\"vel_x\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/vel_x\",\"type\":\"array\"},\"vel_y\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/vel_y\",\"type\":\"array\"},\"vel_yaw\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/vel_yaw\",\"type\":\"array\"},\"vel_z\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/vel_z\",\"type\":\"array\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"flow_comp_m_x\":{\"title\":\"/mavlink/topic/flow_comp_m_x\",\"type\":\"number\"},\"flow_comp_m_y\":{\"title\":\"/mavlink/topic/flow_comp_m_y\",\"type\":\"number\"},\"flow_rate_x\":{\"title\":\"/mavlink/topic/flow_rate_x\",\"type\":\"number\"},\"flow_rate_y\":{\"title\":\"/mavlink/topic/flow_rate_y\",\"type\":\"number\"},\"flow_x\":{\"title\":\"/mavlink/topic/flow_x\",\"type\":\"integer\"},\"flow_y\":{\"title\":\"/mavlink/topic/flow_y\",\"type\":\"integer\"},\"ground_distance\":{\"title\":\"/mavlink/topic/ground_distance\",\"type\":\"number\"},\"quality\":{\"title\":\"/mavlink/topic/quality\",\"type\":\"integer\"},\"sensor_id\":{\"title\":\"/mavlink/topic/sensor_id\",\"type\":\"integer\"},\"time_usec\":{\"title\":\"/mavlink/topic/time_usec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"covariance\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/covariance\",\"type\":\"array\"},\"pitchspeed\":{\"title\":\"/mavlink/topic/pitchspeed\",\"type\":\"number\"},\"q\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/q\",\"type\":\"array\"},\"rollspeed\":{\"title\":\"/mavlink/topic/rollspeed\",\"type\":\"number\"},\"time_usec\":{\"title\":\"/mavlink/topic/time_usec\",\"type\":\"integer\"},\"yawspeed\":{\"title\":\"/mavlink/topic/yawspeed\",\"type\":\"number\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"airspeed\":{\"title\":\"/mavlink/topic/airspeed\",\"type\":\"number\"},\"alt\":{\"title\":\"/mavlink/topic/alt\",\"type\":\"number\"},\"climb\":{\"title\":\"/mavlink/topic/climb\",\"type\":\"number\"},\"groundspeed\":{\"title\":\"/mavlink/topic/groundspeed\",\"type\":\"number\"},\"heading\":{\"title\":\"/mavlink/topic/heading\",\"type\":\"integer\"},\"throttle\":{\"title\":\"/mavlink/topic/throttle\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"hw_unique_id\":{\"title\":\"/mavlink/topic/hw_unique_id\",\"type\":\"string\"},\"hw_version_major\":{\"title\":\"/mavlink/topic/hw_version_major\",\"type\":\"integer\"},\"hw_version_minor\":{\"title\":\"/mavlink/topic/hw_version_minor\",\"type\":\"integer\"},\"name\":{\"title\":\"/mavlink/topic/name\",\"type\":\"string\"},\"sw_vcs_commit\":{\"title\":\"/mavlink/topic/sw_vcs_commit\",\"type\":\"integer\"},\"sw_version_major\":{\"title\":\"/mavlink/topic/sw_version_major\",\"type\":\"integer\"},\"sw_version_minor\":{\"title\":\"/mavlink/topic/sw_version_minor\",\"type\":\"integer\"},\"time_usec\":{\"title\":\"/mavlink/topic/time_usec\",\"type\":\"integer\"},\"uptime_sec\":{\"title\":\"/mavlink/topic/uptime_sec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"param_id\":{\"title\":\"/mavlink/topic/param_id\",\"type\":\"string\"},\"param_index\":{\"title\":\"/mavlink/topic/param_index\",\"type\":\"integer\"},\"target_component\":{\"title\":\"/mavlink/topic/target_component\",\"type\":\"integer\"},\"target_system\":{\"title\":\"/mavlink/topic/target_system\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"angle_offset\":{\"title\":\"/mavlink/topic/angle_offset\",\"type\":\"number\"},\"distances\":{\"items\":{\"type\":\"integer\"},\"title\":\"/mavlink/topic/distances\",\"type\":\"array\"},\"frame\":{\"title\":\"/mavlink/topic/frame\",\"type\":\"integer\"},\"increment\":{\"title\":\"/mavlink/topic/increment\",\"type\":\"integer\"},\"increment_f\":{\"title\":\"/mavlink/topic/increment_f\",\"type\":\"number\"},\"max_distance\":{\"title\":\"/mavlink/topic/max_distance\",\"type\":\"integer\"},\"min_distance\":{\"title\":\"/mavlink/topic/min_distance\",\"type\":\"integer\"},\"sensor_type\":{\"title\":\"/mavlink/topic/sensor_type\",\"type\":\"integer\"},\"time_usec\":{\"title\":\"/mavlink/topic/time_usec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"angle_off\":{\"title\":\"/mavlink/topic/angle_offset\",\"type\":\"number\"},\"distances\":{\"items\":{\"type\":\"integer\"},\"title\":\"/mavlink/topic/distances\",\"type\":\"array\"},\"frame\":{\"title\":\"/mavlink/topic/frame\",\"type\":\"integer\"},\"increment\":{\"title\":\"/mavlink/topic/increment\",\"type\":\"integer\"},\"increment_f\":{\"title\":\"/mavlink/topic/increment_f\",\"type\":\"number\"},\"max_distance\":{\"title\":\"/mavlink/topic/max_distance\",\"type\":\"integer\"},\"min_distance\":{\"title\":\"/mavlink/topic/min_distance\",\"type\":\"integer\"},\"sensor_type\":{\"title\":\"/mavlink/topic/sensor_type\",\"type\":\"integer\"},\"time_usec\":{\"title\":\"/mavlink/topic/time_usec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"angle_offset\":{\"title\":\"/mavlink/topic/angle_offset\",\"type\":\"number\"},\"distances\":{\"items\":{\"type\":\"integer\"},\"title\":\"/mavlink/topic/distances\",\"type\":\"array\"},\"frame\":{\"title\":\"/mavlink/topic/frame\",\"type\":\"integer\"},\"increment\":{\"title\":\"/mavlink/topic/increment\",\"type\":\"integer\"},\"increment_f\":{\"title\":\"/mavlink/topic/increment_f\",\"type\":\"number\"},\"max_distance\":{\"title\":\"/mavlink/topic/max_distance\",\"type\":\"integer\"},\"min_distance\":{\"title\":\"/mavlink/topic/min_distance\",\"type\":\"integer\"},\"sensor_type\":{\"title\":\"/mavlink/topic/sensor_type\",\"type\":\"integer\"},\"time_usec\":{\"title\":\"/mavlink/topic/time_usec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
}
var jsonTest = []string{
	"{\"acc_x\":[1,2,3,4,5],\"acc_y\":[1,2,3,4,5],\"acc_z\":[1,2,3,4,5],\"command\":[1,2,3,4,5],\"pos_x\":[1,2,3,4,5],\"pos_y\":[1,2,3,4,5],\"pos_yaw\":[1,2,3,4,5],\"pos_z\":[1,2,3,4,5],\"time_usec\":1,\"valid_points\":2,\"vel_x\":[1,2,3,4,5],\"vel_y\":[1,2,3,4,5],\"vel_yaw\":[1,2,3,4,5],\"vel_z\":[1,2,3,4,5]}",
	"{\"flow_comp_m_x\":1,\"flow_comp_m_y\":1,\"flow_rate_x\":1,\"flow_rate_y\":1,\"flow_x\":7,\"flow_y\":8,\"ground_distance\":1,\"quality\":10,\"sensor_id\":9,\"time_usec\":3}",
	"{\"covariance\":[1,1,1,1,1,1,1,1,1],\"pitchspeed\":1,\"q\":[1,1,1,1],\"rollspeed\":1,\"time_usec\":2,\"yawspeed\":1}",
	"{\"airspeed\":1234,\"alt\":1234,\"climb\":1234,\"groundspeed\":1234,\"heading\":12,\"throttle\":123}",
	"{\"hw_unique_id\":\"AQIDBAUGBwgJCgsMDQ4PEA==\",\"hw_version_major\":3,\"hw_version_minor\":4,\"name\":\"sapog.px4.io\",\"sw_vcs_commit\":7,\"sw_version_major\":5,\"sw_version_minor\":6,\"time_usec\":1,\"uptime_sec\":2}",
	"{\"param_id\":\"hopefullyWorks\",\"param_index\":333,\"target_component\":2,\"target_system\":1}",
	"{\"angle_offset\":0,\"distances\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72],\"frame\":0,\"increment\":36,\"increment_f\":0,\"max_distance\":50,\"min_distance\":1,\"sensor_type\":2,\"time_usec\":1}",
	"{\"angle_offset\":0,\"distances\":[3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72],\"frame\":0,\"increment\":36,\"increment_f\":0,\"max_distance\":50,\"min_distance\":1,\"sensor_type\":2,\"time_usec\":1}",
	"{\"angle_offset\":\"invalidString\",\"distances\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72],\"frame\":0,\"increment\":36,\"increment_f\":0,\"max_distance\":50,\"min_distance\":1,\"sensor_type\":2,\"time_usec\":1}",
}

// Function that is not exported so copy and pasted here for testing purposes.
func dialectMsgDefToGo(in string) string {
	re := regexp.MustCompile("_[a-z]")
	in = strings.ToLower(in)
	in = re.ReplaceAllStringFunc(in, func(match string) string {
		return strings.ToUpper(match[1:2])
	})
	return strings.ToUpper(in[:1]) + in[1:]
}

// Test case structs to loop through in testing above.
var casesMsgsTest = []struct {
	name   string
	isV2   bool
	parsed gomavlib.Message
	raw    []byte
	id     uint32
}{
	{
		"v1 message with array of enums",
		false,
		&MessageTrajectoryRepresentationWaypoints{
			TimeUsec:    1,
			ValidPoints: 2,
			PosX:        [5]float32{1, 2, 3, 4, 5},
			PosY:        [5]float32{1, 2, 3, 4, 5},
			PosZ:        [5]float32{1, 2, 3, 4, 5},
			VelX:        [5]float32{1, 2, 3, 4, 5},
			VelY:        [5]float32{1, 2, 3, 4, 5},
			VelZ:        [5]float32{1, 2, 3, 4, 5},
			AccX:        [5]float32{1, 2, 3, 4, 5},
			AccY:        [5]float32{1, 2, 3, 4, 5},
			AccZ:        [5]float32{1, 2, 3, 4, 5},
			PosYaw:      [5]float32{1, 2, 3, 4, 5},
			VelYaw:      [5]float32{1, 2, 3, 4, 5},
			Command:     [5]MAV_CMD{1, 2, 3, 4, 5},
		},
		[]byte("\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x00\x00\x80\x3f\x00\x00\x00\x40\x00\x00\x40\x40\x00\x00\x80\x40\x00\x00\xa0\x40\x01\x00\x02\x00\x03\x00\x04\x00\x05\x00\x02"),
		332,
	},
	{
		"v2 message with extensions",
		true,
		&MessageOpticalFlow{
			TimeUsec:       3,
			FlowCompMX:     1,
			FlowCompMY:     1,
			GroundDistance: 1,
			FlowX:          7,
			FlowY:          8,
			SensorId:       9,
			Quality:        0x0A,
			FlowRateX:      1,
			FlowRateY:      1,
		},
		[]byte("\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80\x3F\x00\x00\x80\x3F\x00\x00\x80\x3F\x07\x00\x08\x00\x09\x0A\x00\x00\x80\x3F\x00\x00\x80\x3F"),
		100,
	},
	{
		"v1 message with array",
		false,
		&MessageAttitudeQuaternionCov{
			TimeUsec:   2,
			Q:          [4]float32{1, 1, 1, 1},
			Rollspeed:  1,
			Pitchspeed: 1,
			Yawspeed:   1,
			Covariance: [9]float32{1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		append([]byte("\x02\x00\x00\x00\x00\x00\x00\x00"), bytes.Repeat([]byte("\x00\x00\x80\x3F"), 16)...),
		61,
	},
	{
		"V2 message 74",
		true,
		&MessageVfrHud{
			Airspeed:    1234,
			Groundspeed: 1234,
			Heading:     12,
			Throttle:    123,
			Alt:         1234,
			Climb:       1234,
		},
		[]byte("\x00\x40\x9A\x44\x00\x40\x9A\x44\x00\x40\x9A\x44\x00\x40\x9A\x44\x0C\x00\x7B"),
		74,
	},
	{
		"V2 Message 311 with string",
		true,
		&MessageUavcanNodeInfo{
			TimeUsec:       1,
			UptimeSec:      2,
			Name:           "sapog.px4.io",
			HwVersionMajor: 3,
			HwVersionMinor: 4,
			HwUniqueId:     [16]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			SwVersionMajor: 5,
			SwVersionMinor: 6,
			SwVcsCommit:    7,
		},
		[]byte("\x01\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x07\x00\x00\x00\x73\x61\x70\x6F\x67\x2E\x70\x78\x34\x2E\x69\x6F\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x04\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0A\x0B\x0C\x0D\x0E\x0F\x10\x05\x06"),
		311,
	},
	{
		"V2 Message 320 with string",
		true,
		&MessageParamExtRequestRead{
			TargetSystem:    1,
			TargetComponent: 2,
			ParamId:         "hopefullyWorks",
			ParamIndex:      333,
		},
		[]byte("\x4D\x01\x01\x02\x68\x6F\x70\x65\x66\x75\x6C\x6C\x79\x57\x6F\x72\x6B\x73"),
		320,
	},
	{
		"V2 Message 330 with array of uint16",
		true,
		&MessageObstacleDistance{
			TimeUsec:    1,
			SensorType:  2,
			Distances:   [72]uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72},
			Increment:   36,
			MinDistance: 1,
			MaxDistance: 50,
			IncrementF:  0,
			AngleOffset: 0,
			Frame:       0,
		},
		[]byte("\x01\x00\x00\x00\x00\x00\x00\x00\x01\x00\x02\x00\x03\x00\x04\x00\x05\x00\x06\x00\x07\x00\x08\x00\x09\x00\x0A\x00\x0B\x00\x0C\x00\x0D\x00\x0E\x00\x0F\x00\x10\x00\x11\x00\x12\x00\x13\x00\x14\x00\x15\x00\x16\x00\x17\x00\x18\x00\x19\x00\x1A\x00\x1B\x00\x1C\x00\x1D\x00\x1E\x00\x1F\x00\x20\x00\x21\x00\x22\x00\x23\x00\x24\x00\x25\x00\x26\x00\x27\x00\x28\x00\x29\x00\x2A\x00\x2B\x00\x2C\x00\x2D\x00\x2E\x00\x2F\x00\x30\x00\x31\x00\x32\x00\x33\x00\x34\x00\x35\x00\x36\x00\x37\x00\x38\x00\x39\x00\x3A\x00\x3B\x00\x3C\x00\x3D\x00\x3E\x00\x3F\x00\x40\x00\x41\x00\x42\x00\x43\x00\x44\x00\x45\x00\x46\x00\x47\x00\x48\x00\x01\x00\x32\x00\x02\x24"),
		330,
	},
	{
		"V2 Message 330 with array of uint16 - Schema and JSON given should fail for this test case!",
		true,
		&MessageObstacleDistance{
			TimeUsec:    1,
			SensorType:  2,
			Distances:   [72]uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72},
			Increment:   36,
			MinDistance: 1,
			MaxDistance: 50,
			IncrementF:  0,
			AngleOffset: 0,
			Frame:       0,
		},
		[]byte("\x01\x00\x00\x00\x00\x00\x00\x00\x01\x00\x02\x00\x03\x00\x04\x00\x05\x00\x06\x00\x07\x00\x08\x00\x09\x00\x0A\x00\x0B\x00\x0C\x00\x0D\x00\x0E\x00\x0F\x00\x10\x00\x11\x00\x12\x00\x13\x00\x14\x00\x15\x00\x16\x00\x17\x00\x18\x00\x19\x00\x1A\x00\x1B\x00\x1C\x00\x1D\x00\x1E\x00\x1F\x00\x20\x00\x21\x00\x22\x00\x23\x00\x24\x00\x25\x00\x26\x00\x27\x00\x28\x00\x29\x00\x2A\x00\x2B\x00\x2C\x00\x2D\x00\x2E\x00\x2F\x00\x30\x00\x31\x00\x32\x00\x33\x00\x34\x00\x35\x00\x36\x00\x37\x00\x38\x00\x39\x00\x3A\x00\x3B\x00\x3C\x00\x3D\x00\x3E\x00\x3F\x00\x40\x00\x41\x00\x42\x00\x43\x00\x44\x00\x45\x00\x46\x00\x47\x00\x48\x00\x01\x00\x32\x00\x02\x24"),
		330,
	},
	{
		"V2 Message 330 with array of uint16 - JSONTest[8] should not validate to schemasTest[8] for this test case!",
		true,
		&MessageObstacleDistance{
			TimeUsec:    1,
			SensorType:  2,
			Distances:   [72]uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72},
			Increment:   36,
			MinDistance: 1,
			MaxDistance: 50,
			IncrementF:  0,
			AngleOffset: 0,
			Frame:       0,
		},
		[]byte("\x01\x00\x00\x00\x00\x00\x00\x00\x01\x00\x02\x00\x03\x00\x04\x00\x05\x00\x06\x00\x07\x00\x08\x00\x09\x00\x0A\x00\x0B\x00\x0C\x00\x0D\x00\x0E\x00\x0F\x00\x10\x00\x11\x00\x12\x00\x13\x00\x14\x00\x15\x00\x16\x00\x17\x00\x18\x00\x19\x00\x1A\x00\x1B\x00\x1C\x00\x1D\x00\x1E\x00\x1F\x00\x20\x00\x21\x00\x22\x00\x23\x00\x24\x00\x25\x00\x26\x00\x27\x00\x28\x00\x29\x00\x2A\x00\x2B\x00\x2C\x00\x2D\x00\x2E\x00\x2F\x00\x30\x00\x31\x00\x32\x00\x33\x00\x34\x00\x35\x00\x36\x00\x37\x00\x38\x00\x39\x00\x3A\x00\x3B\x00\x3C\x00\x3D\x00\x3E\x00\x3F\x00\x40\x00\x41\x00\x42\x00\x43\x00\x44\x00\x45\x00\x46\x00\x47\x00\x48\x00\x01\x00\x32\x00\x02\x24"),
		330,
	},
}
var ctMessages = []gomavlib.Message{
	// common.xml
	&MessageHeartbeat{},
	&MessageSysStatus{},
	&MessageSystemTime{},
	&MessagePing{},
	&MessageChangeOperatorControl{},
	&MessageChangeOperatorControlAck{},
	&MessageAuthKey{},
	&MessageLinkNodeStatus{},
	&MessageSetMode{},
	&MessageParamRequestRead{},
	&MessageParamRequestList{},
	&MessageParamValue{},
	&MessageParamSet{},
	&MessageGpsRawInt{},
	&MessageGpsStatus{},
	&MessageScaledImu{},
	&MessageRawImu{},
	&MessageRawPressure{},
	&MessageScaledPressure{},
	&MessageAttitude{},
	&MessageAttitudeQuaternion{},
	&MessageLocalPositionNed{},
	&MessageGlobalPositionInt{},
	&MessageRcChannelsScaled{},
	&MessageRcChannelsRaw{},
	&MessageServoOutputRaw{},
	&MessageMissionRequestPartialList{},
	&MessageMissionWritePartialList{},
	&MessageMissionItem{},
	&MessageMissionRequest{},
	&MessageMissionSetCurrent{},
	&MessageMissionCurrent{},
	&MessageMissionRequestList{},
	&MessageMissionCount{},
	&MessageMissionClearAll{},
	&MessageMissionItemReached{},
	&MessageMissionAck{},
	&MessageSetGpsGlobalOrigin{},
	&MessageGpsGlobalOrigin{},
	&MessageParamMapRc{},
	&MessageMissionRequestInt{},
	&MessageMissionChanged{},
	&MessageSafetySetAllowedArea{},
	&MessageSafetyAllowedArea{},
	&MessageAttitudeQuaternionCov{},
	&MessageNavControllerOutput{},
	&MessageGlobalPositionIntCov{},
	&MessageLocalPositionNedCov{},
	&MessageRcChannels{},
	&MessageRequestDataStream{},
	&MessageDataStream{},
	&MessageManualControl{},
	&MessageRcChannelsOverride{},
	&MessageMissionItemInt{},
	&MessageVfrHud{},
	&MessageCommandInt{},
	&MessageCommandLong{},
	&MessageCommandAck{},
	&MessageManualSetpoint{},
	&MessageSetAttitudeTarget{},
	&MessageAttitudeTarget{},
	&MessageSetPositionTargetLocalNed{},
	&MessagePositionTargetLocalNed{},
	&MessageSetPositionTargetGlobalInt{},
	&MessagePositionTargetGlobalInt{},
	&MessageLocalPositionNedSystemGlobalOffset{},
	&MessageHilState{},
	&MessageHilControls{},
	&MessageHilRcInputsRaw{},
	&MessageHilActuatorControls{},
	&MessageOpticalFlow{},
	&MessageGlobalVisionPositionEstimate{},
	&MessageVisionPositionEstimate{},
	&MessageVisionSpeedEstimate{},
	&MessageViconPositionEstimate{},
	&MessageHighresImu{},
	&MessageOpticalFlowRad{},
	&MessageHilSensor{},
	&MessageSimState{},
	&MessageRadioStatus{},
	&MessageFileTransferProtocol{},
	&MessageTimesync{},
	&MessageCameraTrigger{},
	&MessageHilGps{},
	&MessageHilOpticalFlow{},
	&MessageHilStateQuaternion{},
	&MessageScaledImu2{},
	&MessageLogRequestList{},
	&MessageLogEntry{},
	&MessageLogRequestData{},
	&MessageLogData{},
	&MessageLogErase{},
	&MessageLogRequestEnd{},
	&MessageGpsInjectData{},
	&MessageGps2Raw{},
	&MessagePowerStatus{},
	&MessageSerialControl{},
	&MessageGpsRtk{},
	&MessageGps2Rtk{},
	&MessageScaledImu3{},
	&MessageDataTransmissionHandshake{},
	&MessageEncapsulatedData{},
	&MessageDistanceSensor{},
	&MessageTerrainRequest{},
	&MessageTerrainData{},
	&MessageTerrainCheck{},
	&MessageTerrainReport{},
	&MessageScaledPressure2{},
	&MessageAttPosMocap{},
	&MessageSetActuatorControlTarget{},
	&MessageActuatorControlTarget{},
	&MessageAltitude{},
	&MessageResourceRequest{},
	&MessageScaledPressure3{},
	&MessageFollowTarget{},
	&MessageControlSystemState{},
	&MessageBatteryStatus{},
	&MessageAutopilotVersion{},
	&MessageLandingTarget{},
	&MessageFenceStatus{},
	&MessageEstimatorStatus{},
	&MessageWindCov{},
	&MessageGpsInput{},
	&MessageGpsRtcmData{},
	&MessageHighLatency{},
	&MessageHighLatency2{},
	&MessageVibration{},
	&MessageHomePosition{},
	&MessageSetHomePosition{},
	&MessageMessageInterval{},
	&MessageExtendedSysState{},
	&MessageAdsbVehicle{},
	&MessageCollision{},
	&MessageV2Extension{},
	&MessageMemoryVect{},
	&MessageDebugVect{},
	&MessageNamedValueFloat{},
	&MessageNamedValueInt{},
	&MessageStatustext{},
	&MessageDebug{},
	&MessageSetupSigning{},
	&MessageButtonChange{},
	&MessagePlayTune{},
	&MessageCameraInformation{},
	&MessageCameraSettings{},
	&MessageStorageInformation{},
	&MessageCameraCaptureStatus{},
	&MessageCameraImageCaptured{},
	&MessageFlightInformation{},
	&MessageMountOrientation{},
	&MessageLoggingData{},
	&MessageLoggingDataAcked{},
	&MessageLoggingAck{},
	&MessageVideoStreamInformation{},
	&MessageVideoStreamStatus{},
	&MessageGimbalManagerInformation{},
	&MessageGimbalManagerStatus{},
	&MessageGimbalManagerSetAttitude{},
	&MessageGimbalDeviceInformation{},
	&MessageGimbalDeviceSetAttitude{},
	&MessageGimbalDeviceAttitudeStatus{},
	&MessageAutopilotStateForGimbalDevice{},
	&MessageWifiConfigAp{},
	&MessageProtocolVersion{},
	&MessageAisVessel{},
	&MessageUavcanNodeStatus{},
	&MessageUavcanNodeInfo{},
	&MessageParamExtRequestRead{},
	&MessageParamExtRequestList{},
	&MessageParamExtValue{},
	&MessageParamExtSet{},
	&MessageParamExtAck{},
	&MessageObstacleDistance{},
	&MessageOdometry{},
	&MessageTrajectoryRepresentationWaypoints{},
	&MessageTrajectoryRepresentationBezier{},
	&MessageCellularStatus{},
	&MessageIsbdLinkStatus{},
	&MessageUtmGlobalPosition{},
	&MessageDebugFloatArray{},
	&MessageOrbitExecutionStatus{},
	&MessageSmartBatteryInfo{},
	&MessageSmartBatteryStatus{},
	&MessageActuatorOutputStatus{},
	&MessageTimeEstimateToTarget{},
	&MessageTunnel{},
	&MessageOnboardComputerStatus{},
	&MessageComponentInformation{},
	&MessagePlayTuneV2{},
	&MessageSupportedTunes{},
	&MessageWheelDistance{},
	&MessageOpenDroneIdBasicId{},
	&MessageOpenDroneIdLocation{},
	&MessageOpenDroneIdAuthentication{},
	&MessageOpenDroneIdSelfId{},
	&MessageOpenDroneIdSystem{},
	&MessageOpenDroneIdOperatorId{},
	&MessageOpenDroneIdMessagePack{},
}

// DEFINE PUBLIC RECEIVER FUNCTIONS.
// DEFINE PRIVATE STATIC FUNCTIONS.
// DEFINE PRIVATE RECEIVER FUNCTIONS.
// ALL DONE.
