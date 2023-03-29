package custom

import (
	"os"

	"github.com/blushft/go-diagrams/diagram"
)

type ImageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Image = &ImageContainer{
	path: os.Getenv("PWD"),
	opts: diagram.OptionSet{diagram.Provider("custom"), diagram.NodeShape("none")},
}

func (c *ImageContainer) CustomImage(path string, opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon(path)}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
