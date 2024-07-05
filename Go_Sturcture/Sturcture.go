package main

import "fmt"

type student struct {
	Name  string
	Class int
	No    int
	cc    string
}

// 구조체를 포함하는 구조체
type sp_managedstudent struct {
	stuinfo student
	Score   int
	Subject string
}

// embedded 포함된 필드 방식
type sp_managedbadstudent struct {
	student
	Score   int
	Subject string
}

func main() {
	//구조체 여러 필드를 묶어서 사용하는 타입
	var astud student
	astud.Name = "No jin Gu"
	astud.Class = 3
	astud.No = 1

	var bdstud sp_managedbadstudent = sp_managedbadstudent{
		//여기는 student 로 bdstud 안의 student를 인베디드 하여 사용자의 이름등 std info를 다이렉트 엑세스로 불러올수있음
		student: astud,
		Score:   10,
		Subject: "Math",
	}
	var vipstd sp_managedstudent = sp_managedstudent{
		//여기는 stuinfo
		stuinfo: astud,
		Score:   100,
		Subject: "Math",
	}
	// 이렇게 하면 Student 구조체의 모든 값을 포함해야함
	var k1stud student = student{
		"Yes jin Gu",
		2,
		1,
		"yes",
	}
	// 이렇게 하면 student의 구조체를 다 포함할 필요가 없음 다만 구조체의 값을 안정해주면 각 타입의 기본값으로 변환힘
	var kstud student = student{
		Name:  "Yes jin Gu",
		Class: 2,
		No:    1,
	}
	fmt.Printf("학생의 이름은 %s, 학생의 반은 %d, 학생의 번호는 %d", k1stud.Name, k1stud.Class, k1stud.No)
	fmt.Printf("학생의 이름은 %s, 학생의 반은 %d, 학생의 번호는 %d", kstud.Name, kstud.Class, kstud.No)
	fmt.Printf("학생의 이름은 %s, 학생의 반은 %d, 학생의 번호는 %d", astud.Name, astud.Class, astud.No)
	inform(vipstd)
	inform1(bdstud)
}

func inform(user sp_managedstudent) {

	fmt.Printf("\n학생의 이름은 %s\n, 학생의 반은 %d\n, 학생의 번호는 %d\n, 학생의 과목은 %s\n, 학생의 점수는 %d\n", user.stuinfo.Name, user.stuinfo.Class, user.stuinfo.No, user.Subject, user.Score)
}

func inform1(user sp_managedbadstudent) {

	fmt.Printf("\n학생의 이름은 %s\n, 학생의 반은 %d\n, 학생의 번호는 %d\n, 학생의 과목은 %s\n, 학생의 점수는 %d\n", user.Name, user.Class, user.No, user.Subject, user.Score)
}

// Java의 클래스와 매우 흡사함.
// memory alignment - 메모리 정렬 - sorting, add the empty space in memory with the padding.
//CPU 속도를 고려해서 패딩을 추가 한다,
