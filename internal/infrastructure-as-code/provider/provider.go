package provider

import "github.com/hashicorp/hcl/v2/hclwrite"

type IProvider interface {
	MakeTerraform(body *hclwrite.Body) *hclwrite.Attribute
	MakeProvider(body *hclwrite.Body) (*hclwrite.Block, *hclwrite.Body)
	MakeResource(body *hclwrite.Body, solution ISolution) (*hclwrite.Block, *hclwrite.Body)
}

type ISolution interface {
	MakeResource(body *hclwrite.Body) (*hclwrite.Block, *hclwrite.Body)
}

type Provider struct {
	Name    string
	Version string
	Source  string
}

func New(name string, version string, source string) *Provider {
	return &Provider{
		Name:    name,
		Version: version,
		Source:  source,
	}
}
