package handle

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/atompi/genpass/cmd/options"
)

func Handle(opts *options.Options) {
	password := generatePassword(opts.Length, opts.Letter, opts.Number, opts.Symbol)
	fmt.Fprintln(os.Stdout, password)
}

func generatePassword(length int, withLetters, withNumbers, withSymbols bool) string {
	b := make([]byte, length)
	s := ""
	rand.New(rand.NewSource(time.Now().UnixNano()))

	if withLetters {
		s = s + options.Letters
	}
	if withNumbers {
		s = s + options.Numbers
	}
	if withSymbols {
		s = s + options.Symbols
	}

	for i := range b {
		n := rand.Intn(len(s))
		b[i] = s[n]
	}

	return string(b)
}
