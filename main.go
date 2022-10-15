package main

import (
	"github.com/6691a/infra/internal/infrastructure-as-code/provider/aws"
	"github.com/6691a/infra/internal/infrastructure-as-code/terraform"
)

func main() {
	tf := terraform.New(aws.New("4.34.0", "ap-northeast-2"))

	tf.MakeTerraform(tf.Body)
	tf.MakeProvider(tf.Body)
	tf.MakeResource(tf.Body, aws.S3{Bucket: "test_bucket", Status: true, Name: "main"})
	tf.CreateFile("test")

	//pv := provider.Provider{
	//	Name:    "aws",
	//	Source:  "hashicorp/aws",
	//	Version: "4.34.0",
	//}
	//tf.MakeTerraform(tf.Body, pv)
	//tf.MakeProvider(tf.Body, pv)
	//tf.CreateFile("test")

	// _, body := pk.MakePacker(">= 1.2.0, < 2.0.0")
	// plugins := []plugin.Plugin{
	// 	*plugin.New("myawesomecloud ", ">= 2.7.0", "github.com/azr/myawesomecloud"),
	// 	*plugin.New("happycloud ", ">= 1.1.3", "github.com/azr/happycloud"),
	// }
	// pk.MakePlugins(body, plugins)

	// pk.CreateFile("name")

}
