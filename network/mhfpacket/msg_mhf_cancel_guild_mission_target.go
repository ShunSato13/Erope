package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfCancelGuildMissionTarget represents the MSG_MHF_CANCEL_GUILD_MISSION_TARGET
type MsgMhfCancelGuildMissionTarget struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfCancelGuildMissionTarget) Opcode() network.PacketID {
	return network.MSG_MHF_CANCEL_GUILD_MISSION_TARGET
}

// Parse parses the packet from binary
func (m *MsgMhfCancelGuildMissionTarget) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfCancelGuildMissionTarget) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
