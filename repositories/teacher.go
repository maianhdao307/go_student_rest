package repositories

import (
	"student_rest/db"
)


type TeacherRepositories struct{}

func (_self TeacherRepositories) CreateTeacher(teacher *TeacherEntity) (*TeacherEntity, error) {
	sqlStmt := `INSERT INTO teachers("first_name", "last_name", "date_of_birth") VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err := db.DB.QueryRow(sqlStmt, teacher.FirstName, teacher.LastName, teacher.DateOfBirth).Scan(&id)
	if err != nil {
		return nil, err
	}
	teacher.ID = id
	return teacher, nil
}

func (_self TeacherRepositories) GetTeachers() ([]*TeacherEntity, error) {
	sqlStmt := `SELECT * FROM teachers`
	var teachers = make([]*TeacherEntity, 0)
	rows, err := db.DB.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var teacher TeacherEntity
		rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.DateOfBirth)
		teachers = append(teachers, &teacher)
	}
	return teachers, nil
}

func (_self TeacherRepositories) GetTeacherByID(id string) (*TeacherEntity, error) {
	sqlStmt := `SELECT * FROM teachers WHERE id=$1`
	var teacher TeacherEntity
	err := db.DB.QueryRow(sqlStmt, id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.DateOfBirth)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (_self TeacherRepositories) DeleteTeacher(id string) error {
	sqlStmt := `DELETE FROM teachers WHERE id=$1`
	_, err := db.DB.Exec(sqlStmt, id)
	return err
}

func (_self TeacherRepositories) UpdateTeacher(id string, teacher *TeacherEntity) error {
	sqlStmt := `UPDATE teachers SET "first_name" = $2, "last_name" = $3, "date_of_birth" = $4 WHERE id=$1`
	_, err := db.DB.Exec(sqlStmt, id, teacher.FirstName, teacher.LastName, teacher.DateOfBirth)
	return err
}
