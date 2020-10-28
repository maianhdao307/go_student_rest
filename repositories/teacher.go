package repositories

import (
	"database/sql"
)


type Teacher struct{
	Db *sql.DB
}

type TeacherRepositories interface {
	CreateTeacher(teacher *TeacherEntity) (*TeacherEntity, error)
	GetTeacherByID(id string) (*TeacherEntity, error)
	DeleteTeacher(id string) error
	UpdateTeacher(id string, teacher *TeacherEntity) error
}

func (_self Teacher) CreateTeacher(teacher *TeacherEntity) (*TeacherEntity, error) {
	sqlStmt := `INSERT INTO teachers("first_name", "last_name", "date_of_birth") VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err := _self.Db.QueryRow(sqlStmt, teacher.FirstName, teacher.LastName, teacher.DateOfBirth).Scan(&id)
	if err != nil {
		return nil, err
	}
	teacher.ID = id
	return teacher, nil
}

func (_self Teacher) GetTeacherByID(id string) (*TeacherEntity, error) {
	sqlStmt := `SELECT * FROM teachers WHERE id=$1`
	var teacher TeacherEntity
	err := _self.Db.QueryRow(sqlStmt, id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.DateOfBirth)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (_self Teacher) DeleteTeacher(id string) error {
	sqlStmt := `DELETE FROM teachers WHERE id=$1`
	_, err := _self.Db.Exec(sqlStmt, id)
	return err
}

func (_self Teacher) UpdateTeacher(id string, teacher *TeacherEntity) error {
	sqlStmt := `UPDATE teachers SET "first_name" = $2, "last_name" = $3, "date_of_birth" = $4 WHERE id=$1`
	_, err := _self.Db.Exec(sqlStmt, id, teacher.FirstName, teacher.LastName, teacher.DateOfBirth)
	return err
}
