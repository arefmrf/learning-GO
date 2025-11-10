package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prj/internal/modules/user/service"
)

type AuthController struct {
	jwtService service.JWTService
}

func NewAuthController(jwt service.JWTService) *AuthController {
	return &AuthController{jwtService: jwt}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	token, err := c.jwtService.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		fmt.Println("Failed to encode token", err)
		return
	}
}
