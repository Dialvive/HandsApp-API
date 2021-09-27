package security

import (
	"API/controllers"
	"API/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHashPassword(t *testing.T) {
	cases := []struct {
		name, password   string
		shouldThrowError bool
	}{
		{name: "should hash a normal password", password: "QMXzbVdX83uzUNNcQE9DQkFh", shouldThrowError: false},
		{name: "should hash an empty password", password: "", shouldThrowError: false},
		{name: "should hash a naughty password (face)", password: "・(￣∀￣)・:*:", shouldThrowError: false},
		{name: "should hash a naughty password (chars)", password: "`⁄€‹›ﬁﬂ‡°·‚—±", shouldThrowError: false},
		{name: "should hash a naughty password (js tag)", password: "␡", shouldThrowError: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			hashed, err := HashPassword(c.password)
			if err != nil || c.shouldThrowError {
				t.Fatalf("password: %s shouldThrow: %v; but error was: %v", c.password, c.shouldThrowError, err)
			}
			if c.password == hashed {
				t.Fatalf("with password: %v, got %v, but both should be different", c.password, hashed)
			}
		})
	}
}

func TestPasswordMatches(t *testing.T) {
	passwords := []struct {
		name, password string
	}{
		{name: "normal password", password: "aWjz5nB5"},
		{name: "password manager (firefox)", password: "?Zw!swRv_v<]8GP-"},
		{name: "password manager (safari)", password: "Qswmxs-wlrkxq-2rwlrt"},

		// special cases
		{name: "especial case (uft face)", password: "・(￣∀￣)・:*:"},
		{name: "especial case (japan chars)", password: "田中さんにあげて下さい"},
		{name: "especial case (emojis)", password: "🐵 🙈 🙉 🙊"},
		{name: "especial case (arab)", password: "estالصفحات التّحول"},
		{name: "especial case (especial chars)", password: "¡™£¢∞§¶•ªº–≠"},
	}

	for _, p := range passwords {
		t.Run(p.name, func(t *testing.T) {
			hashed, err := HashPassword(p.password)

			if err != nil {
				t.Fatalf("password: %v had an error hashing; got error: %v", p.password, err)
			}

			cases := []struct {
				name, passedPassword string
				shouldMatch          bool
			}{
				{name: "should match same password", passedPassword: p.password, shouldMatch: true},
				{name: "shouldn't match with trailing space", passedPassword: p.password + " ", shouldMatch: false},
				{name: "shouldn't match cut password", passedPassword: p.password[1:], shouldMatch: false},
			}

			for _, c := range cases {
				t.Run(c.name, func(t *testing.T) {
					matches := PasswordMatches(hashed, c.passedPassword)
					if c.shouldMatch != matches {
						t.Fatalf("PasswordMatches() = %v, want %v, (p: %v)", c.shouldMatch, matches, c.passedPassword)
					}
				})
			}
		})
	}
}

func TestCsrfMiddleware(t *testing.T) {
	r := gin.New()
	r.GET("/csrf", CsrfMiddleware)
	dummyJWT, _ := CreateJWT(models.User{})

	emptyRequest := httptest.NewRequest(http.MethodGet, "/csrf", nil)

	justHeader := emptyRequest.Clone(emptyRequest.Context())
	justHeader.Header.Set(controllers.HandsAppCsrfToken, dummyJWT.CsrfToken)

	justCookie := emptyRequest.Clone(emptyRequest.Context())
	justCookie.Header.Set("Cookie", fmt.Sprint(controllers.HandsAppSession, "=", dummyJWT.Token))

	headerAndCookie := justHeader.Clone(justHeader.Context())
	headerAndCookie.Header.Set("Cookie", fmt.Sprint(controllers.HandsAppSession, "=", dummyJWT.Token))

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

func TestApiKeyMiddleware(t *testing.T) {
	r := gin.New()
	r.GET("", ApiKeyMiddleware)

	emptyRequest := httptest.NewRequest(http.MethodGet, "/", nil)

	wrongApiKey := emptyRequest.Clone(emptyRequest.Context())
	wrongApiKey.Header.Set("x-api-key", "not s3cret")

	correctApiKey := emptyRequest.Clone(emptyRequest.Context())
	correctApiKey.Header.Set("x-api-key", KEY)

	testCases := []struct {
		Name     string
		Expected int
		*http.Request
	}{
		{"without api key", http.StatusNotFound, emptyRequest},
		{"wrong api key", http.StatusNotFound, wrongApiKey},
		{"correct api key", http.StatusOK, correctApiKey},
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
