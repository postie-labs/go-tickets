package app

import (
	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/postie-labs/go-tickets/types"
	"github.com/postie-labs/go-tickets/types/ticket"
)

type Store struct {
	balances   map[crypto.Addr]uint32
	tickets    map[types.Hash]*ticket.Ticket
	ownerships map[types.Hash]crypto.Addr
}

func NewStore() (*Store, error) {
	return &Store{
		balances:   make(map[crypto.Addr]uint32),        // addr: balance
		tickets:    make(map[types.Hash]*ticket.Ticket), // ticket_hash: ticket
		ownerships: make(map[types.Hash]crypto.Addr),    // ticket_hash: owner
	}, nil
}

// accessors
func (s *Store) GetBalance(addr crypto.Addr) uint32 {
	return s.balances[addr]
}

func (s *Store) SetBalance(addr crypto.Addr, amount uint32) {
	s.balances[addr] = amount
}

func (s *Store) GetTicket(hash types.Hash) *ticket.Ticket {
	return s.tickets[hash]
}

func (s *Store) SetTicket(hash types.Hash, tck *ticket.Ticket) {
	s.tickets[hash] = tck
}

func (s *Store) RemoveTicket(hash types.Hash) {
	delete(s.tickets, hash)
}

func (s *Store) GetOwnership(hash types.Hash) crypto.Addr {
	return s.ownerships[hash]
}

func (s *Store) SetOwnership(hash types.Hash, addr crypto.Addr) {
	s.ownerships[hash] = addr
}

func (s *Store) RemoveOwnership(hash types.Hash) {
	delete(s.ownerships, hash)
}
