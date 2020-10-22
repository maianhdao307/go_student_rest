package handlers

import (
	"errors"
	"student_rest/models"
	"student_rest/repositories"
)

type StudentRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
}

func (_self StudentRequest) validation() error {
	if _self.FirstName == "" {
		return errors.New("first name is required")
	}
	if _self.LastName == "" {
		return errors.New("last name is required")
	}
	return nil
}

type StudentsResponse struct {
	Success  bool                         `json:"success"`
	Students []repositories.StudentEntity `json:"students"`
}

type StudentResponse struct {
	Success bool                       `json:"success"`
	Student repositories.StudentEntity `json:"student"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type TeacherRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
}

func (_self TeacherRequest) validation() error {
	if _self.FirstName == "" {
		return errors.New("first name is required")
	}
	if _self.LastName == "" {
		return errors.New("last name is required")
	}
	return nil
}

type TeachersResponse struct {
	Success  bool                         `json:"success"`
	Teachers []repositories.TeacherEntity `json:"teachers"`
}

type TeacherResponse struct {
	Success bool                       `json:"success"`
	Teacher repositories.TeacherEntity `json:"teacher"`
}

type CourseRequest struct {
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TeacherID int    `json:"teacherID"`
}

func (_self CourseRequest) validation() error {
	if _self.Name == "" {
		return errors.New("course name is required")
	}
	return nil
}

type CourseResponse struct {
	Success bool               `json:"success"`
	Course  models.CourseModel `json:"course"`
}

type RegisterCourseRequest struct {
	Student StudentRequest
	Course  CourseRequest
}

func (_self RegisterCourseRequest) validation() error {
	if err := _self.Student.validation(); err != nil  {
		return err
	}
	if err := _self.Course.validation(); err != nil  {
		return err
	}
	return nil
}

type RegisterCourseResponse struct {
	Success bool               `json:"success"`
	Course  models.CourseModel `json:"course"`
	Student models.StudentModel `json:"student"`
}

