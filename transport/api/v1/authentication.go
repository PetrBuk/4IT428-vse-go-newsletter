package v1

import (
	"encoding/json"
	"net/http"
	"vse-go-newsletter-api/transport/api/v1/model"
	"vse-go-newsletter-api/transport/util"

	types "github.com/supabase-community/gotrue-go/types"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var requestData model.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.service.Login(r.Context(), requestData.Email, requestData.Password)

	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, response)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var requestData model.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.service.Register(r.Context(), requestData.Email, requestData.Password)

	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, response)
}

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var requestData model.ChangePasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()

	userData := ctx.Value("user").(map[string]interface{})

	result, err := h.service.ChangePassword(r.Context(), userData["token"].(string), userData["email"].(string), requestData.OldPassword, requestData.NewPassword)

	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, result)
}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	var requestData model.VerifyRequest

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	result, err := h.service.Verify(r.Context(), types.VerificationTypeSignup, requestData.Email, requestData.OTPToken)

	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, result)
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var requestData model.RefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	result, err := h.service.RefreshToken(r.Context(), requestData.RefreshToken)

	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, result)
}
