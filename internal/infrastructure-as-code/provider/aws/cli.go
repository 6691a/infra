package aws

import (
	"fmt"
	"os"
	"path/filepath"
)

type Cli struct {
	AccessKey string
	SecretKey string
	path      string
	fileName  string
	output    string
	AWSProvider
}

// 현재 자격증명 확인
// aws sts get-caller-identity

func (cli Cli) New(provider *AWSProvider, accessKey string, secretKey string) *Cli {
	homeDir, _ := os.UserHomeDir()
	path, _ := filepath.Abs(homeDir + "/.aws")
	return &Cli{
		accessKey,
		secretKey,
		path,
		"config",
		"json",
		*provider,
	}
}

func (cli Cli) CreateConfigFile() {
	filepath := cli.path + "/" + cli.fileName
	err := os.MkdirAll(cli.path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	//if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		defer file.Close()
		return
	}

	fmt.Fprintf(file,
		"[default]\n"+
			"aws_access_key_id = %s\n"+
			"aws_secret_access_key = %s\n"+
			"region = %s\n"+
			"output = %s\n",
		cli.AccessKey, cli.SecretKey, cli.AWSProvider.region, cli.output)
	//} else {
	//	return
	//}
}
