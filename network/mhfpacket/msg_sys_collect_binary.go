package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysCollectBinary represents the MSG_SYS_COLLECT_BINARY
type MsgSysCollectBinary struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysCollectBinary) Opcode() network.PacketID {
	return network.MSG_SYS_COLLECT_BINARY
}

// Parse parses the packet from binary
func (m *MsgSysCollectBinary) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysCollectBinary) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
