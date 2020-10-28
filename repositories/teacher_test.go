package repositories

import (
	"errors"
	"github.com/stretchr/testify/require"
	"student_rest/testhelpers"
	"student_rest/utils"
	"testing"
)

func Test_CreateTeacher(t *testing.T) {
	testCases := []struct {
		name          string
		input         *TeacherEntity
		expectedValue *TeacherEntity
		expectedError error
		giveFixture   string
	}{
		{
			name: "insert teacher fail",
			input: &TeacherEntity{
				FirstName:   "Dao",
				LastName:    "Mai",
			},
			expectedValue: nil,
			expectedError: errors.New("pq: invalid input syntax for type timestamp: \"\""),
			giveFixture:   "./testdata/truncate_data.sql",
		},
		{
			name: "insert teacher successfully",
			input: &TeacherEntity{
				FirstName:   "Dao",
				LastName:    "Mai",
				DateOfBirth: "6/25/1997",
			},
			expectedValue: &TeacherEntity{
				FirstName:   "Dao",
				LastName:    "Mai",
				DateOfBirth: "1997-06-25T00:00:00Z",
			},
			expectedError: nil,
			giveFixture:   "./testdata/truncate_data.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			teacherRepo := Teacher{
				Db: dbMock,
			}

			result, err := teacherRepo.CreateTeacher(testCase.input)

			if testCase.expectedError != nil {
				// For Fail Logic
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				// For Success Logic
				require.NoError(t, err)

				sqlStmt := `SELECT * FROM teachers WHERE id=$1`
				var student StudentEntity
				err := dbMock.QueryRow(sqlStmt, result.ID).Scan(&student.ID, &student.FirstName, &student.LastName, &student.DateOfBirth)
				if err != nil {
					t.Error(err)
				}
				require.Equal(t, testCase.expectedValue.FirstName, student.FirstName)
				require.Equal(t, testCase.expectedValue.LastName, student.LastName)
				require.Equal(t, testCase.expectedValue.DateOfBirth, student.DateOfBirth)
			}
		})
	}
}

func Test_GetTeacherByID(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedValue *TeacherEntity
		expectedError error
		giveFixture   string
	}{
		{
			name:          "get teacher by id fail",
			input:         "3",
			expectedValue: nil,
			expectedError: errors.New("sql: no rows in result set"),
			giveFixture:   "./testdata/teacher/teacher.sql",
		},
		{
			name:  "get teacher by id successfully",
			input: "2",
			expectedValue: &TeacherEntity{
				ID:          2,
				FirstName:   "Duyen",
				LastName:    "Nguyen",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedError: nil,
			giveFixture:   "./testdata/teacher/teacher.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			teacherRepo := Teacher{
				Db: dbMock,
			}

			result, err := teacherRepo.GetTeacherByID(testCase.input)

			if testCase.expectedError != nil {
				// For Fail Logic
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				// For Success Logic
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
		giveFixture   string
	}{
		{
			name:          "delete teacher fail",
			input:         "1",
			expectedError: errors.New("pq: update or delete on table \"teachers\" violates foreign key constraint \"courses_teacher_id_fkey\" on table \"courses\""),
			giveFixture:   "./testdata/teacher/teacher.sql",
		},
		{
			name:          "delete teacher successfully",
			input:         "2",
			expectedError: nil,
			giveFixture:   "./testdata/teacher/teacher.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			teacherRepo := Teacher{
				Db: dbMock,
			}

			err := teacherRepo.DeleteTeacher(testCase.input)

			if testCase.expectedError != nil {
				// For Fail Logic
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				// For Success Logic
				require.NoError(t, err)
			}
		})
	}
}

func Test_UpdateTeacher(t *testing.T) {
	testCases := []struct {
		name          string
		inputID         string
		inputStudent *TeacherEntity
		expectedError error
		giveFixture   string
	}{
		{
			name:  "update teacher fail",
			inputID: "1",
			inputStudent: &TeacherEntity{
				FirstName:   "Dao",
				LastName:    "Mai",
			},
			expectedError: errors.New("pq: invalid input syntax for type timestamp: \"\""),
			giveFixture:   "./testdata/teacher/teacher.sql",
		},
		{
			name:  "update teacher successfully",
			inputID: "1",
			inputStudent: &TeacherEntity{
				FirstName:   "Dao",
				LastName:    "Mai",
				DateOfBirth: "1/12/1997",
			},
			expectedError: nil,
			giveFixture:   "./testdata/teacher/teacher.sql",
		},

	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			teacherRepo := Teacher{
				Db: dbMock,
			}

			err := teacherRepo.UpdateTeacher(testCase.inputID, testCase.inputStudent)

			if testCase.expectedError != nil {
				// For Fail Logic
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				// For Success Logic
				require.NoError(t, err)
			}
		})
	}
}
