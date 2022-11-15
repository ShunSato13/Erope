package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgSysReserve04 represents the MSG_SYS_reserve04
type MsgSysReserve04 struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve04) Opcode() network.PacketID {
	return network.MSG_SYS_reserve04
}

// Parse parses the packet from binary
func (m *MsgSysReserve04) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve04) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
