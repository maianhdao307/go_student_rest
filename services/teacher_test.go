package services

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"student_rest/models"
	"student_rest/repositories"
	"testing"
)

type MockTeacherRepository struct {
	mock.Mock
}

func (m *MockTeacherRepository) CreateTeacher(teacher *repositories.TeacherEntity) (*repositories.TeacherEntity, error) {
	returnArgs := m.Called(teacher)
	return returnArgs.Get(0).(*repositories.TeacherEntity), returnArgs.Error(1)
}

func (m *MockTeacherRepository) GetTeacherByID(id string) (*repositories.TeacherEntity, error) {
	returnArgs := m.Called(id)
	return returnArgs.Get(0).(*repositories.TeacherEntity), returnArgs.Error(1)
}

func (m *MockTeacherRepository) DeleteTeacher(id string) error {
	returnArgs := m.Called(id)
	return returnArgs.Error(0)
}

func (m *MockTeacherRepository) UpdateTeacher(id string, teacher *repositories.TeacherEntity) error {
	returnArgs := m.Called(id, teacher)
	return returnArgs.Error(0)
}

func Test_CreateTeacher(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.TeacherModel
		expectedValue *repositories.TeacherEntity
		expectedError error
		mockRepoInput *repositories.TeacherEntity
		mockRepoResult *repositories.TeacherEntity
		mockRepoError error
	}{
		{
			name:          "create teacher fail",
			input:         &models.TeacherModel{},
			expectedValue: nil,
			expectedError: errors.New("insert teacher fail"),
			mockRepoInput: &repositories.TeacherEntity{},
			mockRepoResult: nil,
			mockRepoError: errors.New("insert teacher fail"),
		},
		{
			name: "create teacher successfully",
			input: &models.TeacherModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedValue: &repositories.TeacherEntity{
				ID:          2,
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedError: nil,
			mockRepoInput: &repositories.TeacherEntity{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockRepoResult: &repositories.TeacherEntity{
				ID:          2,
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MockTeacherRepository)
			mockRepo.On("CreateTeacher", testCase.mockRepoInput).Return(testCase.mockRepoResult, testCase.mockRepoError)

			teacherService := Teacher{
				TeacherRepositories: mockRepo,
			}

			result, err := teacherService.CreateTeacher(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedValue, result)
			}
		})
	}
}

func Test_GetTeacherByID(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedValue *repositories.TeacherEntity
		expectedError error
		mockRepoInput string
		mockRepoResult *repositories.TeacherEntity
		mockRepoError error
	}{
		{
			name:          "get teacher by id fail",
			input:         "1",
			expectedValue: nil,
			expectedError: errors.New("get teacher fail"),
			mockRepoInput: "1",
			mockRepoResult: nil,
			mockRepoError: errors.New("get teacher fail"),
		},
		{
			name:  "get teacher by id successfully",
			input: "2",
			expectedValue: &repositories.TeacherEntity{
				ID:          2,
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedError: nil,
			mockRepoInput: "2",
			mockRepoResult: &repositories.TeacherEntity{
				ID:          2,
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MockTeacherRepository)
			mockRepo.On("GetTeacherByID", testCase.mockRepoInput).Return(testCase.mockRepoResult, testCase.mockRepoError)

			teacherService := Teacher{
				TeacherRepositories: mockRepo,
			}

			result, err := teacherService.GetTeacherByID(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedValue, result)
			}
		})
	}
}

func Test_DeleteTeacher(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
		mockRepoInput string
		mockRepoError error
	}{
		{
			name:          "delete teacher fail",
			input:         "1",
			expectedError: errors.New("delete teacher fail"),
			mockRepoInput: "1",
			mockRepoError: errors.New("delete teacher fail"),
		},
		{
			name:          "delete teacher successfully",
			input:         "2",
			expectedError: nil,
			mockRepoInput: "2",
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MockTeacherRepository)
			mockRepo.On("DeleteTeacher", testCase.mockRepoInput).Return(testCase.mockRepoError)

			teacherService := Teacher{
				TeacherRepositories: mockRepo,
			}

			err := teacherService.DeleteTeacher(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_UpdateTeacher(t *testing.T) {
	testCases := []struct {
		name                 string
		inputID              string
		inputTeacher         *models.TeacherModel
		expectedError        error
		mockRepoInputID      string
		mockRepoInputTeacher *repositories.TeacherEntity
		mockRepoError        error
	}{
		{
			name:                 "update teacher fail",
			inputID:              "1",
			inputTeacher:         &models.TeacherModel{},
			expectedError:        errors.New("update teacher fail"),
			mockRepoInputID:      "1",
			mockRepoInputTeacher: &repositories.TeacherEntity{},
			mockRepoError:        errors.New("update teacher fail"),
		},
		{
			name:    "update teacher successfully",
			inputID: "2",
			inputTeacher: &models.TeacherModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedError:   nil,
			mockRepoInputID: "2",
			mockRepoInputTeacher: &repositories.TeacherEntity{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MockTeacherRepository)
			mockRepo.On("UpdateTeacher", testCase.mockRepoInputID, testCase.mockRepoInputTeacher).Return(testCase.mockRepoError)

			teacherService := Teacher{
				TeacherRepositories: mockRepo,
			}

			err := teacherService.UpdateTeacher(testCase.inputID, testCase.inputTeacher)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

