package main

import (
	"table"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/gate"
	"github.com/davyxu/golog"
)

func main() {

	golog.SetLevelByString("socket", "info")

	table.LoadServiceTable()

	// 后台与前台在两个线程
	backendPipe := cellnet.NewEventPipe()
	frontendPipe := cellnet.NewEventPipe()

	gate.StartBackendAcceptor(backendPipe, table.GetPeerAddress("svc->gate"))
	gate.StartClientAcceptor(frontendPipe, table.GetPeerAddress("client->gate"))

	backendPipe.Start()
	frontendPipe.Start()

	backendPipe.Wait()
	frontendPipe.Wait()
}
