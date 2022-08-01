package main

import "golang/src/other"

func main() {
	const text string = "사용할 텍스트"
	text2 := "text"
	importText("패키지 단위로 동작한다.") // 패키지 단위에서는 모두 참조가능.
	other.ImportTest(text)
	other.ImportTest(text2)
	//other.canNotImport("text");
}
