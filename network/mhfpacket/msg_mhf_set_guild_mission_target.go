package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfSetGuildMissionTarget represents the MSG_MHF_SET_GUILD_MISSION_TARGET
type MsgMhfSetGuildMissionTarget struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSetGuildMissionTarget) Opcode() network.PacketID {
	return network.MSG_MHF_SET_GUILD_MISSION_TARGET
}

// Parse parses the packet from binary
func (m *MsgMhfSetGuildMissionTarget) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSetGuildMissionTarget) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
