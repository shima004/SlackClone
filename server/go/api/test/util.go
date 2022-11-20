package test

import (
	"Slack/server"
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
)

func GetToken(email string, password string) string {
	router := server.NewServer()
	var j = map[string]string{
		"email":    email,
		"password": password,
	}
	b, _ := json.Marshal(j)

	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	var res map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	if err != nil {
		log.Println(err)
	}

	return res["token"]
}
