package testgomavlib

import (
	"bytes"
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
func CreateMessageByIdTest(t *testing.T) {
	includeDirs := []string{"../mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("../mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml", includeDirs)
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

// CreateMessageByIdTest creates a dynamic message based on the input name and checks that the values within it are valid.
func CreateMessageByNameTest(t *testing.T) {
	includeDirs := []string{"../mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("../mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml", includeDirs)
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
func JSONMarshalAndUnmarshalTest(t *testing.T) {
	for i, c := range casesMsgsTest {
		dCT, err := gomavlib.NewDialectCT(3, ctMessages)
		require.NoError(t, err)
		dMsgCT, ok := dCT.Messages[c.id]
		require.Equal(t, true, ok)
		bytesEncoded, err := dMsgCT.Encode(c.parsed, c.isV2)
		require.NoError(t, err)
		require.Equal(t, c.raw, bytesEncoded)

		// Decode bytes using RT
		includeDirs := []string{"../mavlink-upstream/message_definitions/v1.0"}
		defs, version, err := libgen.XMLToFields("../mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml", includeDirs)
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
		require.Equal(t, msgDecoded.(*gomavlib.DynamicMessage).Fields, dm.Fields)
	}
}

// DialectRTCommonXMLTest tests the XMLToFields and RT dialect generation functionality added to gomavlib.
func DialectRTCommonXMLTest(t *testing.T) {
	// Parse the XML file.
	includeDirs := []string{"../mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("../mavlink-upstream/message_definitions/v1.0/common.xml", includeDirs)
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
func DecodeAndEncodeRTTest(t *testing.T) {
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
		includeDirs := []string{"../mavlink-upstream/message_definitions/v1.0"}
		defs, version, err := libgen.XMLToFields("../mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml", includeDirs)
		require.NoError(t, err)

		// Create dialect from the parsed defs.
		dRT, err := gomavlib.NewDialectRT(version, defs)
		dMsgRT := dRT.Messages[c.id]
		require.NoError(t, err)
		require.Equal(t, uint(3), dRT.GetVersion())

		// Decode bytes using RT
		msgDecoded, err := dMsgRT.Decode(bytesEncoded, c.isV2)
		require.NoError(t, err)

		// TODO: Implement the code below properly
		//Make sure all fields of dMsgCT match equivalent values of RT msgDecoded
		//Compare all fields of all RT and CT Messages
		// v := reflect.ValueOf(msgDecodedCT).Elem()
		// for i := 0; i < len(dMsgCT.fields); i++ {
		// 	fCT := dMsgCT.fields[i]
		// 	OriginalName := dialectMsgDefToGo(fCT.name)
		// 	fRT := msgDecoded.(*DynamicMessage).Fields[OriginalName]
		// 	fCTVal := v.FieldByName(OriginalName).Interface()
		// 	fmt.Println(fRT)
		// 	fmt.Println(fCTVal)
		// }

		// Encode using RT
		bytesEncodedByRT, err := dMsgRT.Encode(msgDecoded, c.isV2)
		require.NoError(t, err)
		require.Equal(t, c.raw, bytesEncodedByRT)
	}
}

var schemasTest = []string{
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"AccX\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/AccX\",\"type\":\"array\"},\"AccY\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/AccY\",\"type\":\"array\"},\"AccZ\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/AccZ\",\"type\":\"array\"},\"Command\":{\"items\":{\"type\":\"integer\"},\"title\":\"/mavlink/topic/Command\",\"type\":\"array\"},\"PosX\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/PosX\",\"type\":\"array\"},\"PosY\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/PosY\",\"type\":\"array\"},\"PosYaw\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/PosYaw\",\"type\":\"array\"},\"PosZ\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/PosZ\",\"type\":\"array\"},\"TimeUsec\":{\"title\":\"/mavlink/topic/TimeUsec\",\"type\":\"integer\"},\"ValidPoints\":{\"title\":\"/mavlink/topic/ValidPoints\",\"type\":\"integer\"},\"VelX\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/VelX\",\"type\":\"array\"},\"VelY\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/VelY\",\"type\":\"array\"},\"VelYaw\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/VelYaw\",\"type\":\"array\"},\"VelZ\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/VelZ\",\"type\":\"array\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"FlowCompMX\":{\"title\":\"/mavlink/topic/FlowCompMX\",\"type\":\"number\"},\"FlowCompMY\":{\"title\":\"/mavlink/topic/FlowCompMY\",\"type\":\"number\"},\"FlowRateX\":{\"title\":\"/mavlink/topic/FlowRateX\",\"type\":\"number\"},\"FlowRateY\":{\"title\":\"/mavlink/topic/FlowRateY\",\"type\":\"number\"},\"FlowX\":{\"title\":\"/mavlink/topic/FlowX\",\"type\":\"integer\"},\"FlowY\":{\"title\":\"/mavlink/topic/FlowY\",\"type\":\"integer\"},\"GroundDistance\":{\"title\":\"/mavlink/topic/GroundDistance\",\"type\":\"number\"},\"Quality\":{\"title\":\"/mavlink/topic/Quality\",\"type\":\"integer\"},\"SensorId\":{\"title\":\"/mavlink/topic/SensorId\",\"type\":\"integer\"},\"TimeUsec\":{\"title\":\"/mavlink/topic/TimeUsec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"Covariance\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/Covariance\",\"type\":\"array\"},\"Pitchspeed\":{\"title\":\"/mavlink/topic/Pitchspeed\",\"type\":\"number\"},\"Q\":{\"items\":{\"type\":\"number\"},\"title\":\"/mavlink/topic/Q\",\"type\":\"array\"},\"Rollspeed\":{\"title\":\"/mavlink/topic/Rollspeed\",\"type\":\"number\"},\"TimeUsec\":{\"title\":\"/mavlink/topic/TimeUsec\",\"type\":\"integer\"},\"Yawspeed\":{\"title\":\"/mavlink/topic/Yawspeed\",\"type\":\"number\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"Airspeed\":{\"title\":\"/mavlink/topic/Airspeed\",\"type\":\"number\"},\"Alt\":{\"title\":\"/mavlink/topic/Alt\",\"type\":\"number\"},\"Climb\":{\"title\":\"/mavlink/topic/Climb\",\"type\":\"number\"},\"Groundspeed\":{\"title\":\"/mavlink/topic/Groundspeed\",\"type\":\"number\"},\"Heading\":{\"title\":\"/mavlink/topic/Heading\",\"type\":\"integer\"},\"Throttle\":{\"title\":\"/mavlink/topic/Throttle\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"HwUniqueId\":{\"title\":\"/mavlink/topic/HwUniqueId\",\"type\":\"string\"},\"HwVersionMajor\":{\"title\":\"/mavlink/topic/HwVersionMajor\",\"type\":\"integer\"},\"HwVersionMinor\":{\"title\":\"/mavlink/topic/HwVersionMinor\",\"type\":\"integer\"},\"Name\":{\"title\":\"/mavlink/topic/Name\",\"type\":\"string\"},\"SwVcsCommit\":{\"title\":\"/mavlink/topic/SwVcsCommit\",\"type\":\"integer\"},\"SwVersionMajor\":{\"title\":\"/mavlink/topic/SwVersionMajor\",\"type\":\"integer\"},\"SwVersionMinor\":{\"title\":\"/mavlink/topic/SwVersionMinor\",\"type\":\"integer\"},\"TimeUsec\":{\"title\":\"/mavlink/topic/TimeUsec\",\"type\":\"integer\"},\"UptimeSec\":{\"title\":\"/mavlink/topic/UptimeSec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"ParamId\":{\"title\":\"/mavlink/topic/ParamId\",\"type\":\"string\"},\"ParamIndex\":{\"title\":\"/mavlink/topic/ParamIndex\",\"type\":\"integer\"},\"TargetComponent\":{\"title\":\"/mavlink/topic/TargetComponent\",\"type\":\"integer\"},\"TargetSystem\":{\"title\":\"/mavlink/topic/TargetSystem\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"AngleOffset\":{\"title\":\"/mavlink/topic/AngleOffset\",\"type\":\"number\"},\"Distances\":{\"items\":{\"type\":\"integer\"},\"title\":\"/mavlink/topic/Distances\",\"type\":\"array\"},\"Frame\":{\"title\":\"/mavlink/topic/Frame\",\"type\":\"integer\"},\"Increment\":{\"title\":\"/mavlink/topic/Increment\",\"type\":\"integer\"},\"IncrementF\":{\"title\":\"/mavlink/topic/IncrementF\",\"type\":\"number\"},\"MaxDistance\":{\"title\":\"/mavlink/topic/MaxDistance\",\"type\":\"integer\"},\"MinDistance\":{\"title\":\"/mavlink/topic/MinDistance\",\"type\":\"integer\"},\"SensorType\":{\"title\":\"/mavlink/topic/SensorType\",\"type\":\"integer\"},\"TimeUsec\":{\"title\":\"/mavlink/topic/TimeUsec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"AngleOff\":{\"title\":\"/mavlink/topic/AngleOffset\",\"type\":\"number\"},\"Distances\":{\"items\":{\"type\":\"integer\"},\"title\":\"/mavlink/topic/Distances\",\"type\":\"array\"},\"Frame\":{\"title\":\"/mavlink/topic/Frame\",\"type\":\"integer\"},\"Increment\":{\"title\":\"/mavlink/topic/Increment\",\"type\":\"integer\"},\"IncrementF\":{\"title\":\"/mavlink/topic/IncrementF\",\"type\":\"number\"},\"MaxDistance\":{\"title\":\"/mavlink/topic/MaxDistance\",\"type\":\"integer\"},\"MinDistance\":{\"title\":\"/mavlink/topic/MinDistance\",\"type\":\"integer\"},\"SensorType\":{\"title\":\"/mavlink/topic/SensorType\",\"type\":\"integer\"},\"TimeUsec\":{\"title\":\"/mavlink/topic/TimeUsec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
	"{\"$id\":\"/mavlink/topic\",\"$schema\":\"https://json-schema.org/draft-07/schema#\",\"properties\":{\"AngleOffset\":{\"title\":\"/mavlink/topic/AngleOffset\",\"type\":\"number\"},\"Distances\":{\"items\":{\"type\":\"integer\"},\"title\":\"/mavlink/topic/Distances\",\"type\":\"array\"},\"Frame\":{\"title\":\"/mavlink/topic/Frame\",\"type\":\"integer\"},\"Increment\":{\"title\":\"/mavlink/topic/Increment\",\"type\":\"integer\"},\"IncrementF\":{\"title\":\"/mavlink/topic/IncrementF\",\"type\":\"number\"},\"MaxDistance\":{\"title\":\"/mavlink/topic/MaxDistance\",\"type\":\"integer\"},\"MinDistance\":{\"title\":\"/mavlink/topic/MinDistance\",\"type\":\"integer\"},\"SensorType\":{\"title\":\"/mavlink/topic/SensorType\",\"type\":\"integer\"},\"TimeUsec\":{\"title\":\"/mavlink/topic/TimeUsec\",\"type\":\"integer\"}},\"title\":\"/mavlink/topic\",\"type\":\"object\"}",
}
var jsonTest = []string{
	"{\"AccX\":[1,2,3,4,5],\"AccY\":[1,2,3,4,5],\"AccZ\":[1,2,3,4,5],\"Command\":[1,2,3,4,5],\"PosX\":[1,2,3,4,5],\"PosY\":[1,2,3,4,5],\"PosYaw\":[1,2,3,4,5],\"PosZ\":[1,2,3,4,5],\"TimeUsec\":1,\"ValidPoints\":2,\"VelX\":[1,2,3,4,5],\"VelY\":[1,2,3,4,5],\"VelYaw\":[1,2,3,4,5],\"VelZ\":[1,2,3,4,5]}",
	"{\"FlowCompMX\":1,\"FlowCompMY\":1,\"FlowRateX\":1,\"FlowRateY\":1,\"FlowX\":7,\"FlowY\":8,\"GroundDistance\":1,\"Quality\":10,\"SensorId\":9,\"TimeUsec\":3}",
	"{\"Covariance\":[1,1,1,1,1,1,1,1,1],\"Pitchspeed\":1,\"Q\":[1,1,1,1],\"Rollspeed\":1,\"TimeUsec\":2,\"Yawspeed\":1}",
	"{\"Airspeed\":1234,\"Alt\":1234,\"Climb\":1234,\"Groundspeed\":1234,\"Heading\":12,\"Throttle\":123}",
	"{\"HwUniqueId\":\"AQIDBAUGBwgJCgsMDQ4PEA==\",\"HwVersionMajor\":3,\"HwVersionMinor\":4,\"Name\":\"sapog.px4.io\",\"SwVcsCommit\":7,\"SwVersionMajor\":5,\"SwVersionMinor\":6,\"TimeUsec\":1,\"UptimeSec\":2}",
	"{\"ParamId\":\"hopefullyWorks\",\"ParamIndex\":333,\"TargetComponent\":2,\"TargetSystem\":1}",
	"{\"AngleOffset\":0,\"Distances\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72],\"Frame\":0,\"Increment\":36,\"IncrementF\":0,\"MaxDistance\":50,\"MinDistance\":1,\"SensorType\":2,\"TimeUsec\":1}",
	"{\"AngleOffset\":0,\"Distances\":[3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72],\"Frame\":0,\"Increment\":36,\"IncrementF\":0,\"MaxDistance\":50,\"MinDistance\":1,\"SensorType\":2,\"TimeUsec\":1}",
	"{\"AngleOffset\":\"invalidString\",\"Distances\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72],\"Frame\":0,\"Increment\":36,\"IncrementF\":0,\"MaxDistance\":50,\"MinDistance\":1,\"SensorType\":2,\"TimeUsec\":1}",
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
