package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfGetCaUniqueID represents the MSG_MHF_GET_CA_UNIQUE_ID
type MsgMhfGetCaUniqueID struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetCaUniqueID) Opcode() network.PacketID {
	return network.MSG_MHF_GET_CA_UNIQUE_ID
}

// Parse parses the packet from binary
func (m *MsgMhfGetCaUniqueID) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetCaUniqueID) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
