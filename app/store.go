package app

import (
	"github.com/postie-labs/go-tickets/types"
	"github.com/postie-labs/go-tickets/types/ticket"
)

type Store struct {
	tickets    map[types.Hash]*ticket.Ticket
	ownerships map[types.Hash]*ticket.Ownership
}

func NewStore() (*Store, error) {
	return &Store{
		tickets:    make(map[types.Hash]*ticket.Ticket),
		ownerships: make(map[types.Hash]*ticket.Ownership),
	}, nil
}

// accessors
func (s *Store) GetTicket(hash types.Hash) *ticket.Ticket {
	return s.tickets[hash]
}

func (s *Store) SetTicket(hash types.Hash, tck *ticket.Ticket) {
	s.tickets[hash] = tck
}

func (s *Store) GetOwnership(hash types.Hash) *ticket.Ownership {
	return s.ownerships[hash]
}

func (s *Store) SetOwnership(hash types.Hash, ows *ticket.Ownership) {
	s.ownerships[hash] = ows
}
