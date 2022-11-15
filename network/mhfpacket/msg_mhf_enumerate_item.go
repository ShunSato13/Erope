package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfEnumerateItem represents the MSG_MHF_ENUMERATE_ITEM
type MsgMhfEnumerateItem struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEnumerateItem) Opcode() network.PacketID {
	return network.MSG_MHF_ENUMERATE_ITEM
}

// Parse parses the packet from binary
func (m *MsgMhfEnumerateItem) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEnumerateItem) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
