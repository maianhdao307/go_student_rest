package routes

import (
	"github.com/go-chi/chi"
	"student_rest/handlers"
)

func CreateRoutes() *chi.Mux {
	r := chi.NewRouter()

	studentHandlers := handlers.StudentHandlers{}
	teacherHandlers := handlers.TeacherHandlers{}
	courseHandlers := handlers.CourseHandlers{}

	r.Route("/students", func(r chi.Router) {
		r.MethodFunc("post", "/", studentHandlers.CreateStudent)
		r.MethodFunc("get", "/", studentHandlers.GetStudents)
		r.MethodFunc("get", "/student/{id}", studentHandlers.GetStudentByID)
		r.MethodFunc("delete", "/student/{id}", studentHandlers.DeleteStudent)
		r.MethodFunc("put", "/student/{id}", studentHandlers.UpdateStudent)
		r.MethodFunc("post", "/student/register-course", studentHandlers.UpdateStudent)
	})
	r.Route("/teachers", func(r chi.Router) {
		r.MethodFunc("post", "/", teacherHandlers.CreateTeacher)
		r.MethodFunc("get", "/", teacherHandlers.GetTeachers)
		r.MethodFunc("get", "/teacher/{id}", teacherHandlers.GetTeacherByID)
		r.MethodFunc("delete", "/teacher/{id}", teacherHandlers.DeleteTeacher)
		r.MethodFunc("put", "/teacher/{id}", teacherHandlers.UpdateTeacher)
	})
	r.Route("/courses", func(r chi.Router) {
		r.MethodFunc("post", "/", courseHandlers.CreateCourse)
		r.MethodFunc("get", "/course/{id}", courseHandlers.GetCourseByID)
		r.MethodFunc("delete", "/course/{id}", courseHandlers.DeleteCourse)
		r.MethodFunc("put", "/course/{id}", courseHandlers.UpdateCourse)
	})
	return r
}
