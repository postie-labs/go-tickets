package app

import (
	"context"

	"github.com/postie-labs/go-postie-lib/crypto"
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
func (app *Application) Issue(privKey crypto.PrivKey) error {

	return nil
}

func (app *Application) Transfer() error {
	return nil
}

func (app *Application) Verify() error {
	return nil
}
