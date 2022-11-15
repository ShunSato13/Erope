package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgSysReserve19E represents the MSG_SYS_reserve19E
type MsgSysReserve19E struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve19E) Opcode() network.PacketID {
	return network.MSG_SYS_reserve19E
}

// Parse parses the packet from binary
func (m *MsgSysReserve19E) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve19E) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
