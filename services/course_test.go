package services

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"student_rest/models"
	"student_rest/repositories"
	"testing"
)

type MocCourseRepository struct {
	mock.Mock
}

func (m *MocCourseRepository) CreateCourse(course *repositories.CourseEntity) (*models.CourseModel, error) {
	returnArgs := m.Called(course)
	return returnArgs.Get(0).(*models.CourseModel), returnArgs.Error(1)
}

func (m *MocCourseRepository) GetCourseByID(id string) (*models.CourseModel, error) {
	returnArgs := m.Called(id)
	return returnArgs.Get(0).(*models.CourseModel), returnArgs.Error(1)
}

func (m *MocCourseRepository) DeleteCourse(id string) error {
	returnArgs := m.Called(id)
	return returnArgs.Error(0)
}

func (m *MocCourseRepository) UpdateCourse(id string, course *repositories.CourseEntity) error {
	returnArgs := m.Called(id, course)
	return returnArgs.Error(0)
}

func Test_CreateCourse(t *testing.T) {
	testCases := []struct {
		name           string
		input          *models.CourseModel
		expectedValue  *models.CourseModel
		expectedError  error
		mockRepoInput  *repositories.CourseEntity
		mockRepoResult *models.CourseModel
		mockRepoError  error
	}{
		{
			name: "create course fail",
			input: &models.CourseModel{
				Teacher: &models.TeacherModel{
					ID: 1,
				},
			},
			expectedValue: nil,
			expectedError: errors.New("insert course fail"),
			mockRepoInput: &repositories.CourseEntity{
				TeacherID: 1,
			},
			mockRepoResult: nil,
			mockRepoError:  errors.New("insert course fail"),
		},
		{
			name: "create course successfully",
			input: &models.CourseModel{
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID: 1,
				},
			},
			expectedValue: &models.CourseModel{
				ID:        1,
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID:          1,
					FirstName:   "Anh",
					LastName:    "Le",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			expectedError: nil,
			mockRepoInput: &repositories.CourseEntity{
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				TeacherID: 1,
			},
			mockRepoResult: &models.CourseModel{
				ID:        1,
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID:          1,
					FirstName:   "Anh",
					LastName:    "Le",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MocCourseRepository)
			mockRepo.On("CreateCourse", testCase.mockRepoInput).Return(testCase.mockRepoResult, testCase.mockRepoError)

			courseService := Course{
				CourseRepositories: mockRepo,
			}

			result, err := courseService.CreateCourse(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedValue, result)
			}
		})
	}
}

func Test_GetCourseByID(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedValue  *models.CourseModel
		expectedError  error
		mockRepoInput  string
		mockRepoResult *models.CourseModel
		mockRepoError  error
	}{
		{
			name:           "get course by id fail",
			input:          "1",
			expectedValue:  nil,
			expectedError:  errors.New("get course fail"),
			mockRepoInput:  "1",
			mockRepoResult: nil,
			mockRepoError:  errors.New("get course fail"),
		},
		{
			name:  "get course by id successfully",
			input: "2",
			expectedValue: &models.CourseModel{
				ID:        1,
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID:          1,
					FirstName:   "Anh",
					LastName:    "Le",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			expectedError: nil,
			mockRepoInput: "2",
			mockRepoResult: &models.CourseModel{
				ID:        1,
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID:          1,
					FirstName:   "Anh",
					LastName:    "Le",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MocCourseRepository)
			mockRepo.On("GetCourseByID", testCase.mockRepoInput).Return(testCase.mockRepoResult, testCase.mockRepoError)

			courseService := Course{
				CourseRepositories: mockRepo,
			}

			result, err := courseService.GetCourseByID(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedValue, result)
			}
		})
	}
}

func Test_DeleteCourse(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
		mockRepoInput string
		mockRepoError error
	}{
		{
			name:          "delete course fail",
			input:         "1",
			expectedError: errors.New("delete course fail"),
			mockRepoInput: "1",
			mockRepoError: errors.New("delete course fail"),
		},
		{
			name:          "delete course successfully",
			input:         "2",
			expectedError: nil,
			mockRepoInput: "2",
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MocCourseRepository)
			mockRepo.On("DeleteCourse", testCase.mockRepoInput).Return(testCase.mockRepoError)

			courseService := Course{
				CourseRepositories: mockRepo,
			}

			err := courseService.DeleteCourse(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_UpdateCourse(t *testing.T) {
	testCases := []struct {
		name                 string
		inputID              string
		inputCourse         *models.CourseModel
		expectedError        error
		mockRepoInputID      string
		mockRepoInputCourse *repositories.CourseEntity
		mockRepoError        error
	}{
		{
			name:    "update course fail",
			inputID: "1",
			inputCourse: &models.CourseModel{
				Teacher: &models.TeacherModel{
					ID: 1,
				},
			},
			expectedError:   errors.New("update course fail"),
			mockRepoInputID: "1",
			mockRepoInputCourse: &repositories.CourseEntity{
				TeacherID: 1,
			},
			mockRepoError: errors.New("update course fail"),
		},
		{
			name:    "update course successfully",
			inputID: "2",
			inputCourse: &models.CourseModel{
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID: 1,
				},
			},
			expectedError:   nil,
			mockRepoInputID: "2",
			mockRepoInputCourse: &repositories.CourseEntity{
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				TeacherID: 1,
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MocCourseRepository)
			mockRepo.On("UpdateCourse", testCase.mockRepoInputID, testCase.mockRepoInputCourse).Return(testCase.mockRepoError)

			courseService := Course{
				CourseRepositories: mockRepo,
			}

			err := courseService.UpdateCourse(testCase.inputID, testCase.inputCourse)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
