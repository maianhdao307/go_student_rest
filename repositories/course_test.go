package repositories

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/require"
	"student_rest/models"
	"student_rest/testhelpers"
	"student_rest/utils"
	"testing"
)

func Test_CreateCourse(t *testing.T) {
	testCases := []struct {
		name          string
		input         *CourseEntity
		expectedValue *models.CourseModel
		expectedError error
		giveFixture   string
	}{
		{
			name: "insert course fail",
			input: &CourseEntity{
				Name:      "Math",
				StartTime: "4/23/2020",
				EndTime:   "4/24/2020",
			},
			expectedValue: nil,
			expectedError: errors.New("pq: insert or update on table \"courses\" violates foreign key constraint \"courses_teacher_id_fkey\""),
			giveFixture:   "./testdata/truncate_data.sql",
		},
		{
			name: "insert course successfully",
			input: &CourseEntity{
				Name:      "Math",
				StartTime: "4/23/2020",
				EndTime:   "4/24/2020",
				TeacherID: 1,
			},
			expectedValue: &models.CourseModel{
				Name:      "Math",
				StartTime: "2020-04-23T00:00:00Z",
				EndTime:   "2020-04-24T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID: 1,
				},
			},
			expectedError: nil,
			giveFixture:   "./testdata/teacher/teacher.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			courseRepo := Course{
				Db: dbMock,
			}

			result, err := courseRepo.CreateCourse(testCase.input)

			if testCase.expectedError != nil {
				// For Fail Logic
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				// For Success Logic
				require.NoError(t, err)

				sqlStmt := `SELECT * FROM courses WHERE id=$1`
				var course CourseEntity
				err := dbMock.QueryRow(sqlStmt, result.ID).Scan(&course.ID, &course.Name, &course.StartTime, &course.EndTime, &course.TeacherID)
				if err != nil {
					t.Error(err)
				}
				require.Equal(t, testCase.expectedValue.Name, course.Name)
				require.Equal(t, testCase.expectedValue.StartTime, course.StartTime)
				require.Equal(t, testCase.expectedValue.EndTime, course.EndTime)
				require.Equal(t, testCase.expectedValue.Teacher.ID, course.TeacherID)
			}
		})
	}
}

func Test_GetCourseByID(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedValue *models.CourseModel
		expectedError error
		giveFixture   string
	}{
		{
			name:          "get course by id fail",
			input:         "3",
			expectedValue: nil,
			expectedError: errors.New("sql: no rows in result set"),
			giveFixture:   "./testdata/course/course.sql",
		},
		{
			name:  "get course by id successfully",
			input: "1",
			expectedValue: &models.CourseModel{
				ID:        1,
				Name:      "Math",
				StartTime: "2020-11-02T00:00:00Z",
				EndTime:   "2020-11-03T00:00:00Z",
				Teacher: &models.TeacherModel{
					ID: 1,
					FirstName: "Anh",
					LastName: "Le",
					DateOfBirth: "1998-11-02T00:00:00Z",
				},
			},
			expectedError: nil,
			giveFixture:   "./testdata/course/course.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			courseRepo := Course{
				Db: dbMock,
			}

			result, err := courseRepo.GetCourseByID(testCase.input)

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

func Test_DeleteCourse(t *testing.T) {
	dbMockFail, _ := testhelpers.ConnectDBFailed()

	testCases := []struct {
		name          string
		input         string
		expectedError error
		giveFixture   string
		dbMock *sql.DB
	}{
		{
			name:          "delete course fail",
			input:         "1",
			expectedError: errors.New("pq: password authentication failed for user \"postgres\""),
			giveFixture:   "./testdata/course/course.sql",
			dbMock: dbMockFail,
		},
		{
			name:          "delete course successfully",
			input:         "1",
			expectedError: nil,
			giveFixture:   "./testdata/course/course.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock := testCase.dbMock
			if dbMock == nil {
				dbMock, _ = testhelpers.ConnectDB()
			}

			utils.LoadFixture(dbMock, testCase.giveFixture)

			courseRepo := Course{
				Db: dbMock,
			}

			err := courseRepo.DeleteCourse(testCase.input)

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

func Test_UpdateCourse(t *testing.T) {
	testCases := []struct {
		name          string
		inputID       string
		inputCourse  *CourseEntity
		expectedError error
		giveFixture   string
	}{
		{
			name:    "update course fail",
			inputID: "1",
			inputCourse: &CourseEntity{
				Name:      "Math",
				StartTime: "4/23/2020",
				TeacherID: 1,
			},
			expectedError: errors.New("pq: invalid input syntax for type timestamp: \"\""),
			giveFixture:   "./testdata/course/course.sql",
		},
		{
			name:    "update course successfully",
			inputID: "1",
			inputCourse: &CourseEntity{
				Name:      "Math",
				StartTime: "4/23/2020",
				EndTime:   "4/24/2020",
				TeacherID: 1,
			},
			expectedError: nil,
			giveFixture:   "./testdata/course/course.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			courseRepo := Course{
				Db: dbMock,
			}

			err := courseRepo.UpdateCourse(testCase.inputID, testCase.inputCourse)

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
