package main

import (
	"fmt"
	"math/rand"
	"time" // seed 생성용 패키지
)

/*
func main() {
	//난수 추출된 수의 소수 판정 프로그램 v0.6
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

/*
//난수 추출된 수의 소수 판정 프로그램 v0.1
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0과 1 제외
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}

	if count == 2 {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

/*
//난수 추출된 수의 소수 판정 프로그램 v0.2
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0과 1 제외
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ { // 1과 number일 때 반복문을 돌지않음
		if number%i == 0 {
			count++
		}
	}

	if count == 0 {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

//난수 추출된 수의 소수 판정 프로그램 v0.3
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2 //
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
		}
	}

	if isPrime == true {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}

/*
//난수 추출된 수의 소수 판정 프로그램 v0.4
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2 //
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
		}
	}

	if isPrime { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
*/

/*
package main

import (
	"fmt"
	"math/rand"
	"time" // seed 생성용 패키지
)

//무작위 주사위
func main() {
	// seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	for i := 1; i < 6; i++ {
		dice := rand.Intn(6) + 1
		fmt.Println(dice)
	}
}
*/
