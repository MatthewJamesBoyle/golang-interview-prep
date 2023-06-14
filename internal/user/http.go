package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Svc service
}

type UserCreationSuccessResponse struct {
  Code int `json:"code"`
  Message string `json:"message"`
  UserId string `json:"userId"`
}

type UserCreationFailedResponse struct {
  Code int `json:"code"`
  Message string `json:"message"`
  Reason string `json:"reason"`
}

func (h Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
    response := UserCreationFailedResponse{
      Code: http.StatusBadRequest,
      Message: "Failed to create user",
      Reason: "Invalid request body",
    }
		w.WriteHeader(response.Code)
    w.Header().Add("Content-Type", "application/json")
    bytes, _ := json.Marshal(response)
		w.Write(bytes)
		return
	}
	// Call the AddUser function
	message, err := h.Svc.AddUser(u)
	if err != nil {
    response := UserCreationFailedResponse{
      Code: http.StatusInternalServerError,
      Message: "Failed to create user",
      Reason: "Failed to add user", // Thank you, captain obvious! :)
    }
    w.WriteHeader(response.Code)
    w.Header().Add("Content-Type", "application/json")
    bytes, _ := json.Marshal(response)
		w.Write(bytes)
	}

	// Return a success response
  response := UserCreationSuccessResponse{
    Code: http.StatusCreated, 
    Message: "New user has been created", 
    UserId: message,
  }
	w.WriteHeader(response.Code)
  w.Header().Add("Content-Type", "application/json")
  bytes, _ := json.Marshal(response)
	w.Write(bytes)
}
