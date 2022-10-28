package ssh

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

type SSH struct {
	IP      string
	User    string
	Port    int
	Cert    string
	session *ssh.Session
	client  *ssh.Client
	sftp    *sftp.Client
}

func NewSSH(ip string, user string, cert string, port int) *SSH {
	return &SSH{
		IP:   ip,
		User: user,
		Port: port,
		Cert: cert,
	}
}
func (s *SSH) readCert(file string) ssh.AuthMethod {
	cert, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	key, err := ssh.ParsePrivateKey(cert)
	if err != nil {
		log.Fatal(err)
	}
	return ssh.PublicKeys(key)

}

func (s *SSH) Connect() error {
	// setting ssh config
	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			s.readCert(s.Cert),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connect to ssh
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port), config)
	if err != nil {
		return err
	}

	// create session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println(err)
		client.Close()
		return err
	}

	// create sftp client
	sftp, err := sftp.NewClient(client)
	if err != nil {
		fmt.Println(err)
		client.Close()
		return err
	}

	s.client = client
	s.session = session
	s.sftp = sftp
	return nil
}

func (s *SSH) Close() error {
	if s.session != nil {
		return s.session.Close()
	}
	if s.client != nil {
		return s.client.Close()
	}
	if s.sftp != nil {
		return s.sftp.Close()
	}
	return nil
}

// Run command on remote server
func (s *SSH) Command(cmd string) (string, error) {
	out, err := s.session.CombinedOutput(cmd)
	fmt.Println(string(out))
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// upload file to remote server and print progress bar
// TODO: local file -> AWS S3
func (s *SSH) UploadWithProgressBar(localPath string, remotePath string) error {
	srcFile, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := s.sftp.Create(remotePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// get file size
	fileInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	// create progress bar
	bar := pb.Full.Start64(fileSize)
	barReader := bar.NewProxyReader(srcFile)

	// copy file
	_, err = io.Copy(dstFile, barReader)
	if err != nil {
		return err
	}

	bar.Finish()
	return nil
}
