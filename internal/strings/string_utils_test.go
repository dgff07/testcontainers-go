package strings

import (
	"testing"

	"github.com/dgff07/testcontainers-go/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	tests.InitContainers(m)
}

func TestUppercase(t *testing.T) {
	result := uppercase("aa")

	assert.Equal(t, "AA", result)
}
