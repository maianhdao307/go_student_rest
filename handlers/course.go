package handlers

import (
	"encoding/json"
	"net/http"
	"student_rest/models"

	"github.com/go-chi/chi"
	"student_rest/services"
)

type CourseHandlers struct{}

var courseServices = services.CourseServices{}

// CreateStudent creates new course
func (_self CourseHandlers) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course CourseRequest

	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := course.validation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertedCourse := TransformCourseRequestToCourseModel(course)
	result, err := courseServices.CreateCourse(&convertedCourse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(CourseResponse{
		Success: true,
		Course:  result,
	})
	return
}

func TransformCourseRequestToCourseModel(request CourseRequest) models.CourseModel {
	return models.CourseModel{
		Name:      request.Name,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		Teacher:   &models.TeacherModel{
			ID: request.TeacherID,
		},
	}
}

func (_self CourseHandlers) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result, err := courseServices.GetCourseById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(CourseResponse{
		Success: true,
		Course:  result,
	})
}

func (_self CourseHandlers) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := courseServices.DeleteCourse(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(SuccessResponse{
		Success: true,
	})
}

func (_self CourseHandlers) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var course CourseRequest

	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := course.validation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertedCourse := TransformCourseRequestToCourseModel(course)
	err := courseServices.UpdateCourse(id, &convertedCourse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(SuccessResponse{
		Success: true,
	})
}
