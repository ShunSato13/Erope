package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetCaAchievementHist represents the MSG_MHF_GET_CA_ACHIEVEMENT_HIST
type MsgMhfGetCaAchievementHist struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetCaAchievementHist) Opcode() network.PacketID {
	return network.MSG_MHF_GET_CA_ACHIEVEMENT_HIST
}

// Parse parses the packet from binary
func (m *MsgMhfGetCaAchievementHist) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetCaAchievementHist) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
