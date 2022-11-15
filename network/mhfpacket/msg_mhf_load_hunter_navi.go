package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfLoadHunterNavi represents the MSG_MHF_LOAD_HUNTER_NAVI
type MsgMhfLoadHunterNavi struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfLoadHunterNavi) Opcode() network.PacketID {
	return network.MSG_MHF_LOAD_HUNTER_NAVI
}

// Parse parses the packet from binary
func (m *MsgMhfLoadHunterNavi) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfLoadHunterNavi) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
