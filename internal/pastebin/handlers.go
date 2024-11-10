package pastebin

import (
	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
	"github.com/shoshta73/homehub/internal/auth"
	"github.com/shoshta73/homehub/internal/models/paste"
	"github.com/shoshta73/homehub/internal/models/user"
	"net/http"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/create", create)

	return r
}

func create(w http.ResponseWriter, r *http.Request) {
	logger.Info("Checking request")

	cookie, err := r.Cookie("token")
	if err != nil {
		logger.Error("Cookie not found")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	uc, err := auth.GetClaims(cookie.Value)
	if err != nil {
		logger.Error("Unable to extract claims", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u, err := user.GetUserById(uc.Id)
	if err != nil {
		logger.Error("Unable to retrieve user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !u.HasUserPermission() {
		logger.Warn("User does not have permission", "id", u.ID)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	logger.Info("Got request to create new paste")

	logger.Info("Decoding body")
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("Unable to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if paste.HasTitle(u.ID, body.Title) {
		logger.Warn("User tried to create paste with same title")
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("You already have paste with this title"))
		return
	}

	logger.Info("Creating new paste object")
	np, err := paste.Create(body.Title, body.Content)
	if err != nil {
		logger.Error("Unable to create paste", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Setting paste.OwnerID")
	np.SetOwnerId(u.ID)
	logger.Info("Successfully created paste", "title", np.Title, "content", np.Content)

	err = paste.Insert(np)
	if err != nil {
		logger.Error("Unable to insert paste", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(np.ID))
}
