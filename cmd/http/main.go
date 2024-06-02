package main

import (
	"net/http"
	"ryosantouchh/gym-management-backend/internal/database"
	"ryosantouchh/gym-management-backend/internal/server"

	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading env file")
	}

	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}

	// url := oauthConfig.AuthCodeURL("state")
	// fmt.Printf("url dialog: %v \n", url)

	// tok, err := oauthConfig.Exchange(context.Background(), "authorization-code")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// oauthConfig.Client(context.Background(), tok)

	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		authURL := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(w, r, authURL, http.StatusFound)
		// fmt.Fprintf(w, "<a href='%s'>Login with Google</a>", authURL)
	})
	// Handle callback from authorization server
	http.HandleFunc("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// code := r.URL
	})
	// http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
	// 	// Extract and validate authorization code
	// 	code := r.URL.Query().Get("code")
	// 	if code == "" {
	// 		http.Error(w, "Missing authorization code", http.StatusBadRequest)
	// 		return
	// 	}
	//
	// 	// Exchange authorization code for tokens
	// 	ctx := context.Background()
	// 	tok, err := oauthConfig.Exchange(ctx, code)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	//
	// 	// Use the access token to access protected resources (replace with your logic)
	// 	userInfoURL := "https://www.googleapis.com/oauth2/v3/userinfo"
	// 	client := http.Client{Transport: &oauth2.Transport{Source: tok}}
	// 	userInfoResp, err := client.Get(userInfoURL)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	defer userInfoResp.Body.Close()
	//
	// 	var userInfo map[string]interface{}
	// 	err = json.NewDecoder(userInfoResp.Body).Decode(&userInfo)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	//
	// 	// Return user information to the frontend (consider redirecting instead)
	// 	fmt.Fprintf(w, "Hello, %s!", userInfo["name"].(string))
	// })

	// fmt.Println("Server listening on port 8080")
	// http.ListenAndServe(":8080", nil)

	db := database.Connect()

	app := server.NewServerApp(db)
	app.Start()
}
