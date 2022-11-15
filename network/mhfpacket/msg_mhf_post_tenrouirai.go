package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfPostTenrouirai represents the MSG_MHF_POST_TENROUIRAI
type MsgMhfPostTenrouirai struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfPostTenrouirai) Opcode() network.PacketID {
	return network.MSG_MHF_POST_TENROUIRAI
}

// Parse parses the packet from binary
func (m *MsgMhfPostTenrouirai) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfPostTenrouirai) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
