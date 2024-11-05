package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
	"github.com/shoshta73/homehub/internal/models/user"
)

var insertUserFunc = user.InsertUser

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/register", register)
	r.Post("/login", loginWithEmail)
	r.Post("/validate", validate)

	return r
}

func register(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	logger.Info("Got request to register user")

	logger.Info("Decoding body")
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("Failed to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logger.Info("Body decoded")

	logger.Info("Checking request body")

	if body.Username == "" {
		logger.Error("Username is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(body.Username) < 3 {
		logger.Error("Username is too short")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Email == "" {
		logger.Error("Email is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Password == "" {
		logger.Error("Password is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(body.Password) < 8 {
		logger.Error("Password is too short")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Info("Request body checked")

	logger.Info("Body parsed", "name", body.Name, "username", body.Username, "email", body.Email)

	logger.Info("Checking if username exists")
	exists, err := user.UsernameExists(body.Username)
	if err != nil {
		logger.Error("Failed to check if username exists", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		logger.Error("Username already exists")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logger.Info("Username does not exist")

	logger.Info("Checking if email exists")
	exists, err = user.EmailExists(body.Email)
	if err != nil {
		logger.Error("Failed to check if email exists", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		logger.Error("Email already exists")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logger.Info("Email does not exist")

	logger.Info("Registering user")
	var u *user.User

	if body.Name == "" {
		u = user.CreateUser(body.Username, body.Email, body.Password, map[string]string{})
	} else {
		u = user.CreateUser(body.Username, body.Email, body.Password, map[string]string{
			"name": body.Name,
		})
	}

	err = insertUserFunc(u)
	if err != nil {
		logger.Error("Failed to insert user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tkn := generateToken(u)
	cookie := getCookie(tkn)
	logger.Info(cookie)
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
}

func loginWithEmail(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	logger.Info("Got request to login with email")

	logger.Info("Decoding body")
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error("Failed to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logger.Info("Body decoded")

	logger.Info("Checking request body")

	if body.Email == "" {
		logger.Error("Email is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Password == "" {
		logger.Error("Password is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logger.Info("Request body checked")

	logger.Info("Getting user by email")
	u, err := user.GetUserByEmail(body.Email)
	if err != nil {
		logger.Error("Failed to get user by email", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Got User")

	if !u.VerifyPassword(body.Password) {
		logger.Error("Password does not match")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	logger.Info("User logged in")

	tkn := generateToken(u)
	cookie := getCookie(tkn)
	logger.Info(cookie)
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
}

func validate(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		logger.Info("Cookie: " + cookie.Name)
	}

	cookie, err := r.Cookie("token")
	if err != nil {
		logger.Error("Failed to get token from cookie: " + err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !validateToken(cookie.Value) {
		logger.Error("Token is invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
