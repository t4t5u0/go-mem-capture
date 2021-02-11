// package main

// import (
// 	"log"

// 	"github.com/robotn/xgb/xproto"
// 	"github.com/robotn/xgbutil"
// 	"github.com/robotn/xgbutil/xevent"
// 	"github.com/robotn/xgbutil/xgraphics"
// )

// func main() {
// 	X, err := xgbutil.NewConnDisplay(":0")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ximg, err := xgraphics.NewDrawable(X, xproto.Drawable(X.RootWin()))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Shows the screenshot in a window.
// 	ximg.XShowExtra("Screenshot", true)

// 	// If you'd like to save it as a png, use:
// 	err = ximg.SavePng("screenshot.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// fmt.Println(X.TimeGet())
// 	xevent.Main(X)
// 	// fmt.Println(X.TimeGet())
// }

package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func main() {
	X, err := xgb.NewConn()
	X.Sync()

	if err != nil {
		log.Fatal(err)
	}
	defer X.Close()
	currentDesktop := "_NET_CURRENT_DESKTOP"
	n := getProperty(currentDesktop, X)
	fmt.Println(currentDesktop, n)
	img, err := CaptureScreen(X)
	if err != nil {
		panic(err)
	}

	fileName := time.Now().Format("20060102150405")
	f, err := os.Create("./img/" + fileName + ".png")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

}

func ScreenRect(X *xgb.Conn) (image.Rectangle, error) {
	// c, err := xgb.NewConn()
	// c.Sync()
	// if err != nil {
	// 	return image.Rectangle{}, err
	// }
	// defer c.Close()
	screen := xproto.Setup(X).DefaultScreen(X)
	x := screen.WidthInPixels
	y := screen.HeightInPixels

	return image.Rect(0, 0, int(x), int(y)), nil
}

func CaptureScreen(X *xgb.Conn) (*image.RGBA, error) {
	r, e := ScreenRect(X)
	if e != nil {
		return nil, e
	}
	return CaptureRect(X, r)
}

func CaptureRect(X *xgb.Conn, rect image.Rectangle) (*image.RGBA, error) {
	// c, err := xgb.NewConn()
	// c.Sync()
	// if err != nil {
	// 	return nil, err
	// }
	// defer c.Close()

	screen := xproto.Setup(X).DefaultScreen(X)
	x, y := rect.Dx(), rect.Dy()
	xImg, err := xproto.GetImage(X, xproto.ImageFormatZPixmap, xproto.Drawable(screen.Root), int16(rect.Min.X), int16(rect.Min.Y), uint16(x), uint16(y), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{data, 4 * x, image.Rect(0, 0, x, y)}
	return img, nil
}

func getProperty(prop string, X *xgb.Conn) xproto.Window {
	setup := xproto.Setup(X)
	root := setup.DefaultScreen(X).Root
	// numDesktop := prop
	atomName, err := xproto.InternAtom(X, true, uint16(len(prop)), prop).Reply()
	if err != nil {
		log.Fatal(err)
	}
	reply, err := xproto.GetProperty(X, false, root, atomName.Atom,
		xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		log.Fatal(err)
	}
	result := xproto.Window(xgb.Get32(reply.Value))
	// fmt.Printf("Number of Desktop: %X\n", number)
	return result
}
