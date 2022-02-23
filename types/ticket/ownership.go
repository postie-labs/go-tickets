package ticket

import (
	"github.com/postie-labs/go-tickets/types"
)

type OwnershipProtocolVersion uint16

type Ownership struct {
	// metadata
	Hash            types.Hash
	ProtocolVersion OwnershipProtocolVersion
	Signature       types.Signature // signature of owners

	// data
	Timestamp  types.Timestamp
	TicketHash types.Hash
	Owners     []types.Address
}
