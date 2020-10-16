package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureTransaction()
	err := m.Create(u)
	assert.Nil(t, err)
	assert.Equal(t, 1, u.ID)
	assert.Equal(t, 123.45, u.Amount)
}

func TestGet(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureTransaction()
	err := m.Create(u)
	assert.Nil(t, err)
	saved, err := m.Get(u.ID)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, saved.ID)
	assert.Equal(t, u.AccountID, saved.AccountID)
	assert.Equal(t, u.OperationTypeID, saved.OperationTypeID)
	assert.Equal(t, u.Amount, saved.Amount)
}
