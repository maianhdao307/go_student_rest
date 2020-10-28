package services

import (
	"student_rest/models"
	"student_rest/repositories"
)

type Teacher struct {
	repositories.TeacherRepositories
}

type TeacherServices interface {
	CreateTeacher(teacher *models.TeacherModel) (*repositories.TeacherEntity, error)
	GetTeacherByID(id string) (*repositories.TeacherEntity, error)
	DeleteTeacher(id string) error
	UpdateTeacher(id string, teacher *models.TeacherModel) error
}

func (_self Teacher) CreateTeacher(teacher *models.TeacherModel) (*repositories.TeacherEntity, error) {
	convertedTeacher := transformTeacherModelToTeacherEntity(teacher)
	result, err := _self.TeacherRepositories.CreateTeacher(convertedTeacher)
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

func (_self Teacher) GetTeacherByID(id string) (*repositories.TeacherEntity, error) {
	result, err := _self.TeacherRepositories.GetTeacherByID(id)
	return result, err
}

func (_self Teacher) DeleteTeacher(id string) error {
	err := _self.TeacherRepositories.DeleteTeacher(id)
	return err
}

func (_self Teacher) UpdateTeacher(id string, teacher *models.TeacherModel) error {
	convertedTeacher := transformTeacherModelToTeacherEntity(teacher)
	err := _self.TeacherRepositories.UpdateTeacher(id, convertedTeacher)
	return err
}
