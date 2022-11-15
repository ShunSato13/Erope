package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfPlayNormalGacha represents the MSG_MHF_PLAY_NORMAL_GACHA
type MsgMhfPlayNormalGacha struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfPlayNormalGacha) Opcode() network.PacketID {
	return network.MSG_MHF_PLAY_NORMAL_GACHA
}

// Parse parses the packet from binary
func (m *MsgMhfPlayNormalGacha) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfPlayNormalGacha) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
