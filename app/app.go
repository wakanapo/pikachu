package app

import (
	"net/http"
	"translator"
)

func init() {
	http.HandleFunc("/", translator.HandleTrans)
}

