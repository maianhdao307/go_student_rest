package services

import (
	"student_rest/models"
	"student_rest/repositories"
)

type Course struct {
	repositories.CourseRepositories
}

type CourseServices interface {
	CreateCourse(course *models.CourseModel) (*models.CourseModel, error)
	GetCourseByID(id string) (*models.CourseModel, error)
	DeleteCourse(id string) error
	UpdateCourse(id string, course *models.CourseModel) error
}

func (_self Course) CreateCourse(course *models.CourseModel) (*models.CourseModel, error) {
	convertedCourse := transformCourseModelToCourseEntity(*course)
	result, err := _self.CourseRepositories.CreateCourse(&convertedCourse)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func transformCourseModelToCourseEntity(model models.CourseModel) repositories.CourseEntity {
	return repositories.CourseEntity{
		Name:      model.Name,
		StartTime: model.StartTime,
		EndTime:   model.EndTime,
		TeacherID: model.Teacher.ID,
	}
}

func (_self Course) GetCourseByID(id string) (*models.CourseModel, error) {
	result, err := _self.CourseRepositories.GetCourseByID(id)
	return result, err
}

func (_self Course) DeleteCourse(id string) error {
	err := _self.CourseRepositories.DeleteCourse(id)
	return err
}

func (_self Course) UpdateCourse(id string, course *models.CourseModel) error {
	convertedCourse := transformCourseModelToCourseEntity(*course)
	err := _self.CourseRepositories.UpdateCourse(id, &convertedCourse)
	return err
}
