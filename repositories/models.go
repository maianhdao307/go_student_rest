package repositories

type StudentEntity struct {
	ID          int    `json:"id"`
	StudentID   string `json:"studentID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
}

type TeacherEntity struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
}

type CourseEntity struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TeacherID int    `json:"teacherID"`
}
