package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/mitranim/srv"
)

const help = `srv starts a local HTTP server, serving files from the given
directory (default .), approximary like it would be served by
default Nginx, GitHub Pages, Netlify, and so on. Uses a random
port by default.

Usage:

	srv
	srv -d <dir>
	srv -p <port>

Settings:

`

var conf = struct {
	Port int
	Dir  string
}{
	Port: 0,
	Dir:  ".",
}

var logger = log.New(os.Stderr, "", 0)

func main() {
	flag.IntVar(&conf.Port, "p", conf.Port, "port")
	flag.StringVar(&conf.Dir, "d", conf.Dir, "dir")

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), help)
		flag.PrintDefaults()
	}

	flag.Parse()
	args()

	crit(serve())

}

func serve() error {
	// This allows us to find the OS-provided port.
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.Port))
	if err != nil {
		return err
	}

	port := listener.Addr().(*net.TCPAddr).Port
	dir := conf.Dir
	server := &http.Server{Handler: srv.FileServer(dir)}

	logger.Printf("[srv] serving %q on %v\n", dir, fmt.Sprintf("http://localhost:%v", port))
	return server.Serve(listener)
}

func args() {
	args := flag.Args()
	if len(args) == 0 {
		return
	}

	if args[0] == "help" {
		flag.Usage()
		os.Exit(0)
	}

	crit(fmt.Errorf(`[srv] unexpected arguments %q`, args))
}

func crit(err error) {
	if err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "%+v\n", err)
		os.Exit(1)
	}
}
