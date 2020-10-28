package services

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"student_rest/models"
	"student_rest/repositories"
	"testing"
)

type MockStudentRepository struct {
	mock.Mock
}

func (m *MockStudentRepository) CreateStudent(student *repositories.StudentEntity) (*repositories.StudentEntity, error) {
	returnArgs := m.Called(student)
	return returnArgs.Get(0).(*repositories.StudentEntity), returnArgs.Error(1)
}

func (m *MockStudentRepository) GetStudentByID(id string) (*repositories.StudentEntity, error) {
	returnArgs := m.Called(id)
	return returnArgs.Get(0).(*repositories.StudentEntity), returnArgs.Error(1)
}

func (m *MockStudentRepository) DeleteStudent(id string) error {
	returnArgs := m.Called(id)
	return returnArgs.Error(0)
}

func (m *MockStudentRepository) UpdateStudent(id string, student *repositories.StudentEntity) error {
	returnArgs := m.Called(id, student)
	return returnArgs.Error(0)
}

func (m *MockStudentRepository) RegisterCourse(registerCourseModel *models.RegisterCourseModel) (*models.RegisterCourseModel, error) {
	returnArgs := m.Called(registerCourseModel)
	return returnArgs.Get(0).(*models.RegisterCourseModel), returnArgs.Error(1)
}

type MockUtil struct {
	mock.Mock
}

func (m *MockUtil) GenerateID(regex string, limit int) (string, error) {
	returnArgs := m.Called(regex, limit)
	return returnArgs.String(0), returnArgs.Error(1)
}

func Test_CreateStudent(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.StudentModel
		expectedValue *repositories.StudentEntity
		expectedError error
		mockRepoInput *repositories.StudentEntity
		mockRepoResult *repositories.StudentEntity
		mockRepoError error
		mockGenerateIDError error
	}{
		{
			name:          "generate student id fail",
			input:         &models.StudentModel{},
			expectedValue: nil,
			expectedError: errors.New("generate id fail"),
			mockGenerateIDError: errors.New("generate id fail"),
		},
		{
			name:          "create student fail",
			input:         &models.StudentModel{},
			expectedValue: nil,
			expectedError: errors.New("insert student fail"),
			mockRepoInput: &repositories.StudentEntity{StudentID: "123456"},
			mockRepoResult: nil,
			mockRepoError: errors.New("insert student fail"),
			mockGenerateIDError: nil,
		},
		{
			name: "create student successfully",
			input: &models.StudentModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedValue: &repositories.StudentEntity{
				ID:          2,
				StudentID:   "123456",
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedError: nil,
			mockRepoInput: &repositories.StudentEntity{
				StudentID:   "123456",
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockRepoResult: &repositories.StudentEntity{
				ID:          2,
				StudentID:   "123456",
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockRepoError: nil,
			mockGenerateIDError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockUtil := new(MockUtil)
			mockUtil.On("GenerateID", "[A-Z0-9]{6}", 6).Return("123456", testCase.mockGenerateIDError)

			mockRepo := new(MockStudentRepository)
			mockRepo.On("CreateStudent", testCase.mockRepoInput).Return(testCase.mockRepoResult, testCase.mockRepoError)

			studentService := Student{
				StudentRepositories: mockRepo,
				Utils:               mockUtil,
			}

			result, err := studentService.CreateStudent(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedValue, result)
			}
		})
	}
}

func Test_GetStudentByID(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedValue *repositories.StudentEntity
		expectedError error
		mockRepoInput string
		mockRepoResult *repositories.StudentEntity
		mockRepoError error
	}{
		{
			name:          "get student by id fail",
			input:         "1",
			expectedValue: nil,
			expectedError: errors.New("get student fail"),
			mockRepoInput: "1",
			mockRepoResult: nil,
			mockRepoError: errors.New("get student fail"),
		},
		{
			name:  "get student by id successfully",
			input: "2",
			expectedValue: &repositories.StudentEntity{
				ID:          2,
				StudentID:   "234567",
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedError: nil,
			mockRepoInput: "2",
			mockRepoResult: &repositories.StudentEntity{
				ID:          2,
				StudentID:   "234567",
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MockStudentRepository)
			mockRepo.On("GetStudentByID", testCase.mockRepoInput).Return(testCase.mockRepoResult, testCase.mockRepoError)

			studentService := Student{
				StudentRepositories: mockRepo,
			}

			result, err := studentService.GetStudentByID(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedValue, result)
			}
		})
	}
}

func Test_DeleteStudent(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
		mockRepoInput string
		mockRepoError error
	}{
		{
			name:          "delete student fail",
			input:         "1",
			expectedError: errors.New("delete student fail"),
			mockRepoInput: "1",
			mockRepoError: errors.New("delete student fail"),
		},
		{
			name:          "delete student successfully",
			input:         "2",
			expectedError: nil,
			mockRepoInput: "2",
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MockStudentRepository)
			mockRepo.On("DeleteStudent", testCase.mockRepoInput).Return(testCase.mockRepoError)

			studentService := Student{
				StudentRepositories: mockRepo,
			}

			err := studentService.DeleteStudent(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_UpdateStudent(t *testing.T) {
	testCases := []struct {
		name                 string
		inputID              string
		inputStudent         *models.StudentModel
		expectedError        error
		mockRepoInputID      string
		mockRepoInputStudent *repositories.StudentEntity
		mockRepoError        error
	}{
		{
			name:                 "update student fail",
			inputID:              "1",
			inputStudent:         &models.StudentModel{},
			expectedError:        errors.New("update student fail"),
			mockRepoInputID:      "1",
			mockRepoInputStudent: &repositories.StudentEntity{},
			mockRepoError:        errors.New("update student fail"),
		},
		{
			name:    "update student successfully",
			inputID: "2",
			inputStudent: &models.StudentModel{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedError:   nil,
			mockRepoInputID: "2",
			mockRepoInputStudent: &repositories.StudentEntity{
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(MockStudentRepository)
			mockRepo.On("UpdateStudent", testCase.mockRepoInputID, testCase.mockRepoInputStudent).Return(testCase.mockRepoError)

			studentService := Student{
				StudentRepositories: mockRepo,
			}

			err := studentService.UpdateStudent(testCase.inputID, testCase.inputStudent)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_RegisterCourse(t *testing.T) {
	testCases := []struct {
		name                string
		input               *models.RegisterCourseModel
		expectedValue       *models.RegisterCourseModel
		expectedError       error
		mockRepoInput       *models.RegisterCourseModel
		mockRepoResult      *models.RegisterCourseModel
		mockRepoError       error
		mockGenerateIDError error
	}{
		{
			name: "generate id fail",
			input: &models.RegisterCourseModel{
				Student: &models.StudentModel{},
			},
			expectedValue:       nil,
			expectedError:       errors.New("generate id fail"),
			mockGenerateIDError: errors.New("generate id fail"),
		},
		{
			name: "register course fail",
			input: &models.RegisterCourseModel{
				Student: &models.StudentModel{},
			},
			expectedValue: nil,
			expectedError: errors.New("register course fail"),
			mockRepoInput: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					StudentID: "123456",
				},
			},
			mockRepoResult:      nil,
			mockRepoError:       errors.New("register course fail"),
			mockGenerateIDError: nil,
		},
		{
			name: "register course successfully",
			input: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					FirstName:   "Mai",
					LastName:    "Dao",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			expectedValue: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					ID:          1,
					StudentID:   "123456",
					FirstName:   "Mai",
					LastName:    "Dao",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			expectedError: nil,
			mockRepoInput: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					StudentID:   "123456",
					FirstName:   "Mai",
					LastName:    "Dao",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			mockRepoResult: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					ID:          1,
					StudentID:   "123456",
					FirstName:   "Mai",
					LastName:    "Dao",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			mockRepoError:       nil,
			mockGenerateIDError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Mock repo
			mockRepo := new(MockStudentRepository)
			mockRepo.On("RegisterCourse", testCase.mockRepoInput).Return(testCase.mockRepoResult, testCase.mockRepoError)

			// Mock generate id
			mockUtil := new(MockUtil)
			mockUtil.On("GenerateID", "[A-Z0-9]{6}", 6).Return("123456", testCase.mockGenerateIDError)

			studentService := Student{
				StudentRepositories: mockRepo,
				Utils:               mockUtil,
			}

			result, err := studentService.RegisterCourse(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedValue, result)
			}
		})
	}
}
