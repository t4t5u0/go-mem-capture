package main

import (
	"os/exec"
	"time"

	flags "github.com/jessevdk/go-flags"
)

func main() {
	// f := flag.
	flags.Parse(&Option)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if len(args) == 0 {

	// }
	filepath := "./img/"
	filename := time.Now().Format("2015061200171")
	filetype := ".png"
	// cmd := exec.Command("maim", filepath+filename+filetype)
	cmd := exec.Command("maim", "-i", "-B", filepath+filename+filetype)
	cmd.Start()
	cmd.Output()

}

type param struct {
	width  int
	height int
}

// ffmpeg -f x11grab -i $DISPLAY ./$(date +%s).png
