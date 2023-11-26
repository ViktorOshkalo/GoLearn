package models

type User struct {
	Username string
	Password string
	Id       int
}

type Teacher struct {
	Name string
}

type Student struct {
	Name  string
	Rates map[string]int
}

type Class struct {
	Name     string
	Teacher  Teacher
	Students []Student
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
	success = false
	for _, c := range school.Classes {
		if c.Name == name {
			return c, true
		}
	}
	return
}
