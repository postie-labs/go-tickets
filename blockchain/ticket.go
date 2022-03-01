package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/postie-labs/go-tickets/types"
	"github.com/postie-labs/go-tickets/types/ticket"
	"github.com/postie-labs/go-tickets/util"
)

type TicketProtocolVersion uint16

type Ticket struct {
	Hash            types.Hash            `json:"hash"`
	ProtocolVersion TicketProtocolVersion `json:"protocol_version"`
	Signature       Signature             `json:"siguature"`

	Body        ticket.Ticket `json:"body"`
	encodedBody []byte        `json:"-"`
}

func NewTicket(protocolVersion TicketProtocolVersion, issuer crypto.Addr, ticketType ticket.TicketType, data types.Data) (*Ticket, error) {
	body := ticket.Ticket{
		Timestamp: types.TimestampNow(),
		Issuer:    issuer,
		Type:      ticketType,
		Data:      data,
	}
	encodedBody, err := body.Encode()
	if err != nil {
		return nil, err
	}
	return &Ticket{
		Hash:            util.ToSHA256(encodedBody),
		ProtocolVersion: protocolVersion,
		Signature:       Signature{},

		Body:        body,
		encodedBody: encodedBody,
	}, err
}

func NewTicketFromBytes(data []byte) (*Ticket, error) {
	ticket := Ticket{}
	err := ticket.Decode(data)
	if err != nil {
		return nil, err
	}
	encodedBody, err := ticket.Body.Encode()
	if err != nil {
		return nil, err
	}
	ticket.encodedBody = encodedBody
	return &ticket, nil
}

// ops
func (t *Ticket) Encode() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Ticket) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t *Ticket) Sign(privKey *crypto.PrivKey) error {
	sigBytes, err := privKey.Sign(t.encodedBody)
	if err != nil {
		return err
	}
	t.Signature = Signature{
		PubKey: privKey.PubKey(),
		Bytes:  sigBytes,
	}
	return nil
}

func (t *Ticket) Verify() (bool, error) {
	signature := t.Signature
	if signature.PubKey == nil {
		return false, fmt.Errorf("Ticket.Signature.PubKey is nil")
	}
	if signature.Bytes == nil {
		return false, fmt.Errorf("Ticket.Signature.Bytes is nil")
	}
	return signature.PubKey.Verify(t.encodedBody, signature.Bytes), nil
}

// accessors
func (t *Ticket) GetHash() types.Hash {
	return t.Hash
}

func (t *Ticket) GetProtocolVersion() TicketProtocolVersion {
	return t.ProtocolVersion
}

func (t *Ticket) GetSignature() Signature {
	return t.Signature
}

func (t *Ticket) GetTimestamp() types.Timestamp {
	return t.Body.Timestamp
}

func (t *Ticket) GetIssuer() crypto.Addr {
	return t.Body.Issuer
}

func (t *Ticket) GetType() ticket.TicketType {
	return t.Body.Type
}

func (t *Ticket) GetData() types.Data {
	return t.Body.Data
}
