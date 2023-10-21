package main

import (
	"os"

	"github.com/uri/u"
)

func main() {
	u.PrettyJSON(os.Stdin, os.Stdout)
}
