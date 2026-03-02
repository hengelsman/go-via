package vmware

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type Executor struct {
    Host     string
    User     string
    Password string
}

func NewExecutor(host, user, password string) (*Executor, error) {
    if host == "" || user == "" {
        return nil, fmt.Errorf("invalid ssh executor params")
    }
    return &Executor{Host: host, User: user, Password: password}, nil
}

func (e *Executor) Run(cmd []string) (string, error) {
    config := &ssh.ClientConfig{
        User:            e.User,
        Auth:            []ssh.AuthMethod{ssh.Password(e.Password)},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Timeout:         10 * time.Second,
    }

    addr := e.Host
    if !strings.Contains(addr, ":") {
        addr = addr + ":22"
    }

    client, err := ssh.Dial("tcp", addr, config)
    if err != nil {
        return "", err
    }
    defer client.Close()

    session, err := client.NewSession()
    if err != nil {
        return "", err
    }
    defer session.Close()

    out, err := session.CombinedOutput(strings.Join(cmd, " "))
    return string(out), err
}
