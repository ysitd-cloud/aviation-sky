package flyer

import (
	"bytes"
	"encoding/gob"
)

func init() {
	gob.Register(&Flyer{})
}

func encode(f *Flyer) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(data []byte, v *Flyer) error {
	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(v)
	return err
}
