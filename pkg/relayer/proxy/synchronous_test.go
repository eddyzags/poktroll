package proxy

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSynchronousRPCServer_Forward(t *testing.T) {
	tests := []struct {
		name                string
		inServiceID         string
		inReq               *http.Request
		mockHTTPHandler     http.HandlerFunc
		expectedForwardBody string
		expectedStatusCode  int
	}{
		{
			name: "OK",
			mockHandler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				b, err := io.ReadAll(req.Body)
				require.NoError(t, err)

				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte("pong"))
			}),
			inServiceID: "6bcfd222",
			inReq: &http.Request{
				Body: io.NopCloser(bytes.NewBufferString(`ping`)),
			},
			expectedStatusCode:  http.StatusOK,
			expectedForwardBody: `pong`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
