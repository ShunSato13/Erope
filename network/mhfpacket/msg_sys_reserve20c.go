package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve20C represents the MSG_SYS_reserve20C
type MsgSysReserve20C struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve20C) Opcode() network.PacketID {
	return network.MSG_SYS_reserve20C
}

// Parse parses the packet from binary
func (m *MsgSysReserve20C) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve20C) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
