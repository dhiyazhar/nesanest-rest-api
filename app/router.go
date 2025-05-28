package app

import (
    "nesanest-rest-api/controller"
    "nesanest-rest-api/exception"
    "nesanest-rest-api/helper"
    "nesanest-rest-api/middleware"
    "net/http"
    "strconv"
    "path"
    "strings"
)

type Router struct {
    restoranController controller.RestoranController
    userController     controller.UserController
    reviewController   controller.ReviewController
}

func NewRouter(restoranController controller.RestoranController, userController controller.UserController, reviewController controller.ReviewController) http.Handler {
    return &Router{
        restoranController: restoranController,
        userController:     userController,
        reviewController:   reviewController,
    }
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    defer func() {
        if err := recover(); err != nil {
            exception.ErrorHandler(writer, request, err)
        }
    }()

    cleanPath := path.Clean(request.URL.Path)

    if router.rootHandler(cleanPath, writer) {
        return
    }

    if router.restoranHandler(cleanPath, writer, request) {
        return
    }

    if router.userHandler(cleanPath, writer, request) {
        return
    }

    if router.reviewHandler(cleanPath, writer, request) {
        return
    }

    http.NotFound(writer, request)
}

func (router *Router) rootHandler(cleanPath string, writer http.ResponseWriter) bool {
    if cleanPath == "/" {
        message := "NesaNest RESTful API - 2025"
        helper.WriteToResponseBody(writer, message)
        return true
    }
    return false
}

func (router *Router) restoranHandler(cleanPath string, writer http.ResponseWriter, request *http.Request) bool {
    if cleanPath == "/api/v1/restoran" && request.Method == http.MethodGet {
        router.restoranController.FindAll(writer, request)
        return true
    }

    if cleanPath == "/api/v1/restoran" && request.Method == http.MethodPost {
        middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            router.restoranController.Create(w, r)
        })).ServeHTTP(writer, request)
        return true
    }

    if strings.HasPrefix(cleanPath, "/api/v1/restoran/") {
        id := strings.TrimPrefix(cleanPath, "/api/v1/restoran/")
        if id == "" {
            http.Error(writer, "missing restoran ID", http.StatusBadRequest)
            return true
        }

        switch request.Method {
        case http.MethodGet:
            router.restoranController.FindById(writer, request, id)
        case http.MethodPut:
            middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                router.restoranController.Update(w, r, id)
            })).ServeHTTP(writer, request)
        case http.MethodDelete:
            middleware.NewAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                router.restoranController.Delete(w, r, id)
            })).ServeHTTP(writer, request)
        default:
            writer.WriteHeader(http.StatusMethodNotAllowed)
        }
        return true
    }

    return false
}

func (router *Router) userHandler(cleanPath string, writer http.ResponseWriter, request *http.Request) bool {
    // Register user
    if cleanPath == "/api/v1/users/register" && request.Method == http.MethodPost {
        router.userController.Register(writer, request)
        return true
    }

    // Login user
    if cleanPath == "/api/v1/users/login" && request.Method == http.MethodPost {
        router.userController.Login(writer, request)
        return true
    }

    // Forgot password
    if cleanPath == "/api/v1/users/forgot-password" && request.Method == http.MethodPost {
        router.userController.ForgotPassword(writer, request)
        return true
    }

    // --- Penting: blok profile dan password di atas blok users ---
    // Update profile user (JWT)
    if strings.HasPrefix(cleanPath, "/api/v1/users/profile/") {
        id := strings.TrimPrefix(cleanPath, "/api/v1/users/profile/")
        if id == "" {
            http.Error(writer, "missing user ID", http.StatusBadRequest)
            return true
        }
        if request.Method == http.MethodPut {
            middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                router.userController.UpdateProfile(w, r)
            })).ServeHTTP(writer, request)
            return true
        }
        writer.WriteHeader(http.StatusMethodNotAllowed)
        return true
    }

    // Update password user (JWT)
    if strings.HasPrefix(cleanPath, "/api/v1/users/password/") {
        id := strings.TrimPrefix(cleanPath, "/api/v1/users/password/")
        if id == "" {
            http.Error(writer, "missing user ID", http.StatusBadRequest)
            return true
        }
        if request.Method == http.MethodPut {
            middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                router.userController.UpdatePassword(w, r)
            })).ServeHTTP(writer, request)
            return true
        }
        writer.WriteHeader(http.StatusMethodNotAllowed)
        return true
    }

    // Get user by ID, Delete user
    if strings.HasPrefix(cleanPath, "/api/v1/users/") {
        id := strings.TrimPrefix(cleanPath, "/api/v1/users/")
        if id == "" {
            http.Error(writer, "missing user ID", http.StatusBadRequest)
            return true
        }
        if request.Method == http.MethodGet {
            router.userController.FindById(writer, request, id)
            return true
        }
        if request.Method == http.MethodDelete {
            middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                authHeader := r.Header.Get("Authorization")
                tokenString := ""
                if strings.HasPrefix(authHeader, "Bearer ") {
                    tokenString = strings.TrimPrefix(authHeader, "Bearer ")
                }
                claims, err := helper.ParseJWT(tokenString)
                if err != nil {
                    http.Error(w, "Unauthorized", http.StatusUnauthorized)
                    return
                }
                userId := int(claims["user_id"].(float64))
                router.userController.Delete(w, r, strconv.Itoa(userId))
            })).ServeHTTP(writer, request)
            return true
        }
        writer.WriteHeader(http.StatusMethodNotAllowed)
        return true
    }

    return false
}

func (router *Router) reviewHandler(cleanPath string, writer http.ResponseWriter, request *http.Request) bool {
    // Create review (JWT, upload gambar via form-data)
    if cleanPath == "/api/v1/review" && request.Method == http.MethodPost {
        middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            router.reviewController.CreateReview(w, r)
        })).ServeHTTP(writer, request)
        return true
    }

    // Get review by restoran
    if strings.HasPrefix(cleanPath, "/api/v1/review/restoran/") && request.Method == http.MethodGet {
        id := strings.TrimPrefix(cleanPath, "/api/v1/review/restoran/")
        router.reviewController.GetReviewsByRestoran(writer, request, id)
        return true
    }

    // Get review history by user (JWT)
    if cleanPath == "/api/v1/review/history" && request.Method == http.MethodGet {
        middleware.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            router.reviewController.GetReviewsByUser(w, r)
        })).ServeHTTP(writer, request)
        return true
    }

    return false
}