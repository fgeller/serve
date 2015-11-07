package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
)

func openURL(url string) {
	switch runtime.GOOS {
	case "darwin":
		log.Println("Opening...")
		exec.Command("open", url).Run()
	case "linux":
		log.Println("Opening...")
		exec.Command("xdg-open", url).Run()
	case "windows":
		log.Println("Opening...")
		exec.Command("cmd.exe", "/C", "start", url).Run()
	default:
		log.Printf("Don't know how to open URL on %v.\n", runtime.GOOS)
	}
}

func serve(ln net.Listener) {
	h := http.FileServer(http.Dir("."))
	if err := http.Serve(ln, h); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

func listener(i string, port string) net.Listener {
	ln, err := net.Listen("tcp", i+":"+port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	return ln
}

func open(ln net.Listener, i string) {
	_, p, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		log.Fatal("Failed to identify host and port.", err)
	}
	url := fmt.Sprintf("http://%v:%v", i, p)
	log.Printf("Serving at %v\n", url)
	openURL(url)
}

func main() {
	flag.Parse()
	port := flag.Arg(0)
	if port == "" {
		port = "0"
	}
	i := "0.0.0.0"
	ln := listener(i, port)
	go open(ln, i)
	serve(ln)
}
