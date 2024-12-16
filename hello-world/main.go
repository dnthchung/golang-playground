package main

import "fmt"

// (string, string) => type of return value
func swap(x, y string) (string, string) {
	return y, x
}

// x int, y int => x, y int
func add(x, y int) int {
	return x + y
}

// sum là tên của tham số input, x và y là tên của tham số dc trả về
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i, j int = 1, 2

//ne := 3 // error =>  vì chỉ dùng ":=" ở trong 1 function
//c, python, java := true, false, "no!"

// giá trị mặc định của biến :
/**
 * 0 cho số
 * false cho bool
 * "" cho string
 */

func main() {

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}

	/**
	break: Dừng vòng lặp ngay lập tức.
	continue: Bỏ qua phần còn lại của vòng lặp và bắt đầu vòng mới.
	*/
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // Bỏ qua lần lặp khi i == 3
		}
		if i == 7 {
			break // Dừng vòng lặp khi i == 7
		}
		fmt.Println("Giá trị của i:", i)
	}
}
