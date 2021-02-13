package lib

// Option this is option
var Option struct {
	help      bool `short:"h" long:"help" discription:"this is help"`
	window    bool `short:"w" long:"window" discription:"capture window"`
	area      bool `short:"-a" long:"--area" discription:"capture area"`
	clipboard bool `short:"c" long:"clipboard" discription:"copy to clipboard"`
	delay     bool `short:"d" long:"delay" discription:"delay seconds"`
}
