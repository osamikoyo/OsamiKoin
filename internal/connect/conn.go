package connect

import (
	"log/slog"
	"net"
	"os"
	"sync"

	"osamikoin/cmd/client/tool"
)

func Server(ch chan string, cher chan error) {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	listner, err := net.Listen("tcp", "/client")
	if err != nil {
		cher <- err
	}
	defer listner.Close()

	var wgs sync.WaitGroup


	wgs.Add(1)
	go tool.ConnectionRouting(ch, cher, listner, &wgs)

	select {
	case res := <- cher:
		loger.Error(res.Error())
	case res := <- ch:
		loger.Info(res)
		wgs.Wait()
	}
	
}
