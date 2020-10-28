TRUNCATE TABLE students_courses, students, teachers, courses;

INSERT INTO teachers(
	id, first_name, last_name, date_of_birth)
	VALUES (1, 'Anh', 'Le', '11/2/1998');

INSERT INTO courses(
	id, name, start_time, end_time, teacher_id)
	VALUES (1, 'Math', '11/2/2020', '11/3/2020', 1);



