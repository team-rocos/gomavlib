package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/team-rocos/gomavlib"
	"github.com/team-rocos/gomavlib/dialects/common"
)

var (
	address  string
	messages []string
	stop     float64
	random   bool
)

func messageDetails(input []string, messageName string) (bool, int) {
	messagePeriod := 0
	messageFound := false
	for _, msgString := range input {
		msg := strings.Split(msgString, "@")
		msgNameCheck := strings.ToUpper(msg[0])
		if (msgNameCheck != "HEARTBEAT") && (msgNameCheck != "SYSSTATUS") && (msgNameCheck != "ATTITUDE") {
			err := fmt.Errorf("invalid message name. Messages available are: Heartbeat, Attitude, and SysStatus")
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		if msgNameCheck == messageName {
			messageFound = true
			if len(msg) > 1 {
				if strings.Contains(msg[1], "ms") { // period input in milliseconds
					period := strings.SplitN(msg[1], "ms", 2)
					var err error
					messagePeriod, err = strconv.Atoi(period[0])
					if err != nil {
						fmt.Println("Error: ", err)
						os.Exit(1)
					}
				} else if strings.Contains(msg[1], "s") { // period input in seconds
					period := strings.SplitN(msg[1], "s", 2)
					var err error
					messagePeriod, err = strconv.Atoi(period[0])
					messagePeriod *= 1000 // Convert from seconds to milliseconds
					if err != nil {
						fmt.Println("Error: ", err)
						os.Exit(1)
					}
				}

			}
			break
		}
	}
	if !messageFound {
		return false, 0
	}

	return true, messagePeriod
}

func randomHeartbeat() *common.MessageHeartbeat {
	rand.Seed(time.Now().UnixNano())
	return &common.MessageHeartbeat{
		Type:           common.MAV_TYPE(rand.Intn(34)),
		Autopilot:      common.MAV_AUTOPILOT(rand.Intn(19)) + 1,
		BaseMode:       1,
		CustomMode:     3,
		SystemStatus:   common.MAV_STATE(rand.Intn(9)),
		MavlinkVersion: 2,
	}
}
func randomAttitude() *common.MessageAttitude {
	rand.Seed(time.Now().UnixNano())
	return &common.MessageAttitude{
		TimeBootMs: uint32(time.Now().UnixNano()),
		Roll:       float32(rand.Intn(101)),
		Pitch:      float32(rand.Intn(101)),
		Yaw:        float32(rand.Intn(101)),
		Rollspeed:  float32(rand.Intn(101)),
		Pitchspeed: float32(rand.Intn(101)),
		Yawspeed:   float32(rand.Intn(101)),
	}
}
func randomSysStatus() *common.MessageSysStatus {
	rand.Seed(time.Now().UnixNano())
	return &common.MessageSysStatus{
		OnboardControlSensorsPresent: 1,
		OnboardControlSensorsEnabled: 1,
		OnboardControlSensorsHealth:  1,
		Load:                         uint16(rand.Intn(900)),
		VoltageBattery:               uint16(rand.Intn(101)),
		CurrentBattery:               int16(rand.Intn(101)),
		BatteryRemaining:             int8(rand.Intn(101)),
		DropRateComm:                 0,
		ErrorsComm:                   0,
		ErrorsCount1:                 0,
		ErrorsCount2:                 0,
		ErrorsCount3:                 0,
		ErrorsCount4:                 0,
	}
}

var msg1 = &common.MessageHeartbeat{
	Type:           6,
	Autopilot:      5,
	BaseMode:       4,
	CustomMode:     3,
	SystemStatus:   2,
	MavlinkVersion: 2,
}

var msg2 = &common.MessageSysStatus{
	OnboardControlSensorsPresent: 1,
	OnboardControlSensorsEnabled: 1,
	OnboardControlSensorsHealth:  1,
	Load:                         100,
	VoltageBattery:               50,
	CurrentBattery:               60,
	BatteryRemaining:             50,
	DropRateComm:                 0,
	ErrorsComm:                   0,
	ErrorsCount1:                 0,
	ErrorsCount2:                 0,
	ErrorsCount3:                 0,
	ErrorsCount4:                 0,
}

var msg3 = &common.MessageAttitude{
	TimeBootMs: 2,
	Roll:       0,
	Pitch:      1,
	Yaw:        2,
	Rollspeed:  3,
	Pitchspeed: 4,
	Yawspeed:   5,
}

func main() {
	pflag.StringVarP(&address, "address", "a", "", "Set tcp address to which to send MAVlink messages.")
	pflag.StringSliceVarP(&messages, "msg", "m", []string{""}, "Set message(s) to be sent as well and the period at which they are sent.")
	pflag.Float64VarP(&stop, "stop", "s", 10, "Set the length of time in seconds to send messages. Up to millisecond resolution (3dp).")
	pflag.BoolVarP(&random, "random", "r", false, "Use this to randomise the fields of the messages sent.")
	pflag.Parse()

	fmt.Println("Setting address to: ", address)
	fmt.Println("Setting timeout to: ", stop)
	if random {
		fmt.Println("Sending messages with randomised field values")
	} else {
		fmt.Println("Sending messages with constant field values")
	}

	sendHeartbeat, heartbeatPeriod := messageDetails(messages, "HEARTBEAT")
	sendSysStatus, sysStatusPeriod := messageDetails(messages, "SYSSTATUS")
	sendAttitude, attitudePeriod := messageDetails(messages, "ATTITUDE")

	fmt.Println("sendHeartbeat = ", sendHeartbeat, ", sendSysStatus = ", sendSysStatus, ", sendAttitude = ", sendAttitude)
	fmt.Println("heartbeatPeriod = ", heartbeatPeriod, ", sysStatusPeriod = ", sysStatusPeriod, ", attitudePeriod = ", attitudePeriod)

	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			//gomavlib.EndpointTcpServer{"192.168.1.74:5600"},
			gomavlib.EndpointTcpServer{address},
		},
		D:           common.Dialect,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to write to the target
		OutSystemId: 10,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	nodeReceive, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointTcpClient{address},
		},
		D:           common.Dialect,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to write to the target
		OutSystemId: 10,
	})
	if err != nil {
		panic(err)
	}
	defer nodeReceive.Close()

	stopchan := make(chan struct{})
	stoppedchanHeartbeat := make(chan struct{})
	stoppedchanSysStatus := make(chan struct{})
	stoppedchanAttitude := make(chan struct{})
	stoppedchanReceive := make(chan struct{})

	heartbeatMessagesReceived := 0
	sysStatusMessagesReceived := 0
	attitudeMessagesReceived := 0

	countHeartbeat := 0
	countSysStatus := 0
	countAttitude := 0

	if sendHeartbeat {
		go func() { // Heartbeat
			defer close(stoppedchanHeartbeat)
			for t := range time.NewTicker(time.Duration(heartbeatPeriod) * time.Millisecond).C {
				select {
				default:
					if random {
						node.WriteMessageAll(randomHeartbeat())
					} else {
						node.WriteMessageAll(msg1)
					}
					countHeartbeat++
				case <-stopchan:
					fmt.Println("Time elapsed: ", t)
					fmt.Println("Closing Heartbeat go routine...")
					return
				}
			}
		}()
	}

	if sendSysStatus {
		go func() { // SysStatus
			defer close(stoppedchanSysStatus)
			for range time.NewTicker(time.Duration(sysStatusPeriod) * time.Millisecond).C {
				select {
				default:
					if random {
						node.WriteMessageAll(randomSysStatus())
					} else {
						node.WriteMessageAll(msg2)
					}
					countSysStatus++
				case <-stopchan:
					fmt.Println("Closing SysStatus go routine...")
					return
				}
			}
		}()
	}

	if sendAttitude {
		go func() { // Attitude
			defer close(stoppedchanAttitude)
			for range time.NewTicker(time.Duration(attitudePeriod) * time.Millisecond).C {
				select {
				default:
					if random {
						node.WriteMessageAll(randomAttitude())
					} else {
						node.WriteMessageAll(msg3)
					}
					countAttitude++
				case <-stopchan:
					fmt.Println("Closing Attitude go routine...")
					return
				}
			}
		}()
	}

	go func() { // Receive
		defer close(stoppedchanReceive)
		// print every message we receive
		for evt := range nodeReceive.Events() {
			select {
			default:
				if frm, ok := evt.(*gomavlib.EventFrame); ok {
					if msgReceived, ok := frm.Message().(*common.MessageHeartbeat); ok {
						if msgReceived.Autopilot == 0 {
							fmt.Printf("received: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
						} else {
							heartbeatMessagesReceived++
						}
					} else if _, ok := frm.Message().(*common.MessageSysStatus); ok {
						sysStatusMessagesReceived++
					} else if _, ok := frm.Message().(*common.MessageAttitude); ok {
						attitudeMessagesReceived++
					}
				}
			case <-stopchan:
				fmt.Println("Closing nodeReceive go routine...")
				return
			}
		}
	}()

	time.Sleep(time.Duration(stop*1000) * time.Millisecond)
	close(stopchan)

	if sendHeartbeat {
		<-stoppedchanHeartbeat
	}
	if sendSysStatus {
		<-stoppedchanSysStatus
	}
	if sendAttitude {
		<-stoppedchanAttitude
	}
	<-stoppedchanReceive
	fmt.Println("Closed all go routines!")

	// Print counts
	fmt.Println("countHeartbeat = ", countHeartbeat)
	fmt.Println("countSysStatus = ", countSysStatus)
	fmt.Println("countAttitude = ", countAttitude)

	fmt.Println("Heartbeat Received: ", heartbeatMessagesReceived)
	fmt.Println("SysStatus Received: ", sysStatusMessagesReceived)
	fmt.Println("Attitude Received: ", attitudeMessagesReceived)
}
