package ticket

import (
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
	Timestamp types.Timestamp
	Issuer    types.Address
	Type      TicketType
	Data      types.Data
}

func NewTicket(protocolVersion TicketProtocolVersion, issuer types.Address, ticketType TicketType, data types.Data) *Ticket {
	return &Ticket{
		Hash:            types.Hash{},
		ProtocolVersion: protocolVersion,
		Signature:       make(types.Signature, 0),

		Timestamp: types.TimestampNow(),
		Issuer:    issuer,
		Type:      ticketType,
		Data:      data,
	}
}

// ops
func (t *Ticket) GenerateHash() error {
	return nil
}

func (t *Ticket) Sign() error {
	return nil
}

func (t *Ticket) Verify() error {
	return nil
}

// accessors
func (t *Ticket) GetHash() types.Hash {
	return t.Hash
}

func (t *Ticket) GetProtocolVersion() TicketProtocolVersion {
	return t.ProtocolVersion
}

func (t *Ticket) GetSignature() types.Signature {
	return t.Signature
}

func (t *Ticket) GetTimestamp() types.Timestamp {
	return t.Timestamp
}

func (t *Ticket) GetIssuer() types.Address {
	return t.Issuer
}

func (t *Ticket) GetType() TicketType {
	return t.Type
}

func (t *Ticket) GetData() types.Data {
	return t.Data
}
