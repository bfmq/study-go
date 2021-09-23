package main

import "fmt"

var (
	smr = studentMgr{
		make(map[int64]student, 10),
	}
)

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudents map[int64]student
}

func (s *studentMgr) showStudents() {
	for _, student := range s.allStudents {
		fmt.Println(student.id, student.name)
	}
}

func (s *studentMgr) addStudent(id int64, name string) {
	s.allStudents[id] = student{
		id:   id,
		name: name,
	}
}

func (s *studentMgr) editStudent(id int64, name string) {
	student, err := s.allStudents[id]
	if !err {
		fmt.Println("查无此人")
		return
	}
	student.name = name
	s.allStudents[id] = student
}

func (s *studentMgr) deleteStudent(id int64) {
	delete(s.allStudents, id)
}

func main() {
	smr.addStudent(1, "备份")
	smr.addStudent(2, "双色")
	smr.addStudent(3, "工单")
	smr.addStudent(4, "酷儿")
	smr.addStudent(5, "欧式")
	smr.showStudents()
	fmt.Println("===================")
	smr.deleteStudent(5)
	smr.showStudents()
	fmt.Println("===================")
	smr.editStudent(10, "觉得")
	smr.editStudent(1, "北方")
	smr.showStudents()
	fmt.Println("===================")
}
