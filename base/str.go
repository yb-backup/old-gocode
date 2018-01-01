package main

import "fmt"

func main() {
	var str string
	str = "Hello 世界"
	ch := str[0]
	fmt.Printf("The lenght of %s is %d \n", str, len(str))
	fmt.Printf("The fistr char is %c\n", ch)
	fmt.Printf("%s\n", "hello")
	fmt.Printf("%d\n", 1)
	fmt.Println("abc" + "def")
	for i := 0; i < len(str); i++ {
		c := str[i]
		fmt.Println(i, c)
	}
	str2 := []rune(str) // 转成unicode, 一个中文utf-8为3个长度
	for i := 0; i < len(str2); i++ {
		c := str2[i]
		fmt.Printf("%d %s\n", i, string(c))
	}
	for i, v := range str2 {
		fmt.Println(i, string(v))
	}
}
