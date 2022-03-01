package blockchain

import (
	"testing"

	"github.com/postie-labs/go-postie-lib/crypto"
	"github.com/stretchr/testify/assert"
)

func TestTicket(t *testing.T) {
	issuerPrivKey, err := crypto.GenPrivKey()
	assert.Nil(t, err)
	issuerPubKey := issuerPrivKey.PubKey()
	issuerAddr := issuerPubKey.Address()

	ticketA, err := NewTicket(1, issuerAddr, TicketTypeSingleOwner, []byte("hello world"))
	assert.Nil(t, err)

	err = ticketA.Sign(issuerPrivKey)
	assert.Nil(t, err)

	verified, err := ticketA.Verify()
	assert.Nil(t, err)
	assert.True(t, verified)

	ticketABytes, err := ticketA.Encode()
	assert.Nil(t, err)

	ticketB, err := NewTicketFromBytes(ticketABytes)
	assert.Nil(t, err)

	assert.EqualValues(t, ticketA, ticketB)

	ticketBBytes, err := ticketB.Encode()
	assert.Nil(t, err)

	assert.EqualValues(t, ticketABytes, ticketBBytes)

	verified, err = ticketB.Verify()
	assert.Nil(t, err)
	assert.True(t, verified)
}
