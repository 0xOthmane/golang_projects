package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info found")
	}
	// Authorization: ApiKey {...hex}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part og authorization header")
	}

	return vals[1], nil
}  