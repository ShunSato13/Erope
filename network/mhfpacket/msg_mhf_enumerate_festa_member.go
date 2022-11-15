package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfEnumerateFestaMember represents the MSG_MHF_ENUMERATE_FESTA_MEMBER
type MsgMhfEnumerateFestaMember struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEnumerateFestaMember) Opcode() network.PacketID {
	return network.MSG_MHF_ENUMERATE_FESTA_MEMBER
}

// Parse parses the packet from binary
func (m *MsgMhfEnumerateFestaMember) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEnumerateFestaMember) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
