package mhfpacket

import (
	"errors"

	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfOperateWarehouse represents the MSG_MHF_OPERATE_WAREHOUSE
type MsgMhfOperateWarehouse struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfOperateWarehouse) Opcode() network.PacketID {
	return network.MSG_MHF_OPERATE_WAREHOUSE
}

// Parse parses the packet from binary
func (m *MsgMhfOperateWarehouse) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfOperateWarehouse) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}
