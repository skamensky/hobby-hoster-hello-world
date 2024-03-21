package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<title>Hello World</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body { font-family: Arial, sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; }
		.container { text-align: center; }
	</style>
</head>
<body>
	<div class="container">
		<h1>Hello World</h1>
		<p>You are visiting from: %s</p>
	</div>
</body>
</html>`, r.Host)
}

func main() {
	fmt.Println("Started hello world server!")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":80", nil)
}
