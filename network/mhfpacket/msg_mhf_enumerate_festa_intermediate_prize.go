package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfEnumerateFestaIntermediatePrize represents the MSG_MHF_ENUMERATE_FESTA_INTERMEDIATE_PRIZE
type MsgMhfEnumerateFestaIntermediatePrize struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEnumerateFestaIntermediatePrize) Opcode() network.PacketID {
	return network.MSG_MHF_ENUMERATE_FESTA_INTERMEDIATE_PRIZE
}

// Parse parses the packet from binary
func (m *MsgMhfEnumerateFestaIntermediatePrize) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEnumerateFestaIntermediatePrize) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
