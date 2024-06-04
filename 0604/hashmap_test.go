package _604

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const arraySize = 3571

type Data struct {
	key   string
	value string
	next  *Data
}

type HashMap struct {
	arr [arraySize]*Data
}

func hashFunction(key string) uint32 {
	return uint32(len(key) % 10)
}

func (h *HashMap) add(key string, value string) {
	hashKey := hashFunction(key)
	head := h.arr[hashKey]
	newData := &Data{
		key:   key,
		value: value,
	}

	if head != nil {
		// key가 일치한다면 덮어쓰기
		if head.key == key {
			head.value = value
			return
		}
		newData.next = head
	}
	h.arr[hashKey] = newData
}

func (h *HashMap) get(key string) (string, bool) {
	hashKey := hashFunction(key)
	head := h.arr[hashKey]

	for head != nil {
		if head.key == key {
			return head.value, true
		}
		head = head.next
	}
	return "", false
}

func (h *HashMap) remove(key string) bool {
	hashKey := hashFunction(key)
	head := h.arr[hashKey]

	if head == nil {
		return false
	}

	// 첫번째 것일 경우
	if head.key == key {
		h.arr[hashKey] = head.next
	}

	// 첫번째 다음 부터 탐색
	prev := head
	head = head.next
	for head != nil {
		if head.key == key {
			prev.next = head.next
			return true
		}
		prev = head
		head = head.next
	}
	return false
}

func TestHashmap(t *testing.T) {
	var hm HashMap
	hm.add("abc", "1번")
	hm.add("abcd", "2번")
	hm.add("123", "1-1번")
	hm.add("456", "1-1-1")

	v, _ := hm.get("abc")
	a, _ := hm.get("123")
	assert.Equal(t, v, "1번")
	assert.Equal(t, a, "1-1번")

	hm.remove("123")
	_, ok := hm.get("123")
	assert.Equal(t, ok, false)

}

/*
리스트에 추가하는 방식으로 해쉬 충돌 방지
*/

type DataV struct {
	key   string
	value string
}

type HashMapList struct {
	arr [1234][]DataV
}

func (hl *HashMapList) add(key string, value string) {
	hashKey := hashFunction(key)
	newData := DataV{
		key:   key,
		value: value,
	}
	for _, data := range hl.arr[hashKey] {
		if data.key == key {
			data.value = value
		}
	}
	hl.arr[hashKey] = append(hl.arr[hashKey], newData)
}

func (hl *HashMapList) get(key string) (string, bool) {
	hashKey := hashFunction(key)
	if hl.arr[hashKey] == nil {
		return "", false
	}
	for _, data := range hl.arr[hashKey] {
		if data.key == key {
			return data.value, true
		}
	}
	return "", false
}

func (hl *HashMapList) remove(key string) bool {
	hashKey := hashFunction(key)
	if hl.arr[hashKey] == nil {
		return false
	}

	for i, data := range hl.arr[hashKey] {
		if data.key == key {
			hl.arr[hashKey] = append(hl.arr[hashKey][:i], hl.arr[hashKey][i:]...)
			return true
		}
	}
	return false
}

func TestHashMapList(t *testing.T) {
	hashMap := HashMapList{}
	hashMap.add("123", "1번")
	hashMap.add("456", "1-1번")
	hashMap.add("789", "1-2번")
	hashMap.add("1234", "2번")

	v, _ := hashMap.get("123")
	assert.Equal(t, "1번", v)

	hashMap.remove("456")
}

/*
당근 문제 복기
*/

/*
key = 피신고자
value = []신고자
*/
func Report(input []string) (string, int) {
	reportCount := make(map[string]int)
	uniqueReport := make(map[string]struct{}) // value 는 {} 빈 구조체, map으로 해당 Key(신고내용)이 존재하는지만 확인하기 위함

	for _, report := range input {
		if _, ok := uniqueReport[report]; ok {
			continue
		} else {
			parts := strings.Split(report, " ")
			reporter := parts[0]
			bad := parts[1]

			if reporter == bad {
				continue
			}
			uniqueReport[report] = struct{}{}
			reportCount[bad]++
		}
	}

	var maxReported string
	maxCount := 0
	for key, value := range reportCount {
		if value > maxCount {
			maxCount = value
			maxReported = key
		}
	}
	return maxReported, maxCount
}

func TestRepost(t *testing.T) {
	input := []string{"donis donis", "donis rozy", "donis rozy", "suji rozy", "rozy donis"}
	v, c := Report(input)
	println(v, c)
}
