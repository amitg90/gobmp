package bgp

import (
	"reflect"
	"testing"
)

func TestUnmarshaBaseAttributes(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		expect *BaseAttributes
	}{
		{
			name:  "panic 1",
			input: []byte{0x40, 0x01, 0x01, 0x00, 0x40, 0x02, 0x20, 0x02, 0x06, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x9a, 0x6d, 0x00, 0x00, 0x19, 0x35, 0x00, 0x00, 0x0a, 0x7f, 0x00, 0x00, 0x65, 0x20, 0x00, 0x00, 0x53, 0x4e, 0x01, 0x01, 0x00, 0x00, 0x12, 0xc9, 0x40, 0x03, 0x04, 0xc2, 0x1c, 0x62, 0x25, 0x80, 0x04, 0x04, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x07, 0x08, 0x00, 0x00, 0x65, 0x20, 0xc0, 0x78, 0x51, 0x88, 0xc0, 0x08, 0x18, 0x00, 0x00, 0x9a, 0x6d, 0x19, 0x35, 0x00, 0x56, 0x19, 0x35, 0x0b, 0xb8, 0x19, 0x35, 0x0c, 0x1c, 0x19, 0x35, 0x0c, 0x1e, 0x9a, 0x6d, 0xc2, 0x02, 0xc0, 0x20, 0x30, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0xd3, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x00, 0x64, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x88, 0x38, 0x00, 0x00, 0x00, 0x7a, 0x00, 0x00, 0x00, 0x01},
			expect: &BaseAttributes{
				BaseAttrHash:    "354de7a76afc292b187aa3ea32aa76b9",
				Origin:          "igp",
				ASPath:          []uint32{34872, 39533, 6453, 2687, 25888, 21326, 4809},
				ASPathCount:     7,
				Nexthop:         "194.28.98.37",
				Aggregator:      []byte{0, 0, 101, 32, 192, 120, 81, 136},
				CommunityList:   "0:39533, 6453:86, 6453:3000, 6453:3100, 6453:3102, 39533:49666",
				LgCommunityList: "34872:10:211, 34872:11:1, 34872:100:49, 34872:122:1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalBGPBaseAttributes(tt.input)
			if err != nil {
				t.Fatalf("expected to succeed but failed with error: %+v", err)
			}
			if err == nil {
				if !reflect.DeepEqual(got, tt.expect) {
					t.Errorf("Expected extCommunity %+v does not match to actual extCommunity %+v", tt.expect, got)
				}
			}
		})
	}
}