package controllers

import (
	"bytes"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/skip2/go-qrcode"
)

var url = "https://qrcodeme.herokuapp.com/qrcodeme/"

func (a *App) GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(chi.URLParam(r, "name"))

	png, err := qrcode.Encode(url+name, qrcode.Medium, 256)
	if err != nil {
		log.Fatalf("Failed to generate QR Code: %v", err)
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=qrcode.png")

	buf := bytes.NewReader(png)
	http.ServeContent(w, r, "qrcode.png", time.Time{}, buf)
}
