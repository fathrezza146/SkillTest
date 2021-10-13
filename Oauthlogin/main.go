package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	authGoogle = &oauth2.Config{
		RedirectURL:  "http://localhost:4200/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

var (
	authGoogle  *oauth2.Config
	statestring = "pseudo-random"
)

func main() {
	http.HandleFunc("/", HandleMain)
	http.HandleFunc("/login", GoogleLogin)
	http.HandleFunc("/callback", Googlecallback)
	http.ListenAndServe(":4200", nil)
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	const htmlindeks = `<html><body><a href="/login">Log in with Google</a></body></html>`
	fmt.Fprintf(w, htmlindeks)
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	login := authGoogle.AuthCodeURL(statestring)
	http.Redirect(w, r, login, http.StatusTemporaryRedirect)
}

func Googlecallback(w http.ResponseWriter, r *http.Request) {
	called, err := infouser(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", called)
}

func infouser(state string, code string) ([]byte, error) {
	if state != statestring {
		return nil, fmt.Errorf("no oauthstate")
	}

	token, err := authGoogle.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("tukeran code error : %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("cannot get user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response bodt : %s", err.Error())
	}
	return contents, nil
}
