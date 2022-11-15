package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfSetKiju represents the MSG_MHF_SET_KIJU
type MsgMhfSetKiju struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSetKiju) Opcode() network.PacketID {
	return network.MSG_MHF_SET_KIJU
}

// Parse parses the packet from binary
func (m *MsgMhfSetKiju) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSetKiju) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
