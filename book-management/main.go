package main

import (
	"book-management/services"
	"book-management/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Nhập thông tin người dùng
	var name, gender string
	var age int

	fmt.Println("Chào mừng bạn, nhập thông tin cá nhân trước khi bắt đầu sử dụng phần mềm")
	fmt.Print("Nhập họ và tên: ")
	fmt.Scanln(&name)
	fmt.Print("Nhập giới tính: ")
	fmt.Scanln(&gender)
	fmt.Print("Nhập tuổi: ")
	fmt.Scanln(&age)

	// Giới thiệu người dùng
	utils.IntroduceUser(name, gender, age)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nQuản lý sách:")
		fmt.Println("1. Thêm sách")
		fmt.Println("2. Hiển thị danh sách sách")
		fmt.Println("3. Thoát")
		fmt.Print("Chọn một tùy chọn: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Nhập tên sách: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())

			fmt.Print("Nhập tên tác giả: ")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())

			services.AddBook(title, author)
		case "2":
			services.ListBooks()
		case "3":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại.")
		}
	}
}
