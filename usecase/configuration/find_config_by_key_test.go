package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FindByKey(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	m.Create("active-tab", "getir")
	c := m.FindByKey("active-tab")
	assert.Equal(t, c.Value, "getir")
}