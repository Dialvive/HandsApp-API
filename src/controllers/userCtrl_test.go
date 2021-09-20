package controllers

import (
	"API/models"
	"API/security"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func init() {
	models.ConnectDatabase()
}

func TestPatchUser(t *testing.T) {
	name := fmt.Sprint(t.Name(), "@", time.Now().Unix())
	userJWT, err := userService.SignWithHandsApp(models.User{UserName: name, Mail: name, Password: name, LocaleID: 1})

	if err != nil {
		t.Fatal("error creating the user:", err.Error())
	}

	claims := &models.UserClaim{}
	_ = security.ParseJWT(userJWT.Token, claims)

	r := gin.New()
	r.PATCH("/:ID", security.CsrfMiddleware, PatchUser)

	req := newRequestUsingCsrfToken(
		http.MethodPatch,
		fmt.Sprint("/", claims.Subject),
		strings.NewReader(`{
    "first_name": "Cool",
    "last_name": "Cool Last",
    "user_name": "cool-username",
    "mail": "cool@example.com",
    "password": "secr3t",
    "biography": "cool biu",
    "mailing": "notification",
    "locale_ID": 1,
    "google_sub": "1111",
    "facebook_sub": "2222",
    "apple_sub": "3333",
	"picture": "https://www.example.com"
}`),
		userJWT,
	)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("Expected: ", http.StatusOK, "got: ", w.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	name := fmt.Sprint(t.Name(), "@", time.Now().Unix())
	userJWT, err := userService.SignWithHandsApp(models.User{UserName: name, Mail: name, Password: name, LocaleID: 1})
	if err != nil {
		t.Fatal("error creating user: ", err.Error())
	}

	claims := &models.UserClaim{}
	_ = security.ParseJWT(userJWT.Token, claims)

	r := gin.New()
	r.DELETE("/delete/:ID", security.CsrfMiddleware, DeleteUser)

	req := newRequestUsingCsrfToken(
		http.MethodDelete,
		fmt.Sprint("/delete/", claims.Subject),
		nil,
		userJWT,
	)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
}

func newRequestUsingCsrfToken(method, path string, body io.Reader, userJWT models.HandsAppJWT) *http.Request {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set(HandsAppCsrfToken, userJWT.CsrfToken)
	req.Header.Set("Cookie", fmt.Sprint(HandsAppSession, "=", userJWT.Token))
	return req
}
