package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	err := m.Create("active-tab", "getir")
	assert.Nil(t, err)
}