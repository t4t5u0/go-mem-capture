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

	filepath := "./img/"
	filename := time.Now().Format("20060102150405")
	filetype := ".png"
	var cmd *exec.Cmd
	fmt.Println(args)
	if contains(args, "--help") || contains(args, "-h") {
		return
	}
	if len(args) == 0 {
		// デフォルト。スクリーン全体を取得する
		cmd = exec.Command("maim", filepath+filename+filetype)
	} else if contains(args, "--window") || contains(args, "-w") {
		fmt.Println()
		// アクティブウィンドウをキャプチャ
		cmd = exec.Command("maim", "-i", "-B", filepath+filename+filetype)
	}
	// cmd := exec.Command("maim", filepath+filename+filetype)

	cmd.Start()
	cmd.Output()

}

type param struct {
	width  int
	height int
}

// ffmpeg -f x11grab -i $DISPLAY ./$(date +%s).png
func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
