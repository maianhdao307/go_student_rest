package utils

import (
	"io/ioutil"
	"strings"
	"student_rest/db"
)

func LoadFixture(path string) error {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	DB, _ := db.ConnectDB()

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := DB.Exec(request)
		if err != nil {
			return err
		}
	}
	return nil
}
