package aws

import (
	"github.com/6691a/infra/internal/infrastructure-as-code/provider"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type AWSProvider struct {
	provider.Provider
	region string
}

func New(version string, region string) *AWSProvider {
	return &AWSProvider{
		*provider.New("aws", version, "hashicorp/aws"),
		region,
	}
}

func (aws AWSProvider) MakeTerraform(body *hclwrite.Body) *hclwrite.Attribute {
	return body.SetAttributeValue(aws.Name, cty.ObjectVal(map[string]cty.Value{
		"version": cty.StringVal(aws.Version),
		"source":  cty.StringVal(aws.Source),
	}))

}

func (aws AWSProvider) MakeProvider(body *hclwrite.Body) (*hclwrite.Block, *hclwrite.Body) {
	pvBlock := body.AppendNewBlock("provider", []string{aws.Name})
	pvBody := pvBlock.Body()
	pvBody.SetAttributeValue("region", cty.StringVal(aws.region))
	return pvBlock, pvBody
}

func (aws AWSProvider) MakeResource(body *hclwrite.Body, solution provider.ISolution) (*hclwrite.Block, *hclwrite.Body) {
	var block *hclwrite.Block

	return block, block.Body()
}
