package apps

func Hello(body []byte) ([]byte, error) {
	out := []byte("hello, world")

	return out, nil
}
