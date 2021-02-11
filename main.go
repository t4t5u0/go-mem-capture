package main

import (
	"fmt"
	"go-shutter/lib"
	"os/exec"
	"time"

	flags "github.com/jessevdk/go-flags"
)

func main() {
	// f := flag.
	args, _ := flags.Parse(&lib.Option)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var (
		cmd   *exec.Cmd
		wFlag bool = contains(args, "--widnow") || contains(args, "-w")
		aFlag bool = contains(args, "--area") || contains(args, "-a")
		cFlag bool = contains(args, "--copy") || contains(args, "-c")
	)
	filepath := "./img/"
	filename := time.Now().Format("20060102150405")
	filetype := ".png"
	fmt.Println(args)
	if contains(args, "--help") || contains(args, "-h") {
		return
	}
	if len(args) == 0 {
		// デフォルト。スクリーン全体を取得する
		cmd = exec.Command("maim", filepath+filename+filetype)
	} else if wFlag && aFlag {
		// このパラメータ指定は不正
	} else if wFlag {
		fmt.Println()
		// アクティブウィンドウをキャプチャ
		// xdotoop getactivewindow で id を持ってくる必要あり
		cmd = exec.Command("maim", "-B", "-i", filepath+filename+filetype)
	} else if aFlag {

	}

	if cFlag {
		// クリップボードにコピー
	}
	// cmd := exec.Command("maim", filepath+filename+filetype)

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
