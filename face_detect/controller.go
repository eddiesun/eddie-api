package face_detect

import (
	"appengine"

	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Code    int32
	Message string
}

func init() {
	http.HandleFunc("/face", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("\n\n\n***** New Face Detection Request *****\n")

	// read json from request body
	var image struct {
		image []byte
	}
	var decoder = json.NewDecoder(r.Body)
	var err = decoder.Decode(&image)
	if err != nil {
		respErr(w, http.StatusBadRequest, "Invalid request body received.")
		c.Infof("Bad Request - Error: %v\n", err)
		return
	}

	// base 64 decode the image

	// send to Detect function

	Detect(image.image)

	// json encode the results

	// send back to response

	fmt.Fprint(w, "Face Detection API")
}

func respErr(w http.ResponseWriter, status int, message string) {
	var err struct {
		Success bool
		Message string
	}
	err.Success = false
	err.Message = message

	toReturn, _ := json.Marshal(err)

	http.Error(w, string(toReturn), status)
}
