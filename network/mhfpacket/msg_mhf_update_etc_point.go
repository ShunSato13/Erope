package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateEtcPoint represents the MSG_MHF_UPDATE_ETC_POINT
type MsgMhfUpdateEtcPoint struct {
	AckHandle uint32
	Unk0      uint8
	Unk1      uint16
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateEtcPoint) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_ETC_POINT
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateEtcPoint) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint8()
	m.Unk1 = bf.ReadUint16()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateEtcPoint) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
