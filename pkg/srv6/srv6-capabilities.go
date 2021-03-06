package srv6

import (
	"encoding/binary"

	"github.com/golang/glog"
	"github.com/sbezverk/gobmp/pkg/tools"
)

// CapabilityTLV defines SRv6 Capability TLV object
// No RFC yet
type CapabilityTLV struct {
	Flag uint16 `json:"flag"`
}

// UnmarshalSRv6CapabilityTLV builds SRv6 Capability TLV object
func UnmarshalSRv6CapabilityTLV(b []byte) (*CapabilityTLV, error) {
	if glog.V(6) {
		glog.Infof("SRv6 Capability TLV Raw: %s", tools.MessageHex(b))
	}
	cap := CapabilityTLV{}
	p := 0
	cap.Flag = binary.BigEndian.Uint16(b[p : p+2])

	return &cap, nil
}
