// events are emitted from the server HTTP handlers and are consumed in templ templates on the client.
package events

import "github.com/ItsNotGoodName/go-web-app-example/pkg/htmx"

var (
	HelloWorld htmx.Event = htmx.NewEventString("hello-world")
)

func Toast(message string) htmx.Event {
	return htmx.NewEvent("toast", message)
}
