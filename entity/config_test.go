package entity_test

import (
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := entity.Config{Key: "active-tabs", Value: "getir"}
	assert.Equal(t, c.Key, "active-tabs")
	assert.Equal(t, c.Value, "getir")
}

func TestConfigValidate(t *testing.T) {
	type test struct {
		key    string
		value   interface{}
		want     error
	}

	tests := []test{
		{
			key:    "disable-notifications",
			value:   true,
			want:     nil,
		},
		{
			key:      "enable-cache",
			want:     entity.ErrInvalidEntity,
		},
		{
			value:    true,
			want:     entity.ErrInvalidEntity,
		},
		{
			key:    "enable-cache",
			value:   true,
			want:     nil,
		},
		{
			key:    "default-text",
			value:   "Default Text",
			want:     nil,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewConfig(tc.key, tc.value)
		assert.Equal(t, err, tc.want)
	}
}