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
	"github.com/team-rocos/gomavlib/commands/dialgen/libgen"
)

var (
	address     string
	messages    []string
	stop        float64
	random, udp bool
)

func messageDetails(input []string, messageName string) (bool, int) {
	messagePeriod := 0
	messageFound := false
	for _, msgString := range input {
		msg := strings.Split(msgString, "@")
		msgNameCheck := strings.ToUpper(msg[0])
		if (msgNameCheck != "HEARTBEAT") && (msgNameCheck != "SYS_STATUS") && (msgNameCheck != "ATTITUDE") {
			err := fmt.Errorf("invalid message name. Messages available are: HEARTBEAT, ATTITUDE, and SYS_STATUS")
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		if msgNameCheck == messageName {
			messageFound = true
			if len(msg) > 1 {
				if strings.Contains(msg[1], "us") { // period input in milliseconds
					period := strings.SplitN(msg[1], "us", 2)
					var err error
					messagePeriod, err = strconv.Atoi(period[0])
					if err != nil {
						fmt.Println("Error: ", err)
						os.Exit(1)
					}
				} else if strings.Contains(msg[1], "ms") { // period input in milliseconds
					period := strings.SplitN(msg[1], "ms", 2)
					var err error
					messagePeriod, err = strconv.Atoi(period[0])
					messagePeriod *= 1000
					if err != nil {
						fmt.Println("Error: ", err)
						os.Exit(1)
					}
				} else if strings.Contains(msg[1], "s") { // period input in seconds
					period := strings.SplitN(msg[1], "s", 2)
					var err error
					messagePeriod, err = strconv.Atoi(period[0])
					messagePeriod *= 1000000 // Convert from seconds to milliseconds
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

func randomHeartbeat(dm *gomavlib.DynamicMessage) error {
	rand.Seed(time.Now().UnixNano())
	err := dm.SetField("type", uint8(rand.Intn(34)))
	if err != nil {
		return err
	}
	err = dm.SetField("autopilot", uint8(rand.Intn(19)+1))
	if err != nil {
		return err
	}
	err = dm.SetField("base_mode", uint8(1))
	if err != nil {
		return err
	}
	err = dm.SetField("custom_mode", uint32(3))
	if err != nil {
		return err
	}
	err = dm.SetField("system_status", uint8(rand.Intn(9)))
	if err != nil {
		return err
	}
	err = dm.SetField("mavlink_version", uint8(2))
	if err != nil {
		return err
	}
	return nil
}

func randomAttitude(dm *gomavlib.DynamicMessage) error {
	rand.Seed(time.Now().UnixNano())
	err := dm.SetField("time_boot_ms", uint32(rand.Intn(34)))
	if err != nil {
		return err
	}
	err = dm.SetField("roll", float32(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("pitch", float32(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("yaw", float32(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("rollspeed", float32(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("pitchspeed", float32(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("yawspeed", float32(rand.Intn(101)))
	if err != nil {
		return err
	}
	return nil
}

func randomSysStatus(dm *gomavlib.DynamicMessage) error {
	rand.Seed(time.Now().UnixNano())
	err := dm.SetField("onboard_control_sensors_present", uint32(rand.Intn(34)))
	if err != nil {
		return err
	}
	err = dm.SetField("onboard_control_sensors_enabled", uint32(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("onboard_control_sensors_health", uint32(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("load", uint16(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("voltage_battery", uint16(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("current_battery", int16(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("battery_remaining", int8(rand.Intn(101)))
	if err != nil {
		return err
	}
	err = dm.SetField("drop_rate_comm", uint16(rand.Intn(10)))
	if err != nil {
		return err
	}
	err = dm.SetField("errors_comm", uint16(rand.Intn(10)))
	if err != nil {
		return err
	}
	err = dm.SetField("errors_count1", uint16(rand.Intn(10)))
	if err != nil {
		return err
	}
	err = dm.SetField("errors_count2", uint16(rand.Intn(10)))
	if err != nil {
		return err
	}
	err = dm.SetField("errors_count3", uint16(rand.Intn(10)))
	if err != nil {
		return err
	}
	err = dm.SetField("errors_count4", uint16(rand.Intn(10)))
	if err != nil {
		return err
	}

	return nil
}

var heartbeatMsg *gomavlib.DynamicMessage
var sysStatusMsg *gomavlib.DynamicMessage
var attitudeMsg *gomavlib.DynamicMessage

var node *gomavlib.Node
var nodeReceive *gomavlib.Node

func main() {
	pflag.StringVarP(&address, "address", "a", "", "Set address to which to send MAVlink messages.")
	pflag.StringSliceVarP(&messages, "msg", "m", []string{""}, "Set message(s) to be sent as well and the period at which they are sent.")
	pflag.Float64VarP(&stop, "stop", "s", 10, "Set the length of time in seconds to send messages. Up to millisecond resolution (3dp).")
	pflag.BoolVarP(&random, "random", "r", false, "Use this to randomise the fields of the messages sent.")
	pflag.BoolVarP(&udp, "udp", "u", false, "Set this to send to a udp port. Otherwise defaults to tcp")
	pflag.Parse()

	addressType := "TCP"
	if udp {
		addressType = "UDP"
	}
	fmt.Println("Setting address to: ", addressType, address)
	fmt.Println("Setting timeout to: ", stop)

	random = true // TODO: create constant value messages as alternative to randomised message fields
	if random {
		fmt.Println("Sending messages with randomised field values")
	} else {
		fmt.Println("Sending messages with constant field values")
	}

	sendHeartbeat, heartbeatPeriod := messageDetails(messages, "HEARTBEAT")
	sendSysStatus, sysStatusPeriod := messageDetails(messages, "SYS_STATUS")
	sendAttitude, attitudePeriod := messageDetails(messages, "ATTITUDE")

	heartbeatFreq := 0
	sysStatusFreq := 0
	attitudeFreq := 0
	if sendHeartbeat {
		heartbeatFreq = 1000000 / heartbeatPeriod
	}
	if sendSysStatus {
		sysStatusFreq = 1000000 / sysStatusPeriod
	}
	if sendAttitude {
		attitudeFreq = 1000000 / attitudePeriod
	}
	fmt.Println("heartbeatFreq = ", heartbeatFreq, ", sysStatusFreq = ", sysStatusFreq, ", attitudeFreq = ", attitudeFreq)

	defs, version, err := libgen.XMLToFields("../mavlink-upstream/message_definitions/v1.0/common.xml", []string{"."})
	if err != nil {
		fmt.Println("error creating defs from xml file, err: ", err)
		os.Exit(1)
	}
	// Create dialect from the parsed defs.
	dRT, err := gomavlib.NewDialectRT(version, defs)
	if err != nil {
		fmt.Println("error creating dRT, err: ", err)
		os.Exit(1)
	}
	var nodeEndpoints []gomavlib.EndpointConf
	if udp {
		nodeEndpoints = []gomavlib.EndpointConf{
			gomavlib.EndpointUdpServer{address},
		}
	} else {
		nodeEndpoints = []gomavlib.EndpointConf{
			gomavlib.EndpointTcpServer{address},
		}
	}
	node, err = gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints:   nodeEndpoints,
		D:           dRT,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to write to the target
		OutSystemId: 10,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	var nodeReceiveEndpoints []gomavlib.EndpointConf
	if udp {
		nodeReceiveEndpoints = []gomavlib.EndpointConf{
			gomavlib.EndpointUdpClient{address},
		}
	} else {
		nodeReceiveEndpoints = []gomavlib.EndpointConf{
			gomavlib.EndpointTcpClient{address},
		}
	}
	nodeReceive, err = gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints:   nodeReceiveEndpoints,
		D:           dRT,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to write to the target
		OutSystemId: 10,
	})
	if err != nil {
		panic(err)
	}
	defer nodeReceive.Close()

	if sendHeartbeat {
		heartbeatMsg, err = dRT.CreateMessageByName("HEARTBEAT")
		if err != nil {
			panic(err)
		}
	}
	if sendSysStatus {
		sysStatusMsg, err = dRT.CreateMessageByName("SYS_STATUS")
	}
	if sendAttitude {
		attitudeMsg, err = dRT.CreateMessageByName("ATTITUDE")
	}

	stopchan := make(chan struct{})
	stoppedchanHeartbeat := make(chan struct{})
	stoppedchanSysStatus := make(chan struct{})
	stoppedchanAttitude := make(chan struct{})
	stoppedchanReceive := make(chan struct{})
	stopchanReceive := make(chan struct{})

	heartbeatMessagesReceived := 0
	sysStatusMessagesReceived := 0
	attitudeMessagesReceived := 0

	countHeartbeat := 0
	countSysStatus := 0
	countAttitude := 0

	if sendHeartbeat {
		go func() { // Heartbeat
			defer close(stoppedchanHeartbeat)
			for t := range time.NewTicker(time.Duration(heartbeatPeriod) * time.Microsecond).C {
				select {
				default:
					err := randomHeartbeat(heartbeatMsg)
					if err != nil {
						panic(err)
					}
					node.WriteMessageAll(heartbeatMsg)
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
			for range time.NewTicker(time.Duration(sysStatusPeriod) * time.Microsecond).C {
				select {
				default:
					err := randomSysStatus(sysStatusMsg)
					if err != nil {
						panic(err)
					}
					node.WriteMessageAll(sysStatusMsg)
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
			for range time.NewTicker(time.Duration(attitudePeriod) * time.Microsecond).C {
				select {
				default:
					err := randomAttitude(attitudeMsg)
					if err != nil {
						panic(err)
					}
					node.WriteMessageAll(attitudeMsg)
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
					if msg, ok := frm.Message().(*gomavlib.DynamicMessage); ok {
						if msg.GetName() == "HEARTBEAT" {
							//fmt.Printf("received: id=%d, %+v\n", msg.GetId(), msg)
							heartbeatMessagesReceived++
						} else if msg.GetName() == "SYS_STATUS" {
							sysStatusMessagesReceived++
						} else if msg.GetName() == "ATTITUDE" {
							attitudeMessagesReceived++
						}
					}
				}
			case <-stopchanReceive:
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
	close(stopchanReceive)
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
