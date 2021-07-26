package tests

import (
	"os"
	"net/http/httptest"
	"testing"
	"bytes"
	"net/http"
)


func init() {
	os.Setenv("APP_STATE", "test")

	// if status := services.DropCollection("nodes"); status != true {
    //     log.Fatal("Unable to drop nodes collection")
    // }
}



func TestRegisterNodes(t *testing.T) {

	var jsonStr = []byte(`{"IP_address":"123.0.0.1","Node_id":"0876237872873"}`)

	_, err := http.NewRequest("POST", "/nodes2", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(RegisterNodes)
	// handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// expected := `{"id":4,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}


func TestGetNodes(t *testing.T) {
	_, err := http.NewRequest("GET", "/nodes", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(GetNodes)
	// handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// // Check the response body is what we expect.
	// expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb@g.com","phone_number":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"FirstNameSample","last_name":"LastNameSample","email_address":"lr@gmail.com","phone_number":"1111111111"}]`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}