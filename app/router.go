package app

import (
	"context"
	"fmt"
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
	fmt.Printf("CleanPath DEBUG - %v", cleanPath)
	for _, route := range r.routes {
		fmt.Printf("\nroute DEBUG - %v", route)
		if matches(cleanPath, route.Pattern) && route.Method == req.Method {
			fmt.Printf("\nroute.Pattern DEBUG - %v", route.Pattern)
			handler := route.Handler
			if route.RequireAuth {
				fmt.Printf("\nSUKSES LEWAT AUTH DEBUG")
				handler = func(w http.ResponseWriter, req *http.Request) {
					middleware.JWTAuthMiddleware(http.HandlerFunc(route.Handler)).ServeHTTP(w, req)
				}
			}
			if id := extractID(cleanPath, route.Pattern); id != "" {
				ctx := context.WithValue(req.Context(), "id", id)
				req = req.WithContext(ctx)
			}
			handler(w, req)
			fmt.Printf("\nROUTE SUCCESS DEBUG")
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
		fmt.Printf("\nmatches func base DEBUG - %v", base)
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

// // pre 27 May
// type Router struct {
// 	restoranController controller.RestoranController
// 	userController     controller.UserController
// }

// func NewRouter(restoranController controller.RestoranController, userController controller.UserController) http.Handler {
// 	return &Router{
// 		restoranController: restoranController,
// 		userController:     userController,
// 	}
// }

// func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			exception.ErrorHandler(writer, request, err)
// 		}
// 	}()

// 	cleanPath := path.Clean(request.URL.Path)

// 	if router.rootHandler(cleanPath, writer) {
// 		return
// 	}

// 	if router.restoranHandler(cleanPath, writer, request) {
// 		return
// 	}

// 	if router.userHandler(cleanPath, writer, request) {
// 		return
// 	}

// 	http.NotFound(writer, request)
// 	if router.staticFileServer(cleanPath, writer, request) {
// 		return
// 	}

// 	http.NotFound(writer, request)
// }

// func (router *Router) rootHandler(cleanPath string, writer http.ResponseWriter) bool {
// 	if cleanPath == "/" {
// 		message := "NesaNest RESTful API - 2025"
// 		helper.WriteToResponseBody(writer, message)
// 		return true
// 	}
// 	return false
// }

// func (router *Router) restoranHandler(cleanPath string, writer http.ResponseWriter, request *http.Request) bool {
// 	if cleanPath == "/api/v1/restoran" && request.Method == http.MethodGet {
// 		router.restoranController.FindAll(writer, request)
// 		return true
// 	}

// 	if cleanPath == "/api/v1/restoran" && request.Method == http.MethodPost {
// 		middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			router.restoranController.Create(w, r)
// 		})).ServeHTTP(writer, request)
// 		return true
// 	}

// 	if strings.HasPrefix(cleanPath, "/api/v1/restoran/") {
// 		id := strings.TrimPrefix(cleanPath, "/api/v1/restoran/")
// 		if id == "" {
// 			http.Error(writer, "missing restoran ID", http.StatusBadRequest)
// 			return true
// 		}

// 		switch request.Method {
// 		case http.MethodGet:
// 			router.restoranController.FindById(writer, request, id)
// 		case http.MethodPut:
// 			middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				router.restoranController.Update(w, r, id)
// 			})).ServeHTTP(writer, request)
// 		case http.MethodDelete:
// 			middleware.NewAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				router.restoranController.Delete(w, r, id)
// 			})).ServeHTTP(writer, request)
// 		default:
// 			writer.WriteHeader(http.StatusMethodNotAllowed)
// 		}
// 		return true
// 	}

// 	return false
// }

// func (router *Router) userHandler(cleanPath string, writer http.ResponseWriter, request *http.Request) bool {
// 	// Register user
// 	if cleanPath == "/api/v1/users/register" && request.Method == http.MethodPost {
// 		router.userController.Register(writer, request)
// 		return true
// 	}

// 	// Login user
// 	if cleanPath == "/api/v1/users/login" && request.Method == http.MethodPost {
// 		router.userController.Login(writer, request)
// 		return true
// 	}

// 	// Forgot password
// 	if cleanPath == "/api/v1/users/forgot-password" && request.Method == http.MethodPost {
// 		router.userController.ForgotPassword(writer, request)
// 		return true
// 	}

// 	// Get all users
// 	if cleanPath == "/api/v1/users" && request.Method == http.MethodGet {
// 		router.userController.FindAll(writer, request)
// 		return true
// 	}

// 	// Get user by ID
// 	if strings.HasPrefix(cleanPath, "/api/v1/users/") {
// 		id := strings.TrimPrefix(cleanPath, "/api/v1/users/")
// 		if id == "" {
// 			http.Error(writer, "missing user ID", http.StatusBadRequest)
// 			return true
// 		}
// 		if request.Method == http.MethodDelete {
// 			middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				authHeader := r.Header.Get("Authorization")
// 				tokenString := ""
// 				if strings.HasPrefix(authHeader, "Bearer ") {
// 					tokenString = strings.TrimPrefix(authHeader, "Bearer ")
// 				}
// 				claims, err := helper.ParseJWT(tokenString)
// 				if err != nil {
// 					http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 					return
// 				}
// 				userId := int(claims["user_id"].(float64))
// 				router.userController.Delete(w, r, strconv.Itoa(userId))
// 			})).ServeHTTP(writer, request)
// 			return true
// 		}
// 		writer.WriteHeader(http.StatusMethodNotAllowed)
// 		return true
// 	}

// 	// Update profile user (JWT)
// 	if strings.HasPrefix(cleanPath, "/api/v1/users/profile/") {
// 		id := strings.TrimPrefix(cleanPath, "/api/v1/users/profile/")
// 		if id == "" {
// 			http.Error(writer, "missing user ID", http.StatusBadRequest)
// 			return true
// 		}
// 		if request.Method == http.MethodGet {
// 			middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				authHeader := r.Header.Get("Authorization")
// 				tokenString := ""
// 				if strings.HasPrefix(authHeader, "Bearer ") {
// 					tokenString = strings.TrimPrefix(authHeader, "Bearer ")
// 				}
// 				claims, err := helper.ParseJWT(tokenString)
// 				if err != nil {
// 					http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 					return
// 				}
// 				userId := strconv.FormatFloat(claims["user_id"].(float64), 'f', 0, 64)
// 				router.userController.FindById(w, r, userId)
// 			})).ServeHTTP(writer, request)

// 			return true
// 		}
// 		if request.Method == http.MethodPut {
// 			middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				router.userController.UpdateProfile(w, r)
// 			})).ServeHTTP(writer, request)

// 			return true
// 		}

// 		writer.WriteHeader(http.StatusMethodNotAllowed)

// 		return true
// 	}

// 	// Update password user (JWT)
// 	if strings.HasPrefix(cleanPath, "/api/v1/users/password/") {
// 		id := strings.TrimPrefix(cleanPath, "/api/v1/users/password/")
// 		if id == "" {
// 			http.Error(writer, "missing user ID", http.StatusBadRequest)
// 			return true
// 		}
// 		if request.Method == http.MethodPut {
// 			middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				router.userController.UpdatePassword(w, r)
// 			})).ServeHTTP(writer, request)
// 			return true
// 		}
// 		writer.WriteHeader(http.StatusMethodNotAllowed)
// 		return true
// 	}

// 	return false
// }

// func (router *Router) staticFileServer(cleanPath string, writer http.ResponseWriter, request *http.Request) bool {
// 	if strings.HasPrefix(cleanPath, "/static/") {
// 		fs := http.Dir("static")
// 		http.StripPrefix("/static/", http.FileServer(fs)).ServeHTTP(writer, request)

// 		return true
// 	}

// 	return false
// }
