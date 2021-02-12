package lib

import (
	"os/user"
	"time"
)

func Filename() string {
	usr, _ := user.Current()
	filepath := usr.HomeDir + "/Pictures/"
	filename := time.Now().Format("20060102150405")
	filetype := ".png"
	return filepath + filename + filetype
}
