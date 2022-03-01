package ticket

import (
	"encoding/json"

	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/postie-labs/go-tickets/types"
	"github.com/postie-labs/go-tickets/util"
)

type TicketType uint8

const (
	TicketTypeSingleOwner TicketType = iota
	TicketTypeMultiOwner
)

type Ticket struct {
	Timestamp types.Timestamp `json:"timestamp"`
	Issuer    crypto.Addr     `json:"issuer"`
	Type      TicketType      `json:"type"`
	Data      types.Data      `json:"data"`
}

func NewTicket(issuer crypto.Addr, tckType TicketType, data types.Data) *Ticket {
	return &Ticket{
		Timestamp: types.TimestampNow(),
		Issuer:    issuer,
		Type:      tckType,
		Data:      data,
	}
}

func NewTicketFromBytes(data []byte) (*Ticket, error) {
	ticket := Ticket{}
	err := ticket.Decode(data)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

// ops
func (t *Ticket) Encode() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Ticket) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t *Ticket) Hash() (types.Hash, error) {
	data, err := t.Encode()
	if err != nil {
		return types.EmptyHash, err
	}
	return util.ToSHA256(data), nil
}

// accessors
func (t *Ticket) GetTimestamp() types.Timestamp {
	return t.Timestamp
}

func (t *Ticket) GetIssuer() crypto.Addr {
	return t.Issuer
}

func (t *Ticket) GetType() TicketType {
	return t.Type
}

func (t *Ticket) GetData() types.Data {
	return t.Data
}
