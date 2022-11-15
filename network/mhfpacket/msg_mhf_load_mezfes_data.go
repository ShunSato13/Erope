package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfLoadMezfesData represents the MSG_MHF_LOAD_MEZFES_DATA
type MsgMhfLoadMezfesData struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfLoadMezfesData) Opcode() network.PacketID {
	return network.MSG_MHF_LOAD_MEZFES_DATA
}

// Parse parses the packet from binary
func (m *MsgMhfLoadMezfesData) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfLoadMezfesData) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
