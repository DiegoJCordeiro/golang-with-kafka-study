package webserver

import (
	_ "github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (webServer *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	webServer.Handlers[path] = handler
}

func (webServer *WebServer) Start() error {

	webServer.Router.Use(middleware.Logger)
	webServer.Router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("docs/doc.json"),
	))

	for path, handler := range webServer.Handlers {
		webServer.Router.Handle(path, handler)
	}

	err := http.ListenAndServe(webServer.WebServerPort, webServer.Router)

	if err != nil {
		return err
	}

	return nil
}
