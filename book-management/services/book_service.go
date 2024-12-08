package services

import (
	"book-management/models"
	"fmt"
)

// Danh sách lưu trữ các sách
var Books []models.Book

// Hàm thêm sách mới vào danh sách
func AddBook1(title, author string) {
	book := models.Book{Title: title, Author: author}
	Books = append(Books, book)
	fmt.Println("Sách đã được thêm thành công!")
}

// Hàm hiển thị danh sách sách
func ListBooks() {
	if len(Books) == 0 {
		fmt.Println("Chưa có sách nào trong danh sách.")
		return
	}
	fmt.Println("Danh sách sách:")
	for i, book := range Books {
		fmt.Printf("%d. %s - %s\n", i+1, book.Title, book.Author)
	}
}

func AddBook(title string, author string) {
	book := models.Book{Title: title, Author: author}
	Books = append(Books, book)
	fmt.Println("Sách đã được thêm thành công!")
}
