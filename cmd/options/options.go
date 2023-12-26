package options

import (
	"strconv"

	"github.com/spf13/viper"
)

var Version string = "v1.0.0"

var Letters string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var Numbers string = "0123456789"
var Symbols string = "~!@#$%^&*()"
var Length int = 8

type Options struct {
	All    bool
	Letter bool
	Number bool
	Symbol bool
	Length int
}

func NewOptions(args []string) (opts *Options) {
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
	if !opts.Letter && !opts.Number && !opts.Symbol {
		opts.Letter = true
		opts.Number = true
	}
	if len(args) == 1 {
		l, err := strconv.Atoi(args[0])
		if err != nil {
			l = Length
		}
		opts.Length = l
	} else {
		opts.Length = Length
	}
	if opts.Length < 1 {
		opts.Length = Length
	}

	return
}
