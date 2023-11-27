# 한고

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

[^식]: 사실 한고 자체는 식별자가 아닌 키워드만 한글화하므로 본 문제와 무관합니다. 단, 100% 한글만으로 이루어진 식별자를 공개하는 것이 애시당초 불가능함을 일깨워 줄 수 있습니다.