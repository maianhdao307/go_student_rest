package services

import (
	"github.com/lucasjones/reggen"
	"student_rest/models"
	"student_rest/repositories"
)

type Student struct{
	repositories.StudentRepositories
}

type StudentServices interface {
	RegisterCourse(registerCourseModel *models.RegisterCourseModel) (*models.RegisterCourseModel, error)
}

var (
	studentRandomIDRegex = "[A-Z0-9]{6}"
)

func generateID(regex string, limit int) (string, error) {
	result, err := reggen.Generate("[A-Z0-9]{6}", 6)
	return result, err

}

//func (_self StudentServices) CreateStudent(student *models.StudentModel) (*repositories.StudentEntity, error) {
//	studentID, err := generateID(studentRandomIDRegex, 6)
//	if err != nil {
//		return nil, err
//	}
//	student.StudentID = studentID
//
//	convertedStudent := transformStudenModelToStudentEntity(student)
//	result, err2 := studentRepositories.CreateStudent(convertedStudent)
//	if err2 != nil {
//		return nil, err2
//	}
//	return result, nil
//}
//
//func transformStudenModelToStudentEntity(model *models.StudentModel) *repositories.StudentEntity {
//	return &repositories.StudentEntity{
//		StudentID:   model.StudentID,
//		FirstName:   model.FirstName,
//		LastName:    model.LastName,
//		DateOfBirth: model.DateOfBirth,
//	}
//}
//
//func (_self StudentServices) GetStudents() ([]*repositories.StudentEntity, error) {
//	result, err := studentRepositories.GetStudents()
//	return result, err
//}
//
//func (_self StudentServices) GetStudentByID(id string) (*repositories.StudentEntity, error) {
//	result, err := studentRepositories.GetStudentByID(id)
//	return result, err
//}
//
//func (_self StudentServices) DeleteStudent(id string) error {
//	err := studentRepositories.DeleteStudent(id)
//	return err
//}
//
//func (_self StudentServices) UpdateStudent(id string, student *models.StudentModel) error {
//	convertedStudent := transformStudenModelToStudentEntity(student)
//	err := studentRepositories.UpdateStudent(id, convertedStudent)
//	return err
//}

func (_self Student) RegisterCourse(registerCourseModel *models.RegisterCourseModel) (*models.RegisterCourseModel, error) {
	studentID, err := generateID(studentRandomIDRegex, 6)
	if err != nil {
		return nil, err
	}

	registerCourseModel.Student.StudentID = studentID

	result, err := _self.StudentRepositories.RegisterCourse(registerCourseModel)
	if err != nil {
		return nil, err
	}

	return result, nil
}
