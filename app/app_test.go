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
	tckHash, err := app.Issue(aliceAddr, data)
	assert.NoError(t, err)

	owner := app.store.GetOwnership(tckHash)
	assert.True(t, owner.Equals(aliceAddr))

	tck := app.store.GetTicket(tckHash)
	issuer := tck.GetIssuer()
	assert.True(t, aliceAddr.Equals(issuer))

	tckData := tck.GetData()
	assert.EqualValues(t, data, tckData)

	// TRANSFER
	err = app.Transfer(aliceAddr, bobAddr, tckHash)
	assert.NoError(t, err)

	owner = app.store.GetOwnership(tckHash)
	assert.True(t, bobAddr.Equals(owner))

	// VERIFY
	verified, err := app.Verify(aliceAddr, tckHash)
	assert.NoError(t, err)
	assert.False(t, verified)

	verified, err = app.Verify(bobAddr, tckHash)
	assert.NoError(t, err)
	assert.True(t, verified)
}
