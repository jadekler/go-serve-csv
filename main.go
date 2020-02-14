package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		if req.URL.RequestURI() == "/main.js" {
			outBytes, err := ioutil.ReadFile("main.js")
			if err != nil {
				log.Fatal(err)
			}
			if _, err := resp.Write(outBytes); err != nil {
				log.Fatal(err)
			}
			return
		}

		resp.Write([]byte(`<html>
<head></head>
<body>
Hello world
</body>
</html>
<script type="text/javascript" src="main.js"></script>`))
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}