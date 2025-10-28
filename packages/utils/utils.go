package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func Parsing(r *http.Request, x interface{}) {
	if data, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(data), x); err != nil {
			return
		}
	}
}
