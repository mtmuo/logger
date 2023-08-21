package logger

import (
	"fmt"
	"os"
)

type Client interface {
	Send(subject, context string) error
}

type Notice struct {
	hostName string
	exec     Exec
	clients  []Client
}

func (n *Notice) Send(subject, context string) {
	if len(n.clients) == 0 {
		return
	}

	for _, client := range n.clients {
		_ = client.Send(
			fmt.Sprintf("%s from %s", subject, n.exec.AppName),
			fmt.Sprintf(`
appName: %s run %s 
%s
`,
				n.exec.AppName,
				n.hostName,
				context,
			),
		)
	}
}

func NewNotice() *Notice {
	hostName, _ := os.Hostname()
	return &Notice{
		hostName: hostName,
		exec:     Executable(),
	}
}
