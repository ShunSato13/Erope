package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgSysReserve78 represents the MSG_SYS_reserve78
type MsgSysReserve78 struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve78) Opcode() network.PacketID {
	return network.MSG_SYS_reserve78
}

// Parse parses the packet from binary
func (m *MsgSysReserve78) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve78) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
