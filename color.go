package passbook

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Color struct {
	R uint8 // red
	G uint8 // green
	B uint8 // blue
}

func (c Color) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}

func (c Color) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Color) UnmarshalJSON(data []byte) error {
	_, err := fmt.Fscanf(bytes.NewReader(data), "\"rgb(%d, %d, %d)\"", &(c.R), &(c.G), &(c.B))
	return err
}
