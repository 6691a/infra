package packer

import (
	"github.com/6691a/infra/internal/format/hashicorp"
	"github.com/6691a/infra/internal/image-builder/packer/plugin"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type Packer struct {
	hashicorp.Hashicorp
}

func New() *Packer {
	return &Packer{
		*hashicorp.New(".pkr.hcl"),
	}
}

func (pk Packer) MakePackerBlock(version string) (*hclwrite.Block, *hclwrite.Body) {
	block, body := pk.Hashicorp.NewBlock("packer", nil)
	body.SetAttributeValue("required_version", cty.StringVal(version))
	return block, body
}

func (pk Packer) MakePlugins(body *hclwrite.Body, plugins []plugin.Plugin) {
	body.AppendNewline()

	pluginBlock := body.AppendNewBlock("required_plugins", nil)

	pluginBody := pluginBlock.Body()

	for _, p := range plugins {
		pluginBody.SetAttributeValue(p.Name, cty.ObjectVal(map[string]cty.Value{
			"version": cty.StringVal(p.Version),
			"source":  cty.StringVal(p.Source),
		}))
	}
}

func (pk Packer) MakeSourceBlock(labels []string) (*hclwrite.Block, *hclwrite.Body) {
	block, body := pk.Hashicorp.NewBlock("source", labels)

	switch labels[0] {
	case "procmox":

	}

	return block, body
}
