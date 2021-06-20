package controllers

import (
	"API/models"
	"API/security"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	models.ConnectDatabase()
}

func TestCsrfMiddleware(t *testing.T) {
	r := gin.New()
	r.GET("/csrf", CsrfMiddleware)
	jwtMocked, _ := security.CreateJWT(models.User{})
	req := httptest.NewRequest("GET", "/csrf", nil)
	w := httptest.NewRecorder()

	t.Run("header is not present", func(t *testing.T) {
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Fail()
		}
	})

	req.Header.Set(HandsAppCsrfToken, jwtMocked.CsrfToken)
	t.Run("cookie is not present", func(t *testing.T) {
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Fail()
		}
	})

	req.Header.Set("Cookie", fmt.Sprint(HandsAppSession, "=", jwtMocked.Token))
	t.Run("header is not present", func(t *testing.T) {
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fail()
		}
	})

}

func TestDeleteUser(t *testing.T) {
	name := fmt.Sprint(t.Name(), "@", time.Now().Unix())
	userJWT, _ := userService.SignWithHandsApp(models.User{UserName: name, Mail: name, Password: name, LocaleID: 1})
	claims := &models.UserClaim{}
	_ = security.ParseJWT(userJWT.Token, claims)

	r := gin.New()
	r.DELETE("/delete/:ID", CsrfMiddleware, DeleteUser)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%s", claims.Subject), nil)
	req.Header.Set(HandsAppCsrfToken, userJWT.CsrfToken)
	req.Header.Set("Cookie", fmt.Sprint(HandsAppSession, "=", userJWT.Token))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
}
