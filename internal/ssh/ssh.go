package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
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
	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			s.readCert(s.Cert),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port), config)
	if err != nil {
		return err
	}

	session, err := client.NewSession()
	if err != nil {
		fmt.Println(err)
		client.Close()
		return err
	}
	s.session = session
	s.client = client
	return nil
}

func (s *SSH) Close() error {
	s.session.Close()
	s.client.Close()

	if s.session != nil {
		return s.session.Close()
	}
	if s.client != nil {
		return s.client.Close()
	}
	return nil
}

// Run a command on the remote server
func (s *SSH) Run(cmd string) (string, error) {
	out, err := s.session.CombinedOutput(cmd)
	fmt.Println(string(out))
	if err != nil {
		return "", err
	}
	return string(out), nil
}
