package main

import "github.com/6691a/infra/internal/templates/terraform/vpc"

func main() {
	vpc.ExecuteTemplate("tf.var")
}
