package main

import (
	"fmt"
)

// count sort
/*
배열의 숫자의 인덱스에 1씩 더한다 - 숫자가 몇번 나왔는지 counting   => N 번
counting 결과를 돌면서 몇번 나왔는지에 따라 숫자 출력   => K 번

=> O(N+K) = O(N) 속도 빠름
단, for문이 두개라서 성능 저하
*/

// 문제: 0~10 사이 값을 갖는 배열을 정렬하시오

func countSort() {
	arr := []int{5, 3, 7, 3, 8, 8, 1, 1, 8, 10, 0, 10, 8, 2, 1, 2, 7, 4, 3}
	var count [11]int // 값의 범위만큼 개수의 배열 만들기
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	fmt.Println("count: ", count)

	sorted := make([]int, 0, len(arr))
	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ { // count 개수 만큼 넣어주어야 함
			sorted = append(sorted, i)
		}
	}
	fmt.Println("sorted: ", sorted)
}

/*
위의 과정처럼 숫자만큼 해당 인덱스에 +1 하는 것은 동일
counting 숫자를 이전 인덱스의 값에 더하여 정렬
ex) count = [1, 1, 0, 1] -> [1, 1+1, 1+1+0, 1+1+0+1]=[1, 2, 2, 3]
더한 값이 숫자의 자리가 된다 => 0은 1번째 자리, 1은 2번째 자리, 2는 없고(숫자 변화가 없으니), 3은 세번째 자리
*/

func upgradeCountSort() {
	arr := []int{0, 5, 3, 7, 3, 8, 8, 1, 1, 8, 10, 0, 10, 8, 2, 1, 2, 7, 4, 3}
	var count [11]int
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	fmt.Println("count: ", count)

	for i := 1; i < 11; i++ {
		count[i] += count[i-1] // 이전 값을 하나씩 더하기 -> 인덱스가 아닌 몇번째로 읽힌다
	}
	fmt.Println("count2: ", count)
	sorted := make([]int, 0, len(arr))
	for i := 1; i < len(arr); i++ {
		sorted[count[arr[i]-1]] = arr[i]
		count[arr[i]]--
	}
	fmt.Println("sorted: ", sorted)
}

// 문제: 알파벳 소문자로 이루어진 문자열 중 가장 많이 나오는 알파벳 출력
func countAlphabet() {
	str := "dkqhdklfasedtflaeklgohqjshdjdfklawhw"

	var count [26]int // 알파벳 개수 26
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++ // a가 첫번째로 와야하고, 알파벳에 해당하는 count 증가
	}
	maxCount := 0
	var maxCh byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCount {
			maxCount = count[i]   // 가장 count가 큰 것은 반복문으로 하나씩 비교하면서 maxCount 업데이트 하는 방식!
			maxCh = byte('a' + i) // 숫자를 알파벳으로 바꾸는 방법
		}
	}

	fmt.Printf("Max Character: %c count: %d\n", maxCh, maxCount)

}

// 문제: 한 번의 학생들 중 키가 특정 범위의 학생들만 출력하시오. 범위값은 여러번 입력받을 수 있습니디. 키는 소수점 1자리 까지 주어집니다.
/*
for 반복문 돌면서 if문으로 하나하나 비교할 수 도 있지만.. O(N)
이 방법으로 입력이 M번 들어온다면 NxM => M번이 커질 수록 결국 O(N^2) 이 되어버린다. - 입력이 여러번 들어올 때는 비효율적

*/

func studentHeight() {
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "Kyle", Height: 173.4},
		{Name: "Ken", Height: 164.5},
		{Name: "Ryu", Height: 178.8},
		{Name: "Honda", Height: 154.2},
		{Name: "Hwarang", Height: 188.8},
		{Name: "Lebron", Height: 209.8},
		{Name: "Hodong", Height: 197.7},
		{Name: "Tom", Height: 164.8},
	}
	// 키에 따른 배열을 만든다
	var heightMap [3000][]string // 0 < Height < 3m 범위의 배열을 만듬
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heightMap[idx] = append(heightMap[idx], students[i].Name) // 3000 개의 배열에 각 키 값에 맞는 학생 이름이 들어가있음
	}
	fmt.Println(heightMap)

	// 140 ~ 170
	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name: ", name, "height: ", float64(i)/10)
		}
	}
	// 학생이 많아지고, 질문의 개수가 많아지더라도 배열의 개수인 3천번 밖에 돌지 않기 때문에 훨씬 효율적
}

func main() {
	// countSort()
	upgradeCountSort()
	//countAlphabet()
	// studentHeight()
}
