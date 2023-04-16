package websocketscaler_test

import (
	"testing"

	wsc "github.com/BimaAdi/websocket_scaler"
	"github.com/stretchr/testify/assert"
)

func TestInitWebSocketScaler(t *testing.T) {
	wsTestClient := wsc.InitTestWebsocketClient()
	scaler := wsc.InitScaler(&wsTestClient)
	assert.Equal(t, map[string]wsc.Namespace{}, scaler.Namespaces)
}

func TestServerImplementation(t *testing.T) {
	wsTestClient := wsc.InitTestWebsocketClient()
	app := wsc.InitScaler(&wsTestClient)
	namespace_chat := wsc.InitNamespace()
	namespace_chat.On("/broadcast", func(socket wsc.WebsocketClient, message string) {
		socket.Broadcast(message)
	})
	namespace_chat.On("/personal", func(socket wsc.WebsocketClient, message string) {
		socket.Emit(message)
	})
	app.Of("/chat", namespace_chat)
}

func TestEventEmit(t *testing.T) {
	wsTestClient := wsc.InitTestWebsocketClient()
	app := wsc.InitScaler(&wsTestClient)
	namespace_chat := wsc.InitNamespace()
	namespace_chat.On("/group", func(socket wsc.WebsocketClient, message string) {
		socket.Broadcast(message)
	})
	namespace_chat.On("/personal", func(socket wsc.WebsocketClient, message string) {
		socket.Emit(message)
	})
	app.Of("/chat", namespace_chat)
	app.TestClient("/chat", "/personal", "Hello")
	app.TestClient("/chat", "/group", "Hai")
	assert.Equal(t, []wsc.LogMessageStruct{
		{
			Action:   "Emit",
			SocketId: "",
			Message:  "Hello",
		},
		{
			Action:   "Broadcast",
			SocketId: "",
			Message:  "Hai",
		},
	}, wsTestClient.Log)
}
