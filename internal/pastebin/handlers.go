package pastebin

import (
	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
	"github.com/shoshta73/homehub/internal/auth"
	"github.com/shoshta73/homehub/internal/models/paste"
	"github.com/shoshta73/homehub/internal/models/user"
	"net/http"
	"strconv"
	"time"
)

type Paste struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/create", create)
	r.Get("/created/all", created)
	r.Get("/created/count", createdCount)

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

func createdCount(w http.ResponseWriter, r *http.Request) {
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

	count, err := paste.CreatedCount(u.ID)
	if err != nil {
		logger.Error("Error getting created count", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatInt(count, 10)))
}

func created(w http.ResponseWriter, r *http.Request) {
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

	var resp struct {
		Pastes []Paste `json:"pastes"`
	}

	pastes, err := paste.GetCreatedPastes(u.ID)
	if err != nil {
		logger.Error("Error occurred getting pastes", err)

		resp.Pastes = []Paste{}

		b, err := json.Marshal(resp)
		if err != nil {
			logger.Error("Unable to marshal pastes", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
	}

	p := make([]Paste, len(pastes))
	for idx, pst := range pastes {
		p[idx] = Paste{
			ID:      pst.ID,
			Title:   pst.Title,
			Created: pst.CreatedAt,
			Updated: pst.UpdatedAt,
		}
	}

	resp.Pastes = p

	b, err := json.Marshal(resp)
	if err != nil {
		logger.Error("Unable to marshal pastes", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Info("Successful", "response", resp)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
