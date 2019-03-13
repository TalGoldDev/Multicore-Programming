// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		//Exercise 1.2 - start
		//checking if url contains http://
		if(!strings.Contains(url,"http://")){
			s := []string{"http://", url}
			url = strings.Join(s, "")
		}
		//Exercise 1.2 - end

		//Exercise 1.3 - start
		resp, err := http.Get(url)
		fmt.Println("status code: ",resp.StatusCode)
		//Exercise 1.3 - end


		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//Exercise 1.1 - start
		data := make([]byte, 4096) // our buffer

		if _, err  := io.CopyBuffer(os.Stdout, resp.Body, data); err != nil {
			// checking if we reached EOF.
			fmt.Println(err)
			os.Exit(1)
		}

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", os.Stdout)
		//Exercise 1.1 - end
	}
}
