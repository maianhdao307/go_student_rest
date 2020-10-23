package repositories

import (
	"context"
	"database/sql"
	"student_rest/models"
)

type Student struct{
	Db *sql.DB
}

type StudentRepositories interface {
	RegisterCourse(registerCourseModel *models.RegisterCourseModel) (*models.RegisterCourseModel, error)
}

//
//func (_self Student) CreateStudent(student *StudentEntity) (*StudentEntity, error) {
//	sqlStmt := `INSERT INTO students("student_id", "first_name", "last_name", "date_of_birth") VALUES ($1, $2, $3, $4) RETURNING id`
//	id := 0
//	err := db.DB.QueryRow(sqlStmt, student.StudentID, student.FirstName, student.LastName, student.DateOfBirth).Scan(&id)
//	if err != nil {
//		return nil, err
//	}
//	student.ID = id
//	return student, nil
//}
//
//func (_self Student) GetStudents() ([]*StudentEntity, error) {
//	sqlStmt := `SELECT * FROM students`
//	var students = make([]*StudentEntity, 0)
//	rows, err := db.DB.Query(sqlStmt)
//	if err != nil {
//		return nil, err
//	}
//	for rows.Next() {
//		var student StudentEntity
//		rows.Scan(&student.ID, &student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth)
//		students = append(students, &student)
//	}
//	return students, nil
//}
//
//func (_self Student) GetStudentByID(id string) (*StudentEntity, error) {
//	sqlStmt := `SELECT * FROM students WHERE id=$1`
//	var student StudentEntity
//	err := db.DB.QueryRow(sqlStmt, id).Scan(&student.ID, &student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth)
//	if err != nil {
//		return nil, err
//	}
//	return &student, nil
//}
//
//func (_self Student) DeleteStudent(id string) error {
//	sqlStmt := `DELETE FROM students WHERE id=$1`
//	_, err := db.DB.Exec(sqlStmt, id)
//	return err
//}
//
//func (_self Student) UpdateStudent(id string, student *StudentEntity) error {
//	sqlStmt := `UPDATE students SET "first_name" = $2, "last_name" = $3, "date_of_birth" = $4 WHERE id=$1`
//	_, err := db.DB.Exec(sqlStmt, id, student.FirstName, student.LastName, student.DateOfBirth)
//	return err
//}

func (_self Student) RegisterCourse(registerCourseModel *models.RegisterCourseModel) (*models.RegisterCourseModel, error) {
	student := registerCourseModel.Student
	course := registerCourseModel.Course

	ctx := context.Background()
	tx, err := _self.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	sqlStmt := `INSERT INTO students(student_id, first_name, last_name, date_of_birth) VALUES ($1, $2, $3, $4) RETURNING id`

	newStudent := models.StudentModel{
		StudentID: student.StudentID,
		FirstName: student.FirstName,
		LastName: student.LastName,
		DateOfBirth: student.DateOfBirth,
	}
	err = tx.QueryRowContext(ctx, sqlStmt, student.StudentID, student.FirstName, student.LastName, student.DateOfBirth).Scan(&newStudent.ID)

	if err != nil {
		return nil, err
	}


	sqlStmt = `INSERT INTO teachers(first_name, last_name, date_of_birth) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRowContext(ctx, sqlStmt, course.Teacher.FirstName, course.Teacher.LastName, course.Teacher.DateOfBirth).
		Scan(&course.Teacher.ID)

	if err != nil {
		return nil, err
	}

	sqlStmt = `INSERT INTO courses(name, start_time, end_time, teacher_id) VALUES ($1, $2, $3, $4) RETURNING id`

	err = tx.QueryRowContext(ctx, sqlStmt, course.Name, course.StartTime, course.EndTime, course.Teacher.ID).
		Scan(&course.ID)

	if err != nil {
		return nil, err
	}

	sqlStmt = `INSERT INTO students_courses(student_id, course_id) VALUES ($1, $2)`

	_, err = tx.ExecContext(ctx, sqlStmt, newStudent.ID, course.ID)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &models.RegisterCourseModel{
		Student: &newStudent,
		Course:  course,
	}, nil
}
