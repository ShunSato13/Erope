package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfExchangeWeeklyStamp represents the MSG_MHF_EXCHANGE_WEEKLY_STAMP
type MsgMhfExchangeWeeklyStamp struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfExchangeWeeklyStamp) Opcode() network.PacketID {
	return network.MSG_MHF_EXCHANGE_WEEKLY_STAMP
}

// Parse parses the packet from binary
func (m *MsgMhfExchangeWeeklyStamp) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfExchangeWeeklyStamp) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
