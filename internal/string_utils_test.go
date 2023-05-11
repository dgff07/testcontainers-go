package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUppercase(t *testing.T) {
	result := uppercase("aa")

	assert.Equal(t, "AA", result)
}
