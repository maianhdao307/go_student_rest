package repositories

import (
	"github.com/stretchr/testify/require"
	"student_rest/db"
	"student_rest/models"
	"student_rest/utils"
	"testing"
)



func Test_RegisterCourse(t *testing.T) {

	testCases := []struct {
		name          string
		input         *models.RegisterCourseModel
		expectedValue *models.RegisterCourseModel
		expectedError error
		giveFixture   string
	}{
		{
			name:          "can't create transaction",
			input:         &models.RegisterCourseModel{},
			expectedValue: &models.RegisterCourseModel{},
			expectedError: nil,
			giveFixture: "./testdata/...sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock, _ := db.ConnectDB()

			utils.LoadFixture(dbMock, testCase.giveFixture)

			studenRepo:= Student{
				Db: testCase.dbMock,
			}

			result, err := studenRepo.RegisterCourse(testCase.input)

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
