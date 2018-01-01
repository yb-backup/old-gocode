package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handler(method, params string) *Response {
	return &Response{"OK", "ECHO: " + method + " - " + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	defer client1.Close()
	defer client2.Close()

	resp1, _ := client1.Call("foo", "From client1")
	resp2, _ := client2.Call("foo", "From client2")

	if resp1.Body != "ECHO: foo - From client1" ||
		resp2.Body != "ECHO: foo - From client2" {
		t.Error("IpcClient.Call failed.")
	}

}
