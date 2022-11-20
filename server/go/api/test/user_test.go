package test

import (
	"Slack/apifunc"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"Slack/server"

	"github.com/stretchr/testify/assert"
)

var (
	email    = "test@test.com"
	password = "test"
	name     = "test"
)

func TestUserPost(t *testing.T) {
	router := server.NewServer()
	var j = apifunc.UserPostParams{
		Email:    email,
		Password: password,
		Name:     name,
	}
	b, _ := json.Marshal(j)

	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	var res map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	if err != nil {
		log.Println(err)
	}

	assert.Equal(t, "success", res["message"])
}

func TestLoginPost(t *testing.T) {
	router := server.NewServer()
	var j = map[string]string{
		"email":    "test@test.com",
		"password": "test",
	}
	b, _ := json.Marshal(j)

	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	log.Println(rec.Body.String())
	assert.NotEqual(t, "", rec.Body.String())
}

func TestUserPut(t *testing.T) {
	router := server.NewServer()
	token := GetToken(email, password)

	name = "test2"

	var j = apifunc.UserPutParams{
		Name: name,
	}
	b, _ := json.Marshal(j)

	req := httptest.NewRequest("PUT", "/user", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var res map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, "success", res["message"])
}

func TestUserGet(t *testing.T) {
	router := server.NewServer()
	token := GetToken(email, password)

	req := httptest.NewRequest("GET", "/user", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var res map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	if err != nil {
		log.Println(err)
	}

	assert.Equal(t, name, res["name"])
}

func TestUserDelete(t *testing.T) {
	router := server.NewServer()
	token := GetToken(email, password)

	req := httptest.NewRequest("DELETE", "/user", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	var res map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	if err != nil {
		log.Println(err)
	}

	assert.Equal(t, "success", res["message"])
}
