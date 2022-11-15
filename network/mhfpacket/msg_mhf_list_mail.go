package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// MsgMhfListMail represents the MSG_MHF_LIST_MAIL
type MsgMhfListMail struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfListMail) Opcode() network.PacketID {
	return network.MSG_MHF_LIST_MAIL
}

// Parse parses the packet from binary
func (m *MsgMhfListMail) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfListMail) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
