package apps

import (
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

const classifierXmlFile string = "../../data/haarcascade_frontalface_default.xml"

func RecognizeFace(body []byte) ([]byte, error) {
	c := gocv.NewCascadeClassifier()
	defer c.Close()

	if !c.Load(classifierXmlFile) {
		err := fmt.Errorf("failed to load classifier xml file")
		return nil, err
	}

	img, err := gocv.IMDecode(body, gocv.IMReadUnchanged)
	if err != nil {
		err = fmt.Errorf("failed to decode image file: %w", err)
		return nil, err
	}

	rects := c.DetectMultiScale(img)
	msg := fmt.Sprintf("found %v faces", len(rects))
	log.Print(msg)

	out := []byte(msg)
	return out, nil
}
