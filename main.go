package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	host := flag.String("h", ":3000", "The host to listen on")
	dir := flag.String("d", ".", "The directory to serve")

	flag.Parse()

	fmt.Println("Serving " + *dir + " at http://" + *host)

	fs := http.FileServer(http.Dir(*dir))
	err := http.ListenAndServe(*host, fs)

	if err != nil {
		fmt.Println(fmt.Errorf("failed to start server: %s", err.Error()))

		os.Exit(1)
	}

	fmt.Println("Serving " + *dir + " at http://" + *host)
}
