## Golang - 배열과 슬라이스

- 배열을 선언하는 방법

```
var a = [2]int{1, 2}
```

- 다차원 배열을 선언하는 방법

```
var a = [2][3]int{
    {1, 1, 1},
    {2, 2, 2}
}
```

- 배열의 length와 capacity

배열의 length는 말 그대로 배열 element의 갯수를 의미한다.

배열의 length는 `len` 내장함수를 사용해서 확인할 수 있다.

```
var a = [2]int{1, 2}

// output: 2
fmt.Println(len(a))
```

배열의 capacity는 늘어날 수 있는 최대 용량을 의미한다.

배열의 capacity는 `cap` 내장함수를 사용해서 확인할 수 있다.

기본적으로 배열은 go에서 정적이기에 capacity와 length가 같다.

```
var a = [2]int{1, 2}

// output: 2, 2
fmt.Println(len(a), cap(a))
```

- 배열의 비교

==와 != operator를 통해서 값을 비교할 수 있는데, 참조[1]에는 모든 타입에 대해서 어떻게 비교가 이루어지는지 언급하고있다.

그 중에서 array는 다음과 같이 언급된다.

> Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.

즉, 배열의 타입이 같고 ([2][3]int == [2][3]int), 모든 배열의 요소의 값이 같은지 확인한후 같으면 true를 아니면 false를 결과값으로 리턴한다는 이야기가 된다.

이는 1차원 ~ n차원까지 모두 해당된다.

reference 비교를 하는 다른 언어와는 달리 훨씬 직관적인 비교를 수행하는 점이 마음에 든다.

go에서 배열의 reference 비교를 수행하려면 명확하게 `&`연산자를 통해서 reference 임을 나타내야한다.

```
var a = [4][2]int{ {1, 1}, {2, 2}, {3, 3}, {4, 4} }
var b = [4][2]int{ {1, 1}, {2, 2}, {3, 3}, {4, 4} }

// output : true
fmt.Println(a == b)

//output : false
fmt.Println(&a == &b)
```

[1] comparision operator spec : https://golang.org/ref/spec#Comparison_operators

- 슬라이스란?

슬라이스는 배열을 동적인 크기로 선언할 수 있는 방법을 제공한다. 슬라이스는 배열의 매타정보(length, capacity, pointer)를 통해서 동적인 배열을 다루도록 방법을 제공하는데, 배열을 가리키는 포인터라고 생각하고 사용하면 좋다.

슬라이스를 위한 다음과 같은 go 내장함수들이 제공된다.

- make(type, length, capacity)
- append(target, ...rest)
- copy(target, source)

`make`는 새로운 슬라이스를 생성하는 내장함수다. 이를 통해서 원하는 크기의 슬라이스를 만들 수 있고, 동적인 배열처럼 사용할 수 있는 방법을 제공한다.

`append`는 슬라이스에 새로운 데이터를 추가시키는 내장함수다.

`copy`는 슬라이스 및 참조하는 배열도 복사하는 내장함수다.

슬라이스는 유저에게 똑똑한 방식으로 배열을 사용할 수 있도록 내부 동작이 구현되어있다.

예를 들면 `make`를 통해서 length가 4이고 capacity가 8인 슬라이스를 생성하면, 크기가 8인 새로운 배열(내부 배열 ~ underlying array)을 생성하고 슬라이스의 length를 4로 설정해서 크기가 4인 배열을 갖는것처럼 보이게만든다.

그리고 배열에게 `append`를 통해서 크기를 늘리고싶으면, 슬라이스의 length를 증가시키고 이미 생성되어있는 내부 배열의 요소에 값을 넣는다.

만약 capacity이상으로 크기를 늘리게될 경우 capacity를 증가시키고, 새로운 내부 배열을 만들어서 기존에 있는 배열의 값을 복사시켜서 마치 동적인 배열처럼 사용할 수 있게 만든다.

```
package main

import (
	"fmt"
)

func main() {
	// 배열 선언 방법
	var a = [4][2]int{ {1, 1}, {2, 2}, {3, 3}, {4, 4} }

	fmt.Println("a is", len(a), cap(a))

	// 부분 슬라이스 생성
	sli1 := a[0:2]

	// 배열의 포인터라는 슬라이스 개념의 이해
	sli1[0][1] = 3
	fmt.Println("sli1 is", len(sli1), cap(sli1))

	// 슬라이스 생성
	sli2 := make([][2]int, len(sli1), cap(sli1) * 2)

	// 슬라이스 복사
	copy(sli2, sli1)

	// 2차원으로 구성된 슬라이스 복사시에 깊은 복사 테스트
	if &sli2[0] == &sli1[0] {
		fmt.Println("swallow copy")
	} else {
		fmt.Println("deep copy")
	}

	// 슬라이스 복사 후 원본 배열에 영향 테스트
	// 영향 x
	sli2[0][0] = 3
	fmt.Println(a, "|", sli1, "|", sli2)
}
```

- 슬라이스를 생성하는 방법들

첫째로 크기없는 배열 타입을 선언하면 슬라이스를 만들 수 있다.

```
This is an array literal:

[3]bool{true, true, false}

And this creates the same array as above, then builds a slice that references it:

[]bool{true, true, false}
```

둘째로 부분 슬라이스를 만드는 방법으로 생성할 수 있다.

```
var a = [4]int{1, 2, 3, 4}

// a[0:4]와 같다
b := a[:]
```

셋째로 위에서 언급한 `make` 내장함수를 사용할 수 있다.

```
a := make([]int, 3, 3)
```

## 추가적인 의문

Q) setns(int fd, int nstype) system call의 역할이 뭔가요?

setns는 이미 존재하는 namespace 파일을 현재 프로세스에게 할당하는 방법을 제공한다.

fd는 namespace 파일의 fd를 의미하고, nstype은 namespace 타입을 의미한다.

nstype은 namespace type 체크를 해주는 역할이고 기술하지않는다면 fd의 namespace type을 체크하지않는다.

Q) go에서 == 는 Equal 레퍼런스 비교인가요? 값만 비교하나요?

위의 `배열의 비교` 목차에서 이 내용을 다루고 있다.

Q) bind mount가 뭔가요?

```
mount --bind [source] [target]
```

source의 자원을 그대로 target에 복사해서 동기화시키는 마운트를 수행한다.

target을 변경하면 그대로 source에 반영된다. (+ vice versa)

Q) SIGCHILD를 왜 clone할때 flag로 넣나요?

```
## The child termination signal
When the child process terminates, a signal may be sent to the parent.  The termination signal is specified in the low byte of flags (clone()) or in cl_args.exit_signal (clone3()).
If this signal is specified as anything other than SIGCHLD, then the parent process must specify the __WALL or __WCLONE options when waiting for the child with wait(2).
If no signal (i.e., zero) is specified, then the parent process is not signaled when the child terminates.
```

해당 플래그를 넣으면 child process가 종료될때 parent process에게 종료되었음을 알려준다.

플래그를 넣지않으면, 종료되어도 parent process에게 알려주지않는다.

[1] namespace in go : https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a
