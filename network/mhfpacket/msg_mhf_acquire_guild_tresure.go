package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireGuildTresure represents the MSG_MHF_ACQUIRE_GUILD_TRESURE
type MsgMhfAcquireGuildTresure struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireGuildTresure) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_GUILD_TRESURE
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireGuildTresure) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireGuildTresure) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
