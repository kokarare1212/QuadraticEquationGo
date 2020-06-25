package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4{
		fmt.Println(os.Args[0]+": エラー: 不正な引数です")
		helpMsg()
		os.Exit(-1)
	}
	if !checkVal(os.Args[1], os.Args[2], os.Args[3]){
		fmt.Println(os.Args[0]+": エラー: 不正な引数です")
		helpMsg()
		os.Exit(-1)
	}
	a, _ := strconv.ParseFloat(os.Args[1], 64)
	b, _ := strconv.ParseFloat(os.Args[2], 64)
	c, _ := strconv.ParseFloat(os.Args[3], 64)
	ir := (b * b - 4 * a * c)
	irh := math.Sqrt(float64(ir))
	if math.IsNaN(irh){
		fmt.Println(os.Args[0]+": エラー: 不正な引数です")
		helpMsg()
		os.Exit(-1)
	}
	ca1 := (-b + irh) / (2 * a)
	ca2 := (-b - irh) / (2 * a)
	fmt.Println(os.Args[1]+"x²+"+os.Args[2]+"x+"+os.Args[3]+"=0")
	fmt.Println("x="+strconv.FormatFloat(ca1, 'f', 0, 64)+","+strconv.FormatFloat(ca2, 'f', 0, 64))
}
func helpMsg() {
	fmt.Println("使用方法\n"+os.Args[0]+" (x²の係数) (xの係数) (x以外の係数)")
}
func checkVal(a, b, c string) bool {
	a1, a2 := strconv.ParseFloat(a, 64)
	b1, b2 := strconv.ParseFloat(b, 64)
	c1, c2 := strconv.ParseFloat(c, 64)
	if a2 == nil && b2 == nil && c2 == nil{
		if a1 != 0 && b1 != 0 && c1 != 0{
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}