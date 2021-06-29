package utils

import (
	"net/http"
	"strconv"
	"strings"
)

func IsValidStatus(status int) bool {

	if status == http.StatusTooManyRequests {
		return true
	}

	statusStr := strconv.Itoa(status)
	statusSplited := strings.Split(statusStr, "")

	return statusSplited[0] == "2"

}
