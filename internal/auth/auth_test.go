package auth

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAPIKeyValid(t *testing.T) {
	w := httptest.NewRecorder()
	w.Header().Set("Authorization", "ApiKey cwnjleks1i")

	got, _ := GetAPIKey(w.Header())
	want := "cwnjleks1i"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKeyNoAuthHeader(t *testing.T) {
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	errorMsg := "no authorization header included"

	_, err := GetAPIKey(w.Header())
	if errorMsg != err.Error() {
		t.Fatalf("expected error: %v, got: %v", errorMsg, err.Error())
	}
}
