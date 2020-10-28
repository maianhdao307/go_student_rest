package repositories

import (
	"errors"
	"github.com/stretchr/testify/require"
	"student_rest/models"
	"student_rest/testhelpers"
	"student_rest/utils"
	"testing"
	"database/sql"
)

func Test_CreateStudent(t *testing.T) {
	testCases := []struct {
		name          string
		input         *StudentEntity
		expectedValue *StudentEntity
		expectedError error
		giveFixture   string
	}{
		{
			name: "insert student fail",
			input: &StudentEntity{
				StudentID:   "12345678",
				FirstName:   "Dao",
				LastName:    "Mai",
				DateOfBirth: "6/25/1997",
			},
			expectedValue: nil,
			expectedError: errors.New("pq: value too long for type character varying(6)"),
			giveFixture:   "./testdata/truncate_data.sql",
		},
		{
			name: "insert student successfully",
			input: &StudentEntity{
				StudentID:   "123456",
				FirstName:   "Dao",
				LastName:    "Mai",
				DateOfBirth: "6/25/1997",
			},
			expectedValue: &StudentEntity{
				StudentID:   "123456",
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

			studentRepo := Student{
				Db: dbMock,
			}

			result, err := studentRepo.CreateStudent(testCase.input)

			if testCase.expectedError != nil {
				// For Fail Logic
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				// For Success Logic
				require.NoError(t, err)

				sqlStmt := `SELECT * FROM students WHERE id=$1`
				var student StudentEntity
				err := dbMock.QueryRow(sqlStmt, result.ID).Scan(&student.ID, &student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth)
				if err != nil {
					t.Error(err)
				}
				require.Equal(t, testCase.expectedValue.StudentID, student.StudentID)
				require.Equal(t, testCase.expectedValue.FirstName, student.FirstName)
				require.Equal(t, testCase.expectedValue.LastName, student.LastName)
				require.Equal(t, testCase.expectedValue.DateOfBirth, student.DateOfBirth)
			}
		})
	}
}

func Test_GetStudentByID(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedValue *StudentEntity
		expectedError error
		giveFixture   string
	}{
		{
			name:          "get student by id fail",
			input:         "3",
			expectedValue: nil,
			expectedError: errors.New("sql: no rows in result set"),
			giveFixture:   "./testdata/student/student.sql",
		},
		{
			name:  "get student by id successfully",
			input: "2",
			expectedValue: &StudentEntity{
				ID:          2,
				StudentID:   "234567",
				FirstName:   "Mai",
				LastName:    "Dao",
				DateOfBirth: "1998-11-02T00:00:00Z",
			},
			expectedError: nil,
			giveFixture:   "./testdata/student/student.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			studentRepo := Student{
				Db: dbMock,
			}

			result, err := studentRepo.GetStudentByID(testCase.input)

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

func Test_DeleteStudent(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
		giveFixture   string
	}{
		{
			name:          "delete student fail",
			input:         "1",
			expectedError: errors.New("pq: update or delete on table \"students\" violates foreign key constraint \"students_courses_student_id_fkey\" on table \"students_courses\""),
			giveFixture:   "./testdata/student/student.sql",
		},
		{
			name:          "delete student successfully",
			input:         "2",
			expectedError: nil,
			giveFixture:   "./testdata/student/student.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			studentRepo := Student{
				Db: dbMock,
			}

			err := studentRepo.DeleteStudent(testCase.input)

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

func Test_UpdateStudent(t *testing.T) {
	testCases := []struct {
		name          string
		inputID         string
		inputStudent *StudentEntity
		expectedError error
		giveFixture   string
	}{
		{
			name:  "update student fail",
			inputID: "1",
			inputStudent: &StudentEntity{
				StudentID:   "123456",
				FirstName:   "Dao",
				LastName:    "Mai",
			},
			expectedError: errors.New("pq: invalid input syntax for type timestamp: \"\""),
			giveFixture:   "./testdata/student/student.sql",
		},
		{
			name:  "update student successfully",
			inputID: "1",
			inputStudent: &StudentEntity{
				StudentID:   "123456",
				FirstName:   "Dao",
				LastName:    "Mai",
				DateOfBirth: "1/12/1997",
			},
			expectedError: nil,
			giveFixture:   "./testdata/student/student.sql",
		},

	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := testhelpers.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			studentRepo := Student{
				Db: dbMock,
			}

			err := studentRepo.UpdateStudent(testCase.inputID, testCase.inputStudent)

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

func Test_RegisterCourse(t *testing.T) {
	dbMockFail, _ := testhelpers.ConnectDBFailed()

	testCases := []struct {
		name          string
		input         *models.RegisterCourseModel
		expectedValue *models.RegisterCourseModel
		expectedError error
		giveFixture   string
		dbMock        *sql.DB
	}{
		{
			name:          "can't create transaction",
			input:         &models.RegisterCourseModel{},
			expectedValue: nil,
			expectedError: errors.New("pq: password authentication failed for user \"postgres\""),
			giveFixture:   "./testdata/truncate_data.sql",
			dbMock:        dbMockFail,
		},
		{
			name: "insert student fail",
			input: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					StudentID:   "12345678",
					DateOfBirth: "6/25/1997",
				},
			},
			expectedValue: nil,
			expectedError: errors.New("pq: value too long for type character varying(6)"),
			giveFixture:   "./testdata/truncate_data.sql",
		},
		{
			name: "insert teacher fail",
			input: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					StudentID:   "123456",
					FirstName:   "Dao",
					LastName:    "Mai",
					DateOfBirth: "6/25/1997",
				},
				Course: &models.CourseModel{
					Teacher: &models.TeacherModel{},
				},
			},
			expectedValue: nil,
			expectedError: errors.New("pq: invalid input syntax for type timestamp: \"\""),
			giveFixture:   "./testdata/truncate_data.sql",
		},
		{
			name: "insert course fail",
			input: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					StudentID:   "123456",
					FirstName:   "Dao",
					LastName:    "Mai",
					DateOfBirth: "6/25/1997",
				},
				Course: &models.CourseModel{
					Teacher: &models.TeacherModel{
						FirstName:   "Dao",
						LastName:    "Mai",
						DateOfBirth: "6/25/1997",
					},
				},
			},
			expectedValue: nil,
			expectedError: errors.New("pq: invalid input syntax for type timestamp: \"\""),
			giveFixture:   "./testdata/truncate_data.sql",
		},
		{
			name: "register course successfully",
			input: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					StudentID:   "123456",
					FirstName:   "Dao",
					LastName:    "Mai",
					DateOfBirth: "6/25/1997",
				},
				Course: &models.CourseModel{
					Name:      "Math",
					StartTime: "4/23/2020",
					EndTime:   "4/24/2020",
					Teacher: &models.TeacherModel{
						FirstName:   "Dao",
						LastName:    "Mai",
						DateOfBirth: "6/25/1997",
					},
				},
			},
			expectedValue: &models.RegisterCourseModel{
				Student: &models.StudentModel{
					StudentID:   "123456",
					FirstName:   "Dao",
					LastName:    "Mai",
					DateOfBirth: "1997-06-25T00:00:00Z",
				},
				Course: &models.CourseModel{
					Name:      "Math",
					StartTime: "2020-04-23T00:00:00Z",
					EndTime:   "2020-04-24T00:00:00Z",
					Teacher: &models.TeacherModel{
						FirstName:   "Dao",
						LastName:    "Mai",
						DateOfBirth: "1997-06-25T00:00:00Z",
					},
				},
			},
			expectedError: nil,
			giveFixture:   "./testdata/truncate_data.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock := testCase.dbMock
			if dbMock == nil {
				dbMock, _ = testhelpers.ConnectDB()
			}

			utils.LoadFixture(dbMock, testCase.giveFixture)

			studentRepo := Student{
				Db: dbMock,
			}

			result, err := studentRepo.RegisterCourse(testCase.input)

			if testCase.expectedError != nil {
				// For Fail Logic
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				// For Success Logic
				require.NoError(t, err)

				sqlStmt := `SELECT * FROM students WHERE id=$1`
				var student StudentEntity
				err := dbMock.QueryRow(sqlStmt, result.Student.ID).Scan(&student.ID, &student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth)
				if err != nil {
					t.Error(err)
				}
				require.Equal(t, testCase.expectedValue.Student.StudentID, student.StudentID)
				require.Equal(t, testCase.expectedValue.Student.FirstName, student.FirstName)
				require.Equal(t, testCase.expectedValue.Student.LastName, student.LastName)
				require.Equal(t, testCase.expectedValue.Student.DateOfBirth, student.DateOfBirth)

				sqlStmt = `SELECT * FROM teachers WHERE id=$1`
				var teacher TeacherEntity
				err = dbMock.QueryRow(sqlStmt, result.Course.Teacher.ID).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.DateOfBirth)
				if err != nil {
					t.Error(err)
				}
				require.Equal(t, testCase.expectedValue.Course.Teacher.FirstName, teacher.FirstName)
				require.Equal(t, testCase.expectedValue.Course.Teacher.LastName, teacher.LastName)
				require.Equal(t, testCase.expectedValue.Course.Teacher.DateOfBirth, teacher.DateOfBirth)

				sqlStmt = `SELECT * FROM courses WHERE id=$1`
				var course CourseEntity
				err = dbMock.QueryRow(sqlStmt, result.Course.ID).Scan(&course.ID, &course.Name, &course.StartTime, &course.EndTime, &course.TeacherID)
				if err != nil {
					t.Error(err)
				}
				require.Equal(t, testCase.expectedValue.Course.Name, course.Name)
				require.Equal(t, testCase.expectedValue.Course.StartTime, course.StartTime)
				require.Equal(t, testCase.expectedValue.Course.EndTime, course.EndTime)
				require.Equal(t, teacher.ID, course.TeacherID)
			}
		})
	}
}
