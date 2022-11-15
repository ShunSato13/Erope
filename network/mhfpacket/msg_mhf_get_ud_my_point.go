package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfGetUdMyPoint represents the MSG_MHF_GET_UD_MY_POINT
type MsgMhfGetUdMyPoint struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetUdMyPoint) Opcode() network.PacketID {
	return network.MSG_MHF_GET_UD_MY_POINT
}

// Parse parses the packet from binary
func (m *MsgMhfGetUdMyPoint) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetUdMyPoint) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
