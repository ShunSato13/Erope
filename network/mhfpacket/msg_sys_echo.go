package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysEcho represents the MSG_SYS_ECHO
type MsgSysEcho struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysEcho) Opcode() network.PacketID {
	return network.MSG_SYS_ECHO
}

// Parse parses the packet from binary
func (m *MsgSysEcho) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysEcho) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
