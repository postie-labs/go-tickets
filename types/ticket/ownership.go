package ticket

import (
	crypto "github.com/postie-labs/go-crypto-lib"
	"github.com/postie-labs/go-tickets/types"
)

type OwnershipProtocolVersion uint16

type Ownership struct {
	// metadata
	Hash            types.Hash
	ProtocolVersion OwnershipProtocolVersion
	Signatures      []types.Signature // signature of owners

	// data
	Timestamp  types.Timestamp
	TicketHash types.Hash
	Owners     []crypto.Addr
}

func NewOwnership(protocolVersion OwnershipProtocolVersion, ticketHash types.Hash, owners []crypto.Addr) *Ownership {
	return &Ownership{
		Hash:            types.Hash{},
		ProtocolVersion: protocolVersion,
		Signatures:      make([]types.Signature, 0),

		Timestamp:  types.TimestampNow(),
		TicketHash: ticketHash,
		Owners:     owners,
	}
}

// ops
func (o *Ownership) GenerateHash() error {
	return nil
}

func (o *Ownership) Sign() error {
	return nil
}

func (o *Ownership) Verify() error {
	return nil
}

// accessors
func (o *Ownership) GetHash() types.Hash {
	return o.Hash
}

func (o *Ownership) GetProtocolVersion() OwnershipProtocolVersion {
	return o.ProtocolVersion
}

func (o *Ownership) GetSignatures() []types.Signature {
	return o.Signatures
}

func (o *Ownership) GetTimestamp() types.Timestamp {
	return o.Timestamp
}

func (o *Ownership) GetTicketHash() types.Hash {
	return o.TicketHash
}

func (o *Ownership) GetOwners() []crypto.Addr {
	return o.Owners
}
