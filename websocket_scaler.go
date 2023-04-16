package websocketscaler

type Scaler struct {
	Namespaces      map[string]Namespace
	WebsocketClient WebsocketClient
}

type WebsocketClient interface {
	Emit(message string)
	Broadcast(message string)
	ToSocketId(socketId string, message string)
}

func InitScaler(websocketclient WebsocketClient) Scaler {
	return Scaler{
		Namespaces:      map[string]Namespace{},
		WebsocketClient: websocketclient,
	}
}

func (scaler *Scaler) TestClient(namespace_name string, event_name string, message string) {
	namespace, is_exists := scaler.Namespaces[namespace_name]
	if is_exists {
		event, is_exists := namespace.Event[event_name]
		if is_exists {
			event(scaler.WebsocketClient, message)
		}
	}
}

func (scaler *Scaler) Of(namespace_name string, namespace Namespace) {
	scaler.Namespaces[namespace_name] = namespace
}

type Namespace struct {
	Event map[string]func(WebsocketClient, string)
}

func InitNamespace() Namespace {
	return Namespace{
		Event: map[string]func(WebsocketClient, string){},
	}
}

func (namespace *Namespace) On(event_name string, method func(WebsocketClient, string)) {
	namespace.Event[event_name] = method
}
