package cfgloader_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ashep/go-cfgloader"
)

func TestLoadJSON(tt *testing.T) {
	tt.Run("InvalidInJSON", func(t *testing.T) {
		out := outStruct{}
		err := cfgloader.LoadJSON([]byte("{]"), &out, nil)

		assert.EqualError(t, err, "invalid character ']' looking for beginning of object key string")
	})

	tt.Run("InvalidSchema", func(t *testing.T) {
		out := outStruct{}
		err := cfgloader.LoadJSON([]byte(`{"foo":"bar"}`), &out, []byte(`{]`))

		assert.NotErrorIs(t, err, cfgloader.SchemaValidationError{})
		assert.EqualError(t, err, "invalid character ']' looking for beginning of object key string")
	})

	tt.Run("SchemaCheckFailed", func(t *testing.T) {
		out := outStruct{}
		err := cfgloader.LoadJSON([]byte(`{"foo":"bar","bar":""}`), &out, testSchema)

		assert.ErrorIs(t, err, cfgloader.SchemaValidationError{})
		assert.EqualError(t, err, "bar: String length must be greater than or equal to 1; foo: String length must be less than or equal to 2")
	})

	tt.Run("OkEmptySchema", func(t *testing.T) {
		out := outStruct{}
		err := cfgloader.LoadJSON([]byte(`{"foo":"bar","bar":""}`), &out, nil)

		require.NoError(t, err)
		assert.Equal(t, "bar", out.Foo)
		assert.Equal(t, "", out.Bar)
	})

	tt.Run("Ok", func(t *testing.T) {
		out := outStruct{}
		err := cfgloader.LoadJSON([]byte(`{"foo":"ba","bar":"baz"}`), &out, testSchema)

		require.NoError(t, err)
		assert.Equal(t, "ba", out.Foo)
		assert.Equal(t, "baz", out.Bar)
	})
}

func TestLoadJSONFromPath(tt *testing.T) {
	tt.Run("OkEmptySchema", func(t *testing.T) {
		p, err := writeTempFile(t, []byte(`{"foo":"ba","bar":"baz"}`), "")
		require.NoError(t, err)

		out := outStruct{}
		err = cfgloader.LoadJSONFromPath(p, &out, testSchema)

		require.NoError(t, err)
		assert.Equal(t, "ba", out.Foo)
		assert.Equal(t, "baz", out.Bar)
	})
}