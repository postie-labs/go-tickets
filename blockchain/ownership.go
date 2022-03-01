package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/postie-labs/go-tickets/types"
	"github.com/postie-labs/go-tickets/util"
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
		Hash:            util.ToSHA256(encodedBody),
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

func (o *Ownership) Sign(privKey *crypto.PrivKey) error {
	sigBytes, err := privKey.Sign(o.encodedBody)
	if err != nil {
		return err
	}
	o.Signature = Signature{
		PubKey: privKey.PubKey(),
		Bytes:  sigBytes,
	}
	return nil
}

func (o *Ownership) Verify() (bool, error) {
	signature := o.Signature
	if signature.PubKey == nil {
		return false, fmt.Errorf("Ticket.Signature.PubKey is nil")
	}
	if signature.Bytes == nil {
		return false, fmt.Errorf("Ticket.Signature.Bytes is nil")
	}
	return signature.PubKey.Verify(o.encodedBody, signature.Bytes), nil
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
