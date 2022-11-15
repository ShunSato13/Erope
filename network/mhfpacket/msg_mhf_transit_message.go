package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfTransitMessage represents the MSG_MHF_TRANSIT_MESSAGE
type MsgMhfTransitMessage struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfTransitMessage) Opcode() network.PacketID {
	return network.MSG_MHF_TRANSIT_MESSAGE
}

// Parse parses the packet from binary
func (m *MsgMhfTransitMessage) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfTransitMessage) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
