package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfGetUdTacticsRemainingPoint represents the MSG_MHF_GET_UD_TACTICS_REMAINING_POINT
type MsgMhfGetUdTacticsRemainingPoint struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetUdTacticsRemainingPoint) Opcode() network.PacketID {
	return network.MSG_MHF_GET_UD_TACTICS_REMAINING_POINT
}

// Parse parses the packet from binary
func (m *MsgMhfGetUdTacticsRemainingPoint) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetUdTacticsRemainingPoint) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
