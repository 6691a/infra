package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	// create new file on system
	tfFile, err := os.Create("servicelist.tf")
	if err != nil {
		fmt.Println(err)
		return
	}

	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()
	tfBlock := rootBody.AppendNewBlock("terraform", nil)
	tfBlockBody := tfBlock.Body()
	reqProvsBlock := tfBlockBody.AppendNewBlock("required_providers",
		nil)

	reqProvsBlockBody := reqProvsBlock.Body()

	reqProvsBlockBody.SetAttributeValue("pagerduty", cty.ObjectVal(map[string]cty.Value{
		"source":  cty.StringVal("PagerDuty/pagerduty"),
		"version": cty.StringVal("2.3.0"),
	}))
	rootBody.AppendNewline()

	tfFile.Write(f.Bytes())
}
