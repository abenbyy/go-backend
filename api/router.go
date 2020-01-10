package api


import(
	"github.com/abenbyy/go-backend/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r := mux.NewRouter()
	r.Use(middleware.LogMiddleware);




	return r
}

