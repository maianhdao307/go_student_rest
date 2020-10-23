package repositories

import (
	"student_rest/db"
	"student_rest/models"
)

type CourseRepositories struct{}

func (_self CourseRepositories) CreateCourse(course *CourseEntity) (*models.CourseModel, error) {
	sqlStmt := `INSERT INTO courses("name", "start_time", "end_time", "teacher_id") VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0
	err := db.DB.QueryRow(sqlStmt, course.Name, course.StartTime, course.EndTime, course.TeacherID).Scan(&id)

	if err != nil {
		return nil, err
	}
	course.ID = id

	sqlStmt = `SELECT * FROM teachers WHERE id=$1`
	var teacher models.TeacherModel
	err = db.DB.QueryRow(sqlStmt, course.TeacherID).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.DateOfBirth)

	if err != nil {
		return nil, err
	}

	newCourse := &models.CourseModel{
		ID:        course.ID,
		Name:      course.Name,
		StartTime: course.StartTime,
		EndTime:   course.EndTime,
		Teacher:   &teacher,
	}
	return newCourse, nil
}

func (_self CourseRepositories) GetCourseById(id string) (*models.CourseModel, error) {
	sqlStmt := `SELECT * FROM courses WHERE id = $1`
	var course CourseEntity
	err := db.DB.QueryRow(sqlStmt, id).Scan(&course.ID, &course.Name, &course.StartTime, &course.EndTime, &course.TeacherID)
	if err != nil {
		return nil, err
	}

	sqlStmt = `SELECT * FROM teachers WHERE id=$1`
	var teacher models.TeacherModel
	err = db.DB.QueryRow(sqlStmt, course.TeacherID).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.DateOfBirth)
	if err != nil {
		return nil, err
	}
	return &models.CourseModel{
		ID:        course.ID,
		Name:      course.Name,
		StartTime: course.StartTime,
		EndTime:   course.EndTime,
		Teacher:   &teacher,
	}, nil
}

func (_self CourseRepositories) DeleteCourse(id string) error {
	sqlStmt := `DELETE FROM courses WHERE id=$1`
	_, err := db.DB.Exec(sqlStmt, id)
	return err
}

func (_self CourseRepositories) UpdateCourse(id string, course *CourseEntity) error {
	sqlStmt := `UPDATE courses SET "name" = $2, "start_time" = $3, "end_time" = $4, "teacher_id" = $5 WHERE id=$1`
	_, err := db.DB.Exec(sqlStmt, id, course.Name, course.StartTime, course.EndTime, course.TeacherID)
	return err
}
