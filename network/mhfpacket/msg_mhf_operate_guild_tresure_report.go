package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfOperateGuildTresureReport represents the MSG_MHF_OPERATE_GUILD_TRESURE_REPORT
type MsgMhfOperateGuildTresureReport struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfOperateGuildTresureReport) Opcode() network.PacketID {
	return network.MSG_MHF_OPERATE_GUILD_TRESURE_REPORT
}

// Parse parses the packet from binary
func (m *MsgMhfOperateGuildTresureReport) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfOperateGuildTresureReport) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
