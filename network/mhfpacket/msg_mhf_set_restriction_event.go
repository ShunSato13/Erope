package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfSetRestrictionEvent represents the MSG_MHF_SET_RESTRICTION_EVENT
type MsgMhfSetRestrictionEvent struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSetRestrictionEvent) Opcode() network.PacketID {
	return network.MSG_MHF_SET_RESTRICTION_EVENT
}

// Parse parses the packet from binary
func (m *MsgMhfSetRestrictionEvent) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSetRestrictionEvent) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
