package services

import (
	"github.com/lucasjones/reggen"
	"student_rest/models"
	"student_rest/repositories"
)

type StudentServices struct {}

var studentRepositories = repositories.StudentRepositories {}

func (_self StudentServices) CreateStudent(student models.StudentModel) (repositories.StudentEntity, error) {
	studentID, _ := reggen.Generate("[A-Z0-9]{6}", 6)
	student.StudentID = studentID

	convertedStudent := transformStudenModelToStudentEntity(student)
	result, err := studentRepositories.CreateStudent(convertedStudent)
	if err != nil {
		return repositories.StudentEntity{}, err
	}
	return result, nil
}

func transformStudenModelToStudentEntity(model models.StudentModel) repositories.StudentEntity {
	return repositories.StudentEntity{
		StudentID:   model.StudentID,
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		DateOfBirth: model.DateOfBirth,
	}
}

func (_self StudentServices) GetStudents() ([]repositories.StudentEntity, error) {
	result, err := studentRepositories.GetStudents()
	return result, err
}

func (_self StudentServices)  GetStudentByID(id string) (repositories.StudentEntity, error) {
	result, err := studentRepositories.GetStudentByID(id)
	return result, err
}

func (_self StudentServices)  DeleteStudent(id string) error {
	err := studentRepositories.DeleteStudent(id)
	return err
}

func (_self StudentServices) UpdateStudent(id string, student models.StudentModel) error {
	convertedStudent := transformStudenModelToStudentEntity(student)
	err := studentRepositories.UpdateStudent(id, convertedStudent)
	return err
}
