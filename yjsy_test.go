package yjsy

import (
	"fmt"
	"os"
	"testing"

	"github.com/west2-online/yjsy/utils"
)

var (
	username = os.Getenv("YJSY_USERNAME") // 学号
	password = os.Getenv("YJSY_PASSWORD") // 密码
)
var (
	islogin = false
	stu     = NewStudent().WithUser(username, password)
)

func login() error {
	err := stu.Login()
	if err != nil {
		return err
	}

	islogin = true
	return nil
}
func TestMain(m *testing.M) {
	err := login()
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		os.Exit(1)
	}

	// 运行测试
	code := m.Run()

	// 在所有测试结束后执行清理
	os.Exit(code)
}

func TestGetExamRoomInfo(t *testing.T) {
	examRoom, err := stu.GetExamRoom(ExamRoomReq{
		Term: "XNXQ='2023-2024-1'",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(utils.PrintStruct(examRoom))
}

func Test_GetMarks(t *testing.T) {
	_, err := stu.GetMarks()
	if err != nil {
		t.Error(err)
	}

	// 不允许输出成绩

}

func Test_GetTerms(t *testing.T) {
	terms, err := stu.GetTerms()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("term :", terms.Terms)

}

func Test_GetCourse(t *testing.T) {
	terms, err := stu.GetTerms()
	if err != nil {
		t.Error(err)
	}

	list, err := stu.GetSemesterCourses(terms.Terms[0])
	if err != nil {
		t.Error(err)
	}

	fmt.Println("course num:", len(list))
}
