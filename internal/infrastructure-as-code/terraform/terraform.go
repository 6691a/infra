package terraform

import (
	"github.com/6691a/infra/internal/format/hashicorp"
	"github.com/6691a/infra/internal/infrastructure-as-code/provider"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func New(provider provider.IProvider) *Terraform {
	return &Terraform{
		*hashicorp.New(".tf"),
		provider,
	}
}

type Terraform struct {
	hashicorp.Hashicorp
	provider provider.IProvider
}

func (tf Terraform) MakeTerraform(body *hclwrite.Body) (*hclwrite.Block, *hclwrite.Body) {
	tfBlock := body.AppendNewBlock("terraform", nil)
	tfBody := tfBlock.Body()

	pvBlock := tfBody.AppendNewBlock("required_providers", nil)
	pvBody := pvBlock.Body()

	tf.provider.MakeTerraform(pvBody)
	return tfBlock, tfBody
}

func (tf Terraform) MakeProvider(body *hclwrite.Body) (*hclwrite.Block, *hclwrite.Body) {
	body.AppendNewline()

	pvBlock, pvBody := tf.provider.MakeProvider(body)

	return pvBlock, pvBody
}

func (tf Terraform) MakeResource(body *hclwrite.Body, solution provider.ISolution) (*hclwrite.Block, *hclwrite.Body) {
	body.AppendNewline()
	resBlock, resBody := solution.MakeResource(body)
	return resBlock, resBody
}
