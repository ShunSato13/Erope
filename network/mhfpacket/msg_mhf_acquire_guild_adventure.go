package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfAcquireGuildAdventure represents the MSG_MHF_ACQUIRE_GUILD_ADVENTURE
type MsgMhfAcquireGuildAdventure struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireGuildAdventure) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_GUILD_ADVENTURE
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireGuildAdventure) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireGuildAdventure) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
