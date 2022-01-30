package delivery

import (
	"github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/internal/user"
	"net/http"
)

type UserHandler struct {
	userUse user.UserUsecase
}

func NewUserHandler(use *user.UserUsecase) *UserHandler {
	return &UserHandler{
		userUse: use,
	}
}

func (uh *UserHandler) Configure(mux *mux.Router) {
	mux.HandleFunc("/api/v1/user/{id:[0-9]+}", GetUserById).Methods("GET")
	mux.HandleFunc("/api/v1/user/register", RegisterUser).Methods("PUT")
	mux.HandleFunc("/api/v1/user/password", ChangePassword).Methods("POST")
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {

}





