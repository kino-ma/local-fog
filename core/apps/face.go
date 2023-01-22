package apps

import (
	"fmt"
	"log"

	"github.com/Kagami/go-face"
)

const imageFile string = "./data/sample.jpeg"
const modelsDir string = "data/models"

func RecognizeFace(body []byte) ([]byte, error) {
	// Init the recognizer.
	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	// Free the resources when you're finished.
	defer rec.Close()

	// Test image with 10 faces.
	// Recognize faces on that image.
	faces, err := rec.RecognizeFile(imageFile)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	msg := fmt.Sprintf("found %v faces", len(faces))
	log.Print(msg)

	out := []byte(msg)
	return out, nil
}
