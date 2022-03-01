package app

import (
	"context"

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
func (app *Application) Issue(privKey *crypto.PrivKey, data []byte) (types.Hash, error) {
	// generate and sign a new ticket
	tck, err := ticket.NewTicket(
		app.ticketProtocolVersion,
		privKey.PubKey().Address(),
		ticket.TicketTypeSingleOwner,
		data,
	)
	if err != nil {
		return types.EmptyHash, err
	}
	err = tck.Sign(privKey)
	if err != nil {
		return types.EmptyHash, err
	}

	// generate and sign a new ownership
	ows, err := ticket.NewOwnership(
		app.ownershipProtocolVersion,
		tck.Hash,
		privKey.PubKey().Address(),
	)
	if err != nil {
		return types.EmptyHash, err
	}
	err = ows.Sign(privKey)
	if err != nil {
		return types.EmptyHash, err
	}

	// store
	app.store.SetTicket(tck.Hash, tck)
	app.store.SetOwnership(ows.Hash, ows)

	return ows.Hash, nil
}

func (app *Application) Transfer() error {
	return nil
}

func (app *Application) Verify() error {
	return nil
}
