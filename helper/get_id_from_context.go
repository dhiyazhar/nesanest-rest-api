package helper

import (
	"fmt"
	"net/http"
)

func GetRestoranIDFromCtx(request *http.Request) int {
	idFromCtx := request.Context().Value("restoran_id")
	restoran_id := idFromCtx.(int)

	return restoran_id
}

func GetUserIDFromCtx(request *http.Request) int {
	idFromCtx := request.Context().Value("user_id")
	user_id := idFromCtx.(int)
	fmt.Printf("GetUserIDFromCtx DEBUG - %v", user_id)

	return user_id
}
