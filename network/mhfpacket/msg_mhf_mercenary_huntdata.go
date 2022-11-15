package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfMercenaryHuntdata represents the MSG_MHF_MERCENARY_HUNTDATA
type MsgMhfMercenaryHuntdata struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfMercenaryHuntdata) Opcode() network.PacketID {
	return network.MSG_MHF_MERCENARY_HUNTDATA
}

// Parse parses the packet from binary
func (m *MsgMhfMercenaryHuntdata) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfMercenaryHuntdata) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
