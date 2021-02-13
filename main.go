package main

import (
	"fmt"
	"go-shutter/lib"
	"log"
	"os/exec"

	flags "github.com/jessevdk/go-flags"
)

func main() {
	args, err := flags.Parse(&lib.Option)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(args)
	var (
		cmd   *exec.Cmd
		wFlag bool = contains(args, "--widnow") || contains(args, "-w")
		aFlag bool = contains(args, "--area") || contains(args, "-a")
		cFlag bool = contains(args, "--clipboard") || contains(args, "-c")
		dFlag bool = contains(args, "--delay") || contains(args, "-d")
		hFlag bool = contains(args, "--help") || contains(args, "-h")
	)

	fn := lib.Filename()
	if hFlag {
		return
	}
	if len(args) == 0 {
		// デフォルト。スクリーン全体を取得する
		cmd = lib.FullScreen(fn)

	}
	if wFlag && aFlag {
		// このパラメータ指定は不正
		fmt.Println(wFlag, aFlag)
		log.Fatal("this paramator is invalid")
	}

	if cFlag {
		if wFlag {
			// アクティブウィンドウをキャプチャ
			// xdotool getactivewindow で id を持ってくる必要あり
			cmd = lib.CopyWindowScreen()
		} else if aFlag {
			cmd = lib.CopyAreaScreen()
		}
		if dFlag {
			// n秒ディレイ
		}
		// クリップボードにコピー
		// xclip -se c -t image/png
		clip := exec.Command("xclip", "-se", "c", "-t", "image/png")
		clip.Stdin, _ = cmd.StdoutPipe()
		clip.Run()

	} else {
		if wFlag && aFlag {
			// このパラメータ指定は不正
		} else if wFlag {
			// アクティブウィンドウをキャプチャ
			// xdotool getactivewindow で id を持ってくる必要あり
			cmd = lib.WindowScreen(fn)
		} else if aFlag {
			cmd = lib.AreaScreen(fn)
		}
		if dFlag {
			// n秒ディレイ
		}
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
