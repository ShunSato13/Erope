package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgCaExchangeItem represents the MSG_CA_EXCHANGE_ITEM
type MsgCaExchangeItem struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgCaExchangeItem) Opcode() network.PacketID {
	return network.MSG_CA_EXCHANGE_ITEM
}

// Parse parses the packet from binary
func (m *MsgCaExchangeItem) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgCaExchangeItem) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
