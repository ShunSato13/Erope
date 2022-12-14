package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfStartBoostTime represents the MSG_MHF_START_BOOST_TIME
type MsgMhfStartBoostTime struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfStartBoostTime) Opcode() network.PacketID {
	return network.MSG_MHF_START_BOOST_TIME
}

// Parse parses the packet from binary
func (m *MsgMhfStartBoostTime) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfStartBoostTime) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
