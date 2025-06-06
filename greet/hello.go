package greet

import "fmt"

const jap string = "こんにちは,おやすみ" //この変数は同じパッケージからアクセスできる

func Hello() {
	fmt.Println("Hello from greet!")
}

func HelloJapanese() {
	fmt.Println("こんにちは from greet!")
}

func anotherHello() {
	fmt.Println("小文字の関数からのこんにちは from greet!")
}

func AnotherHelloJapanese() {
	anotherHello() //外部アクセスできない関数を使った
	fmt.Println("別のこんにちは from greet!")
}
