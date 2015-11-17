package service

import (
	"io/ioutil"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	claims := make(map[string]interface{})
	claims["id"] = "test@leaguelog.ca"

	key, err := ioutil.ReadFile("testdata/hmac_key")
	if err != nil {
		t.Fatal("Unable to read HMAC key.")
	}

	j := InitializeJwt(key)
	token := j.Generate(claims)

	expectedToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RAbGVhZ3VlbG9nLmNhIn0.mkGPonVdQeQ1nTezFMVKGHvZiAY9L1dLrLDmhcV-4gvZ4bGOm8J1jSlh5eRd-eSWrOkiZqnIHRU4i1gELq2S2A"
	if token != expectedToken {
		t.Error("Incorrect token - expected: %s, received: %s", expectedToken, token)
	}
}

func TestValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RAbGVhZ3VlbG9nLmNhIn0.mkGPonVdQeQ1nTezFMVKGHvZiAY9L1dLrLDmhcV-4gvZ4bGOm8J1jSlh5eRd-eSWrOkiZqnIHRU4i1gELq2S2A"

	key, err := ioutil.ReadFile("testdata/hmac_key")
	if err != nil {
		t.Fatal("Unable to read HMAC key.")
	}

	j := InitializeJwt(key)
	err = j.Validate(token)
	if err != nil {
		t.Errorf("Token should be valid (%v): %s", err, token)
	}
}

func TestInvalidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImludmFsaWRAbGVhZ3VlbG9nLmNhIn0.mkGPonVdQeQ1nTezFMVKGHvZiAY9L1dLrLDmhcV-4gvZ4bGOm8J1jSlh5eRd-eSWrOkiZqnIHRU4i1gELq2S2A"

	key, err := ioutil.ReadFile("testdata/hmac_key")
	if err != nil {
		t.Fatal("Unable to read HMAC key.")
	}

	j := InitializeJwt(key)
	err = j.Validate(token)
	if err == nil {
		t.Errorf("Token should be invalid: %s", token)
	}
}
