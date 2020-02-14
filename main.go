package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		switch req.URL.RequestURI() {
		case "/main.js":
			outBytes, err := ioutil.ReadFile("main.js")
			if err != nil {
				log.Fatal(err)
			}
			if _, err := resp.Write(outBytes); err != nil {
				log.Fatal(err)
			}
			return
		case "/foo.csv":
			if _, err := resp.Write([]byte(`1,2,3
3,4,5
6,7,8`)); err != nil {
				log.Fatal(err)
			}
			return
		}

		resp.Write([]byte(`<html>
<head></head>
<body>
Hello world

<button id="clickme">click me</button>
</body>
</html>
<script type="text/javascript" src="main.js"></script>`))
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
