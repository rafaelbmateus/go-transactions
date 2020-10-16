package operation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureOperation()
	err := m.Create(u)
	assert.Nil(t, err)
	assert.Equal(t, 1, u.ID)
	assert.Equal(t, "compra a vista", u.Description)
}

func Test_Get(t *testing.T) {
	repo := NewInmemRepository()
	m := NewManager(repo)
	u := NewFixtureOperation()
	err := m.Create(u)
	assert.Nil(t, err)
	uu, err := m.Get(u.ID)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, uu.ID)
	assert.Equal(t, u.Description, uu.Description)
}
