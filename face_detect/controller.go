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
		Base64DecodedImage []byte `json:"image"`
	}

	var decoder = json.NewDecoder(r.Body)
	if err := decoder.Decode(&image); err != nil {
		respErr(w, http.StatusBadRequest, "Invalid request body received.")
		c.Infof("Bad Request - Error: %v\n", err)
		return
	}

	// send to Detect function
	fmt.Fprintf(w, "%v", Detect(c, image.Base64DecodedImage))

	// json encode the results

	// send back to response

	// fmt.Fprint(w, "Face Detection API")
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
