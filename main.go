package main

import (
	"os/exec"
	"time"
)

func main() {
	filepath := "./img/"
	filename := time.Now().Format("2015061200171")
	filetype := ".png"
	cmd := exec.Command("maim", filepath+filename+filetype)
	cmd.Start()
	cmd.Output()

}

type param struct {
	width  int
	height int
}

// ffmpeg -f x11grab -i $DISPLAY ./$(date +%s).png
