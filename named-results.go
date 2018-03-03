package main

import "fmt"

// 戻り値となる変数に名前をつけることが出来る
// 名前を付けた戻り値の変数を使うとreturnに何も書かずに戻せる
// これをnaked returnといいます
func split(sum int)(x, y int){
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(split(17))
}
