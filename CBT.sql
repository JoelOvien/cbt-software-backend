CREATE TABLE "Students" (
  "id" integer PRIMARY KEY,
  "full_name" varchar(255) NOT NULL,
  "matric_number" varchar(50) UNIQUE NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "password" varchar(50) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Examiners" (
  "id" integer PRIMARY KEY,
  "full_name" varchar(255) NOT NULL,
  "staff_number" varchar(50) UNIQUE NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "password" varchar(50) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Administrators" (
  "id" integer PRIMARY KEY,
  "full_name" varchar(255) NOT NULL,
  "matric_number" varchar(50) UNIQUE NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "password" varchar(50) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Exams" (
  "id" integer PRIMARY KEY,
  "course_title" TEXT NOT NULL,
  "course_code" TEXT NOT NULL,
  "course_duration" integer NOT NULL,
  "examiner" Examiners NOT NULL,
  "passing_score" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Questions" (
  "id" integer PRIMARY KEY,
  "course_code" TEXT UNIQUE NOT NULL,
  "question" TEXT,
  "type" varchar(5) NOT NULL,
  "answer" TEXT,
  "options" TEXT[],
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Answers" (
  "id" integer PRIMARY KEY,
  "session_id" int UNIQUE NOT NULL,
  "question_id" int UNIQUE NOT NULL,
  "type" varchar(5) NOT NULL,
  "answer" TEXT NOT NULL
);

CREATE TABLE "ExamSessions" (
  "id" integer PRIMARY KEY,
  "exam_id" int UNIQUE NOT NULL,
  "start_time" int NOT NULL,
  "end_time" int NOT NULL,
  "score" int NOT NULL
);
