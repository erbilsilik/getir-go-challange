package entity_test

import (
	"testing"

	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewRecord(t *testing.T) {
	counts := []int{150, 160}
	r, err := entity.NewRecord("Getir test", counts)
	assert.Nil(t, err)
	assert.Equal(t, r.Value, "Getir test")
	assert.NotNil(t, r.Key)
	assert.NotNil(t, r.CreatedAt)
	assert.Equal(t,0, r.TotalCount)
	assert.Equal(t, r.Counts, counts)
}
