# linked list vs array

## array

- 연속 메모리 -> 한번에 할당되고 한 번에 해제된다
- Random Access 에 강하다
    - a[33] = 100 세팅한다고 하면, index 33 이라는 메모리 주소를 찾아야 한다
    - 100개의 공간을 확보 -> a는 그 공간의 시작주소이므로 100 + 33 X 8(type size)
    - +, x 한 번 만 하면 되므로 컴뷰터에서는 빠른 속도로 가능 -> 랜덤에 강하다
    - O(1) - 크기와 index 와 상관 없이 +, x 계산 한 번에 되므로 O(1)
- 삽입/삭제에 약하다(배열 끝 추가와 삭제는 괜찮다)
    - 중간에 삽입한다고 하면 - 앞의 것들 하나씩 복붙 + 삽입 + 뒤의 것들 하나씩 복붙
    - 따라서 삽입과 삭제는 배열의 요소 개수만큼 시간이 소요된다
    - O(N) - 배열의 개수만큼 시간이 소요된다

# Linked List 역순
## double linked list 역순
- node previous 와 node next 를 바꿔주고 처음과 끝을 바꿔준다

## Single linked list 역순
- ex. 1-> 2-> 3
- popFront 1
- popFront 2 해서 1 앞에 붙인다 2 -> 1
- popFront 2 해서 2 앞에 붙인다 3 -> 2-> 1
- 