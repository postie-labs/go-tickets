package app

import (
	"testing"

	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/postie-labs/go-tickets/types/ticket"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	// common
	alice, err := crypto.GenPrivKey()
	assert.NoError(t, err)
	aliceAddr := alice.PubKey().Address()
	bob, err := crypto.GenPrivKey()
	assert.NoError(t, err)
	bobAddr := bob.PubKey().Address()
	sto, err := NewStore()
	assert.NoError(t, err)

	// generate dummy tickets
	tckA, err := ticket.NewTicket(
		DefaultTicketProtocolVersion,
		aliceAddr,
		ticket.TicketTypeSingleOwner,
		[]byte("hello world 0"),
	)
	assert.NoError(t, err)
	tckB, err := ticket.NewTicket(
		DefaultTicketProtocolVersion,
		bobAddr,
		ticket.TicketTypeSingleOwner,
		[]byte("hello world 1"),
	)
	assert.NoError(t, err)

	// store and check dummy tickets
	sto.SetTicket(tckA.Hash, tckA)
	assert.Equal(t, 1, len(sto.tickets))
	tckAToVerify := sto.GetTicket(tckA.Hash)
	assert.EqualValues(t, tckA, tckAToVerify)
	sto.SetTicket(tckB.Hash, tckB)
	assert.Equal(t, 2, len(sto.tickets))
	tckBToVerify := sto.GetTicket(tckB.Hash)
	assert.EqualValues(t, tckB, tckBToVerify)

	// store and check dummy ownerships
	sto.SetOwnership(tckA.Hash, aliceAddr)
	assert.Equal(t, 1, len(sto.ownerships))
	tmpAddr := sto.GetOwnership(tckA.Hash)
	assert.True(t, tmpAddr.Equals(aliceAddr))

	sto.SetOwnership(tckB.Hash, aliceAddr)
	assert.Equal(t, 2, len(sto.ownerships))
	tmpAddr = sto.GetOwnership(tckB.Hash)
	assert.True(t, tmpAddr.Equals(aliceAddr))

	sto.SetOwnership(tckB.Hash, bobAddr)
	assert.Equal(t, 2, len(sto.ownerships))
	tmpAddr = sto.GetOwnership(tckB.Hash)
	assert.True(t, tmpAddr.Equals(bobAddr))

	// remove and check dummy ownerships
	sto.RemoveOwnership(tckA.Hash)
	assert.Equal(t, 1, len(sto.ownerships))
	tmpAddr = sto.GetOwnership(tckA.Hash)
	assert.EqualValues(t, tmpAddr, "")
	sto.RemoveOwnership(tckB.Hash)
	assert.Equal(t, 0, len(sto.ownerships))
	tmpAddr = sto.GetOwnership(tckB.Hash)
	assert.EqualValues(t, tmpAddr, "")

	// remove and check dummy tickets
	sto.RemoveTicket(tckA.Hash)
	assert.Equal(t, 1, len(sto.tickets))
	tckAToVerify = sto.GetTicket(tckA.Hash)
	assert.Nil(t, tckAToVerify)
	sto.RemoveTicket(tckB.Hash)
	assert.Equal(t, 0, len(sto.tickets))
	tckBToVerify = sto.GetTicket(tckB.Hash)
	assert.Nil(t, tckBToVerify)

}
