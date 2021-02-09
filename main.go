package main

import (
	"log"

	"github.com/robotn/xgb/xproto"
	"github.com/robotn/xgbutil"
	"github.com/robotn/xgbutil/xevent"
	"github.com/robotn/xgbutil/xgraphics"
)

func main() {
	X, err := xgbutil.NewConnDisplay(":0")
	if err != nil {
		log.Fatal(err)
	}
	ximg, err := xgraphics.NewDrawable(X, xproto.Drawable(X.RootWin()))
	if err != nil {
		log.Fatal(err)
	}
	// Shows the screenshot in a window.
	ximg.XShowExtra("Screenshot", !true)

	// If you'd like to save it as a png, use:
	err = ximg.SavePng("screenshot.png")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(X.TimeGet())
	xevent.Main(X)
	// fmt.Println(X.TimeGet())
}
