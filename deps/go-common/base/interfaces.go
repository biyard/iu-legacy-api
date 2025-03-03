package base

type Router interface {
	OnInit(router PathRouter)
}

type PathRouter interface {
	Group(string, Router)

	Any(string, any)
	GET(string, any)
	POST(string, any)
	DELETE(string, any)
	PATCH(string, any)
	PUT(string, any)
	OPTIONS(string, any)
	HEAD(string, any)

	// TODO: implement these
	// Use(...HandlerFunc) IRoutes
	// Handle(string, string, any) IRoutes
	// Match([]string, string, any) IRoutes

	// StaticFile(string, string) IRoutes
	// StaticFileFS(string, string, http.FileSystem) IRoutes
	// Static(string, string) IRoutes
	// StaticFS(string, http.FileSystem) IRoutes
}
