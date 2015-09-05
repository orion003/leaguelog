package controller

import (
    "fmt"
    "net/http/httptest"
    "testing"
    
    "recleague/server"
)

var *httptest.Server

func init() {
    c := NewController()
    
    server = httptest.NewServer(server.NewRouter(c))   
}

func TestAddEmail(t *testing.T) {
    url = fmt.Sprintf("%s/users", server.URL)
}