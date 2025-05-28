package app

import (
	"context"
	"net/http"
	"path"
	"strings"

	"nesanest-rest-api/controller"
	"nesanest-rest-api/exception"
	"nesanest-rest-api/helper"
	"nesanest-rest-api/middleware"
)

type Route struct {
	Method      string
	Pattern     string
	Handler     http.HandlerFunc
	RequireAuth bool
}

type Router struct {
	routes []Route
}

func NewRouter(restoranController controller.RestoranController, userController controller.UserController) http.Handler {
	r := &Router{routes: []Route{}}

	// root
	r.Handle("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		helper.WriteToResponseBody(w, map[string]string{
			"message": "NesaNest RESTful API - 2025",
		})
	}, false)

	// static file server
	r.Handle("GET", "/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP, false)

	// restoran - public
	r.Handle("GET", "/api/v1/restoran", restoranController.FindAll, false)
	r.Handle("GET", "/api/v1/restoran/{id}", restoranController.FindById, false)

	// restoran - protected
	r.Handle("POST", "/api/v1/restoran", restoranController.Create, true)
	r.Handle("PUT", "/api/v1/restoran/{id}", restoranController.Update, true)
	r.Handle("DELETE", "/api/v1/restoran/{id}", restoranController.Delete, true)

	// user
	r.Handle("POST", "/api/v1/users/register", userController.Register, false)
	r.Handle("POST", "/api/v1/users/login", userController.Login, false)
	r.Handle("POST", "/api/v1/users/forgot-password", userController.ForgotPassword, false)

	//user - protected
	r.Handle("GET", "/api/v1/users", userController.FindAll, true)
	r.Handle("GET", "/api/v1/users/profile", userController.FindById, true)
	r.Handle("PUT", "/api/v1/users/profile", userController.UpdateProfile, true)
	r.Handle("PUT", "/api/v1/users/password", userController.UpdatePassword, true)
	r.Handle("DELETE", "/api/v1/users/{id}", userController.Delete, true)

	return r
}

func (r *Router) Handle(method, pattern string, h http.HandlerFunc, auth bool) {
	r.routes = append(r.routes, Route{method, pattern, h, auth})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			exception.ErrorHandler(w, req, err)
		}
	}()

	cleanPath := path.Clean(req.URL.Path)
	for _, route := range r.routes {
		if matches(cleanPath, route.Pattern) && route.Method == req.Method {
			handler := route.Handler
			if route.RequireAuth {
				handler = func(w http.ResponseWriter, req *http.Request) {
					middleware.JWTAuthMiddleware(http.HandlerFunc(route.Handler)).ServeHTTP(w, req)
				}
			}
			if idParams := extractID(cleanPath, route.Pattern); idParams != "" {
				var contextKey string
				if strings.Contains(route.Pattern, "restoran") {
					contextKey = "restoran_id"
				}

				ctx := context.WithValue(req.Context(), contextKey, idParams)
				req = req.WithContext(ctx)
			}
			handler(w, req)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	helper.WriteToResponseBody(w, map[string]string{
		"error": "resource not found",
	})
}

func matches(urlPath, pattern string) bool {
	if strings.Contains(pattern, "{id}") {
		base := strings.TrimSuffix(pattern, "/{id}")
		if strings.HasPrefix(urlPath, base+"/") {
			id := strings.TrimPrefix(urlPath, base+"/")
			return id != ""
		}
		return false
	}

	if strings.HasSuffix(pattern, "/") && pattern != "/" {
		return strings.HasPrefix(urlPath, pattern)
	}

	return urlPath == pattern
}

func extractID(urlPath, pattern string) string {
	if !strings.Contains(pattern, "{id}") {
		return ""
	}
	base := strings.TrimSuffix(pattern, "/{id}")
	return strings.TrimPrefix(urlPath, base+"/")
}
