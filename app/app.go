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

	// generate and sign a new ownership
	ows, err := ticket.NewOwnership(
		app.ownershipProtocolVersion,
		tck.Hash,
		issuer.PubKey().Address(),
	)
	if err != nil {
		return types.EmptyHash, err
	}
	err = ows.Sign(issuer)
	if err != nil {
		return types.EmptyHash, err
	}

	// store
	app.store.SetTicket(tck.Hash, tck)
	app.store.SetOwnership(ows.Hash, ows)

	return ows.Hash, nil
}

// subject: owner
func (app *Application) Transfer(fromOwner, toOwner *crypto.PrivKey, fromOwsHash types.Hash) (types.Hash, error) {
	fromOwnerAddr := fromOwner.PubKey().Address()
	fromOws := app.store.GetOwnership(fromOwsHash)
	if fromOws == nil {
		return types.EmptyHash, fmt.Errorf("failed to find ownership: %s", fromOwsHash)
	}

	// verify fromOws
	ok, err := fromOws.Verify()
	if err != nil {
		return types.EmptyHash, err
	}
	if !ok {
		return types.EmptyHash, fmt.Errorf("failed to verify ownership")
	}

	// check permission
	if !fromOwnerAddr.Equals(fromOws.GetOwner()) {
		return types.EmptyHash, fmt.Errorf("owner doesn't match: %s != %s", fromOwner, fromOws.GetOwner())
	}

	// generate and sign a new owership
	toOws, err := ticket.NewOwnership(
		app.ownershipProtocolVersion,
		fromOws.GetTicketHash(),
		toOwner.PubKey().Address(),
	)
	if err != nil {
		return types.EmptyHash, err
	}
	err = toOws.Sign(toOwner)
	if err != nil {
		return types.EmptyHash, err
	}

	// store
	app.store.RemoveOwnership(fromOwsHash)
	app.store.SetOwnership(toOws.Hash, toOws)

	return toOws.Hash, nil
}

func (app *Application) Verify() error {
	return nil
}
