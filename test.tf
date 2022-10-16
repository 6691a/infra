terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.34.0"
    }
  }
}

provider "aws" {
  region = "ap-northeast-2"
}

resource "aws_s3_bucket" "main" {
  bucket = "test_golang_s3_bucket"
  status = "Enabled"
  tags {
    Name        = "test_tag"
    Environment = "test_dev"
  }
}
