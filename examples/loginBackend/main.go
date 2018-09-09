package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/services"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Store struct {
	Users map[string]User
}

var store = Store{Users: make(map[string]User)}

type User struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Company      string    `json:"company"`
	Nationality  string    `json:"nationality"`
	Expiry       time.Time `json:"-"`
	AccessToken  string    `json:"-"`
	RefreshToken string    `json:"-"`
	TokenType    string    `json:"-"`
}

type GenericHTTPResponse struct {
	StatusCode int
	Message    string
}

func getToken(ID string) oauth2.Token {
	user := store.Users[ID]
	token := oauth2.Token{}
	token.RefreshToken = user.RefreshToken
	token.AccessToken = user.AccessToken
	token.TokenType = user.TokenType
	token.Expiry = user.Expiry
	return token
}

func newUser(u User) error {
	if u.ID == "" {
		return errors.New("User is uninitialized")
	}
	attemptUser := store.Users[u.ID]
	if attemptUser.ID != "" {
		store.Users[u.ID] = u
		log.Println("User added")
		return nil
	}
	log.Println("User already exists")
	return errors.New("User already exists")
}

func login(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("auth.json")
	if err != nil {
		log.Fatal(err)
	}
	conf, err := google.ConfigFromJSON(data, "https://www.googleapis.com/auth/bigquery")
	if err != nil {
		log.Fatal(err)
	}
	url := conf.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func oauthCallback(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("auth.json")
	if err != nil {
		log.Fatal(err)
	}
	conf, err := google.ConfigFromJSON(data, "https://www.googleapis.com/auth/bigquery")
	if err != nil {
		log.Fatal(err)
	}

	type callbackRequest struct {
		Code string `json:"code"`
	}

	hehe, _ := ioutil.ReadAll(r.Body)
	lol := callbackRequest{}
	err = json.Unmarshal(hehe, &lol)
	if err != nil {
		log.Println("Error in trying to get oauth callback right. Will return immedietately with error")
		zzz, _ := json.Marshal(GenericHTTPResponse{StatusCode: http.StatusInternalServerError, Message: "Error in marshalling"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(zzz)
		return
	}

	// code, ok := r.URL.Query()["code"]
	code := lol.Code
	if code == "" {
		response := GenericHTTPResponse{StatusCode: http.StatusUnauthorized, Message: "User has not allowed app access"}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			log.Println("Error in Marshalling Unauthorized Error response")
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(responseJSON)
		return
	}

	ctx := context.Background()
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("ERROR ERROR. Cannot exchange token"))
		return
	}

	log.Println("Login Successful")

	user := User{}
	userID, err := uuid.NewV4()
	if err != nil {
		log.Println("Unable to generate proper used id in the request")
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(GenericHTTPResponse{StatusCode: http.StatusInternalServerError, Message: "Internal Error"})
		w.Write(response)
		return
	}
	user.ID = userID.String()
	user.RefreshToken = tok.RefreshToken
	user.AccessToken = tok.AccessToken
	user.TokenType = tok.TokenType
	user.Expiry = tok.Expiry

	log.Printf("%v\n", store)
	store.Users[user.ID] = user
	log.Printf("%v\n", store)

	jwtToken, err := services.NewToken(user.ID, 8000, "temp", "app")
	if err != nil {
		log.Fatalf("Error in generating the jwt token. %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		responseJSON, _ := json.Marshal(GenericHTTPResponse{StatusCode: http.StatusInternalServerError, Message: "JWT Token Error"})
		w.Write(responseJSON)
		return
	}

	log.Printf("JWT Token: %v\n", jwtToken)

	type successfulLogin struct {
		AuthToken  string `json:"auth_token"`
		StatusCode int    `json:"status_code"`
		Message    string `json:"msg"`
	}

	w.WriteHeader(http.StatusOK)
	responseJSON, _ := json.Marshal(successfulLogin{AuthToken: jwtToken, StatusCode: http.StatusOK, Message: "Successful Login"})
	w.Write(responseJSON)
	return
}

func logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func userDetails(w http.ResponseWriter, r *http.Request) {
	authenticationToken := r.Header.Get("Authentication")
	if authenticationToken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		response, _ := json.Marshal(GenericHTTPResponse{StatusCode: http.StatusUnauthorized, Message: "Unauthorized user"})
		w.Write(response)
		return
	}

	user_id, err := services.ExtractToken(authenticationToken, "temp")
	if err != nil {
		log.Printf("Error in extracting user id from authentication jwt token. Err: %v\n", err)
		w.WriteHeader(http.StatusUnauthorized)
		jsonResponse, _ := json.Marshal(GenericHTTPResponse{StatusCode: http.StatusUnauthorized, Message: "Cannot extract user id from field"})
		w.Write(jsonResponse)
		return
	}
	user := store.Users[user_id]
	if user.ID == "" {
		log.Printf("User id not available. Commit suicide here")
		w.WriteHeader(http.StatusUnauthorized)
		jsonResponse, _ := json.Marshal(GenericHTTPResponse{StatusCode: http.StatusUnauthorized, Message: "Cannot extract user id from field"})
		w.Write(jsonResponse)
	}

	data, err := ioutil.ReadFile("auth.json")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error"))
		return
	}
	conf, err := google.ConfigFromJSON(data, "https://www.googleapis.com/auth/bigquery")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error"))
		return
	}

	tok := oauth2.Token{}
	tok.RefreshToken = user.RefreshToken
	tok.AccessToken = user.AccessToken
	tok.Expiry = user.Expiry
	tok.TokenType = user.TokenType

	client := conf.Client(context.Background(), &tok)
	_ = client
	log.Println("Successful creation of client")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
	return
}

func printUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	jsonResponse, _ := json.Marshal(store)
	w.Write(jsonResponse)
	return
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	type VersionResponse struct {
		Version   string    `json:"version"`
		BuildDate time.Time `json:"build_date"`
	}
	w.WriteHeader(http.StatusOK)
	versionResponseJSON, _ := json.Marshal(VersionResponse{Version: "v2.0.13", BuildDate: time.Now()})
	w.Write(versionResponseJSON)
	return
}

func main() {
	log.Println("Begin")
	defer log.Println("Dieded")
	r := mux.NewRouter()
	r.HandleFunc("/version", getVersion)
	r.HandleFunc("/login", login)
	r.HandleFunc("/oauth_callback", oauthCallback)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/user", userDetails)
	r.HandleFunc("/dumpUsers", printUsers)
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	if err := http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		panic(err)
	}

}
