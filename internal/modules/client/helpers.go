package client

import (
	"encoding/json"
	"io"
)

func parseClientFromBody(b io.Reader) (Client, error) {
	c := Client{}
	err := json.NewDecoder(b).Decode(&c)
	return c, err
}
