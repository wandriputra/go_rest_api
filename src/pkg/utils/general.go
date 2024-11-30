package utils

import "fmt"

type GeneralResponse struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func PanicIfNeeded(err any, message ...string) {
	if err != nil {
		if fmt.Sprintf("%s", err) == "record not found" && len(message) > 0 {
			panic(message[0])
		} else {
			panic(err)
		}
	}
}
