package cfgloader_test

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ashep/go-cfgloader"
)

func TestLoadYAML(tt *testing.T) {
	tt.Run("InvalidInYAML", func(t *testing.T) {
		out := outStruct{}
		err := cfgloader.LoadYAML([]byte("foo"), &out, nil)

		assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `foo` into cfgloader_test.outStruct")
	})

	tt.Run("InvalidSchema", func(t *testing.T) {
		out := outStruct{}
		err := cfgloader.LoadYAML([]byte("foo: bar"), &out, []byte(`{]`))

		assert.EqualError(t, err, "invalid character ']' looking for beginning of object key string")
	})

	tt.Run("SchemaCheckFailed", func(t *testing.T) {
		out := outStruct{}
		err := cfgloader.LoadYAML([]byte("foo: bar"), &out, testSchema)

		assert.EqualError(t, err, "bar: String length must be greater than or equal to 1; foo: String length must be less than or equal to 2")
	})

	tt.Run("Ok", func(t *testing.T) {
		out := outStruct{}
		err := cfgloader.LoadYAML([]byte("foo: ba\nbar: baz"), &out, testSchema)

		assert.NoError(t, err)
	})
}

func TestLoadYAMLFromFile(tt *testing.T) {
	tt.Run("InvalidInYAML", func(t *testing.T) {
		p, err := writeTempFile(t, []byte("foo"))
		require.NoError(t, err)

		out := outStruct{}
		err = cfgloader.LoadYAMLFromPath(p, &out, nil)
		assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `foo` into cfgloader_test.outStruct")
	})

	tt.Run("InvalidSchema", func(t *testing.T) {
		p, err := writeTempFile(t, []byte("foo: bar"))
		require.NoError(t, err)

		out := outStruct{}
		err = cfgloader.LoadYAMLFromPath(p, &out, []byte(`{]`))

		assert.EqualError(t, err, "invalid character ']' looking for beginning of object key string")
	})

	tt.Run("SchemaCheckFailed", func(t *testing.T) {
		p, err := writeTempFile(t, []byte("foo: bar"))
		require.NoError(t, err)

		out := outStruct{}
		err = cfgloader.LoadYAMLFromPath(p, &out, testSchema)

		assert.EqualError(t, err, "bar: String length must be greater than or equal to 1; foo: String length must be less than or equal to 2")
	})

	tt.Run("Ok", func(t *testing.T) {
		p, err := writeTempFile(t, []byte("foo: ba\nbar: baz"))
		require.NoError(t, err)

		out := outStruct{}
		err = cfgloader.LoadYAMLFromPath(p, &out, testSchema)

		assert.NoError(t, err)
	})
}

func writeTempFile(t *testing.T, b []byte) (string, error) {
	p := path.Join(t.TempDir(), "testFile")

	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err = f.Write(b); err != nil {
		return "", err
	}

	return p, nil
}
