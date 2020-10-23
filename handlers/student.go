package handlers

import (
	"encoding/json"
	"net/http"
	"student_rest/models"

	"student_rest/services"
)

type StudentHandlers struct{
	services.StudentServices
}

// CreateStudent creates new student
//func (_self StudentHandlers) CreateStudent(w http.ResponseWriter, r *http.Request) {
//	var student StudentRequest
//
//	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	if err := student.validation(); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	convertedStudent := transformStudentRequestToStudentModel(student)
//	result, err := studentServices.CreateStudent(&convertedStudent)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	json.NewEncoder(w).Encode(StudentResponse{
//		Success: true,
//		Student: result,
//	})
//	return
//}
//
func transformStudentRequestToStudentModel(request StudentRequest) *models.StudentModel {
	return &models.StudentModel{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBirth,
	}
}
//
//func (_self StudentHandlers) GetStudents(w http.ResponseWriter, r *http.Request) {
//	result, err := studentServices.GetStudents()
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	json.NewEncoder(w).Encode(StudentsResponse{
//		Success:  true,
//		Students: result,
//	})
//}
//
//func (_self StudentHandlers) GetStudentByID(w http.ResponseWriter, r *http.Request) {
//	id := chi.URLParam(r, "id")
//	result, err := studentServices.GetStudentByID(id)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	json.NewEncoder(w).Encode(StudentResponse{
//		Success: true,
//		Student: result,
//	})
//}
//
//func (_self StudentHandlers) DeleteStudent(w http.ResponseWriter, r *http.Request) {
//	id := chi.URLParam(r, "id")
//
//	if err := studentServices.DeleteStudent(id); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	json.NewEncoder(w).Encode(SuccessResponse{
//		Success: true,
//	})
//}
//
//func (_self StudentHandlers) UpdateStudent(w http.ResponseWriter, r *http.Request) {
//	id := chi.URLParam(r, "id")
//
//	var student StudentRequest
//
//	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	if err := student.validation(); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	convertedStudent := transformStudentRequestToStudentModel(student)
//	err := studentServices.UpdateStudent(id, &convertedStudent)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	json.NewEncoder(w).Encode(SuccessResponse{
//		Success: true,
//	})
//}



func (_self StudentHandlers) RegisterCourse(w http.ResponseWriter, r *http.Request) {
	var registerCourseRequest RegisterCourseRequest

	if err := json.NewDecoder(r.Body).Decode(&registerCourseRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := registerCourseRequest.validation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registerCourseModel := models.RegisterCourseModel{
		Student: transformStudentRequestToStudentModel(*registerCourseRequest.Student),
		Course: &models.CourseModel{
			Name: registerCourseRequest.Course.Name,
			StartTime: registerCourseRequest.Course.StartTime,
			EndTime: registerCourseRequest.Course.EndTime,
			Teacher: &models.TeacherModel{
				FirstName: registerCourseRequest.Course.Teacher.FirstName,
				LastName: registerCourseRequest.Course.Teacher.LastName,
				DateOfBirth: registerCourseRequest.Course.Teacher.DateOfBirth,
			},
		},
	}
	result, err := _self.StudentServices.RegisterCourse(&registerCourseModel)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(RegisterCourseResponse{
		Success: true,
		Student: result.Student,
		Course: result.Course,
	})
	return
}
