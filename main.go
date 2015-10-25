package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
)

func open(url string) {
	switch runtime.GOOS {
	case "darwin":
		log.Println("Opening...")
		exec.Command("open", url).Run()
	case "linux":
		log.Println("Opening...")
		exec.Command("xdg-open", url).Run()
	}
}

func serve(ln net.Listener) {
	h := http.FileServer(http.Dir("."))
	if err := http.Serve(ln, h); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

func listener(i string) net.Listener {
	ln, err := net.Listen("tcp", i+":0")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	return ln
}

func main() {
	i := "0.0.0.0"
	ln := listener(i)
	go func() {
		_, p, err := net.SplitHostPort(ln.Addr().String())
		if err != nil {
			log.Fatal("Failed to identify host and port.", err)
		}
		url := fmt.Sprintf("http://%v:%v", i, p)
		log.Printf("Serving at %v\n", url)
		open(url)
	}()
	serve(ln)
}
