package services

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_GenerateID(t *testing.T) {
	testCases := []struct {
		name          string
		inputRegex    string
		inputLimit    int
		expectedError error
	}{
		{
			name:          "generate id fail",
			inputRegex:    "(",
			inputLimit:    6,
			expectedError: errors.New("error parsing regexp: missing closing ): `(`"),
		},
		{
			name:          "generate id successfully",
			inputRegex:    "[A-Z0-9]{6}",
			inputLimit:    6,
			expectedError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			utils := Utils{}

			result, err := utils.GenerateID(testCase.inputRegex, testCase.inputLimit)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.LessOrEqual(t, len(result), testCase.inputLimit)
			}
		})
	}
}
