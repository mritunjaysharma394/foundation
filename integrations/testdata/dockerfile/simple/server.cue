server: {
	name: "myserver"

	integration: docker: dockerfile: "Dockerfile"

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: {
				internetFacing: true
				httpRoutes: "*": ["/"]
			}
		}
	}
}