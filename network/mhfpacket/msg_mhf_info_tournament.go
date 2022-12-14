package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfInfoTournament represents the MSG_MHF_INFO_TOURNAMENT
type MsgMhfInfoTournament struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfInfoTournament) Opcode() network.PacketID {
	return network.MSG_MHF_INFO_TOURNAMENT
}

// Parse parses the packet from binary
func (m *MsgMhfInfoTournament) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfInfoTournament) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
