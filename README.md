<h1 align="center" style="font-size: 200em;">한고</h1>

<p align="center">Go로 컴파일되는 한글 프로그래밍 언어</p>

## 개요

한고는 Go에 기반한 한글 프로그래밍 언어로, 다음 개념의 증명을 위해 만들어졌습니다.

1. [라틴 문자에 불필요하게 의존하는 식별자 공개 원칙](https://go.dev/doc/faq#unicode_identifiers)에 대한 신선한 비판[^식]
2. 외부 라이브러리니 직접 짠 렉서 혹은 파서 없이 내장 text/scanner만으로 구현이 가능함의 증명

## 예시

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./examples/hello.go) -->
<!-- The below code snippet is automatically added from ./examples/hello.go -->

```go
묶음 main

사용 "fmt"

함수 main() {
	fmt.Println("안녕, 세상!")
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

## 키워드 변환 테이블

| 한고 키워드 | Go 키워드   | 한고 키워드 | Go 키워드     |
| ----------- | ----------- | ----------- | ------------- |
| `정지`      | `break`     | `묶음`      | `package`     |
| `기본`      | `default`   | `비교`      | `switch`      |
| `함수`      | `func`      | `상수`      | `const`       |
| `특성`      | `interface` | `이어서`    | `fallthrough` |
| `선택`      | `select`    | `만약`      | `if`          |
| `경우`      | `case`      | `범위`      | `range`       |
| `지연`      | `defer`     | `형`        | `type`        |
| `출발`      | `go`        | `계속`      | `continue`    |
| `사전`      | `map`       | `반복`      | `for`         |
| `구조체`    | `struct`    | `사용`      | `import`      |
| `통로`      | `chan`      | `반환`      | `return`      |
| `아니라면`  | `else`      | `선언`      | `var`         |
| `이동`      | `goto`      |             |               |

## 고민중인 기능

본 프로젝트의 원 목적은 <ins>표준 라이브러리만 사용하여</ins> 100줄 이내의 짧은 코드만으로 Go와 1:1 변환 가능한 언어를 구현하는 것이었습니다.

당초 목적한 구현이 표준 라이브러리 없이 가능함이 증명되었으므로 앞으로는 다음과 같은 기능을 차차 추가할 예정입니다:

-   [ ] `공개` 키워드
-   [ ] Context-sensitive transformations: 자체 파서 구현 필요
-   [ ] Syntax highlighting support
-   [ ] 더 적은 키워드
-   [ ] 기본 타입(`int`, `string`), 빌트인 함수(`len`, `make`) 지원: name resolution needed

[^식]: 사실 한고 자체는 식별자가 아닌 키워드만 한글화하므로 본 문제와 무관합니다. 단, 100% 한글만으로 이루어진 식별자를 공개하는 것이 애시당초 불가능함을 일깨워 줄 수 있습니다.
