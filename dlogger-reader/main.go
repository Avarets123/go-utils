package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type CmdReaders interface {
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
}

func main() {

	socket := os.Getenv("SSH_AUTH_SOCK")
	conn, err := net.Dial("unix", socket)
	if err != nil {
		panic(err)
	}

	ipWithPort := ""

	agCl := agent.NewClient(conn)

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agCl.Signers),
		},
		Timeout:         10 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshClient, err := ssh.Dial("tcp", ipWithPort, config)
	if err != nil {
		panic(err)
	}

	ses, _ := sshClient.NewSession()

	stdout, err := ses.StdoutPipe()
	if err != nil {
		panic(err)
	}

	stderr, err := ses.StderrPipe()
	if err != nil {
		panic(err)
	}

	ses.Start("docker logs game-runner -f")

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Fprint(os.Stderr, scanner.Text())
		}

	}()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	ses.Wait()
}

func ReadFromLocal() {

	containerName := "db_main"

	cmd := exec.Command("docker", "logs", "-f", containerName)

	readStdoutAndStdErr(cmd)

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	cmd.Wait()

}

func readStdoutAndStdErr(cmd CmdReaders) {
	sdtout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	sdterr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	go func() {
		scanner := bufio.NewScanner(sdterr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

	}()

	go func() {
		scanner := bufio.NewScanner(sdtout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
}
