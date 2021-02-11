// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/jezek/xgb"
// 	"github.com/jezek/xgb/xproto"
// )

// func main() {
// 	X, err := xgb.NewConn()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer X.Close()

// 	// setup := xproto.Setup(X)
// 	// root := setup.DefaultScreen(X).Root

// 	numDesktop := "_NET_NUMBER_OF_DESKTOPS"
// 	currentDesktop := "_NET_CURRENT_DESKTOP"
// 	number := getProperty(numDesktop, X)
// 	current := getProperty(currentDesktop, X)
// 	fmt.Printf("Number of Desktop: %X\n", number)
// 	fmt.Printf("Current Desktop %X\n", current)
// }

// func getProperty(prop string, X *xgb.Conn) xproto.Window {
// 	setup := xproto.Setup(X)
// 	root := setup.DefaultScreen(X).Root
// 	// numDesktop := prop
// 	atomName, err := xproto.InternAtom(X, true, uint16(len(prop)), prop).Reply()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	reply, err := xproto.GetProperty(X, false, root, atomName.Atom,
// 		xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	result := xproto.Window(xgb.Get32(reply.Value))
// 	// fmt.Printf("Number of Desktop: %X\n", number)
// 	return result
// }
