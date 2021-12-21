package user

import "github.com/go-chi/chi"

func (ac *AuthController) NewRouter() (r chi.Router){
	r = chi.NewRouter()

	r.Group("/auth", func(r chi.Router) {
		r.Post("/login", ac.HLogin)
		r.Post("/register", ac.HRegister)
	})

	return r
}





