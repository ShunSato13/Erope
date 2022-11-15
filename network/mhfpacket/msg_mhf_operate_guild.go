package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

type OperateGuildAction uint8

const (
	OPERATE_GUILD_ACTION_DISBAND        = 0x01
	OPERATE_GUILD_ACTION_APPLY          = 0x02
	OPERATE_GUILD_ACTION_LEAVE          = 0x03
	OPERATE_GUILD_ACTION_UPDATE_COMMENT = 0x09
	OPERATE_GUILD_ACTION_DONATE         = 0x0a
)

// MsgMhfOperateGuild represents the MSG_MHF_OPERATE_GUILD
type MsgMhfOperateGuild struct {
	AckHandle uint32
	GuildID   uint32
	Action    OperateGuildAction
	UnkData   []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfOperateGuild) Opcode() network.PacketID {
	return network.MSG_MHF_OPERATE_GUILD
}

// Parse parses the packet from binary
func (m *MsgMhfOperateGuild) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.GuildID = bf.ReadUint32()
	m.Action = OperateGuildAction(bf.ReadUint8())
	m.UnkData = bf.DataFromCurrent()

	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfOperateGuild) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
