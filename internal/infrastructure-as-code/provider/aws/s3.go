package aws

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type S3 struct {
	Bucket string
	Status bool
	Name   string
	Tag
}

func (s3 S3) New(bucket string, status bool, name string, tag *Tag) *S3 {
	return &S3{
		Bucket: bucket,
		Status: status,
		Name:   name,
		Tag:    *tag,
	}
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
	if s3.Tag != (Tag{}) {
		tagBlock := s3Body.AppendNewBlock("tags", nil)
		tagBody := tagBlock.Body()
		tagBody.SetAttributeValue("Name", cty.StringVal(s3.Tag.Name))
		tagBody.SetAttributeValue("Environment", cty.StringVal(s3.Tag.Env))
	}

	return s3Block, s3Body
}
