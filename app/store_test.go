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
	tckA := ticket.NewTicket(
		aliceAddr,
		ticket.TicketTypeSingleOwner,
		[]byte("hello world 0"),
	)
	tckAHash, err := tckA.Hash()
	assert.NoError(t, err)
	tckB := ticket.NewTicket(
		bobAddr,
		ticket.TicketTypeSingleOwner,
		[]byte("hello world 1"),
	)
	tckBHash, err := tckB.Hash()
	assert.NoError(t, err)

	// store and check dummy tickets
	sto.SetTicket(tckAHash, tckA)
	assert.Equal(t, 1, len(sto.tickets))
	tckAToVerify := sto.GetTicket(tckAHash)
	assert.EqualValues(t, tckA, tckAToVerify)
	sto.SetTicket(tckBHash, tckB)
	assert.Equal(t, 2, len(sto.tickets))
	tckBToVerify := sto.GetTicket(tckBHash)
	assert.EqualValues(t, tckB, tckBToVerify)

	// store and check dummy ownerships
	sto.SetOwnership(tckAHash, aliceAddr)
	assert.Equal(t, 1, len(sto.ownerships))
	tmpAddr := sto.GetOwnership(tckAHash)
	assert.True(t, tmpAddr.Equals(aliceAddr))

	sto.SetOwnership(tckBHash, aliceAddr)
	assert.Equal(t, 2, len(sto.ownerships))
	tmpAddr = sto.GetOwnership(tckBHash)
	assert.True(t, tmpAddr.Equals(aliceAddr))

	sto.SetOwnership(tckBHash, bobAddr)
	assert.Equal(t, 2, len(sto.ownerships))
	tmpAddr = sto.GetOwnership(tckBHash)
	assert.True(t, tmpAddr.Equals(bobAddr))

	// remove and check dummy ownerships
	sto.RemoveOwnership(tckAHash)
	assert.Equal(t, 1, len(sto.ownerships))
	tmpAddr = sto.GetOwnership(tckAHash)
	assert.EqualValues(t, tmpAddr, "")
	sto.RemoveOwnership(tckBHash)
	assert.Equal(t, 0, len(sto.ownerships))
	tmpAddr = sto.GetOwnership(tckBHash)
	assert.EqualValues(t, tmpAddr, "")

	// remove and check dummy tickets
	sto.RemoveTicket(tckAHash)
	assert.Equal(t, 1, len(sto.tickets))
	tckAToVerify = sto.GetTicket(tckAHash)
	assert.Nil(t, tckAToVerify)
	sto.RemoveTicket(tckBHash)
	assert.Equal(t, 0, len(sto.tickets))
	tckBToVerify = sto.GetTicket(tckBHash)
	assert.Nil(t, tckBToVerify)

}
