package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUseGachaPoint represents the MSG_MHF_USE_GACHA_POINT
type MsgMhfUseGachaPoint struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUseGachaPoint) Opcode() network.PacketID {
	return network.MSG_MHF_USE_GACHA_POINT
}

// Parse parses the packet from binary
func (m *MsgMhfUseGachaPoint) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUseGachaPoint) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
