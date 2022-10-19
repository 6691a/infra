resource "aws_vpc" "{{.Name}}" {
 cidr_block = "{{.CidrBlock}}"
 enable_dns_hostnames = true
 enable_dns_support = true
 instance_tenancy = "{{.InstanceTenancy}}"
 tags = {
    Name = "{{.TagName}}"
 }
}
