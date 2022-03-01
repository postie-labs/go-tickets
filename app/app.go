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
}

func NewApplication() (*Application, error) {
	store, err := NewStore()
	if err != nil {
		return nil, err
	}
	return &Application{
		ctx:   context.Background(),
		store: store,
	}, nil
}

// ops

// subject: issuer
func (app *Application) Issue(issuer *crypto.PrivKey, data []byte) (types.Hash, error) {
	// generate and sign a new ticket
	tck := ticket.NewTicket(
		issuer.PubKey().Address(),
		ticket.TicketTypeSingleOwner,
		data,
	)

	tckHash, err := tck.Hash()
	if err != nil {
		return types.EmptyHash, err
	}

	// store
	app.store.SetTicket(tckHash, tck)
	app.store.SetOwnership(tckHash, issuer.PubKey().Address())

	return tckHash, nil
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

// subject: validator
func (app *Application) Verify(owner crypto.Addr, tckHash types.Hash) (bool, error) {
	realOwner := app.store.GetOwnership(tckHash)
	if !realOwner.Equals(owner) {
		return false, fmt.Errorf("failed to match address: %s != %s", realOwner, owner)
	}
	return true, nil
}
