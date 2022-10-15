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
		Hashicorp: *hashicorp.New(".pkr.hcl"),
	}
}

func (pk Packer) MakePacker(version string) (*hclwrite.Block, *hclwrite.Body) {
	block, body := pk.Hashicorp.NewBlock("packer", nil)
	body.SetAttributeValue("required_version", cty.StringVal(version))
	return block, body
}

func (pk Packer) MakePlugins(body *hclwrite.Body, plugins []plugin.Plugin) (*hclwrite.Block, *hclwrite.Body) {
	body.AppendNewline()

	pgBlock := body.AppendNewBlock("required_plugins", nil)
	pgBody := pgBlock.Body()

	for _, p := range plugins {
		pgBody.SetAttributeValue(p.Name, cty.ObjectVal(map[string]cty.Value{
			"version": cty.StringVal(p.Version),
			"source":  cty.StringVal(p.Source),
		}))
	}
	return pgBlock, pgBody
}

func (pk Packer) MakeSourceBlock(labels []string) (*hclwrite.Block, *hclwrite.Body) {
	block, body := pk.Hashicorp.NewBlock("source", labels)

	switch labels[0] {
	case "proxmox":

	}
	return block, body
}
