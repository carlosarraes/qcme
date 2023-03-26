package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/carlosarraes/qcmeback/models"
	"github.com/carlosarraes/qcmeback/utils"
)

const url = "https://qrcodeme.herokuapp.com/qrcodeme/"

func (a *App) GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(user.Name)

	buffer, err := utils.DrawQrCode(url, strings.ToLower(user.Name))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=qrcode.png")
	http.ServeContent(w, r, "qrcode.png", time.Time{}, bytes.NewReader(buffer.Bytes()))
}
