package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfOperationInvGuild represents the MSG_MHF_OPERATION_INV_GUILD
type MsgMhfOperationInvGuild struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfOperationInvGuild) Opcode() network.PacketID {
	return network.MSG_MHF_OPERATION_INV_GUILD
}

// Parse parses the packet from binary
func (m *MsgMhfOperationInvGuild) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfOperationInvGuild) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
