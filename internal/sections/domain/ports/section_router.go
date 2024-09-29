package ports

import "github.com/gorilla/mux"

// TODO: Es una dependencia,en el dominio a futuro se trabara sin embargo debo ser aun mas fuerte indicando que quiero que se trabaje con Mux y no con multiples routers
type SectionRouter interface {
	SectionHandler
	*mux.Router

	InitialiceSectionRouter(...mux.MiddlewareFunc)
}
