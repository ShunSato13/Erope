package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfReserve010F represents the MSG_MHF_reserve010F
type MsgMhfReserve010F struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfReserve010F) Opcode() network.PacketID {
	return network.MSG_MHF_reserve010F
}

// Parse parses the packet from binary
func (m *MsgMhfReserve010F) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfReserve010F) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
