package ticket

import (
	"encoding/hex"
	"testing"

	"github.com/postie-labs/go-crypto-lib"
	"github.com/postie-labs/go-tickets/types"
	"github.com/stretchr/testify/assert"
)

func TestOwnership(t *testing.T) {
	ownerPrivKey, err := crypto.GenPrivKey()
	assert.Nil(t, err)
	ownerPubKey := ownerPrivKey.PubKey()
	ownerAddr := ownerPubKey.Address()

	ticketHash := types.Hash{}
	tmp, err := hex.DecodeString("56d875c7a087ddae10dc92051dd7195cf0b03ff1e05370f2347ff6509c8cc884")
	assert.Nil(t, err)
	copy(ticketHash[:], tmp)

	ownershipA, err := NewOwnership(1, ticketHash, ownerAddr)
	assert.Nil(t, err)

	err = ownershipA.Sign(ownerPrivKey)
	assert.Nil(t, err)

	verified, err := ownershipA.Verify()
	assert.Nil(t, err)
	assert.True(t, verified)

	ownershipABytes, err := ownershipA.Encode()
	assert.Nil(t, err)

	ownershipB, err := NewOwnershipFromBytes(ownershipABytes)
	assert.Nil(t, err)

	assert.EqualValues(t, ownershipA, ownershipB)

	ownershipBBytes, err := ownershipB.Encode()
	assert.Nil(t, err)

	assert.EqualValues(t, ownershipABytes, ownershipBBytes)

	verified, err = ownershipB.Verify()
	assert.Nil(t, err)
	assert.True(t, verified)
}
