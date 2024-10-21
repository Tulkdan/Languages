package main

import "example/web-server/src"

func main() {
	src.NewHTTPServer(":8000").ListenAndServe()
}
