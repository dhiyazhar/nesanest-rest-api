package helper

import (
	"log/slog"
)

func PanicIfError(err error) {
	if err != nil {
		slog.Error("", slog.Any("err", err))
		panic(err)
	}
}
