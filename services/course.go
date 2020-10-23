package services

import (
	"student_rest/models"
	"student_rest/repositories"
)

type CourseServices struct{}

var courseRepositories = repositories.CourseRepositories{}

func (_self CourseServices) CreateCourse(course *models.CourseModel) (*models.CourseModel, error) {
	convertedCourse := transformCourseModelToCourseEntity(*course)
	result, err := courseRepositories.CreateCourse(&convertedCourse)
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

func (_self CourseServices) GetCourseById(id string) (*models.CourseModel, error) {
	result, err := courseRepositories.GetCourseById(id)
	return result, err
}

func (_self CourseServices) DeleteCourse(id string) error {
	err := courseRepositories.DeleteCourse(id)
	return err
}

func (_self CourseServices) UpdateCourse(id string, course *models.CourseModel) error {
	convertedCourse := transformCourseModelToCourseEntity(*course)
	err := courseRepositories.UpdateCourse(id, &convertedCourse)
	return err
}
