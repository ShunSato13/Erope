package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgSysReserve0C represents the MSG_SYS_reserve0C
type MsgSysReserve0C struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve0C) Opcode() network.PacketID {
	return network.MSG_SYS_reserve0C
}

// Parse parses the packet from binary
func (m *MsgSysReserve0C) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve0C) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
