package main

import (
	"GoRecipes/utilities"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
)

// cli-id : 366129222642-0f2m0k82m497rag5ls6jvftd9hj6is3b.apps.googleusercontent.com
// ssecret w2DgP9O9J7vgtVOjtME8bWqA
var googleOauthConfig *oauth2.Config

func setup_oauth(router *gin.Engine) {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/google_session",
		ClientID:     "366129222642-0f2m0k82m497rag5ls6jvftd9hj6is3b.apps.googleusercontent.com",
		ClientSecret: "w2DgP9O9J7vgtVOjtME8bWqA",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	router.GET("/login", func(c *gin.Context) {
		state := utilities.RandoString(50)
		sess := sessions.Default(c)
		sess.Set("state", state)
		sess.Save()
		url := googleOauthConfig.AuthCodeURL(state)
		c.Redirect(302, url)
	})

	router.GET("/logout", func(c *gin.Context) {
		if IsAuth(c) {
			sess := sessions.Default(c)
			sess.Clear()
		}
		c.Redirect(302, "/home")
	})
}

func setup_sessions(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))
	router.GET("google_session", func(c *gin.Context) {
		sess := sessions.Default(c)
		if sess.Get("state").(string) != c.Query("state") {
			log.Fatal("invalid oauth state")
		}
		token, err := googleOauthConfig.Exchange(oauth2.NoContext, c.Query("code"))
		if err != nil {
			log.Fatal(err)
		}
		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents : %s\n", contents)
		sess.Set("IsAuth", true)
		sess.Save()
		c.Redirect(302, "/home")
	})
}

func auth_routes(router *gin.Engine) {
	setup_sessions(router)
	setup_oauth(router)
}

func IsAuth(c *gin.Context) bool {
	session := sessions.Default(c)
	auth := session.Get("IsAuth")
	return auth != nil && auth.(bool) == true
}
