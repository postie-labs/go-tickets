package ticket

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

	data := []byte("hello world")
	ticketA := NewTicket(issuerAddr, TicketTypeSingleOwner, data)

	ticketABytes, err := ticketA.Encode()
	assert.Nil(t, err)

	ticketB, err := NewTicketFromBytes(ticketABytes)
	assert.Nil(t, err)

	assert.EqualValues(t, ticketA, ticketB)

	ticketBBytes, err := ticketB.Encode()
	assert.Nil(t, err)

	assert.EqualValues(t, ticketABytes, ticketBBytes)
}
