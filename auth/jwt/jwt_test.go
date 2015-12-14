package jwt

import (
	"testing"
)

var hmac = []byte("579760E50509F2F28324421C7509741F5BF03B9158161076B3C6B39FB028D9E2C251490A3F8BD1F59728259A673668CAFEB49C9E9499F8386B147D7260B6937A")

func TestGenerateToken(t *testing.T) {
	claims := make(map[string]interface{})
	claims["id"] = "test@leaguelog.ca"

	j := InitializeJwt(hmac)
	token := j.Generate(claims)

	expectedToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RAbGVhZ3VlbG9nLmNhIn0.Xx8WqprVmafi1ZaoMWiL43dC5H-Im0hTtJydprrFGWWndgeTQB91aWSan37NrIBP_rBL06Axv-MO0WJNji70kw"
	if token != expectedToken {
		t.Errorf("Incorrect token - expected: %s, received: %s", expectedToken, token)
	}
}

func TestValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RAbGVhZ3VlbG9nLmNhIn0.Xx8WqprVmafi1ZaoMWiL43dC5H-Im0hTtJydprrFGWWndgeTQB91aWSan37NrIBP_rBL06Axv-MO0WJNji70kw"

	j := InitializeJwt(hmac)
	err := j.Validate(token)
	if err != nil {
		t.Errorf("Token should be valid (%v): %s", err, token)
	}
}

func TestInvalidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImludmFsaWRAbGVhZ3VlbG9nLmNhIn0.mkGPonVdQeQ1nTezFMVKGHvZiAY9L1dLrLDmhcV-4gvZ4bGOm8J1jSlh5eRd-eSWrOkiZqnIHRU4i1gELq2S2A"

	j := InitializeJwt(hmac)
	err := j.Validate(token)
	if err == nil {
		t.Errorf("Token should be invalid: %s", token)
	}
}
