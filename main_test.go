package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"."
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func TestLoginNonExistentUser(t *testing.T) {
	clearTable()
	var loginReq = []byte(`{"userid":"noexist","password":"1234"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginReq))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Invalid User Information" {
		t.Errorf("Expected the 'error' key of the response to be set to 'err'. Got '%s'", m["error"])
	}
}

func TestSignupAndLoginSuccess(t *testing.T) {
	clearTable()
	//signup
	var signUpReq = []byte(`{"id":1,"userid":"oyt","password":"1234"}`)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(signUpReq))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var loginReq = []byte(`{"id":1,"userid":"oyt","password":"1234"}`)
	secondReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginReq))
	secondReq.Header.Set("Content-Type", "application/json")
	secondResponse := executeRequest(secondReq)

	checkResponseCode(t, http.StatusOK, secondResponse.Code)
}

func TestPutAndGetService(t *testing.T) {
	clearTable()
	var putReqPayload = []byte(`{"id":1,"name":"test service","userid":"oyt"}`)
	putReq, _ := http.NewRequest("POST", "/services", bytes.NewBuffer(putReqPayload))
	putReq.Header.Set("Content-Type", "application/json")
	putResponse := executeRequest(putReq)
	checkResponseCode(t, http.StatusOK, putResponse.Code)

	getReq, _ := http.NewRequest("GET", "/services/oyt", bytes.NewBuffer([]byte("")))
	getResponse := executeRequest(getReq)

	checkResponseCode(t, http.StatusOK, getResponse.Code)
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
id SERIAL,
userid TEXT NOT NULL,
password TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS services 
(
id SERIAL,
name TEXT NOT NULL,
userid TEXT NOT NULL
);
`

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
	a.DB.Exec("DELETE FROM services")
	a.DB.Exec("ALTER SEQUENCE services_id_seq RESTART WITH 1")
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
