package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfAcquireFestaIntermediatePrize represents the MSG_MHF_ACQUIRE_FESTA_INTERMEDIATE_PRIZE
type MsgMhfAcquireFestaIntermediatePrize struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireFestaIntermediatePrize) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_FESTA_INTERMEDIATE_PRIZE
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireFestaIntermediatePrize) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireFestaIntermediatePrize) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
