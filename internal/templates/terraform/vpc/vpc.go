package vpc

import (
	"os"
	"text/template"
)

type vpc struct {
	Name            string
	CidrBlock       string
	TagName         string
	InstanceTenancy string
}

func NewVPC(cidrBlock string, name string, instanceTenancy string, tagName string) *vpc {
	switch instanceTenancy {
	case "dedicated":
		return &vpc{CidrBlock: cidrBlock, Name: name, InstanceTenancy: "dedicated", TagName: tagName}
	default:
		return &vpc{CidrBlock: cidrBlock, Name: name, InstanceTenancy: "default", TagName: tagName}
	}
}

//func ReadTemplate() *template.Template {
//	//return template.Must(template.ParseFiles("internal/templates/terraform/vpc/vpc.tf"))
//}

func ExecuteTemplate(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	tmp, err := template.New("").ParseFiles("internal/templates/terraform/vpc/vpc.tf")
	if err != nil {
		panic(err)
	}
	tmp.ExecuteTemplate(file, "vpc.tf", vpc{CidrBlock: "123", Name: "test"})

}
