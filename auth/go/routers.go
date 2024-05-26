package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreatePostPost",
		strings.ToUpper("Post"),
		"/create_post",
		CreatePostPost,
	},

	Route{
		"DeletePostPostIdDelete",
		strings.ToUpper("Delete"),
		"/delete_post/{post_id}",
		DeletePostPostIdDelete,
	},

	Route{
		"GetPostPostIdGet",
		strings.ToUpper("Get"),
		"/get_post/{post_id}",
		GetPostPostIdGet,
	},

	Route{
		"GetPostsGet",
		strings.ToUpper("Get"),
		"/get_posts",
		GetPostsGet,
	},

	Route{
		"LoginPost",
		strings.ToUpper("Post"),
		"/login",
		LoginPost,
	},

	Route{
		"RegisterPost",
		strings.ToUpper("Post"),
		"/register",
		RegisterPost,
	},

	Route{
		"UpdatePostPostIdPut",
		strings.ToUpper("Put"),
		"/update_post/{post_id}",
		UpdatePostPostIdPut,
	},

	Route{
		"UpdatePut",
		strings.ToUpper("Put"),
		"/update",
		UpdatePut,
	},
}
