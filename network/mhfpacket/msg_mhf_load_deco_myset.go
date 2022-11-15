package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfLoadDecoMyset represents the MSG_MHF_LOAD_DECO_MYSET
type MsgMhfLoadDecoMyset struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfLoadDecoMyset) Opcode() network.PacketID {
	return network.MSG_MHF_LOAD_DECO_MYSET
}

// Parse parses the packet from binary
func (m *MsgMhfLoadDecoMyset) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfLoadDecoMyset) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
