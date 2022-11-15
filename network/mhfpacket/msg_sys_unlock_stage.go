package mhfpacket

import (
	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgSysUnlockStage represents the MSG_SYS_UNLOCK_STAGE
type MsgSysUnlockStage struct {
	Unk0 uint16 // Hardcoded 0 in the binary.
}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysUnlockStage) Opcode() network.PacketID {
	return network.MSG_SYS_UNLOCK_STAGE
}

// Parse parses the packet from binary
func (m *MsgSysUnlockStage) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.Unk0 = bf.ReadUint16()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgSysUnlockStage) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	bf.WriteUint16(m.Unk0)
	return nil
}
