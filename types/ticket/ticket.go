package ticket

import (
	"time"

	"github.com/postie-labs/go-tickets/types"
)

type TicketType uint8

const (
	TicketTypeSingleOwner TicketType = iota
	TicketTypeMultiOwner
)

type TicketProtocolVersion uint16

type Ticket struct {
	// metadata
	Hash            types.Hash
	ProtocolVersion TicketProtocolVersion
	Signature       types.Signature // signature of issuer

	// data
	Timestamp time.Time
	Issuer    types.Address
	Type      TicketType
	Data      types.Data
}
