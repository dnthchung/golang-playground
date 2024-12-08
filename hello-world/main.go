package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Cấu trúc lưu trữ thông tin sách
type Book struct {
	Title  string
	Author string
}

// Danh sách lưu trữ các sách
var books []Book

// Hàm thêm sách mới vào danh sách
func addBook(title, author string) {
	book := Book{Title: title, Author: author}
	books = append(books, book)
	fmt.Println("Sách đã được thêm thành công!")
}

// Hàm hiển thị danh sách sách
func listBooks() {
	if len(books) == 0 {
		fmt.Println("Chưa có sách nào trong danh sách.")
		return
	}
	fmt.Println("Danh sách sách:")
	for i, book := range books {
		fmt.Printf("%d. %s - %s\n", i+1, book.Title, book.Author)
	}
}

// Hàm giới thiệu người dùng
func introduceUser(fullName, gender string, age int) {
	fmt.Println("==============================")
	fmt.Println("Họ và tên:", fullName)
	fmt.Println("Giới tính:", gender)
	fmt.Println("Tuổi:", age)
	// Sửa fmt.Sprintf thành fmt.Printf để in ra màn hình
	fmt.Printf("Xin chào %s, bạn là %s, hiện đang %d tuổi\n", fullName, gender, age)
}

// Hàm chính
func main() {

	// User name
	var name string
	var gender string
	var age int

	fmt.Println("Chào mừng bạn, nhập thông tin cá nhân trước khi bắt đầu sử dụng phần mềm")
	fmt.Print("Nhập họ và tên: ")
	fmt.Scanln(&name)
	fmt.Print("Nhập giới tính: ")
	fmt.Scanln(&gender)
	fmt.Print("Nhập tuổi: ")
	fmt.Scanln(&age)

	//introduct user
	introduceUser(name, gender, age)

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

			addBook(title, author)
		case "2":
			listBooks()
		case "3":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại.")
		}
	}
}
