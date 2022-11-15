package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve7B represents the MSG_SYS_reserve7B
type MsgSysReserve7B struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve7B) Opcode() network.PacketID {
	return network.MSG_SYS_reserve7B
}

// Parse parses the packet from binary
func (m *MsgSysReserve7B) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve7B) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
