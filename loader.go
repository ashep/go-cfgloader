package cfgloader

import (
	"errors"
	"strings"
)

func LoadFromFile(path string, out interface{}, schema []byte) error {
	switch {
	case strings.HasSuffix(path, ".yml"), strings.HasSuffix(path, ".yaml"):
		return LoadYAMLFromPath(path, out, schema)
	case strings.HasSuffix(path, ".json"):
		return errors.New("not implemented")
	}

	return nil
}
