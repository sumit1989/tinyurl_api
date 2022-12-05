package integrationtest

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
)

func getMockServer(address string) (*httptest.Server, error) {
	testListener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	server := httptest.NewUnstartedServer(handler)
	server.Listener.Close()
	server.Listener = testListener
	return server, nil

}
func performRequest(r http.Handler, method, path string, body io.Reader) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		fmt.Errorf("error creating request: %-v", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w, nil
}
