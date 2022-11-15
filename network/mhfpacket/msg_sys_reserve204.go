package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgSysReserve204 represents the MSG_SYS_reserve204
type MsgSysReserve204 struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve204) Opcode() network.PacketID {
	return network.MSG_SYS_reserve204
}

// Parse parses the packet from binary
func (m *MsgSysReserve204) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve204) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
