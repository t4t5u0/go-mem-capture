package lib

// Option this is option
var Option struct {
	Help      bool `short:"h" long:"help" discription:"this is help"`
	Window    bool `short:"w" long:"window" discription:"capture window"`
	Area      bool `short:"a" long:"area" discription:"capture area"`
	Clipboard bool `short:"c" long:"clipboard" discription:"copy to clipboard"`
	Delay     bool `short:"d" long:"delay" discription:"delay seconds"`
}
