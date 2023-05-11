package strings

import (
	"strings"

	"github.com/dgff07/testcontainers-go/internal/tests"
)

func uppercase(value string) string {
	tests.DB.Ping()
	return strings.ToUpper(value)
}
