package main

import (
	"fmt"
	"go-shutter/lib"
	"os/exec"

	flags "github.com/jessevdk/go-flags"
)

func main() {
	args, _ := flags.Parse(&lib.Option)
	var (
		cmd   *exec.Cmd
		wFlag bool = contains(args, "--widnow") || contains(args, "-w")
		aFlag bool = contains(args, "--area") || contains(args, "-a")
		cFlag bool = contains(args, "--clipboard") || contains(args, "-c")
		dFlag bool = contains(args, "--delay") || contains(args, "-d")
	)
	// filepath := "./img/"
	fn := lib.Filename()
	fmt.Println(args)
	if contains(args, "--help") || contains(args, "-h") {
		return
	}
	if len(args) == 0 {
		// デフォルト。スクリーン全体を取得する
		cmd = lib.FullScreen(fn)
	} else if wFlag && aFlag {
		// このパラメータ指定は不正
	} else if wFlag {
		// アクティブウィンドウをキャプチャ
		// xdotool getactivewindow で id を持ってくる必要あり
		cmd = lib.WindowScreen(fn)
	} else if aFlag {
		cmd = lib.AreaScreen(fn)
	}

	if cFlag {
		// クリップボードにコピー
	}
	// cmd := exec.Command("maim", filepath+filename+filetype)
	if dFlag {
		// n秒ディレイ
	}
	cmd.Start()
	cmd.Output()

}

// ffmpeg -f x11grab -i $DISPLAY ./$(date +%s).png

// 配列に要素が含まれているか判定する
func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
