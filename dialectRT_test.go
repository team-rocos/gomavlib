package gomavlib

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	libgen "github.com/team-rocos/gomavlib/commands/dialgen/libgen"
)

// DEFINE PUBLIC TYPES AND STRUCTURES.

// DEFINE PRIVATE TYPES AND STRUCTURES.
// DEFINE PUBLIC STATIC FUNCTIONS.
func TestCreateMessageById(t *testing.T) {
	includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("./mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml", includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := NewDialectRT(version, defs)
	require.NoError(t, err)

	// Create dynamic message using id of each message in dRT
	for _, mRT := range dRT.messages {
		dm, err := dRT.CreateMessageById(uint32(mRT.msg.Id))
		require.NoError(t, err)
		require.Equal(t, mRT, dm.t)
	}

	// CreateMessageById using invalid id. Assert that error is returned
	_, err = dRT.CreateMessageById(40000000)
	assert.Error(t, err)
}

func TestCreateMessageByName(t *testing.T) {
	includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("./mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml", includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := NewDialectRT(version, defs)
	require.NoError(t, err)

	// Create dynamic message by name using name from each mRT in dRT
	for _, mRT := range dRT.messages {
		dm, err := dRT.CreateMessageByName(mRT.msg.Name)
		require.NoError(t, err)
		require.Equal(t, mRT, dm.t)
	}

	// Create dynamic message using invalid name. Assert that error is returned
	_, err = dRT.CreateMessageByName("abcdefghijklmnop***")
	assert.Error(t, err)
}

func TestJSONMarshal(t *testing.T) {
	for i, c := range casesMsgsTest {
		dMsgCT, err := newDialectMessage(c.parsed)
		require.NoError(t, err)
		bytesEncoded, err := dMsgCT.encode(c.parsed, c.isV2)
		require.NoError(t, err)
		require.Equal(t, c.raw, bytesEncoded)

		// Decode bytes using RT
		includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
		defs, version, err := libgen.XMLToFields("./mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml", includeDirs)
		require.NoError(t, err)

		fmt.Println(i, "...................................................................")

		// Create dialect from the parsed defs.
		dRT, err := NewDialectRT(version, defs)
		dMsgRT := dRT.messages[c.id]
		require.NoError(t, err)
		require.Equal(t, uint(3), dRT.getVersion())

		// Decode bytes using RT
		msgDecoded, err := dMsgRT.decode(c.raw, c.isV2)
		require.NoError(t, err)

		bytesCreated, err := msgDecoded.(*DynamicMessage).MarshalJSON()
		require.NoError(t, err)
		fmt.Println(string(bytesCreated))
		fmt.Println()
		schemaBytes, err := msgDecoded.(*DynamicMessage).GenerateJSONSchema("/mavlink", "/topic")
		require.NoError(t, err)
		fmt.Println(string(schemaBytes))
	}
}

func TestDialectRTCommonXML(t *testing.T) {
	// Parse the XML file.
	includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("./mavlink-upstream/message_definitions/v1.0/common.xml", includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := NewDialectRT(version, defs)
	require.NoError(t, err)
	require.Equal(t, uint(3), dRT.getVersion())

	// Check Individual Messages for RT
	// Checking Message 5
	msg := dRT.messages[5].msg
	require.Equal(t, "ChangeOperatorControl", msg.Name)
	require.Equal(t, 5, msg.Id)
	field := msg.Fields[0]
	require.Equal(t, "TargetSystem", field.Name)
	require.Equal(t, "uint8", field.Type)
	field = msg.Fields[3]
	require.Equal(t, "Passkey", field.Name)
	require.Equal(t, "string", field.Type)

	// Checking Message 82 - Has float[4] array as a field
	msg = dRT.messages[82].msg
	require.Equal(t, "SetAttitudeTarget", msg.Name)
	require.Equal(t, 82, msg.Id)
	field = msg.Fields[1]
	require.Equal(t, "Q", field.Name)
	require.Equal(t, "float32", field.Type)

	// Compare with DialectCT
	dCT, err := NewDialectCT(3, CTMessages)
	require.NoError(t, err)

	require.Equal(t, len(dCT.messages), len(dRT.messages))

	// Compare RT and CT for all messages
	for _, m := range CTMessages {
		index := m.GetId()
		// Compare dCT with dRT
		mCT := dCT.messages[index]
		mRT := dRT.messages[index]
		require.Equal(t, mCT.sizeNormal, byte(mRT.sizeNormal))
		require.Equal(t, mCT.sizeExtended, byte(mRT.sizeExtended))
		require.Equal(t, mCT.crcExtra, mRT.crcExtra)

		// Compare all fields of all RT and CT Messages
		for i := 0; i < len(mCT.fields); i++ {
			fCT := mCT.fields[i]
			fRT := mRT.msg.Fields[i]
			require.Equal(t, fCT.isEnum, fRT.IsEnum)
			require.Equal(t, fCT.ftype, dialectFieldTypeFromGo[fRT.Type])
			require.Equal(t, fCT.name, fRT.OriginalName)
			require.Equal(t, fCT.arrayLength, byte(fRT.ArrayLength))
			require.Equal(t, fCT.index, fRT.Index)
			require.Equal(t, fCT.isExtension, fRT.IsExtension)
		}
	}
}

func TestDecodeAndEncodeRT0(t *testing.T) {
	// Encode using CT
	c := casesMsgsTest[0]
	dMsgCT, err := newDialectMessage(c.parsed)
	require.NoError(t, err)
	bytesEncoded, err := dMsgCT.encode(c.parsed, c.isV2)
	require.NoError(t, err)
	require.Equal(t, c.raw, bytesEncoded)

	// Decode bytes using RT
	includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("./mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml", includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := NewDialectRT(version, defs)
	dMsgRT := dRT.messages[332]
	require.NoError(t, err)
	require.Equal(t, uint(3), dRT.getVersion())

	// Decode bytes using RT
	msgDecoded, err := dMsgRT.decode(bytesEncoded, c.isV2)
	require.NoError(t, err)
	fmt.Printf("%+v\n", msgDecoded)

	// Try encoding using RT
	bytesEncodedByRT, err := dMsgRT.encode(msgDecoded, c.isV2)
	require.NoError(t, err)
	fmt.Println(bytesEncodedByRT)
	require.Equal(t, c.raw, bytesEncodedByRT)
}

func TestDecodeAndEncodeRT1(t *testing.T) {
	// Encode using CT
	c := casesMsgsTest[1]
	dMsgCT, err := newDialectMessage(c.parsed)
	require.NoError(t, err)
	bytesEncoded, err := dMsgCT.encode(c.parsed, c.isV2)
	require.NoError(t, err)
	require.Equal(t, c.raw, bytesEncoded)

	// Decode bytes using RT
	includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("TestingXMLFiles/testing1.xml", includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := NewDialectRT(version, defs)
	dMsgRT := dRT.messages[100]
	require.NoError(t, err)
	require.Equal(t, uint(3), dRT.getVersion())

	// Decode Bytes using RT
	msgDecoded, err := dMsgRT.decode(bytesEncoded, c.isV2)
	require.NoError(t, err)
	fmt.Printf("%+v\n", msgDecoded)

	// Try encoding using RT
	bytesEncodedByRT, err := dMsgRT.encode(msgDecoded, c.isV2)
	require.NoError(t, err)
	fmt.Println(bytesEncodedByRT)
	require.Equal(t, c.raw, bytesEncodedByRT)
}

func TestDecodeAndEncodeRT2(t *testing.T) {
	// Encode using CT
	c := casesMsgsTest[2]
	dMsgCT, err := newDialectMessage(c.parsed)
	require.NoError(t, err)
	bytesEncoded, err := dMsgCT.encode(c.parsed, c.isV2)
	require.NoError(t, err)
	require.Equal(t, c.raw, bytesEncoded)

	// Decode bytes using RT
	includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("TestingXMLFiles/testing2.xml", includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := NewDialectRT(version, defs)
	dMsgRT := dRT.messages[61]
	require.NoError(t, err)
	require.Equal(t, uint(3), dRT.getVersion())

	// Decode Bytes using RT
	msgDecoded, err := dMsgRT.decode(bytesEncoded, c.isV2)
	require.NoError(t, err)
	fmt.Printf("%+v\n", msgDecoded)

	// Try encoding using RT
	bytesEncodedByRT, err := dMsgRT.encode(msgDecoded, c.isV2)
	require.NoError(t, err)
	fmt.Println(bytesEncodedByRT)
	require.Equal(t, c.raw, bytesEncodedByRT)
}

func TestDecodeAndEncodeRT3(t *testing.T) {
	// Encode using CT
	c := casesMsgsTest[3]
	dMsgCT, err := newDialectMessage(c.parsed)
	require.NoError(t, err)
	bytesEncoded, err := dMsgCT.encode(c.parsed, c.isV2)
	require.NoError(t, err)
	fmt.Println([]byte(bytesEncoded))

	b := "\\"
	for _, integer := range bytesEncoded {
		b += fmt.Sprintf("x%X", integer)
		b += "\\"
	}
	fmt.Println("manual bytes conversion:")
	fmt.Println(b)
	fmt.Println("...................................")

	require.Equal(t, c.raw, bytesEncoded)

	// Decode bytes using RT
	includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
	defs, version, err := libgen.XMLToFields("TestingXMLFiles/testing3.xml", includeDirs)
	require.NoError(t, err)

	// Create dialect from the parsed defs.
	dRT, err := NewDialectRT(version, defs)
	dMsgRT := dRT.messages[74]
	require.NoError(t, err)
	require.Equal(t, uint(3), dRT.getVersion())

	// Decode Bytes using RT
	msgDecoded, err := dMsgRT.decode(bytesEncoded, c.isV2)
	require.NoError(t, err)
	fmt.Printf("%+v\n", msgDecoded)

	// Try encoding using RT
	bytesEncodedByRT, err := dMsgRT.encode(msgDecoded, c.isV2)
	require.NoError(t, err)
	fmt.Println(bytesEncodedByRT)
	require.Equal(t, c.raw, bytesEncodedByRT)
}

var casesMsgsTest = []struct {
	name   string
	isV2   bool
	parsed Message
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
}

var CTMessages = []Message{
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
