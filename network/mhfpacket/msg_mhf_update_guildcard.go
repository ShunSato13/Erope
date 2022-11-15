package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfUpdateGuildcard represents the MSG_MHF_UPDATE_GUILDCARD
type MsgMhfUpdateGuildcard struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateGuildcard) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_GUILDCARD
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateGuildcard) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateGuildcard) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
