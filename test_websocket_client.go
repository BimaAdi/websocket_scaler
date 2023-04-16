package websocketscaler

type LogMessageStruct struct {
	Action   string
	SocketId string
	Message  string
}

type TestWebsocketClient struct {
	Log []LogMessageStruct
}

func InitTestWebsocketClient() TestWebsocketClient {
	return TestWebsocketClient{
		Log: []LogMessageStruct{},
	}
}

func (socket *TestWebsocketClient) Emit(message string) {
	socket.Log = append(socket.Log, LogMessageStruct{
		Action:   "Emit",
		SocketId: "",
		Message:  message,
	})
}

func (socket *TestWebsocketClient) Broadcast(message string) {
	socket.Log = append(socket.Log, LogMessageStruct{
		Action:   "Broadcast",
		SocketId: "",
		Message:  message,
	})
}

func (socket *TestWebsocketClient) ToSocketId(socketId string, message string) {
	socket.Log = append(socket.Log, LogMessageStruct{
		Action:   "ToSocketId",
		SocketId: socketId,
		Message:  message,
	})
}
