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

const url = "https://qcme.vercel.app/"

func (a *App) GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "invalid request")
		return
	}

	if user.Name == "" || user.Linkedin == "" || user.Github == "" {
		utils.WriteResponse(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := a.DB.CheckUser(user.Name); err == nil {
		utils.WriteResponse(w, http.StatusConflict, "user already exists")
		return
	}

	buffer, err := utils.DrawQrCode(url, strings.ToLower(user.Name))
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := a.DB.CreateUser(user); err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, "error creating user")
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=qrcode.png")
	http.ServeContent(w, r, "qrcode.png", time.Time{}, bytes.NewReader(buffer.Bytes()))
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	user, err := a.DB.GetUser(name)
	if err != nil || user.Name == "" {
		utils.WriteResponse(w, http.StatusNotFound, "user not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
