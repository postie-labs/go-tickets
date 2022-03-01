package app

import (
	"testing"

	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/stretchr/testify/assert"
)

func TestApplication(t *testing.T) {
	// common
	privKey, err := crypto.GenPrivKey()
	assert.NoError(t, err)
	app, err := NewApplication()
	assert.NoError(t, err)

	data := []byte("hello world 0")
	owsHash, err := app.Issue(privKey, data)
	assert.NoError(t, err)

	ows := app.store.GetOwnership(owsHash)
	owsOwner := ows.GetOwner()
	assert.EqualValues(t, privKey.PubKey().Address(), owsOwner)

	owsVerified, err := ows.Verify()
	assert.NoError(t, err)
	assert.True(t, owsVerified)

	tckHash := ows.GetTicketHash()
	tck := app.store.GetTicket(tckHash)

	tckIssuer := tck.GetIssuer()
	assert.EqualValues(t, privKey.PubKey().Address(), tckIssuer)

	tckVerified, err := tck.Verify()
	assert.NoError(t, err)
	assert.True(t, tckVerified)

	tckData := tck.GetData()
	assert.EqualValues(t, data, tckData)
}
