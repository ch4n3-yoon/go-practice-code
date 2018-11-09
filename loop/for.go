package main

import "fmt"

func main() {
	var num = 0
	var primes []int

	fmt.Print("몇 번 반복하실건가요? ")
	fmt.Scanf("%d", &num)

	fmt.Printf("%d번 반복을 시작합니다..\n", num)
	for i := 1; i <= num; i++ {
		if isPrime(i) {
			fmt.Printf("%d는 소수입니다. \n", i)
			primes = append(primes, i)
		}
	}

	fmt.Println("지정한 숫자까지 찾은 소수 리스트입니다.")
	fmt.Printf("[*] 총 %d 개가 발견되었습니다. \n", len(primes))
	fmt.Println(primes)
}

func isPrime(num int) bool {

	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}