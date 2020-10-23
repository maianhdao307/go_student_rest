package routes

import (
	"database/sql"
	"github.com/go-chi/chi"
	"student_rest/handlers"
	"student_rest/repositories"
	"student_rest/services"
)

func CreateRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/students", func(r chi.Router) {
		studentHandlers := handlers.StudentHandlers{
			StudentServices: services.Student{
				StudentRepositories: repositories.Student{
					Db : db,
				},
			},
		}

		//r.MethodFunc("post", "/", studentHandlers.CreateStudent)
		//r.MethodFunc("get", "/", studentHandlers.GetStudents)
		//r.MethodFunc("get", "/student/{id}", studentHandlers.GetStudentByID)
		//r.MethodFunc("delete", "/student/{id}", studentHandlers.DeleteStudent)
		//r.MethodFunc("put", "/student/{id}", studentHandlers.UpdateStudent)
		r.MethodFunc("post", "/register-course", studentHandlers.RegisterCourse)
	})

	r.Route("/teachers", func(r chi.Router) {
		//r.MethodFunc("post", "/", teacherHandlers.CreateTeacher)
		//r.MethodFunc("get", "/", teacherHandlers.GetTeachers)
		//r.MethodFunc("get", "/teacher/{id}", teacherHandlers.GetTeacherByID)
		//r.MethodFunc("delete", "/teacher/{id}", teacherHandlers.DeleteTeacher)
		//r.MethodFunc("put", "/teacher/{id}", teacherHandlers.UpdateTeacher)
	})
	r.Route("/courses", func(r chi.Router) {
		//r.MethodFunc("post", "/", courseHandlers.CreateCourse)
		//r.MethodFunc("get", "/course/{id}", courseHandlers.GetCourseByID)
		//r.MethodFunc("delete", "/course/{id}", courseHandlers.DeleteCourse)
		//r.MethodFunc("put", "/course/{id}", courseHandlers.UpdateCourse)
	})
	return r
}
