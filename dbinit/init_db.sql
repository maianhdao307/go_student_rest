CREATE TABLE students (
	id serial PRIMARY KEY,
	student_id varchar(6) NOT NULL,
	first_name text NOT NULL,
	last_name text NOT NULL,
	date_of_birth timestamp NOT NULL
);

CREATE TABLE teachers (
	id serial PRIMARY KEY,
	first_name text NOT NULL,
	last_name text NOT NULL,
	date_of_birth timestamp NOT NULL
);

CREATE TABLE courses (
	id serial PRIMARY KEY,
	name text NOT NULL,
	start_time timestamp NOT NULL,
	end_time timestamp NOT NULL,
	teacher_id int NOT NULL,

	FOREIGN KEY (teacher_id) REFERENCES teachers(id)
);

CREATE TABLE students_courses (
	id serial PRIMARY KEY,
	student_id int NOT NULL,
	course_id int NOT NULL,

	FOREIGN KEY (student_id) REFERENCES students(id),
	FOREIGN KEY (course_id) REFERENCES courses(id)

);