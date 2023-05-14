package tiny

import (
	"encoding/json"
	"net/http"
)

func SendStatus(w http.ResponseWriter,status int) {
   w.WriteHeader(status)
}

func SendString(w http.ResponseWriter, text string) {
    w.Write([]byte(text))
}

func SendJSON(w http.ResponseWriter, obj any) error {
    data, err := json.Marshal(obj)
    if err != nil {
        return err
    }
    w.Header().Add("Content-Type", "application/json")
    w.Write(data)
    return nil
}


