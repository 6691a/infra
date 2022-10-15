package hashicorp

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Hashicorp struct {
	// 파일 확장자
	fileExtension string
	File          *hclwrite.File
	Body          *hclwrite.Body
}

func New(fileExtension string) *Hashicorp {
	file := *hclwrite.NewEmptyFile()
	body := file.Body()

	return &Hashicorp{
		fileExtension: fileExtension,
		File:          &file,
		Body:          body,
	}
}

//func (hc Hashicorp) NewBlock(name string, labels []string) (*hclwrite.Block, *hclwrite.Body) {
//	block := hc.Body.AppendNewBlock(name, []string{})
//	return block, block.Body()
//}

func (hc Hashicorp) CreateFile(fileName string) {
	fileName = fmt.Sprintf("%s%s", fileName, hc.fileExtension)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	file.Write(hc.File.Bytes())
}
