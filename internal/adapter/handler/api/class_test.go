package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestClassModel struct {
	ClassType string
	Duration  int
}

type MockClassRepository struct {
	db interface{}
}

type MockClassService struct {
	repo MockClassRepository
}

func TestHandlerCreateClass(t *testing.T) {
	class := TestClassModel{
		ClassType: "boxing",
		Duration:  60,
	}
	jsonData, err := json.Marshal(class)
	if err != nil {
		t.Fatal(err)
	}

	body := bytes.NewBuffer(jsonData)
	req := httptest.NewRequest(http.MethodPost, "/v1/class", body)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	// var mockDB interface{}
	// mockRepo := MockClassRepository{db: mockDB}
	// mockService := service.ClassService{repo: mo}
	// mockService := MockClassService{repo: mockRepo}
	mockCtx := &MockHTTPContext{t: t, ResponseRecorder: recorder}
	mockCtx.SetRequest(req)

	// handler := api.NewClassHandler(mockService)
	// handler.CreateClass(mockCtx)

	// if mockCtx.Result().StatusCode != http.StatusCreated {
	// t.Errorf("Expected status code 201, got %d", mockCtx.Code)
	// }
}
