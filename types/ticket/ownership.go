package ticket

import (
	"encoding/json"

	crypto "github.com/postie-labs/go-crypto-lib"
	"github.com/postie-labs/go-tickets/types"
)

type OwnershipProtocolVersion uint16

type Ownership struct {
	Hash            types.Hash               `json:"hash"`
	ProtocolVersion OwnershipProtocolVersion `json:"protocol_version"`
	Signature       Signature                `json:"signature"`

	Body        OwnershipBody `json:"body"`
	encodedBody []byte        `json:"-"`
}

type OwnershipBody struct {
	Timestamp  types.Timestamp `json:"timestamp"`
	TicketHash types.Hash      `json:"ticket_hash"`
	Owner      crypto.Addr     `json:"owner"`
}

func (ob *OwnershipBody) encode() ([]byte, error) {
	return json.Marshal(ob)
}

func NewOwnership(protocolVersion OwnershipProtocolVersion, ticketHash types.Hash, owner crypto.Addr) (*Ownership, error) {
	body := OwnershipBody{
		Timestamp:  types.TimestampNow(),
		TicketHash: ticketHash,
		Owner:      owner,
	}
	encodedBody, err := body.encode()
	if err != nil {
		return nil, err
	}
	return &Ownership{
		Hash:            types.Hash{},
		ProtocolVersion: protocolVersion,
		Signature:       Signature{},

		Body:        body,
		encodedBody: encodedBody,
	}, nil
}

func NewOwnershipFromBytes(data []byte) (*Ownership, error) {
	ownership := Ownership{}
	err := ownership.Decode(data)
	if err != nil {
		return nil, err
	}
	encodedBody, err := ownership.Body.encode()
	if err != nil {
		return nil, err
	}
	ownership.encodedBody = encodedBody
	return &ownership, nil
}

// ops
func (o *Ownership) Encode() ([]byte, error) {
	return json.Marshal(o)
}

func (o *Ownership) Decode(data []byte) error {
	return json.Unmarshal(data, o)
}

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

func (o *Ownership) GetSignature() Signature {
	return o.Signature
}

func (o *Ownership) GetTimestamp() types.Timestamp {
	return o.Body.Timestamp
}

func (o *Ownership) GetTicketHash() types.Hash {
	return o.Body.TicketHash
}

func (o *Ownership) GetOwner() crypto.Addr {
	return o.Body.Owner
}
