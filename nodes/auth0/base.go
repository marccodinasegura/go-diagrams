package auth0

import "github.com/marccodinasegura/go-diagrams/diagram"

type baseContainer struct {
	path string
	opts []diagram.NodeOption
}

var Base = &baseContainer{
	opts: diagram.OptionSet{diagram.Provider("auth0"), diagram.NodeShape("none")},
	path: "assets/auth0",
}

func (c *baseContainer) MongoDB(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/auth0/mongodb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
