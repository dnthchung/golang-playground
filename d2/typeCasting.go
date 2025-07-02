package main

import "fmt"

func main() {
    totalScore := 85 // Kiểu int
    subjectCount := 10 // Kiểu int

    // Sai: Chia hai số int sẽ cho ra kết quả là int (85 / 10 = 8)
    // average := totalScore / subjectCount

    // TODO: Sửa dòng dưới đây
    // Ép kiểu totalScore và subjectCount thành float64 trước khi chia

	average := float64(totalScore) / float64(subjectCount)


    fmt.Println("Điểm trung bình là:", average) // Kết quả mong đợi: 8.5
}
