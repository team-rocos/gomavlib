package testgomavlib

import (
	"reflect"

	"github.com/team-rocos/gomavlib"
)

// this file contains a test dialect used in other tests.
// it's better not to import real dialects, but to use a separate one

var testDialect = gomavlib.MustDialectCT(3, []gomavlib.Message{
	&MessageTest5{},
	&MessageTest6{},
	&MessageTest8{},
	&MessageHeartbeat{},
	&MessageOpticalFlow{},
})

type MessageTest5 struct {
	TestByte byte
	TestUint uint32
}

func (m *MessageTest5) GetId() uint32 {
	return 5
}

func (m *MessageTest5) SetField(field string, value interface{}) error {
	reflect.ValueOf(m).Elem().FieldByName(field).Set(reflect.ValueOf(value))
	return nil
}

type MessageTest6 struct {
	TestByte byte
	TestUint uint32
}

func (m *MessageTest6) GetId() uint32 {
	return 0x0607
}

func (m *MessageTest6) SetField(field string, value interface{}) error {
	reflect.ValueOf(m).Elem().FieldByName(field).Set(reflect.ValueOf(value))
	return nil
}

type MessageTest8 struct {
	TestByte byte
	TestUint uint32
}

func (m *MessageTest8) GetId() uint32 {
	return 8
}

func (m *MessageTest8) SetField(field string, value interface{}) error {
	reflect.ValueOf(m).Elem().FieldByName(field).Set(reflect.ValueOf(value))
	return nil
}

type MessageAhrs struct {
	OmegaIx     float32 `mavname:"omegaIx"`
	OmegaIy     float32 `mavname:"omegaIy"`
	OmegaIz     float32 `mavname:"omegaIz"`
	AccelWeight float32
	RenormVal   float32
	ErrorRp     float32
	ErrorYaw    float32
}

func (*MessageAhrs) GetId() uint32 {
	return 163
}

func (m *MessageAhrs) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Enumeration of the ADSB altimeter types
type ADSB_ALTITUDE_TYPE int

const (
	// Altitude reported from a Baro source using QNH reference
	ADSB_ALTITUDE_TYPE_PRESSURE_QNH ADSB_ALTITUDE_TYPE = 0
	// Altitude reported from a GNSS source
	ADSB_ALTITUDE_TYPE_GEOMETRIC ADSB_ALTITUDE_TYPE = 1
)

// ADSB classification for the type of vehicle emitting the transponder signal
type ADSB_EMITTER_TYPE int

const (
	//
	ADSB_EMITTER_TYPE_NO_INFO ADSB_EMITTER_TYPE = 0
	//
	ADSB_EMITTER_TYPE_LIGHT ADSB_EMITTER_TYPE = 1
	//
	ADSB_EMITTER_TYPE_SMALL ADSB_EMITTER_TYPE = 2
	//
	ADSB_EMITTER_TYPE_LARGE ADSB_EMITTER_TYPE = 3
	//
	ADSB_EMITTER_TYPE_HIGH_VORTEX_LARGE ADSB_EMITTER_TYPE = 4
	//
	ADSB_EMITTER_TYPE_HEAVY ADSB_EMITTER_TYPE = 5
	//
	ADSB_EMITTER_TYPE_HIGHLY_MANUV ADSB_EMITTER_TYPE = 6
	//
	ADSB_EMITTER_TYPE_ROTOCRAFT ADSB_EMITTER_TYPE = 7
	//
	ADSB_EMITTER_TYPE_UNASSIGNED ADSB_EMITTER_TYPE = 8
	//
	ADSB_EMITTER_TYPE_GLIDER ADSB_EMITTER_TYPE = 9
	//
	ADSB_EMITTER_TYPE_LIGHTER_AIR ADSB_EMITTER_TYPE = 10
	//
	ADSB_EMITTER_TYPE_PARACHUTE ADSB_EMITTER_TYPE = 11
	//
	ADSB_EMITTER_TYPE_ULTRA_LIGHT ADSB_EMITTER_TYPE = 12
	//
	ADSB_EMITTER_TYPE_UNASSIGNED2 ADSB_EMITTER_TYPE = 13
	//
	ADSB_EMITTER_TYPE_UAV ADSB_EMITTER_TYPE = 14
	//
	ADSB_EMITTER_TYPE_SPACE ADSB_EMITTER_TYPE = 15
	//
	ADSB_EMITTER_TYPE_UNASSGINED3 ADSB_EMITTER_TYPE = 16
	//
	ADSB_EMITTER_TYPE_EMERGENCY_SURFACE ADSB_EMITTER_TYPE = 17
	//
	ADSB_EMITTER_TYPE_SERVICE_SURFACE ADSB_EMITTER_TYPE = 18
	//
	ADSB_EMITTER_TYPE_POINT_OBSTACLE ADSB_EMITTER_TYPE = 19
)

// These flags indicate status such as data validity of each data source. Set = data valid
type ADSB_FLAGS int

const (
	//
	ADSB_FLAGS_VALID_COORDS ADSB_FLAGS = 1
	//
	ADSB_FLAGS_VALID_ALTITUDE ADSB_FLAGS = 2
	//
	ADSB_FLAGS_VALID_HEADING ADSB_FLAGS = 4
	//
	ADSB_FLAGS_VALID_VELOCITY ADSB_FLAGS = 8
	//
	ADSB_FLAGS_VALID_CALLSIGN ADSB_FLAGS = 16
	//
	ADSB_FLAGS_VALID_SQUAWK ADSB_FLAGS = 32
	//
	ADSB_FLAGS_SIMULATED ADSB_FLAGS = 64
	//
	ADSB_FLAGS_VERTICAL_VELOCITY_VALID ADSB_FLAGS = 128
	//
	ADSB_FLAGS_BARO_VALID ADSB_FLAGS = 256
	//
	ADSB_FLAGS_SOURCE_UAT ADSB_FLAGS = 32768
)

// These flags are used in the AIS_VESSEL.fields bitmask to indicate validity of data in the other message fields. When set, the data is valid.
type AIS_FLAGS int

const (
	// 1 = Position accuracy less than 10m, 0 = position accuracy greater than 10m.
	AIS_FLAGS_POSITION_ACCURACY AIS_FLAGS = 1
	//
	AIS_FLAGS_VALID_COG AIS_FLAGS = 2
	//
	AIS_FLAGS_VALID_VELOCITY AIS_FLAGS = 4
	// 1 = Velocity over 52.5765m/s (102.2 knots)
	AIS_FLAGS_HIGH_VELOCITY AIS_FLAGS = 8
	//
	AIS_FLAGS_VALID_TURN_RATE AIS_FLAGS = 16
	// Only the sign of the returned turn rate value is valid, either greater than 5deg/30s or less than -5deg/30s
	AIS_FLAGS_TURN_RATE_SIGN_ONLY AIS_FLAGS = 32
	//
	AIS_FLAGS_VALID_DIMENSIONS AIS_FLAGS = 64
	// Distance to bow is larger than 511m
	AIS_FLAGS_LARGE_BOW_DIMENSION AIS_FLAGS = 128
	// Distance to stern is larger than 511m
	AIS_FLAGS_LARGE_STERN_DIMENSION AIS_FLAGS = 256
	// Distance to port side is larger than 63m
	AIS_FLAGS_LARGE_PORT_DIMENSION AIS_FLAGS = 512
	// Distance to starboard side is larger than 63m
	AIS_FLAGS_LARGE_STARBOARD_DIMENSION AIS_FLAGS = 1024
	//
	AIS_FLAGS_VALID_CALLSIGN AIS_FLAGS = 2048
	//
	AIS_FLAGS_VALID_NAME AIS_FLAGS = 4096
)

// Navigational status of AIS vessel, enum duplicated from AIS standard, https://gpsd.gitlab.io/gpsd/AIVDM.html
type AIS_NAV_STATUS int

const (
	// Under way using engine.
	UNDER_WAY AIS_NAV_STATUS = 0
	//
	AIS_NAV_ANCHORED AIS_NAV_STATUS = 1
	//
	AIS_NAV_UN_COMMANDED AIS_NAV_STATUS = 2
	//
	AIS_NAV_RESTRICTED_MANOEUVERABILITY AIS_NAV_STATUS = 3
	//
	AIS_NAV_DRAUGHT_CONSTRAINED AIS_NAV_STATUS = 4
	//
	AIS_NAV_MOORED AIS_NAV_STATUS = 5
	//
	AIS_NAV_AGROUND AIS_NAV_STATUS = 6
	//
	AIS_NAV_FISHING AIS_NAV_STATUS = 7
	//
	AIS_NAV_SAILING AIS_NAV_STATUS = 8
	//
	AIS_NAV_RESERVED_HSC AIS_NAV_STATUS = 9
	//
	AIS_NAV_RESERVED_WIG AIS_NAV_STATUS = 10
	//
	AIS_NAV_RESERVED_1 AIS_NAV_STATUS = 11
	//
	AIS_NAV_RESERVED_2 AIS_NAV_STATUS = 12
	//
	AIS_NAV_RESERVED_3 AIS_NAV_STATUS = 13
	// Search And Rescue Transponder.
	AIS_NAV_AIS_SART AIS_NAV_STATUS = 14
	// Not available (default).
	AIS_NAV_UNKNOWN AIS_NAV_STATUS = 15
)

// Type of AIS vessel, enum duplicated from AIS standard, https://gpsd.gitlab.io/gpsd/AIVDM.html
type AIS_TYPE int

const (
	// Not available (default).
	AIS_TYPE_UNKNOWN AIS_TYPE = 0
	//
	AIS_TYPE_RESERVED_1 AIS_TYPE = 1
	//
	AIS_TYPE_RESERVED_2 AIS_TYPE = 2
	//
	AIS_TYPE_RESERVED_3 AIS_TYPE = 3
	//
	AIS_TYPE_RESERVED_4 AIS_TYPE = 4
	//
	AIS_TYPE_RESERVED_5 AIS_TYPE = 5
	//
	AIS_TYPE_RESERVED_6 AIS_TYPE = 6
	//
	AIS_TYPE_RESERVED_7 AIS_TYPE = 7
	//
	AIS_TYPE_RESERVED_8 AIS_TYPE = 8
	//
	AIS_TYPE_RESERVED_9 AIS_TYPE = 9
	//
	AIS_TYPE_RESERVED_10 AIS_TYPE = 10
	//
	AIS_TYPE_RESERVED_11 AIS_TYPE = 11
	//
	AIS_TYPE_RESERVED_12 AIS_TYPE = 12
	//
	AIS_TYPE_RESERVED_13 AIS_TYPE = 13
	//
	AIS_TYPE_RESERVED_14 AIS_TYPE = 14
	//
	AIS_TYPE_RESERVED_15 AIS_TYPE = 15
	//
	AIS_TYPE_RESERVED_16 AIS_TYPE = 16
	//
	AIS_TYPE_RESERVED_17 AIS_TYPE = 17
	//
	AIS_TYPE_RESERVED_18 AIS_TYPE = 18
	//
	AIS_TYPE_RESERVED_19 AIS_TYPE = 19
	// Wing In Ground effect.
	AIS_TYPE_WIG AIS_TYPE = 20
	//
	AIS_TYPE_WIG_HAZARDOUS_A AIS_TYPE = 21
	//
	AIS_TYPE_WIG_HAZARDOUS_B AIS_TYPE = 22
	//
	AIS_TYPE_WIG_HAZARDOUS_C AIS_TYPE = 23
	//
	AIS_TYPE_WIG_HAZARDOUS_D AIS_TYPE = 24
	//
	AIS_TYPE_WIG_RESERVED_1 AIS_TYPE = 25
	//
	AIS_TYPE_WIG_RESERVED_2 AIS_TYPE = 26
	//
	AIS_TYPE_WIG_RESERVED_3 AIS_TYPE = 27
	//
	AIS_TYPE_WIG_RESERVED_4 AIS_TYPE = 28
	//
	AIS_TYPE_WIG_RESERVED_5 AIS_TYPE = 29
	//
	AIS_TYPE_FISHING AIS_TYPE = 30
	//
	AIS_TYPE_TOWING AIS_TYPE = 31
	// Towing: length exceeds 200m or breadth exceeds 25m.
	AIS_TYPE_TOWING_LARGE AIS_TYPE = 32
	// Dredging or other underwater ops.
	AIS_TYPE_DREDGING AIS_TYPE = 33
	//
	AIS_TYPE_DIVING AIS_TYPE = 34
	//
	AIS_TYPE_MILITARY AIS_TYPE = 35
	//
	AIS_TYPE_SAILING AIS_TYPE = 36
	//
	AIS_TYPE_PLEASURE AIS_TYPE = 37
	//
	AIS_TYPE_RESERVED_20 AIS_TYPE = 38
	//
	AIS_TYPE_RESERVED_21 AIS_TYPE = 39
	// High Speed Craft.
	AIS_TYPE_HSC AIS_TYPE = 40
	//
	AIS_TYPE_HSC_HAZARDOUS_A AIS_TYPE = 41
	//
	AIS_TYPE_HSC_HAZARDOUS_B AIS_TYPE = 42
	//
	AIS_TYPE_HSC_HAZARDOUS_C AIS_TYPE = 43
	//
	AIS_TYPE_HSC_HAZARDOUS_D AIS_TYPE = 44
	//
	AIS_TYPE_HSC_RESERVED_1 AIS_TYPE = 45
	//
	AIS_TYPE_HSC_RESERVED_2 AIS_TYPE = 46
	//
	AIS_TYPE_HSC_RESERVED_3 AIS_TYPE = 47
	//
	AIS_TYPE_HSC_RESERVED_4 AIS_TYPE = 48
	//
	AIS_TYPE_HSC_UNKNOWN AIS_TYPE = 49
	//
	AIS_TYPE_PILOT AIS_TYPE = 50
	// Search And Rescue vessel.
	AIS_TYPE_SAR AIS_TYPE = 51
	//
	AIS_TYPE_TUG AIS_TYPE = 52
	//
	AIS_TYPE_PORT_TENDER AIS_TYPE = 53
	// Anti-pollution equipment.
	AIS_TYPE_ANTI_POLLUTION AIS_TYPE = 54
	//
	AIS_TYPE_LAW_ENFORCEMENT AIS_TYPE = 55
	//
	AIS_TYPE_SPARE_LOCAL_1 AIS_TYPE = 56
	//
	AIS_TYPE_SPARE_LOCAL_2 AIS_TYPE = 57
	//
	AIS_TYPE_MEDICAL_TRANSPORT AIS_TYPE = 58
	// Noncombatant ship according to RR Resolution No. 18.
	AIS_TYPE_NONECOMBATANT AIS_TYPE = 59
	//
	AIS_TYPE_PASSENGER AIS_TYPE = 60
	//
	AIS_TYPE_PASSENGER_HAZARDOUS_A AIS_TYPE = 61
	//
	AIS_TYPE_PASSENGER_HAZARDOUS_B AIS_TYPE = 62
	//
	AIS_TYPE_AIS_TYPE_PASSENGER_HAZARDOUS_C AIS_TYPE = 63
	//
	AIS_TYPE_PASSENGER_HAZARDOUS_D AIS_TYPE = 64
	//
	AIS_TYPE_PASSENGER_RESERVED_1 AIS_TYPE = 65
	//
	AIS_TYPE_PASSENGER_RESERVED_2 AIS_TYPE = 66
	//
	AIS_TYPE_PASSENGER_RESERVED_3 AIS_TYPE = 67
	//
	AIS_TYPE_AIS_TYPE_PASSENGER_RESERVED_4 AIS_TYPE = 68
	//
	AIS_TYPE_PASSENGER_UNKNOWN AIS_TYPE = 69
	//
	AIS_TYPE_CARGO AIS_TYPE = 70
	//
	AIS_TYPE_CARGO_HAZARDOUS_A AIS_TYPE = 71
	//
	AIS_TYPE_CARGO_HAZARDOUS_B AIS_TYPE = 72
	//
	AIS_TYPE_CARGO_HAZARDOUS_C AIS_TYPE = 73
	//
	AIS_TYPE_CARGO_HAZARDOUS_D AIS_TYPE = 74
	//
	AIS_TYPE_CARGO_RESERVED_1 AIS_TYPE = 75
	//
	AIS_TYPE_CARGO_RESERVED_2 AIS_TYPE = 76
	//
	AIS_TYPE_CARGO_RESERVED_3 AIS_TYPE = 77
	//
	AIS_TYPE_CARGO_RESERVED_4 AIS_TYPE = 78
	//
	AIS_TYPE_CARGO_UNKNOWN AIS_TYPE = 79
	//
	AIS_TYPE_TANKER AIS_TYPE = 80
	//
	AIS_TYPE_TANKER_HAZARDOUS_A AIS_TYPE = 81
	//
	AIS_TYPE_TANKER_HAZARDOUS_B AIS_TYPE = 82
	//
	AIS_TYPE_TANKER_HAZARDOUS_C AIS_TYPE = 83
	//
	AIS_TYPE_TANKER_HAZARDOUS_D AIS_TYPE = 84
	//
	AIS_TYPE_TANKER_RESERVED_1 AIS_TYPE = 85
	//
	AIS_TYPE_TANKER_RESERVED_2 AIS_TYPE = 86
	//
	AIS_TYPE_TANKER_RESERVED_3 AIS_TYPE = 87
	//
	AIS_TYPE_TANKER_RESERVED_4 AIS_TYPE = 88
	//
	AIS_TYPE_TANKER_UNKNOWN AIS_TYPE = 89
	//
	AIS_TYPE_OTHER AIS_TYPE = 90
	//
	AIS_TYPE_OTHER_HAZARDOUS_A AIS_TYPE = 91
	//
	AIS_TYPE_OTHER_HAZARDOUS_B AIS_TYPE = 92
	//
	AIS_TYPE_OTHER_HAZARDOUS_C AIS_TYPE = 93
	//
	AIS_TYPE_OTHER_HAZARDOUS_D AIS_TYPE = 94
	//
	AIS_TYPE_OTHER_RESERVED_1 AIS_TYPE = 95
	//
	AIS_TYPE_OTHER_RESERVED_2 AIS_TYPE = 96
	//
	AIS_TYPE_OTHER_RESERVED_3 AIS_TYPE = 97
	//
	AIS_TYPE_OTHER_RESERVED_4 AIS_TYPE = 98
	//
	AIS_TYPE_OTHER_UNKNOWN AIS_TYPE = 99
)

// Camera capability flags (Bitmap)
type CAMERA_CAP_FLAGS int

const (
	// Camera is able to record video
	CAMERA_CAP_FLAGS_CAPTURE_VIDEO CAMERA_CAP_FLAGS = 1
	// Camera is able to capture images
	CAMERA_CAP_FLAGS_CAPTURE_IMAGE CAMERA_CAP_FLAGS = 2
	// Camera has separate Video and Image/Photo modes (MAV_CMD_SET_CAMERA_MODE)
	CAMERA_CAP_FLAGS_HAS_MODES CAMERA_CAP_FLAGS = 4
	// Camera can capture images while in video mode
	CAMERA_CAP_FLAGS_CAN_CAPTURE_IMAGE_IN_VIDEO_MODE CAMERA_CAP_FLAGS = 8
	// Camera can capture videos while in Photo/Image mode
	CAMERA_CAP_FLAGS_CAN_CAPTURE_VIDEO_IN_IMAGE_MODE CAMERA_CAP_FLAGS = 16
	// Camera has image survey mode (MAV_CMD_SET_CAMERA_MODE)
	CAMERA_CAP_FLAGS_HAS_IMAGE_SURVEY_MODE CAMERA_CAP_FLAGS = 32
	// Camera has basic zoom control (MAV_CMD_SET_CAMERA_ZOOM)
	CAMERA_CAP_FLAGS_HAS_BASIC_ZOOM CAMERA_CAP_FLAGS = 64
	// Camera has basic focus control (MAV_CMD_SET_CAMERA_FOCUS)
	CAMERA_CAP_FLAGS_HAS_BASIC_FOCUS CAMERA_CAP_FLAGS = 128
	// Camera has video streaming capabilities (use MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION for video streaming info)
	CAMERA_CAP_FLAGS_HAS_VIDEO_STREAM CAMERA_CAP_FLAGS = 256
)

// Camera Modes.
type CAMERA_MODE int

const (
	// Camera is in image/photo capture mode.
	CAMERA_MODE_IMAGE CAMERA_MODE = 0
	// Camera is in video capture mode.
	CAMERA_MODE_VIDEO CAMERA_MODE = 1
	// Camera is in image survey capture mode. It allows for camera controller to do specific settings for surveys.
	CAMERA_MODE_IMAGE_SURVEY CAMERA_MODE = 2
)

// Zoom types for MAV_CMD_SET_CAMERA_ZOOM
type CAMERA_ZOOM_TYPE int

const (
	// Zoom one step increment (-1 for wide, 1 for tele)
	ZOOM_TYPE_STEP CAMERA_ZOOM_TYPE = 0
	// Continuous zoom up/down until stopped (-1 for wide, 1 for tele, 0 to stop zooming)
	ZOOM_TYPE_CONTINUOUS CAMERA_ZOOM_TYPE = 1
	// Zoom value as proportion of full camera range (a value between 0.0 and 100.0)
	ZOOM_TYPE_RANGE CAMERA_ZOOM_TYPE = 2
	// Zoom value/variable focal length in milimetres. Note that there is no message to get the valid zoom range of the camera, so this can type can only be used for cameras where the zoom range is known (implying that this cannot reliably be used in a GCS for an arbitrary camera)
	ZOOM_TYPE_FOCAL_LENGTH CAMERA_ZOOM_TYPE = 3
)

// Cellular network radio type
type CELLULAR_NETWORK_RADIO_TYPE int

const (
	//
	CELLULAR_NETWORK_RADIO_TYPE_NONE CELLULAR_NETWORK_RADIO_TYPE = 0
	//
	CELLULAR_NETWORK_RADIO_TYPE_GSM CELLULAR_NETWORK_RADIO_TYPE = 1
	//
	CELLULAR_NETWORK_RADIO_TYPE_CDMA CELLULAR_NETWORK_RADIO_TYPE = 2
	//
	CELLULAR_NETWORK_RADIO_TYPE_WCDMA CELLULAR_NETWORK_RADIO_TYPE = 3
	//
	CELLULAR_NETWORK_RADIO_TYPE_LTE CELLULAR_NETWORK_RADIO_TYPE = 4
)

// These flags encode the cellular network status
type CELLULAR_NETWORK_STATUS_FLAG int

const (
	// Roaming is active
	CELLULAR_NETWORK_STATUS_FLAG_ROAMING CELLULAR_NETWORK_STATUS_FLAG = 1
)

// Component capability flags (Bitmap)
type COMPONENT_CAP_FLAGS int

const (
	// Component has parameters, and supports the parameter protocol (PARAM messages).
	COMPONENT_CAP_FLAGS_PARAM COMPONENT_CAP_FLAGS = 1
	// Component has parameters, and supports the extended parameter protocol (PARAM_EXT messages).
	COMPONENT_CAP_FLAGS_PARAM_EXT COMPONENT_CAP_FLAGS = 2
)

// Flags in EKF_STATUS message
type ESTIMATOR_STATUS_FLAGS int

const (
	// True if the attitude estimate is good
	ESTIMATOR_ATTITUDE ESTIMATOR_STATUS_FLAGS = 1
	// True if the horizontal velocity estimate is good
	ESTIMATOR_VELOCITY_HORIZ ESTIMATOR_STATUS_FLAGS = 2
	// True if the  vertical velocity estimate is good
	ESTIMATOR_VELOCITY_VERT ESTIMATOR_STATUS_FLAGS = 4
	// True if the horizontal position (relative) estimate is good
	ESTIMATOR_POS_HORIZ_REL ESTIMATOR_STATUS_FLAGS = 8
	// True if the horizontal position (absolute) estimate is good
	ESTIMATOR_POS_HORIZ_ABS ESTIMATOR_STATUS_FLAGS = 16
	// True if the vertical position (absolute) estimate is good
	ESTIMATOR_POS_VERT_ABS ESTIMATOR_STATUS_FLAGS = 32
	// True if the vertical position (above ground) estimate is good
	ESTIMATOR_POS_VERT_AGL ESTIMATOR_STATUS_FLAGS = 64
	// True if the EKF is in a constant position mode and is not using external measurements (eg GPS or optical flow)
	ESTIMATOR_CONST_POS_MODE ESTIMATOR_STATUS_FLAGS = 128
	// True if the EKF has sufficient data to enter a mode that will provide a (relative) position estimate
	ESTIMATOR_PRED_POS_HORIZ_REL ESTIMATOR_STATUS_FLAGS = 256
	// True if the EKF has sufficient data to enter a mode that will provide a (absolute) position estimate
	ESTIMATOR_PRED_POS_HORIZ_ABS ESTIMATOR_STATUS_FLAGS = 512
	// True if the EKF has detected a GPS glitch
	ESTIMATOR_GPS_GLITCH ESTIMATOR_STATUS_FLAGS = 1024
	// True if the EKF has detected bad accelerometer data
	ESTIMATOR_ACCEL_ERROR ESTIMATOR_STATUS_FLAGS = 2048
)

// List of possible failure type to inject.
type FAILURE_TYPE int

const (
	// No failure injected, used to reset a previous failure.
	FAILURE_TYPE_OK FAILURE_TYPE = 0
	// Sets unit off, so completely non-responsive.
	FAILURE_TYPE_OFF FAILURE_TYPE = 1
	// Unit is stuck e.g. keeps reporting the same value.
	FAILURE_TYPE_STUCK FAILURE_TYPE = 2
	// Unit is reporting complete garbage.
	FAILURE_TYPE_GARBAGE FAILURE_TYPE = 3
	// Unit is consistently wrong.
	FAILURE_TYPE_WRONG FAILURE_TYPE = 4
	// Unit is slow, so e.g. reporting at slower than expected rate.
	FAILURE_TYPE_SLOW FAILURE_TYPE = 5
	// Data of unit is delayed in time.
	FAILURE_TYPE_DELAYED FAILURE_TYPE = 6
	// Unit is sometimes working, sometimes not.
	FAILURE_TYPE_INTERMITTENT FAILURE_TYPE = 7
)

// List of possible units where failures can be injected.
type FAILURE_UNIT int

const (
	//
	FAILURE_UNIT_SENSOR_GYRO FAILURE_UNIT = 0
	//
	FAILURE_UNIT_SENSOR_ACCEL FAILURE_UNIT = 1
	//
	FAILURE_UNIT_SENSOR_MAG FAILURE_UNIT = 2
	//
	FAILURE_UNIT_SENSOR_BARO FAILURE_UNIT = 3
	//
	FAILURE_UNIT_SENSOR_GPS FAILURE_UNIT = 4
	//
	FAILURE_UNIT_SENSOR_OPTICAL_FLOW FAILURE_UNIT = 5
	//
	FAILURE_UNIT_SENSOR_VIO FAILURE_UNIT = 6
	//
	FAILURE_UNIT_SENSOR_DISTANCE_SENSOR FAILURE_UNIT = 7
	//
	FAILURE_UNIT_SYSTEM_BATTERY FAILURE_UNIT = 100
	//
	FAILURE_UNIT_SYSTEM_MOTOR FAILURE_UNIT = 101
	//
	FAILURE_UNIT_SYSTEM_SERVO FAILURE_UNIT = 102
	//
	FAILURE_UNIT_SYSTEM_AVOIDANCE FAILURE_UNIT = 103
	//
	FAILURE_UNIT_SYSTEM_RC_SIGNAL FAILURE_UNIT = 104
	//
	FAILURE_UNIT_SYSTEM_MAVLINK_SIGNAL FAILURE_UNIT = 105
)

//
type FENCE_ACTION int

const (
	// Disable fenced mode
	FENCE_ACTION_NONE FENCE_ACTION = 0
	// Switched to guided mode to return point (fence point 0)
	FENCE_ACTION_GUIDED FENCE_ACTION = 1
	// Report fence breach, but don't take action
	FENCE_ACTION_REPORT FENCE_ACTION = 2
	// Switched to guided mode to return point (fence point 0) with manual throttle control
	FENCE_ACTION_GUIDED_THR_PASS FENCE_ACTION = 3
	// Switch to RTL (return to launch) mode and head for the return point.
	FENCE_ACTION_RTL FENCE_ACTION = 4
)

//
type FENCE_BREACH int

const (
	// No last fence breach
	FENCE_BREACH_NONE FENCE_BREACH = 0
	// Breached minimum altitude
	FENCE_BREACH_MINALT FENCE_BREACH = 1
	// Breached maximum altitude
	FENCE_BREACH_MAXALT FENCE_BREACH = 2
	// Breached fence boundary
	FENCE_BREACH_BOUNDARY FENCE_BREACH = 3
)

// Actions being taken to mitigate/prevent fence breach
type FENCE_MITIGATE int

const (
	// Unknown
	FENCE_MITIGATE_UNKNOWN FENCE_MITIGATE = 0
	// No actions being taken
	FENCE_MITIGATE_NONE FENCE_MITIGATE = 1
	// Velocity limiting active to prevent breach
	FENCE_MITIGATE_VEL_LIMIT FENCE_MITIGATE = 2
)

// These values define the type of firmware release.  These values indicate the first version or release of this type.  For example the first alpha release would be 64, the second would be 65.
type FIRMWARE_VERSION_TYPE int

const (
	// development release
	FIRMWARE_VERSION_TYPE_DEV FIRMWARE_VERSION_TYPE = 0
	// alpha release
	FIRMWARE_VERSION_TYPE_ALPHA FIRMWARE_VERSION_TYPE = 64
	// beta release
	FIRMWARE_VERSION_TYPE_BETA FIRMWARE_VERSION_TYPE = 128
	// release candidate
	FIRMWARE_VERSION_TYPE_RC FIRMWARE_VERSION_TYPE = 192
	// official stable release
	FIRMWARE_VERSION_TYPE_OFFICIAL FIRMWARE_VERSION_TYPE = 255
)

// Gimbal device (low level) capability flags (bitmap)
type GIMBAL_DEVICE_CAP_FLAGS int

const (
	// Gimbal device supports a retracted position
	GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT GIMBAL_DEVICE_CAP_FLAGS = 1
	// Gimbal device supports a horizontal, forward looking position, stabilized
	GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL GIMBAL_DEVICE_CAP_FLAGS = 2
	// Gimbal device supports rotating around roll axis.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS GIMBAL_DEVICE_CAP_FLAGS = 4
	// Gimbal device supports to follow a roll angle relative to the vehicle
	GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW GIMBAL_DEVICE_CAP_FLAGS = 8
	// Gimbal device supports locking to an roll angle (generally that's the default with roll stabilized)
	GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK GIMBAL_DEVICE_CAP_FLAGS = 16
	// Gimbal device supports rotating around pitch axis.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS GIMBAL_DEVICE_CAP_FLAGS = 32
	// Gimbal device supports to follow a pitch angle relative to the vehicle
	GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW GIMBAL_DEVICE_CAP_FLAGS = 64
	// Gimbal device supports locking to an pitch angle (generally that's the default with pitch stabilized)
	GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK GIMBAL_DEVICE_CAP_FLAGS = 128
	// Gimbal device supports rotating around yaw axis.
	GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS GIMBAL_DEVICE_CAP_FLAGS = 256
	// Gimbal device supports to follow a yaw angle relative to the vehicle (generally that's the default)
	GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW GIMBAL_DEVICE_CAP_FLAGS = 512
	// Gimbal device supports locking to an absolute heading (often this is an option available)
	GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK GIMBAL_DEVICE_CAP_FLAGS = 1024
	// Gimbal device supports yawing/panning infinetely (e.g. using slip disk).
	GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW GIMBAL_DEVICE_CAP_FLAGS = 2048
)

// Gimbal device (low level) error flags (bitmap, 0 means no error)
type GIMBAL_DEVICE_ERROR_FLAGS int

const (
	// Gimbal device is limited by hardware roll limit.
	GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 1
	// Gimbal device is limited by hardware pitch limit.
	GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 2
	// Gimbal device is limited by hardware yaw limit.
	GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 4
	// There is an error with the gimbal encoders.
	GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 8
	// There is an error with the gimbal power source.
	GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 16
	// There is an error with the gimbal motor's.
	GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 32
	// There is an error with the gimbal's software.
	GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 64
	// There is an error with the gimbal's communication.
	GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 128
)

// Flags for gimbal device (lower level) operation.
type GIMBAL_DEVICE_FLAGS int

const (
	// Set to retracted safe position (no stabilization), takes presedence over all other flags.
	GIMBAL_DEVICE_FLAGS_RETRACT GIMBAL_DEVICE_FLAGS = 1
	// Set to neutral position (horizontal, forward looking, with stabiliziation), takes presedence over all other flags except RETRACT.
	GIMBAL_DEVICE_FLAGS_NEUTRAL GIMBAL_DEVICE_FLAGS = 2
	// Lock roll angle to absolute angle relative to horizon (not relative to drone). This is generally the default with a stabilizing gimbal.
	GIMBAL_DEVICE_FLAGS_ROLL_LOCK GIMBAL_DEVICE_FLAGS = 4
	// Lock pitch angle to absolute angle relative to horizon (not relative to drone). This is generally the default.
	GIMBAL_DEVICE_FLAGS_PITCH_LOCK GIMBAL_DEVICE_FLAGS = 8
	// Lock yaw angle to absolute angle relative to North (not relative to drone). If this flag is set, the quaternion is in the Earth frame with the x-axis pointing North (yaw absolute). If this flag is not set, the quaternion frame is in the Earth frame rotated so that the x-axis is pointing forward (yaw relative to vehicle).
	GIMBAL_DEVICE_FLAGS_YAW_LOCK GIMBAL_DEVICE_FLAGS = 16
)

// Gimbal manager high level capability flags (bitmap). The first 16 bits are identical to the GIMBAL_DEVICE_CAP_FLAGS which are identical with GIMBAL_DEVICE_FLAGS. However, the gimbal manager does not need to copy the flags from the gimbal but can also enhance the capabilities and thus add flags.
type GIMBAL_MANAGER_CAP_FLAGS int

const (
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT GIMBAL_MANAGER_CAP_FLAGS = 1
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL GIMBAL_MANAGER_CAP_FLAGS = 2
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS GIMBAL_MANAGER_CAP_FLAGS = 4
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = 8
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK GIMBAL_MANAGER_CAP_FLAGS = 16
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS GIMBAL_MANAGER_CAP_FLAGS = 32
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = 64
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK GIMBAL_MANAGER_CAP_FLAGS = 128
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS GIMBAL_MANAGER_CAP_FLAGS = 256
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = 512
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK GIMBAL_MANAGER_CAP_FLAGS = 1024
	// Based on GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW.
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW GIMBAL_MANAGER_CAP_FLAGS = 2048
	// Gimbal manager supports to point to a local position.
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL GIMBAL_MANAGER_CAP_FLAGS = 65536
	// Gimbal manager supports to point to a global latitude, longitude, altitude position.
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL GIMBAL_MANAGER_CAP_FLAGS = 131072
	// Gimbal manager supports tracking of a point on the camera.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_TRACKING_POINT GIMBAL_MANAGER_CAP_FLAGS = 262144
	// Gimbal manager supports tracking of a point on the camera.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_TRACKING_RECTANGLE GIMBAL_MANAGER_CAP_FLAGS = 524288
	// Gimbal manager supports pitching and yawing at an angular velocity scaled by focal length (the more zoomed in, the slower the movement).
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_FOCAL_LENGTH_SCALE GIMBAL_MANAGER_CAP_FLAGS = 1048576
	// Gimbal manager supports nudging when pointing to a location or tracking.
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_NUDGING GIMBAL_MANAGER_CAP_FLAGS = 2097152
	// Gimbal manager supports overriding when pointing to a location or tracking.
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_OVERRIDE GIMBAL_MANAGER_CAP_FLAGS = 4194304
)

// Flags for high level gimbal manager operation The first 16 bytes are identical to the GIMBAL_DEVICE_FLAGS.
type GIMBAL_MANAGER_FLAGS int

const (
	// Based on GIMBAL_DEVICE_FLAGS_RETRACT
	GIMBAL_MANAGER_FLAGS_RETRACT GIMBAL_MANAGER_FLAGS = 1
	// Based on GIMBAL_DEVICE_FLAGS_NEUTRAL
	GIMBAL_MANAGER_FLAGS_NEUTRAL GIMBAL_MANAGER_FLAGS = 2
	// Based on GIMBAL_DEVICE_FLAGS_ROLL_LOCK
	GIMBAL_MANAGER_FLAGS_ROLL_LOCK GIMBAL_MANAGER_FLAGS = 4
	// Based on GIMBAL_DEVICE_FLAGS_PITCH_LOCK
	GIMBAL_MANAGER_FLAGS_PITCH_LOCK GIMBAL_MANAGER_FLAGS = 8
	// Based on GIMBAL_DEVICE_FLAGS_YAW_LOCK
	GIMBAL_MANAGER_FLAGS_YAW_LOCK GIMBAL_MANAGER_FLAGS = 16
	// Scale angular velocity relative to focal length. This means the gimbal moves slower if it is zoomed in.
	GIMBAL_MANAGER_FLAGS_ANGULAR_VELOCITY_RELATIVE_TO_FOCAL_LENGTH GIMBAL_MANAGER_FLAGS = 1048576
	// Interpret attitude control on top of pointing to a location or tracking. If this flag is set, the quaternion is relative to the existing tracking angle.
	GIMBAL_MANAGER_FLAGS_NUDGE GIMBAL_MANAGER_FLAGS = 2097152
	// Completely override pointing to a location or tracking. If this flag is set, the quaternion is (as usual) according to GIMBAL_MANAGER_FLAGS_YAW_LOCK.
	GIMBAL_MANAGER_FLAGS_OVERRIDE GIMBAL_MANAGER_FLAGS = 4194304
	// This flag can be set to give up control previously set using MAV_CMD_DO_GIMBAL_MANAGER_ATTITUDE. This flag must not be combined with other flags.
	GIMBAL_MANAGER_FLAGS_NONE GIMBAL_MANAGER_FLAGS = 8388608
)

// Type of GPS fix
type GPS_FIX_TYPE int

const (
	// No GPS connected
	GPS_FIX_TYPE_NO_GPS GPS_FIX_TYPE = 0
	// No position information, GPS is connected
	GPS_FIX_TYPE_NO_FIX GPS_FIX_TYPE = 1
	// 2D position
	GPS_FIX_TYPE_2D_FIX GPS_FIX_TYPE = 2
	// 3D position
	GPS_FIX_TYPE_3D_FIX GPS_FIX_TYPE = 3
	// DGPS/SBAS aided 3D position
	GPS_FIX_TYPE_DGPS GPS_FIX_TYPE = 4
	// RTK float, 3D position
	GPS_FIX_TYPE_RTK_FLOAT GPS_FIX_TYPE = 5
	// RTK Fixed, 3D position
	GPS_FIX_TYPE_RTK_FIXED GPS_FIX_TYPE = 6
	// Static fixed, typically used for base stations
	GPS_FIX_TYPE_STATIC GPS_FIX_TYPE = 7
	// PPP, 3D position.
	GPS_FIX_TYPE_PPP GPS_FIX_TYPE = 8
)

//
type GPS_INPUT_IGNORE_FLAGS int

const (
	// ignore altitude field
	GPS_INPUT_IGNORE_FLAG_ALT GPS_INPUT_IGNORE_FLAGS = 1
	// ignore hdop field
	GPS_INPUT_IGNORE_FLAG_HDOP GPS_INPUT_IGNORE_FLAGS = 2
	// ignore vdop field
	GPS_INPUT_IGNORE_FLAG_VDOP GPS_INPUT_IGNORE_FLAGS = 4
	// ignore horizontal velocity field (vn and ve)
	GPS_INPUT_IGNORE_FLAG_VEL_HORIZ GPS_INPUT_IGNORE_FLAGS = 8
	// ignore vertical velocity field (vd)
	GPS_INPUT_IGNORE_FLAG_VEL_VERT GPS_INPUT_IGNORE_FLAGS = 16
	// ignore speed accuracy field
	GPS_INPUT_IGNORE_FLAG_SPEED_ACCURACY GPS_INPUT_IGNORE_FLAGS = 32
	// ignore horizontal accuracy field
	GPS_INPUT_IGNORE_FLAG_HORIZONTAL_ACCURACY GPS_INPUT_IGNORE_FLAGS = 64
	// ignore vertical accuracy field
	GPS_INPUT_IGNORE_FLAG_VERTICAL_ACCURACY GPS_INPUT_IGNORE_FLAGS = 128
)

// Flags to report failure cases over the high latency telemtry.
type HL_FAILURE_FLAG int

const (
	// GPS failure.
	HL_FAILURE_FLAG_GPS HL_FAILURE_FLAG = 1
	// Differential pressure sensor failure.
	HL_FAILURE_FLAG_DIFFERENTIAL_PRESSURE HL_FAILURE_FLAG = 2
	// Absolute pressure sensor failure.
	HL_FAILURE_FLAG_ABSOLUTE_PRESSURE HL_FAILURE_FLAG = 4
	// Accelerometer sensor failure.
	HL_FAILURE_FLAG_3D_ACCEL HL_FAILURE_FLAG = 8
	// Gyroscope sensor failure.
	HL_FAILURE_FLAG_3D_GYRO HL_FAILURE_FLAG = 16
	// Magnetometer sensor failure.
	HL_FAILURE_FLAG_3D_MAG HL_FAILURE_FLAG = 32
	// Terrain subsystem failure.
	HL_FAILURE_FLAG_TERRAIN HL_FAILURE_FLAG = 64
	// Battery failure/critical low battery.
	HL_FAILURE_FLAG_BATTERY HL_FAILURE_FLAG = 128
	// RC receiver failure/no rc connection.
	HL_FAILURE_FLAG_RC_RECEIVER HL_FAILURE_FLAG = 256
	// Offboard link failure.
	HL_FAILURE_FLAG_OFFBOARD_LINK HL_FAILURE_FLAG = 512
	// Engine failure.
	HL_FAILURE_FLAG_ENGINE HL_FAILURE_FLAG = 1024
	// Geofence violation.
	HL_FAILURE_FLAG_GEOFENCE HL_FAILURE_FLAG = 2048
	// Estimator failure, for example measurement rejection or large variances.
	HL_FAILURE_FLAG_ESTIMATOR HL_FAILURE_FLAG = 4096
	// Mission failure.
	HL_FAILURE_FLAG_MISSION HL_FAILURE_FLAG = 8192
)

// Type of landing target
type LANDING_TARGET_TYPE int

const (
	// Landing target signaled by light beacon (ex: IR-LOCK)
	LANDING_TARGET_TYPE_LIGHT_BEACON LANDING_TARGET_TYPE = 0
	// Landing target signaled by radio beacon (ex: ILS, NDB)
	LANDING_TARGET_TYPE_RADIO_BEACON LANDING_TARGET_TYPE = 1
	// Landing target represented by a fiducial marker (ex: ARTag)
	LANDING_TARGET_TYPE_VISION_FIDUCIAL LANDING_TARGET_TYPE = 2
	// Landing target represented by a pre-defined visual shape/feature (ex: X-marker, H-marker, square)
	LANDING_TARGET_TYPE_VISION_OTHER LANDING_TARGET_TYPE = 3
)

//
type MAVLINK_DATA_STREAM_TYPE int

const (
	//
	MAVLINK_DATA_STREAM_IMG_JPEG MAVLINK_DATA_STREAM_TYPE = 0
	//
	MAVLINK_DATA_STREAM_IMG_BMP MAVLINK_DATA_STREAM_TYPE = 1
	//
	MAVLINK_DATA_STREAM_IMG_RAW8U MAVLINK_DATA_STREAM_TYPE = 2
	//
	MAVLINK_DATA_STREAM_IMG_RAW32U MAVLINK_DATA_STREAM_TYPE = 3
	//
	MAVLINK_DATA_STREAM_IMG_PGM MAVLINK_DATA_STREAM_TYPE = 4
	//
	MAVLINK_DATA_STREAM_IMG_PNG MAVLINK_DATA_STREAM_TYPE = 5
)

//
type MAV_ARM_AUTH_DENIED_REASON int

const (
	// Not a specific reason
	MAV_ARM_AUTH_DENIED_REASON_GENERIC MAV_ARM_AUTH_DENIED_REASON = 0
	// Authorizer will send the error as string to GCS
	MAV_ARM_AUTH_DENIED_REASON_NONE MAV_ARM_AUTH_DENIED_REASON = 1
	// At least one waypoint have a invalid value
	MAV_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT MAV_ARM_AUTH_DENIED_REASON = 2
	// Timeout in the authorizer process(in case it depends on network)
	MAV_ARM_AUTH_DENIED_REASON_TIMEOUT MAV_ARM_AUTH_DENIED_REASON = 3
	// Airspace of the mission in use by another vehicle, second result parameter can have the waypoint id that caused it to be denied.
	MAV_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE MAV_ARM_AUTH_DENIED_REASON = 4
	// Weather is not good to fly
	MAV_ARM_AUTH_DENIED_REASON_BAD_WEATHER MAV_ARM_AUTH_DENIED_REASON = 5
)

// Micro air vehicle / autopilot classes. This identifies the individual model.
type MAV_AUTOPILOT int

const (
	// Generic autopilot, full support for everything
	MAV_AUTOPILOT_GENERIC MAV_AUTOPILOT = 0
	// Reserved for future use.
	MAV_AUTOPILOT_RESERVED MAV_AUTOPILOT = 1
	// SLUGS autopilot, http://slugsuav.soe.ucsc.edu
	MAV_AUTOPILOT_SLUGS MAV_AUTOPILOT = 2
	// ArduPilot - Plane/Copter/Rover/Sub/Tracker, https://ardupilot.org
	MAV_AUTOPILOT_ARDUPILOTMEGA MAV_AUTOPILOT = 3
	// OpenPilot, http://openpilot.org
	MAV_AUTOPILOT_OPENPILOT MAV_AUTOPILOT = 4
	// Generic autopilot only supporting simple waypoints
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY MAV_AUTOPILOT = 5
	// Generic autopilot supporting waypoints and other simple navigation commands
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY MAV_AUTOPILOT = 6
	// Generic autopilot supporting the full mission command set
	MAV_AUTOPILOT_GENERIC_MISSION_FULL MAV_AUTOPILOT = 7
	// No valid autopilot, e.g. a GCS or other MAVLink component
	MAV_AUTOPILOT_INVALID MAV_AUTOPILOT = 8
	// PPZ UAV - http://nongnu.org/paparazzi
	MAV_AUTOPILOT_PPZ MAV_AUTOPILOT = 9
	// UAV Dev Board
	MAV_AUTOPILOT_UDB MAV_AUTOPILOT = 10
	// FlexiPilot
	MAV_AUTOPILOT_FP MAV_AUTOPILOT = 11
	// PX4 Autopilot - http://px4.io/
	MAV_AUTOPILOT_PX4 MAV_AUTOPILOT = 12
	// SMACCMPilot - http://smaccmpilot.org
	MAV_AUTOPILOT_SMACCMPILOT MAV_AUTOPILOT = 13
	// AutoQuad -- http://autoquad.org
	MAV_AUTOPILOT_AUTOQUAD MAV_AUTOPILOT = 14
	// Armazila -- http://armazila.com
	MAV_AUTOPILOT_ARMAZILA MAV_AUTOPILOT = 15
	// Aerob -- http://aerob.ru
	MAV_AUTOPILOT_AEROB MAV_AUTOPILOT = 16
	// ASLUAV autopilot -- http://www.asl.ethz.ch
	MAV_AUTOPILOT_ASLUAV MAV_AUTOPILOT = 17
	// SmartAP Autopilot - http://sky-drones.com
	MAV_AUTOPILOT_SMARTAP MAV_AUTOPILOT = 18
	// AirRails - http://uaventure.com
	MAV_AUTOPILOT_AIRRAILS MAV_AUTOPILOT = 19
)

// Enumeration for battery charge states.
type MAV_BATTERY_CHARGE_STATE int

const (
	// Low battery state is not provided
	MAV_BATTERY_CHARGE_STATE_UNDEFINED MAV_BATTERY_CHARGE_STATE = 0
	// Battery is not in low state. Normal operation.
	MAV_BATTERY_CHARGE_STATE_OK MAV_BATTERY_CHARGE_STATE = 1
	// Battery state is low, warn and monitor close.
	MAV_BATTERY_CHARGE_STATE_LOW MAV_BATTERY_CHARGE_STATE = 2
	// Battery state is critical, return or abort immediately.
	MAV_BATTERY_CHARGE_STATE_CRITICAL MAV_BATTERY_CHARGE_STATE = 3
	// Battery state is too low for ordinary abort sequence. Perform fastest possible emergency stop to prevent damage.
	MAV_BATTERY_CHARGE_STATE_EMERGENCY MAV_BATTERY_CHARGE_STATE = 4
	// Battery failed, damage unavoidable.
	MAV_BATTERY_CHARGE_STATE_FAILED MAV_BATTERY_CHARGE_STATE = 5
	// Battery is diagnosed to be defective or an error occurred, usage is discouraged / prohibited.
	MAV_BATTERY_CHARGE_STATE_UNHEALTHY MAV_BATTERY_CHARGE_STATE = 6
	// Battery is charging.
	MAV_BATTERY_CHARGE_STATE_CHARGING MAV_BATTERY_CHARGE_STATE = 7
)

// Enumeration of battery functions
type MAV_BATTERY_FUNCTION int

const (
	// Battery function is unknown
	MAV_BATTERY_FUNCTION_UNKNOWN MAV_BATTERY_FUNCTION = 0
	// Battery supports all flight systems
	MAV_BATTERY_FUNCTION_ALL MAV_BATTERY_FUNCTION = 1
	// Battery for the propulsion system
	MAV_BATTERY_FUNCTION_PROPULSION MAV_BATTERY_FUNCTION = 2
	// Avionics battery
	MAV_BATTERY_FUNCTION_AVIONICS MAV_BATTERY_FUNCTION = 3
	// Payload battery
	MAV_BATTERY_TYPE_PAYLOAD MAV_BATTERY_FUNCTION = 4
)

// Enumeration of battery types
type MAV_BATTERY_TYPE int

const (
	// Not specified.
	MAV_BATTERY_TYPE_UNKNOWN MAV_BATTERY_TYPE = 0
	// Lithium polymer battery
	MAV_BATTERY_TYPE_LIPO MAV_BATTERY_TYPE = 1
	// Lithium-iron-phosphate battery
	MAV_BATTERY_TYPE_LIFE MAV_BATTERY_TYPE = 2
	// Lithium-ION battery
	MAV_BATTERY_TYPE_LION MAV_BATTERY_TYPE = 3
	// Nickel metal hydride battery
	MAV_BATTERY_TYPE_NIMH MAV_BATTERY_TYPE = 4
)

// Commands to be executed by the MAV. They can be executed on user request, or as part of a mission script. If the action is used in a mission, the parameter mapping to the waypoint/mission message is as follows: Param 1, Param 2, Param 3, Param 4, X: Param 5, Y:Param 6, Z:Param 7. This command list is similar what ARINC 424 is for commercial aircraft: A data format how to interpret waypoint/mission data. NaN and INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current yaw or latitude rather than a specific value). See https://mavlink.io/en/guide/xml_schema.html#MAV_CMD for information about the structure of the MAV_CMD entries
type MAV_CMD int

const (
	// Navigate to waypoint.
	MAV_CMD_NAV_WAYPOINT MAV_CMD = 16
	// Loiter around this waypoint an unlimited amount of time
	MAV_CMD_NAV_LOITER_UNLIM MAV_CMD = 17
	// Loiter around this waypoint for X turns
	MAV_CMD_NAV_LOITER_TURNS MAV_CMD = 18
	// Loiter around this waypoint for X seconds
	MAV_CMD_NAV_LOITER_TIME MAV_CMD = 19
	// Return to launch location
	MAV_CMD_NAV_RETURN_TO_LAUNCH MAV_CMD = 20
	// Land at location.
	MAV_CMD_NAV_LAND MAV_CMD = 21
	// Takeoff from ground / hand
	MAV_CMD_NAV_TAKEOFF MAV_CMD = 22
	// Land at local position (local frame only)
	MAV_CMD_NAV_LAND_LOCAL MAV_CMD = 23
	// Takeoff from local position (local frame only)
	MAV_CMD_NAV_TAKEOFF_LOCAL MAV_CMD = 24
	// Vehicle following, i.e. this waypoint represents the position of a moving vehicle
	MAV_CMD_NAV_FOLLOW MAV_CMD = 25
	// Continue on the current course and climb/descend to specified altitude.  When the altitude is reached continue to the next command (i.e., don't proceed to the next command until the desired altitude is reached.
	MAV_CMD_NAV_CONTINUE_AND_CHANGE_ALT MAV_CMD = 30
	// Begin loiter at the specified Latitude and Longitude.  If Lat=Lon=0, then loiter at the current position.  Don't consider the navigation command complete (don't leave loiter) until the altitude has been reached.  Additionally, if the Heading Required parameter is non-zero the  aircraft will not leave the loiter until heading toward the next waypoint.
	MAV_CMD_NAV_LOITER_TO_ALT MAV_CMD = 31
	// Begin following a target
	MAV_CMD_DO_FOLLOW MAV_CMD = 32
	// Reposition the MAV after a follow target command has been sent
	MAV_CMD_DO_FOLLOW_REPOSITION MAV_CMD = 33
	// Start orbiting on the circumference of a circle defined by the parameters. Setting any value NaN results in using defaults.
	MAV_CMD_DO_ORBIT MAV_CMD = 34
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	MAV_CMD_NAV_ROI MAV_CMD = 80
	// Control autonomous path planning on the MAV.
	MAV_CMD_NAV_PATHPLANNING MAV_CMD = 81
	// Navigate to waypoint using a spline path.
	MAV_CMD_NAV_SPLINE_WAYPOINT MAV_CMD = 82
	// Takeoff from ground using VTOL mode, and transition to forward flight with specified heading.
	MAV_CMD_NAV_VTOL_TAKEOFF MAV_CMD = 84
	// Land using VTOL mode
	MAV_CMD_NAV_VTOL_LAND MAV_CMD = 85
	// hand control over to an external controller
	MAV_CMD_NAV_GUIDED_ENABLE MAV_CMD = 92
	// Delay the next navigation command a number of seconds or until a specified time
	MAV_CMD_NAV_DELAY MAV_CMD = 93
	// Descend and place payload. Vehicle moves to specified location, descends until it detects a hanging payload has reached the ground, and then releases the payload. If ground is not detected before the reaching the maximum descent value (param1), the command will complete without releasing the payload.
	MAV_CMD_NAV_PAYLOAD_PLACE MAV_CMD = 94
	// NOP - This command is only used to mark the upper limit of the NAV/ACTION commands in the enumeration
	MAV_CMD_NAV_LAST MAV_CMD = 95
	// Delay mission state machine.
	MAV_CMD_CONDITION_DELAY MAV_CMD = 112
	// Ascend/descend at rate.  Delay mission state machine until desired altitude reached.
	MAV_CMD_CONDITION_CHANGE_ALT MAV_CMD = 113
	// Delay mission state machine until within desired distance of next NAV point.
	MAV_CMD_CONDITION_DISTANCE MAV_CMD = 114
	// Reach a certain target angle.
	MAV_CMD_CONDITION_YAW MAV_CMD = 115
	// NOP - This command is only used to mark the upper limit of the CONDITION commands in the enumeration
	MAV_CMD_CONDITION_LAST MAV_CMD = 159
	// Set system mode.
	MAV_CMD_DO_SET_MODE MAV_CMD = 176
	// Jump to the desired command in the mission list.  Repeat this action only the specified number of times
	MAV_CMD_DO_JUMP MAV_CMD = 177
	// Change speed and/or throttle set points.
	MAV_CMD_DO_CHANGE_SPEED MAV_CMD = 178
	// Changes the home location either to the current location or a specified location.
	MAV_CMD_DO_SET_HOME MAV_CMD = 179
	// Set a system parameter.  Caution!  Use of this command requires knowledge of the numeric enumeration value of the parameter.
	MAV_CMD_DO_SET_PARAMETER MAV_CMD = 180
	// Set a relay to a condition.
	MAV_CMD_DO_SET_RELAY MAV_CMD = 181
	// Cycle a relay on and off for a desired number of cycles with a desired period.
	MAV_CMD_DO_REPEAT_RELAY MAV_CMD = 182
	// Set a servo to a desired PWM value.
	MAV_CMD_DO_SET_SERVO MAV_CMD = 183
	// Cycle a between its nominal setting and a desired PWM for a desired number of cycles with a desired period.
	MAV_CMD_DO_REPEAT_SERVO MAV_CMD = 184
	// Terminate flight immediately
	MAV_CMD_DO_FLIGHTTERMINATION MAV_CMD = 185
	// Change altitude set point.
	MAV_CMD_DO_CHANGE_ALTITUDE MAV_CMD = 186
	// Mission command to perform a landing. This is used as a marker in a mission to tell the autopilot where a sequence of mission items that represents a landing starts. It may also be sent via a COMMAND_LONG to trigger a landing, in which case the nearest (geographically) landing sequence in the mission will be used. The Latitude/Longitude is optional, and may be set to 0 if not needed. If specified then it will be used to help find the closest landing sequence.
	MAV_CMD_DO_LAND_START MAV_CMD = 189
	// Mission command to perform a landing from a rally point.
	MAV_CMD_DO_RALLY_LAND MAV_CMD = 190
	// Mission command to safely abort an autonomous landing.
	MAV_CMD_DO_GO_AROUND MAV_CMD = 191
	// Reposition the vehicle to a specific WGS84 global position.
	MAV_CMD_DO_REPOSITION MAV_CMD = 192
	// If in a GPS controlled position mode, hold the current position or continue.
	MAV_CMD_DO_PAUSE_CONTINUE MAV_CMD = 193
	// Set moving direction to forward or reverse.
	MAV_CMD_DO_SET_REVERSE MAV_CMD = 194
	// Sets the region of interest (ROI) to a location. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal is not to react to this message.
	MAV_CMD_DO_SET_ROI_LOCATION MAV_CMD = 195
	// Sets the region of interest (ROI) to be toward next waypoint, with optional pitch/roll/yaw offset. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message.
	MAV_CMD_DO_SET_ROI_WPNEXT_OFFSET MAV_CMD = 196
	// Cancels any previous ROI command returning the vehicle/sensors to default flight characteristics. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message. After this command the gimbal manager should go back to manual input if available, and otherwise assume a neutral position.
	MAV_CMD_DO_SET_ROI_NONE MAV_CMD = 197
	// Mount tracks system with specified system ID. Determination of target vehicle position may be done with GLOBAL_POSITION_INT or any other means. This command can be sent to a gimbal manager but not to a gimbal device. A gimbal device is not to react to this message.
	MAV_CMD_DO_SET_ROI_SYSID MAV_CMD = 198
	// Control onboard camera system.
	MAV_CMD_DO_CONTROL_VIDEO MAV_CMD = 200
	// Sets the region of interest (ROI) for a sensor set or the vehicle itself. This can then be used by the vehicle's control system to control the vehicle attitude and the attitude of various sensors such as cameras.
	MAV_CMD_DO_SET_ROI MAV_CMD = 201
	// Configure digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	MAV_CMD_DO_DIGICAM_CONFIGURE MAV_CMD = 202
	// Control digital camera. This is a fallback message for systems that have not yet implemented PARAM_EXT_XXX messages and camera definition files (see https://mavlink.io/en/services/camera_def.html ).
	MAV_CMD_DO_DIGICAM_CONTROL MAV_CMD = 203
	// Mission command to configure a camera or antenna mount
	MAV_CMD_DO_MOUNT_CONFIGURE MAV_CMD = 204
	// Mission command to control a camera or antenna mount
	MAV_CMD_DO_MOUNT_CONTROL MAV_CMD = 205
	// Mission command to set camera trigger distance for this flight. The camera is triggered each time this distance is exceeded. This command can also be used to set the shutter integration time for the camera.
	MAV_CMD_DO_SET_CAM_TRIGG_DIST MAV_CMD = 206
	// Mission command to enable the geofence
	MAV_CMD_DO_FENCE_ENABLE MAV_CMD = 207
	// Mission command to trigger a parachute
	MAV_CMD_DO_PARACHUTE MAV_CMD = 208
	// Mission command to perform motor test.
	MAV_CMD_DO_MOTOR_TEST MAV_CMD = 209
	// Change to/from inverted flight.
	MAV_CMD_DO_INVERTED_FLIGHT MAV_CMD = 210
	// Sets a desired vehicle turn angle and speed change.
	MAV_CMD_NAV_SET_YAW_SPEED MAV_CMD = 213
	// Mission command to set camera trigger interval for this flight. If triggering is enabled, the camera is triggered each time this interval expires. This command can also be used to set the shutter integration time for the camera.
	MAV_CMD_DO_SET_CAM_TRIGG_INTERVAL MAV_CMD = 214
	// Mission command to control a camera or antenna mount, using a quaternion as reference.
	MAV_CMD_DO_MOUNT_CONTROL_QUAT MAV_CMD = 220
	// set id of master controller
	MAV_CMD_DO_GUIDED_MASTER MAV_CMD = 221
	// Set limits for external control
	MAV_CMD_DO_GUIDED_LIMITS MAV_CMD = 222
	// Control vehicle engine. This is interpreted by the vehicles engine controller to change the target engine state. It is intended for vehicles with internal combustion engines
	MAV_CMD_DO_ENGINE_CONTROL MAV_CMD = 223
	// Set the mission item with sequence number seq as current item. This means that the MAV will continue to this mission item on the shortest path (not following the mission items in-between).
	MAV_CMD_DO_SET_MISSION_CURRENT MAV_CMD = 224
	// NOP - This command is only used to mark the upper limit of the DO commands in the enumeration
	MAV_CMD_DO_LAST MAV_CMD = 240
	// Trigger calibration. This command will be only accepted if in pre-flight mode. Except for Temperature Calibration, only one sensor should be set in a single message and all others should be zero.
	MAV_CMD_PREFLIGHT_CALIBRATION MAV_CMD = 241
	// Set sensor offsets. This command will be only accepted if in pre-flight mode.
	MAV_CMD_PREFLIGHT_SET_SENSOR_OFFSETS MAV_CMD = 242
	// Trigger UAVCAN config. This command will be only accepted if in pre-flight mode.
	MAV_CMD_PREFLIGHT_UAVCAN MAV_CMD = 243
	// Request storage of different parameter values and logs. This command will be only accepted if in pre-flight mode.
	MAV_CMD_PREFLIGHT_STORAGE MAV_CMD = 245
	// Request the reboot or shutdown of system components.
	MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN MAV_CMD = 246
	// Override current mission with command to pause mission, pause mission and move to position, continue/resume mission. When param 1 indicates that the mission is paused (MAV_GOTO_DO_HOLD), param 2 defines whether it holds in place or moves to another position.
	MAV_CMD_OVERRIDE_GOTO MAV_CMD = 252
	// start running a mission
	MAV_CMD_MISSION_START MAV_CMD = 300
	// Arms / Disarms a component
	MAV_CMD_COMPONENT_ARM_DISARM MAV_CMD = 400
	// Turns illuminators ON/OFF. An illuminator is a light source that is used for lighting up dark areas external to the sytstem: e.g. a torch or searchlight (as opposed to a light source for illuminating the system itself, e.g. an indicator light).
	MAV_CMD_ILLUMINATOR_ON_OFF MAV_CMD = 405
	// Request the home position from the vehicle.
	MAV_CMD_GET_HOME_POSITION MAV_CMD = 410
	// Inject artificial failure for testing purposes. Note that autopilots should implement an additional protection before accepting this command such as a specific param setting.
	MAV_CMD_INJECT_FAILURE MAV_CMD = 420
	// Starts receiver pairing.
	MAV_CMD_START_RX_PAIR MAV_CMD = 500
	// Request the interval between messages for a particular MAVLink message ID. The receiver should ACK the command and then emit its response in a MESSAGE_INTERVAL message.
	MAV_CMD_GET_MESSAGE_INTERVAL MAV_CMD = 510
	// Set the interval between messages for a particular MAVLink message ID. This interface replaces REQUEST_DATA_STREAM.
	MAV_CMD_SET_MESSAGE_INTERVAL MAV_CMD = 511
	// Request the target system(s) emit a single instance of a specified message (i.e. a "one-shot" version of MAV_CMD_SET_MESSAGE_INTERVAL).
	MAV_CMD_REQUEST_MESSAGE MAV_CMD = 512
	// Request MAVLink protocol version compatibility
	MAV_CMD_REQUEST_PROTOCOL_VERSION MAV_CMD = 519
	// Request autopilot capabilities. The receiver should ACK the command and then emit its capabilities in an AUTOPILOT_VERSION message
	MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES MAV_CMD = 520
	// Request camera information (CAMERA_INFORMATION).
	MAV_CMD_REQUEST_CAMERA_INFORMATION MAV_CMD = 521
	// Request camera settings (CAMERA_SETTINGS).
	MAV_CMD_REQUEST_CAMERA_SETTINGS MAV_CMD = 522
	// Request storage information (STORAGE_INFORMATION). Use the command's target_component to target a specific component's storage.
	MAV_CMD_REQUEST_STORAGE_INFORMATION MAV_CMD = 525
	// Format a storage medium. Once format is complete, a STORAGE_INFORMATION message is sent. Use the command's target_component to target a specific component's storage.
	MAV_CMD_STORAGE_FORMAT MAV_CMD = 526
	// Request camera capture status (CAMERA_CAPTURE_STATUS)
	MAV_CMD_REQUEST_CAMERA_CAPTURE_STATUS MAV_CMD = 527
	// Request flight information (FLIGHT_INFORMATION)
	MAV_CMD_REQUEST_FLIGHT_INFORMATION MAV_CMD = 528
	// Reset all camera settings to Factory Default
	MAV_CMD_RESET_CAMERA_SETTINGS MAV_CMD = 529
	// Set camera running mode. Use NaN for reserved values. GCS will send a MAV_CMD_REQUEST_VIDEO_STREAM_STATUS command after a mode change if the camera supports video streaming.
	MAV_CMD_SET_CAMERA_MODE MAV_CMD = 530
	// Set camera zoom. Camera must respond with a CAMERA_SETTINGS message (on success). Use NaN for reserved values.
	MAV_CMD_SET_CAMERA_ZOOM MAV_CMD = 531
	// Set camera focus. Camera must respond with a CAMERA_SETTINGS message (on success). Use NaN for reserved values.
	MAV_CMD_SET_CAMERA_FOCUS MAV_CMD = 532
	// Tagged jump target. Can be jumped to with MAV_CMD_DO_JUMP_TAG.
	MAV_CMD_JUMP_TAG MAV_CMD = 600
	// Jump to the matching tag in the mission list. Repeat this action for the specified number of times. A mission should contain a single matching tag for each jump. If this is not the case then a jump to a missing tag should complete the mission, and a jump where there are multiple matching tags should always select the one with the lowest mission sequence number.
	MAV_CMD_DO_JUMP_TAG MAV_CMD = 601
	// High level setpoint to be sent to a gimbal manager to set a gimbal attitude. It is possible to set combinations of the values below. E.g. an angle as well as a desired angular rate can be used to get to this angle at a certain angular rate, or an angular rate only will result in continuous turning. NaN is to be used to signal unset. Note: a gimbal is never to react to this command but only the gimbal manager.
	MAV_CMD_DO_GIMBAL_MANAGER_ATTITUDE MAV_CMD = 1000
	// If the gimbal manager supports visual tracking (GIMBAL_MANAGER_CAP_FLAGS_HAS_TRACKING_POINT is set), this command allows to initiate the tracking. Such a tracking gimbal manager would usually be an integrated camera/gimbal, or alternatively a companion computer connected to a camera.
	MAV_CMD_DO_GIMBAL_MANAGER_TRACK_POINT MAV_CMD = 1001
	// If the gimbal supports visual tracking (GIMBAL_MANAGER_CAP_FLAGS_HAS_TRACKING_RECTANGLE is set), this command allows to initiate the tracking. Such a tracking gimbal manager would usually be an integrated camera/gimbal, or alternatively a companion computer connected to a camera.
	MAV_CMD_DO_GIMBAL_MANAGER_TRACK_RECTANGLE MAV_CMD = 1002
	// Start image capture sequence. Sends CAMERA_IMAGE_CAPTURED after each capture. Use NaN for reserved values.
	MAV_CMD_IMAGE_START_CAPTURE MAV_CMD = 2000
	// Stop image capture sequence Use NaN for reserved values.
	MAV_CMD_IMAGE_STOP_CAPTURE MAV_CMD = 2001
	// Re-request a CAMERA_IMAGE_CAPTURE message. Use NaN for reserved values.
	MAV_CMD_REQUEST_CAMERA_IMAGE_CAPTURE MAV_CMD = 2002
	// Enable or disable on-board camera triggering system.
	MAV_CMD_DO_TRIGGER_CONTROL MAV_CMD = 2003
	// Starts video capture (recording). Use NaN for reserved values.
	MAV_CMD_VIDEO_START_CAPTURE MAV_CMD = 2500
	// Stop the current video capture (recording). Use NaN for reserved values.
	MAV_CMD_VIDEO_STOP_CAPTURE MAV_CMD = 2501
	// Start video streaming
	MAV_CMD_VIDEO_START_STREAMING MAV_CMD = 2502
	// Stop the given video stream
	MAV_CMD_VIDEO_STOP_STREAMING MAV_CMD = 2503
	// Request video stream information (VIDEO_STREAM_INFORMATION)
	MAV_CMD_REQUEST_VIDEO_STREAM_INFORMATION MAV_CMD = 2504
	// Request video stream status (VIDEO_STREAM_STATUS)
	MAV_CMD_REQUEST_VIDEO_STREAM_STATUS MAV_CMD = 2505
	// Request to start streaming logging data over MAVLink (see also LOGGING_DATA message)
	MAV_CMD_LOGGING_START MAV_CMD = 2510
	// Request to stop streaming log data over MAVLink
	MAV_CMD_LOGGING_STOP MAV_CMD = 2511
	//
	MAV_CMD_AIRFRAME_CONFIGURATION MAV_CMD = 2520
	// Request to start/stop transmitting over the high latency telemetry
	MAV_CMD_CONTROL_HIGH_LATENCY MAV_CMD = 2600
	// Create a panorama at the current position
	MAV_CMD_PANORAMA_CREATE MAV_CMD = 2800
	// Request VTOL transition
	MAV_CMD_DO_VTOL_TRANSITION MAV_CMD = 3000
	// Request authorization to arm the vehicle to a external entity, the arm authorizer is responsible to request all data that is needs from the vehicle before authorize or deny the request. If approved the progress of command_ack message should be set with period of time that this authorization is valid in seconds or in case it was denied it should be set with one of the reasons in ARM_AUTH_DENIED_REASON.
	MAV_CMD_ARM_AUTHORIZATION_REQUEST MAV_CMD = 3001
	// This command sets the submode to standard guided when vehicle is in guided mode. The vehicle holds position and altitude and the user can input the desired velocities along all three axes.
	MAV_CMD_SET_GUIDED_SUBMODE_STANDARD MAV_CMD = 4000
	// This command sets submode circle when vehicle is in guided mode. Vehicle flies along a circle facing the center of the circle. The user can input the velocity along the circle and change the radius. If no input is given the vehicle will hold position.
	MAV_CMD_SET_GUIDED_SUBMODE_CIRCLE MAV_CMD = 4001
	// Delay mission state machine until gate has been reached.
	MAV_CMD_CONDITION_GATE MAV_CMD = 4501
	// Fence return point. There can only be one fence return point.
	MAV_CMD_NAV_FENCE_RETURN_POINT MAV_CMD = 5000
	// Fence vertex for an inclusion polygon (the polygon must not be self-intersecting). The vehicle must stay within this area. Minimum of 3 vertices required.
	MAV_CMD_NAV_FENCE_POLYGON_VERTEX_INCLUSION MAV_CMD = 5001
	// Fence vertex for an exclusion polygon (the polygon must not be self-intersecting). The vehicle must stay outside this area. Minimum of 3 vertices required.
	MAV_CMD_NAV_FENCE_POLYGON_VERTEX_EXCLUSION MAV_CMD = 5002
	// Circular fence area. The vehicle must stay inside this area.
	MAV_CMD_NAV_FENCE_CIRCLE_INCLUSION MAV_CMD = 5003
	// Circular fence area. The vehicle must stay outside this area.
	MAV_CMD_NAV_FENCE_CIRCLE_EXCLUSION MAV_CMD = 5004
	// Rally point. You can have multiple rally points defined.
	MAV_CMD_NAV_RALLY_POINT MAV_CMD = 5100
	// Commands the vehicle to respond with a sequence of messages UAVCAN_NODE_INFO, one message per every UAVCAN node that is online. Note that some of the response messages can be lost, which the receiver can detect easily by checking whether every received UAVCAN_NODE_STATUS has a matching message UAVCAN_NODE_INFO received earlier; if not, this command should be sent again in order to request re-transmission of the node information messages.
	MAV_CMD_UAVCAN_GET_NODE_INFO MAV_CMD = 5200
	// Deploy payload on a Lat / Lon / Alt position. This includes the navigation to reach the required release position and velocity.
	MAV_CMD_PAYLOAD_PREPARE_DEPLOY MAV_CMD = 30001
	// Control the payload deployment.
	MAV_CMD_PAYLOAD_CONTROL_DEPLOY MAV_CMD = 30002
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_1 MAV_CMD = 31000
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_2 MAV_CMD = 31001
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_3 MAV_CMD = 31002
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_4 MAV_CMD = 31003
	// User defined waypoint item. Ground Station will show the Vehicle as flying through this item.
	MAV_CMD_WAYPOINT_USER_5 MAV_CMD = 31004
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_1 MAV_CMD = 31005
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_2 MAV_CMD = 31006
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_3 MAV_CMD = 31007
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_4 MAV_CMD = 31008
	// User defined spatial item. Ground Station will not show the Vehicle as flying through this item. Example: ROI item.
	MAV_CMD_SPATIAL_USER_5 MAV_CMD = 31009
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_1 MAV_CMD = 31010
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_2 MAV_CMD = 31011
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_3 MAV_CMD = 31012
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_4 MAV_CMD = 31013
	// User defined command. Ground Station will not show the Vehicle as flying through this item. Example: MAV_CMD_DO_SET_PARAMETER item.
	MAV_CMD_USER_5 MAV_CMD = 31014
)

// ACK / NACK / ERROR values as a result of MAV_CMDs and for mission item transmission.
type MAV_CMD_ACK int

const (
	// Command / mission item is ok.
	MAV_CMD_ACK_OK MAV_CMD_ACK = 0
	// Generic error message if none of the other reasons fails or if no detailed error reporting is implemented.
	MAV_CMD_ACK_ERR_FAIL MAV_CMD_ACK = 1
	// The system is refusing to accept this command from this source / communication partner.
	MAV_CMD_ACK_ERR_ACCESS_DENIED MAV_CMD_ACK = 2
	// Command or mission item is not supported, other commands would be accepted.
	MAV_CMD_ACK_ERR_NOT_SUPPORTED MAV_CMD_ACK = 3
	// The coordinate frame of this command / mission item is not supported.
	MAV_CMD_ACK_ERR_COORDINATE_FRAME_NOT_SUPPORTED MAV_CMD_ACK = 4
	// The coordinate frame of this command is ok, but he coordinate values exceed the safety limits of this system. This is a generic error, please use the more specific error messages below if possible.
	MAV_CMD_ACK_ERR_COORDINATES_OUT_OF_RANGE MAV_CMD_ACK = 5
	// The X or latitude value is out of range.
	MAV_CMD_ACK_ERR_X_LAT_OUT_OF_RANGE MAV_CMD_ACK = 6
	// The Y or longitude value is out of range.
	MAV_CMD_ACK_ERR_Y_LON_OUT_OF_RANGE MAV_CMD_ACK = 7
	// The Z or altitude value is out of range.
	MAV_CMD_ACK_ERR_Z_ALT_OUT_OF_RANGE MAV_CMD_ACK = 8
)

// Possible actions an aircraft can take to avoid a collision.
type MAV_COLLISION_ACTION int

const (
	// Ignore any potential collisions
	MAV_COLLISION_ACTION_NONE MAV_COLLISION_ACTION = 0
	// Report potential collision
	MAV_COLLISION_ACTION_REPORT MAV_COLLISION_ACTION = 1
	// Ascend or Descend to avoid threat
	MAV_COLLISION_ACTION_ASCEND_OR_DESCEND MAV_COLLISION_ACTION = 2
	// Move horizontally to avoid threat
	MAV_COLLISION_ACTION_MOVE_HORIZONTALLY MAV_COLLISION_ACTION = 3
	// Aircraft to move perpendicular to the collision's velocity vector
	MAV_COLLISION_ACTION_MOVE_PERPENDICULAR MAV_COLLISION_ACTION = 4
	// Aircraft to fly directly back to its launch point
	MAV_COLLISION_ACTION_RTL MAV_COLLISION_ACTION = 5
	// Aircraft to stop in place
	MAV_COLLISION_ACTION_HOVER MAV_COLLISION_ACTION = 6
)

// Source of information about this collision.
type MAV_COLLISION_SRC int

const (
	// ID field references ADSB_VEHICLE packets
	MAV_COLLISION_SRC_ADSB MAV_COLLISION_SRC = 0
	// ID field references MAVLink SRC ID
	MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT MAV_COLLISION_SRC = 1
)

// Aircraft-rated danger from this threat.
type MAV_COLLISION_THREAT_LEVEL int

const (
	// Not a threat
	MAV_COLLISION_THREAT_LEVEL_NONE MAV_COLLISION_THREAT_LEVEL = 0
	// Craft is mildly concerned about this threat
	MAV_COLLISION_THREAT_LEVEL_LOW MAV_COLLISION_THREAT_LEVEL = 1
	// Craft is panicking, and may take actions to avoid threat
	MAV_COLLISION_THREAT_LEVEL_HIGH MAV_COLLISION_THREAT_LEVEL = 2
)

// Component ids (values) for the different types and instances of onboard hardware/software that might make up a MAVLink system (autopilot, cameras, servos, GPS systems, avoidance systems etc.).      Components must use the appropriate ID in their source address when sending messages. Components can also use IDs to determine if they are the intended recipient of an incoming message. The MAV_COMP_ID_ALL value is used to indicate messages that must be processed by all components.      When creating new entries, components that can have multiple instances (e.g. cameras, servos etc.) should be allocated sequential values. An appropriate number of values should be left free after these components to allow the number of instances to be expanded.
type MAV_COMPONENT int

const (
	// Target id (target_component) used to broadcast messages to all components of the receiving system. Components should attempt to process messages with this component ID and forward to components on any other interfaces. Note: This is not a valid *source* component id for a message.
	MAV_COMP_ID_ALL MAV_COMPONENT = 0
	// System flight controller component ("autopilot"). Only one autopilot is expected in a particular system.
	MAV_COMP_ID_AUTOPILOT1 MAV_COMPONENT = 1
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER1 MAV_COMPONENT = 25
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER2 MAV_COMPONENT = 26
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER3 MAV_COMPONENT = 27
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER4 MAV_COMPONENT = 28
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER5 MAV_COMPONENT = 29
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER6 MAV_COMPONENT = 30
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER7 MAV_COMPONENT = 31
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER8 MAV_COMPONENT = 32
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER9 MAV_COMPONENT = 33
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER10 MAV_COMPONENT = 34
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER11 MAV_COMPONENT = 35
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER12 MAV_COMPONENT = 36
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER13 MAV_COMPONENT = 37
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER14 MAV_COMPONENT = 38
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER15 MAV_COMPONENT = 39
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USE16 MAV_COMPONENT = 40
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER17 MAV_COMPONENT = 41
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER18 MAV_COMPONENT = 42
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER19 MAV_COMPONENT = 43
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER20 MAV_COMPONENT = 44
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER21 MAV_COMPONENT = 45
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER22 MAV_COMPONENT = 46
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER23 MAV_COMPONENT = 47
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER24 MAV_COMPONENT = 48
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER25 MAV_COMPONENT = 49
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER26 MAV_COMPONENT = 50
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER27 MAV_COMPONENT = 51
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER28 MAV_COMPONENT = 52
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER29 MAV_COMPONENT = 53
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER30 MAV_COMPONENT = 54
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER31 MAV_COMPONENT = 55
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER32 MAV_COMPONENT = 56
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER33 MAV_COMPONENT = 57
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER34 MAV_COMPONENT = 58
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER35 MAV_COMPONENT = 59
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER36 MAV_COMPONENT = 60
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER37 MAV_COMPONENT = 61
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER38 MAV_COMPONENT = 62
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER39 MAV_COMPONENT = 63
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER40 MAV_COMPONENT = 64
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER41 MAV_COMPONENT = 65
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER42 MAV_COMPONENT = 66
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER43 MAV_COMPONENT = 67
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER44 MAV_COMPONENT = 68
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER45 MAV_COMPONENT = 69
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER46 MAV_COMPONENT = 70
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER47 MAV_COMPONENT = 71
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER48 MAV_COMPONENT = 72
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER49 MAV_COMPONENT = 73
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER50 MAV_COMPONENT = 74
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER51 MAV_COMPONENT = 75
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER52 MAV_COMPONENT = 76
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER53 MAV_COMPONENT = 77
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER54 MAV_COMPONENT = 78
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER55 MAV_COMPONENT = 79
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER56 MAV_COMPONENT = 80
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER57 MAV_COMPONENT = 81
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER58 MAV_COMPONENT = 82
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER59 MAV_COMPONENT = 83
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER60 MAV_COMPONENT = 84
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER61 MAV_COMPONENT = 85
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER62 MAV_COMPONENT = 86
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER63 MAV_COMPONENT = 87
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER64 MAV_COMPONENT = 88
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER65 MAV_COMPONENT = 89
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER66 MAV_COMPONENT = 90
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER67 MAV_COMPONENT = 91
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER68 MAV_COMPONENT = 92
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER69 MAV_COMPONENT = 93
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER70 MAV_COMPONENT = 94
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER71 MAV_COMPONENT = 95
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER72 MAV_COMPONENT = 96
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER73 MAV_COMPONENT = 97
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER74 MAV_COMPONENT = 98
	// Id for a component on privately managed MAVLink network. Can be used for any purpose but may not be published by components outside of the private network.
	MAV_COMP_ID_USER75 MAV_COMPONENT = 99
	// Camera #1.
	MAV_COMP_ID_CAMERA MAV_COMPONENT = 100
	// Camera #2.
	MAV_COMP_ID_CAMERA2 MAV_COMPONENT = 101
	// Camera #3.
	MAV_COMP_ID_CAMERA3 MAV_COMPONENT = 102
	// Camera #4.
	MAV_COMP_ID_CAMERA4 MAV_COMPONENT = 103
	// Camera #5.
	MAV_COMP_ID_CAMERA5 MAV_COMPONENT = 104
	// Camera #6.
	MAV_COMP_ID_CAMERA6 MAV_COMPONENT = 105
	// Servo #1.
	MAV_COMP_ID_SERVO1 MAV_COMPONENT = 140
	// Servo #2.
	MAV_COMP_ID_SERVO2 MAV_COMPONENT = 141
	// Servo #3.
	MAV_COMP_ID_SERVO3 MAV_COMPONENT = 142
	// Servo #4.
	MAV_COMP_ID_SERVO4 MAV_COMPONENT = 143
	// Servo #5.
	MAV_COMP_ID_SERVO5 MAV_COMPONENT = 144
	// Servo #6.
	MAV_COMP_ID_SERVO6 MAV_COMPONENT = 145
	// Servo #7.
	MAV_COMP_ID_SERVO7 MAV_COMPONENT = 146
	// Servo #8.
	MAV_COMP_ID_SERVO8 MAV_COMPONENT = 147
	// Servo #9.
	MAV_COMP_ID_SERVO9 MAV_COMPONENT = 148
	// Servo #10.
	MAV_COMP_ID_SERVO10 MAV_COMPONENT = 149
	// Servo #11.
	MAV_COMP_ID_SERVO11 MAV_COMPONENT = 150
	// Servo #12.
	MAV_COMP_ID_SERVO12 MAV_COMPONENT = 151
	// Servo #13.
	MAV_COMP_ID_SERVO13 MAV_COMPONENT = 152
	// Servo #14.
	MAV_COMP_ID_SERVO14 MAV_COMPONENT = 153
	// Gimbal #1.
	MAV_COMP_ID_GIMBAL MAV_COMPONENT = 154
	// Logging component.
	MAV_COMP_ID_LOG MAV_COMPONENT = 155
	// Automatic Dependent Surveillance-Broadcast (ADS-B) component.
	MAV_COMP_ID_ADSB MAV_COMPONENT = 156
	// On Screen Display (OSD) devices for video links.
	MAV_COMP_ID_OSD MAV_COMPONENT = 157
	// Generic autopilot peripheral component ID. Meant for devices that do not implement the parameter microservice.
	MAV_COMP_ID_PERIPHERAL MAV_COMPONENT = 158
	// Gimbal ID for QX1.
	MAV_COMP_ID_QX1_GIMBAL MAV_COMPONENT = 159
	// FLARM collision alert component.
	MAV_COMP_ID_FLARM MAV_COMPONENT = 160
	// Gimbal #2.
	MAV_COMP_ID_GIMBAL2 MAV_COMPONENT = 171
	// Gimbal #3.
	MAV_COMP_ID_GIMBAL3 MAV_COMPONENT = 172
	// Gimbal #4
	MAV_COMP_ID_GIMBAL4 MAV_COMPONENT = 173
	// Gimbal #5.
	MAV_COMP_ID_GIMBAL5 MAV_COMPONENT = 174
	// Gimbal #6.
	MAV_COMP_ID_GIMBAL6 MAV_COMPONENT = 175
	// Component that can generate/supply a mission flight plan (e.g. GCS or developer API).
	MAV_COMP_ID_MISSIONPLANNER MAV_COMPONENT = 190
	// Component that finds an optimal path between points based on a certain constraint (e.g. minimum snap, shortest path, cost, etc.).
	MAV_COMP_ID_PATHPLANNER MAV_COMPONENT = 195
	// Component that plans a collision free path between two points.
	MAV_COMP_ID_OBSTACLE_AVOIDANCE MAV_COMPONENT = 196
	// Component that provides position estimates using VIO techniques.
	MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY MAV_COMPONENT = 197
	// Component that manages pairing of vehicle and GCS.
	MAV_COMP_ID_PAIRING_MANAGER MAV_COMPONENT = 198
	// Inertial Measurement Unit (IMU) #1.
	MAV_COMP_ID_IMU MAV_COMPONENT = 200
	// Inertial Measurement Unit (IMU) #2.
	MAV_COMP_ID_IMU_2 MAV_COMPONENT = 201
	// Inertial Measurement Unit (IMU) #3.
	MAV_COMP_ID_IMU_3 MAV_COMPONENT = 202
	// GPS #1.
	MAV_COMP_ID_GPS MAV_COMPONENT = 220
	// GPS #2.
	MAV_COMP_ID_GPS2 MAV_COMPONENT = 221
	// Component to bridge MAVLink to UDP (i.e. from a UART).
	MAV_COMP_ID_UDP_BRIDGE MAV_COMPONENT = 240
	// Component to bridge to UART (i.e. from UDP).
	MAV_COMP_ID_UART_BRIDGE MAV_COMPONENT = 241
	// Component handling TUNNEL messages (e.g. vendor specific GUI of a component).
	MAV_COMP_ID_TUNNEL_NODE MAV_COMPONENT = 242
	// Component for handling system messages (e.g. to ARM, takeoff, etc.).
	MAV_COMP_ID_SYSTEM_CONTROL MAV_COMPONENT = 250
)

// A data stream is not a fixed set of messages, but rather a     recommendation to the autopilot software. Individual autopilots may or may not obey     the recommended messages.
type MAV_DATA_STREAM int

const (
	// Enable all data streams
	MAV_DATA_STREAM_ALL MAV_DATA_STREAM = 0
	// Enable IMU_RAW, GPS_RAW, GPS_STATUS packets.
	MAV_DATA_STREAM_RAW_SENSORS MAV_DATA_STREAM = 1
	// Enable GPS_STATUS, CONTROL_STATUS, AUX_STATUS
	MAV_DATA_STREAM_EXTENDED_STATUS MAV_DATA_STREAM = 2
	// Enable RC_CHANNELS_SCALED, RC_CHANNELS_RAW, SERVO_OUTPUT_RAW
	MAV_DATA_STREAM_RC_CHANNELS MAV_DATA_STREAM = 3
	// Enable ATTITUDE_CONTROLLER_OUTPUT, POSITION_CONTROLLER_OUTPUT, NAV_CONTROLLER_OUTPUT.
	MAV_DATA_STREAM_RAW_CONTROLLER MAV_DATA_STREAM = 4
	// Enable LOCAL_POSITION, GLOBAL_POSITION/GLOBAL_POSITION_INT messages.
	MAV_DATA_STREAM_POSITION MAV_DATA_STREAM = 6
	// Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA1 MAV_DATA_STREAM = 10
	// Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA2 MAV_DATA_STREAM = 11
	// Dependent on the autopilot
	MAV_DATA_STREAM_EXTRA3 MAV_DATA_STREAM = 12
)

// Enumeration of distance sensor types
type MAV_DISTANCE_SENSOR int

const (
	// Laser rangefinder, e.g. LightWare SF02/F or PulsedLight units
	MAV_DISTANCE_SENSOR_LASER MAV_DISTANCE_SENSOR = 0
	// Ultrasound rangefinder, e.g. MaxBotix units
	MAV_DISTANCE_SENSOR_ULTRASOUND MAV_DISTANCE_SENSOR = 1
	// Infrared rangefinder, e.g. Sharp units
	MAV_DISTANCE_SENSOR_INFRARED MAV_DISTANCE_SENSOR = 2
	// Radar type, e.g. uLanding units
	MAV_DISTANCE_SENSOR_RADAR MAV_DISTANCE_SENSOR = 3
	// Broken or unknown type, e.g. analog units
	MAV_DISTANCE_SENSOR_UNKNOWN MAV_DISTANCE_SENSOR = 4
)

// Bitmap of options for the MAV_CMD_DO_REPOSITION
type MAV_DO_REPOSITION_FLAGS int

const (
	// The aircraft should immediately transition into guided. This should not be set for follow me applications
	MAV_DO_REPOSITION_FLAGS_CHANGE_MODE MAV_DO_REPOSITION_FLAGS = 1
)

// Enumeration of estimator types
type MAV_ESTIMATOR_TYPE int

const (
	// Unknown type of the estimator.
	MAV_ESTIMATOR_TYPE_UNKNOWN MAV_ESTIMATOR_TYPE = 0
	// This is a naive estimator without any real covariance feedback.
	MAV_ESTIMATOR_TYPE_NAIVE MAV_ESTIMATOR_TYPE = 1
	// Computer vision based estimate. Might be up to scale.
	MAV_ESTIMATOR_TYPE_VISION MAV_ESTIMATOR_TYPE = 2
	// Visual-inertial estimate.
	MAV_ESTIMATOR_TYPE_VIO MAV_ESTIMATOR_TYPE = 3
	// Plain GPS estimate.
	MAV_ESTIMATOR_TYPE_GPS MAV_ESTIMATOR_TYPE = 4
	// Estimator integrating GPS and inertial sensing.
	MAV_ESTIMATOR_TYPE_GPS_INS MAV_ESTIMATOR_TYPE = 5
	// Estimate from external motion capturing system.
	MAV_ESTIMATOR_TYPE_MOCAP MAV_ESTIMATOR_TYPE = 6
	// Estimator based on lidar sensor input.
	MAV_ESTIMATOR_TYPE_LIDAR MAV_ESTIMATOR_TYPE = 7
	// Estimator on autopilot.
	MAV_ESTIMATOR_TYPE_AUTOPILOT MAV_ESTIMATOR_TYPE = 8
)

//
type MAV_FRAME int

const (
	// Global (WGS84) coordinate frame + MSL altitude. First value / x: latitude, second value / y: longitude, third value / z: positive altitude over mean sea level (MSL).
	MAV_FRAME_GLOBAL MAV_FRAME = 0
	// Local coordinate frame, Z-down (x: north, y: east, z: down).
	MAV_FRAME_LOCAL_NED MAV_FRAME = 1
	// NOT a coordinate frame, indicates a mission command.
	MAV_FRAME_MISSION MAV_FRAME = 2
	// Global (WGS84) coordinate frame + altitude relative to the home position. First value / x: latitude, second value / y: longitude, third value / z: positive altitude with 0 being at the altitude of the home location.
	MAV_FRAME_GLOBAL_RELATIVE_ALT MAV_FRAME = 3
	// Local coordinate frame, Z-up (x: east, y: north, z: up).
	MAV_FRAME_LOCAL_ENU MAV_FRAME = 4
	// Global (WGS84) coordinate frame (scaled) + MSL altitude. First value / x: latitude in degrees*1.0e-7, second value / y: longitude in degrees*1.0e-7, third value / z: positive altitude over mean sea level (MSL).
	MAV_FRAME_GLOBAL_INT MAV_FRAME = 5
	// Global (WGS84) coordinate frame (scaled) + altitude relative to the home position. First value / x: latitude in degrees*10e-7, second value / y: longitude in degrees*10e-7, third value / z: positive altitude with 0 being at the altitude of the home location.
	MAV_FRAME_GLOBAL_RELATIVE_ALT_INT MAV_FRAME = 6
	// Offset to the current local frame. Anything expressed in this frame should be added to the current local frame position.
	MAV_FRAME_LOCAL_OFFSET_NED MAV_FRAME = 7
	// Setpoint in body NED frame. This makes sense if all position control is externalized - e.g. useful to command 2 m/s^2 acceleration to the right.
	MAV_FRAME_BODY_NED MAV_FRAME = 8
	// Offset in body NED frame. This makes sense if adding setpoints to the current flight path, to avoid an obstacle - e.g. useful to command 2 m/s^2 acceleration to the east.
	MAV_FRAME_BODY_OFFSET_NED MAV_FRAME = 9
	// Global (WGS84) coordinate frame with AGL altitude (at the waypoint coordinate). First value / x: latitude in degrees, second value / y: longitude in degrees, third value / z: positive altitude in meters with 0 being at ground level in terrain model.
	MAV_FRAME_GLOBAL_TERRAIN_ALT MAV_FRAME = 10
	// Global (WGS84) coordinate frame (scaled) with AGL altitude (at the waypoint coordinate). First value / x: latitude in degrees*10e-7, second value / y: longitude in degrees*10e-7, third value / z: positive altitude in meters with 0 being at ground level in terrain model.
	MAV_FRAME_GLOBAL_TERRAIN_ALT_INT MAV_FRAME = 11
	// Body fixed frame of reference, Z-down (x: forward, y: right, z: down).
	MAV_FRAME_BODY_FRD MAV_FRAME = 12
	// Body fixed frame of reference, Z-up (x: forward, y: left, z: up).
	MAV_FRAME_BODY_FLU MAV_FRAME = 13
	// Odometry local coordinate frame of data given by a motion capture system, Z-down (x: north, y: east, z: down).
	MAV_FRAME_MOCAP_NED MAV_FRAME = 14
	// Odometry local coordinate frame of data given by a motion capture system, Z-up (x: east, y: north, z: up).
	MAV_FRAME_MOCAP_ENU MAV_FRAME = 15
	// Odometry local coordinate frame of data given by a vision estimation system, Z-down (x: north, y: east, z: down).
	MAV_FRAME_VISION_NED MAV_FRAME = 16
	// Odometry local coordinate frame of data given by a vision estimation system, Z-up (x: east, y: north, z: up).
	MAV_FRAME_VISION_ENU MAV_FRAME = 17
	// Odometry local coordinate frame of data given by an estimator running onboard the vehicle, Z-down (x: north, y: east, z: down).
	MAV_FRAME_ESTIM_NED MAV_FRAME = 18
	// Odometry local coordinate frame of data given by an estimator running onboard the vehicle, Z-up (x: east, y: noth, z: up).
	MAV_FRAME_ESTIM_ENU MAV_FRAME = 19
	// Forward, Right, Down coordinate frame. This is a local frame with Z-down and arbitrary F/R alignment (i.e. not aligned with NED/earth frame).
	MAV_FRAME_LOCAL_FRD MAV_FRAME = 20
	// Forward, Left, Up coordinate frame. This is a local frame with Z-up and arbitrary F/L alignment (i.e. not aligned with ENU/earth frame).
	MAV_FRAME_LOCAL_FLU MAV_FRAME = 21
)

// Actions that may be specified in MAV_CMD_OVERRIDE_GOTO to override mission execution.
type MAV_GOTO int

const (
	// Hold at the current position.
	MAV_GOTO_DO_HOLD MAV_GOTO = 0
	// Continue with the next item in mission execution.
	MAV_GOTO_DO_CONTINUE MAV_GOTO = 1
	// Hold at the current position of the system
	MAV_GOTO_HOLD_AT_CURRENT_POSITION MAV_GOTO = 2
	// Hold at the position specified in the parameters of the DO_HOLD action
	MAV_GOTO_HOLD_AT_SPECIFIED_POSITION MAV_GOTO = 3
)

// Enumeration of landed detector states
type MAV_LANDED_STATE int

const (
	// MAV landed state is unknown
	MAV_LANDED_STATE_UNDEFINED MAV_LANDED_STATE = 0
	// MAV is landed (on ground)
	MAV_LANDED_STATE_ON_GROUND MAV_LANDED_STATE = 1
	// MAV is in air
	MAV_LANDED_STATE_IN_AIR MAV_LANDED_STATE = 2
	// MAV currently taking off
	MAV_LANDED_STATE_TAKEOFF MAV_LANDED_STATE = 3
	// MAV currently landing
	MAV_LANDED_STATE_LANDING MAV_LANDED_STATE = 4
)

// Result of mission operation (in a MISSION_ACK message).
type MAV_MISSION_RESULT int

const (
	// mission accepted OK
	MAV_MISSION_ACCEPTED MAV_MISSION_RESULT = 0
	// Generic error / not accepting mission commands at all right now.
	MAV_MISSION_ERROR MAV_MISSION_RESULT = 1
	// Coordinate frame is not supported.
	MAV_MISSION_UNSUPPORTED_FRAME MAV_MISSION_RESULT = 2
	// Command is not supported.
	MAV_MISSION_UNSUPPORTED MAV_MISSION_RESULT = 3
	// Mission item exceeds storage space.
	MAV_MISSION_NO_SPACE MAV_MISSION_RESULT = 4
	// One of the parameters has an invalid value.
	MAV_MISSION_INVALID MAV_MISSION_RESULT = 5
	// param1 has an invalid value.
	MAV_MISSION_INVALID_PARAM1 MAV_MISSION_RESULT = 6
	// param2 has an invalid value.
	MAV_MISSION_INVALID_PARAM2 MAV_MISSION_RESULT = 7
	// param3 has an invalid value.
	MAV_MISSION_INVALID_PARAM3 MAV_MISSION_RESULT = 8
	// param4 has an invalid value.
	MAV_MISSION_INVALID_PARAM4 MAV_MISSION_RESULT = 9
	// x / param5 has an invalid value.
	MAV_MISSION_INVALID_PARAM5_X MAV_MISSION_RESULT = 10
	// y / param6 has an invalid value.
	MAV_MISSION_INVALID_PARAM6_Y MAV_MISSION_RESULT = 11
	// z / param7 has an invalid value.
	MAV_MISSION_INVALID_PARAM7 MAV_MISSION_RESULT = 12
	// Mission item received out of sequence
	MAV_MISSION_INVALID_SEQUENCE MAV_MISSION_RESULT = 13
	// Not accepting any mission commands from this communication partner.
	MAV_MISSION_DENIED MAV_MISSION_RESULT = 14
	// Current mission operation cancelled (e.g. mission upload, mission download).
	MAV_MISSION_OPERATION_CANCELLED MAV_MISSION_RESULT = 15
)

// Type of mission items being requested/sent in mission protocol.
type MAV_MISSION_TYPE int

const (
	// Items are mission commands for main mission.
	MAV_MISSION_TYPE_MISSION MAV_MISSION_TYPE = 0
	// Specifies GeoFence area(s). Items are MAV_CMD_NAV_FENCE_ GeoFence items.
	MAV_MISSION_TYPE_FENCE MAV_MISSION_TYPE = 1
	// Specifies the rally points for the vehicle. Rally points are alternative RTL points. Items are MAV_CMD_NAV_RALLY_POINT rally point items.
	MAV_MISSION_TYPE_RALLY MAV_MISSION_TYPE = 2
	// Only used in MISSION_CLEAR_ALL to clear all mission types.
	MAV_MISSION_TYPE_ALL MAV_MISSION_TYPE = 255
)

// These defines are predefined OR-combined mode flags. There is no need to use values from this enum, but it               simplifies the use of the mode flags. Note that manual input is enabled in all modes as a safety override.
type MAV_MODE int

const (
	// System is not ready to fly, booting, calibrating, etc. No flag is set.
	MAV_MODE_PREFLIGHT MAV_MODE = 0
	// System is allowed to be active, under assisted RC control.
	MAV_MODE_STABILIZE_DISARMED MAV_MODE = 80
	// System is allowed to be active, under assisted RC control.
	MAV_MODE_STABILIZE_ARMED MAV_MODE = 208
	// System is allowed to be active, under manual (RC) control, no stabilization
	MAV_MODE_MANUAL_DISARMED MAV_MODE = 64
	// System is allowed to be active, under manual (RC) control, no stabilization
	MAV_MODE_MANUAL_ARMED MAV_MODE = 192
	// System is allowed to be active, under autonomous control, manual setpoint
	MAV_MODE_GUIDED_DISARMED MAV_MODE = 88
	// System is allowed to be active, under autonomous control, manual setpoint
	MAV_MODE_GUIDED_ARMED MAV_MODE = 216
	// System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints)
	MAV_MODE_AUTO_DISARMED MAV_MODE = 92
	// System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints)
	MAV_MODE_AUTO_ARMED MAV_MODE = 220
	// UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only.
	MAV_MODE_TEST_DISARMED MAV_MODE = 66
	// UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only.
	MAV_MODE_TEST_ARMED MAV_MODE = 194
)

// These flags encode the MAV mode.
type MAV_MODE_FLAG int

const (
	// 0b10000000 MAV safety set to armed. Motors are enabled / running / can start. Ready to fly. Additional note: this flag is to be ignore when sent in the command MAV_CMD_DO_SET_MODE and MAV_CMD_COMPONENT_ARM_DISARM shall be used instead. The flag can still be used to report the armed state.
	MAV_MODE_FLAG_SAFETY_ARMED MAV_MODE_FLAG = 128
	// 0b01000000 remote control input is enabled.
	MAV_MODE_FLAG_MANUAL_INPUT_ENABLED MAV_MODE_FLAG = 64
	// 0b00100000 hardware in the loop simulation. All motors / actuators are blocked, but internal software is full operational.
	MAV_MODE_FLAG_HIL_ENABLED MAV_MODE_FLAG = 32
	// 0b00010000 system stabilizes electronically its attitude (and optionally position). It needs however further control inputs to move around.
	MAV_MODE_FLAG_STABILIZE_ENABLED MAV_MODE_FLAG = 16
	// 0b00001000 guided mode enabled, system flies waypoints / mission items.
	MAV_MODE_FLAG_GUIDED_ENABLED MAV_MODE_FLAG = 8
	// 0b00000100 autonomous mode enabled, system finds its own goal positions. Guided flag can be set or not, depends on the actual implementation.
	MAV_MODE_FLAG_AUTO_ENABLED MAV_MODE_FLAG = 4
	// 0b00000010 system has a test mode enabled. This flag is intended for temporary system tests and should not be used for stable implementations.
	MAV_MODE_FLAG_TEST_ENABLED MAV_MODE_FLAG = 2
	// 0b00000001 Reserved for future use.
	MAV_MODE_FLAG_CUSTOM_MODE_ENABLED MAV_MODE_FLAG = 1
)

// These values encode the bit positions of the decode position. These values can be used to read the value of a flag bit by combining the base_mode variable with AND with the flag position value. The result will be either 0 or 1, depending on if the flag is set or not.
type MAV_MODE_FLAG_DECODE_POSITION int

const (
	// First bit:  10000000
	MAV_MODE_FLAG_DECODE_POSITION_SAFETY MAV_MODE_FLAG_DECODE_POSITION = 128
	// Second bit: 01000000
	MAV_MODE_FLAG_DECODE_POSITION_MANUAL MAV_MODE_FLAG_DECODE_POSITION = 64
	// Third bit:  00100000
	MAV_MODE_FLAG_DECODE_POSITION_HIL MAV_MODE_FLAG_DECODE_POSITION = 32
	// Fourth bit: 00010000
	MAV_MODE_FLAG_DECODE_POSITION_STABILIZE MAV_MODE_FLAG_DECODE_POSITION = 16
	// Fifth bit:  00001000
	MAV_MODE_FLAG_DECODE_POSITION_GUIDED MAV_MODE_FLAG_DECODE_POSITION = 8
	// Sixth bit:   00000100
	MAV_MODE_FLAG_DECODE_POSITION_AUTO MAV_MODE_FLAG_DECODE_POSITION = 4
	// Seventh bit: 00000010
	MAV_MODE_FLAG_DECODE_POSITION_TEST MAV_MODE_FLAG_DECODE_POSITION = 2
	// Eighth bit: 00000001
	MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE MAV_MODE_FLAG_DECODE_POSITION = 1
)

// Enumeration of possible mount operation modes. This message is used by obsolete/deprecated gimbal messages.
type MAV_MOUNT_MODE int

const (
	// Load and keep safe position (Roll,Pitch,Yaw) from permant memory and stop stabilization
	MAV_MOUNT_MODE_RETRACT MAV_MOUNT_MODE = 0
	// Load and keep neutral position (Roll,Pitch,Yaw) from permanent memory.
	MAV_MOUNT_MODE_NEUTRAL MAV_MOUNT_MODE = 1
	// Load neutral position and start MAVLink Roll,Pitch,Yaw control with stabilization
	MAV_MOUNT_MODE_MAVLINK_TARGETING MAV_MOUNT_MODE = 2
	// Load neutral position and start RC Roll,Pitch,Yaw control with stabilization
	MAV_MOUNT_MODE_RC_TARGETING MAV_MOUNT_MODE = 3
	// Load neutral position and start to point to Lat,Lon,Alt
	MAV_MOUNT_MODE_GPS_POINT MAV_MOUNT_MODE = 4
	// Gimbal tracks system with specified system ID
	MAV_MOUNT_MODE_SYSID_TARGET MAV_MOUNT_MODE = 5
)

//
type MAV_ODID_AUTH_TYPE int

const (
	// No authentication type is specified.
	MAV_ODID_AUTH_TYPE_NONE MAV_ODID_AUTH_TYPE = 0
	// Signature for the UAS (Unmanned Aircraft System) ID.
	MAV_ODID_AUTH_TYPE_UAS_ID_SIGNATURE MAV_ODID_AUTH_TYPE = 1
	// Signature for the Operator ID.
	MAV_ODID_AUTH_TYPE_OPERATOR_ID_SIGNATURE MAV_ODID_AUTH_TYPE = 2
	// Signature for the entire message set.
	MAV_ODID_AUTH_TYPE_MESSAGE_SET_SIGNATURE MAV_ODID_AUTH_TYPE = 3
	// Authentication is provided by Network Remote ID.
	MAV_ODID_AUTH_TYPE_NETWORK_REMOTE_ID MAV_ODID_AUTH_TYPE = 4
)

//
type MAV_ODID_DESC_TYPE int

const (
	// Free-form text description of the purpose of the flight.
	MAV_ODID_DESC_TYPE_TEXT MAV_ODID_DESC_TYPE = 0
)

//
type MAV_ODID_HEIGHT_REF int

const (
	// The height field is relative to the take-off location.
	MAV_ODID_HEIGHT_REF_OVER_TAKEOFF MAV_ODID_HEIGHT_REF = 0
	// The height field is relative to ground.
	MAV_ODID_HEIGHT_REF_OVER_GROUND MAV_ODID_HEIGHT_REF = 1
)

//
type MAV_ODID_HOR_ACC int

const (
	// The horizontal accuracy is unknown.
	MAV_ODID_HOR_ACC_UNKNOWN MAV_ODID_HOR_ACC = 0
	// The horizontal accuracy is smaller than 10 Nautical Miles. 18.52 km.
	MAV_ODID_HOR_ACC_10NM MAV_ODID_HOR_ACC = 1
	// The horizontal accuracy is smaller than 4 Nautical Miles. 7.408 km.
	MAV_ODID_HOR_ACC_4NM MAV_ODID_HOR_ACC = 2
	// The horizontal accuracy is smaller than 2 Nautical Miles. 3.704 km.
	MAV_ODID_HOR_ACC_2NM MAV_ODID_HOR_ACC = 3
	// The horizontal accuracy is smaller than 1 Nautical Miles. 1.852 km.
	MAV_ODID_HOR_ACC_1NM MAV_ODID_HOR_ACC = 4
	// The horizontal accuracy is smaller than 0.5 Nautical Miles. 926 m.
	MAV_ODID_HOR_ACC_0_5NM MAV_ODID_HOR_ACC = 5
	// The horizontal accuracy is smaller than 0.3 Nautical Miles. 555.6 m.
	MAV_ODID_HOR_ACC_0_3NM MAV_ODID_HOR_ACC = 6
	// The horizontal accuracy is smaller than 0.1 Nautical Miles. 185.2 m.
	MAV_ODID_HOR_ACC_0_1NM MAV_ODID_HOR_ACC = 7
	// The horizontal accuracy is smaller than 0.05 Nautical Miles. 92.6 m.
	MAV_ODID_HOR_ACC_0_05NM MAV_ODID_HOR_ACC = 8
	// The horizontal accuracy is smaller than 30 meter.
	MAV_ODID_HOR_ACC_30_METER MAV_ODID_HOR_ACC = 9
	// The horizontal accuracy is smaller than 10 meter.
	MAV_ODID_HOR_ACC_10_METER MAV_ODID_HOR_ACC = 10
	// The horizontal accuracy is smaller than 3 meter.
	MAV_ODID_HOR_ACC_3_METER MAV_ODID_HOR_ACC = 11
	// The horizontal accuracy is smaller than 1 meter.
	MAV_ODID_HOR_ACC_1_METER MAV_ODID_HOR_ACC = 12
)

//
type MAV_ODID_ID_TYPE int

const (
	// No type defined.
	MAV_ODID_ID_TYPE_NONE MAV_ODID_ID_TYPE = 0
	// Manufacturer Serial Number (ANSI/CTA-2063 format).
	MAV_ODID_ID_TYPE_SERIAL_NUMBER MAV_ODID_ID_TYPE = 1
	// CAA (Civil Aviation Authority) registered ID. Format: [ICAO Country Code].[CAA Assigned ID].
	MAV_ODID_ID_TYPE_CAA_REGISTRATION_ID MAV_ODID_ID_TYPE = 2
	// UTM (Unmanned Traffic Management) assigned UUID (RFC4122).
	MAV_ODID_ID_TYPE_UTM_ASSIGNED_UUID MAV_ODID_ID_TYPE = 3
)

//
type MAV_ODID_LOCATION_SRC int

const (
	// The location of the operator is the same as the take-off location.
	MAV_ODID_LOCATION_SRC_TAKEOFF MAV_ODID_LOCATION_SRC = 0
	// The location of the operator is based on live GNSS data.
	MAV_ODID_LOCATION_SRC_LIVE_GNSS MAV_ODID_LOCATION_SRC = 1
	// The location of the operator is a fixed location.
	MAV_ODID_LOCATION_SRC_FIXED MAV_ODID_LOCATION_SRC = 2
)

//
type MAV_ODID_OPERATOR_ID_TYPE int

const (
	// CAA (Civil Aviation Authority) registered operator ID.
	MAV_ODID_OPERATOR_ID_TYPE_CAA MAV_ODID_OPERATOR_ID_TYPE = 0
)

//
type MAV_ODID_SPEED_ACC int

const (
	// The speed accuracy is unknown.
	MAV_ODID_SPEED_ACC_UNKNOWN MAV_ODID_SPEED_ACC = 0
	// The speed accuracy is smaller than 10 meters per second.
	MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 1
	// The speed accuracy is smaller than 3 meters per second.
	MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 2
	// The speed accuracy is smaller than 1 meters per second.
	MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 3
	// The speed accuracy is smaller than 0.3 meters per second.
	MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 4
)

//
type MAV_ODID_STATUS int

const (
	// The status of the (UA) Unmanned Aircraft is undefined.
	MAV_ODID_STATUS_UNDECLARED MAV_ODID_STATUS = 0
	// The UA is on the ground.
	MAV_ODID_STATUS_GROUND MAV_ODID_STATUS = 1
	// The UA is in the air.
	MAV_ODID_STATUS_AIRBORNE MAV_ODID_STATUS = 2
)

//
type MAV_ODID_TIME_ACC int

const (
	// The timestamp accuracy is unknown.
	MAV_ODID_TIME_ACC_UNKNOWN MAV_ODID_TIME_ACC = 0
	// The timestamp accuracy is smaller than 0.1 second.
	MAV_ODID_TIME_ACC_0_1_SECOND MAV_ODID_TIME_ACC = 1
	// The timestamp accuracy is smaller than 0.2 second.
	MAV_ODID_TIME_ACC_0_2_SECOND MAV_ODID_TIME_ACC = 2
	// The timestamp accuracy is smaller than 0.3 second.
	MAV_ODID_TIME_ACC_0_3_SECOND MAV_ODID_TIME_ACC = 3
	// The timestamp accuracy is smaller than 0.4 second.
	MAV_ODID_TIME_ACC_0_4_SECOND MAV_ODID_TIME_ACC = 4
	// The timestamp accuracy is smaller than 0.5 second.
	MAV_ODID_TIME_ACC_0_5_SECOND MAV_ODID_TIME_ACC = 5
	// The timestamp accuracy is smaller than 0.6 second.
	MAV_ODID_TIME_ACC_0_6_SECOND MAV_ODID_TIME_ACC = 6
	// The timestamp accuracy is smaller than 0.7 second.
	MAV_ODID_TIME_ACC_0_7_SECOND MAV_ODID_TIME_ACC = 7
	// The timestamp accuracy is smaller than 0.8 second.
	MAV_ODID_TIME_ACC_0_8_SECOND MAV_ODID_TIME_ACC = 8
	// The timestamp accuracy is smaller than 0.9 second.
	MAV_ODID_TIME_ACC_0_9_SECOND MAV_ODID_TIME_ACC = 9
	// The timestamp accuracy is smaller than 1.0 second.
	MAV_ODID_TIME_ACC_1_0_SECOND MAV_ODID_TIME_ACC = 10
	// The timestamp accuracy is smaller than 1.1 second.
	MAV_ODID_TIME_ACC_1_1_SECOND MAV_ODID_TIME_ACC = 11
	// The timestamp accuracy is smaller than 1.2 second.
	MAV_ODID_TIME_ACC_1_2_SECOND MAV_ODID_TIME_ACC = 12
	// The timestamp accuracy is smaller than 1.3 second.
	MAV_ODID_TIME_ACC_1_3_SECOND MAV_ODID_TIME_ACC = 13
	// The timestamp accuracy is smaller than 1.4 second.
	MAV_ODID_TIME_ACC_1_4_SECOND MAV_ODID_TIME_ACC = 14
	// The timestamp accuracy is smaller than 1.5 second.
	MAV_ODID_TIME_ACC_1_5_SECOND MAV_ODID_TIME_ACC = 15
)

//
type MAV_ODID_UA_TYPE int

const (
	// No UA (Unmanned Aircraft) type defined.
	MAV_ODID_UA_TYPE_NONE MAV_ODID_UA_TYPE = 0
	// Aeroplane/Airplane. Fixed wing.
	MAV_ODID_UA_TYPE_AEROPLANE MAV_ODID_UA_TYPE = 1
	// Helicopter or multirotor.
	MAV_ODID_UA_TYPE_HELICOPTER_OR_MULTIROTOR MAV_ODID_UA_TYPE = 2
	// Gyroplane.
	MAV_ODID_UA_TYPE_GYROPLANE MAV_ODID_UA_TYPE = 3
	// VTOL (Vertical Take-Off and Landing). Fixed wing aircraft that can take off vertically.
	MAV_ODID_UA_TYPE_HYBRID_LIFT MAV_ODID_UA_TYPE = 4
	// Ornithopter.
	MAV_ODID_UA_TYPE_ORNITHOPTER MAV_ODID_UA_TYPE = 5
	// Glider.
	MAV_ODID_UA_TYPE_GLIDER MAV_ODID_UA_TYPE = 6
	// Kite.
	MAV_ODID_UA_TYPE_KITE MAV_ODID_UA_TYPE = 7
	// Free Balloon.
	MAV_ODID_UA_TYPE_FREE_BALLOON MAV_ODID_UA_TYPE = 8
	// Captive Balloon.
	MAV_ODID_UA_TYPE_CAPTIVE_BALLOON MAV_ODID_UA_TYPE = 9
	// Airship. E.g. a blimp.
	MAV_ODID_UA_TYPE_AIRSHIP MAV_ODID_UA_TYPE = 10
	// Free Fall/Parachute (unpowered).
	MAV_ODID_UA_TYPE_FREE_FALL_PARACHUTE MAV_ODID_UA_TYPE = 11
	// Rocket.
	MAV_ODID_UA_TYPE_ROCKET MAV_ODID_UA_TYPE = 12
	// Tethered powered aircraft.
	MAV_ODID_UA_TYPE_TETHERED_POWERED_AIRCRAFT MAV_ODID_UA_TYPE = 13
	// Ground Obstacle.
	MAV_ODID_UA_TYPE_GROUND_OBSTACLE MAV_ODID_UA_TYPE = 14
	// Other type of aircraft not listed earlier.
	MAV_ODID_UA_TYPE_OTHER MAV_ODID_UA_TYPE = 15
)

//
type MAV_ODID_VER_ACC int

const (
	// The vertical accuracy is unknown.
	MAV_ODID_VER_ACC_UNKNOWN MAV_ODID_VER_ACC = 0
	// The vertical accuracy is smaller than 150 meter.
	MAV_ODID_VER_ACC_150_METER MAV_ODID_VER_ACC = 1
	// The vertical accuracy is smaller than 45 meter.
	MAV_ODID_VER_ACC_45_METER MAV_ODID_VER_ACC = 2
	// The vertical accuracy is smaller than 25 meter.
	MAV_ODID_VER_ACC_25_METER MAV_ODID_VER_ACC = 3
	// The vertical accuracy is smaller than 10 meter.
	MAV_ODID_VER_ACC_10_METER MAV_ODID_VER_ACC = 4
	// The vertical accuracy is smaller than 3 meter.
	MAV_ODID_VER_ACC_3_METER MAV_ODID_VER_ACC = 5
	// The vertical accuracy is smaller than 1 meter.
	MAV_ODID_VER_ACC_1_METER MAV_ODID_VER_ACC = 6
)

// Specifies the datatype of a MAVLink extended parameter.
type MAV_PARAM_EXT_TYPE int

const (
	// 8-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT8 MAV_PARAM_EXT_TYPE = 1
	// 8-bit signed integer
	MAV_PARAM_EXT_TYPE_INT8 MAV_PARAM_EXT_TYPE = 2
	// 16-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT16 MAV_PARAM_EXT_TYPE = 3
	// 16-bit signed integer
	MAV_PARAM_EXT_TYPE_INT16 MAV_PARAM_EXT_TYPE = 4
	// 32-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT32 MAV_PARAM_EXT_TYPE = 5
	// 32-bit signed integer
	MAV_PARAM_EXT_TYPE_INT32 MAV_PARAM_EXT_TYPE = 6
	// 64-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT64 MAV_PARAM_EXT_TYPE = 7
	// 64-bit signed integer
	MAV_PARAM_EXT_TYPE_INT64 MAV_PARAM_EXT_TYPE = 8
	// 32-bit floating-point
	MAV_PARAM_EXT_TYPE_REAL32 MAV_PARAM_EXT_TYPE = 9
	// 64-bit floating-point
	MAV_PARAM_EXT_TYPE_REAL64 MAV_PARAM_EXT_TYPE = 10
	// Custom Type
	MAV_PARAM_EXT_TYPE_CUSTOM MAV_PARAM_EXT_TYPE = 11
)

// Specifies the datatype of a MAVLink parameter.
type MAV_PARAM_TYPE int

const (
	// 8-bit unsigned integer
	MAV_PARAM_TYPE_UINT8 MAV_PARAM_TYPE = 1
	// 8-bit signed integer
	MAV_PARAM_TYPE_INT8 MAV_PARAM_TYPE = 2
	// 16-bit unsigned integer
	MAV_PARAM_TYPE_UINT16 MAV_PARAM_TYPE = 3
	// 16-bit signed integer
	MAV_PARAM_TYPE_INT16 MAV_PARAM_TYPE = 4
	// 32-bit unsigned integer
	MAV_PARAM_TYPE_UINT32 MAV_PARAM_TYPE = 5
	// 32-bit signed integer
	MAV_PARAM_TYPE_INT32 MAV_PARAM_TYPE = 6
	// 64-bit unsigned integer
	MAV_PARAM_TYPE_UINT64 MAV_PARAM_TYPE = 7
	// 64-bit signed integer
	MAV_PARAM_TYPE_INT64 MAV_PARAM_TYPE = 8
	// 32-bit floating-point
	MAV_PARAM_TYPE_REAL32 MAV_PARAM_TYPE = 9
	// 64-bit floating-point
	MAV_PARAM_TYPE_REAL64 MAV_PARAM_TYPE = 10
)

// Power supply status flags (bitmask)
type MAV_POWER_STATUS int

const (
	// main brick power supply valid
	MAV_POWER_STATUS_BRICK_VALID MAV_POWER_STATUS = 1
	// main servo power supply valid for FMU
	MAV_POWER_STATUS_SERVO_VALID MAV_POWER_STATUS = 2
	// USB power is connected
	MAV_POWER_STATUS_USB_CONNECTED MAV_POWER_STATUS = 4
	// peripheral supply is in over-current state
	MAV_POWER_STATUS_PERIPH_OVERCURRENT MAV_POWER_STATUS = 8
	// hi-power peripheral supply is in over-current state
	MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT MAV_POWER_STATUS = 16
	// Power status has changed since boot
	MAV_POWER_STATUS_CHANGED MAV_POWER_STATUS = 32
)

// Bitmask of (optional) autopilot capabilities (64 bit). If a bit is set, the autopilot supports this capability.
type MAV_PROTOCOL_CAPABILITY int

const (
	// Autopilot supports MISSION float message type.
	MAV_PROTOCOL_CAPABILITY_MISSION_FLOAT MAV_PROTOCOL_CAPABILITY = 1
	// Autopilot supports the new param float message type.
	MAV_PROTOCOL_CAPABILITY_PARAM_FLOAT MAV_PROTOCOL_CAPABILITY = 2
	// Autopilot supports MISSION_INT scaled integer message type.
	MAV_PROTOCOL_CAPABILITY_MISSION_INT MAV_PROTOCOL_CAPABILITY = 4
	// Autopilot supports COMMAND_INT scaled integer message type.
	MAV_PROTOCOL_CAPABILITY_COMMAND_INT MAV_PROTOCOL_CAPABILITY = 8
	// Autopilot supports the new param union message type.
	MAV_PROTOCOL_CAPABILITY_PARAM_UNION MAV_PROTOCOL_CAPABILITY = 16
	// Autopilot supports the new FILE_TRANSFER_PROTOCOL message type.
	MAV_PROTOCOL_CAPABILITY_FTP MAV_PROTOCOL_CAPABILITY = 32
	// Autopilot supports commanding attitude offboard.
	MAV_PROTOCOL_CAPABILITY_SET_ATTITUDE_TARGET MAV_PROTOCOL_CAPABILITY = 64
	// Autopilot supports commanding position and velocity targets in local NED frame.
	MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_LOCAL_NED MAV_PROTOCOL_CAPABILITY = 128
	// Autopilot supports commanding position and velocity targets in global scaled integers.
	MAV_PROTOCOL_CAPABILITY_SET_POSITION_TARGET_GLOBAL_INT MAV_PROTOCOL_CAPABILITY = 256
	// Autopilot supports terrain protocol / data handling.
	MAV_PROTOCOL_CAPABILITY_TERRAIN MAV_PROTOCOL_CAPABILITY = 512
	// Autopilot supports direct actuator control.
	MAV_PROTOCOL_CAPABILITY_SET_ACTUATOR_TARGET MAV_PROTOCOL_CAPABILITY = 1024
	// Autopilot supports the flight termination command.
	MAV_PROTOCOL_CAPABILITY_FLIGHT_TERMINATION MAV_PROTOCOL_CAPABILITY = 2048
	// Autopilot supports onboard compass calibration.
	MAV_PROTOCOL_CAPABILITY_COMPASS_CALIBRATION MAV_PROTOCOL_CAPABILITY = 4096
	// Autopilot supports MAVLink version 2.
	MAV_PROTOCOL_CAPABILITY_MAVLINK2 MAV_PROTOCOL_CAPABILITY = 8192
	// Autopilot supports mission fence protocol.
	MAV_PROTOCOL_CAPABILITY_MISSION_FENCE MAV_PROTOCOL_CAPABILITY = 16384
	// Autopilot supports mission rally point protocol.
	MAV_PROTOCOL_CAPABILITY_MISSION_RALLY MAV_PROTOCOL_CAPABILITY = 32768
	// Autopilot supports the flight information protocol.
	MAV_PROTOCOL_CAPABILITY_FLIGHT_INFORMATION MAV_PROTOCOL_CAPABILITY = 65536
)

// Result from a MAVLink command (MAV_CMD)
type MAV_RESULT int

const (
	// Command is valid (is supported and has valid parameters), and was executed.
	MAV_RESULT_ACCEPTED MAV_RESULT = 0
	// Command is valid, but cannot be executed at this time. This is used to indicate a problem that should be fixed just by waiting (e.g. a state machine is busy, can't arm because have not got GPS lock, etc.). Retrying later should work.
	MAV_RESULT_TEMPORARILY_REJECTED MAV_RESULT = 1
	// Command is invalid (is supported but has invalid parameters). Retrying same command and parameters will not work.
	MAV_RESULT_DENIED MAV_RESULT = 2
	// Command is not supported (unknown).
	MAV_RESULT_UNSUPPORTED MAV_RESULT = 3
	// Command is valid, but execution has failed. This is used to indicate any non-temporary or unexpected problem, i.e. any problem that must be fixed before the command can succeed/be retried. For example, attempting to write a file when out of memory, attempting to arm when sensors are not calibrated, etc.
	MAV_RESULT_FAILED MAV_RESULT = 4
	// Command is valid and is being executed. This will be followed by further progress updates, i.e. the component may send further COMMAND_ACK messages with result MAV_RESULT_IN_PROGRESS (at a rate decided by the implementation), and must terminate by sending a COMMAND_ACK message with final result of the operation. The COMMAND_ACK.progress field can be used to indicate the progress of the operation. There is no need for the sender to retry the command, but if done during execution, the component will return MAV_RESULT_IN_PROGRESS with an updated progress.
	MAV_RESULT_IN_PROGRESS MAV_RESULT = 5
)

// The ROI (region of interest) for the vehicle. This can be                be used by the vehicle for camera/vehicle attitude alignment (see                MAV_CMD_NAV_ROI).
type MAV_ROI int

const (
	// No region of interest.
	MAV_ROI_NONE MAV_ROI = 0
	// Point toward next waypoint, with optional pitch/roll/yaw offset.
	MAV_ROI_WPNEXT MAV_ROI = 1
	// Point toward given waypoint.
	MAV_ROI_WPINDEX MAV_ROI = 2
	// Point toward fixed location.
	MAV_ROI_LOCATION MAV_ROI = 3
	// Point toward of given id.
	MAV_ROI_TARGET MAV_ROI = 4
)

// Enumeration of sensor orientation, according to its rotations
type MAV_SENSOR_ORIENTATION int

const (
	// Roll: 0, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_NONE MAV_SENSOR_ORIENTATION = 0
	// Roll: 0, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_YAW_45 MAV_SENSOR_ORIENTATION = 1
	// Roll: 0, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_YAW_90 MAV_SENSOR_ORIENTATION = 2
	// Roll: 0, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_YAW_135 MAV_SENSOR_ORIENTATION = 3
	// Roll: 0, Pitch: 0, Yaw: 180
	MAV_SENSOR_ROTATION_YAW_180 MAV_SENSOR_ORIENTATION = 4
	// Roll: 0, Pitch: 0, Yaw: 225
	MAV_SENSOR_ROTATION_YAW_225 MAV_SENSOR_ORIENTATION = 5
	// Roll: 0, Pitch: 0, Yaw: 270
	MAV_SENSOR_ROTATION_YAW_270 MAV_SENSOR_ORIENTATION = 6
	// Roll: 0, Pitch: 0, Yaw: 315
	MAV_SENSOR_ROTATION_YAW_315 MAV_SENSOR_ORIENTATION = 7
	// Roll: 180, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_180 MAV_SENSOR_ORIENTATION = 8
	// Roll: 180, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_ROLL_180_YAW_45 MAV_SENSOR_ORIENTATION = 9
	// Roll: 180, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_180_YAW_90 MAV_SENSOR_ORIENTATION = 10
	// Roll: 180, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_ROLL_180_YAW_135 MAV_SENSOR_ORIENTATION = 11
	// Roll: 0, Pitch: 180, Yaw: 0
	MAV_SENSOR_ROTATION_PITCH_180 MAV_SENSOR_ORIENTATION = 12
	// Roll: 180, Pitch: 0, Yaw: 225
	MAV_SENSOR_ROTATION_ROLL_180_YAW_225 MAV_SENSOR_ORIENTATION = 13
	// Roll: 180, Pitch: 0, Yaw: 270
	MAV_SENSOR_ROTATION_ROLL_180_YAW_270 MAV_SENSOR_ORIENTATION = 14
	// Roll: 180, Pitch: 0, Yaw: 315
	MAV_SENSOR_ROTATION_ROLL_180_YAW_315 MAV_SENSOR_ORIENTATION = 15
	// Roll: 90, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90 MAV_SENSOR_ORIENTATION = 16
	// Roll: 90, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_ROLL_90_YAW_45 MAV_SENSOR_ORIENTATION = 17
	// Roll: 90, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_90_YAW_90 MAV_SENSOR_ORIENTATION = 18
	// Roll: 90, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_ROLL_90_YAW_135 MAV_SENSOR_ORIENTATION = 19
	// Roll: 270, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270 MAV_SENSOR_ORIENTATION = 20
	// Roll: 270, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_ROLL_270_YAW_45 MAV_SENSOR_ORIENTATION = 21
	// Roll: 270, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_270_YAW_90 MAV_SENSOR_ORIENTATION = 22
	// Roll: 270, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_ROLL_270_YAW_135 MAV_SENSOR_ORIENTATION = 23
	// Roll: 0, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_PITCH_90 MAV_SENSOR_ORIENTATION = 24
	// Roll: 0, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_PITCH_270 MAV_SENSOR_ORIENTATION = 25
	// Roll: 0, Pitch: 180, Yaw: 90
	MAV_SENSOR_ROTATION_PITCH_180_YAW_90 MAV_SENSOR_ORIENTATION = 26
	// Roll: 0, Pitch: 180, Yaw: 270
	MAV_SENSOR_ROTATION_PITCH_180_YAW_270 MAV_SENSOR_ORIENTATION = 27
	// Roll: 90, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_90 MAV_SENSOR_ORIENTATION = 28
	// Roll: 180, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_180_PITCH_90 MAV_SENSOR_ORIENTATION = 29
	// Roll: 270, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270_PITCH_90 MAV_SENSOR_ORIENTATION = 30
	// Roll: 90, Pitch: 180, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_180 MAV_SENSOR_ORIENTATION = 31
	// Roll: 270, Pitch: 180, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270_PITCH_180 MAV_SENSOR_ORIENTATION = 32
	// Roll: 90, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_270 MAV_SENSOR_ORIENTATION = 33
	// Roll: 180, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_180_PITCH_270 MAV_SENSOR_ORIENTATION = 34
	// Roll: 270, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270_PITCH_270 MAV_SENSOR_ORIENTATION = 35
	// Roll: 90, Pitch: 180, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_180_YAW_90 MAV_SENSOR_ORIENTATION = 36
	// Roll: 90, Pitch: 0, Yaw: 270
	MAV_SENSOR_ROTATION_ROLL_90_YAW_270 MAV_SENSOR_ORIENTATION = 37
	// Roll: 90, Pitch: 68, Yaw: 293
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_68_YAW_293 MAV_SENSOR_ORIENTATION = 38
	// Pitch: 315
	MAV_SENSOR_ROTATION_PITCH_315 MAV_SENSOR_ORIENTATION = 39
	// Roll: 90, Pitch: 315
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_315 MAV_SENSOR_ORIENTATION = 40
	// Custom orientation
	MAV_SENSOR_ROTATION_CUSTOM MAV_SENSOR_ORIENTATION = 100
)

// Indicates the severity level, generally used for status messages to indicate their relative urgency. Based on RFC-5424 using expanded definitions at: http://www.kiwisyslog.com/kb/info:-syslog-message-levels/.
type MAV_SEVERITY int

const (
	// System is unusable. This is a "panic" condition.
	MAV_SEVERITY_EMERGENCY MAV_SEVERITY = 0
	// Action should be taken immediately. Indicates error in non-critical systems.
	MAV_SEVERITY_ALERT MAV_SEVERITY = 1
	// Action must be taken immediately. Indicates failure in a primary system.
	MAV_SEVERITY_CRITICAL MAV_SEVERITY = 2
	// Indicates an error in secondary/redundant systems.
	MAV_SEVERITY_ERROR MAV_SEVERITY = 3
	// Indicates about a possible future error if this is not resolved within a given timeframe. Example would be a low battery warning.
	MAV_SEVERITY_WARNING MAV_SEVERITY = 4
	// An unusual event has occurred, though not an error condition. This should be investigated for the root cause.
	MAV_SEVERITY_NOTICE MAV_SEVERITY = 5
	// Normal operational messages. Useful for logging. No action is required for these messages.
	MAV_SEVERITY_INFO MAV_SEVERITY = 6
	// Useful non-operational messages that can assist in debugging. These should not occur during normal operation.
	MAV_SEVERITY_DEBUG MAV_SEVERITY = 7
)

// Smart battery supply status/fault flags (bitmask) for health indication.
type MAV_SMART_BATTERY_FAULT int

const (
	// Battery has deep discharged.
	MAV_SMART_BATTERY_FAULT_DEEP_DISCHARGE MAV_SMART_BATTERY_FAULT = 1
	// Voltage spikes.
	MAV_SMART_BATTERY_FAULT_SPIKES MAV_SMART_BATTERY_FAULT = 2
	// Single cell has failed.
	MAV_SMART_BATTERY_FAULT_SINGLE_CELL_FAIL MAV_SMART_BATTERY_FAULT = 4
	// Over-current fault.
	MAV_SMART_BATTERY_FAULT_OVER_CURRENT MAV_SMART_BATTERY_FAULT = 8
	// Over-temperature fault.
	MAV_SMART_BATTERY_FAULT_OVER_TEMPERATURE MAV_SMART_BATTERY_FAULT = 16
	// Under-temperature fault.
	MAV_SMART_BATTERY_FAULT_UNDER_TEMPERATURE MAV_SMART_BATTERY_FAULT = 32
)

//
type MAV_STATE int

const (
	// Uninitialized system, state is unknown.
	MAV_STATE_UNINIT MAV_STATE = 0
	// System is booting up.
	MAV_STATE_BOOT MAV_STATE = 1
	// System is calibrating and not flight-ready.
	MAV_STATE_CALIBRATING MAV_STATE = 2
	// System is grounded and on standby. It can be launched any time.
	MAV_STATE_STANDBY MAV_STATE = 3
	// System is active and might be already airborne. Motors are engaged.
	MAV_STATE_ACTIVE MAV_STATE = 4
	// System is in a non-normal flight mode. It can however still navigate.
	MAV_STATE_CRITICAL MAV_STATE = 5
	// System is in a non-normal flight mode. It lost control over parts or over the whole airframe. It is in mayday and going down.
	MAV_STATE_EMERGENCY MAV_STATE = 6
	// System just initialized its power-down sequence, will shut down now.
	MAV_STATE_POWEROFF MAV_STATE = 7
	// System is terminating itself.
	MAV_STATE_FLIGHT_TERMINATION MAV_STATE = 8
)

// These encode the sensors whose status is sent as part of the SYS_STATUS message.
type MAV_SYS_STATUS_SENSOR int

const (
	// 0x01 3D gyro
	MAV_SYS_STATUS_SENSOR_3D_GYRO MAV_SYS_STATUS_SENSOR = 1
	// 0x02 3D accelerometer
	MAV_SYS_STATUS_SENSOR_3D_ACCEL MAV_SYS_STATUS_SENSOR = 2
	// 0x04 3D magnetometer
	MAV_SYS_STATUS_SENSOR_3D_MAG MAV_SYS_STATUS_SENSOR = 4
	// 0x08 absolute pressure
	MAV_SYS_STATUS_SENSOR_ABSOLUTE_PRESSURE MAV_SYS_STATUS_SENSOR = 8
	// 0x10 differential pressure
	MAV_SYS_STATUS_SENSOR_DIFFERENTIAL_PRESSURE MAV_SYS_STATUS_SENSOR = 16
	// 0x20 GPS
	MAV_SYS_STATUS_SENSOR_GPS MAV_SYS_STATUS_SENSOR = 32
	// 0x40 optical flow
	MAV_SYS_STATUS_SENSOR_OPTICAL_FLOW MAV_SYS_STATUS_SENSOR = 64
	// 0x80 computer vision position
	MAV_SYS_STATUS_SENSOR_VISION_POSITION MAV_SYS_STATUS_SENSOR = 128
	// 0x100 laser based position
	MAV_SYS_STATUS_SENSOR_LASER_POSITION MAV_SYS_STATUS_SENSOR = 256
	// 0x200 external ground truth (Vicon or Leica)
	MAV_SYS_STATUS_SENSOR_EXTERNAL_GROUND_TRUTH MAV_SYS_STATUS_SENSOR = 512
	// 0x400 3D angular rate control
	MAV_SYS_STATUS_SENSOR_ANGULAR_RATE_CONTROL MAV_SYS_STATUS_SENSOR = 1024
	// 0x800 attitude stabilization
	MAV_SYS_STATUS_SENSOR_ATTITUDE_STABILIZATION MAV_SYS_STATUS_SENSOR = 2048
	// 0x1000 yaw position
	MAV_SYS_STATUS_SENSOR_YAW_POSITION MAV_SYS_STATUS_SENSOR = 4096
	// 0x2000 z/altitude control
	MAV_SYS_STATUS_SENSOR_Z_ALTITUDE_CONTROL MAV_SYS_STATUS_SENSOR = 8192
	// 0x4000 x/y position control
	MAV_SYS_STATUS_SENSOR_XY_POSITION_CONTROL MAV_SYS_STATUS_SENSOR = 16384
	// 0x8000 motor outputs / control
	MAV_SYS_STATUS_SENSOR_MOTOR_OUTPUTS MAV_SYS_STATUS_SENSOR = 32768
	// 0x10000 rc receiver
	MAV_SYS_STATUS_SENSOR_RC_RECEIVER MAV_SYS_STATUS_SENSOR = 65536
	// 0x20000 2nd 3D gyro
	MAV_SYS_STATUS_SENSOR_3D_GYRO2 MAV_SYS_STATUS_SENSOR = 131072
	// 0x40000 2nd 3D accelerometer
	MAV_SYS_STATUS_SENSOR_3D_ACCEL2 MAV_SYS_STATUS_SENSOR = 262144
	// 0x80000 2nd 3D magnetometer
	MAV_SYS_STATUS_SENSOR_3D_MAG2 MAV_SYS_STATUS_SENSOR = 524288
	// 0x100000 geofence
	MAV_SYS_STATUS_GEOFENCE MAV_SYS_STATUS_SENSOR = 1048576
	// 0x200000 AHRS subsystem health
	MAV_SYS_STATUS_AHRS MAV_SYS_STATUS_SENSOR = 2097152
	// 0x400000 Terrain subsystem health
	MAV_SYS_STATUS_TERRAIN MAV_SYS_STATUS_SENSOR = 4194304
	// 0x800000 Motors are reversed
	MAV_SYS_STATUS_REVERSE_MOTOR MAV_SYS_STATUS_SENSOR = 8388608
	// 0x1000000 Logging
	MAV_SYS_STATUS_LOGGING MAV_SYS_STATUS_SENSOR = 16777216
	// 0x2000000 Battery
	MAV_SYS_STATUS_SENSOR_BATTERY MAV_SYS_STATUS_SENSOR = 33554432
	// 0x4000000 Proximity
	MAV_SYS_STATUS_SENSOR_PROXIMITY MAV_SYS_STATUS_SENSOR = 67108864
	// 0x8000000 Satellite Communication
	MAV_SYS_STATUS_SENSOR_SATCOM MAV_SYS_STATUS_SENSOR = 134217728
	// 0x10000000 pre-arm check status. Always healthy when armed
	MAV_SYS_STATUS_PREARM_CHECK MAV_SYS_STATUS_SENSOR = 268435456
	// 0x20000000 Avoidance/collision prevention
	MAV_SYS_STATUS_OBSTACLE_AVOIDANCE MAV_SYS_STATUS_SENSOR = 536870912
)

//
type MAV_TUNNEL_PAYLOAD_TYPE int

const (
	// Encoding of payload unknown.
	MAV_TUNNEL_PAYLOAD_TYPE_UNKNOWN MAV_TUNNEL_PAYLOAD_TYPE = 0
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED0 MAV_TUNNEL_PAYLOAD_TYPE = 200
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED1 MAV_TUNNEL_PAYLOAD_TYPE = 201
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED2 MAV_TUNNEL_PAYLOAD_TYPE = 202
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED3 MAV_TUNNEL_PAYLOAD_TYPE = 203
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED4 MAV_TUNNEL_PAYLOAD_TYPE = 204
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED5 MAV_TUNNEL_PAYLOAD_TYPE = 205
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED6 MAV_TUNNEL_PAYLOAD_TYPE = 206
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED7 MAV_TUNNEL_PAYLOAD_TYPE = 207
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED8 MAV_TUNNEL_PAYLOAD_TYPE = 208
	// Registered for STorM32 gimbal controller.
	MAV_TUNNEL_PAYLOAD_TYPE_STORM32_RESERVED9 MAV_TUNNEL_PAYLOAD_TYPE = 209
)

// MAVLINK component type reported in HEARTBEAT message. Flight controllers must report the type of the vehicle on which they are mounted (e.g. MAV_TYPE_OCTOROTOR). All other components must report a value appropriate for their type (e.g. a camera must use MAV_TYPE_CAMERA).
type MAV_TYPE int

const (
	// Generic micro air vehicle
	MAV_TYPE_GENERIC MAV_TYPE = 0
	// Fixed wing aircraft.
	MAV_TYPE_FIXED_WING MAV_TYPE = 1
	// Quadrotor
	MAV_TYPE_QUADROTOR MAV_TYPE = 2
	// Coaxial helicopter
	MAV_TYPE_COAXIAL MAV_TYPE = 3
	// Normal helicopter with tail rotor.
	MAV_TYPE_HELICOPTER MAV_TYPE = 4
	// Ground installation
	MAV_TYPE_ANTENNA_TRACKER MAV_TYPE = 5
	// Operator control unit / ground control station
	MAV_TYPE_GCS MAV_TYPE = 6
	// Airship, controlled
	MAV_TYPE_AIRSHIP MAV_TYPE = 7
	// Free balloon, uncontrolled
	MAV_TYPE_FREE_BALLOON MAV_TYPE = 8
	// Rocket
	MAV_TYPE_ROCKET MAV_TYPE = 9
	// Ground rover
	MAV_TYPE_GROUND_ROVER MAV_TYPE = 10
	// Surface vessel, boat, ship
	MAV_TYPE_SURFACE_BOAT MAV_TYPE = 11
	// Submarine
	MAV_TYPE_SUBMARINE MAV_TYPE = 12
	// Hexarotor
	MAV_TYPE_HEXAROTOR MAV_TYPE = 13
	// Octorotor
	MAV_TYPE_OCTOROTOR MAV_TYPE = 14
	// Tricopter
	MAV_TYPE_TRICOPTER MAV_TYPE = 15
	// Flapping wing
	MAV_TYPE_FLAPPING_WING MAV_TYPE = 16
	// Kite
	MAV_TYPE_KITE MAV_TYPE = 17
	// Onboard companion controller
	MAV_TYPE_ONBOARD_CONTROLLER MAV_TYPE = 18
	// Two-rotor VTOL using control surfaces in vertical operation in addition. Tailsitter.
	MAV_TYPE_VTOL_DUOROTOR MAV_TYPE = 19
	// Quad-rotor VTOL using a V-shaped quad config in vertical operation. Tailsitter.
	MAV_TYPE_VTOL_QUADROTOR MAV_TYPE = 20
	// Tiltrotor VTOL
	MAV_TYPE_VTOL_TILTROTOR MAV_TYPE = 21
	// VTOL reserved 2
	MAV_TYPE_VTOL_RESERVED2 MAV_TYPE = 22
	// VTOL reserved 3
	MAV_TYPE_VTOL_RESERVED3 MAV_TYPE = 23
	// VTOL reserved 4
	MAV_TYPE_VTOL_RESERVED4 MAV_TYPE = 24
	// VTOL reserved 5
	MAV_TYPE_VTOL_RESERVED5 MAV_TYPE = 25
	// Gimbal
	MAV_TYPE_GIMBAL MAV_TYPE = 26
	// ADSB system
	MAV_TYPE_ADSB MAV_TYPE = 27
	// Steerable, nonrigid airfoil
	MAV_TYPE_PARAFOIL MAV_TYPE = 28
	// Dodecarotor
	MAV_TYPE_DODECAROTOR MAV_TYPE = 29
	// Camera
	MAV_TYPE_CAMERA MAV_TYPE = 30
	// Charging station
	MAV_TYPE_CHARGING_STATION MAV_TYPE = 31
	// FLARM collision avoidance system
	MAV_TYPE_FLARM MAV_TYPE = 32
	// Servo
	MAV_TYPE_SERVO MAV_TYPE = 33
)

// Enumeration of VTOL states
type MAV_VTOL_STATE int

const (
	// MAV is not configured as VTOL
	MAV_VTOL_STATE_UNDEFINED MAV_VTOL_STATE = 0
	// VTOL is in transition from multicopter to fixed-wing
	MAV_VTOL_STATE_TRANSITION_TO_FW MAV_VTOL_STATE = 1
	// VTOL is in transition from fixed-wing to multicopter
	MAV_VTOL_STATE_TRANSITION_TO_MC MAV_VTOL_STATE = 2
	// VTOL is in multicopter state
	MAV_VTOL_STATE_MC MAV_VTOL_STATE = 3
	// VTOL is in fixed-wing state
	MAV_VTOL_STATE_FW MAV_VTOL_STATE = 4
)

//
type MOTOR_TEST_ORDER int

const (
	// default autopilot motor test method
	MOTOR_TEST_ORDER_DEFAULT MOTOR_TEST_ORDER = 0
	// motor numbers are specified as their index in a predefined vehicle-specific sequence
	MOTOR_TEST_ORDER_SEQUENCE MOTOR_TEST_ORDER = 1
	// motor numbers are specified as the output as labeled on the board
	MOTOR_TEST_ORDER_BOARD MOTOR_TEST_ORDER = 2
)

//
type MOTOR_TEST_THROTTLE_TYPE int

const (
	// throttle as a percentage from 0 ~ 100
	MOTOR_TEST_THROTTLE_PERCENT MOTOR_TEST_THROTTLE_TYPE = 0
	// throttle as an absolute PWM value (normally in range of 1000~2000)
	MOTOR_TEST_THROTTLE_PWM MOTOR_TEST_THROTTLE_TYPE = 1
	// throttle pass-through from pilot's transmitter
	MOTOR_TEST_THROTTLE_PILOT MOTOR_TEST_THROTTLE_TYPE = 2
	// per-motor compass calibration test
	MOTOR_TEST_COMPASS_CAL MOTOR_TEST_THROTTLE_TYPE = 3
)

// Yaw behaviour during orbit flight.
type ORBIT_YAW_BEHAVIOUR int

const (
	// Vehicle front points to the center (default).
	ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER ORBIT_YAW_BEHAVIOUR = 0
	// Vehicle front holds heading when message received.
	ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING ORBIT_YAW_BEHAVIOUR = 1
	// Yaw uncontrolled.
	ORBIT_YAW_BEHAVIOUR_UNCONTROLLED ORBIT_YAW_BEHAVIOUR = 2
	// Vehicle front follows flight path (tangential to circle).
	ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE ORBIT_YAW_BEHAVIOUR = 3
	// Yaw controlled by RC input.
	ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED ORBIT_YAW_BEHAVIOUR = 4
)

//
type PARACHUTE_ACTION int

const (
	// Disable parachute release.
	PARACHUTE_DISABLE PARACHUTE_ACTION = 0
	// Enable parachute release.
	PARACHUTE_ENABLE PARACHUTE_ACTION = 1
	// Release parachute.
	PARACHUTE_RELEASE PARACHUTE_ACTION = 2
)

// Result from a PARAM_EXT_SET message.
type PARAM_ACK int

const (
	// Parameter value ACCEPTED and SET
	PARAM_ACK_ACCEPTED PARAM_ACK = 0
	// Parameter value UNKNOWN/UNSUPPORTED
	PARAM_ACK_VALUE_UNSUPPORTED PARAM_ACK = 1
	// Parameter failed to set
	PARAM_ACK_FAILED PARAM_ACK = 2
	// Parameter value received but not yet validated or set. A subsequent PARAM_EXT_ACK will follow once operation is completed with the actual result. These are for parameters that may take longer to set. Instead of waiting for an ACK and potentially timing out, you will immediately receive this response to let you know it was received.
	PARAM_ACK_IN_PROGRESS PARAM_ACK = 3
)

// Bitmap to indicate which dimensions should be ignored by the vehicle: a value of 0b0000000000000000 or 0b0000001000000000 indicates that none of the setpoint dimensions should be ignored. If bit 9 is set the floats afx afy afz should be interpreted as force instead of acceleration.
type POSITION_TARGET_TYPEMASK int

const (
	// Ignore position x
	POSITION_TARGET_TYPEMASK_X_IGNORE POSITION_TARGET_TYPEMASK = 1
	// Ignore position y
	POSITION_TARGET_TYPEMASK_Y_IGNORE POSITION_TARGET_TYPEMASK = 2
	// Ignore position z
	POSITION_TARGET_TYPEMASK_Z_IGNORE POSITION_TARGET_TYPEMASK = 4
	// Ignore velocity x
	POSITION_TARGET_TYPEMASK_VX_IGNORE POSITION_TARGET_TYPEMASK = 8
	// Ignore velocity y
	POSITION_TARGET_TYPEMASK_VY_IGNORE POSITION_TARGET_TYPEMASK = 16
	// Ignore velocity z
	POSITION_TARGET_TYPEMASK_VZ_IGNORE POSITION_TARGET_TYPEMASK = 32
	// Ignore acceleration x
	POSITION_TARGET_TYPEMASK_AX_IGNORE POSITION_TARGET_TYPEMASK = 64
	// Ignore acceleration y
	POSITION_TARGET_TYPEMASK_AY_IGNORE POSITION_TARGET_TYPEMASK = 128
	// Ignore acceleration z
	POSITION_TARGET_TYPEMASK_AZ_IGNORE POSITION_TARGET_TYPEMASK = 256
	// Use force instead of acceleration
	POSITION_TARGET_TYPEMASK_FORCE_SET POSITION_TARGET_TYPEMASK = 512
	// Ignore yaw
	POSITION_TARGET_TYPEMASK_YAW_IGNORE POSITION_TARGET_TYPEMASK = 1024
	// Ignore yaw rate
	POSITION_TARGET_TYPEMASK_YAW_RATE_IGNORE POSITION_TARGET_TYPEMASK = 2048
)

// Precision land modes (used in MAV_CMD_NAV_LAND).
type PRECISION_LAND_MODE int

const (
	// Normal (non-precision) landing.
	PRECISION_LAND_MODE_DISABLED PRECISION_LAND_MODE = 0
	// Use precision landing if beacon detected when land command accepted, otherwise land normally.
	PRECISION_LAND_MODE_OPPORTUNISTIC PRECISION_LAND_MODE = 1
	// Use precision landing, searching for beacon if not found when land command accepted (land normally if beacon cannot be found).
	PRECISION_LAND_MODE_REQUIRED PRECISION_LAND_MODE = 2
)

// RC type
type RC_TYPE int

const (
	// Spektrum DSM2
	RC_TYPE_SPEKTRUM_DSM2 RC_TYPE = 0
	// Spektrum DSMX
	RC_TYPE_SPEKTRUM_DSMX RC_TYPE = 1
)

// RTK GPS baseline coordinate system, used for RTK corrections
type RTK_BASELINE_COORDINATE_SYSTEM int

const (
	// Earth-centered, Earth-fixed
	RTK_BASELINE_COORDINATE_SYSTEM_ECEF RTK_BASELINE_COORDINATE_SYSTEM = 0
	// RTK basestation centered, north, east, down
	RTK_BASELINE_COORDINATE_SYSTEM_NED RTK_BASELINE_COORDINATE_SYSTEM = 1
)

// SERIAL_CONTROL device types
type SERIAL_CONTROL_DEV int

const (
	// First telemetry port
	SERIAL_CONTROL_DEV_TELEM1 SERIAL_CONTROL_DEV = 0
	// Second telemetry port
	SERIAL_CONTROL_DEV_TELEM2 SERIAL_CONTROL_DEV = 1
	// First GPS port
	SERIAL_CONTROL_DEV_GPS1 SERIAL_CONTROL_DEV = 2
	// Second GPS port
	SERIAL_CONTROL_DEV_GPS2 SERIAL_CONTROL_DEV = 3
	// system shell
	SERIAL_CONTROL_DEV_SHELL SERIAL_CONTROL_DEV = 10
	// SERIAL0
	SERIAL_CONTROL_SERIAL0 SERIAL_CONTROL_DEV = 100
	// SERIAL1
	SERIAL_CONTROL_SERIAL1 SERIAL_CONTROL_DEV = 101
	// SERIAL2
	SERIAL_CONTROL_SERIAL2 SERIAL_CONTROL_DEV = 102
	// SERIAL3
	SERIAL_CONTROL_SERIAL3 SERIAL_CONTROL_DEV = 103
	// SERIAL4
	SERIAL_CONTROL_SERIAL4 SERIAL_CONTROL_DEV = 104
	// SERIAL5
	SERIAL_CONTROL_SERIAL5 SERIAL_CONTROL_DEV = 105
	// SERIAL6
	SERIAL_CONTROL_SERIAL6 SERIAL_CONTROL_DEV = 106
	// SERIAL7
	SERIAL_CONTROL_SERIAL7 SERIAL_CONTROL_DEV = 107
	// SERIAL8
	SERIAL_CONTROL_SERIAL8 SERIAL_CONTROL_DEV = 108
	// SERIAL9
	SERIAL_CONTROL_SERIAL9 SERIAL_CONTROL_DEV = 109
)

// SERIAL_CONTROL flags (bitmask)
type SERIAL_CONTROL_FLAG int

const (
	// Set if this is a reply
	SERIAL_CONTROL_FLAG_REPLY SERIAL_CONTROL_FLAG = 1
	// Set if the sender wants the receiver to send a response as another SERIAL_CONTROL message
	SERIAL_CONTROL_FLAG_RESPOND SERIAL_CONTROL_FLAG = 2
	// Set if access to the serial port should be removed from whatever driver is currently using it, giving exclusive access to the SERIAL_CONTROL protocol. The port can be handed back by sending a request without this flag set
	SERIAL_CONTROL_FLAG_EXCLUSIVE SERIAL_CONTROL_FLAG = 4
	// Block on writes to the serial port
	SERIAL_CONTROL_FLAG_BLOCKING SERIAL_CONTROL_FLAG = 8
	// Send multiple replies until port is drained
	SERIAL_CONTROL_FLAG_MULTI SERIAL_CONTROL_FLAG = 16
)

// Focus types for MAV_CMD_SET_CAMERA_FOCUS
type SET_FOCUS_TYPE int

const (
	// Focus one step increment (-1 for focusing in, 1 for focusing out towards infinity).
	FOCUS_TYPE_STEP SET_FOCUS_TYPE = 0
	// Continuous focus up/down until stopped (-1 for focusing in, 1 for focusing out towards infinity, 0 to stop focusing)
	FOCUS_TYPE_CONTINUOUS SET_FOCUS_TYPE = 1
	// Focus value as proportion of full camera focus range (a value between 0.0 and 100.0)
	FOCUS_TYPE_RANGE SET_FOCUS_TYPE = 2
	// Focus value in metres. Note that there is no message to get the valid focus range of the camera, so this can type can only be used for cameras where the range is known (implying that this cannot reliably be used in a GCS for an arbitrary camera).
	FOCUS_TYPE_METERS SET_FOCUS_TYPE = 3
)

// Flags to indicate the status of camera storage.
type STORAGE_STATUS int

const (
	// Storage is missing (no microSD card loaded for example.)
	STORAGE_STATUS_EMPTY STORAGE_STATUS = 0
	// Storage present but unformatted.
	STORAGE_STATUS_UNFORMATTED STORAGE_STATUS = 1
	// Storage present and ready.
	STORAGE_STATUS_READY STORAGE_STATUS = 2
	// Camera does not supply storage status information. Capacity information in STORAGE_INFORMATION fields will be ignored.
	STORAGE_STATUS_NOT_SUPPORTED STORAGE_STATUS = 3
)

// Tune formats (used for vehicle buzzer/tone generation).
type TUNE_FORMAT int

const (
	// Format is QBasic 1.1 Play: https://www.qbasic.net/en/reference/qb11/Statement/PLAY-006.htm.
	TUNE_FORMAT_QBASIC1_1 TUNE_FORMAT = 1
	// Format is Modern Music Markup Language (MML): https://en.wikipedia.org/wiki/Music_Macro_Language#Modern_MML.
	TUNE_FORMAT_MML_MODERN TUNE_FORMAT = 2
)

// Generalized UAVCAN node health
type UAVCAN_NODE_HEALTH int

const (
	// The node is functioning properly.
	UAVCAN_NODE_HEALTH_OK UAVCAN_NODE_HEALTH = 0
	// A critical parameter went out of range or the node has encountered a minor failure.
	UAVCAN_NODE_HEALTH_WARNING UAVCAN_NODE_HEALTH = 1
	// The node has encountered a major failure.
	UAVCAN_NODE_HEALTH_ERROR UAVCAN_NODE_HEALTH = 2
	// The node has suffered a fatal malfunction.
	UAVCAN_NODE_HEALTH_CRITICAL UAVCAN_NODE_HEALTH = 3
)

// Generalized UAVCAN node mode
type UAVCAN_NODE_MODE int

const (
	// The node is performing its primary functions.
	UAVCAN_NODE_MODE_OPERATIONAL UAVCAN_NODE_MODE = 0
	// The node is initializing; this mode is entered immediately after startup.
	UAVCAN_NODE_MODE_INITIALIZATION UAVCAN_NODE_MODE = 1
	// The node is under maintenance.
	UAVCAN_NODE_MODE_MAINTENANCE UAVCAN_NODE_MODE = 2
	// The node is in the process of updating its software.
	UAVCAN_NODE_MODE_SOFTWARE_UPDATE UAVCAN_NODE_MODE = 3
	// The node is no longer available online.
	UAVCAN_NODE_MODE_OFFLINE UAVCAN_NODE_MODE = 7
)

// Flags for the global position report.
type UTM_DATA_AVAIL_FLAGS int

const (
	// The field time contains valid data.
	UTM_DATA_AVAIL_FLAGS_TIME_VALID UTM_DATA_AVAIL_FLAGS = 1
	// The field uas_id contains valid data.
	UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE UTM_DATA_AVAIL_FLAGS = 2
	// The fields lat, lon and h_acc contain valid data.
	UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE UTM_DATA_AVAIL_FLAGS = 4
	// The fields alt and v_acc contain valid data.
	UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE UTM_DATA_AVAIL_FLAGS = 8
	// The field relative_alt contains valid data.
	UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE UTM_DATA_AVAIL_FLAGS = 16
	// The fields vx and vy contain valid data.
	UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE UTM_DATA_AVAIL_FLAGS = 32
	// The field vz contains valid data.
	UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE UTM_DATA_AVAIL_FLAGS = 64
	// The fields next_lat, next_lon and next_alt contain valid data.
	UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE UTM_DATA_AVAIL_FLAGS = 128
)

// Airborne status of UAS.
type UTM_FLIGHT_STATE int

const (
	// The flight state can't be determined.
	UTM_FLIGHT_STATE_UNKNOWN UTM_FLIGHT_STATE = 1
	// UAS on ground.
	UTM_FLIGHT_STATE_GROUND UTM_FLIGHT_STATE = 2
	// UAS airborne.
	UTM_FLIGHT_STATE_AIRBORNE UTM_FLIGHT_STATE = 3
	// UAS is in an emergency flight state.
	UTM_FLIGHT_STATE_EMERGENCY UTM_FLIGHT_STATE = 16
	// UAS has no active controls.
	UTM_FLIGHT_STATE_NOCTRL UTM_FLIGHT_STATE = 32
)

// Stream status flags (Bitmap)
type VIDEO_STREAM_STATUS_FLAGS int

const (
	// Stream is active (running)
	VIDEO_STREAM_STATUS_FLAGS_RUNNING VIDEO_STREAM_STATUS_FLAGS = 1
	// Stream is thermal imaging
	VIDEO_STREAM_STATUS_FLAGS_THERMAL VIDEO_STREAM_STATUS_FLAGS = 2
)

// Video stream types
type VIDEO_STREAM_TYPE int

const (
	// Stream is RTSP
	VIDEO_STREAM_TYPE_RTSP VIDEO_STREAM_TYPE = 0
	// Stream is RTP UDP (URI gives the port number)
	VIDEO_STREAM_TYPE_RTPUDP VIDEO_STREAM_TYPE = 1
	// Stream is MPEG on TCP
	VIDEO_STREAM_TYPE_TCP_MPEG VIDEO_STREAM_TYPE = 2
	// Stream is h.264 on MPEG TS (URI gives the port number)
	VIDEO_STREAM_TYPE_MPEG_TS_H264 VIDEO_STREAM_TYPE = 3
)

// Direction of VTOL transition
type VTOL_TRANSITION_HEADING int

const (
	// Respect the heading configuration of the vehicle.
	VTOL_TRANSITION_HEADING_VEHICLE_DEFAULT VTOL_TRANSITION_HEADING = 0
	// Use the heading pointing towards the next waypoint.
	VTOL_TRANSITION_HEADING_NEXT_WAYPOINT VTOL_TRANSITION_HEADING = 1
	// Use the heading on takeoff (while sitting on the ground).
	VTOL_TRANSITION_HEADING_TAKEOFF VTOL_TRANSITION_HEADING = 2
	// Use the specified heading in parameter 4.
	VTOL_TRANSITION_HEADING_SPECIFIED VTOL_TRANSITION_HEADING = 3
	// Use the current heading when reaching takeoff altitude (potentially facing the wind when weather-vaning is active).
	VTOL_TRANSITION_HEADING_ANY VTOL_TRANSITION_HEADING = 4
)

// common.xml

// The heartbeat message shows that a system or component is present and responding. The type and autopilot fields (along with the message component id), allow the receiving system to treat further messages from this system appropriately (e.g. by laying out the user interface based on the autopilot). This microservice is documented at https://mavlink.io/en/services/heartbeat.html
type MessageHeartbeat struct {
	// Vehicle or component type. For a flight controller component the vehicle type (quadrotor, helicopter, etc.). For other components the component type (e.g. camera, gimbal, etc.). This should be used in preference to component id for identifying the component type.
	Type MAV_TYPE `mavenum:"uint8"`
	// Autopilot type / class. Use MAV_AUTOPILOT_INVALID for components that are not flight controllers.
	Autopilot MAV_AUTOPILOT `mavenum:"uint8"`
	// System mode bitmap.
	BaseMode MAV_MODE_FLAG `mavenum:"uint8"`
	// A bitfield for use for autopilot-specific flags
	CustomMode uint32
	// System status flag.
	SystemStatus MAV_STATE `mavenum:"uint8"`
	// MAVLink version, not writable by user, gets added by protocol because of magic data type: uint8_t_mavlink_version
	MavlinkVersion uint8
}

func (m *MessageHeartbeat) GetId() uint32 {
	return 0
}

func (m *MessageHeartbeat) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The general system state. If the system is following the MAVLink standard, the system state is mainly defined by three orthogonal states/modes: The system mode, which is either LOCKED (motors shut down and locked), MANUAL (system under RC control), GUIDED (system with autonomous position control, position setpoint controlled manually) or AUTO (system guided by path/waypoint planner). The NAV_MODE defined the current flight state: LIFTOFF (often an open-loop maneuver), LANDING, WAYPOINTS or VECTOR. This represents the internal navigation state machine. The system status shows whether the system is currently active or not and if an emergency occurred. During the CRITICAL and EMERGENCY states the MAV is still considered to be active, but should start emergency procedures autonomously. After a failure occurred it should first move from active to critical to allow manual intervention and then move to emergency after a certain timeout.
type MessageSysStatus struct {
	// Bitmap showing which onboard controllers and sensors are present. Value of 0: not present. Value of 1: present.
	OnboardControlSensorsPresent MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	// Bitmap showing which onboard controllers and sensors are enabled:  Value of 0: not enabled. Value of 1: enabled.
	OnboardControlSensorsEnabled MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	// Bitmap showing which onboard controllers and sensors have an error (or are operational). Value of 0: error. Value of 1: healthy.
	OnboardControlSensorsHealth MAV_SYS_STATUS_SENSOR `mavenum:"uint32"`
	// Maximum usage in percent of the mainloop time. Values: [0-1000] - should always be below 1000
	Load uint16
	// Battery voltage, UINT16_MAX: Voltage not sent by autopilot
	VoltageBattery uint16
	// Battery current, -1: Current not sent by autopilot
	CurrentBattery int16
	// Battery energy remaining, -1: Battery remaining energy not sent by autopilot
	BatteryRemaining int8
	// Communication drop rate, (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)
	DropRateComm uint16
	// Communication errors (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)
	ErrorsComm uint16
	// Autopilot-specific errors
	ErrorsCount1 uint16
	// Autopilot-specific errors
	ErrorsCount2 uint16
	// Autopilot-specific errors
	ErrorsCount3 uint16
	// Autopilot-specific errors
	ErrorsCount4 uint16
}

func (m *MessageSysStatus) GetId() uint32 {
	return 1
}

func (m *MessageSysStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The system time is the time of the master clock, typically the computer clock of the main onboard computer.
type MessageSystemTime struct {
	// Timestamp (UNIX epoch time).
	TimeUnixUsec uint64
	// Timestamp (time since system boot).
	TimeBootMs uint32
}

func (m *MessageSystemTime) GetId() uint32 {
	return 2
}

func (m *MessageSystemTime) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// A ping message either requesting or responding to a ping. This allows to measure the system latencies, including serial port, radio modem and UDP connections. The ping microservice is documented at https://mavlink.io/en/services/ping.html
type MessagePing struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// PING sequence
	Seq uint32
	// 0: request ping from all receiving systems. If greater than 0: message is a ping response and number is the system id of the requesting system
	TargetSystem uint8
	// 0: request ping from all receiving components. If greater than 0: message is a ping response and number is the component id of the requesting component.
	TargetComponent uint8
}

func (m *MessagePing) GetId() uint32 {
	return 4
}

func (m *MessagePing) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request to control this MAV
type MessageChangeOperatorControl struct {
	// System the GCS requests control for
	TargetSystem uint8
	// 0: request control of this MAV, 1: Release control of this MAV
	ControlRequest uint8
	// 0: key as plaintext, 1-255: future, different hashing/encryption variants. The GCS should in general use the safest mode possible initially and then gradually move down the encryption level if it gets a NACK message indicating an encryption mismatch.
	Version uint8
	// Password / Key, depending on version plaintext or encrypted. 25 or less characters, NULL terminated. The characters may involve A-Z, a-z, 0-9, and "!?,.-"
	Passkey string `mavlen:"25"`
}

func (m *MessageChangeOperatorControl) GetId() uint32 {
	return 5
}

func (m *MessageChangeOperatorControl) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Accept / deny control of this MAV
type MessageChangeOperatorControlAck struct {
	// ID of the GCS this message
	GcsSystemId uint8
	// 0: request control of this MAV, 1: Release control of this MAV
	ControlRequest uint8
	// 0: ACK, 1: NACK: Wrong passkey, 2: NACK: Unsupported passkey encryption method, 3: NACK: Already under control
	Ack uint8
}

func (m *MessageChangeOperatorControlAck) GetId() uint32 {
	return 6
}

func (m *MessageChangeOperatorControlAck) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Emit an encrypted signature / key identifying this system. PLEASE NOTE: This protocol has been kept simple, so transmitting the key requires an encrypted channel for true safety.
type MessageAuthKey struct {
	// key
	Key string `mavlen:"32"`
}

func (m *MessageAuthKey) GetId() uint32 {
	return 7
}

func (m *MessageAuthKey) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Status generated in each node in the communication chain and injected into MAVLink stream.
type MessageLinkNodeStatus struct {
	// Timestamp (time since system boot).
	Timestamp uint64
	// Remaining free transmit buffer space
	TxBuf uint8
	// Remaining free receive buffer space
	RxBuf uint8
	// Transmit rate
	TxRate uint32
	// Receive rate
	RxRate uint32
	// Number of bytes that could not be parsed correctly.
	RxParseErr uint16
	// Transmit buffer overflows. This number wraps around as it reaches UINT16_MAX
	TxOverflows uint16
	// Receive buffer overflows. This number wraps around as it reaches UINT16_MAX
	RxOverflows uint16
	// Messages sent
	MessagesSent uint32
	// Messages received (estimated from counting seq)
	MessagesReceived uint32
	// Messages lost (estimated from counting seq)
	MessagesLost uint32
}

func (m *MessageLinkNodeStatus) GetId() uint32 {
	return 8
}

func (m *MessageLinkNodeStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Set the system mode, as defined by enum MAV_MODE. There is no target component id as the mode is by definition for the overall aircraft, not only for one component.
type MessageSetMode struct {
	// The system setting the mode
	TargetSystem uint8
	// The new base mode.
	BaseMode MAV_MODE `mavenum:"uint8"`
	// The new autopilot-specific mode. This field can be ignored by an autopilot.
	CustomMode uint32
}

func (m *MessageSetMode) GetId() uint32 {
	return 11
}

func (m *MessageSetMode) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request to read the onboard parameter with the param_id string id. Onboard parameters are stored as key[const char*] -> value[float]. This allows to send a parameter to any other component (such as the GCS) without the need of previous knowledge of possible parameter names. Thus the same GCS can store different parameters for different autopilots. See also https://mavlink.io/en/services/parameter.html for a full documentation of QGroundControl and IMU code.
type MessageParamRequestRead struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored)
	ParamIndex int16
}

func (m *MessageParamRequestRead) GetId() uint32 {
	return 20
}

func (m *MessageParamRequestRead) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request all parameters of this component. After this request, all parameters are emitted. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html
type MessageParamRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

func (m *MessageParamRequestList) GetId() uint32 {
	return 21
}

func (m *MessageParamRequestList) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Emit the value of a onboard parameter. The inclusion of param_count and param_index in the message allows the recipient to keep track of received parameters and allows him to re-request missing parameters after a loss or timeout. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html
type MessageParamValue struct {
	// Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Onboard parameter value
	ParamValue float32
	// Onboard parameter type.
	ParamType MAV_PARAM_TYPE `mavenum:"uint8"`
	// Total number of onboard parameters
	ParamCount uint16
	// Index of this onboard parameter
	ParamIndex uint16
}

func (m *MessageParamValue) GetId() uint32 {
	return 22
}

func (m *MessageParamValue) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Set a parameter value (write new value to permanent storage). IMPORTANT: The receiving component should acknowledge the new parameter value by sending a PARAM_VALUE message to all communication partners. This will also ensure that multiple GCS all have an up-to-date list of all parameters. If the sending GCS did not receive a PARAM_VALUE message within its timeout time, it should re-send the PARAM_SET message. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html
type MessageParamSet struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Onboard parameter value
	ParamValue float32
	// Onboard parameter type.
	ParamType MAV_PARAM_TYPE `mavenum:"uint8"`
}

func (m *MessageParamSet) GetId() uint32 {
	return 23
}

func (m *MessageParamSet) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The global position, as returned by the Global Positioning System (GPS). This is                NOT the global position estimate of the system, but rather a RAW sensor value. See message GLOBAL_POSITION for the global position estimate.
type MessageGpsRawInt struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// GPS fix type.
	FixType GPS_FIX_TYPE `mavenum:"uint8"`
	// Latitude (WGS84, EGM96 ellipsoid)
	Lat int32
	// Longitude (WGS84, EGM96 ellipsoid)
	Lon int32
	// Altitude (MSL). Positive for up. Note that virtually all GPS modules provide the MSL altitude in addition to the WGS84 altitude.
	Alt int32
	// GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX
	Eph uint16
	// GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX
	Epv uint16
	// GPS ground speed. If unknown, set to: UINT16_MAX
	Vel uint16
	// Course over ground (NOT heading, but direction of movement) in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	Cog uint16
	// Number of satellites visible. If unknown, set to 255
	SatellitesVisible uint8
	// Altitude (above WGS84, EGM96 ellipsoid). Positive for up.
	AltEllipsoid int32 `mavext:"true"`
	// Position uncertainty.
	HAcc uint32 `mavext:"true"`
	// Altitude uncertainty.
	VAcc uint32 `mavext:"true"`
	// Speed uncertainty.
	VelAcc uint32 `mavext:"true"`
	// Heading / track uncertainty
	HdgAcc uint32 `mavext:"true"`
	// Yaw in earth frame from north. Use 0 if this GPS does not provide yaw. Use 65535 if this GPS is configured to provide yaw and is currently unable to provide it. Use 36000 for north.
	Yaw uint16 `mavext:"true"`
}

func (m *MessageGpsRawInt) GetId() uint32 {
	return 24
}

func (m *MessageGpsRawInt) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The positioning status, as reported by GPS. This message is intended to display status information about each satellite visible to the receiver. See message GLOBAL_POSITION for the global position estimate. This message can contain information for up to 20 satellites.
type MessageGpsStatus struct {
	// Number of satellites visible
	SatellitesVisible uint8
	// Global satellite ID
	SatellitePrn [20]uint8
	// 0: Satellite not used, 1: used for localization
	SatelliteUsed [20]uint8
	// Elevation (0: right on top of receiver, 90: on the horizon) of satellite
	SatelliteElevation [20]uint8
	// Direction of satellite, 0: 0 deg, 255: 360 deg.
	SatelliteAzimuth [20]uint8
	// Signal to noise ratio of satellite
	SatelliteSnr [20]uint8
}

func (m *MessageGpsStatus) GetId() uint32 {
	return 25
}

func (m *MessageGpsStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The RAW IMU readings for the usual 9DOF sensor setup. This message should contain the scaled values to the described units
type MessageScaledImu struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// X acceleration
	Xacc int16
	// Y acceleration
	Yacc int16
	// Z acceleration
	Zacc int16
	// Angular speed around X axis
	Xgyro int16
	// Angular speed around Y axis
	Ygyro int16
	// Angular speed around Z axis
	Zgyro int16
	// X Magnetic field
	Xmag int16
	// Y Magnetic field
	Ymag int16
	// Z Magnetic field
	Zmag int16
	// Temperature, 0: IMU does not provide temperature values. If the IMU is at 0C it must send 1 (0.01C).
	Temperature int16 `mavext:"true"`
}

func (m *MessageScaledImu) GetId() uint32 {
	return 26
}

func (m *MessageScaledImu) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The RAW IMU readings for a 9DOF sensor, which is identified by the id (default IMU1). This message should always contain the true raw values without any scaling to allow data capture and system debugging.
type MessageRawImu struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// X acceleration (raw)
	Xacc int16
	// Y acceleration (raw)
	Yacc int16
	// Z acceleration (raw)
	Zacc int16
	// Angular speed around X axis (raw)
	Xgyro int16
	// Angular speed around Y axis (raw)
	Ygyro int16
	// Angular speed around Z axis (raw)
	Zgyro int16
	// X Magnetic field (raw)
	Xmag int16
	// Y Magnetic field (raw)
	Ymag int16
	// Z Magnetic field (raw)
	Zmag int16
	// Id. Ids are numbered from 0 and map to IMUs numbered from 1 (e.g. IMU1 will have a message with id=0)
	Id uint8 `mavext:"true"`
	// Temperature, 0: IMU does not provide temperature values. If the IMU is at 0C it must send 1 (0.01C).
	Temperature int16 `mavext:"true"`
}

func (m *MessageRawImu) GetId() uint32 {
	return 27
}

func (m *MessageRawImu) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The RAW pressure readings for the typical setup of one absolute pressure and one differential pressure sensor. The sensor values should be the raw, UNSCALED ADC values.
type MessageRawPressure struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Absolute pressure (raw)
	PressAbs int16
	// Differential pressure 1 (raw, 0 if nonexistent)
	PressDiff1 int16
	// Differential pressure 2 (raw, 0 if nonexistent)
	PressDiff2 int16
	// Raw Temperature measurement (raw)
	Temperature int16
}

func (m *MessageRawPressure) GetId() uint32 {
	return 28
}

func (m *MessageRawPressure) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The pressure readings for the typical setup of one absolute and differential pressure sensor. The units are as specified in each field.
type MessageScaledPressure struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Absolute pressure
	PressAbs float32
	// Differential pressure 1
	PressDiff float32
	// Temperature
	Temperature int16
}

func (m *MessageScaledPressure) GetId() uint32 {
	return 29
}

func (m *MessageScaledPressure) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The attitude in the aeronautical frame (right-handed, Z-down, X-front, Y-right).
type MessageAttitude struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Roll angle (-pi..+pi)
	Roll float32
	// Pitch angle (-pi..+pi)
	Pitch float32
	// Yaw angle (-pi..+pi)
	Yaw float32
	// Roll angular speed
	Rollspeed float32
	// Pitch angular speed
	Pitchspeed float32
	// Yaw angular speed
	Yawspeed float32
}

func (m *MessageAttitude) GetId() uint32 {
	return 30
}

func (m *MessageAttitude) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The attitude in the aeronautical frame (right-handed, Z-down, X-front, Y-right), expressed as quaternion. Quaternion order is w, x, y, z and a zero rotation would be expressed as (1 0 0 0).
type MessageAttitudeQuaternion struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Quaternion component 1, w (1 in null-rotation)
	Q1 float32
	// Quaternion component 2, x (0 in null-rotation)
	Q2 float32
	// Quaternion component 3, y (0 in null-rotation)
	Q3 float32
	// Quaternion component 4, z (0 in null-rotation)
	Q4 float32
	// Roll angular speed
	Rollspeed float32
	// Pitch angular speed
	Pitchspeed float32
	// Yaw angular speed
	Yawspeed float32
	// Rotation offset by which the attitude quaternion and angular speed vector should be rotated for user display (quaternion with [w, x, y, z] order, zero-rotation is [1, 0, 0, 0], send [0, 0, 0, 0] if field not supported). This field is intended for systems in which the reference attitude may change during flight. For example, tailsitters VTOLs rotate their reference attitude by 90 degrees between hover mode and fixed wing mode, thus repr_offset_q is equal to [1, 0, 0, 0] in hover mode and equal to [0.7071, 0, 0.7071, 0] in fixed wing mode.
	ReprOffsetQ [4]float32 `mavext:"true"`
}

func (m *MessageAttitudeQuaternion) GetId() uint32 {
	return 31
}

func (m *MessageAttitudeQuaternion) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The filtered local position (e.g. fused computer vision and accelerometers). Coordinate frame is right-handed, Z-axis down (aeronautical frame, NED / north-east-down convention)
type MessageLocalPositionNed struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// X Position
	X float32
	// Y Position
	Y float32
	// Z Position
	Z float32
	// X Speed
	Vx float32
	// Y Speed
	Vy float32
	// Z Speed
	Vz float32
}

func (m *MessageLocalPositionNed) GetId() uint32 {
	return 32
}

func (m *MessageLocalPositionNed) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The filtered global position (e.g. fused GPS and accelerometers). The position is in GPS-frame (right-handed, Z-up). It               is designed as scaled integer message since the resolution of float is not sufficient.
type MessageGlobalPositionInt struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Latitude, expressed
	Lat int32
	// Longitude, expressed
	Lon int32
	// Altitude (MSL). Note that virtually all GPS modules provide both WGS84 and MSL.
	Alt int32
	// Altitude above ground
	RelativeAlt int32
	// Ground X Speed (Latitude, positive north)
	Vx int16
	// Ground Y Speed (Longitude, positive east)
	Vy int16
	// Ground Z Speed (Altitude, positive down)
	Vz int16
	// Vehicle heading (yaw angle), 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	Hdg uint16
}

func (m *MessageGlobalPositionInt) GetId() uint32 {
	return 33
}

func (m *MessageGlobalPositionInt) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The scaled values of the RC channels received: (-100%) -10000, (0%) 0, (100%) 10000. Channels that are inactive should be set to UINT16_MAX.
type MessageRcChannelsScaled struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Port uint8
	// RC channel 1 value scaled.
	Chan1Scaled int16
	// RC channel 2 value scaled.
	Chan2Scaled int16
	// RC channel 3 value scaled.
	Chan3Scaled int16
	// RC channel 4 value scaled.
	Chan4Scaled int16
	// RC channel 5 value scaled.
	Chan5Scaled int16
	// RC channel 6 value scaled.
	Chan6Scaled int16
	// RC channel 7 value scaled.
	Chan7Scaled int16
	// RC channel 8 value scaled.
	Chan8Scaled int16
	// Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Rssi uint8
}

func (m *MessageRcChannelsScaled) GetId() uint32 {
	return 34
}

func (m *MessageRcChannelsScaled) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The RAW values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. A value of UINT16_MAX implies the channel is unused. Individual receivers/transmitters might violate this specification.
type MessageRcChannelsRaw struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Port uint8
	// RC channel 1 value.
	Chan1Raw uint16
	// RC channel 2 value.
	Chan2Raw uint16
	// RC channel 3 value.
	Chan3Raw uint16
	// RC channel 4 value.
	Chan4Raw uint16
	// RC channel 5 value.
	Chan5Raw uint16
	// RC channel 6 value.
	Chan6Raw uint16
	// RC channel 7 value.
	Chan7Raw uint16
	// RC channel 8 value.
	Chan8Raw uint16
	// Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Rssi uint8
}

func (m *MessageRcChannelsRaw) GetId() uint32 {
	return 35
}

func (m *MessageRcChannelsRaw) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Superseded by ACTUATOR_OUTPUT_STATUS. The RAW values of the servo outputs (for RC input from the remote, use the RC_CHANNELS messages). The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%.
type MessageServoOutputRaw struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint32
	// Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Port uint8
	// Servo output 1 value
	Servo1Raw uint16
	// Servo output 2 value
	Servo2Raw uint16
	// Servo output 3 value
	Servo3Raw uint16
	// Servo output 4 value
	Servo4Raw uint16
	// Servo output 5 value
	Servo5Raw uint16
	// Servo output 6 value
	Servo6Raw uint16
	// Servo output 7 value
	Servo7Raw uint16
	// Servo output 8 value
	Servo8Raw uint16
	// Servo output 9 value
	Servo9Raw uint16 `mavext:"true"`
	// Servo output 10 value
	Servo10Raw uint16 `mavext:"true"`
	// Servo output 11 value
	Servo11Raw uint16 `mavext:"true"`
	// Servo output 12 value
	Servo12Raw uint16 `mavext:"true"`
	// Servo output 13 value
	Servo13Raw uint16 `mavext:"true"`
	// Servo output 14 value
	Servo14Raw uint16 `mavext:"true"`
	// Servo output 15 value
	Servo15Raw uint16 `mavext:"true"`
	// Servo output 16 value
	Servo16Raw uint16 `mavext:"true"`
}

func (m *MessageServoOutputRaw) GetId() uint32 {
	return 36
}

func (m *MessageServoOutputRaw) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request a partial list of mission items from the system/component. https://mavlink.io/en/services/mission.html. If start and end index are the same, just send one waypoint.
type MessageMissionRequestPartialList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Start index
	StartIndex int16
	// End index, -1 by default (-1: send list to end). Else a valid index of the list
	EndIndex int16
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionRequestPartialList) GetId() uint32 {
	return 37
}

func (m *MessageMissionRequestPartialList) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// This message is sent to the MAV to write a partial list. If start index == end index, only one item will be transmitted / updated. If the start index is NOT 0 and above the current list size, this request should be REJECTED!
type MessageMissionWritePartialList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Start index. Must be smaller / equal to the largest index of the current onboard list.
	StartIndex int16
	// End index, equal or greater than start index.
	EndIndex int16
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionWritePartialList) GetId() uint32 {
	return 38
}

func (m *MessageMissionWritePartialList) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Message encoding a mission item. This message is emitted to announce                the presence of a mission item and to set a mission item on the system. The mission item can be either in x, y, z meters (type: LOCAL) or x:lat, y:lon, z:altitude. Local frame is Z-down, right handed (NED), global frame is Z-up, right handed (ENU). NaN may be used to indicate an optional/default value (e.g. to use the system's current latitude or yaw rather than a specific value). See also https://mavlink.io/en/services/mission.html.
type MessageMissionItem struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Sequence
	Seq uint16
	// The coordinate system of the waypoint.
	Frame MAV_FRAME `mavenum:"uint8"`
	// The scheduled action for the waypoint.
	Command MAV_CMD `mavenum:"uint16"`
	// false:0, true:1
	Current uint8
	// Autocontinue to next waypoint
	Autocontinue uint8
	// PARAM1, see MAV_CMD enum
	Param1 float32
	// PARAM2, see MAV_CMD enum
	Param2 float32
	// PARAM3, see MAV_CMD enum
	Param3 float32
	// PARAM4, see MAV_CMD enum
	Param4 float32
	// PARAM5 / local: X coordinate, global: latitude
	X float32
	// PARAM6 / local: Y coordinate, global: longitude
	Y float32
	// PARAM7 / local: Z coordinate, global: altitude (relative or absolute, depending on frame).
	Z float32
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionItem) GetId() uint32 {
	return 39
}

func (m *MessageMissionItem) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request the information of the mission item with the sequence number seq. The response of the system to this message should be a MISSION_ITEM message. https://mavlink.io/en/services/mission.html
type MessageMissionRequest struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Sequence
	Seq uint16
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionRequest) GetId() uint32 {
	return 40
}

func (m *MessageMissionRequest) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Set the mission item with sequence number seq as current item. This means that the MAV will continue to this mission item on the shortest path (not following the mission items in-between).
type MessageMissionSetCurrent struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Sequence
	Seq uint16
}

func (m *MessageMissionSetCurrent) GetId() uint32 {
	return 41
}

func (m *MessageMissionSetCurrent) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Message that announces the sequence number of the current active mission item. The MAV will fly towards this mission item.
type MessageMissionCurrent struct {
	// Sequence
	Seq uint16
}

func (m *MessageMissionCurrent) GetId() uint32 {
	return 42
}

func (m *MessageMissionCurrent) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request the overall list of mission items from the system/component.
type MessageMissionRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionRequestList) GetId() uint32 {
	return 43
}

func (m *MessageMissionRequestList) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// This message is emitted as response to MISSION_REQUEST_LIST by the MAV and to initiate a write transaction. The GCS can then request the individual mission item based on the knowledge of the total number of waypoints.
type MessageMissionCount struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Number of mission items in the sequence
	Count uint16
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionCount) GetId() uint32 {
	return 44
}

func (m *MessageMissionCount) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Delete all mission items at once.
type MessageMissionClearAll struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionClearAll) GetId() uint32 {
	return 45
}

func (m *MessageMissionClearAll) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// A certain mission item has been reached. The system will either hold this position (or circle on the orbit) or (if the autocontinue on the WP was set) continue to the next waypoint.
type MessageMissionItemReached struct {
	// Sequence
	Seq uint16
}

func (m *MessageMissionItemReached) GetId() uint32 {
	return 46
}

func (m *MessageMissionItemReached) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Acknowledgment message during waypoint handling. The type field states if this message is a positive ack (type=0) or if an error happened (type=non-zero).
type MessageMissionAck struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Mission result.
	Type MAV_MISSION_RESULT `mavenum:"uint8"`
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionAck) GetId() uint32 {
	return 47
}

func (m *MessageMissionAck) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Sets the GPS co-ordinates of the vehicle local origin (0,0,0) position. Vehicle should emit GPS_GLOBAL_ORIGIN irrespective of whether the origin is changed. This enables transform between the local coordinate frame and the global (GPS) coordinate frame, which may be necessary when (for example) indoor and outdoor settings are connected and the MAV should move from in- to outdoor.
type MessageSetGpsGlobalOrigin struct {
	// System ID
	TargetSystem uint8
	// Latitude (WGS84)
	Latitude int32
	// Longitude (WGS84)
	Longitude int32
	// Altitude (MSL). Positive for up.
	Altitude int32
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64 `mavext:"true"`
}

func (m *MessageSetGpsGlobalOrigin) GetId() uint32 {
	return 48
}

func (m *MessageSetGpsGlobalOrigin) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Publishes the GPS co-ordinates of the vehicle local origin (0,0,0) position. Emitted whenever a new GPS-Local position mapping is requested or set - e.g. following SET_GPS_GLOBAL_ORIGIN message.
type MessageGpsGlobalOrigin struct {
	// Latitude (WGS84)
	Latitude int32
	// Longitude (WGS84)
	Longitude int32
	// Altitude (MSL). Positive for up.
	Altitude int32
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64 `mavext:"true"`
}

func (m *MessageGpsGlobalOrigin) GetId() uint32 {
	return 49
}

func (m *MessageGpsGlobalOrigin) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Bind a RC channel to a parameter. The parameter should change according to the RC channel value.
type MessageParamMapRc struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored), send -2 to disable any existing map for this rc_channel_index.
	ParamIndex int16
	// Index of parameter RC channel. Not equal to the RC channel id. Typically corresponds to a potentiometer-knob on the RC.
	ParameterRcChannelIndex uint8
	// Initial parameter value
	ParamValue0 float32
	// Scale, maps the RC range [-1, 1] to a parameter value
	Scale float32
	// Minimum param value. The protocol does not define if this overwrites an onboard minimum value. (Depends on implementation)
	ParamValueMin float32
	// Maximum param value. The protocol does not define if this overwrites an onboard maximum value. (Depends on implementation)
	ParamValueMax float32
}

func (m *MessageParamMapRc) GetId() uint32 {
	return 50
}

func (m *MessageParamMapRc) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request the information of the mission item with the sequence number seq. The response of the system to this message should be a MISSION_ITEM_INT message. https://mavlink.io/en/services/mission.html
type MessageMissionRequestInt struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Sequence
	Seq uint16
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionRequestInt) GetId() uint32 {
	return 51
}

func (m *MessageMissionRequestInt) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// A broadcast message to notify any ground station or SDK if a mission, geofence or safe points have changed on the vehicle.
type MessageMissionChanged struct {
	// Start index for partial mission change (-1 for all items).
	StartIndex int16
	// End index of a partial mission change. -1 is a synonym for the last mission item (i.e. selects all items from start_index). Ignore field if start_index=-1.
	EndIndex int16
	// System ID of the author of the new mission.
	OriginSysid uint8
	// Compnent ID of the author of the new mission.
	OriginCompid MAV_COMPONENT `mavenum:"uint8"`
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8"`
}

func (m *MessageMissionChanged) GetId() uint32 {
	return 52
}

func (m *MessageMissionChanged) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Set a safety zone (volume), which is defined by two corners of a cube. This message can be used to tell the MAV which setpoints/waypoints to accept and which to reject. Safety areas are often enforced by national or competition regulations.
type MessageSafetySetAllowedArea struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Coordinate frame. Can be either global, GPS, right-handed with Z axis up or local, right handed, Z axis down.
	Frame MAV_FRAME `mavenum:"uint8"`
	// x position 1 / Latitude 1
	P1x float32
	// y position 1 / Longitude 1
	P1y float32
	// z position 1 / Altitude 1
	P1z float32
	// x position 2 / Latitude 2
	P2x float32
	// y position 2 / Longitude 2
	P2y float32
	// z position 2 / Altitude 2
	P2z float32
}

func (m *MessageSafetySetAllowedArea) GetId() uint32 {
	return 54
}

func (m *MessageSafetySetAllowedArea) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Read out the safety zone the MAV currently assumes.
type MessageSafetyAllowedArea struct {
	// Coordinate frame. Can be either global, GPS, right-handed with Z axis up or local, right handed, Z axis down.
	Frame MAV_FRAME `mavenum:"uint8"`
	// x position 1 / Latitude 1
	P1x float32
	// y position 1 / Longitude 1
	P1y float32
	// z position 1 / Altitude 1
	P1z float32
	// x position 2 / Latitude 2
	P2x float32
	// y position 2 / Longitude 2
	P2y float32
	// z position 2 / Altitude 2
	P2z float32
}

func (m *MessageSafetyAllowedArea) GetId() uint32 {
	return 55
}

func (m *MessageSafetyAllowedArea) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The attitude in the aeronautical frame (right-handed, Z-down, X-front, Y-right), expressed as quaternion. Quaternion order is w, x, y, z and a zero rotation would be expressed as (1 0 0 0).
type MessageAttitudeQuaternionCov struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation)
	Q [4]float32
	// Roll angular speed
	Rollspeed float32
	// Pitch angular speed
	Pitchspeed float32
	// Yaw angular speed
	Yawspeed float32
	// Row-major representation of a 3x3 attitude covariance matrix (states: roll, pitch, yaw; first three entries are the first ROW, next three entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [9]float32
}

func (m *MessageAttitudeQuaternionCov) GetId() uint32 {
	return 61
}

func (m *MessageAttitudeQuaternionCov) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The state of the fixed wing navigation and position controller.
type MessageNavControllerOutput struct {
	// Current desired roll
	NavRoll float32
	// Current desired pitch
	NavPitch float32
	// Current desired heading
	NavBearing int16
	// Bearing to current waypoint/target
	TargetBearing int16
	// Distance to active waypoint
	WpDist uint16
	// Current altitude error
	AltError float32
	// Current airspeed error
	AspdError float32
	// Current crosstrack error on x-y plane
	XtrackError float32
}

func (m *MessageNavControllerOutput) GetId() uint32 {
	return 62
}

func (m *MessageNavControllerOutput) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The filtered global position (e.g. fused GPS and accelerometers). The position is in GPS-frame (right-handed, Z-up). It  is designed as scaled integer message since the resolution of float is not sufficient. NOTE: This message is intended for onboard networks / companion computers and higher-bandwidth links and optimized for accuracy and completeness. Please use the GLOBAL_POSITION_INT message for a minimal subset.
type MessageGlobalPositionIntCov struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Class id of the estimator this estimate originated from.
	EstimatorType MAV_ESTIMATOR_TYPE `mavenum:"uint8"`
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// Altitude in meters above MSL
	Alt int32
	// Altitude above ground
	RelativeAlt int32
	// Ground X Speed (Latitude)
	Vx float32
	// Ground Y Speed (Longitude)
	Vy float32
	// Ground Z Speed (Altitude)
	Vz float32
	// Row-major representation of a 6x6 position and velocity 6x6 cross-covariance matrix (states: lat, lon, alt, vx, vy, vz; first six entries are the first ROW, next six entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [36]float32
}

func (m *MessageGlobalPositionIntCov) GetId() uint32 {
	return 63
}

func (m *MessageGlobalPositionIntCov) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The filtered local position (e.g. fused computer vision and accelerometers). Coordinate frame is right-handed, Z-axis down (aeronautical frame, NED / north-east-down convention)
type MessageLocalPositionNedCov struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Class id of the estimator this estimate originated from.
	EstimatorType MAV_ESTIMATOR_TYPE `mavenum:"uint8"`
	// X Position
	X float32
	// Y Position
	Y float32
	// Z Position
	Z float32
	// X Speed
	Vx float32
	// Y Speed
	Vy float32
	// Z Speed
	Vz float32
	// X Acceleration
	Ax float32
	// Y Acceleration
	Ay float32
	// Z Acceleration
	Az float32
	// Row-major representation of position, velocity and acceleration 9x9 cross-covariance matrix upper right triangle (states: x, y, z, vx, vy, vz, ax, ay, az; first nine entries are the first ROW, next eight entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [45]float32
}

func (m *MessageLocalPositionNedCov) GetId() uint32 {
	return 64
}

func (m *MessageLocalPositionNedCov) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The PPM values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%.  A value of UINT16_MAX implies the channel is unused. Individual receivers/transmitters might violate this specification.
type MessageRcChannels struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Total number of RC channels being received. This can be larger than 18, indicating that more channels are available but not given in this message. This value should be 0 when no RC channels are available.
	Chancount uint8
	// RC channel 1 value.
	Chan1Raw uint16
	// RC channel 2 value.
	Chan2Raw uint16
	// RC channel 3 value.
	Chan3Raw uint16
	// RC channel 4 value.
	Chan4Raw uint16
	// RC channel 5 value.
	Chan5Raw uint16
	// RC channel 6 value.
	Chan6Raw uint16
	// RC channel 7 value.
	Chan7Raw uint16
	// RC channel 8 value.
	Chan8Raw uint16
	// RC channel 9 value.
	Chan9Raw uint16
	// RC channel 10 value.
	Chan10Raw uint16
	// RC channel 11 value.
	Chan11Raw uint16
	// RC channel 12 value.
	Chan12Raw uint16
	// RC channel 13 value.
	Chan13Raw uint16
	// RC channel 14 value.
	Chan14Raw uint16
	// RC channel 15 value.
	Chan15Raw uint16
	// RC channel 16 value.
	Chan16Raw uint16
	// RC channel 17 value.
	Chan17Raw uint16
	// RC channel 18 value.
	Chan18Raw uint16
	// Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Rssi uint8
}

func (m *MessageRcChannels) GetId() uint32 {
	return 65
}

func (m *MessageRcChannels) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request a data stream.
type MessageRequestDataStream struct {
	// The target requested to send the message stream.
	TargetSystem uint8
	// The target requested to send the message stream.
	TargetComponent uint8
	// The ID of the requested data stream
	ReqStreamId uint8
	// The requested message rate
	ReqMessageRate uint16
	// 1 to start sending, 0 to stop sending.
	StartStop uint8
}

func (m *MessageRequestDataStream) GetId() uint32 {
	return 66
}

func (m *MessageRequestDataStream) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Data stream status information.
type MessageDataStream struct {
	// The ID of the requested data stream
	StreamId uint8
	// The message rate
	MessageRate uint16
	// 1 stream is enabled, 0 stream is stopped.
	OnOff uint8
}

func (m *MessageDataStream) GetId() uint32 {
	return 67
}

func (m *MessageDataStream) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// This message provides an API for manually controlling the vehicle using standard joystick axes nomenclature, along with a joystick-like input device. Unused axes can be disabled an buttons are also transmit as boolean values of their
type MessageManualControl struct {
	// The system to be controlled.
	Target uint8
	// X-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to forward(1000)-backward(-1000) movement on a joystick and the pitch of a vehicle.
	X int16
	// Y-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to left(-1000)-right(1000) movement on a joystick and the roll of a vehicle.
	Y int16
	// Z-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to a separate slider movement with maximum being 1000 and minimum being -1000 on a joystick and the thrust of a vehicle. Positive values are positive thrust, negative values are negative thrust.
	Z int16
	// R-axis, normalized to the range [-1000,1000]. A value of INT16_MAX indicates that this axis is invalid. Generally corresponds to a twisting of the joystick, with counter-clockwise being 1000 and clockwise being -1000, and the yaw of a vehicle.
	R int16
	// A bitfield corresponding to the joystick buttons' current state, 1 for pressed, 0 for released. The lowest bit corresponds to Button 1.
	Buttons uint16
}

func (m *MessageManualControl) GetId() uint32 {
	return 69
}

func (m *MessageManualControl) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The RAW values of the RC channels sent to the MAV to override info received from the RC radio. A value of UINT16_MAX means no change to that channel. A value of 0 means control of that channel should be released back to the RC radio. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. Individual receivers/transmitters might violate this specification.
type MessageRcChannelsOverride struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// RC channel 1 value. A value of UINT16_MAX means to ignore this field.
	Chan1Raw uint16
	// RC channel 2 value. A value of UINT16_MAX means to ignore this field.
	Chan2Raw uint16
	// RC channel 3 value. A value of UINT16_MAX means to ignore this field.
	Chan3Raw uint16
	// RC channel 4 value. A value of UINT16_MAX means to ignore this field.
	Chan4Raw uint16
	// RC channel 5 value. A value of UINT16_MAX means to ignore this field.
	Chan5Raw uint16
	// RC channel 6 value. A value of UINT16_MAX means to ignore this field.
	Chan6Raw uint16
	// RC channel 7 value. A value of UINT16_MAX means to ignore this field.
	Chan7Raw uint16
	// RC channel 8 value. A value of UINT16_MAX means to ignore this field.
	Chan8Raw uint16
	// RC channel 9 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan9Raw uint16 `mavext:"true"`
	// RC channel 10 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan10Raw uint16 `mavext:"true"`
	// RC channel 11 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan11Raw uint16 `mavext:"true"`
	// RC channel 12 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan12Raw uint16 `mavext:"true"`
	// RC channel 13 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan13Raw uint16 `mavext:"true"`
	// RC channel 14 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan14Raw uint16 `mavext:"true"`
	// RC channel 15 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan15Raw uint16 `mavext:"true"`
	// RC channel 16 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan16Raw uint16 `mavext:"true"`
	// RC channel 17 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan17Raw uint16 `mavext:"true"`
	// RC channel 18 value. A value of 0 or UINT16_MAX means to ignore this field.
	Chan18Raw uint16 `mavext:"true"`
}

func (m *MessageRcChannelsOverride) GetId() uint32 {
	return 70
}

func (m *MessageRcChannelsOverride) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Message encoding a mission item. This message is emitted to announce                the presence of a mission item and to set a mission item on the system. The mission item can be either in x, y, z meters (type: LOCAL) or x:lat, y:lon, z:altitude. Local frame is Z-down, right handed (NED), global frame is Z-up, right handed (ENU). NaN or INT32_MAX may be used in float/integer params (respectively) to indicate optional/default values (e.g. to use the component's current latitude, yaw rather than a specific value). See also https://mavlink.io/en/services/mission.html.
type MessageMissionItemInt struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Waypoint ID (sequence number). Starts at zero. Increases monotonically for each waypoint, no gaps in the sequence (0,1,2,3,4).
	Seq uint16
	// The coordinate system of the waypoint.
	Frame MAV_FRAME `mavenum:"uint8"`
	// The scheduled action for the waypoint.
	Command MAV_CMD `mavenum:"uint16"`
	// false:0, true:1
	Current uint8
	// Autocontinue to next waypoint
	Autocontinue uint8
	// PARAM1, see MAV_CMD enum
	Param1 float32
	// PARAM2, see MAV_CMD enum
	Param2 float32
	// PARAM3, see MAV_CMD enum
	Param3 float32
	// PARAM4, see MAV_CMD enum
	Param4 float32
	// PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	X int32
	// PARAM6 / y position: local: x position in meters * 1e4, global: longitude in degrees *10^7
	Y int32
	// PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame.
	Z float32
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageMissionItemInt) GetId() uint32 {
	return 73
}

func (m *MessageMissionItemInt) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Metrics typically displayed on a HUD for fixed wing aircraft.
type MessageVfrHud struct {
	// Current indicated airspeed (IAS).
	Airspeed float32
	// Current ground speed.
	Groundspeed float32
	// Current heading in compass units (0-360, 0=north).
	Heading int16
	// Current throttle setting (0 to 100).
	Throttle uint16
	// Current altitude (MSL).
	Alt float32
	// Current climb rate.
	Climb float32
}

func (m *MessageVfrHud) GetId() uint32 {
	return 74
}

func (m *MessageVfrHud) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Message encoding a command with parameters as scaled integers. Scaling depends on the actual command value. The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandInt struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// The coordinate system of the COMMAND.
	Frame MAV_FRAME `mavenum:"uint8"`
	// The scheduled action for the mission item.
	Command MAV_CMD `mavenum:"uint16"`
	// false:0, true:1
	Current uint8
	// autocontinue to next wp
	Autocontinue uint8
	// PARAM1, see MAV_CMD enum
	Param1 float32
	// PARAM2, see MAV_CMD enum
	Param2 float32
	// PARAM3, see MAV_CMD enum
	Param3 float32
	// PARAM4, see MAV_CMD enum
	Param4 float32
	// PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	X int32
	// PARAM6 / local: y position in meters * 1e4, global: longitude in degrees * 10^7
	Y int32
	// PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame).
	Z float32
}

func (m *MessageCommandInt) GetId() uint32 {
	return 75
}

func (m *MessageCommandInt) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Send a command with up to seven parameters to the MAV. The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandLong struct {
	// System which should execute the command
	TargetSystem uint8
	// Component which should execute the command, 0 for all components
	TargetComponent uint8
	// Command ID (of command to send).
	Command MAV_CMD `mavenum:"uint16"`
	// 0: First transmission of this command. 1-255: Confirmation transmissions (e.g. for kill command)
	Confirmation uint8
	// Parameter 1 (for the specific command).
	Param1 float32
	// Parameter 2 (for the specific command).
	Param2 float32
	// Parameter 3 (for the specific command).
	Param3 float32
	// Parameter 4 (for the specific command).
	Param4 float32
	// Parameter 5 (for the specific command).
	Param5 float32
	// Parameter 6 (for the specific command).
	Param6 float32
	// Parameter 7 (for the specific command).
	Param7 float32
}

func (m *MessageCommandLong) GetId() uint32 {
	return 76
}

func (m *MessageCommandLong) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Report status of a command. Includes feedback whether the command was executed. The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandAck struct {
	// Command ID (of acknowledged command).
	Command MAV_CMD `mavenum:"uint16"`
	// Result of command.
	Result MAV_RESULT `mavenum:"uint8"`
	// WIP: Also used as result_param1, it can be set with a enum containing the errors reasons of why the command was denied or the progress percentage or 255 if unknown the progress when result is MAV_RESULT_IN_PROGRESS.
	Progress uint8 `mavext:"true"`
	// WIP: Additional parameter of the result, example: which parameter of MAV_CMD_NAV_WAYPOINT caused it to be denied.
	ResultParam2 int32 `mavext:"true"`
	// WIP: System which requested the command to be executed
	TargetSystem uint8 `mavext:"true"`
	// WIP: Component which requested the command to be executed
	TargetComponent uint8 `mavext:"true"`
}

func (m *MessageCommandAck) GetId() uint32 {
	return 77
}

func (m *MessageCommandAck) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Setpoint in roll, pitch, yaw and thrust from the operator
type MessageManualSetpoint struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Desired roll rate
	Roll float32
	// Desired pitch rate
	Pitch float32
	// Desired yaw rate
	Yaw float32
	// Collective thrust, normalized to 0 .. 1
	Thrust float32
	// Flight mode switch position, 0.. 255
	ModeSwitch uint8
	// Override mode switch position, 0.. 255
	ManualOverrideSwitch uint8
}

func (m *MessageManualSetpoint) GetId() uint32 {
	return 81
}

func (m *MessageManualSetpoint) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Sets a desired vehicle attitude. Used by an external controller to command the vehicle (manual controller or other system).
type MessageSetAttitudeTarget struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Mappings: If any of these bits are set, the corresponding input should be ignored: bit 1: body roll rate, bit 2: body pitch rate, bit 3: body yaw rate. bit 4-bit 6: reserved, bit 7: throttle, bit 8: attitude
	TypeMask uint8
	// Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	Q [4]float32
	// Body roll rate
	BodyRollRate float32
	// Body pitch rate
	BodyPitchRate float32
	// Body yaw rate
	BodyYawRate float32
	// Collective thrust, normalized to 0 .. 1 (-1 .. 1 for vehicles capable of reverse trust)
	Thrust float32
}

func (m *MessageSetAttitudeTarget) GetId() uint32 {
	return 82
}

func (m *MessageSetAttitudeTarget) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Reports the current commanded attitude of the vehicle as specified by the autopilot. This should match the commands sent in a SET_ATTITUDE_TARGET message if the vehicle is being controlled this way.
type MessageAttitudeTarget struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Mappings: If any of these bits are set, the corresponding input should be ignored: bit 1: body roll rate, bit 2: body pitch rate, bit 3: body yaw rate. bit 4-bit 7: reserved, bit 8: attitude
	TypeMask uint8
	// Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	Q [4]float32
	// Body roll rate
	BodyRollRate float32
	// Body pitch rate
	BodyPitchRate float32
	// Body yaw rate
	BodyYawRate float32
	// Collective thrust, normalized to 0 .. 1 (-1 .. 1 for vehicles capable of reverse trust)
	Thrust float32
}

func (m *MessageAttitudeTarget) GetId() uint32 {
	return 83
}

func (m *MessageAttitudeTarget) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Sets a desired vehicle position in a local north-east-down coordinate frame. Used by an external controller to command the vehicle (manual controller or other system).
type MessageSetPositionTargetLocalNed struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Valid options are: MAV_FRAME_LOCAL_NED = 1, MAV_FRAME_LOCAL_OFFSET_NED = 7, MAV_FRAME_BODY_NED = 8, MAV_FRAME_BODY_OFFSET_NED = 9
	CoordinateFrame MAV_FRAME `mavenum:"uint8"`
	// Bitmap to indicate which dimensions should be ignored by the vehicle.
	TypeMask POSITION_TARGET_TYPEMASK `mavenum:"uint16"`
	// X Position in NED frame
	X float32
	// Y Position in NED frame
	Y float32
	// Z Position in NED frame (note, altitude is negative in NED)
	Z float32
	// X velocity in NED frame
	Vx float32
	// Y velocity in NED frame
	Vy float32
	// Z velocity in NED frame
	Vz float32
	// X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afx float32
	// Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy float32
	// Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz float32
	// yaw setpoint
	Yaw float32
	// yaw rate setpoint
	YawRate float32
}

func (m *MessageSetPositionTargetLocalNed) GetId() uint32 {
	return 84
}

func (m *MessageSetPositionTargetLocalNed) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Reports the current commanded vehicle position, velocity, and acceleration as specified by the autopilot. This should match the commands sent in SET_POSITION_TARGET_LOCAL_NED if the vehicle is being controlled this way.
type MessagePositionTargetLocalNed struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Valid options are: MAV_FRAME_LOCAL_NED = 1, MAV_FRAME_LOCAL_OFFSET_NED = 7, MAV_FRAME_BODY_NED = 8, MAV_FRAME_BODY_OFFSET_NED = 9
	CoordinateFrame MAV_FRAME `mavenum:"uint8"`
	// Bitmap to indicate which dimensions should be ignored by the vehicle.
	TypeMask POSITION_TARGET_TYPEMASK `mavenum:"uint16"`
	// X Position in NED frame
	X float32
	// Y Position in NED frame
	Y float32
	// Z Position in NED frame (note, altitude is negative in NED)
	Z float32
	// X velocity in NED frame
	Vx float32
	// Y velocity in NED frame
	Vy float32
	// Z velocity in NED frame
	Vz float32
	// X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afx float32
	// Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy float32
	// Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz float32
	// yaw setpoint
	Yaw float32
	// yaw rate setpoint
	YawRate float32
}

func (m *MessagePositionTargetLocalNed) GetId() uint32 {
	return 85
}

func (m *MessagePositionTargetLocalNed) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Sets a desired vehicle position, velocity, and/or acceleration in a global coordinate system (WGS84). Used by an external controller to command the vehicle (manual controller or other system).
type MessageSetPositionTargetGlobalInt struct {
	// Timestamp (time since system boot). The rationale for the timestamp in the setpoint is to allow the system to compensate for the transport delay of the setpoint. This allows the system to compensate processing latency.
	TimeBootMs uint32
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Valid options are: MAV_FRAME_GLOBAL_INT = 5, MAV_FRAME_GLOBAL_RELATIVE_ALT_INT = 6, MAV_FRAME_GLOBAL_TERRAIN_ALT_INT = 11
	CoordinateFrame MAV_FRAME `mavenum:"uint8"`
	// Bitmap to indicate which dimensions should be ignored by the vehicle.
	TypeMask POSITION_TARGET_TYPEMASK `mavenum:"uint16"`
	// X Position in WGS84 frame
	LatInt int32
	// Y Position in WGS84 frame
	LonInt int32
	// Altitude (MSL, Relative to home, or AGL - depending on frame)
	Alt float32
	// X velocity in NED frame
	Vx float32
	// Y velocity in NED frame
	Vy float32
	// Z velocity in NED frame
	Vz float32
	// X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afx float32
	// Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy float32
	// Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz float32
	// yaw setpoint
	Yaw float32
	// yaw rate setpoint
	YawRate float32
}

func (m *MessageSetPositionTargetGlobalInt) GetId() uint32 {
	return 86
}

func (m *MessageSetPositionTargetGlobalInt) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Reports the current commanded vehicle position, velocity, and acceleration as specified by the autopilot. This should match the commands sent in SET_POSITION_TARGET_GLOBAL_INT if the vehicle is being controlled this way.
type MessagePositionTargetGlobalInt struct {
	// Timestamp (time since system boot). The rationale for the timestamp in the setpoint is to allow the system to compensate for the transport delay of the setpoint. This allows the system to compensate processing latency.
	TimeBootMs uint32
	// Valid options are: MAV_FRAME_GLOBAL_INT = 5, MAV_FRAME_GLOBAL_RELATIVE_ALT_INT = 6, MAV_FRAME_GLOBAL_TERRAIN_ALT_INT = 11
	CoordinateFrame MAV_FRAME `mavenum:"uint8"`
	// Bitmap to indicate which dimensions should be ignored by the vehicle.
	TypeMask POSITION_TARGET_TYPEMASK `mavenum:"uint16"`
	// X Position in WGS84 frame
	LatInt int32
	// Y Position in WGS84 frame
	LonInt int32
	// Altitude (MSL, AGL or relative to home altitude, depending on frame)
	Alt float32
	// X velocity in NED frame
	Vx float32
	// Y velocity in NED frame
	Vy float32
	// Z velocity in NED frame
	Vz float32
	// X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afx float32
	// Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy float32
	// Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz float32
	// yaw setpoint
	Yaw float32
	// yaw rate setpoint
	YawRate float32
}

func (m *MessagePositionTargetGlobalInt) GetId() uint32 {
	return 87
}

func (m *MessagePositionTargetGlobalInt) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The offset in X, Y, Z and yaw between the LOCAL_POSITION_NED messages of MAV X and the global coordinate frame in NED coordinates. Coordinate frame is right-handed, Z-axis down (aeronautical frame, NED / north-east-down convention)
type MessageLocalPositionNedSystemGlobalOffset struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// X Position
	X float32
	// Y Position
	Y float32
	// Z Position
	Z float32
	// Roll
	Roll float32
	// Pitch
	Pitch float32
	// Yaw
	Yaw float32
}

func (m *MessageLocalPositionNedSystemGlobalOffset) GetId() uint32 {
	return 89
}

func (m *MessageLocalPositionNedSystemGlobalOffset) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Sent from simulation to autopilot. This packet is useful for high throughput applications such as hardware in the loop simulations.
type MessageHilState struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Roll angle
	Roll float32
	// Pitch angle
	Pitch float32
	// Yaw angle
	Yaw float32
	// Body frame roll / phi angular speed
	Rollspeed float32
	// Body frame pitch / theta angular speed
	Pitchspeed float32
	// Body frame yaw / psi angular speed
	Yawspeed float32
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// Altitude
	Alt int32
	// Ground X Speed (Latitude)
	Vx int16
	// Ground Y Speed (Longitude)
	Vy int16
	// Ground Z Speed (Altitude)
	Vz int16
	// X acceleration
	Xacc int16
	// Y acceleration
	Yacc int16
	// Z acceleration
	Zacc int16
}

func (m *MessageHilState) GetId() uint32 {
	return 90
}

func (m *MessageHilState) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Sent from autopilot to simulation. Hardware in the loop control outputs
type MessageHilControls struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Control output -1 .. 1
	RollAilerons float32
	// Control output -1 .. 1
	PitchElevator float32
	// Control output -1 .. 1
	YawRudder float32
	// Throttle 0 .. 1
	Throttle float32
	// Aux 1, -1 .. 1
	Aux1 float32
	// Aux 2, -1 .. 1
	Aux2 float32
	// Aux 3, -1 .. 1
	Aux3 float32
	// Aux 4, -1 .. 1
	Aux4 float32
	// System mode.
	Mode MAV_MODE `mavenum:"uint8"`
	// Navigation mode (MAV_NAV_MODE)
	NavMode uint8
}

func (m *MessageHilControls) GetId() uint32 {
	return 91
}

func (m *MessageHilControls) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Sent from simulation to autopilot. The RAW values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. Individual receivers/transmitters might violate this specification.
type MessageHilRcInputsRaw struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// RC channel 1 value
	Chan1Raw uint16
	// RC channel 2 value
	Chan2Raw uint16
	// RC channel 3 value
	Chan3Raw uint16
	// RC channel 4 value
	Chan4Raw uint16
	// RC channel 5 value
	Chan5Raw uint16
	// RC channel 6 value
	Chan6Raw uint16
	// RC channel 7 value
	Chan7Raw uint16
	// RC channel 8 value
	Chan8Raw uint16
	// RC channel 9 value
	Chan9Raw uint16
	// RC channel 10 value
	Chan10Raw uint16
	// RC channel 11 value
	Chan11Raw uint16
	// RC channel 12 value
	Chan12Raw uint16
	// Receive signal strength indicator in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Rssi uint8
}

func (m *MessageHilRcInputsRaw) GetId() uint32 {
	return 92
}

func (m *MessageHilRcInputsRaw) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Sent from autopilot to simulation. Hardware in the loop control outputs (replacement for HIL_CONTROLS)
type MessageHilActuatorControls struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Control outputs -1 .. 1. Channel assignment depends on the simulated hardware.
	Controls [16]float32
	// System mode. Includes arming state.
	Mode MAV_MODE_FLAG `mavenum:"uint8"`
	// Flags as bitfield, 1: indicate simulation using lockstep.
	Flags uint64
}

func (m *MessageHilActuatorControls) GetId() uint32 {
	return 93
}

func (m *MessageHilActuatorControls) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Optical flow from a flow sensor (e.g. optical mouse sensor)
type MessageOpticalFlow struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Sensor ID
	SensorId uint8
	// Flow in x-sensor direction
	FlowX int16
	// Flow in y-sensor direction
	FlowY int16
	// Flow in x-sensor direction, angular-speed compensated
	FlowCompMX float32
	// Flow in y-sensor direction, angular-speed compensated
	FlowCompMY float32
	// Optical flow quality / confidence. 0: bad, 255: maximum quality
	Quality uint8
	// Ground distance. Positive value: distance known. Negative value: Unknown distance
	GroundDistance float32
	// Flow rate about X axis
	FlowRateX float32 `mavext:"true"`
	// Flow rate about Y axis
	FlowRateY float32 `mavext:"true"`
}

func (m *MessageOpticalFlow) GetId() uint32 {
	return 100
}

func (m *MessageOpticalFlow) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Global position/attitude estimate from a vision source.
type MessageGlobalVisionPositionEstimate struct {
	// Timestamp (UNIX time or since system boot)
	Usec uint64
	// Global X position
	X float32
	// Global Y position
	Y float32
	// Global Z position
	Z float32
	// Roll angle
	Roll float32
	// Pitch angle
	Pitch float32
	// Yaw angle
	Yaw float32
	// Row-major representation of pose 6x6 cross-covariance matrix upper right triangle (states: x_global, y_global, z_global, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [21]float32 `mavext:"true"`
	// Estimate reset counter. This should be incremented when the estimate resets in any of the dimensions (position, velocity, attitude, angular speed). This is designed to be used when e.g an external SLAM system detects a loop-closure and the estimate jumps.
	ResetCounter uint8 `mavext:"true"`
}

func (m *MessageGlobalVisionPositionEstimate) GetId() uint32 {
	return 101
}

func (m *MessageGlobalVisionPositionEstimate) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Local position/attitude estimate from a vision source.
type MessageVisionPositionEstimate struct {
	// Timestamp (UNIX time or time since system boot)
	Usec uint64
	// Local X position
	X float32
	// Local Y position
	Y float32
	// Local Z position
	Z float32
	// Roll angle
	Roll float32
	// Pitch angle
	Pitch float32
	// Yaw angle
	Yaw float32
	// Row-major representation of pose 6x6 cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [21]float32 `mavext:"true"`
	// Estimate reset counter. This should be incremented when the estimate resets in any of the dimensions (position, velocity, attitude, angular speed). This is designed to be used when e.g an external SLAM system detects a loop-closure and the estimate jumps.
	ResetCounter uint8 `mavext:"true"`
}

func (m *MessageVisionPositionEstimate) GetId() uint32 {
	return 102
}

func (m *MessageVisionPositionEstimate) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Speed estimate from a vision source.
type MessageVisionSpeedEstimate struct {
	// Timestamp (UNIX time or time since system boot)
	Usec uint64
	// Global X speed
	X float32
	// Global Y speed
	Y float32
	// Global Z speed
	Z float32
	// Row-major representation of 3x3 linear velocity covariance matrix (states: vx, vy, vz; 1st three entries - 1st row, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [9]float32 `mavext:"true"`
	// Estimate reset counter. This should be incremented when the estimate resets in any of the dimensions (position, velocity, attitude, angular speed). This is designed to be used when e.g an external SLAM system detects a loop-closure and the estimate jumps.
	ResetCounter uint8 `mavext:"true"`
}

func (m *MessageVisionSpeedEstimate) GetId() uint32 {
	return 103
}

func (m *MessageVisionSpeedEstimate) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Global position estimate from a Vicon motion system source.
type MessageViconPositionEstimate struct {
	// Timestamp (UNIX time or time since system boot)
	Usec uint64
	// Global X position
	X float32
	// Global Y position
	Y float32
	// Global Z position
	Z float32
	// Roll angle
	Roll float32
	// Pitch angle
	Pitch float32
	// Yaw angle
	Yaw float32
	// Row-major representation of 6x6 pose cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [21]float32 `mavext:"true"`
}

func (m *MessageViconPositionEstimate) GetId() uint32 {
	return 104
}

func (m *MessageViconPositionEstimate) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The IMU readings in SI units in NED body frame
type MessageHighresImu struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// X acceleration
	Xacc float32
	// Y acceleration
	Yacc float32
	// Z acceleration
	Zacc float32
	// Angular speed around X axis
	Xgyro float32
	// Angular speed around Y axis
	Ygyro float32
	// Angular speed around Z axis
	Zgyro float32
	// X Magnetic field
	Xmag float32
	// Y Magnetic field
	Ymag float32
	// Z Magnetic field
	Zmag float32
	// Absolute pressure
	AbsPressure float32
	// Differential pressure
	DiffPressure float32
	// Altitude calculated from pressure
	PressureAlt float32
	// Temperature
	Temperature float32
	// Bitmap for fields that have updated since last message, bit 0 = xacc, bit 12: temperature
	FieldsUpdated uint16
	// Id. Ids are numbered from 0 and map to IMUs numbered from 1 (e.g. IMU1 will have a message with id=0)
	Id uint8 `mavext:"true"`
}

func (m *MessageHighresImu) GetId() uint32 {
	return 105
}

func (m *MessageHighresImu) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Optical flow from an angular rate flow sensor (e.g. PX4FLOW or mouse sensor)
type MessageOpticalFlowRad struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Sensor ID
	SensorId uint8
	// Integration time. Divide integrated_x and integrated_y by the integration time to obtain average flow. The integration time also indicates the.
	IntegrationTimeUs uint32
	// Flow around X axis (Sensor RH rotation about the X axis induces a positive flow. Sensor linear motion along the positive Y axis induces a negative flow.)
	IntegratedX float32
	// Flow around Y axis (Sensor RH rotation about the Y axis induces a positive flow. Sensor linear motion along the positive X axis induces a positive flow.)
	IntegratedY float32
	// RH rotation around X axis
	IntegratedXgyro float32
	// RH rotation around Y axis
	IntegratedYgyro float32
	// RH rotation around Z axis
	IntegratedZgyro float32
	// Temperature
	Temperature int16
	// Optical flow quality / confidence. 0: no valid flow, 255: maximum quality
	Quality uint8
	// Time since the distance was sampled.
	TimeDeltaDistanceUs uint32
	// Distance to the center of the flow field. Positive value (including zero): distance known. Negative value: Unknown distance.
	Distance float32
}

func (m *MessageOpticalFlowRad) GetId() uint32 {
	return 106
}

func (m *MessageOpticalFlowRad) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The IMU readings in SI units in NED body frame
type MessageHilSensor struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// X acceleration
	Xacc float32
	// Y acceleration
	Yacc float32
	// Z acceleration
	Zacc float32
	// Angular speed around X axis in body frame
	Xgyro float32
	// Angular speed around Y axis in body frame
	Ygyro float32
	// Angular speed around Z axis in body frame
	Zgyro float32
	// X Magnetic field
	Xmag float32
	// Y Magnetic field
	Ymag float32
	// Z Magnetic field
	Zmag float32
	// Absolute pressure
	AbsPressure float32
	// Differential pressure (airspeed)
	DiffPressure float32
	// Altitude calculated from pressure
	PressureAlt float32
	// Temperature
	Temperature float32
	// Bitmap for fields that have updated since last message, bit 0 = xacc, bit 12: temperature, bit 31: full reset of attitude/position/velocities/etc was performed in sim.
	FieldsUpdated uint32
}

func (m *MessageHilSensor) GetId() uint32 {
	return 107
}

func (m *MessageHilSensor) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Status of simulation environment, if used
type MessageSimState struct {
	// True attitude quaternion component 1, w (1 in null-rotation)
	Q1 float32
	// True attitude quaternion component 2, x (0 in null-rotation)
	Q2 float32
	// True attitude quaternion component 3, y (0 in null-rotation)
	Q3 float32
	// True attitude quaternion component 4, z (0 in null-rotation)
	Q4 float32
	// Attitude roll expressed as Euler angles, not recommended except for human-readable outputs
	Roll float32
	// Attitude pitch expressed as Euler angles, not recommended except for human-readable outputs
	Pitch float32
	// Attitude yaw expressed as Euler angles, not recommended except for human-readable outputs
	Yaw float32
	// X acceleration
	Xacc float32
	// Y acceleration
	Yacc float32
	// Z acceleration
	Zacc float32
	// Angular speed around X axis
	Xgyro float32
	// Angular speed around Y axis
	Ygyro float32
	// Angular speed around Z axis
	Zgyro float32
	// Latitude
	Lat float32
	// Longitude
	Lon float32
	// Altitude
	Alt float32
	// Horizontal position standard deviation
	StdDevHorz float32
	// Vertical position standard deviation
	StdDevVert float32
	// True velocity in north direction in earth-fixed NED frame
	Vn float32
	// True velocity in east direction in earth-fixed NED frame
	Ve float32
	// True velocity in down direction in earth-fixed NED frame
	Vd float32
}

func (m *MessageSimState) GetId() uint32 {
	return 108
}

func (m *MessageSimState) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Status generated by radio and injected into MAVLink stream.
type MessageRadioStatus struct {
	// Local (message sender) recieved signal strength indication in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Rssi uint8
	// Remote (message receiver) signal strength indication in device-dependent units/scale. Values: [0-254], 255: invalid/unknown.
	Remrssi uint8
	// Remaining free transmitter buffer space.
	Txbuf uint8
	// Local background noise level. These are device dependent RSSI values (scale as approx 2x dB on SiK radios). Values: [0-254], 255: invalid/unknown.
	Noise uint8
	// Remote background noise level. These are device dependent RSSI values (scale as approx 2x dB on SiK radios). Values: [0-254], 255: invalid/unknown.
	Remnoise uint8
	// Count of radio packet receive errors (since boot).
	Rxerrors uint16
	// Count of error corrected radio packets (since boot).
	Fixed uint16
}

func (m *MessageRadioStatus) GetId() uint32 {
	return 109
}

func (m *MessageRadioStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// File transfer message
type MessageFileTransferProtocol struct {
	// Network ID (0 for broadcast)
	TargetNetwork uint8
	// System ID (0 for broadcast)
	TargetSystem uint8
	// Component ID (0 for broadcast)
	TargetComponent uint8
	// Variable length payload. The length is defined by the remaining message length when subtracting the header and other fields.  The entire content of this block is opaque unless you understand any the encoding message_type.  The particular encoding used can be extension specific and might not always be documented as part of the mavlink specification.
	Payload [251]uint8
}

func (m *MessageFileTransferProtocol) GetId() uint32 {
	return 110
}

func (m *MessageFileTransferProtocol) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Time synchronization message.
type MessageTimesync struct {
	// Time sync timestamp 1
	Tc1 int64
	// Time sync timestamp 2
	Ts1 int64
}

func (m *MessageTimesync) GetId() uint32 {
	return 111
}

func (m *MessageTimesync) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Camera-IMU triggering and synchronisation message.
type MessageCameraTrigger struct {
	// Timestamp for image frame (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Image frame sequence
	Seq uint32
}

func (m *MessageCameraTrigger) GetId() uint32 {
	return 112
}

func (m *MessageCameraTrigger) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The global position, as returned by the Global Positioning System (GPS). This is                 NOT the global position estimate of the sytem, but rather a RAW sensor value. See message GLOBAL_POSITION for the global position estimate.
type MessageHilGps struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// 0-1: no fix, 2: 2D fix, 3: 3D fix. Some applications will not use the value of this field unless it is at least two, so always correctly fill in the fix.
	FixType uint8
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (MSL). Positive for up.
	Alt int32
	// GPS HDOP horizontal dilution of position. If unknown, set to: 65535
	Eph uint16
	// GPS VDOP vertical dilution of position. If unknown, set to: 65535
	Epv uint16
	// GPS ground speed. If unknown, set to: 65535
	Vel uint16
	// GPS velocity in north direction in earth-fixed NED frame
	Vn int16
	// GPS velocity in east direction in earth-fixed NED frame
	Ve int16
	// GPS velocity in down direction in earth-fixed NED frame
	Vd int16
	// Course over ground (NOT heading, but direction of movement), 0.0..359.99 degrees. If unknown, set to: 65535
	Cog uint16
	// Number of satellites visible. If unknown, set to 255
	SatellitesVisible uint8
}

func (m *MessageHilGps) GetId() uint32 {
	return 113
}

func (m *MessageHilGps) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Simulated optical flow from a flow sensor (e.g. PX4FLOW or optical mouse sensor)
type MessageHilOpticalFlow struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Sensor ID
	SensorId uint8
	// Integration time. Divide integrated_x and integrated_y by the integration time to obtain average flow. The integration time also indicates the.
	IntegrationTimeUs uint32
	// Flow in radians around X axis (Sensor RH rotation about the X axis induces a positive flow. Sensor linear motion along the positive Y axis induces a negative flow.)
	IntegratedX float32
	// Flow in radians around Y axis (Sensor RH rotation about the Y axis induces a positive flow. Sensor linear motion along the positive X axis induces a positive flow.)
	IntegratedY float32
	// RH rotation around X axis
	IntegratedXgyro float32
	// RH rotation around Y axis
	IntegratedYgyro float32
	// RH rotation around Z axis
	IntegratedZgyro float32
	// Temperature
	Temperature int16
	// Optical flow quality / confidence. 0: no valid flow, 255: maximum quality
	Quality uint8
	// Time since the distance was sampled.
	TimeDeltaDistanceUs uint32
	// Distance to the center of the flow field. Positive value (including zero): distance known. Negative value: Unknown distance.
	Distance float32
}

func (m *MessageHilOpticalFlow) GetId() uint32 {
	return 114
}

func (m *MessageHilOpticalFlow) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Sent from simulation to autopilot, avoids in contrast to HIL_STATE singularities. This packet is useful for high throughput applications such as hardware in the loop simulations.
type MessageHilStateQuaternion struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Vehicle attitude expressed as normalized quaternion in w, x, y, z order (with 1 0 0 0 being the null-rotation)
	AttitudeQuaternion [4]float32
	// Body frame roll / phi angular speed
	Rollspeed float32
	// Body frame pitch / theta angular speed
	Pitchspeed float32
	// Body frame yaw / psi angular speed
	Yawspeed float32
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// Altitude
	Alt int32
	// Ground X Speed (Latitude)
	Vx int16
	// Ground Y Speed (Longitude)
	Vy int16
	// Ground Z Speed (Altitude)
	Vz int16
	// Indicated airspeed
	IndAirspeed uint16
	// True airspeed
	TrueAirspeed uint16
	// X acceleration
	Xacc int16
	// Y acceleration
	Yacc int16
	// Z acceleration
	Zacc int16
}

func (m *MessageHilStateQuaternion) GetId() uint32 {
	return 115
}

func (m *MessageHilStateQuaternion) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The RAW IMU readings for secondary 9DOF sensor setup. This message should contain the scaled values to the described units
type MessageScaledImu2 struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// X acceleration
	Xacc int16
	// Y acceleration
	Yacc int16
	// Z acceleration
	Zacc int16
	// Angular speed around X axis
	Xgyro int16
	// Angular speed around Y axis
	Ygyro int16
	// Angular speed around Z axis
	Zgyro int16
	// X Magnetic field
	Xmag int16
	// Y Magnetic field
	Ymag int16
	// Z Magnetic field
	Zmag int16
	// Temperature, 0: IMU does not provide temperature values. If the IMU is at 0C it must send 1 (0.01C).
	Temperature int16 `mavext:"true"`
}

func (m *MessageScaledImu2) GetId() uint32 {
	return 116
}

func (m *MessageScaledImu2) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request a list of available logs. On some systems calling this may stop on-board logging until LOG_REQUEST_END is called.
type MessageLogRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// First log id (0 for first available)
	Start uint16
	// Last log id (0xffff for last available)
	End uint16
}

func (m *MessageLogRequestList) GetId() uint32 {
	return 117
}

func (m *MessageLogRequestList) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Reply to LOG_REQUEST_LIST
type MessageLogEntry struct {
	// Log id
	Id uint16
	// Total number of logs
	NumLogs uint16
	// High log number
	LastLogNum uint16
	// UTC timestamp of log since 1970, or 0 if not available
	TimeUtc uint32
	// Size of the log (may be approximate)
	Size uint32
}

func (m *MessageLogEntry) GetId() uint32 {
	return 118
}

func (m *MessageLogEntry) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request a chunk of a log
type MessageLogRequestData struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Log id (from LOG_ENTRY reply)
	Id uint16
	// Offset into the log
	Ofs uint32
	// Number of bytes
	Count uint32
}

func (m *MessageLogRequestData) GetId() uint32 {
	return 119
}

func (m *MessageLogRequestData) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Reply to LOG_REQUEST_DATA
type MessageLogData struct {
	// Log id (from LOG_ENTRY reply)
	Id uint16
	// Offset into the log
	Ofs uint32
	// Number of bytes (zero for end of log)
	Count uint8
	// log data
	Data [90]uint8
}

func (m *MessageLogData) GetId() uint32 {
	return 120
}

func (m *MessageLogData) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Erase all logs
type MessageLogErase struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

func (m *MessageLogErase) GetId() uint32 {
	return 121
}

func (m *MessageLogErase) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Stop log transfer and resume normal logging
type MessageLogRequestEnd struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

func (m *MessageLogRequestEnd) GetId() uint32 {
	return 122
}

func (m *MessageLogRequestEnd) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Data for injecting into the onboard GPS (used for DGPS)
type MessageGpsInjectData struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Data length
	Len uint8
	// Raw data (110 is enough for 12 satellites of RTCMv2)
	Data [110]uint8
}

func (m *MessageGpsInjectData) GetId() uint32 {
	return 123
}

func (m *MessageGpsInjectData) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Second GPS data.
type MessageGps2Raw struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// GPS fix type.
	FixType GPS_FIX_TYPE `mavenum:"uint8"`
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (MSL). Positive for up.
	Alt int32
	// GPS HDOP horizontal dilution of position. If unknown, set to: UINT16_MAX
	Eph uint16
	// GPS VDOP vertical dilution of position. If unknown, set to: UINT16_MAX
	Epv uint16
	// GPS ground speed. If unknown, set to: UINT16_MAX
	Vel uint16
	// Course over ground (NOT heading, but direction of movement): 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	Cog uint16
	// Number of satellites visible. If unknown, set to 255
	SatellitesVisible uint8
	// Number of DGPS satellites
	DgpsNumch uint8
	// Age of DGPS info
	DgpsAge uint32
	// Yaw in earth frame from north. Use 0 if this GPS does not provide yaw. Use 65535 if this GPS is configured to provide yaw and is currently unable to provide it. Use 36000 for north.
	Yaw uint16 `mavext:"true"`
}

func (m *MessageGps2Raw) GetId() uint32 {
	return 124
}

func (m *MessageGps2Raw) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Power supply status
type MessagePowerStatus struct {
	// 5V rail voltage.
	Vcc uint16 `mavname:"Vcc"`
	// Servo rail voltage.
	Vservo uint16 `mavname:"Vservo"`
	// Bitmap of power supply status flags.
	Flags MAV_POWER_STATUS `mavenum:"uint16"`
}

func (m *MessagePowerStatus) GetId() uint32 {
	return 125
}

func (m *MessagePowerStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Control a serial port. This can be used for raw access to an onboard serial peripheral such as a GPS or telemetry radio. It is designed to make it possible to update the devices firmware via MAVLink messages or change the devices settings. A message with zero bytes can be used to change just the baudrate.
type MessageSerialControl struct {
	// Serial control device type.
	Device SERIAL_CONTROL_DEV `mavenum:"uint8"`
	// Bitmap of serial control flags.
	Flags SERIAL_CONTROL_FLAG `mavenum:"uint8"`
	// Timeout for reply data
	Timeout uint16
	// Baudrate of transfer. Zero means no change.
	Baudrate uint32
	// how many bytes in this transfer
	Count uint8
	// serial data
	Data [70]uint8
}

func (m *MessageSerialControl) GetId() uint32 {
	return 126
}

func (m *MessageSerialControl) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// RTK GPS data. Gives information on the relative baseline calculation the GPS is reporting
type MessageGpsRtk struct {
	// Time since boot of last baseline message received.
	TimeLastBaselineMs uint32
	// Identification of connected RTK receiver.
	RtkReceiverId uint8
	// GPS Week Number of last baseline
	Wn uint16
	// GPS Time of Week of last baseline
	Tow uint32
	// GPS-specific health report for RTK data.
	RtkHealth uint8
	// Rate of baseline messages being received by GPS
	RtkRate uint8
	// Current number of sats used for RTK calculation.
	Nsats uint8
	// Coordinate system of baseline
	BaselineCoordsType RTK_BASELINE_COORDINATE_SYSTEM `mavenum:"uint8"`
	// Current baseline in ECEF x or NED north component.
	BaselineAMm int32
	// Current baseline in ECEF y or NED east component.
	BaselineBMm int32
	// Current baseline in ECEF z or NED down component.
	BaselineCMm int32
	// Current estimate of baseline accuracy.
	Accuracy uint32
	// Current number of integer ambiguity hypotheses.
	IarNumHypotheses int32
}

func (m *MessageGpsRtk) GetId() uint32 {
	return 127
}

func (m *MessageGpsRtk) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// RTK GPS data. Gives information on the relative baseline calculation the GPS is reporting
type MessageGps2Rtk struct {
	// Time since boot of last baseline message received.
	TimeLastBaselineMs uint32
	// Identification of connected RTK receiver.
	RtkReceiverId uint8
	// GPS Week Number of last baseline
	Wn uint16
	// GPS Time of Week of last baseline
	Tow uint32
	// GPS-specific health report for RTK data.
	RtkHealth uint8
	// Rate of baseline messages being received by GPS
	RtkRate uint8
	// Current number of sats used for RTK calculation.
	Nsats uint8
	// Coordinate system of baseline
	BaselineCoordsType RTK_BASELINE_COORDINATE_SYSTEM `mavenum:"uint8"`
	// Current baseline in ECEF x or NED north component.
	BaselineAMm int32
	// Current baseline in ECEF y or NED east component.
	BaselineBMm int32
	// Current baseline in ECEF z or NED down component.
	BaselineCMm int32
	// Current estimate of baseline accuracy.
	Accuracy uint32
	// Current number of integer ambiguity hypotheses.
	IarNumHypotheses int32
}

func (m *MessageGps2Rtk) GetId() uint32 {
	return 128
}

func (m *MessageGps2Rtk) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The RAW IMU readings for 3rd 9DOF sensor setup. This message should contain the scaled values to the described units
type MessageScaledImu3 struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// X acceleration
	Xacc int16
	// Y acceleration
	Yacc int16
	// Z acceleration
	Zacc int16
	// Angular speed around X axis
	Xgyro int16
	// Angular speed around Y axis
	Ygyro int16
	// Angular speed around Z axis
	Zgyro int16
	// X Magnetic field
	Xmag int16
	// Y Magnetic field
	Ymag int16
	// Z Magnetic field
	Zmag int16
	// Temperature, 0: IMU does not provide temperature values. If the IMU is at 0C it must send 1 (0.01C).
	Temperature int16 `mavext:"true"`
}

func (m *MessageScaledImu3) GetId() uint32 {
	return 129
}

func (m *MessageScaledImu3) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Handshake message to initiate, control and stop image streaming when using the Image Transmission Protocol: https://mavlink.io/en/services/image_transmission.html.
type MessageDataTransmissionHandshake struct {
	// Type of requested/acknowledged data.
	Type MAVLINK_DATA_STREAM_TYPE `mavenum:"uint8"`
	// total data size (set on ACK only).
	Size uint32
	// Width of a matrix or image.
	Width uint16
	// Height of a matrix or image.
	Height uint16
	// Number of packets being sent (set on ACK only).
	Packets uint16
	// Payload size per packet (normally 253 byte, see DATA field size in message ENCAPSULATED_DATA) (set on ACK only).
	Payload uint8
	// JPEG quality. Values: [1-100].
	JpgQuality uint8
}

func (m *MessageDataTransmissionHandshake) GetId() uint32 {
	return 130
}

func (m *MessageDataTransmissionHandshake) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Data packet for images sent using the Image Transmission Protocol: https://mavlink.io/en/services/image_transmission.html.
type MessageEncapsulatedData struct {
	// sequence number (starting with 0 on every transmission)
	Seqnr uint16
	// image data bytes
	Data [253]uint8
}

func (m *MessageEncapsulatedData) GetId() uint32 {
	return 131
}

func (m *MessageEncapsulatedData) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Distance sensor information for an onboard rangefinder.
type MessageDistanceSensor struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Minimum distance the sensor can measure
	MinDistance uint16
	// Maximum distance the sensor can measure
	MaxDistance uint16
	// Current distance reading
	CurrentDistance uint16
	// Type of distance sensor.
	Type MAV_DISTANCE_SENSOR `mavenum:"uint8"`
	// Onboard ID of the sensor
	Id uint8
	// Direction the sensor faces. downward-facing: ROTATION_PITCH_270, upward-facing: ROTATION_PITCH_90, backward-facing: ROTATION_PITCH_180, forward-facing: ROTATION_NONE, left-facing: ROTATION_YAW_90, right-facing: ROTATION_YAW_270
	Orientation MAV_SENSOR_ORIENTATION `mavenum:"uint8"`
	// Measurement variance. Max standard deviation is 6cm. 255 if unknown.
	Covariance uint8
	// Horizontal Field of View (angle) where the distance measurement is valid and the field of view is known. Otherwise this is set to 0.
	HorizontalFov float32 `mavext:"true"`
	// Vertical Field of View (angle) where the distance measurement is valid and the field of view is known. Otherwise this is set to 0.
	VerticalFov float32 `mavext:"true"`
	// Quaternion of the sensor orientation in vehicle body frame (w, x, y, z order, zero-rotation is 1, 0, 0, 0). Zero-rotation is along the vehicle body x-axis. This field is required if the orientation is set to MAV_SENSOR_ROTATION_CUSTOM. Set it to 0 if invalid."
	Quaternion [4]float32 `mavext:"true"`
}

func (m *MessageDistanceSensor) GetId() uint32 {
	return 132
}

func (m *MessageDistanceSensor) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request for terrain data and terrain status
type MessageTerrainRequest struct {
	// Latitude of SW corner of first grid
	Lat int32
	// Longitude of SW corner of first grid
	Lon int32
	// Grid spacing
	GridSpacing uint16
	// Bitmask of requested 4x4 grids (row major 8x7 array of grids, 56 bits)
	Mask uint64
}

func (m *MessageTerrainRequest) GetId() uint32 {
	return 133
}

func (m *MessageTerrainRequest) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Terrain data sent from GCS. The lat/lon and grid_spacing must be the same as a lat/lon from a TERRAIN_REQUEST
type MessageTerrainData struct {
	// Latitude of SW corner of first grid
	Lat int32
	// Longitude of SW corner of first grid
	Lon int32
	// Grid spacing
	GridSpacing uint16
	// bit within the terrain request mask
	Gridbit uint8
	// Terrain data MSL
	Data [16]int16
}

func (m *MessageTerrainData) GetId() uint32 {
	return 134
}

func (m *MessageTerrainData) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request that the vehicle report terrain height at the given location. Used by GCS to check if vehicle has all terrain data needed for a mission.
type MessageTerrainCheck struct {
	// Latitude
	Lat int32
	// Longitude
	Lon int32
}

func (m *MessageTerrainCheck) GetId() uint32 {
	return 135
}

func (m *MessageTerrainCheck) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Response from a TERRAIN_CHECK request
type MessageTerrainReport struct {
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// grid spacing (zero if terrain at this location unavailable)
	Spacing uint16
	// Terrain height MSL
	TerrainHeight float32
	// Current vehicle height above lat/lon terrain height
	CurrentHeight float32
	// Number of 4x4 terrain blocks waiting to be received or read from disk
	Pending uint16
	// Number of 4x4 terrain blocks in memory
	Loaded uint16
}

func (m *MessageTerrainReport) GetId() uint32 {
	return 136
}

func (m *MessageTerrainReport) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Barometer readings for 2nd barometer
type MessageScaledPressure2 struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Absolute pressure
	PressAbs float32
	// Differential pressure
	PressDiff float32
	// Temperature measurement
	Temperature int16
}

func (m *MessageScaledPressure2) GetId() uint32 {
	return 137
}

func (m *MessageScaledPressure2) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Motion capture attitude and position
type MessageAttPosMocap struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	Q [4]float32
	// X position (NED)
	X float32
	// Y position (NED)
	Y float32
	// Z position (NED)
	Z float32
	// Row-major representation of a pose 6x6 cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [21]float32 `mavext:"true"`
}

func (m *MessageAttPosMocap) GetId() uint32 {
	return 138
}

func (m *MessageAttPosMocap) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Set the vehicle attitude and body angular rates.
type MessageSetActuatorControlTarget struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Actuator group. The "_mlx" indicates this is a multi-instance message and a MAVLink parser should use this field to difference between instances.
	GroupMlx uint8
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Actuator controls. Normed to -1..+1 where 0 is neutral position. Throttle for single rotation direction motors is 0..1, negative range for reverse direction. Standard mapping for attitude controls (group 0): (index 0-7): roll, pitch, yaw, throttle, flaps, spoilers, airbrakes, landing gear. Load a pass-through mixer to repurpose them as generic outputs.
	Controls [8]float32
}

func (m *MessageSetActuatorControlTarget) GetId() uint32 {
	return 139
}

func (m *MessageSetActuatorControlTarget) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Set the vehicle attitude and body angular rates.
type MessageActuatorControlTarget struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Actuator group. The "_mlx" indicates this is a multi-instance message and a MAVLink parser should use this field to difference between instances.
	GroupMlx uint8
	// Actuator controls. Normed to -1..+1 where 0 is neutral position. Throttle for single rotation direction motors is 0..1, negative range for reverse direction. Standard mapping for attitude controls (group 0): (index 0-7): roll, pitch, yaw, throttle, flaps, spoilers, airbrakes, landing gear. Load a pass-through mixer to repurpose them as generic outputs.
	Controls [8]float32
}

func (m *MessageActuatorControlTarget) GetId() uint32 {
	return 140
}

func (m *MessageActuatorControlTarget) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The current system altitude.
type MessageAltitude struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// This altitude measure is initialized on system boot and monotonic (it is never reset, but represents the local altitude change). The only guarantee on this field is that it will never be reset and is consistent within a flight. The recommended value for this field is the uncorrected barometric altitude at boot time. This altitude will also drift and vary between flights.
	AltitudeMonotonic float32
	// This altitude measure is strictly above mean sea level and might be non-monotonic (it might reset on events like GPS lock or when a new QNH value is set). It should be the altitude to which global altitude waypoints are compared to. Note that it is *not* the GPS altitude, however, most GPS modules already output MSL by default and not the WGS84 altitude.
	AltitudeAmsl float32
	// This is the local altitude in the local coordinate frame. It is not the altitude above home, but in reference to the coordinate origin (0, 0, 0). It is up-positive.
	AltitudeLocal float32
	// This is the altitude above the home position. It resets on each change of the current home position.
	AltitudeRelative float32
	// This is the altitude above terrain. It might be fed by a terrain database or an altimeter. Values smaller than -1000 should be interpreted as unknown.
	AltitudeTerrain float32
	// This is not the altitude, but the clear space below the system according to the fused clearance estimate. It generally should max out at the maximum range of e.g. the laser altimeter. It is generally a moving target. A negative value indicates no measurement available.
	BottomClearance float32
}

func (m *MessageAltitude) GetId() uint32 {
	return 141
}

func (m *MessageAltitude) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The autopilot is requesting a resource (file, binary, other type of data)
type MessageResourceRequest struct {
	// Request ID. This ID should be re-used when sending back URI contents
	RequestId uint8
	// The type of requested URI. 0 = a file via URL. 1 = a UAVCAN binary
	UriType uint8
	// The requested unique resource identifier (URI). It is not necessarily a straight domain name (depends on the URI type enum)
	Uri [120]uint8
	// The way the autopilot wants to receive the URI. 0 = MAVLink FTP. 1 = binary stream.
	TransferType uint8
	// The storage path the autopilot wants the URI to be stored in. Will only be valid if the transfer_type has a storage associated (e.g. MAVLink FTP).
	Storage [120]uint8
}

func (m *MessageResourceRequest) GetId() uint32 {
	return 142
}

func (m *MessageResourceRequest) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Barometer readings for 3rd barometer
type MessageScaledPressure3 struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Absolute pressure
	PressAbs float32
	// Differential pressure
	PressDiff float32
	// Temperature measurement
	Temperature int16
}

func (m *MessageScaledPressure3) GetId() uint32 {
	return 143
}

func (m *MessageScaledPressure3) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Current motion information from a designated system
type MessageFollowTarget struct {
	// Timestamp (time since system boot).
	Timestamp uint64
	// bit positions for tracker reporting capabilities (POS = 0, VEL = 1, ACCEL = 2, ATT + RATES = 3)
	EstCapabilities uint8
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (MSL)
	Alt float32
	// target velocity (0,0,0) for unknown
	Vel [3]float32
	// linear target acceleration (0,0,0) for unknown
	Acc [3]float32
	// (1 0 0 0 for unknown)
	AttitudeQ [4]float32
	// (0 0 0 for unknown)
	Rates [3]float32
	// eph epv
	PositionCov [3]float32
	// button states or switches of a tracker device
	CustomState uint64
}

func (m *MessageFollowTarget) GetId() uint32 {
	return 144
}

func (m *MessageFollowTarget) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The smoothed, monotonic system state used to feed the control loops of the system.
type MessageControlSystemState struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// X acceleration in body frame
	XAcc float32
	// Y acceleration in body frame
	YAcc float32
	// Z acceleration in body frame
	ZAcc float32
	// X velocity in body frame
	XVel float32
	// Y velocity in body frame
	YVel float32
	// Z velocity in body frame
	ZVel float32
	// X position in local frame
	XPos float32
	// Y position in local frame
	YPos float32
	// Z position in local frame
	ZPos float32
	// Airspeed, set to -1 if unknown
	Airspeed float32
	// Variance of body velocity estimate
	VelVariance [3]float32
	// Variance in local position
	PosVariance [3]float32
	// The attitude, represented as Quaternion
	Q [4]float32
	// Angular rate in roll axis
	RollRate float32
	// Angular rate in pitch axis
	PitchRate float32
	// Angular rate in yaw axis
	YawRate float32
}

func (m *MessageControlSystemState) GetId() uint32 {
	return 146
}

func (m *MessageControlSystemState) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Battery information. Updates GCS with flight controller battery status. Use SMART_BATTERY_* messages instead for smart batteries.
type MessageBatteryStatus struct {
	// Battery ID
	Id uint8
	// Function of the battery
	BatteryFunction MAV_BATTERY_FUNCTION `mavenum:"uint8"`
	// Type (chemistry) of the battery
	Type MAV_BATTERY_TYPE `mavenum:"uint8"`
	// Temperature of the battery. INT16_MAX for unknown temperature.
	Temperature int16
	// Battery voltage of cells. Cells above the valid cell count for this battery should have the UINT16_MAX value. If individual cell voltages are unknown or not measured for this battery, then the overall battery voltage should be filled in cell 0, with all others set to UINT16_MAX. If the voltage of the battery is greater than (UINT16_MAX - 1), then cell 0 should be set to (UINT16_MAX - 1), and cell 1 to the remaining voltage. This can be extended to multiple cells if the total voltage is greater than 2 * (UINT16_MAX - 1).
	Voltages [10]uint16
	// Battery current, -1: autopilot does not measure the current
	CurrentBattery int16
	// Consumed charge, -1: autopilot does not provide consumption estimate
	CurrentConsumed int32
	// Consumed energy, -1: autopilot does not provide energy consumption estimate
	EnergyConsumed int32
	// Remaining battery energy. Values: [0-100], -1: autopilot does not estimate the remaining battery.
	BatteryRemaining int8
	// Remaining battery time, 0: autopilot does not provide remaining battery time estimate
	TimeRemaining int32 `mavext:"true"`
	// State for extent of discharge, provided by autopilot for warning or external reactions
	ChargeState MAV_BATTERY_CHARGE_STATE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageBatteryStatus) GetId() uint32 {
	return 147
}

func (m *MessageBatteryStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Version and capability of autopilot software. This should be emitted in response to a MAV_CMD_REQUEST_AUTOPILOT_CAPABILITIES command.
type MessageAutopilotVersion struct {
	// Bitmap of capabilities
	Capabilities MAV_PROTOCOL_CAPABILITY `mavenum:"uint64"`
	// Firmware version number
	FlightSwVersion uint32
	// Middleware version number
	MiddlewareSwVersion uint32
	// Operating system version number
	OsSwVersion uint32
	// HW / board version (last 8 bytes should be silicon ID, if any)
	BoardVersion uint32
	// Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	FlightCustomVersion [8]uint8
	// Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	MiddlewareCustomVersion [8]uint8
	// Custom version field, commonly the first 8 bytes of the git hash. This is not an unique identifier, but should allow to identify the commit using the main version number even for very large code bases.
	OsCustomVersion [8]uint8
	// ID of the board vendor
	VendorId uint16
	// ID of the product
	ProductId uint16
	// UID if provided by hardware (see uid2)
	Uid uint64
	// UID if provided by hardware (supersedes the uid field. If this is non-zero, use this field, otherwise use uid)
	Uid2 [18]uint8 `mavext:"true"`
}

func (m *MessageAutopilotVersion) GetId() uint32 {
	return 148
}

func (m *MessageAutopilotVersion) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The location of a landing target. See: https://mavlink.io/en/services/landing_target.html
type MessageLandingTarget struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// The ID of the target if multiple targets are present
	TargetNum uint8
	// Coordinate frame used for following fields.
	Frame MAV_FRAME `mavenum:"uint8"`
	// X-axis angular offset of the target from the center of the image
	AngleX float32
	// Y-axis angular offset of the target from the center of the image
	AngleY float32
	// Distance to the target from the vehicle
	Distance float32
	// Size of target along x-axis
	SizeX float32
	// Size of target along y-axis
	SizeY float32
	// X Position of the landing target in MAV_FRAME
	X float32 `mavext:"true"`
	// Y Position of the landing target in MAV_FRAME
	Y float32 `mavext:"true"`
	// Z Position of the landing target in MAV_FRAME
	Z float32 `mavext:"true"`
	// Quaternion of landing target orientation (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	Q [4]float32 `mavext:"true"`
	// Type of landing target
	Type LANDING_TARGET_TYPE `mavenum:"uint8" mavext:"true"`
	// Boolean indicating whether the position fields (x, y, z, q, type) contain valid target position information (valid: 1, invalid: 0). Default is 0 (invalid).
	PositionValid uint8 `mavext:"true"`
}

func (m *MessageLandingTarget) GetId() uint32 {
	return 149
}

func (m *MessageLandingTarget) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Status of geo-fencing. Sent in extended status stream when fencing enabled.
type MessageFenceStatus struct {
	// Breach status (0 if currently inside fence, 1 if outside).
	BreachStatus uint8
	// Number of fence breaches.
	BreachCount uint16
	// Last breach type.
	BreachType FENCE_BREACH `mavenum:"uint8"`
	// Time (since boot) of last breach.
	BreachTime uint32
	// Active action to prevent fence breach
	BreachMitigation FENCE_MITIGATE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageFenceStatus) GetId() uint32 {
	return 162
}

func (m *MessageFenceStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Estimator status message including flags, innovation test ratios and estimated accuracies. The flags message is an integer bitmask containing information on which EKF outputs are valid. See the ESTIMATOR_STATUS_FLAGS enum definition for further information. The innovation test ratios show the magnitude of the sensor innovation divided by the innovation check threshold. Under normal operation the innovation test ratios should be below 0.5 with occasional values up to 1.0. Values greater than 1.0 should be rare under normal operation and indicate that a measurement has been rejected by the filter. The user should be notified if an innovation test ratio greater than 1.0 is recorded. Notifications for values in the range between 0.5 and 1.0 should be optional and controllable by the user.
type MessageEstimatorStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Bitmap indicating which EKF outputs are valid.
	Flags ESTIMATOR_STATUS_FLAGS `mavenum:"uint16"`
	// Velocity innovation test ratio
	VelRatio float32
	// Horizontal position innovation test ratio
	PosHorizRatio float32
	// Vertical position innovation test ratio
	PosVertRatio float32
	// Magnetometer innovation test ratio
	MagRatio float32
	// Height above terrain innovation test ratio
	HaglRatio float32
	// True airspeed innovation test ratio
	TasRatio float32
	// Horizontal position 1-STD accuracy relative to the EKF local origin
	PosHorizAccuracy float32
	// Vertical position 1-STD accuracy relative to the EKF local origin
	PosVertAccuracy float32
}

func (m *MessageEstimatorStatus) GetId() uint32 {
	return 230
}

func (m *MessageEstimatorStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Wind covariance estimate from vehicle.
type MessageWindCov struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Wind in X (NED) direction
	WindX float32
	// Wind in Y (NED) direction
	WindY float32
	// Wind in Z (NED) direction
	WindZ float32
	// Variability of the wind in XY. RMS of a 1 Hz lowpassed wind estimate.
	VarHoriz float32
	// Variability of the wind in Z. RMS of a 1 Hz lowpassed wind estimate.
	VarVert float32
	// Altitude (MSL) that this measurement was taken at
	WindAlt float32
	// Horizontal speed 1-STD accuracy
	HorizAccuracy float32
	// Vertical speed 1-STD accuracy
	VertAccuracy float32
}

func (m *MessageWindCov) GetId() uint32 {
	return 231
}

func (m *MessageWindCov) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// GPS sensor input message.  This is a raw sensor value sent by the GPS. This is NOT the global position estimate of the system.
type MessageGpsInput struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// ID of the GPS for multiple GPS inputs
	GpsId uint8
	// Bitmap indicating which GPS input flags fields to ignore.  All other fields must be provided.
	IgnoreFlags GPS_INPUT_IGNORE_FLAGS `mavenum:"uint16"`
	// GPS time (from start of GPS week)
	TimeWeekMs uint32
	// GPS week number
	TimeWeek uint16
	// 0-1: no fix, 2: 2D fix, 3: 3D fix. 4: 3D with DGPS. 5: 3D with RTK
	FixType uint8
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (MSL). Positive for up.
	Alt float32
	// GPS HDOP horizontal dilution of position
	Hdop float32
	// GPS VDOP vertical dilution of position
	Vdop float32
	// GPS velocity in north direction in earth-fixed NED frame
	Vn float32
	// GPS velocity in east direction in earth-fixed NED frame
	Ve float32
	// GPS velocity in down direction in earth-fixed NED frame
	Vd float32
	// GPS speed accuracy
	SpeedAccuracy float32
	// GPS horizontal accuracy
	HorizAccuracy float32
	// GPS vertical accuracy
	VertAccuracy float32
	// Number of satellites visible.
	SatellitesVisible uint8
	// Yaw of vehicle relative to Earth's North, zero means not available, use 36000 for north
	Yaw uint16 `mavext:"true"`
}

func (m *MessageGpsInput) GetId() uint32 {
	return 232
}

func (m *MessageGpsInput) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// RTCM message for injecting into the onboard GPS (used for DGPS)
type MessageGpsRtcmData struct {
	// LSB: 1 means message is fragmented, next 2 bits are the fragment ID, the remaining 5 bits are used for the sequence ID. Messages are only to be flushed to the GPS when the entire message has been reconstructed on the autopilot. The fragment ID specifies which order the fragments should be assembled into a buffer, while the sequence ID is used to detect a mismatch between different buffers. The buffer is considered fully reconstructed when either all 4 fragments are present, or all the fragments before the first fragment with a non full payload is received. This management is used to ensure that normal GPS operation doesn't corrupt RTCM data, and to recover from a unreliable transport delivery order.
	Flags uint8
	// data length
	Len uint8
	// RTCM message (may be fragmented)
	Data [180]uint8
}

func (m *MessageGpsRtcmData) GetId() uint32 {
	return 233
}

func (m *MessageGpsRtcmData) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Message appropriate for high latency connections like Iridium
type MessageHighLatency struct {
	// Bitmap of enabled system modes.
	BaseMode MAV_MODE_FLAG `mavenum:"uint8"`
	// A bitfield for use for autopilot-specific flags.
	CustomMode uint32
	// The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown.
	LandedState MAV_LANDED_STATE `mavenum:"uint8"`
	// roll
	Roll int16
	// pitch
	Pitch int16
	// heading
	Heading uint16
	// throttle (percentage)
	Throttle int8
	// heading setpoint
	HeadingSp int16
	// Latitude
	Latitude int32
	// Longitude
	Longitude int32
	// Altitude above mean sea level
	AltitudeAmsl int16
	// Altitude setpoint relative to the home position
	AltitudeSp int16
	// airspeed
	Airspeed uint8
	// airspeed setpoint
	AirspeedSp uint8
	// groundspeed
	Groundspeed uint8
	// climb rate
	ClimbRate int8
	// Number of satellites visible. If unknown, set to 255
	GpsNsat uint8
	// GPS Fix type.
	GpsFixType GPS_FIX_TYPE `mavenum:"uint8"`
	// Remaining battery (percentage)
	BatteryRemaining uint8
	// Autopilot temperature (degrees C)
	Temperature int8
	// Air temperature (degrees C) from airspeed sensor
	TemperatureAir int8
	// failsafe (each bit represents a failsafe where 0=ok, 1=failsafe active (bit0:RC, bit1:batt, bit2:GPS, bit3:GCS, bit4:fence)
	Failsafe uint8
	// current waypoint number
	WpNum uint8
	// distance to target
	WpDistance uint16
}

func (m *MessageHighLatency) GetId() uint32 {
	return 234
}

func (m *MessageHighLatency) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Message appropriate for high latency connections like Iridium (version 2)
type MessageHighLatency2 struct {
	// Timestamp (milliseconds since boot or Unix epoch)
	Timestamp uint32
	// Type of the MAV (quadrotor, helicopter, etc.)
	Type MAV_TYPE `mavenum:"uint8"`
	// Autopilot type / class. Use MAV_AUTOPILOT_INVALID for components that are not flight controllers.
	Autopilot MAV_AUTOPILOT `mavenum:"uint8"`
	// A bitfield for use for autopilot-specific flags (2 byte version).
	CustomMode uint16
	// Latitude
	Latitude int32
	// Longitude
	Longitude int32
	// Altitude above mean sea level
	Altitude int16
	// Altitude setpoint
	TargetAltitude int16
	// Heading
	Heading uint8
	// Heading setpoint
	TargetHeading uint8
	// Distance to target waypoint or position
	TargetDistance uint16
	// Throttle
	Throttle uint8
	// Airspeed
	Airspeed uint8
	// Airspeed setpoint
	AirspeedSp uint8
	// Groundspeed
	Groundspeed uint8
	// Windspeed
	Windspeed uint8
	// Wind heading
	WindHeading uint8
	// Maximum error horizontal position since last message
	Eph uint8
	// Maximum error vertical position since last message
	Epv uint8
	// Air temperature from airspeed sensor
	TemperatureAir int8
	// Maximum climb rate magnitude since last message
	ClimbRate int8
	// Battery level (-1 if field not provided).
	Battery int8
	// Current waypoint number
	WpNum uint16
	// Bitmap of failure flags.
	FailureFlags HL_FAILURE_FLAG `mavenum:"uint16"`
	// Field for custom payload.
	Custom0 int8
	// Field for custom payload.
	Custom1 int8
	// Field for custom payload.
	Custom2 int8
}

func (m *MessageHighLatency2) GetId() uint32 {
	return 235
}

func (m *MessageHighLatency2) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Vibration levels and accelerometer clipping
type MessageVibration struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Vibration levels on X-axis
	VibrationX float32
	// Vibration levels on Y-axis
	VibrationY float32
	// Vibration levels on Z-axis
	VibrationZ float32
	// first accelerometer clipping count
	Clipping_0 uint32
	// second accelerometer clipping count
	Clipping_1 uint32
	// third accelerometer clipping count
	Clipping_2 uint32
}

func (m *MessageVibration) GetId() uint32 {
	return 241
}

func (m *MessageVibration) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// This message can be requested by sending the MAV_CMD_GET_HOME_POSITION command. The position the system will return to and land on. The position is set automatically by the system during the takeoff in case it was not explicitly set by the operator before or after. The global and local positions encode the position in the respective coordinate frames, while the q parameter encodes the orientation of the surface. Under normal conditions it describes the heading and terrain slope, which can be used by the aircraft to adjust the approach. The approach 3D vector describes the point to which the system should fly in normal flight mode and then perform a landing sequence along the vector.
type MessageHomePosition struct {
	// Latitude (WGS84)
	Latitude int32
	// Longitude (WGS84)
	Longitude int32
	// Altitude (MSL). Positive for up.
	Altitude int32
	// Local X position of this position in the local coordinate frame
	X float32
	// Local Y position of this position in the local coordinate frame
	Y float32
	// Local Z position of this position in the local coordinate frame
	Z float32
	// World to surface normal and heading transformation of the takeoff position. Used to indicate the heading and slope of the ground
	Q [4]float32
	// Local X position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachX float32
	// Local Y position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachY float32
	// Local Z position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachZ float32
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64 `mavext:"true"`
}

func (m *MessageHomePosition) GetId() uint32 {
	return 242
}

func (m *MessageHomePosition) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The position the system will return to and land on. The position is set automatically by the system during the takeoff in case it was not explicitly set by the operator before or after. The global and local positions encode the position in the respective coordinate frames, while the q parameter encodes the orientation of the surface. Under normal conditions it describes the heading and terrain slope, which can be used by the aircraft to adjust the approach. The approach 3D vector describes the point to which the system should fly in normal flight mode and then perform a landing sequence along the vector.
type MessageSetHomePosition struct {
	// System ID.
	TargetSystem uint8
	// Latitude (WGS84)
	Latitude int32
	// Longitude (WGS84)
	Longitude int32
	// Altitude (MSL). Positive for up.
	Altitude int32
	// Local X position of this position in the local coordinate frame
	X float32
	// Local Y position of this position in the local coordinate frame
	Y float32
	// Local Z position of this position in the local coordinate frame
	Z float32
	// World to surface normal and heading transformation of the takeoff position. Used to indicate the heading and slope of the ground
	Q [4]float32
	// Local X position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachX float32
	// Local Y position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachY float32
	// Local Z position of the end of the approach vector. Multicopters should set this position based on their takeoff path. Grass-landing fixed wing aircraft should set it the same way as multicopters. Runway-landing fixed wing aircraft should set it to the opposite direction of the takeoff, assuming the takeoff happened from the threshold / touchdown zone.
	ApproachZ float32
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64 `mavext:"true"`
}

func (m *MessageSetHomePosition) GetId() uint32 {
	return 243
}

func (m *MessageSetHomePosition) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The interval between messages for a particular MAVLink message ID. This message is the response to the MAV_CMD_GET_MESSAGE_INTERVAL command. This interface replaces DATA_STREAM.
type MessageMessageInterval struct {
	// The ID of the requested MAVLink message. v1.0 is limited to 254 messages.
	MessageId uint16
	// The interval between two messages. A value of -1 indicates this stream is disabled, 0 indicates it is not available, &gt; 0 indicates the interval at which it is sent.
	IntervalUs int32
}

func (m *MessageMessageInterval) GetId() uint32 {
	return 244
}

func (m *MessageMessageInterval) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Provides state for additional features
type MessageExtendedSysState struct {
	// The VTOL state if applicable. Is set to MAV_VTOL_STATE_UNDEFINED if UAV is not in VTOL configuration.
	VtolState MAV_VTOL_STATE `mavenum:"uint8"`
	// The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown.
	LandedState MAV_LANDED_STATE `mavenum:"uint8"`
}

func (m *MessageExtendedSysState) GetId() uint32 {
	return 245
}

func (m *MessageExtendedSysState) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The location and information of an ADSB vehicle
type MessageAdsbVehicle struct {
	// ICAO address
	IcaoAddress uint32 `mavname:"ICAO_address"`
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// ADSB altitude type.
	AltitudeType ADSB_ALTITUDE_TYPE `mavenum:"uint8"`
	// Altitude(ASL)
	Altitude int32
	// Course over ground
	Heading uint16
	// The horizontal velocity
	HorVelocity uint16
	// The vertical velocity. Positive is up
	VerVelocity int16
	// The callsign, 8+null
	Callsign string `mavlen:"9"`
	// ADSB emitter type.
	EmitterType ADSB_EMITTER_TYPE `mavenum:"uint8"`
	// Time since last communication in seconds
	Tslc uint8
	// Bitmap to indicate various statuses including valid data fields
	Flags ADSB_FLAGS `mavenum:"uint16"`
	// Squawk code
	Squawk uint16
}

func (m *MessageAdsbVehicle) GetId() uint32 {
	return 246
}

func (m *MessageAdsbVehicle) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about a potential collision
type MessageCollision struct {
	// Collision data source
	Src MAV_COLLISION_SRC `mavenum:"uint8"`
	// Unique identifier, domain based on src field
	Id uint32
	// Action that is being taken to avoid this collision
	Action MAV_COLLISION_ACTION `mavenum:"uint8"`
	// How concerned the aircraft is about this collision
	ThreatLevel MAV_COLLISION_THREAT_LEVEL `mavenum:"uint8"`
	// Estimated time until collision occurs
	TimeToMinimumDelta float32
	// Closest vertical distance between vehicle and object
	AltitudeMinimumDelta float32
	// Closest horizontal distance between vehicle and object
	HorizontalMinimumDelta float32
}

func (m *MessageCollision) GetId() uint32 {
	return 247
}

func (m *MessageCollision) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Message implementing parts of the V2 payload specs in V1 frames for transitional support.
type MessageV2Extension struct {
	// Network ID (0 for broadcast)
	TargetNetwork uint8
	// System ID (0 for broadcast)
	TargetSystem uint8
	// Component ID (0 for broadcast)
	TargetComponent uint8
	// A code that identifies the software component that understands this message (analogous to USB device classes or mime type strings). If this code is less than 32768, it is considered a 'registered' protocol extension and the corresponding entry should be added to https://github.com/mavlink/mavlink/definition_files/extension_message_ids.xml. Software creators can register blocks of message IDs as needed (useful for GCS specific metadata, etc...). Message_types greater than 32767 are considered local experiments and should not be checked in to any widely distributed codebase.
	MessageType uint16
	// Variable length payload. The length must be encoded in the payload as part of the message_type protocol, e.g. by including the length as payload data, or by terminating the payload data with a non-zero marker. This is required in order to reconstruct zero-terminated payloads that are (or otherwise would be) trimmed by MAVLink 2 empty-byte truncation. The entire content of the payload block is opaque unless you understand the encoding message_type. The particular encoding used can be extension specific and might not always be documented as part of the MAVLink specification.
	Payload [249]uint8
}

func (m *MessageV2Extension) GetId() uint32 {
	return 248
}

func (m *MessageV2Extension) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Send raw controller memory. The use of this message is discouraged for normal packets, but a quite efficient way for testing new messages and getting experimental debug output.
type MessageMemoryVect struct {
	// Starting address of the debug variables
	Address uint16
	// Version code of the type variable. 0=unknown, type ignored and assumed int16_t. 1=as below
	Ver uint8
	// Type code of the memory variables. for ver = 1: 0=16 x int16_t, 1=16 x uint16_t, 2=16 x Q15, 3=16 x 1Q14
	Type uint8
	// Memory contents at specified address
	Value [32]int8
}

func (m *MessageMemoryVect) GetId() uint32 {
	return 249
}

func (m *MessageMemoryVect) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// To debug something using a named 3D vector.
type MessageDebugVect struct {
	// Name
	Name string `mavlen:"10"`
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// x
	X float32
	// y
	Y float32
	// z
	Z float32
}

func (m *MessageDebugVect) GetId() uint32 {
	return 250
}

func (m *MessageDebugVect) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Send a key-value pair as float. The use of this message is discouraged for normal packets, but a quite efficient way for testing new messages and getting experimental debug output.
type MessageNamedValueFloat struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the debug variable
	Name string `mavlen:"10"`
	// Floating point value
	Value float32
}

func (m *MessageNamedValueFloat) GetId() uint32 {
	return 251
}

func (m *MessageNamedValueFloat) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Send a key-value pair as integer. The use of this message is discouraged for normal packets, but a quite efficient way for testing new messages and getting experimental debug output.
type MessageNamedValueInt struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the debug variable
	Name string `mavlen:"10"`
	// Signed integer value
	Value int32
}

func (m *MessageNamedValueInt) GetId() uint32 {
	return 252
}

func (m *MessageNamedValueInt) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Status text message. These messages are printed in yellow in the COMM console of QGroundControl. WARNING: They consume quite some bandwidth, so use only for important status and error messages. If implemented wisely, these messages are buffered on the MCU and sent only at a limited rate (e.g. 10 Hz).
type MessageStatustext struct {
	// Severity of status. Relies on the definitions within RFC-5424.
	Severity MAV_SEVERITY `mavenum:"uint8"`
	// Status text message, without null termination character
	Text string `mavlen:"50"`
	// Unique (opaque) identifier for this statustext message.  May be used to reassemble a logical long-statustext message from a sequence of chunks.  A value of zero indicates this is the only chunk in the sequence and the message can be emitted immediately.
	Id uint16 `mavext:"true"`
	// This chunk's sequence number; indexing is from zero.  Any null character in the text field is taken to mean this was the last chunk.
	ChunkSeq uint8 `mavext:"true"`
}

func (m *MessageStatustext) GetId() uint32 {
	return 253
}

func (m *MessageStatustext) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Send a debug value. The index is used to discriminate between values. These values show up in the plot of QGroundControl as DEBUG N.
type MessageDebug struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// index of debug variable
	Ind uint8
	// DEBUG value
	Value float32
}

func (m *MessageDebug) GetId() uint32 {
	return 254
}

func (m *MessageDebug) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Setup a MAVLink2 signing key. If called with secret_key of all zero and zero initial_timestamp will disable signing
type MessageSetupSigning struct {
	// system id of the target
	TargetSystem uint8
	// component ID of the target
	TargetComponent uint8
	// signing key
	SecretKey [32]uint8
	// initial timestamp
	InitialTimestamp uint64
}

func (m *MessageSetupSigning) GetId() uint32 {
	return 256
}

func (m *MessageSetupSigning) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Report button state change.
type MessageButtonChange struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Time of last change of button state.
	LastChangeMs uint32
	// Bitmap for state of buttons.
	State uint8
}

func (m *MessageButtonChange) GetId() uint32 {
	return 257
}

func (m *MessageButtonChange) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Control vehicle tone generation (buzzer).
type MessagePlayTune struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// tune in board specific format
	Tune string `mavlen:"30"`
	// tune extension (appended to tune)
	Tune2 string `mavext:"true" mavlen:"200"`
}

func (m *MessagePlayTune) GetId() uint32 {
	return 258
}

func (m *MessagePlayTune) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about a camera
type MessageCameraInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the camera vendor
	VendorName [32]uint8
	// Name of the camera model
	ModelName [32]uint8
	// Version of the camera firmware (v &lt;&lt; 24 &amp; 0xff = Dev, v &lt;&lt; 16 &amp; 0xff = Patch, v &lt;&lt; 8 &amp; 0xff = Minor, v &amp; 0xff = Major)
	FirmwareVersion uint32
	// Focal length
	FocalLength float32
	// Image sensor size horizontal
	SensorSizeH float32
	// Image sensor size vertical
	SensorSizeV float32
	// Horizontal image resolution
	ResolutionH uint16
	// Vertical image resolution
	ResolutionV uint16
	// Reserved for a lens ID
	LensId uint8
	// Bitmap of camera capability flags.
	Flags CAMERA_CAP_FLAGS `mavenum:"uint32"`
	// Camera definition version (iteration)
	CamDefinitionVersion uint16
	// Camera definition URI (if any, otherwise only basic functions will be available). HTTP- (http://) and MAVLink FTP- (mavlinkftp://) formatted URIs are allowed (and both must be supported by any GCS that implements the Camera Protocol).
	CamDefinitionUri string `mavlen:"140"`
}

func (m *MessageCameraInformation) GetId() uint32 {
	return 259
}

func (m *MessageCameraInformation) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Settings of a camera, can be requested using MAV_CMD_REQUEST_CAMERA_SETTINGS.
type MessageCameraSettings struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Camera mode
	ModeId CAMERA_MODE `mavenum:"uint8"`
	// Current zoom level (0.0 to 100.0, NaN if not known)
	Zoomlevel float32 `mavext:"true" mavname:"zoomLevel"`
	// Current focus level (0.0 to 100.0, NaN if not known)
	Focuslevel float32 `mavext:"true" mavname:"focusLevel"`
}

func (m *MessageCameraSettings) GetId() uint32 {
	return 260
}

func (m *MessageCameraSettings) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about a storage medium. This message is sent in response to a request and whenever the status of the storage changes (STORAGE_STATUS).
type MessageStorageInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Storage ID (1 for first, 2 for second, etc.)
	StorageId uint8
	// Number of storage devices
	StorageCount uint8
	// Status of storage
	Status STORAGE_STATUS `mavenum:"uint8"`
	// Total capacity. If storage is not ready (STORAGE_STATUS_READY) value will be ignored.
	TotalCapacity float32
	// Used capacity. If storage is not ready (STORAGE_STATUS_READY) value will be ignored.
	UsedCapacity float32
	// Available storage capacity. If storage is not ready (STORAGE_STATUS_READY) value will be ignored.
	AvailableCapacity float32
	// Read speed.
	ReadSpeed float32
	// Write speed.
	WriteSpeed float32
}

func (m *MessageStorageInformation) GetId() uint32 {
	return 261
}

func (m *MessageStorageInformation) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about the status of a capture.
type MessageCameraCaptureStatus struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Current status of image capturing (0: idle, 1: capture in progress, 2: interval set but idle, 3: interval set and capture in progress)
	ImageStatus uint8
	// Current status of video capturing (0: idle, 1: capture in progress)
	VideoStatus uint8
	// Image capture interval
	ImageInterval float32
	// Time since recording started
	RecordingTimeMs uint32
	// Available storage capacity.
	AvailableCapacity float32
}

func (m *MessageCameraCaptureStatus) GetId() uint32 {
	return 262
}

func (m *MessageCameraCaptureStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about a captured image
type MessageCameraImageCaptured struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Timestamp (time since UNIX epoch) in UTC. 0 for unknown.
	TimeUtc uint64
	// Camera ID (1 for first, 2 for second, etc.)
	CameraId uint8
	// Latitude where image was taken
	Lat int32
	// Longitude where capture was taken
	Lon int32
	// Altitude (MSL) where image was taken
	Alt int32
	// Altitude above ground
	RelativeAlt int32
	// Quaternion of camera orientation (w, x, y, z order, zero-rotation is 0, 0, 0, 0)
	Q [4]float32
	// Zero based index of this image (image count since armed -1)
	ImageIndex int32
	// Boolean indicating success (1) or failure (0) while capturing this image.
	CaptureResult int8
	// URL of image taken. Either local storage or http://foo.jpg if camera provides an HTTP interface.
	FileUrl string `mavlen:"205"`
}

func (m *MessageCameraImageCaptured) GetId() uint32 {
	return 263
}

func (m *MessageCameraImageCaptured) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about flight since last arming.
type MessageFlightInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Timestamp at arming (time since UNIX epoch) in UTC, 0 for unknown
	ArmingTimeUtc uint64
	// Timestamp at takeoff (time since UNIX epoch) in UTC, 0 for unknown
	TakeoffTimeUtc uint64
	// Universally unique identifier (UUID) of flight, should correspond to name of log files
	FlightUuid uint64
}

func (m *MessageFlightInformation) GetId() uint32 {
	return 264
}

func (m *MessageFlightInformation) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Orientation of a mount
type MessageMountOrientation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Roll in global frame (set to NaN for invalid).
	Roll float32
	// Pitch in global frame (set to NaN for invalid).
	Pitch float32
	// Yaw relative to vehicle (set to NaN for invalid).
	Yaw float32
	// Yaw in absolute frame relative to Earth's North, north is 0 (set to NaN for invalid).
	YawAbsolute float32 `mavext:"true"`
}

func (m *MessageMountOrientation) GetId() uint32 {
	return 265
}

func (m *MessageMountOrientation) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// A message containing logged data (see also MAV_CMD_LOGGING_START)
type MessageLoggingData struct {
	// system ID of the target
	TargetSystem uint8
	// component ID of the target
	TargetComponent uint8
	// sequence number (can wrap)
	Sequence uint16
	// data length
	Length uint8
	// offset into data where first message starts. This can be used for recovery, when a previous message got lost (set to 255 if no start exists).
	FirstMessageOffset uint8
	// logged data
	Data [249]uint8
}

func (m *MessageLoggingData) GetId() uint32 {
	return 266
}

func (m *MessageLoggingData) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// A message containing logged data which requires a LOGGING_ACK to be sent back
type MessageLoggingDataAcked struct {
	// system ID of the target
	TargetSystem uint8
	// component ID of the target
	TargetComponent uint8
	// sequence number (can wrap)
	Sequence uint16
	// data length
	Length uint8
	// offset into data where first message starts. This can be used for recovery, when a previous message got lost (set to 255 if no start exists).
	FirstMessageOffset uint8
	// logged data
	Data [249]uint8
}

func (m *MessageLoggingDataAcked) GetId() uint32 {
	return 267
}

func (m *MessageLoggingDataAcked) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// An ack for a LOGGING_DATA_ACKED message
type MessageLoggingAck struct {
	// system ID of the target
	TargetSystem uint8
	// component ID of the target
	TargetComponent uint8
	// sequence number (must match the one in LOGGING_DATA_ACKED)
	Sequence uint16
}

func (m *MessageLoggingAck) GetId() uint32 {
	return 268
}

func (m *MessageLoggingAck) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about video stream
type MessageVideoStreamInformation struct {
	// Video Stream ID (1 for first, 2 for second, etc.)
	StreamId uint8
	// Number of streams available.
	Count uint8
	// Type of stream.
	Type VIDEO_STREAM_TYPE `mavenum:"uint8"`
	// Bitmap of stream status flags.
	Flags VIDEO_STREAM_STATUS_FLAGS `mavenum:"uint16"`
	// Frame rate.
	Framerate float32
	// Horizontal resolution.
	ResolutionH uint16
	// Vertical resolution.
	ResolutionV uint16
	// Bit rate.
	Bitrate uint32
	// Video image rotation clockwise.
	Rotation uint16
	// Horizontal Field of view.
	Hfov uint16
	// Stream name.
	Name string `mavlen:"32"`
	// Video stream URI (TCP or RTSP URI ground station should connect to) or port number (UDP port ground station should listen to).
	Uri string `mavlen:"160"`
}

func (m *MessageVideoStreamInformation) GetId() uint32 {
	return 269
}

func (m *MessageVideoStreamInformation) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about the status of a video stream.
type MessageVideoStreamStatus struct {
	// Video Stream ID (1 for first, 2 for second, etc.)
	StreamId uint8
	// Bitmap of stream status flags
	Flags VIDEO_STREAM_STATUS_FLAGS `mavenum:"uint16"`
	// Frame rate
	Framerate float32
	// Horizontal resolution
	ResolutionH uint16
	// Vertical resolution
	ResolutionV uint16
	// Bit rate
	Bitrate uint32
	// Video image rotation clockwise
	Rotation uint16
	// Horizontal Field of view
	Hfov uint16
}

func (m *MessageVideoStreamStatus) GetId() uint32 {
	return 270
}

func (m *MessageVideoStreamStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about a high level gimbal manager. This message should be requested by a ground station using MAV_CMD_REQUEST_MESSAGE.
type MessageGimbalManagerInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Bitmap of gimbal capability flags.
	CapFlags GIMBAL_MANAGER_CAP_FLAGS `mavenum:"uint32"`
	// Gimbal device ID that this gimbal manager is responsible for.
	GimbalDeviceId uint8
	// Maximum tilt/pitch angle (positive: up, negative: down)
	TiltMax float32
	// Minimum tilt/pitch angle (positive: up, negative: down)
	TiltMin float32
	// Maximum tilt/pitch angular rate (positive: up, negative: down)
	TiltRateMax float32
	// Maximum pan/yaw angle (positive: to the right, negative: to the left)
	PanMax float32
	// Minimum pan/yaw angle (positive: to the right, negative: to the left)
	PanMin float32
	// Minimum pan/yaw angular rate (positive: to the right, negative: to the left)
	PanRateMax float32
}

func (m *MessageGimbalManagerInformation) GetId() uint32 {
	return 280
}

func (m *MessageGimbalManagerInformation) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Current status about a high level gimbal manager. This message should be broadcast at a low regular rate (e.g. 5Hz).
type MessageGimbalManagerStatus struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// High level gimbal manager flags currently applied.
	Flags GIMBAL_MANAGER_FLAGS `mavenum:"uint32"`
	// Gimbal device ID that this gimbal manager is responsible for.
	GimbalDeviceId uint8
}

func (m *MessageGimbalManagerStatus) GetId() uint32 {
	return 281
}

func (m *MessageGimbalManagerStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// High level message to control a gimbal's attitude. This message is to be sent to the gimbal manager (e.g. from a ground station). Angles and rates can be set to NaN according to use case.
type MessageGimbalManagerSetAttitude struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// High level gimbal manager flags to use.
	Flags GIMBAL_MANAGER_FLAGS `mavenum:"uint32"`
	// Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. (Send command multiple times for more than one but not all gimbals.)
	GimbalDeviceId uint8
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation, the frame is depends on whether the flag GIMBAL_MANAGER_FLAGS_YAW_LOCK is set)
	Q [4]float32
	// X component of angular velocity, positive is banking to the right, NaN to be ignored.
	AngularVelocityX float32
	// Y component of angular velocity, positive is tilting up, NaN to be ignored.
	AngularVelocityY float32
	// Z component of angular velocity, positive is panning to the right, NaN to be ignored.
	AngularVelocityZ float32
}

func (m *MessageGimbalManagerSetAttitude) GetId() uint32 {
	return 282
}

func (m *MessageGimbalManagerSetAttitude) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about a low level gimbal. This message should be requested by the gimbal manager or a ground station using MAV_CMD_REQUEST_MESSAGE.
type MessageGimbalDeviceInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the gimbal vendor
	VendorName [32]uint8
	// Name of the gimbal model
	ModelName [32]uint8
	// Version of the gimbal firmware (v &lt;&lt; 24 &amp; 0xff = Dev, v &lt;&lt; 16 &amp; 0xff = Patch, v &lt;&lt; 8 &amp; 0xff = Minor, v &amp; 0xff = Major)
	FirmwareVersion uint32
	// Bitmap of gimbal capability flags.
	CapFlags GIMBAL_DEVICE_CAP_FLAGS `mavenum:"uint16"`
	// Maximum tilt/pitch angle (positive: up, negative: down)
	TiltMax float32
	// Minimum tilt/pitch angle (positive: up, negative: down)
	TiltMin float32
	// Maximum tilt/pitch angular rate (positive: up, negative: down)
	TiltRateMax float32
	// Maximum pan/yaw angle (positive: to the right, negative: to the left)
	PanMax float32
	// Minimum pan/yaw angle (positive: to the right, negative: to the left)
	PanMin float32
	// Minimum pan/yaw angular rate (positive: to the right, negative: to the left)
	PanRateMax float32
}

func (m *MessageGimbalDeviceInformation) GetId() uint32 {
	return 283
}

func (m *MessageGimbalDeviceInformation) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Low level message to control a gimbal device's attitude. This message is to be sent from the gimbal manager to the gimbal device component. Angles and rates can be set to NaN according to use case.
type MessageGimbalDeviceSetAttitude struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Low level gimbal flags.
	Flags GIMBAL_DEVICE_FLAGS `mavenum:"uint16"`
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation, the frame is depends on whether the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set, set all fields to NaN if only angular velocity should be used)
	Q [4]float32
	// X component of angular velocity, positive is banking to the right, NaN to be ignored.
	AngularVelocityX float32
	// Y component of angular velocity, positive is tilting up, NaN to be ignored.
	AngularVelocityY float32
	// Z component of angular velocity, positive is panning to the right, NaN to be ignored.
	AngularVelocityZ float32
}

func (m *MessageGimbalDeviceSetAttitude) GetId() uint32 {
	return 284
}

func (m *MessageGimbalDeviceSetAttitude) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Message reporting the status of a gimbal device. This message should be broadcasted by a gimbal device component. The angles encoded in the quaternion are in the global frame (roll: positive is tilt to the right, pitch: positive is tilting up, yaw is turn to the right). This message should be broadcast at a low regular rate (e.g. 10Hz).
type MessageGimbalDeviceAttitudeStatus struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Current gimbal flags set.
	Flags GIMBAL_DEVICE_FLAGS `mavenum:"uint16"`
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation, the frame is depends on whether the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set)
	Q [4]float32
	// X component of angular velocity (NaN if unknown)
	AngularVelocityX float32
	// Y component of angular velocity (NaN if unknown)
	AngularVelocityY float32
	// Z component of angular velocity (NaN if unknown)
	AngularVelocityZ float32
	// Failure flags (0 for no failure)
	FailureFlags GIMBAL_DEVICE_ERROR_FLAGS `mavenum:"uint32"`
}

func (m *MessageGimbalDeviceAttitudeStatus) GetId() uint32 {
	return 285
}

func (m *MessageGimbalDeviceAttitudeStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Low level message containing autopilot state relevant for a gimbal device. This message is to be sent from the gimbal manager to the gimbal device component. The data of this message server for the gimbal's estimator corrections in particular horizon compensation, as well as the autopilot's control intention e.g. feed forward angular control in z-axis.
type MessageAutopilotStateForGimbalDevice struct {
	// Timestamp (time since system boot).
	TimeBootUs uint64
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Quaternion components of autopilot attitude: w, x, y, z (1 0 0 0 is the null-rotation, Hamiltonian convention).
	Q [4]float32
	// Estimated delay of the attitude data.
	QEstimatedDelayUs uint32
	// X Speed in NED (North, East, Down).
	Vx float32
	// Y Speed in NED (North, East, Down).
	Vy float32
	// Z Speed in NED (North, East, Down).
	Vz float32
	// Estimated delay of the speed data.
	VEstimatedDelayUs uint32
	// Feed forward Z component of angular velocity, positive is yawing to the right, NaN to be ignored. This is to indicate if the autopilot is actively yawing.
	FeedForwardAngularVelocityZ float32
}

func (m *MessageAutopilotStateForGimbalDevice) GetId() uint32 {
	return 286
}

func (m *MessageAutopilotStateForGimbalDevice) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Configure AP SSID and Password.
type MessageWifiConfigAp struct {
	// Name of Wi-Fi network (SSID). Leave it blank to leave it unchanged.
	Ssid string `mavlen:"32"`
	// Password. Leave it blank for an open AP.
	Password string `mavlen:"64"`
}

func (m *MessageWifiConfigAp) GetId() uint32 {
	return 299
}

func (m *MessageWifiConfigAp) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Version and capability of protocol version. This message is the response to REQUEST_PROTOCOL_VERSION and is used as part of the handshaking to establish which MAVLink version should be used on the network. Every node should respond to REQUEST_PROTOCOL_VERSION to enable the handshaking. Library implementers should consider adding this into the default decoding state machine to allow the protocol core to respond directly.
type MessageProtocolVersion struct {
	// Currently active MAVLink version number * 100: v1.0 is 100, v2.0 is 200, etc.
	Version uint16
	// Minimum MAVLink version supported
	MinVersion uint16
	// Maximum MAVLink version supported (set to the same value as version by default)
	MaxVersion uint16
	// The first 8 bytes (not characters printed in hex!) of the git hash.
	SpecVersionHash [8]uint8
	// The first 8 bytes (not characters printed in hex!) of the git hash.
	LibraryVersionHash [8]uint8
}

func (m *MessageProtocolVersion) GetId() uint32 {
	return 300
}

func (m *MessageProtocolVersion) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The location and information of an AIS vessel
type MessageAisVessel struct {
	// Mobile Marine Service Identifier, 9 decimal digits
	Mmsi uint32 `mavname:"MMSI"`
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// Course over ground
	Cog uint16 `mavname:"COG"`
	// True heading
	Heading uint16
	// Speed over ground
	Velocity uint16
	// Turn rate
	TurnRate int8
	// Navigational status
	NavigationalStatus AIS_NAV_STATUS `mavenum:"uint8"`
	// Type of vessels
	Type AIS_TYPE `mavenum:"uint8"`
	// Distance from lat/lon location to bow
	DimensionBow uint16
	// Distance from lat/lon location to stern
	DimensionStern uint16
	// Distance from lat/lon location to port side
	DimensionPort uint8
	// Distance from lat/lon location to starboard side
	DimensionStarboard uint8
	// The vessel callsign
	Callsign string `mavlen:"7"`
	// The vessel name
	Name string `mavlen:"20"`
	// Time since last communication in seconds
	Tslc uint16
	// Bitmask to indicate various statuses including valid data fields
	Flags AIS_FLAGS `mavenum:"uint16"`
}

func (m *MessageAisVessel) GetId() uint32 {
	return 301
}

func (m *MessageAisVessel) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// General status information of an UAVCAN node. Please refer to the definition of the UAVCAN message "uavcan.protocol.NodeStatus" for the background information. The UAVCAN specification is available at http://uavcan.org.
type MessageUavcanNodeStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Time since the start-up of the node.
	UptimeSec uint32
	// Generalized node health status.
	Health UAVCAN_NODE_HEALTH `mavenum:"uint8"`
	// Generalized operating mode.
	Mode UAVCAN_NODE_MODE `mavenum:"uint8"`
	// Not used currently.
	SubMode uint8
	// Vendor-specific status information.
	VendorSpecificStatusCode uint16
}

func (m *MessageUavcanNodeStatus) GetId() uint32 {
	return 310
}

func (m *MessageUavcanNodeStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// General information describing a particular UAVCAN node. Please refer to the definition of the UAVCAN service "uavcan.protocol.GetNodeInfo" for the background information. This message should be emitted by the system whenever a new node appears online, or an existing node reboots. Additionally, it can be emitted upon request from the other end of the MAVLink channel (see MAV_CMD_UAVCAN_GET_NODE_INFO). It is also not prohibited to emit this message unconditionally at a low frequency. The UAVCAN specification is available at http://uavcan.org.
type MessageUavcanNodeInfo struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Time since the start-up of the node.
	UptimeSec uint32
	// Node name string. For example, "sapog.px4.io".
	Name string `mavlen:"80"`
	// Hardware major version number.
	HwVersionMajor uint8
	// Hardware minor version number.
	HwVersionMinor uint8
	// Hardware unique 128-bit ID.
	HwUniqueId [16]uint8
	// Software major version number.
	SwVersionMajor uint8
	// Software minor version number.
	SwVersionMinor uint8
	// Version control system (VCS) revision identifier (e.g. git short commit hash). Zero if unknown.
	SwVcsCommit uint32
}

func (m *MessageUavcanNodeInfo) GetId() uint32 {
	return 311
}

func (m *MessageUavcanNodeInfo) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request to read the value of a parameter with the either the param_id string id or param_index.
type MessageParamExtRequestRead struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter index. Set to -1 to use the Parameter ID field as identifier (else param_id will be ignored)
	ParamIndex int16
}

func (m *MessageParamExtRequestRead) GetId() uint32 {
	return 320
}

func (m *MessageParamExtRequestRead) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Request all parameters of this component. After this request, all parameters are emitted.
type MessageParamExtRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

func (m *MessageParamExtRequestList) GetId() uint32 {
	return 321
}

func (m *MessageParamExtRequestList) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Emit the value of a parameter. The inclusion of param_count and param_index in the message allows the recipient to keep track of received parameters and allows them to re-request missing parameters after a loss or timeout.
type MessageParamExtValue struct {
	// Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter value
	ParamValue string `mavlen:"128"`
	// Parameter type.
	ParamType MAV_PARAM_EXT_TYPE `mavenum:"uint8"`
	// Total number of parameters
	ParamCount uint16
	// Index of this parameter
	ParamIndex uint16
}

func (m *MessageParamExtValue) GetId() uint32 {
	return 322
}

func (m *MessageParamExtValue) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Set a parameter value. In order to deal with message loss (and retransmission of PARAM_EXT_SET), when setting a parameter value and the new value is the same as the current value, you will immediately get a PARAM_ACK_ACCEPTED response. If the current state is PARAM_ACK_IN_PROGRESS, you will accordingly receive a PARAM_ACK_IN_PROGRESS in response.
type MessageParamExtSet struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter value
	ParamValue string `mavlen:"128"`
	// Parameter type.
	ParamType MAV_PARAM_EXT_TYPE `mavenum:"uint8"`
}

func (m *MessageParamExtSet) GetId() uint32 {
	return 323
}

func (m *MessageParamExtSet) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Response from a PARAM_EXT_SET message.
type MessageParamExtAck struct {
	// Parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter value (new value if PARAM_ACK_ACCEPTED, current value otherwise)
	ParamValue string `mavlen:"128"`
	// Parameter type.
	ParamType MAV_PARAM_EXT_TYPE `mavenum:"uint8"`
	// Result code.
	ParamResult PARAM_ACK `mavenum:"uint8"`
}

func (m *MessageParamExtAck) GetId() uint32 {
	return 324
}

func (m *MessageParamExtAck) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Obstacle distances in front of the sensor, starting from the left in increment degrees to the right
type MessageObstacleDistance struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Class id of the distance sensor type.
	SensorType MAV_DISTANCE_SENSOR `mavenum:"uint8"`
	// Distance of obstacles around the vehicle with index 0 corresponding to north + angle_offset, unless otherwise specified in the frame. A value of 0 is valid and means that the obstacle is practically touching the sensor. A value of max_distance +1 means no obstacle is present. A value of UINT16_MAX for unknown/not used. In a array element, one unit corresponds to 1cm.
	Distances [72]uint16
	// Angular width in degrees of each array element. Increment direction is clockwise. This field is ignored if increment_f is non-zero.
	Increment uint8
	// Minimum distance the sensor can measure.
	MinDistance uint16
	// Maximum distance the sensor can measure.
	MaxDistance uint16
	// Angular width in degrees of each array element as a float. If non-zero then this value is used instead of the uint8_t increment field. Positive is clockwise direction, negative is counter-clockwise.
	IncrementF float32 `mavext:"true"`
	// Relative angle offset of the 0-index element in the distances array. Value of 0 corresponds to forward. Positive is clockwise direction, negative is counter-clockwise.
	AngleOffset float32 `mavext:"true"`
	// Coordinate frame of reference for the yaw rotation and offset of the sensor data. Defaults to MAV_FRAME_GLOBAL, which is north aligned. For body-mounted sensors use MAV_FRAME_BODY_FRD, which is vehicle front aligned.
	Frame MAV_FRAME `mavenum:"uint8" mavext:"true"`
}

func (m *MessageObstacleDistance) GetId() uint32 {
	return 330
}

func (m *MessageObstacleDistance) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Odometry message to communicate odometry information with an external interface. Fits ROS REP 147 standard for aerial vehicles (http://www.ros.org/reps/rep-0147.html).
type MessageOdometry struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Coordinate frame of reference for the pose data.
	FrameId MAV_FRAME `mavenum:"uint8"`
	// Coordinate frame of reference for the velocity in free space (twist) data.
	ChildFrameId MAV_FRAME `mavenum:"uint8"`
	// X Position
	X float32
	// Y Position
	Y float32
	// Z Position
	Z float32
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation)
	Q [4]float32
	// X linear speed
	Vx float32
	// Y linear speed
	Vy float32
	// Z linear speed
	Vz float32
	// Roll angular speed
	Rollspeed float32
	// Pitch angular speed
	Pitchspeed float32
	// Yaw angular speed
	Yawspeed float32
	// Row-major representation of a 6x6 pose cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	PoseCovariance [21]float32
	// Row-major representation of a 6x6 velocity cross-covariance matrix upper right triangle (states: vx, vy, vz, rollspeed, pitchspeed, yawspeed; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	VelocityCovariance [21]float32
	// Estimate reset counter. This should be incremented when the estimate resets in any of the dimensions (position, velocity, attitude, angular speed). This is designed to be used when e.g an external SLAM system detects a loop-closure and the estimate jumps.
	ResetCounter uint8 `mavext:"true"`
	// Type of estimator that is providing the odometry.
	EstimatorType MAV_ESTIMATOR_TYPE `mavenum:"uint8" mavext:"true"`
}

func (m *MessageOdometry) GetId() uint32 {
	return 331
}

func (m *MessageOdometry) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Describe a trajectory using an array of up-to 5 waypoints in the local frame (MAV_FRAME_LOCAL_NED).
type MessageTrajectoryRepresentationWaypoints struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Number of valid points (up-to 5 waypoints are possible)
	ValidPoints uint8
	// X-coordinate of waypoint, set to NaN if not being used
	PosX [5]float32
	// Y-coordinate of waypoint, set to NaN if not being used
	PosY [5]float32
	// Z-coordinate of waypoint, set to NaN if not being used
	PosZ [5]float32
	// X-velocity of waypoint, set to NaN if not being used
	VelX [5]float32
	// Y-velocity of waypoint, set to NaN if not being used
	VelY [5]float32
	// Z-velocity of waypoint, set to NaN if not being used
	VelZ [5]float32
	// X-acceleration of waypoint, set to NaN if not being used
	AccX [5]float32
	// Y-acceleration of waypoint, set to NaN if not being used
	AccY [5]float32
	// Z-acceleration of waypoint, set to NaN if not being used
	AccZ [5]float32
	// Yaw angle, set to NaN if not being used
	PosYaw [5]float32
	// Yaw rate, set to NaN if not being used
	VelYaw [5]float32
	// Scheduled action for each waypoint, UINT16_MAX if not being used.
	Command [5]MAV_CMD `mavenum:"uint16"`
}

func (m *MessageTrajectoryRepresentationWaypoints) GetId() uint32 {
	return 332
}

func (m *MessageTrajectoryRepresentationWaypoints) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Describe a trajectory using an array of up-to 5 bezier control points in the local frame (MAV_FRAME_LOCAL_NED).
type MessageTrajectoryRepresentationBezier struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Number of valid control points (up-to 5 points are possible)
	ValidPoints uint8
	// X-coordinate of bezier control points. Set to NaN if not being used
	PosX [5]float32
	// Y-coordinate of bezier control points. Set to NaN if not being used
	PosY [5]float32
	// Z-coordinate of bezier control points. Set to NaN if not being used
	PosZ [5]float32
	// Bezier time horizon. Set to NaN if velocity/acceleration should not be incorporated
	Delta [5]float32
	// Yaw. Set to NaN for unchanged
	PosYaw [5]float32
}

func (m *MessageTrajectoryRepresentationBezier) GetId() uint32 {
	return 333
}

func (m *MessageTrajectoryRepresentationBezier) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Report current used cellular network status
type MessageCellularStatus struct {
	// Status bitmap
	Status CELLULAR_NETWORK_STATUS_FLAG `mavenum:"uint16"`
	// Cellular network radio type: gsm, cdma, lte...
	Type CELLULAR_NETWORK_RADIO_TYPE `mavenum:"uint8"`
	// Cellular network RSSI/RSRP in dBm, absolute value
	Quality uint8
	// Mobile country code. If unknown, set to: UINT16_MAX
	Mcc uint16
	// Mobile network code. If unknown, set to: UINT16_MAX
	Mnc uint16
	// Location area code. If unknown, set to: 0
	Lac uint16
	// Cell ID. If unknown, set to: UINT32_MAX
	Cid uint32
}

func (m *MessageCellularStatus) GetId() uint32 {
	return 334
}

func (m *MessageCellularStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Status of the Iridium SBD link.
type MessageIsbdLinkStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	Timestamp uint64
	// Timestamp of the last successful sbd session. The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	LastHeartbeat uint64
	// Number of failed SBD sessions.
	FailedSessions uint16
	// Number of successful SBD sessions.
	SuccessfulSessions uint16
	// Signal quality equal to the number of bars displayed on the ISU signal strength indicator. Range is 0 to 5, where 0 indicates no signal and 5 indicates maximum signal strength.
	SignalQuality uint8
	// 1: Ring call pending, 0: No call pending.
	RingPending uint8
	// 1: Transmission session pending, 0: No transmission session pending.
	TxSessionPending uint8
	// 1: Receiving session pending, 0: No receiving session pending.
	RxSessionPending uint8
}

func (m *MessageIsbdLinkStatus) GetId() uint32 {
	return 335
}

func (m *MessageIsbdLinkStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The global position resulting from GPS and sensor fusion.
type MessageUtmGlobalPosition struct {
	// Time of applicability of position (microseconds since UNIX epoch).
	Time uint64
	// Unique UAS ID.
	UasId [18]uint8
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (WGS84)
	Alt int32
	// Altitude above ground
	RelativeAlt int32
	// Ground X speed (latitude, positive north)
	Vx int16
	// Ground Y speed (longitude, positive east)
	Vy int16
	// Ground Z speed (altitude, positive down)
	Vz int16
	// Horizontal position uncertainty (standard deviation)
	HAcc uint16
	// Altitude uncertainty (standard deviation)
	VAcc uint16
	// Speed uncertainty (standard deviation)
	VelAcc uint16
	// Next waypoint, latitude (WGS84)
	NextLat int32
	// Next waypoint, longitude (WGS84)
	NextLon int32
	// Next waypoint, altitude (WGS84)
	NextAlt int32
	// Time until next update. Set to 0 if unknown or in data driven mode.
	UpdateRate uint16
	// Flight state
	FlightState UTM_FLIGHT_STATE `mavenum:"uint8"`
	// Bitwise OR combination of the data available flags.
	Flags UTM_DATA_AVAIL_FLAGS `mavenum:"uint8"`
}

func (m *MessageUtmGlobalPosition) GetId() uint32 {
	return 340
}

func (m *MessageUtmGlobalPosition) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Large debug/prototyping array. The message uses the maximum available payload for data. The array_id and name fields are used to discriminate between messages in code and in user interfaces (respectively). Do not use in production code.
type MessageDebugFloatArray struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Name, for human-friendly display in a Ground Control Station
	Name string `mavlen:"10"`
	// Unique ID used to discriminate between arrays
	ArrayId uint16
	// data
	Data [58]float32 `mavext:"true"`
}

func (m *MessageDebugFloatArray) GetId() uint32 {
	return 350
}

func (m *MessageDebugFloatArray) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Vehicle status report that is sent out while orbit execution is in progress (see MAV_CMD_DO_ORBIT).
type MessageOrbitExecutionStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Radius of the orbit circle. Positive values orbit clockwise, negative values orbit counter-clockwise.
	Radius float32
	// The coordinate system of the fields: x, y, z.
	Frame MAV_FRAME `mavenum:"uint8"`
	// X coordinate of center point. Coordinate system depends on frame field: local = x position in meters * 1e4, global = latitude in degrees * 1e7.
	X int32
	// Y coordinate of center point.  Coordinate system depends on frame field: local = x position in meters * 1e4, global = latitude in degrees * 1e7.
	Y int32
	// Altitude of center point. Coordinate system depends on frame field.
	Z float32
}

func (m *MessageOrbitExecutionStatus) GetId() uint32 {
	return 360
}

func (m *MessageOrbitExecutionStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Smart Battery information (static/infrequent update). Use for updates from: smart battery to flight stack, flight stack to GCS. Use instead of BATTERY_STATUS for smart batteries.
type MessageSmartBatteryInfo struct {
	// Battery ID
	Id uint8
	// Capacity when full according to manufacturer, -1: field not provided.
	CapacityFullSpecification int32
	// Capacity when full (accounting for battery degradation), -1: field not provided.
	CapacityFull int32
	// Charge/discharge cycle count. -1: field not provided.
	CycleCount uint16
	// Serial number. -1: field not provided.
	SerialNumber int32
	// Static device name. Encode as manufacturer and product names separated using an underscore.
	DeviceName string `mavlen:"50"`
	// Battery weight. 0: field not provided.
	Weight uint16
	// Minimum per-cell voltage when discharging. If not supplied set to UINT16_MAX value.
	DischargeMinimumVoltage uint16
	// Minimum per-cell voltage when charging. If not supplied set to UINT16_MAX value.
	ChargingMinimumVoltage uint16
	// Minimum per-cell voltage when resting. If not supplied set to UINT16_MAX value.
	RestingMinimumVoltage uint16
}

func (m *MessageSmartBatteryInfo) GetId() uint32 {
	return 370
}

func (m *MessageSmartBatteryInfo) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Smart Battery information (dynamic). Use for updates from: smart battery to flight stack, flight stack to GCS. Use instead of BATTERY_STATUS for smart batteries.
type MessageSmartBatteryStatus struct {
	// Battery ID
	Id uint16
	// Remaining battery energy. Values: [0-100], -1: field not provided.
	CapacityRemaining int16
	// Battery current (through all cells/loads). Positive if discharging, negative if charging. UINT16_MAX: field not provided.
	Current int16
	// Battery temperature. -1: field not provided.
	Temperature int16
	// Fault/health indications.
	FaultBitmask MAV_SMART_BATTERY_FAULT `mavenum:"int32"`
	// Estimated remaining battery time. -1: field not provided.
	TimeRemaining int32
	// The cell number of the first index in the 'voltages' array field. Using this field allows you to specify cell voltages for batteries with more than 16 cells.
	CellOffset uint16
	// Individual cell voltages. Batteries with more 16 cells can use the cell_offset field to specify the cell offset for the array specified in the current message . Index values above the valid cell count for this battery should have the UINT16_MAX value.
	Voltages [16]uint16
}

func (m *MessageSmartBatteryStatus) GetId() uint32 {
	return 371
}

func (m *MessageSmartBatteryStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// The raw values of the actuator outputs (e.g. on Pixhawk, from MAIN, AUX ports). This message supersedes SERVO_OUTPUT_RAW.
type MessageActuatorOutputStatus struct {
	// Timestamp (since system boot).
	TimeUsec uint64
	// Active outputs
	Active uint32
	// Servo / motor output array values. Zero values indicate unused channels.
	Actuator [32]float32
}

func (m *MessageActuatorOutputStatus) GetId() uint32 {
	return 375
}

func (m *MessageActuatorOutputStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Time/duration estimates for various events and actions given the current vehicle state and position.
type MessageTimeEstimateToTarget struct {
	// Estimated time to complete the vehicle's configured "safe return" action from its current position (e.g. RTL, Smart RTL, etc.). -1 indicates that the vehicle is landed, or that no time estimate available.
	SafeReturn int32
	// Estimated time for vehicle to complete the LAND action from its current position. -1 indicates that the vehicle is landed, or that no time estimate available.
	Land int32
	// Estimated time for reaching/completing the currently active mission item. -1 means no time estimate available.
	MissionNextItem int32
	// Estimated time for completing the current mission. -1 means no mission active and/or no estimate available.
	MissionEnd int32
	// Estimated time for completing the current commanded action (i.e. Go To, Takeoff, Land, etc.). -1 means no action active and/or no estimate available.
	CommandedAction int32
}

func (m *MessageTimeEstimateToTarget) GetId() uint32 {
	return 380
}

func (m *MessageTimeEstimateToTarget) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Message for transporting "arbitrary" variable-length data from one component to another (broadcast is not forbidden, but discouraged). The encoding of the data is usually extension specific, i.e. determined by the source, and is usually not documented as part of the MAVLink specification.
type MessageTunnel struct {
	// System ID (can be 0 for broadcast, but this is discouraged)
	TargetSystem uint8
	// Component ID (can be 0 for broadcast, but this is discouraged)
	TargetComponent uint8
	// A code that identifies the content of the payload (0 for unknown, which is the default). If this code is less than 32768, it is a 'registered' payload type and the corresponding code should be added to the MAV_TUNNEL_PAYLOAD_TYPE enum. Software creators can register blocks of types as needed. Codes greater than 32767 are considered local experiments and should not be checked in to any widely distributed codebase.
	PayloadType MAV_TUNNEL_PAYLOAD_TYPE `mavenum:"uint16"`
	// Length of the data transported in payload
	PayloadLength uint8
	// Variable length payload. The payload length is defined by payload_length. The entire content of this block is opaque unless you understand the encoding specified by payload_type.
	Payload [128]uint8
}

func (m *MessageTunnel) GetId() uint32 {
	return 385
}

func (m *MessageTunnel) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Hardware status sent by an onboard computer.
type MessageOnboardComputerStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Time since system boot.
	Uptime uint32
	// Type of the onboard computer: 0: Mission computer primary, 1: Mission computer backup 1, 2: Mission computer backup 2, 3: Compute node, 4-5: Compute spares, 6-9: Payload computers.
	Type uint8
	// CPU usage on the component in percent (100 - idle). A value of UINT8_MAX implies the field is unused.
	CpuCores [8]uint8
	// Combined CPU usage as the last 10 slices of 100 MS (a histogram). This allows to identify spikes in load that max out the system, but only for a short amount of time. A value of UINT8_MAX implies the field is unused.
	CpuCombined [10]uint8
	// GPU usage on the component in percent (100 - idle). A value of UINT8_MAX implies the field is unused.
	GpuCores [4]uint8
	// Combined GPU usage as the last 10 slices of 100 MS (a histogram). This allows to identify spikes in load that max out the system, but only for a short amount of time. A value of UINT8_MAX implies the field is unused.
	GpuCombined [10]uint8
	// Temperature of the board. A value of INT8_MAX implies the field is unused.
	TemperatureBoard int8
	// Temperature of the CPU core. A value of INT8_MAX implies the field is unused.
	TemperatureCore [8]int8
	// Fan speeds. A value of INT16_MAX implies the field is unused.
	FanSpeed [4]int16
	// Amount of used RAM on the component system. A value of UINT32_MAX implies the field is unused.
	RamUsage uint32
	// Total amount of RAM on the component system. A value of UINT32_MAX implies the field is unused.
	RamTotal uint32
	// Storage type: 0: HDD, 1: SSD, 2: EMMC, 3: SD card (non-removable), 4: SD card (removable). A value of UINT32_MAX implies the field is unused.
	StorageType [4]uint32
	// Amount of used storage space on the component system. A value of UINT32_MAX implies the field is unused.
	StorageUsage [4]uint32
	// Total amount of storage space on the component system. A value of UINT32_MAX implies the field is unused.
	StorageTotal [4]uint32
	// Link type: 0-9: UART, 10-19: Wired network, 20-29: Wifi, 30-39: Point-to-point proprietary, 40-49: Mesh proprietary
	LinkType [6]uint32
	// Network traffic from the component system. A value of UINT32_MAX implies the field is unused.
	LinkTxRate [6]uint32
	// Network traffic to the component system. A value of UINT32_MAX implies the field is unused.
	LinkRxRate [6]uint32
	// Network capacity from the component system. A value of UINT32_MAX implies the field is unused.
	LinkTxMax [6]uint32
	// Network capacity to the component system. A value of UINT32_MAX implies the field is unused.
	LinkRxMax [6]uint32
}

func (m *MessageOnboardComputerStatus) GetId() uint32 {
	return 390
}

func (m *MessageOnboardComputerStatus) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Information about a component. For camera components instead use CAMERA_INFORMATION, and for autopilots use AUTOPILOT_VERSION. Components including GCSes should consider supporting requests of this message via MAV_CMD_REQUEST_MESSAGE.
type MessageComponentInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the component vendor
	VendorName [32]uint8
	// Name of the component model
	ModelName [32]uint8
	// Version of the component firmware (v &lt;&lt; 24 &amp; 0xff = Dev, v &lt;&lt; 16 &amp; 0xff = Patch, v &lt;&lt; 8 &amp; 0xff = Minor, v &amp; 0xff = Major)
	FirmwareVersion uint32
	// Version of the component hardware (v &lt;&lt; 24 &amp; 0xff = Dev, v &lt;&lt; 16 &amp; 0xff = Patch, v &lt;&lt; 8 &amp; 0xff = Minor, v &amp; 0xff = Major)
	HardwareVersion uint32
	// Bitmap of component capability flags.
	CapabilityFlags COMPONENT_CAP_FLAGS `mavenum:"uint32"`
	// Component definition version (iteration)
	ComponentDefinitionVersion uint16
	// Component definition URI (if any, otherwise only basic functions will be available). The XML format is not yet specified and work in progress.
	ComponentDefinitionUri string `mavlen:"140"`
}

func (m *MessageComponentInformation) GetId() uint32 {
	return 395
}

func (m *MessageComponentInformation) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Play vehicle tone/tune (buzzer). Supersedes message PLAY_TUNE.
type MessagePlayTuneV2 struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Tune format
	Format TUNE_FORMAT `mavenum:"uint32"`
	// Tune definition as a NULL-terminated string.
	Tune string `mavlen:"248"`
}

func (m *MessagePlayTuneV2) GetId() uint32 {
	return 400
}

func (m *MessagePlayTuneV2) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Tune formats supported by vehicle. This should be emitted as response to MAV_CMD_REQUEST_MESSAGE.
type MessageSupportedTunes struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Bitfield of supported tune formats.
	Format TUNE_FORMAT `mavenum:"uint32"`
}

func (m *MessageSupportedTunes) GetId() uint32 {
	return 401
}

func (m *MessageSupportedTunes) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Cumulative distance traveled for each reported wheel.
type MessageWheelDistance struct {
	// Timestamp (synced to UNIX time or since system boot).
	TimeUsec uint64
	// Number of wheels reported.
	Count uint8
	// Distance reported by individual wheel encoders. Forward rotations increase values, reverse rotations decrease them. Not all wheels will necessarily have wheel encoders; the mapping of encoders to wheel positions must be agreed/understood by the endpoints.
	Distance [16]float64
}

func (m *MessageWheelDistance) GetId() uint32 {
	return 9000
}

func (m *MessageWheelDistance) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Data for filling the OpenDroneID Basic ID message. This and the below messages are primarily meant for feeding data to/from an OpenDroneID implementation. E.g. https://github.com/opendroneid/opendroneid-core-c. See also the ASTM Remote ID standard at https://www.astm.org/Standards/F3411.htm. The usage of these messages is documented at https://mavlink.io/en/services/opendroneid.html.
type MessageOpenDroneIdBasicId struct {
	// Indicates the format for the uas_id field of this message.
	IdType MAV_ODID_ID_TYPE `mavenum:"uint8"`
	// Indicates the type of UA (Unmanned Aircraft).
	UaType MAV_ODID_UA_TYPE `mavenum:"uint8"`
	// UAS (Unmanned Aircraft System) ID following the format specified by id_type. Shall be filled with nulls in the unused portion of the field.
	UasId [20]uint8
}

func (m *MessageOpenDroneIdBasicId) GetId() uint32 {
	return 12900
}

func (m *MessageOpenDroneIdBasicId) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Data for filling the OpenDroneID Location message. The float data types are 32-bit IEEE 754. The Location message provides the location, altitude, direction and speed of the aircraft.
type MessageOpenDroneIdLocation struct {
	// Indicates whether the unmanned aircraft is on the ground or in the air.
	Status MAV_ODID_STATUS `mavenum:"uint8"`
	// Direction over ground (not heading, but direction of movement) measured clockwise from true North: 0 - 35999 centi-degrees. If unknown: 36100 centi-degrees.
	Direction uint16
	// Ground speed. Positive only. If unknown: 25500 cm/s. If speed is larger than 25425 cm/s, use 25425 cm/s.
	SpeedHorizontal uint16
	// The vertical speed. Up is positive. If unknown: 6300 cm/s. If speed is larger than 6200 cm/s, use 6200 cm/s. If lower than -6200 cm/s, use -6200 cm/s.
	SpeedVertical int16
	// Current latitude of the unmanned aircraft. If unknown: 0 (both Lat/Lon).
	Latitude int32
	// Current longitude of the unmanned aircraft. If unknown: 0 (both Lat/Lon).
	Longitude int32
	// The altitude calculated from the barometric pressue. Reference is against 29.92inHg or 1013.2mb. If unknown: -1000 m.
	AltitudeBarometric float32
	// The geodetic altitude as defined by WGS84. If unknown: -1000 m.
	AltitudeGeodetic float32
	// Indicates the reference point for the height field.
	HeightReference MAV_ODID_HEIGHT_REF `mavenum:"uint8"`
	// The current height of the unmanned aircraft above the take-off location or the ground as indicated by height_reference. If unknown: -1000 m.
	Height float32
	// The accuracy of the horizontal position.
	HorizontalAccuracy MAV_ODID_HOR_ACC `mavenum:"uint8"`
	// The accuracy of the vertical position.
	VerticalAccuracy MAV_ODID_VER_ACC `mavenum:"uint8"`
	// The accuracy of the barometric altitude.
	BarometerAccuracy MAV_ODID_VER_ACC `mavenum:"uint8"`
	// The accuracy of the horizontal and vertical speed.
	SpeedAccuracy MAV_ODID_SPEED_ACC `mavenum:"uint8"`
	// Seconds after the full hour with reference to UTC time. Typically the GPS outputs a time-of-week value in milliseconds. First convert that to UTC and then convert for this field using ((float) (time_week_ms % (60*60*1000))) / 1000.
	Timestamp float32
	// The accuracy of the timestamps.
	TimestampAccuracy MAV_ODID_TIME_ACC `mavenum:"uint8"`
}

func (m *MessageOpenDroneIdLocation) GetId() uint32 {
	return 12901
}

func (m *MessageOpenDroneIdLocation) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Data for filling the OpenDroneID Authentication message. The Authentication Message defines a field that can provide a means of authenticity for the identity of the UAS (Unmanned Aircraft System). The Authentication message can have two different formats. Five data pages are supported. For data page 0, the fields PageCount, Length and TimeStamp are present and AuthData is only 17 bytes. For data page 1 through 4, PageCount, Length and TimeStamp are not present and the size of AuthData is 23 bytes.
type MessageOpenDroneIdAuthentication struct {
	// Indicates the type of authentication.
	AuthenticationType MAV_ODID_AUTH_TYPE `mavenum:"uint8"`
	// Allowed range is 0 - 4.
	DataPage uint8
	// This field is only present for page 0. Allowed range is 0 - 5.
	PageCount uint8
	// This field is only present for page 0. Total bytes of authentication_data from all data pages. Allowed range is 0 - 109 (17 + 23*4).
	Length uint8
	// This field is only present for page 0. 32 bit Unix Timestamp in seconds since 00:00:00 01/01/2019.
	Timestamp uint32
	// Opaque authentication data. For page 0, the size is only 17 bytes. For other pages, the size is 23 bytes. Shall be filled with nulls in the unused portion of the field.
	AuthenticationData [23]uint8
}

func (m *MessageOpenDroneIdAuthentication) GetId() uint32 {
	return 12902
}

func (m *MessageOpenDroneIdAuthentication) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Data for filling the OpenDroneID Self ID message. The Self ID Message is an opportunity for the operator to (optionally) declare their identity and purpose of the flight. This message can provide additional information that could reduce the threat profile of a UA (Unmanned Aircraft) flying in a particular area or manner.
type MessageOpenDroneIdSelfId struct {
	// Indicates the type of the description field.
	DescriptionType MAV_ODID_DESC_TYPE `mavenum:"uint8"`
	// Text description or numeric value expressed as ASCII characters. Shall be filled with nulls in the unused portion of the field.
	Description string `mavlen:"23"`
}

func (m *MessageOpenDroneIdSelfId) GetId() uint32 {
	return 12903
}

func (m *MessageOpenDroneIdSelfId) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Data for filling the OpenDroneID System message. The System Message contains general system information including the operator location and possible aircraft group information.
type MessageOpenDroneIdSystem struct {
	// Specifies the location source for the operator location.
	Flags MAV_ODID_LOCATION_SRC `mavenum:"uint8"`
	// Latitude of the operator. If unknown: 0 (both Lat/Lon).
	OperatorLatitude int32
	// Longitude of the operator. If unknown: 0 (both Lat/Lon).
	OperatorLongitude int32
	// Number of aircraft in the area, group or formation (default 1).
	AreaCount uint16
	// Radius of the cylindrical area of the group or formation (default 0).
	AreaRadius uint16
	// Area Operations Ceiling relative to WGS84. If unknown: -1000 m.
	AreaCeiling float32
	// Area Operations Floor relative to WGS84. If unknown: -1000 m.
	AreaFloor float32
}

func (m *MessageOpenDroneIdSystem) GetId() uint32 {
	return 12904
}

func (m *MessageOpenDroneIdSystem) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// Data for filling the OpenDroneID Operator ID message, which contains the CAA (Civil Aviation Authority) issued operator ID.
type MessageOpenDroneIdOperatorId struct {
	// Indicates the type of the operator_id field.
	OperatorIdType MAV_ODID_OPERATOR_ID_TYPE `mavenum:"uint8"`
	// Text description or numeric value expressed as ASCII characters. Shall be filled with nulls in the unused portion of the field.
	OperatorId string `mavlen:"20"`
}

func (m *MessageOpenDroneIdOperatorId) GetId() uint32 {
	return 12905
}

func (m *MessageOpenDroneIdOperatorId) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}

// An OpenDroneID message pack is a container for multiple encoded OpenDroneID messages (i.e. not in the format given for the above messages descriptions but after encoding into the compressed OpenDroneID byte format). Used e.g. when transmitting on Bluetooth 5.0 Long Range/Extended Advertising or on WiFi Neighbor Aware Networking.
type MessageOpenDroneIdMessagePack struct {
	// This field must currently always be equal to 25 bytes, since all encoded OpenDroneID messages are specificed to have this length.
	SingleMessageSize uint8
	// Number of encoded messages in the pack (not the number of bytes). Allowed range is 1 - 10.
	MsgPackSize uint8
	// Concatenation of encoded OpenDroneID messages. Shall be filled with nulls in the unused portion of the field.
	Messages [250]uint8
}

func (m *MessageOpenDroneIdMessagePack) GetId() uint32 {
	return 12915
}

func (m *MessageOpenDroneIdMessagePack) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}
