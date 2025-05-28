package helper

import (
	"fmt"
	"net/http"
)

func GetRestoranIDFromCtx(request *http.Request) string {
	idFromCtx := request.Context().Value("id")
	restoran_id := idFromCtx.(string)

	return restoran_id
}

func GetUserIDFromCtx(request *http.Request) int {
	idFromCtx := request.Context().Value("user_id")
	user_id := idFromCtx.(int)
	fmt.Printf("GetUserIDFromCtx DEBUG - %v", user_id)

	return user_id
}
