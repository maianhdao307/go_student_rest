package services

import (
	"student_rest/models"
	"student_rest/repositories"
)

type TeacherServices struct {}

var teacherRepositories = repositories.TeacherRepositories {}

func (_self TeacherServices) CreateTeacher(teacher *models.TeacherModel) (*repositories.TeacherEntity, error) {
	convertedTeacher := transformTeacherModelToTeacherEntity(teacher)
	result, err := teacherRepositories.CreateTeacher(convertedTeacher)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func transformTeacherModelToTeacherEntity(model *models.TeacherModel) *repositories.TeacherEntity {
	return &repositories.TeacherEntity{
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		DateOfBirth: model.DateOfBirth,
	}
}

func (_self TeacherServices) GetTeachers() ([]*repositories.TeacherEntity, error) {
	result, err := teacherRepositories.GetTeachers()
	return result, err
}

func (_self TeacherServices) GetTeacherByID(id string) (*repositories.TeacherEntity, error) {
	result, err := teacherRepositories.GetTeacherByID(id)
	return result, err
}

func (_self TeacherServices) DeleteTeacher(id string) error {
	err := teacherRepositories.DeleteTeacher(id)
	return err
}

func (_self TeacherServices) UpdateTeacher(id string, teacher *models.TeacherModel) error {
	convertedTeacher := transformTeacherModelToTeacherEntity(teacher)
	err := teacherRepositories.UpdateTeacher(id, convertedTeacher)
	return err
}
