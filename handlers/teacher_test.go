package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"student_rest/models"
	"student_rest/repositories"
	"testing"
)

type MockTeacherService struct {
	mock.Mock
}

func (m *MockTeacherService) CreateTeacher(teacher *models.TeacherModel) (*repositories.TeacherEntity, error) {
	returnArgs := m.Called(teacher)
	return returnArgs.Get(0).(*repositories.TeacherEntity), returnArgs.Error(1)
}

func (m *MockTeacherService) GetTeacherByID(id string) (*repositories.TeacherEntity, error) {
	returnArgs := m.Called(id)
	return returnArgs.Get(0).(*repositories.TeacherEntity), returnArgs.Error(1)
}

func (m *MockTeacherService) DeleteTeacher(id string) error {
	returnArgs := m.Called(id)
	return returnArgs.Error(0)
}

func (m *MockTeacherService) UpdateTeacher(id string, teacher *models.TeacherModel) error {
	returnArgs := m.Called(id, teacher)
	return returnArgs.Error(0)
}

func Test_CreateTeacher(t *testing.T) {
	testCases := []struct {
		name                 string
		requestBody          map[string]interface{}
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     *models.TeacherModel
		mockServiceResult    *repositories.TeacherEntity
		mockServiceError     error
	}{
		{
			name: "decode request body fail",
			requestBody: map[string]interface{}{
				"firstName": 1,
			},
			expectedResponseBody: "json: cannot unmarshal number into Go struct field TeacherRequest.firstName of type string\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body fail",
			requestBody: map[string]interface{}{
				"firstName": "Dao",
			},
			expectedResponseBody: "last name is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "create teacher fail",
			requestBody: map[string]interface{}{
				"firstName":   "Mai",
				"lastName":    "Dao",
				"dateOfBirth": "1998-11-02T00:00:00Z",
			},
			expectedResponseBody: "create teacher fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput: &models.TeacherModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceResult: nil,
			mockServiceError:  errors.New("create teacher fail"),
		},
		{
			name: "create teacher successfully",
			requestBody: map[string]interface{}{
				"firstName":   "Mai",
				"lastName":    "Dao",
				"dateOfBirth": "1998-11-02T00:00:00Z",
			},
			expectedResponseBody: "{\"success\":true,\"teacher\":{\"id\":1,\"firstName\":\"Mai\",\"lastName\":\"Dao\",\"dateOfBirth\":\"1998-11-02T00:00:00Z\"}}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput: &models.TeacherModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceResult: &repositories.TeacherEntity{
				ID:          1,
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockTeacherService)
			mockService.On("CreateTeacher", testCase.mockServiceInput).Return(testCase.mockServiceResult, testCase.mockServiceError)

			teacherHandler := TeacherHandlers{
				TeacherServices: mockService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/teachers", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(teacherHandler.CreateTeacher)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_GetTeacherByID(t *testing.T) {
	testCases := []struct {
		name                 string
		paramID              string
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     string
		mockServiceResult    *repositories.TeacherEntity
		mockServiceError     error
	}{
		{
			name:                 "get teacher by id fail",
			paramID:              "1",
			expectedResponseBody: "get teacher fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput:     "1",
			mockServiceResult:    nil,
			mockServiceError:     errors.New("get teacher fail"),
		},
		{
			name:                 "get teacher by id successfully",
			paramID:              "2",
			expectedResponseBody: "{\"success\":true,\"teacher\":{\"id\":2,\"firstName\":\"Mai\",\"lastName\":\"Dao\",\"dateOfBirth\":\"1998-11-02T00:00:00Z\"}}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput:     "2",
			mockServiceResult: &repositories.TeacherEntity{
				ID:          2,
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockTeacherService)
			mockService.On("GetTeacherByID", testCase.mockServiceInput).Return(testCase.mockServiceResult, testCase.mockServiceError)

			teacherHandler := TeacherHandlers{
				TeacherServices: mockService,
			}

			req, err := http.NewRequest(http.MethodGet, "/teachers/teacher/{id}", nil)
			if err != nil {
				t.Error(err)
			}

			chiCtx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			chiCtx.URLParams.Add("id", testCase.paramID)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(teacherHandler.GetTeacherByID)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_DeleteTeacher(t *testing.T) {
	testCases := []struct {
		name                 string
		paramID              string
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     string
		mockServiceError     error
	}{
		{
			name:                 "delete teacher fail",
			paramID:              "1",
			expectedResponseBody: "delete teacher fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput:     "1",
			mockServiceError:     errors.New("delete teacher fail"),
		},
		{
			name:                 "delete teacher successfully",
			paramID:              "2",
			expectedResponseBody: "{\"success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput:     "2",
			mockServiceError:     nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockTeacherService)
			mockService.On("DeleteTeacher", testCase.mockServiceInput).Return(testCase.mockServiceError)

			teacherHandler := TeacherHandlers{
				TeacherServices: mockService,
			}

			req, err := http.NewRequest(http.MethodDelete, "/teachers/teacher/{id}", nil)
			if err != nil {
				t.Error(err)
			}

			chiCtx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			chiCtx.URLParams.Add("id", testCase.paramID)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(teacherHandler.DeleteTeacher)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_UpdateTeacher(t *testing.T) {
	testCases := []struct {
		name                    string
		paramID                 string
		requestBody             map[string]interface{}
		expectedResponseBody    string
		expectedStatus          int
		mockServiceInputID      string
		mockServiceInputTeacher *models.TeacherModel
		mockServiceError        error
	}{
		{
			name:    "decode request body fail",
			paramID: "1",
			requestBody: map[string]interface{}{
				"firstName": 1,
			},
			expectedResponseBody: "json: cannot unmarshal number into Go struct field TeacherRequest.firstName of type string\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name:    "validate request body fail",
			paramID: "1",
			requestBody: map[string]interface{}{
				"firstName": "Dao",
			},
			expectedResponseBody: "last name is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name:    "update teacher fail",
			paramID: "1",
			requestBody: map[string]interface{}{
				"firstName":   "Mai",
				"lastName":    "Dao",
				"dateOfBirth": "1998-11-02T00:00:00Z",
			},
			expectedResponseBody: "update teacher fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInputID:   "1",
			mockServiceInputTeacher: &models.TeacherModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceError: errors.New("update teacher fail"),
		},
		{
			name:    "update teacher successfully",
			paramID: "2",
			requestBody: map[string]interface{}{
				"firstName":   "Mai",
				"lastName":    "Dao",
				"dateOfBirth": "1998-11-02T00:00:00Z",
			},
			expectedResponseBody: "{\"success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInputID:   "2",
			mockServiceInputTeacher: &models.TeacherModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockTeacherService)
			mockService.On("UpdateTeacher", testCase.mockServiceInputID, testCase.mockServiceInputTeacher).Return(testCase.mockServiceError)

			teacherHandler := TeacherHandlers{
				TeacherServices: mockService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(http.MethodPut, "/teachers/teacher/{id}", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			chiCtx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			chiCtx.URLParams.Add("id", testCase.paramID)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(teacherHandler.UpdateTeacher)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}
