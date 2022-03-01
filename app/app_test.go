package app

import (
	"testing"

	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/stretchr/testify/assert"
)

func TestApplication(t *testing.T) {
	// common
	alice, err := crypto.GenPrivKey()
	assert.NoError(t, err)
	aliceAddr := alice.PubKey().Address()
	bob, err := crypto.GenPrivKey()
	assert.NoError(t, err)
	bobAddr := bob.PubKey().Address()
	app, err := NewApplication()
	assert.NoError(t, err)

	// ISSUE
	data := []byte("hello world 0")
	owsAHash, err := app.Issue(alice, data)
	assert.NoError(t, err)

	owsA := app.store.GetOwnership(owsAHash)
	assert.NotNil(t, owsA)

	owsAOwner := owsA.GetOwner()
	assert.True(t, aliceAddr.Equals(owsAOwner))

	verified, err := owsA.Verify()
	assert.NoError(t, err)
	assert.True(t, verified)

	tckAHash := owsA.GetTicketHash()
	tckA := app.store.GetTicket(tckAHash)

	tckAIssuer := tckA.GetIssuer()
	assert.True(t, aliceAddr.Equals(tckAIssuer))

	verified, err = tckA.Verify()
	assert.NoError(t, err)
	assert.True(t, verified)

	tckAData := tckA.GetData()
	assert.EqualValues(t, data, tckAData)

	// TRANSFER
	owsBHash, err := app.Transfer(alice, bob, owsAHash)
	assert.NoError(t, err)

	owsA = app.store.GetOwnership(owsAHash)
	assert.Nil(t, owsA)
	owsB := app.store.GetOwnership(owsBHash)
	assert.NotNil(t, owsB)

	owsBOwner := owsB.GetOwner()
	assert.True(t, bobAddr.Equals(owsBOwner))

	verified, err = owsB.Verify()
	assert.NoError(t, err)
	assert.True(t, verified)

	tckBHash := owsB.GetTicketHash()
	assert.True(t, tckAHash.Equals(tckBHash))
	tckB := app.store.GetTicket(tckBHash)
	assert.EqualValues(t, tckA, tckB)
}
