package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfExchangeFpoint2Item represents the MSG_MHF_EXCHANGE_FPOINT_2_ITEM
type MsgMhfExchangeFpoint2Item struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfExchangeFpoint2Item) Opcode() network.PacketID {
	return network.MSG_MHF_EXCHANGE_FPOINT_2_ITEM
}

// Parse parses the packet from binary
func (m *MsgMhfExchangeFpoint2Item) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfExchangeFpoint2Item) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
