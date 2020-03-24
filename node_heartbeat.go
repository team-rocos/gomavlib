package gomavlib

import (
	"time"
)

type nodeHeartbeat struct {
	n         *Node
	terminate chan struct{}
}

func newNodeHeartbeat(n *Node) *nodeHeartbeat {
	// module is disabled
	if n.conf.HeartbeatDisable == true {
		return nil
	}

	// dialect must be enabled
	if n.conf.D == nil {
		return nil
	}

	// heartbeat message must exist in dialect and correspond to standard
	mp, ok := (n.conf.D).getMsgById(0)
	if ok == false || (*mp).getCRCExtra() != 50 {
		return nil
	}

	h := &nodeHeartbeat{
		n:         n,
		terminate: make(chan struct{}, 1),
	}

	return h
}

func (h *nodeHeartbeat) close() {
	h.terminate <- struct{}{}
}

func (h *nodeHeartbeat) run() {
	// take version from dialect if possible
	mavlinkVersion := uint64(3)
	if h.n.conf.D != nil {
		mavlinkVersion = uint64(h.n.conf.D.getVersion())
	}

	ticker := time.NewTicker(h.n.conf.HeartbeatPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			dm, _ := h.n.conf.D.getMsgById(0)
			msg := (*dm).newMsg()
			msg.SetField("Type", int8(h.n.conf.HeartbeatSystemType))
			msg.SetField("Autopilot", int8(h.n.conf.HeartbeatAutopilotType))
			msg.SetField("BaseMode", int8(0))
			msg.SetField("CustomMode", uint32(0))
			msg.SetField("SystemStatus", int8(4)) // MAV_STATE_ACTIVE
			msg.SetField("MavlinkVersion", uint8(mavlinkVersion))
			h.n.WriteMessageAll(msg)

		case <-h.terminate:
			return
		}
	}
}
