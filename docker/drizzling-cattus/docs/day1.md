### package와 import

```
package main


// factored import statement
// 공식문서에서는 이 방법으로 import하는걸 추천합니다.
import (
  "fmt"
  "math/rand"
)

func main() {
  // ...
}
```

package는 모듈이라고 생각할 수 있습니다.

관습적으로 package이름은 import path의 마지막 인자를 사용합니다.

그 중에서 main package는 특별한 package로서 main함수를 포함하고 실행파일이 되는 패키지입니다.

Go에서는 첫 문자가 대문자로 시작하면 그 패키지를 사용하는 곳에서 접근할 수 있는 exported name이 됩니다.

### 함수 선언

```
package test

import "fmt"

// func 함수명(매개변수...) 리턴타입 을 지니는 기본적인 함수 선언 방법
func foo1(sum int) int {
}

// 리턴 타입에 이름을 붙일 수 있습니다.
// 그러면 리턴만 해도 무엇을 리턴할지 자동으로 컴파일러에서 추리해줍니다.
func foo2(sum int) (x int) {
  return
}

// 리턴을 여러개할 수 도 있습니다.
// a, b, := foo3(1)
func foo3(sum int) (x int, y int) {
}

// 중복된 타입을 한번에 묶어서 사용할 수 있습니다.
func foo4(sum, hello int) {
}
```

### 배열 선언 및 slice

```
package main

import "fmt"

// arr[srt:dst]
// index가 srt부터 dst전까지 array를 slice합니다.
// srt나 dst는 생략할 수 있고 배열의 시작과 끝을 의미하게 됩니다.
// arr[0:2] == { arr[0], arr[1] }
func splice(arr []int, index int) (x, y []int) {
  x = arr[:index]
  y = arr[index:]
  return
}

func main() {
  // int형 배열을 선언하고 초기화합니다.
  arr := []int { 1, 2, 3, 4 }

  // [1, 2] [3, 4]이 출력됩니다.
  fmt.Println(splice(arr, 2))
}
```

### 변수 선언 및 타입 추론

```
// var를 통해서 변수를 선언할 수 있습니다.
// 초기와된 값에 따라 자동적으로 타입추론을 합니다.
var a, b, c = 1, 2, 3

// 명시적으로 타입을 기술할 수 있습니다.
var a, b, c int = 1, 2, 3

// 하나의 var에 각각 다른 타입을 기술할 수 는 없습니다.
var a, b, c int, bool, int

// := 는 `short assignment statement`라고 합니다.
// var와 type을 생략하고 기술할 수 있습니다.
// 기본적인 초기화 할당과 같습니다.
d := 1

// 재할당은 불가합니다.
var e = 2
e := 3

// 상수를 선언할 수 있습니다.
const g = "hello"
const h = 1.0

// 상수를 사용할때 := 를 사용할 수 없습니다.
const a := "wrong"
```

### if문 & if-else

```
// 기본적인 if문, 특이한 점으로 소괄호가 없습니다.
if a < 10 {

}

// 안쓸 것 같은 문법
// if 조건식 안에 세미콜론을 통해서 여러개의 표현식을 넣을 수 있습니다.
// 표현식에서 사용한 변수들은 if 및 if-else안에서 사용할 수 있습니다.
if v := 1; v > 10 {

} else if v < 10 {

} else {

}
```

### 반복문

```
// 기본적인 for문
// 짧은 선언이 유용합니다.
for a := 1; a < 10; a++ {

}

// while문처럼 사용할 수 있습니다.
for a < 10 {

}

// 무한루프를 사용할 수 있습니다.
for {

}
```
