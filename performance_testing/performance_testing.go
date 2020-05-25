package performancetesting

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/team-rocos/gomavlib"
	"github.com/team-rocos/gomavlib/commands/dialgen/libgen"
)

var (
	address             string
	messages            []string
	stop                float64
	udp, useReceiveNode bool
	dRT                 *gomavlib.DialectRT
	dmMutex             = sync.RWMutex{}
	node, nodeReceive   *gomavlib.Node
)

func messageDetails(input []string) ([]*gomavlib.DynamicMessage, []int, error) {
	dm := make([]*gomavlib.DynamicMessage, 0)
	var periods []int
	messagePeriod := 0
	if len(input) == 0 {
		fmt.Println("input []string to messageDetails function must be greater than 0")
		os.Exit(1)
	}
	for _, msgString := range input {
		msg := strings.Split(msgString, "@")
		msgName := strings.ToUpper(msg[0])
		dmInstance, err := dRT.CreateMessageByName(msgName)
		if err != nil {
			return dm, periods, err
		}
		dm = append(dm, dmInstance)

		if len(msg) > 1 {
			if strings.Contains(msg[1], "us") { // period input in milliseconds
				period := strings.SplitN(msg[1], "us", 2)
				var err error
				messagePeriod, err = strconv.Atoi(period[0])
				if err != nil {
					return dm, periods, err
				}
				periods = append(periods, messagePeriod)
			} else if strings.Contains(msg[1], "ms") { // period input in milliseconds
				period := strings.SplitN(msg[1], "ms", 2)
				var err error
				messagePeriod, err = strconv.Atoi(period[0])
				messagePeriod *= 1000 // Concert from milliseconds to microseconds
				if err != nil {
					return dm, periods, err
				}
				periods = append(periods, messagePeriod)
			} else if strings.Contains(msg[1], "s") { // period input in seconds
				period := strings.SplitN(msg[1], "s", 2)
				var err error
				messagePeriod, err = strconv.Atoi(period[0])
				messagePeriod *= 1000000 // Convert from seconds to microseconds
				if err != nil {
					return dm, periods, err
				}
				periods = append(periods, messagePeriod)
			}
		}
	}
	return dm, periods, nil
}

func randomDynamicMessage(dm *gomavlib.DynamicMessage) error {
	rand.Seed(time.Now().UnixNano())
	for _, fieldInfo := range dm.T.Msg.Fields {
		fieldName := fieldInfo.OriginalName
		switch fieldInfo.Type {
		case "int8":
			if fieldInfo.ArrayLength != 0 {
				result := make([]int8, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = int8(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, int8(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "uint8":
			if fieldInfo.ArrayLength != 0 {
				result := make([]uint8, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = uint8(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, uint8(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "int16":
			if fieldInfo.ArrayLength != 0 {
				result := make([]int16, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = int16(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, int16(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "uint16":
			if fieldInfo.ArrayLength != 0 {
				result := make([]uint16, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = uint16(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, uint16(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "int32":
			if fieldInfo.ArrayLength != 0 {
				result := make([]int32, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = int32(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, int32(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "uint32":
			if fieldInfo.ArrayLength != 0 {
				result := make([]uint32, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = uint32(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, uint32(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "int64":
			if fieldInfo.ArrayLength != 0 {
				result := make([]int64, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = int64(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, int64(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "uint64":
			if fieldInfo.ArrayLength != 0 {
				result := make([]uint64, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = uint64(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, uint64(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "float64":
			if fieldInfo.ArrayLength != 0 {
				result := make([]float64, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = float64(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, float64(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "float32":
			if fieldInfo.ArrayLength != 0 {
				result := make([]float32, fieldInfo.ArrayLength)
				for i := 0; i < fieldInfo.ArrayLength; i++ {
					result[i] = float32(rand.Intn(100))
				}
				err := dm.SetField(fieldName, result)
				if err != nil {
					return err
				}
			} else {
				err := dm.SetField(fieldName, float32(rand.Intn(100)))
				if err != nil {
					return err
				}
			}
		case "string":
			if fieldInfo.ArrayLength == 0 {
				return errors.New("DynamicMessage string field has an array lenght of 0")
			}
			result := ""
			for i := 0; i < fieldInfo.ArrayLength; i++ {
				if i%3 == 0 {
					result += "c"
				} else if i%2 == 0 {
					result += "b"
				} else {
					result += "a"
				}
			}
			dm.SetField(fieldName, result)
		default:
			return errors.New("unsupported field type in dynamic MAVLink message")
		}
	}
	return nil
}

// func main() {
// 	pflag.StringVarP(&address, "address", "a", "", "Set address to which to send MAVlink messages.")
// 	pflag.StringSliceVarP(&messages, "msg", "m", []string{""}, "Set message(s) to be sent as well and the period at which they are sent.")
// 	pflag.Float64VarP(&stop, "stop", "s", 10, "Set the length of time in seconds to send messages. Up to millisecond resolution (3dp).")
// 	pflag.BoolVarP(&udp, "udp", "u", false, "Set this to send to a udp port. Otherwise defaults to tcp")
// 	pflag.BoolVarP(&useReceiveNode, "receive", "r", false, "Set this to set up a receive node, which can be used to confirm the number of messages sent successfully.")
// 	pflag.Parse()

// 	err := MessageSender(address, messages, stop, udp, "../mavlink-upstream/message_definitions/v1.0/common.xml", []string{"."}, useReceiveNode)
// 	if err != nil {
// 		fmt.Println("error sending messages in MessageSender, err: ", err)
// 		os.Exit(1)
// 	}
// }

// MessageSender sends MAVLink messages to the specified address, at the specified frequency, until the stop timeout value is reached.
// The messages send have randomised field values.
// Inputs:
// - address: A string showing the ip:port address, e.g. "192.168.1.84:5600"
// - messages: Slice of strings specifying the MAVLink message name, and the period between messages sent separated with '@', e.g. "ATTITUDE@10ms".
// Period units of seconds (s), milliseconds (ms), or microseconds (us) can be specified.
// - stop: A float64 value giving the amount of time in seconds that messages should be sent. E.g, 10, would stop MessageSender after 10 seconds.
// udp: A boolean variable. If set to true, the address spcified will be interpreted as udp. If false, it will be interpreted as a tcp address.
// xmlPath: A string showing the path to the dialect XML file used to create the MAVLink dialect.
// includeDirs: A string slice providing any applicable xml files included by the dialect XML file specified in the xmlPath input.
// receiveNode: Boolean variable. If true, a local receive node is created which can be used to confirm the number of messages sent successfully.
func MessageSender(address string, messages []string, stop float64, udp bool, xmlPath string, includeDirs []string, useReceiveNode bool) error {

	addressType := "TCP"
	if udp {
		addressType = "UDP"
	}
	fmt.Println("Setting address to: ", addressType, address)
	fmt.Println("Setting timeout to: ", stop)

	defs, version, err := libgen.XMLToFields(xmlPath, includeDirs)
	if err != nil {
		return err
	}
	// Create dialect from the parsed defs.
	dRT, err = gomavlib.NewDialectRT(version, defs)
	if err != nil {
		return err
	}

	// Create nodes for sending and receiving messages
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
		return err
	}
	defer node.Close()

	if useReceiveNode {
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
			return err
		}
		defer nodeReceive.Close()
	}
	// Create dynamicMessage
	dynamicMsgSlice, periodSlice, err := messageDetails(messages)
	if err != nil {
		return err
	}
	for i, msg := range dynamicMsgSlice {
		fmt.Printf("Sending %v at %v Hz\n", msg.GetName(), 1000000/periodSlice[i])
	}

	stopchan := make(chan struct{})
	stoppedchanReceive := make(chan struct{})
	stopchanReceive := make(chan struct{})

	sentMessages := make(map[string]int)
	for _, m := range dynamicMsgSlice {
		sentMessages[m.GetName()] = 0
	}

	receivedMessages := make(map[string]int)
	for _, m := range dynamicMsgSlice {
		receivedMessages[m.GetName()] = 0
	}

	stoppedSendMessageChannels := make(map[string]chan struct{})
	for _, m := range dynamicMsgSlice {
		stoppedSendMessageChannels[m.GetName()] = make(chan struct{})
	}

	if useReceiveNode {
		go func() { // Receive
			fmt.Println("Starting go routine to recieve all messages")
			defer close(stoppedchanReceive)
			// print every message we receive
			for evt := range nodeReceive.Events() {
				select {
				default:
					if frm, ok := evt.(*gomavlib.EventFrame); ok {
						dmMutex.Lock()
						if msg, ok := frm.Message().(*gomavlib.DynamicMessage); ok {
							name := msg.GetName()
							receivedMessages[name]++
						}
						dmMutex.Unlock()
					}
				case <-stopchanReceive:
					fmt.Println("Closing nodeReceive go routine...")
					return
				}
			}
		}()
	}

	// Send DynamicMessages, creating a go routine for each topic.
	for i, msgToSend := range dynamicMsgSlice {
		msg := msgToSend
		index := i
		dmMutex.Lock()
		msgName := msg.GetName()
		dmMutex.Unlock()

		go func() {
			fmt.Printf("Starting go routine to send %v\n", msgName)
			period := periodSlice[index]
			defer close(stoppedSendMessageChannels[msgName])
			for range time.NewTicker(time.Duration(period) * time.Microsecond).C {
				select {
				default:
					dmMutex.Lock()
					err := randomDynamicMessage(msg)
					if err != nil {
						fmt.Println("error creating random values for DynamicMessage fields: err", err)
						os.Exit(1)
					}
					node.WriteMessageAll(msg)
					dmMutex.Unlock()

					sentMessages[msgName]++

				case <-stopchan:
					fmt.Printf("Closing send %v go routine...\n", msgName)
					return
				}
			}
		}()
	}

	time.Sleep(time.Duration(stop*1000) * time.Millisecond)

	// Close send go routine
	close(stopchan)
	for _, m := range dynamicMsgSlice {
		channel := stoppedSendMessageChannels[m.GetName()]
		<-channel // Wait for channel to close
	}

	// Now close receive go routine
	if useReceiveNode {
		close(stopchanReceive)
		<-stoppedchanReceive // Wait for channel to close
	}
	fmt.Println("Closed all go routines!")

	// Print counts
	for _, m := range dynamicMsgSlice {
		fmt.Printf("%v Messages Sent    : %v\n", m.GetName(), sentMessages[m.GetName()])
		if useReceiveNode {
			fmt.Printf("%v Messages Received: %v\n", m.GetName(), receivedMessages[m.GetName()])
		}
	}

	// All done, so return no errors.
	return nil
}
