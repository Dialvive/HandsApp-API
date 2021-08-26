package controllers

import (
	"API/models"
	"API/security"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func init() {
	models.ConnectDatabase()
}

func TestCsrfMiddleware(t *testing.T) {
	r := gin.New()
	r.GET("/csrf", CsrfMiddleware)
	dummyJWT, _ := security.CreateJWT(models.User{})

	emptyRequest := httptest.NewRequest(http.MethodGet, "/csrf", nil)

	justHeader := emptyRequest.Clone(emptyRequest.Context())
	justHeader.Header.Set(HandsAppCsrfToken, dummyJWT.CsrfToken)

	justCookie := emptyRequest.Clone(emptyRequest.Context())
	justCookie.Header.Set("Cookie", fmt.Sprint(HandsAppSession, "=", dummyJWT.Token))

	headerAndCookie := justHeader.Clone(justHeader.Context())
	headerAndCookie.Header.Set("Cookie", fmt.Sprint(HandsAppSession, "=", dummyJWT.Token))

	testCases := []struct {
		Name     string
		Expected int
		*http.Request
	}{
		{"without authentication", http.StatusBadRequest, emptyRequest},
		{"header is not present", http.StatusBadRequest, justCookie},
		{"cookie is not present", http.StatusBadRequest, justHeader},
		{"should pass when cookie and header are set", http.StatusOK, headerAndCookie},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			w := httptest.NewRecorder()

			r.ServeHTTP(w, tc.Request)
			if w.Code != tc.Expected {
				t.Fatal("Expected: ", tc.Expected, "got: ", w.Code)
			}
		})
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
	r.DELETE("/delete/:ID", CsrfMiddleware, DeleteUser)

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
