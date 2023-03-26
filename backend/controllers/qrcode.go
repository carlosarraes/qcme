package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/carlosarraes/qcmeback/models"
	"github.com/carlosarraes/qcmeback/utils"
	"github.com/go-chi/chi/v5"
)

const url = "https://qrcodeme.herokuapp.com/qrcodeme/"

func (a *App) GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checkUser, _ := a.DB.GetUser(user.Name)
	if checkUser.Name == user.Name {
		http.Error(w, "user already exists", http.StatusConflict)
		return
	}

	buffer, err := utils.DrawQrCode(url, strings.ToLower(user.Name))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := a.DB.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=qrcode.png")
	http.ServeContent(w, r, "qrcode.png", time.Time{}, bytes.NewReader(buffer.Bytes()))
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	user, err := a.DB.GetUser(name)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
