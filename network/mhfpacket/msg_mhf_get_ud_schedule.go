package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetUdSchedule represents the MSG_MHF_GET_UD_SCHEDULE
type MsgMhfGetUdSchedule struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetUdSchedule) Opcode() network.PacketID {
	return network.MSG_MHF_GET_UD_SCHEDULE
}

// Parse parses the packet from binary
func (m *MsgMhfGetUdSchedule) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetUdSchedule) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
