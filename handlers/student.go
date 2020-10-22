package handlers

import (
	"encoding/json"
	"net/http"
	"student_rest/models"

	"github.com/go-chi/chi"
	"student_rest/services"
)

type StudentHandlers struct{}

var studentServices = services.StudentServices{}

// CreateStudent creates new student
func (_self StudentHandlers) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student StudentRequest

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := student.validation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertedStudent := transformStudentRequestToStudentModel(student)
	result, err := studentServices.CreateStudent(convertedStudent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(StudentResponse{
		Success: true,
		Student: result,
	})
	return
}

func transformStudentRequestToStudentModel(request StudentRequest) models.StudentModel {
	return models.StudentModel{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBirth,
	}
}

func (_self StudentHandlers) GetStudents(w http.ResponseWriter, r *http.Request) {
	result, err := studentServices.GetStudents()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(StudentsResponse{
		Success:  true,
		Students: result,
	})
}

func (_self StudentHandlers) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result, err := studentServices.GetStudentByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(StudentResponse{
		Success: true,
		Student: result,
	})
}

func (_self StudentHandlers) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := studentServices.DeleteStudent(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(SuccessResponse{
		Success: true,
	})
}

func (_self StudentHandlers) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var student StudentRequest

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := student.validation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertedStudent := transformStudentRequestToStudentModel(student)
	err := studentServices.UpdateStudent(id, convertedStudent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(SuccessResponse{
		Success: true,
	})
}
//
//func (_self StudentHandlers) RegisterCourse(w http.ResponseWriter, r *http.Request) {
//	var body RegisterCourseRequest
//
//	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	if err := body.validation(); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	convertedBody := transformRegisterCourseRequestToRegisterCourseModel(body)
//	result, err := studentServices.CreateStudent(convertedBody)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	json.NewEncoder(w).Encode(RegisterCourseResponse{
//		Success: true,
//		Student: result,
//		Course:
//	})
//	return
//}
//
//func transformRegisterCourseRequestToRegisterCourseModel(request RegisterCourseRequest) models.RegisterCourseModel {
//	return models.RegisterCourseModel{
//		Course: models.CourseModel{
//			Name: request.Course.Name,
//			StartTime: request.Course.StartTime,
//			EndTime: request.Course.EndTime,
//			Teacher: models.TeacherModel{
//				ID: request.Course.TeacherID,
//			},
//		},
//		Student: models.StudentModel{
//			FirstName:   request.Student.FirstName,
//			LastName:    request.Student.LastName,
//			DateOfBirth: request.Student.DateOfBirth,
//		},
//	}
//}
