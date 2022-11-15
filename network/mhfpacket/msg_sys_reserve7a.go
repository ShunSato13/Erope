package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgSysReserve7A represents the MSG_SYS_reserve7A
type MsgSysReserve7A struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve7A) Opcode() network.PacketID {
	return network.MSG_SYS_reserve7A
}

// Parse parses the packet from binary
func (m *MsgSysReserve7A) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve7A) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
