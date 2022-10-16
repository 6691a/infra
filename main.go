package main

import "github.com/6691a/infra/internal/infrastructure-as-code/provider/aws"

func main() {
	provider := aws.AWSProvider{}.New("4.34.0", "ap-northeast-2")
	cli := aws.Cli{}.New(provider, "12311", "aaa")
	cli.CreateConfigFile()
	//tf := terraform.Terraform{}.New(provider)
	//
	//tf.MakeTerraform(tf.Body)
	//tf.MakeProvider(tf.Body)
	//tag := aws.Tag{}.New("test_tag", "test_dev")
	//s3 := aws.S3{}.New("test_golang_s3_bucket", true, "main", tag)
	//
	//tf.MakeResource(tf.Body, s3)
	//tf.CreateFile("test")

}
