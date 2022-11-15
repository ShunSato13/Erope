package mhfpacket

import (
	"github.com/ShunSato13/Erupe/network"
	"github.com/ShunSato13/Erupe/network/clientctx"
	"github.com/ShunSato13/byteframe"
)

// Parser is the interface that wraps the Parse method.
type Parser interface {
	Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error
}

// Builder is the interface that wraps the Build method.
type Builder interface {
	Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error
}

// Opcoder is the interface that wraps the Opcode method.
type Opcoder interface {
	Opcode() network.PacketID
}

// MHFPacket is the interface that groups the Parse, Build, and Opcode methods.
type MHFPacket interface {
	Parser
	Builder
	Opcoder
}
