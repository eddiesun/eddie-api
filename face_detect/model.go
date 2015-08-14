package face_detect

import (
	"appengine"
	"appengine/urlfetch"

	_ "net/http"
	"net/url"
)

const (
	faceplusplusURL = "https://faceplusplus-faceplusplus.p.mashape.com/detection/detect"
	testingImageUrl = "http://www.faceplusplus.com/wp-content/themes/faceplusplus/assets/img/demo/1.jpg"
)

type Face struct {
	Raw []byte
}

func Detect(c appengine.Context, image []byte) []Face {
	var urlValues = url.Values{}
	urlValues.Set("url", testingImageUrl)
	var u, _ = url.Parse(faceplusplusURL)
	u.RawQuery = urlValues.Encode()

	var client = urlfetch.Client(c)
	resp, err := client.Get(u.String())

	c.Infof("resp: %v\n", resp)
	c.Infof("err: %v\n", err)

	var faces = make([]Face, 0)
	faces = append(faces, Face{image})
	return faces
}
