package utils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidStatus(t *testing.T) {
	assert.True(t, IsValidStatus(http.StatusOK), "Status valid. Incorrect response")
	assert.True(t, IsValidStatus(http.StatusTooManyRequests), "Status valid. Incorrect response")
	assert.False(t, IsValidStatus(http.StatusNotFound), "Status valid. Incorrect response")
	assert.False(t, IsValidStatus(http.StatusInternalServerError), "Status valid. Incorrect response")
}
