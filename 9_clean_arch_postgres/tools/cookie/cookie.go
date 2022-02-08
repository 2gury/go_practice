package cookieHelper

import (
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/models"
	"net/http"
	"time"
)

func CreateCookie(sess *models.Session) *http.Cookie {
	return &http.Cookie{
		Name: consts.SessionName,
		Value: sess.Value,
		Path: "/",
		SameSite: http.SameSiteStrictMode,
		Expires: time.Now().Add(sess.TimeDuration),
		HttpOnly: true,
	}
}

func CreateExpiredCookie(cookie *http.Cookie) *http.Cookie {
	return &http.Cookie{
		Name: consts.SessionName,
		Value: cookie.Value,
		Path: "/",
		SameSite: http.SameSiteStrictMode,
		Expires: time.Now().AddDate(0, 0, -1),
		HttpOnly: true,
	}
}

func SetCookie(w http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(w, cookie)
}

func GetCookie(r *http.Request, cookieName string) (*http.Cookie, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func DeleteCookie(w http.ResponseWriter, r *http.Request, cookieName string) error {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return err
	}
	expiredCookie := CreateExpiredCookie(cookie)
	http.SetCookie(w, expiredCookie)
	return nil
}
