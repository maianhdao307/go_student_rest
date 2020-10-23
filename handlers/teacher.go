package handlers

import (
	"encoding/json"
	"net/http"
	"student_rest/models"

	"github.com/go-chi/chi"
	"student_rest/services"
)

type TeacherHandlers struct {}

var teacherServices = services.TeacherServices {}

// CreateStudent creates new teacher
func (_self TeacherHandlers) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher TeacherRequest

	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := teacher.validation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertedTeacher := transformTeacherRequestToTeacherModel(teacher)
	result, err := teacherServices.CreateTeacher(&convertedTeacher)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(TeacherResponse{
		Success: true,
		Teacher: result,
	})
	return
}

func transformTeacherRequestToTeacherModel(request TeacherRequest) models.TeacherModel {
	return models.TeacherModel{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBirth,
	}
}

func (_self TeacherHandlers) GetTeachers(w http.ResponseWriter, r *http.Request) {
	result, err := teacherServices.GetTeachers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(TeachersResponse{
		Success:  true,
		Teachers: result,
	})
}

func (_self TeacherHandlers) GetTeacherByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result, err := teacherServices.GetTeacherByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(TeacherResponse{
		Success: true,
		Teacher: result,
	})
}

func (_self TeacherHandlers) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := teacherServices.DeleteTeacher(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(SuccessResponse{
		Success: true,
	})
}

func (_self TeacherHandlers) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var teacher TeacherRequest

	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := teacher.validation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertedTeacher := transformTeacherRequestToTeacherModel(teacher)
	err := teacherServices.UpdateTeacher(id, &convertedTeacher)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(SuccessResponse{
		Success: true,
	})
}
