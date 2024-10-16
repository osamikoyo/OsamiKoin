package route

import (
	"strconv"
	"strings"

	"osamikoin/internal/route/handler"
)

func CommandRoute(ch chan string) {
	res := <-ch

	arguments := strings.Split(res, " ")
	switch arguments[0] {
	case "register":
		handler.RegisterCLI()
	case "mysending":
		handler.GetSending(arguments[1])
	case "send":
		count,_ := strconv.Atoi(arguments[3]) 
		handler.Send(arguments[1], arguments[2], count)
	}

}
