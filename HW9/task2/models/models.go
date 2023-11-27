package models

type User struct {
	Username string
	Password string
	Id       int
}

type Teacher struct {
	Name   string
	UserId int
}

type Student struct {
	Name  string
	Id    int
	Rates map[string]int
}

type Class struct {
	Name     string
	Teacher  Teacher
	Students []Student
}

func (class Class) GetStudentById(id int) (student Student, success bool) {
	for _, s := range class.Students {
		if s.Id == id {
			return s, true
		}
	}
	return Student{}, false
}

type ClassPreview struct {
	Name    string
	Teacher string
}

type SchoolPreview struct {
	Name           string
	ClassesPreview []ClassPreview
}

type School struct {
	Name    string
	Classes []Class
}

func (school School) GetPreview() SchoolPreview {
	classesPreview := []ClassPreview{}
	for _, class := range school.Classes {
		classPreview := ClassPreview{
			Name:    class.Name,
			Teacher: class.Teacher.Name,
		}
		classesPreview = append(classesPreview, classPreview)
	}
	schoolPreview := SchoolPreview{
		Name:           school.Name,
		ClassesPreview: classesPreview,
	}
	return schoolPreview
}

func (school School) GetClassByName(name string) (class Class, success bool) {
	for _, c := range school.Classes {
		if c.Name == name {
			return c, true
		}
	}
	return Class{}, false
}

func (school School) GetTeachersStudentById(teacher Teacher, studentId int) (student Student, found bool) {
	for _, c := range school.Classes {
		if c.Teacher.Name == teacher.Name {
			student, found = c.GetStudentById(studentId)
			if found {
				return student, found
			}
		}
	}
	return Student{}, false
}
