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
	"testing"
)

type MockCourseService struct {
	mock.Mock
}

func (m *MockCourseService) CreateCourse(course *models.CourseModel) (*models.CourseModel, error) {
	returnArgs := m.Called(course)
	return returnArgs.Get(0).(*models.CourseModel), returnArgs.Error(1)
}

func (m *MockCourseService) GetCourseByID(id string) (*models.CourseModel, error) {
	returnArgs := m.Called(id)
	return returnArgs.Get(0).(*models.CourseModel), returnArgs.Error(1)
}

func (m *MockCourseService) DeleteCourse(id string) error {
	returnArgs := m.Called(id)
	return returnArgs.Error(0)
}

func (m *MockCourseService) UpdateCourse(id string, course *models.CourseModel) error {
	returnArgs := m.Called(id, course)
	return returnArgs.Error(0)
}

func Test_CreateCourse(t *testing.T) {
	testCases := []struct {
		name                 string
		requestBody          map[string]interface{}
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     *models.CourseModel
		mockServiceResult    *models.CourseModel
		mockServiceError     error
	}{
		{
			name: "decode request body fail",
			requestBody: map[string]interface{}{
				"startTime": 1,
			},
			expectedResponseBody: "json: cannot unmarshal number into Go struct field CourseRequest.startTime of type string\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body fail",
			requestBody: map[string]interface{}{
				"startTime": "1998-11-02T00:00:00Z",
			},
			expectedResponseBody: "course name is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "create course fail",
			requestBody: map[string]interface{}{
				"name":      "Math",
				"startTime": "2020-11-02T00:00:00Z",
				"endTime":   "2020-11-03T00:00:00Z",
				"teacherID": 1,
			},
			expectedResponseBody: "create course fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput: &models.CourseModel{
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID: 1,
				},
			},
			mockServiceResult: nil,
			mockServiceError:  errors.New("create course fail"),
		},
		{
			name: "create course successfully",
			requestBody: map[string]interface{}{
				"name":      "Physics",
				"startTime": "2020-11-02T00:00:00Z",
				"endTime":   "2020-11-03T00:00:00Z",
				"teacherID": 1,
			},
			expectedResponseBody: "{\"success\":true,\"course\":{\"ID\":1,\"Name\":\"Physics\",\"StartTime\":\"2020-11-02T00:00:00Z\",\"EndTime\":\"2020-11-03T00:00:00Z\",\"Teacher\":{\"ID\":1,\"FirstName\":\"Mai\",\"LastName\":\"Dao\",\"DateOfBirth\":\"1998-11-02T00:00:00Z\"}}}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput: &models.CourseModel{
				Name:      "Physics",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID: 1,
				},
			},
			mockServiceResult: &models.CourseModel{
				ID:        1,
				Name:      "Physics",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID:          1,
					FirstName:   "Mai",
					LastName:    "Dao",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockCourseService)
			mockService.On("CreateCourse", testCase.mockServiceInput).Return(testCase.mockServiceResult, testCase.mockServiceError)

			courseHandler := CourseHandlers{
				CourseServices: mockService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(courseHandler.CreateCourse)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_GetCourseByID(t *testing.T) {
	testCases := []struct {
		name                 string
		paramID              string
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     string
		mockServiceResult    *models.CourseModel
		mockServiceError     error
	}{
		{
			name:                 "get course by id fail",
			paramID:              "1",
			expectedResponseBody: "get course fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput:     "1",
			mockServiceResult:    nil,
			mockServiceError:     errors.New("get course fail"),
		},
		{
			name:                 "get course by id successfully",
			paramID:              "2",
			expectedResponseBody: "{\"success\":true,\"course\":{\"ID\":1,\"Name\":\"Math\",\"StartTime\":\"2020-11-02T00:00:00Z\",\"EndTime\":\"2020-11-03T00:00:00Z\",\"Teacher\":{\"ID\":1,\"FirstName\":\"Mai\",\"LastName\":\"Dao\",\"DateOfBirth\":\"1998-11-02T00:00:00Z\"}}}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput:     "2",
			mockServiceResult: &models.CourseModel{
				ID:        1,
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID:          1,
					FirstName:   "Mai",
					LastName:    "Dao",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockCourseService)
			mockService.On("GetCourseByID", testCase.mockServiceInput).Return(testCase.mockServiceResult, testCase.mockServiceError)

			courseHandler := CourseHandlers{
				CourseServices: mockService,
			}

			req, err := http.NewRequest(http.MethodGet, "/courses/course/{id}", nil)
			if err != nil {
				t.Error(err)
			}

			chiCtx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			chiCtx.URLParams.Add("id", testCase.paramID)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(courseHandler.GetCourseByID)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_DeleteCourse(t *testing.T) {
	testCases := []struct {
		name                 string
		paramID              string
		expectedResponseBody string
		expectedStatus       int
		mockServiceInput     string
		mockServiceError     error
	}{
		{
			name:                 "delete course fail",
			paramID:              "1",
			expectedResponseBody: "delete course fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInput:     "1",
			mockServiceError:     errors.New("delete course fail"),
		},
		{
			name:                 "delete course successfully",
			paramID:              "2",
			expectedResponseBody: "{\"success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInput:     "2",
			mockServiceError:     nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockCourseService)
			mockService.On("DeleteCourse", testCase.mockServiceInput).Return(testCase.mockServiceError)

			courseHandler := CourseHandlers{
				CourseServices: mockService,
			}

			req, err := http.NewRequest(http.MethodDelete, "/courses/course/{id}", nil)
			if err != nil {
				t.Error(err)
			}

			chiCtx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			chiCtx.URLParams.Add("id", testCase.paramID)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(courseHandler.DeleteCourse)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func Test_UpdateCourse(t *testing.T) {
	testCases := []struct {
		name                   string
		paramID                string
		requestBody            map[string]interface{}
		expectedResponseBody   string
		expectedStatus         int
		mockServiceInputID     string
		mockServiceInputCourse *models.CourseModel
		mockServiceError       error
	}{
		{
			name:    "decode request body fail",
			paramID: "1",
			requestBody: map[string]interface{}{
				"startTime": 1,
			},
			expectedResponseBody: "json: cannot unmarshal number into Go struct field CourseRequest.startTime of type string\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name:    "validate request body fail",
			paramID: "1",
			requestBody: map[string]interface{}{
				"startTime": "2020-11-02T00:00:00Z",
			},
			expectedResponseBody: "course name is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name:    "update course fail",
			paramID: "1",
			requestBody: map[string]interface{}{
				"name":      "Math",
				"startTime": "2020-11-02T00:00:00Z",
				"endTime":   "2020-11-03T00:00:00Z",
				"teacherID": 1,
			},
			expectedResponseBody: "update course fail\n",
			expectedStatus:       http.StatusInternalServerError,
			mockServiceInputID:   "1",
			mockServiceInputCourse: &models.CourseModel{
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID: 1,
				},
			},
			mockServiceError: errors.New("update course fail"),
		},
		{
			name:    "update course successfully",
			paramID: "2",
			requestBody: map[string]interface{}{
				"name":      "Physics",
				"startTime": "2020-11-02T00:00:00Z",
				"endTime":   "2020-11-03T00:00:00Z",
				"teacherID": 1,
			},
			expectedResponseBody: "{\"success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockServiceInputID:   "2",
			mockServiceInputCourse: &models.CourseModel{
				Name:      "Physics",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID: 1,
				},
			},
			mockServiceError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(MockCourseService)
			mockService.On("UpdateCourse", testCase.mockServiceInputID, testCase.mockServiceInputCourse).Return(testCase.mockServiceError)

			courseHandler := CourseHandlers{
				CourseServices: mockService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}
			req, err := http.NewRequest(http.MethodPut, "/courses/course/{id}", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			chiCtx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			chiCtx.URLParams.Add("id", testCase.paramID)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(courseHandler.UpdateCourse)

			handler.ServeHTTP(rr, req)

			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}
