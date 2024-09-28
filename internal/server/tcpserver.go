package server

import (
	"log/slog"
	"net"
	"os"
	"sync"

	"osamikoin/cmd/client/tool"
)

func TCPserver(cher chan error, ch chan string) {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	listner, err := net.Listen("tcp", "/conn")
	if err != nil {
		cher <- err
	}
	defer listner.Close()

	var wgs sync.WaitGroup

	wgs.Add(1)
	go tool.ConnectionRouting(ch, cher, listner, &wgs)

	select {
	case res := <-cher:
		loger.Error(res.Error())
	case res := <-ch:
		loger.Info(res)
		wgs.Wait()
	}
}
