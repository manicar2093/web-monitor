package connections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSqliteConection(t *testing.T) {
	dbName := "testing.db"
	got := NewSqliteConection(dbName)

	assert.NotNil(t, got, "There is no database connection")
}
