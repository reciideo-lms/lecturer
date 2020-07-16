package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlugString(t *testing.T) {
	t.Parallel()
	var flagtest = []struct {
		in  string
		out string
	}{
		{"Hello", "hello"},
		{"Tschüss", "tschuss"},
		{"Hello There", "hellothere"},
		{"Tschüß", "tschuss"},
	}

	for _, tt := range flagtest {
		t.Run(tt.out, func(t *testing.T) {
			t.Parallel()
			result, err := SlugString(tt.in)

			assert.NotNil(t, result)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, result)
		})
	}
}

func TestConcatStrings(t *testing.T) {
	t.Parallel()
	var flagtest = []struct {
		in1 string
		in2 string
		out string
	}{
		{"1", "2", "2"},
		{"test", "works", "testworks"},
		{"helps", "always", "helpsalways"},
	}

	for _, tt := range flagtest {
		t.Run(tt.out, func(t *testing.T) {
			t.Parallel()
			result := ConcatStrings(tt.in1, tt.in2)

			assert.NotNil(t, result)
			assert.Equal(t, tt.out, result)
		})
	}
}
