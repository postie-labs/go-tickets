package app

import "context"

type Application struct {
	ctx context.Context
}

func NewApplication() (*Application, error) {
	return &Application{
		ctx: context.Background(),
	}, nil
}

// ops
func (app *Application) Issue() error {
	return nil
}

func (app *Application) Transfer() error {
	return nil
}

func (app *Application) Verify() error {
	return nil
}

// accessors
