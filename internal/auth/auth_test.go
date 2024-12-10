package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	gotHeader := http.Header{}
	gotHeader.Add("Authorization", "ApiKey absc234asdadasda")
	got, err := GetAPIKey(gotHeader)
	if err != nil {
		t.Fatalf("got: %v", got)
	}
}
