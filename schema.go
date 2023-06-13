package cfgloader

import (
	"errors"
	"sort"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func formatSchemaErrors(res *gojsonschema.Result) error {
	r := make([]string, 0)
	for _, v := range res.Errors() {
		r = append(r, v.String())
	}

	sort.Strings(r)

	return errors.New(strings.Join(r, "; "))
}
