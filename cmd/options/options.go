package options

import (
	"github.com/spf13/viper"
)

var Version string = "v1.0.0"

var Letters string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var Numbers string = "0123456789"
var Symbols string = "~!@#$%^&*()"

type Options struct {
	All    bool
	Letter bool
	Number bool
	Symbol bool
	Length int
}

func NewOptions() (opts *Options) {
	opts = &Options{
		All:    viper.GetBool("all"),
		Letter: viper.GetBool("letter"),
		Number: viper.GetBool("number"),
		Symbol: viper.GetBool("symbol"),
		Length: viper.GetInt("length"),
	}
	if opts.All {
		opts.Letter = true
		opts.Number = true
		opts.Symbol = true
	}
	if opts.Length < 1 {
		opts.Length = 8
	}
	return
}
