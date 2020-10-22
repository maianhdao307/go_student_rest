package repositories

import (
	"student_rest/db"
)

type StudentRepositories struct{}

func (_self StudentRepositories) CreateStudent(student StudentEntity) (StudentEntity, error) {
	sqlStmt := `INSERT INTO students("student_id", "first_name", "last_name", "date_of_birth") VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0
	err := db.DB.QueryRow(sqlStmt, student.StudentID, student.FirstName, student.LastName, student.DateOfBirth).Scan(&id)
	if err != nil {
		return StudentEntity{}, err
	}
	student.ID = id
	return student, nil
}

func (_self StudentRepositories) GetStudents() ([]StudentEntity, error) {
	sqlStmt := `SELECT * FROM students`
	var students = make([]StudentEntity, 0)
	rows, err := db.DB.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student StudentEntity
		rows.Scan(&student.ID, &student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth)
		students = append(students, student)
	}
	return students, nil
}

func (_self StudentRepositories) GetStudentByID(id string) (StudentEntity, error) {
	sqlStmt := `SELECT * FROM students WHERE id=$1`
	var student StudentEntity
	err := db.DB.QueryRow(sqlStmt, id).Scan(&student.ID, &student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth)
	if err != nil {
		return StudentEntity{}, err
	}
	return student, nil
}

func (_self StudentRepositories) DeleteStudent(id string) error {
	sqlStmt := `DELETE FROM students WHERE id=$1`
	_, err := db.DB.Exec(sqlStmt, id)
	return err
}

func (_self StudentRepositories) UpdateStudent(id string, student StudentEntity) error {
	sqlStmt := `UPDATE students SET "first_name" = $2, "last_name" = $3, "date_of_birth" = $4 WHERE id=$1`
	_, err := db.DB.Exec(sqlStmt, id, student.FirstName, student.LastName, student.DateOfBirth)
	return err
}
