package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfStateFestaU represents the MSG_MHF_STATE_FESTA_U
type MsgMhfStateFestaU struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfStateFestaU) Opcode() network.PacketID {
	return network.MSG_MHF_STATE_FESTA_U
}

// Parse parses the packet from binary
func (m *MsgMhfStateFestaU) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfStateFestaU) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
