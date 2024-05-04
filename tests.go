package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {

	os.Setenv("DB_USER", "test_user")
	os.Setenv("DB_NAME", "test_db")
	os.Setenv("DB_PASSWORD", "test_password")
	os.Setenv("DB_SSLMODE", "disable")
	os.Setenv("DB_HOST", "localhost")

	code := m.Run()

	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_SSLMODE")
	os.Unsetenv("DB_HOST")

	os.Exit(code)
}

func TestHandleIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleIndex)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `<!DOCTYPE html>
<html>
...
</html>
`
	if rr.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
