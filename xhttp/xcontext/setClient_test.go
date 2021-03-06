package xcontext

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Comcast/webpa-common/xhttp"
	"github.com/stretchr/testify/assert"
)

func testSetClientDefault(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = SetClient(nil)(context.Background(), httptest.NewRequest("GET", "/", nil))
	)

	assert.Equal(http.DefaultClient, xhttp.GetClient(ctx))
}

func testSetClientCustom(t *testing.T) {
	var (
		assert = assert.New(t)

		expected = new(http.Client)
		ctx      = SetClient(expected)(context.Background(), httptest.NewRequest("GET", "/", nil))
	)

	assert.Equal(expected, xhttp.GetClient(ctx))
}

func TestSetClient(t *testing.T) {
	t.Run("Default", testSetClientDefault)
	t.Run("Custom", testSetClientCustom)
}
