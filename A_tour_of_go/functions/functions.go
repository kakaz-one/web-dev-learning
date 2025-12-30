package main

import "fmt"

// TODO: ここに calculateArea 関数を作成してください
// ヒント: func 関数名(引数1, 引数2 型) 戻り値の型 { ... }
func calculateArea(x, y int) int {
	return x * y
}

func main() {
	// テスト用の数値（縦10, 横20）
	h := 10
	w := 20

	// 関数を呼び出して結果を受け取る
	result := calculateArea(h, w)

	// 結果を表示
	fmt.Println("長方形の面積:", result)
}
