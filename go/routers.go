/*
 * Фильмотека API
 *
 * API Фильмотеки
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func WithDB(db *sql.DB) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next(w, r.WithContext(ctx))
		}
	}
}

func NewRouter(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()
	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.HandleFunc(route.Pattern, WithDB(db)(Logger(route.HandlerFunc, route.Name, db)))
		case "POST":
			router.HandleFunc(route.Pattern, WithDB(db)(Logger(route.HandlerFunc, route.Name, db)))
		case "PUT":
			router.HandleFunc(route.Pattern, WithDB(db)(Logger(route.HandlerFunc, route.Name, db)))
		case "DELETE":
			router.HandleFunc(route.Pattern, WithDB(db)(Logger(route.HandlerFunc, route.Name, db)))
		}
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
		"ActorDelete",
		"DELETE",
		"/actor/",
		ActorDelete,
	},

	Route{
		"ActorGet",
		"GET",
		"/actor/",
		ActorGet,
	},

	Route{
		"ActorPost",
		"POST",
		"/actor/",
		ActorPost,
	},

	Route{
		"ActorPut",
		"PUT",
		"/actor/",
		ActorPut,
	},

	Route{
		"ActorsGet",
		"GET",
		"/actors/",
		ActorsGet,
	},

	Route{
		"MoviesGet",
		"GET",
		"/movies/",
		MoviesGet,
	},

	Route{
		"MovieDelete",
		"DELETE",
		"/movie/",
		MovieDelete,
	},

	Route{
		"MovieGet",
		"GET",
		"/movie/",
		MovieGet,
	},

	Route{
		"MoviePost",
		"POST",
		"/movie/",
		MoviePost,
	},

	Route{
		"MoviePut",
		"PUT",
		"/movie/",
		MoviePut,
	},
}
