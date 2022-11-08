package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

var usage_wanted = `MXlat

Usage:
  mxlat init [--apikey=<apikey>]
  mxlat languages
  mxlat translate (--all|--lang=<lang>...) <text>
  mxlat -h | --help
  mxlat --version

Options:
  -h --help						Show this screen.
  --version						Show version.
  -a --all						Take all languages availables for translation.
  -l LANG --lang=<LANG>		Select language.s in ISO 629-1 separate by comma.
  -k APIKEY --apikey=<APIKEY> 	Configure DeepL API key without interactionwith init mode.`

func main() {
	usage := `MXlat

Usage:
  mxlat -h | --help
  mxlat --version

Options:
  -h --help						Show this screen.
  --version						Show version.`

	arguments, _ := docopt.ParseArgs(usage, nil, "0.0.1")
	fmt.Println(arguments)
}
