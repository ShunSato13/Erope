package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysSetStatus represents the MSG_SYS_SET_STATUS
type MsgSysSetStatus struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysSetStatus) Opcode() network.PacketID {
	return network.MSG_SYS_SET_STATUS
}

// Parse parses the packet from binary
func (m *MsgSysSetStatus) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysSetStatus) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
