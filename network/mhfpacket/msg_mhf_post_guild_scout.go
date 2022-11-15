package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfPostGuildScout represents the MSG_MHF_POST_GUILD_SCOUT
type MsgMhfPostGuildScout struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfPostGuildScout) Opcode() network.PacketID {
	return network.MSG_MHF_POST_GUILD_SCOUT
}

// Parse parses the packet from binary
func (m *MsgMhfPostGuildScout) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfPostGuildScout) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
