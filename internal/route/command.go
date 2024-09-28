package route

import "osamikoin/internal/route/handler"

func CommandRoute(ch chan string) {
	res := <-ch
	switch res{
	case "register":
		handler.RegisterCLI()
	case "mysending":
		
	case "send":
		
	}

} 