package main

import (
	"ascii-art-web/Controller"
	"ascii-art-web/model"
	"fmt"
	"log"
	"net/http"
)
func main() {
	fmt.Println("Server running at http://localhost:3001")
	http.Handle("/View/static/", http.StripPrefix("/View/static/", http.FileServer(http.Dir("./View/static/"))))
	http.Handle("/View/image/", http.StripPrefix("/View/image/", http.FileServer(http.Dir("./View/image/"))))
	http.HandleFunc("/", Controller.RenderHomePage)
	http.HandleFunc("/asciiArt", Controller.RenderAsciiArtPage)
	http.HandleFunc("/download",model.DownloadHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}