package iapi

import (
	"os"
	"strconv"
	"testing"
)

var ICINGA2_API_USER = os.Getenv("ICINGA2_API_USER")
var ICINGA2_API_PASSWORD = os.Getenv("ICINGA2_API_PASSWORD")
var ICINGA2_API_URL = os.Getenv("ICINGA2_API_URL")
var ICINGA2_INSECURE_SKIP_TLS_VERIFY, _ = strconv.ParseBool(os.Getenv("ICINGA2_INSECURE_SKIP_TLS_VERIFY"))

var Icinga2_Server = Server{ICINGA2_API_USER, ICINGA2_API_PASSWORD, ICINGA2_API_URL, ICINGA2_INSECURE_SKIP_TLS_VERIFY, nil}

func TestConnect(t *testing.T) {

	v := os.Getenv("ICINGA2_API_URL")
	if v == "" {
		t.Fatal("ICINGA2_API_URL must be set for acceptance tests")
	}

	v = os.Getenv("ICINGA2_API_USER")
	if v == "" {
		t.Fatal("ICINGA2_API_USER must be set for acceptance tests")
	}

	v = os.Getenv("ICINGA2_API_PASSWORD")
	if v == "" {
		t.Fatal("ICINGA2_API_PASSWORD must be set for acceptance tests")
	}

	var Icinga2_Server = Server{"icinga-test", "icinga", ICINGA2_API_URL, ICINGA2_INSECURE_SKIP_TLS_VERIFY, nil}
	Icinga2_Server.Connect()

	if Icinga2_Server.httpClient == nil {
		t.Errorf("Failed to succesfully connect to Icinga Server")
	}
}

func TestConnectServerUnavailable(t *testing.T) {

	var Icinga2_Server = Server{"icinga-test", "icinga", "https://127.0.0.1:4665/v1", ICINGA2_INSECURE_SKIP_TLS_VERIFY, nil}
	err := Icinga2_Server.Connect()

	if err == nil {
		t.Errorf("Error : Did not get error connecting to unavailable server.")
	}
}

func TestConnectWithBadCredential(t *testing.T) {

	var Icinga2_Server = Server{"unknownUser", "unknownPW", ICINGA2_API_URL, ICINGA2_INSECURE_SKIP_TLS_VERIFY, nil}
	err := Icinga2_Server.Connect()
	if err != nil {
		t.Errorf("Did not fail with bad credentials : %s", err)
	}
}

func TestNewAPIRequest(t *testing.T) {

	result, _ := Icinga2_Server.NewAPIRequest("GET", "/status", nil)

	if result.Code != 200 {
		t.Errorf("%s", result.Status)
	}
}

func TestConnectServerBadURINoVersion(t *testing.T) {

	var Icinga2_Server = Server{ICINGA2_API_USER, ICINGA2_API_PASSWORD, "https://127.0.0.1:5665", ICINGA2_INSECURE_SKIP_TLS_VERIFY, nil}
	result, _ := Icinga2_Server.NewAPIRequest("GET", "/status", nil)

	if result.Code != 404 {
		t.Errorf("Error : Did not get expected 404 error connection to bad URI, with no version.")
	}
}
