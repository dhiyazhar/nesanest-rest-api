// filepath: [user_controller_impl.go](http://_vscodecontentref_/0)
package controller

import (
    "net/http"
    "nesanest-rest-api/helper"
    "nesanest-rest-api/model/web"
    "nesanest-rest-api/service"
    "strconv"
    "strings"
)

type UserControllerImpl struct {
    UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
    return &UserControllerImpl{
        UserService: userService,
    }
}

func (controller *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
    var request web.UserRegisterRequest
    helper.ReadFromRequestBody(r, &request)
    response := controller.UserService.Register(r.Context(), request)
    helper.WriteToResponseBody(w, response)
}

func (controller *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
    var request web.UserLoginRequest
    helper.ReadFromRequestBody(r, &request)
    response, token := controller.UserService.Login(r.Context(), request)
    helper.WriteToResponseBody(w, map[string]interface{}{
        "user":  response,
        "token": token,
    })
}

func (controller *UserControllerImpl) UpdateProfile(w http.ResponseWriter, r *http.Request) {
    // Ambil JWT dari header Authorization
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

    var request web.UserUpdateUsernameRequest
    helper.ReadFromRequestBody(r, &request)
    request.Id = userId // Pakai ID dari JWT, bukan dari body

    response := controller.UserService.UpdateProfile(r.Context(), request)
    helper.WriteToResponseBody(w, response)
}

func (controller *UserControllerImpl) UpdatePassword(w http.ResponseWriter, r *http.Request) {
    var request web.UserUpdatePasswordRequest
    helper.ReadFromRequestBody(r, &request)
    controller.UserService.UpdatePassword(r.Context(), request)
    helper.WriteToResponseBody(w, map[string]string{"message": "Password berhasil diupdate"})
}

func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, id string) {
    userId, err := strconv.Atoi(id)
    helper.PanicIfError(err)
    controller.UserService.Delete(r.Context(), userId)
    helper.WriteToResponseBody(w, map[string]string{"message": "User berhasil dihapus"})
}