package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysLockGlobalSema represents the MSG_SYS_LOCK_GLOBAL_SEMA
type MsgSysLockGlobalSema struct {
	AckHandle             uint32
	UnkIDString0          string
	ServerChannelIDString string
}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysLockGlobalSema) Opcode() network.PacketID {
	return network.MSG_SYS_LOCK_GLOBAL_SEMA
}

// Parse parses the packet from binary
func (m *MsgSysLockGlobalSema) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()

	stageNameLength := bf.ReadUint16()
	serverIDLength := bf.ReadUint16()

	m.UnkIDString0 = string(bf.ReadBytes(uint(stageNameLength)))
	m.ServerChannelIDString = string(bf.ReadBytes(uint(serverIDLength)))

	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgSysLockGlobalSema) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
