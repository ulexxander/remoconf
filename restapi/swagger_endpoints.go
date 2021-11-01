package restapi

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (h *Handler) getSwaggerWebInterface() http.HandlerFunc {
	return httpSwagger.Handler(httpSwagger.URL("http://localhost:4000/swagger/docs.json"))
}

func (h *Handler) getSwaggerDocs(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write(h.swaggerDocs); err != nil {
		httpError(w, http.StatusInternalServerError)
		h.logger.Printf("error writing swagger docs: %s", err)
	}
}
