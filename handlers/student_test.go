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

type MockStudentService struct {
	mock.Mock
}

func (m *MockStudentService) CreateStudent(student *models.StudentModel) (*repositories.StudentEntity, error) {
	returnArgs := m.Called(student)
	return returnArgs.Get(0).(*repositories.StudentEntity), returnArgs.Error(1)
}

func (m *MockStudentService) GetStudentByID(id string) (*repositories.StudentEntity, error) {
	returnArgs := m.Called(id)
	return returnArgs.Get(0).(*repositories.StudentEntity), returnArgs.Error(1)
}

func (m *MockStudentService) DeleteStudent(id string) error {
	returnArgs := m.Called(id)
	return returnArgs.Error(0)
}

func (m *MockStudentService) UpdateStudent(id string, student *models.StudentModel) error {
	returnArgs := m.Called(id, student)
	return returnArgs.Error(0)
}

func (m *MockStudentService) RegisterCourse(registerCourseModel *models.RegisterCourseModel) (*models.RegisterCourseModel, error) {
	returnArgs := m.Called(registerCourseModel)
	return returnArgs.Get(0).(*models.RegisterCourseModel), returnArgs.Error(1)

}

func Test_CreateStudent(t *testing.T) {
	testCases := []struct {
		name                 string
		requestBody          map[string]interface{}
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     *models.StudentModel
		mockServiceResult    *repositories.StudentEntity
		mockServiceError     error
	}{
		{
			name: "decode request body fail",
			requestBody: map[string]interface{}{
				"firstName": 1,
			},
			expectedResponseBody: "json: cannot unmarshal number into Go struct field StudentRequest.firstName of type string\n",
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
			name: "create student fail",
			requestBody: map[string]interface{}{
				"firstName":   "Mai",
				"lastName":    "Dao",
				"dateOfBirth": "1998-11-02T00:00:00Z",
			},
			expectedResponseBody: "create student fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput: &models.StudentModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceResult: nil,
			mockServiceError:  errors.New("create student fail"),
		},
		{
			name: "create student successfully",
			requestBody: map[string]interface{}{
				"firstName":   "Mai",
				"lastName":    "Dao",
				"dateOfBirth": "1998-11-02T00:00:00Z",
			},
			expectedResponseBody: "{\"success\":true,\"student\":{\"id\":1,\"studentID\":\"123456\",\"firstName\":\"Mai\",\"lastName\":\"Dao\",\"dateOfBirth\":\"1998-11-02T00:00:00Z\"}}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput: &models.StudentModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceResult: &repositories.StudentEntity{
				ID:          1,
				StudentID:   "123456",
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockStudentService)
			mockService.On("CreateStudent", testCase.mockServiceInput).Return(testCase.mockServiceResult, testCase.mockServiceError)

			studentHandler := StudentHandlers{
				StudentServices: mockService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/students", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(studentHandler.CreateStudent)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_GetStudentByID(t *testing.T) {
	testCases := []struct {
		name                 string
		paramID                   string
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     string
		mockServiceResult    *repositories.StudentEntity
		mockServiceError     error
	}{
		{
			name:                 "get student by id fail",
			paramID:                   "1",
			expectedResponseBody: "get student fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput:     "1",
			mockServiceResult:    nil,
			mockServiceError:     errors.New("get student fail"),
		},
		{
			name:                 "get student by id successfully",
			paramID:                   "2",
			expectedResponseBody: "{\"success\":true,\"student\":{\"id\":2,\"studentID\":\"123456\",\"firstName\":\"Mai\",\"lastName\":\"Dao\",\"dateOfBirth\":\"1998-11-02T00:00:00Z\"}}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput:     "2",
			mockServiceResult: &repositories.StudentEntity{
				ID:          2,
				StudentID:   "123456",
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockStudentService)
			mockService.On("GetStudentByID", testCase.mockServiceInput).Return(testCase.mockServiceResult, testCase.mockServiceError)

			studentHandler := StudentHandlers{
				StudentServices: mockService,
			}

			req, err := http.NewRequest(http.MethodGet, "/students/student/{id}", nil)
			if err != nil {
				t.Error(err)
			}

			chiCtx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			chiCtx.URLParams.Add("id", testCase.paramID)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(studentHandler.GetStudentByID)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_DeleteStudent(t *testing.T) {
	testCases := []struct {
		name                 string
		paramID                   string
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     string
		mockServiceError     error
	}{
		{
			name:                 "delete student fail",
			paramID:                   "1",
			expectedResponseBody: "delete student fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput:     "1",
			mockServiceError:     errors.New("delete student fail"),
		},
		{
			name:                 "delete student successfully",
			paramID:                   "2",
			expectedResponseBody: "{\"success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput:     "2",
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockStudentService)
			mockService.On("DeleteStudent", testCase.mockServiceInput).Return(testCase.mockServiceError)

			studentHandler := StudentHandlers{
				StudentServices: mockService,
			}

			req, err := http.NewRequest(http.MethodDelete, "/students/student/{id}", nil)
			if err != nil {
				t.Error(err)
			}

			chiCtx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			chiCtx.URLParams.Add("id", testCase.paramID)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(studentHandler.DeleteStudent)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_UpdateStudent(t *testing.T) {
	testCases := []struct {
		name                 string
		paramID                   string
		requestBody          map[string]interface{}
		expectedResponseBody string
		expectedStatus       int
		mockServiceInputID     string
		mockServiceInputStudent *models.StudentModel
		mockServiceError     error
	}{
		{
			name: "decode request body fail",
			paramID: "1",
			requestBody: map[string]interface{}{
				"firstName": 1,
			},
			expectedResponseBody: "json: cannot unmarshal number into Go struct field StudentRequest.firstName of type string\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body fail",
			paramID: "1",
			requestBody: map[string]interface{}{
				"firstName": "Dao",
			},
			expectedResponseBody: "last name is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "update student fail",
			paramID: "1",
			requestBody: map[string]interface{}{
				"firstName":   "Mai",
				"lastName":    "Dao",
				"dateOfBirth": "1998-11-02T00:00:00Z",
			},
			expectedResponseBody: "update student fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInputID: "1",
			mockServiceInputStudent: &models.StudentModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceError:  errors.New("update student fail"),
		},
		{
			name: "update student successfully",
			paramID: "2",
			requestBody: map[string]interface{}{
				"firstName":   "Mai",
				"lastName":    "Dao",
				"dateOfBirth": "1998-11-02T00:00:00Z",
			},
			expectedResponseBody: "{\"success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInputID: "2",
			mockServiceInputStudent: &models.StudentModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockStudentService)
			mockService.On("UpdateStudent", testCase.mockServiceInputID, testCase.mockServiceInputStudent).Return(testCase.mockServiceError)

			studentHandler := StudentHandlers{
				StudentServices: mockService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(http.MethodPut, "/students/student/{id}", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			chiCtx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			chiCtx.URLParams.Add("id", testCase.paramID)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(studentHandler.UpdateStudent)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_RegisterCourse(t *testing.T) {
	testCases := []struct {
		name                 string
		requestBody          map[string]interface{}
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     *models.RegisterCourseModel
		mockServiceResult    *models.RegisterCourseModel
		mockServiceError     error
	}{
		{
			name: "decode request body fail",
			requestBody: map[string]interface{}{
				"student": 1,
			},
			expectedResponseBody: "json: cannot unmarshal number into Go struct field RegisterCourseRequest.Student of type handlers.StudentRequest\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body fail",
			requestBody: map[string]interface{}{
				"student": map[string]interface{} {
					"firstName":   "Mai",
					"lastName":    "Dao",
					"dateOfBirth": "1998-11-02T00:00:00Z",
				},
			},
			expectedResponseBody: "course is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "register course fail",
			requestBody: map[string]interface{}{
				"student": map[string]interface{} {
					"firstName":   "Dao",
					"lastName":    "Mai",
					"dateOfBirth": "1998-11-02T00:00:00Z",
				},
				"course": map[string]interface{} {
					"name": "Math",
					"startTime":"1998-11-02T00:00:00Z",
					"endTime": "1998-11-02T00:00:00Z",
					"teacher": map[string]interface{}{
						"firstName": "Dao",
						"lastName": "Mai",
						"dateOfBirth": "1998-11-02T00:00:00Z",
					},
				},
			},
			expectedResponseBody: "register course fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					FirstName:   "Dao",
					LastName:    "Mai",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
				Course: &models.CourseModel{
					Name:      "Math",
					StartTime: "1998-11-02T00:00:00Z",
					EndTime:   "1998-11-02T00:00:00Z",
					Teacher: &models.TeacherModel{
						FirstName:   "Dao",
						LastName:    "Mai",
						DateOfBirth: "1998-11-02T00:00:00Z",
					},
				},
			},
			mockServiceResult: nil,
			mockServiceError:  errors.New("register course fail"),
		},
		{
			name: "create student successfully",
			requestBody: map[string]interface{}{
				"student": map[string]interface{} {
					"firstName":   "Dao",
					"lastName":    "Mai",
					"dateOfBirth": "1998-11-02T00:00:00Z",
				},
				"course": map[string]interface{} {
					"name": "Math",
					"startTime":"1998-11-02T00:00:00Z",
					"endTime": "1998-11-02T00:00:00Z",
					"teacher": map[string]interface{}{
						"firstName": "Dao",
						"lastName": "Mai",
						"dateOfBirth": "1998-11-02T00:00:00Z",
					},
				},
			},
			expectedResponseBody: "{\"success\":true,\"course\":{\"ID\":1,\"Name\":\"Math\",\"StartTime\":\"1998-11-02T00:00:00Z\",\"EndTime\":\"1998-11-02T00:00:00Z\",\"Teacher\":{\"ID\":1,\"FirstName\":\"Dao\",\"LastName\":\"Mai\",\"DateOfBirth\":\"1998-11-02T00:00:00Z\"}},\"student\":{\"ID\":1,\"StudentID\":\"123456\",\"FirstName\":\"Dao\",\"LastName\":\"Mai\",\"DateOfBirth\":\"1998-11-02T00:00:00Z\"}}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					FirstName:   "Dao",
					LastName:    "Mai",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
				Course: &models.CourseModel{
					Name:      "Math",
					StartTime: "1998-11-02T00:00:00Z",
					EndTime:   "1998-11-02T00:00:00Z",
					Teacher: &models.TeacherModel{
						FirstName:   "Dao",
						LastName:    "Mai",
						DateOfBirth: "1998-11-02T00:00:00Z",
					},
				},
			},
			mockServiceResult: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					ID: 1,
					StudentID: "123456",
					FirstName:   "Dao",
					LastName:    "Mai",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
				Course: &models.CourseModel{
					ID: 1,
					Name:      "Math",
					StartTime: "1998-11-02T00:00:00Z",
					EndTime:   "1998-11-02T00:00:00Z",
					Teacher: &models.TeacherModel{
						ID: 1,
						FirstName:   "Dao",
						LastName:    "Mai",
						DateOfBirth: "1998-11-02T00:00:00Z",
					},
				},
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockStudentService)
			mockService.On("RegisterCourse", testCase.mockServiceInput).Return(testCase.mockServiceResult, testCase.mockServiceError)

			studentHandler := StudentHandlers{
				StudentServices: mockService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(http.MethodGet, "/students/register-course", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(studentHandler.RegisterCourse)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

