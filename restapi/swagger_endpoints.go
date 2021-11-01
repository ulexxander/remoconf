package restapi

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (e *Endpoints) GetSwaggerDocs(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write(e.SwaggerDocs); err != nil {
		httpError(w, http.StatusInternalServerError)
		e.Logger.Printf("error writing swagger docs: %s", err)
	}
}

func (e *Endpoints) GetSwaggerWebInterface() http.HandlerFunc {
	return httpSwagger.Handler(httpSwagger.URL("http://localhost:4000/swagger/docs.json"))
}
