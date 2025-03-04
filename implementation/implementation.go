package implementation

import (
	"graphql-go/compatibility-unit-tests/types"
)

var GraphqlGoImplementation = types.Implementation{
	Repo: types.Repository{
		Name:          "graphql-go-graphql",
		URL:           "https://github.com/graphql-go/graphql",
		ReferenceName: "v0.8.1",
		Dir:           "./repos/graphql-go-graphql/",
	},
	Type: types.GoImplementationType,
}

var GraphqlJSImplementation = types.Implementation{
	Repo: types.Repository{
		Name:          "graphql-graphql-js",
		URL:           "https://github.com/graphql/graphql-js",
		ReferenceName: "v0.6.0",
		Dir:           "./repos/graphql-graphql-js/",
	},
	Type:              types.RefImplementationType,
	TestNamesFilePath: "./puller-js/unit-tests.txt",
}

var RefImplementation = GraphqlJSImplementation

var Implementations = []types.Implementation{GraphqlGoImplementation}

var gqlGoImplURL = GraphqlGoImplementation.MapKey("Implementation")
var jsImplURL = GraphqlJSImplementation.MapKey("Implementation")

var ImplementationsMap = map[string]types.Implementation{
	gqlGoImplURL: GraphqlGoImplementation,
	jsImplURL:    GraphqlJSImplementation,
}
