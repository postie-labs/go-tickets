package app

import (
	"context"
	"fmt"

	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/postie-labs/go-tickets/types"
	"github.com/postie-labs/go-tickets/types/ticket"
)

const (
	DefaultTicketProtocolVersion    = 1
	DefaultOwnershipProtocolVersion = 1
)

type Application struct {
	ctx   context.Context
	store *Store

	ticketProtocolVersion    ticket.TicketProtocolVersion
	ownershipProtocolVersion ticket.OwnershipProtocolVersion
}

func NewApplication() (*Application, error) {
	store, err := NewStore()
	if err != nil {
		return nil, err
	}
	return &Application{
		ctx:   context.Background(),
		store: store,

		ticketProtocolVersion:    DefaultTicketProtocolVersion,
		ownershipProtocolVersion: DefaultOwnershipProtocolVersion,
	}, nil
}

// ops

// subject: issuer
func (app *Application) Issue(issuer *crypto.PrivKey, data []byte) (types.Hash, error) {
	// generate and sign a new ticket
	tck, err := ticket.NewTicket(
		app.ticketProtocolVersion,
		issuer.PubKey().Address(),
		ticket.TicketTypeSingleOwner,
		data,
	)
	if err != nil {
		return types.EmptyHash, err
	}
	err = tck.Sign(issuer)
	if err != nil {
		return types.EmptyHash, err
	}

	// store
	app.store.SetTicket(tck.Hash, tck)
	app.store.SetOwnership(tck.Hash, issuer.PubKey().Address())

	return tck.Hash, nil
}

// subject: owner
func (app *Application) Transfer(from, to crypto.Addr, tckHash types.Hash) error {
	// check permission
	owner := app.store.GetOwnership(tckHash)
	if owner.Equals("") {
		return fmt.Errorf("faild to find ownership: %s", tckHash)
	}
	if !from.Equals(owner) {
		return fmt.Errorf("failed to match address: %s != %s", from, owner)
	}

	// store
	app.store.SetOwnership(tckHash, to)

	return nil
}

func (app *Application) Verify() error {
	return nil
}
