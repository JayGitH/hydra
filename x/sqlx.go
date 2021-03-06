package x

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	jose "gopkg.in/square/go-jose.v2"
)

// swagger:type JSONWebKeySet
type JoseJSONWebKeySet struct {
	// swagger:ignore
	*jose.JSONWebKeySet
}

func (n *JoseJSONWebKeySet) Scan(value interface{}) error {
	v := fmt.Sprintf("%s", value)
	if len(v) == 0 {
		return nil
	}
	return errors.WithStack(json.Unmarshal([]byte(v), n))
}

func (n *JoseJSONWebKeySet) Value() (driver.Value, error) {
	value, err := json.Marshal(n)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return string(value), nil
}
