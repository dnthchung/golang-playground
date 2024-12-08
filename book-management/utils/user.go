package utils

import "fmt"

// Hàm giới thiệu người dùng
func IntroduceUser1(fullName, gender string, age int) {
	fmt.Println("==============================")
	fmt.Println("Họ và tên:", fullName)
	fmt.Println("Giới tính:", gender)
	fmt.Println("Tuổi:", age)
	fmt.Printf("Xin chào %s, bạn là %s, hiện đang %d tuổi\n", fullName, gender, age)
}

func IntroduceUser(name string, gender string, age int) {
	fmt.Println("==============================")
	fmt.Println("Họ và tên:", name)
	fmt.Println("Giới tính:", gender)
	fmt.Println("Tuổi:", age)
	fmt.Printf("Xin chào %s, bạn là %s, hiện đang %d tuổi\n", name, gender, age)
}
