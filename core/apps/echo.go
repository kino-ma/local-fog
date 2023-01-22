package apps

func Echo(body []byte) ([]byte, error) {
	out := make([]byte, len(body))
	copy(out, body)

	return out, nil
}
