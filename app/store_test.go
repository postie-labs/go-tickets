package app

import (
	"testing"

	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/postie-labs/go-tickets/types/ticket"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	// common
	issuerPrivKey, err := crypto.GenPrivKey()
	assert.NoError(t, err)
	ownerPrivKey, err := crypto.GenPrivKey()
	assert.NoError(t, err)
	sto, err := NewStore()
	assert.NoError(t, err)

	// generate dummy tickets
	tckA, err := ticket.NewTicket(
		DefaultTicketProtocolVersion,
		issuerPrivKey.PubKey().Address(),
		ticket.TicketTypeSingleOwner,
		[]byte("hello world 0"),
	)
	assert.NoError(t, err)
	tckB, err := ticket.NewTicket(
		DefaultTicketProtocolVersion,
		issuerPrivKey.PubKey().Address(),
		ticket.TicketTypeSingleOwner,
		[]byte("hello world 1"),
	)
	assert.NoError(t, err)

	// store and check dummy tickets
	sto.SetTicket(tckA.Hash, tckA)
	tckAToVerify := sto.GetTicket(tckA.Hash)
	assert.EqualValues(t, tckA, tckAToVerify)
	sto.SetTicket(tckB.Hash, tckB)
	tckBToVerify := sto.GetTicket(tckB.Hash)
	assert.EqualValues(t, tckB, tckBToVerify)

	// generate dummy ownerships
	owsA, err := ticket.NewOwnership(
		DefaultOwnershipProtocolVersion,
		tckA.Hash,
		ownerPrivKey.PubKey().Address(),
	)
	assert.NoError(t, err)
	owsB, err := ticket.NewOwnership(
		DefaultOwnershipProtocolVersion,
		tckB.Hash,
		ownerPrivKey.PubKey().Address(),
	)
	assert.NoError(t, err)

	// store and check dummy ownerships
	sto.SetOwnership(owsA.Hash, owsA)
	owsAToVerify := sto.GetOwnership(owsA.Hash)
	assert.EqualValues(t, owsA, owsAToVerify)
	sto.SetOwnership(owsB.Hash, owsB)
	owsBToVerify := sto.GetOwnership(owsB.Hash)
	assert.EqualValues(t, owsB, owsBToVerify)
}
