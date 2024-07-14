package model

import (
	"net/http"
	"strconv"
	// "strings"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
    dataParam := r.URL.Query().Get("data")
    exportFormat := r.URL.Query().Get("exportFormat")

    var contentType string
    var fileExtension string
    // if exportFormat == "html" {
    //     // Convert to Windows-style newlines if necessary
    //     dataParam = strings.ReplaceAll(dataParam, "\n", "\r\n")
    // }
    switch exportFormat {
    case "csv":
        contentType = "text/csv"
        fileExtension = "csv"
    default:
        contentType = "text/plain; charset=utf-8"
        fileExtension = "txt"
    }

    // Update content length after potential modifications
    contentLength := len(dataParam)

    w.Header().Set("Content-Disposition", "attachment; filename=ascii_output."+fileExtension)
    w.Header().Set("Content-Type", contentType)
    w.Header().Set("Content-Length", strconv.Itoa(contentLength))

    // Write the modified data to the response
    _, err := w.Write([]byte(dataParam))
    if err != nil {
        http.Error(w, "Failed to write data to response", http.StatusInternalServerError)
        return
    }
}