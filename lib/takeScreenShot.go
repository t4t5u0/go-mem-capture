package lib

import (
	"fmt"
	"os/exec"
)

// FullScreen default
func FullScreen(filename string) *exec.Cmd {
	cmd := exec.Command("maim", filename)
	return cmd
}

// WindowScreen アクティブウィンドウを取得
func WindowScreen(filename string) *exec.Cmd {
	id, _ := exec.Command("xdotool", "getactivewindow").Output()
	strid := string(id)
	fmt.Println(strid)
	cmd := exec.Command("maim", "-B", "-i", strid, filename)
	return cmd
}

// AreaScreen 画面の一部分を取得
func AreaScreen(filename string) *exec.Cmd {
	cmd := exec.Command("maim", "-B", "-s", filename)
	return cmd
}
