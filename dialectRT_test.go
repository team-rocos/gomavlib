package gomavlib
import (
    "testing"
    "github.com/stretchr/testify/require"
	libgen "github.com/team-rocos/gomavlib/commands/dialgen/libgen"
)
// DEFINE PUBLIC TYPES AND STRUCTURES.

// DEFINE PRIVATE TYPES AND STRUCTURES.
// DEFINE PUBLIC STATIC FUNCTIONS.
func TestDialectRTCommonXML(t *testing.T) {
    // Parse the XML file.
    includeDirs := []string{"./mavlink-upstream/message_definitions/v1.0"}
    defs, version, err := libgen.XMLToFields("./mavlink-upstream/message_definitions/v1.0/common.xml", includeDirs)
	require.NoError(t, err)
	
    // Create dialect from the parsed defs.
    dRT, err := NewDialectRT(version, defs)
    require.NoError(t, err)
	require.Equal(t, uint(3), dRT.getVersion())
	
    // Check Individual Messages
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
	
	// Compare RT and CT for all messages
	for _, m := range CTMessages {
		index := m.GetId()
		// Compare dCT with dRT
		mCT := dCT.messages[index]
		mRT := dRT.messages[index]
		require.Equal(t, mCT.sizeNormal, byte(mRT.sizeNormal))
		require.Equal(t, mCT.sizeExtended, byte(mRT.sizeExtended))
		require.Equal(t, mCT.crcExtra, mRT.crcExtra)
	}
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