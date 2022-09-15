// Temporary server definition so ns wouldn't ignore this file.
server: {
	name: "tmp"

	image: "redis:6.2.6-alpine@sha256:132337b9d7744ffee4fae83f51de53c3530935ad3ba528b7110f2d805f55cbf5"
}

resourceClasses: {
	"Database": {
		intent: {
			type:   "foundation.integrations.testdata.resources.classes.DatabaseIntent"
			source: "./proto1.proto"
		}
		produces: {
			type:   "foundation.integrations.testdata.resources.classes.DatabaseInstance"
			source: "./protos/proto2.proto"
		}
	}
}
