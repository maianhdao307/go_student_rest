package services

import (
	"student_rest/models"
	"student_rest/repositories"
)

type Student struct{
	StudentRepositories repositories.StudentRepositories
	Utils UtilsService
}

type StudentServices interface {
	CreateStudent(student *models.StudentModel) (*repositories.StudentEntity, error)
	GetStudentByID(id string) (*repositories.StudentEntity, error)
	DeleteStudent(id string) error
	UpdateStudent(id string, student *models.StudentModel) error
	RegisterCourse(registerCourseModel *models.RegisterCourseModel) (*models.RegisterCourseModel, error)
}

var (
	studentRandomIDRegex = "[A-Z0-9]{6}"
)

func (_self Student) CreateStudent(student *models.StudentModel) (*repositories.StudentEntity, error) {
	studentID, err := _self.Utils.GenerateID(studentRandomIDRegex, 6)
	if err != nil {
		return nil, err
	}
	student.StudentID = studentID

	convertedStudent := transformStudentModelToStudentEntity(student)
	result, err2 := _self.StudentRepositories.CreateStudent(convertedStudent)
	if err2 != nil {
		return nil, err2
	}
	return result, nil
}

func transformStudentModelToStudentEntity(model *models.StudentModel) *repositories.StudentEntity {
	return &repositories.StudentEntity{
		StudentID:   model.StudentID,
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		DateOfBirth: model.DateOfBirth,
	}
}

func (_self Student) GetStudentByID(id string) (*repositories.StudentEntity, error) {
	result, err := _self.StudentRepositories.GetStudentByID(id)
	return result, err
}

func (_self Student) DeleteStudent(id string) error {
	err := _self.StudentRepositories.DeleteStudent(id)
	return err
}

func (_self Student) UpdateStudent(id string, student *models.StudentModel) error {
	convertedStudent := transformStudentModelToStudentEntity(student)
	err := _self.StudentRepositories.UpdateStudent(id, convertedStudent)
	return err
}

func (_self Student) RegisterCourse(registerCourseModel *models.RegisterCourseModel) (*models.RegisterCourseModel, error) {
	studentID, err := _self.Utils.GenerateID(studentRandomIDRegex, 6)
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
