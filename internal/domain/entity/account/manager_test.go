package account

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureAccount()
	created, err := m.Create(u)
	assert.Nil(t, err)
	assert.Equal(t, created.ID, u.ID)
	assert.Equal(t, created.DocumentNumber, u.DocumentNumber)
}

func TestGet(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureAccount()
	_, err := m.Create(u)
	assert.Nil(t, err)
	saved, err := m.Get(u.ID)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, saved.ID)
	assert.Equal(t, u.DocumentNumber, saved.DocumentNumber)
	assert.Equal(t, u.AvailableCreditLimit, saved.AvailableCreditLimit)
}

func TestUpdate(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureAccount()
	_, err := m.Create(u)
	assert.Nil(t, err)
	saved, err := m.Get(1)
	assert.Nil(t, err)
	saved.DocumentNumber = "00987654321"
	assert.Nil(t, m.Update(saved))
	updated, err := m.Get(saved.ID)
	assert.Nil(t, err)
	assert.Equal(t, "00987654321", updated.DocumentNumber)
}

func TestDelete(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureAccount()
	_, err := m.Create(u)
	assert.Nil(t, err)

	err = m.Delete(u.ID)
	assert.Nil(t, err)

	err = m.Delete(u.ID)
	assert.Equal(t, errors.New("Not found"), err)
	_, err = m.Get(u.ID)
	assert.Equal(t, errors.New("Not found"), err)
}
