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

// CopyAreaScreen 画面の一部分を取得
func CopyAreaScreen() *exec.Cmd {
	cmd := exec.Command("maim", "-B", "-s")
	return cmd
}

// CopyFullScreen default
func CopyFullScreen() *exec.Cmd {
	cmd := exec.Command("maim")
	return cmd
}

// CopyWindowScreen アクティブウィンドウを取得
func CopyWindowScreen() *exec.Cmd {
	id, _ := exec.Command("xdotool", "getactivewindow").Output()
	strid := string(id)
	fmt.Println(strid)
	cmd := exec.Command("maim", "-B", "-i", strid)
	return cmd
}
