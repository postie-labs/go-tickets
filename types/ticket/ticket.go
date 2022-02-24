package ticket

import (
	"encoding/json"
	"fmt"

	crypto "github.com/postie-labs/go-crypto-lib"
	"github.com/postie-labs/go-tickets/types"
	"github.com/postie-labs/go-tickets/util"
)

type TicketType uint8

const (
	TicketTypeSingleOwner TicketType = iota
	TicketTypeMultiOwner
)

type TicketProtocolVersion uint16

type Ticket struct {
	Hash            types.Hash            `json:"hash"`
	ProtocolVersion TicketProtocolVersion `json:"protocol_version"`
	Signature       Signature             `json:"siguature"`
	Body            TicketBody            `json:"body"`
	EncodedBody     []byte                `json:"-"`
}

type TicketBody struct {
	Timestamp types.Timestamp `json:"timestamp"`
	Issuer    crypto.Addr     `json:"issuer"`
	Type      TicketType      `json:"type"`
	Data      types.Data      `json:"data"`
}

func (tb *TicketBody) encode() ([]byte, error) {
	return json.Marshal(tb)
}

func (tb *TicketBody) decode(data []byte) error {
	return json.Unmarshal(data, tb)
}

func NewTicket(protocolVersion TicketProtocolVersion, issuer crypto.Addr, ticketType TicketType, data types.Data) (*Ticket, error) {
	body := TicketBody{
		Timestamp: types.TimestampNow(),
		Issuer:    issuer,
		Type:      ticketType,
		Data:      data,
	}
	encodedBody, err := body.encode()
	if err != nil {
		return nil, err
	}
	return &Ticket{
		Hash:            util.ToSHA256(encodedBody),
		ProtocolVersion: protocolVersion,
		Signature:       Signature{},
		Body:            body,
		EncodedBody:     encodedBody,
	}, err
}

// ops
func (t *Ticket) Sign(privKey *crypto.PrivKey) error {
	sigBytes, err := privKey.Sign(t.EncodedBody)
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
	return signature.PubKey.Verify(t.Body.Data, signature.Bytes), nil
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

func (t *Ticket) GetType() TicketType {
	return t.Body.Type
}

func (t *Ticket) GetData() types.Data {
	return t.Body.Data
}
