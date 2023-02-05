package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	// 向上滚动：3行
	robotgo.ScrollMouse(3, `up`)
	//// 向下滚动：2行
	//robotgo.ScrollMouse(2, `down`)
	//
	//// 按下鼠标左键
	//// 第1个参数：left(左键) / center(中键，即：滚轮) / right(右键)
	//// 第2个参数：是否双击
	//robotgo.MouseClick(`left`, false)
	//
	//// 按住鼠标左键
	//robotgo.MouseToggle(`down`, `left`)
	//// 解除按住鼠标左键
	//robotgo.MouseToggle(`up`, `left`)
	//
	//robotgo.MoveSmooth(500, 300)
	//
	//robotgo.Click(`left`, false)
	//
	//time.Sleep(1 * time.Second)
	//
	//robotgo.MoveSmooth(600, 350)
	//
	//robotgo.Click(`left`, false)
	//
	//time.Sleep(1 * time.Second)
	//
	//robotgo.MoveSmooth(700, 500)
	//
	//time.Sleep(1 * time.Second)

	robotgo.Click(`left`, false)

	robotgo.ShowAlert("dddd", "sdfsdf")

	robotgo.TypeStr("mingyuan")



}