package aws

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type S3 struct {
	Bucket string
	Status bool
	Name   string
}

func (s3 S3) MakeResource(body *hclwrite.Body) (*hclwrite.Block, *hclwrite.Body) {
	s3Block := body.AppendNewBlock("resource", []string{"aws_s3_bucket", s3.Name})
	s3Body := s3Block.Body()
	s3Body.SetAttributeValue("bucket", cty.StringVal(s3.Bucket))
	if s3.Status {
		s3Body.SetAttributeValue("status", cty.StringVal("Enabled"))
	} else {
		s3Body.SetAttributeValue("status", cty.StringVal("Suspended"))
	}

	return s3Block, s3Body
}
