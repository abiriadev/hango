묶음 main

사용 "fmt"

함수 총합(목록 ...int) int {
	합 := 0

	반복 _, 값 := 범위 목록 {
		합 += 값
	}

	반환 합
}

함수 main() {
	fmt.Println(총합(1, 2, 3, 4, 5))
}
