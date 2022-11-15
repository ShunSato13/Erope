package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfGetBoxGachaInfo represents the MSG_MHF_GET_BOX_GACHA_INFO
type MsgMhfGetBoxGachaInfo struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetBoxGachaInfo) Opcode() network.PacketID {
	return network.MSG_MHF_GET_BOX_GACHA_INFO
}

// Parse parses the packet from binary
func (m *MsgMhfGetBoxGachaInfo) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetBoxGachaInfo) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
