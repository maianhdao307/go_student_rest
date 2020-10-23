package models

type StudentModel struct {
	ID          int
	StudentID   string
	FirstName   string
	LastName    string
	DateOfBirth string
}

type UpdateStudentModel struct {
	FirstName   string
	LastName    string
	DateOfBirth string
}

type TeacherModel struct {
	ID          int
	FirstName   string
	LastName    string
	DateOfBirth string
}

type CourseModel struct {
	ID        int
	Name      string
	StartTime string
	EndTime   string
	Teacher   *TeacherModel
}

type RegisterCourseModel struct {
	Student *StudentModel
	Course  *CourseModel

}
