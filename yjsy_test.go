package yjsy

import (
	"fmt"
	"os"
	"testing"
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

func Test_GetCourse(t *testing.T) {
	list, err := stu.GetSemesterCourses()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("course num:", len(list))

	// 不允许输出具体课程
}
