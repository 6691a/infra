package storages

type S3 struct {
	Region    string
	accesskey string
	secretkey string
	client    *s3.Client
}

func NewS3(region string, accesskey string, secretkey string) *S3 {
	return &S3{
		Region:    region,
		accesskey: accesskey,
		secretkey: secretkey,
	}
}
