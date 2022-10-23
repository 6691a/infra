package main

import "github.com/6691a/infra/internal/ssh"

func main() {
	//vpc.ExecuteTemplate("tf.var")
	s := ssh.NewSSH("3.36.127.65", "ubuntu", "./mysql.pem", 22)
	s.Connect()
	s.Run("ls -al")
	s.Close()
}
