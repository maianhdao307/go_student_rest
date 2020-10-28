TRUNCATE TABLE students_courses, students, teachers, courses;

INSERT INTO public.students(
	id, student_id, first_name, last_name, date_of_birth)
	VALUES (1, '123456', 'Anh', 'Le', '11/2/1998');

INSERT INTO students(
	id, student_id, first_name, last_name, date_of_birth)
	VALUES (2, '234567', 'Mai', 'Dao', '11/2/1998');


INSERT INTO teachers(
	id, first_name, last_name, date_of_birth)
	VALUES (1, 'Anh', 'Le', '11/2/1998');

INSERT INTO teachers(
	id, first_name, last_name, date_of_birth)
	VALUES (2, 'Duyen', 'Nguyen', '11/2/1998');


INSERT INTO courses(
	id, name, start_time, end_time, teacher_id)
	VALUES (1, 'Math', '11/2/2020', '11/3/2020', 1);

INSERT INTO courses(
	id, name, start_time, end_time, teacher_id)
	VALUES (2, 'Physics', '11/2/2020', '11/3/2020', 2);


INSERT INTO students_courses(
	id, student_id, course_id)
	VALUES (1, 1, 1);




