package main

import "fmt"

// 関数の二つ以上の引数が同じ型である場合は，最後の方を型を省略して記述出来る
func add(x, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
}
