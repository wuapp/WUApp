package wuapp

import "C"
import "encoding/json"

type Message struct {
	Url     string      `json:"url"`
	Data    string `json:"data"`
	Success string      `json:"success"`
	Error   string      `json:"error"`
}

var clientHandler = "wuapp.receive"

// Service is to add a backend service for frontend to request.
// params:
//	url - the url act as an unique identifier of a service, for example, "user/login", "blog/get/:id".
//	handler - the function that handle the client request.
func Service(url string, action func(*Context)) {
	route := new(route)
	route.action = action
	parseRoute(url, route)
}

func Request(msg Message)  {
	s, err := json.Marshal(msg)
	if err != nil {
		return
	}

	invokeJavascript(clientHandler + "("+ string(s) +")")
	return
}

func receive(msg *C.char)  {
	goMsg :=C.GoString(msg)
	//Log("ClientHandler:", message)
	message := new(Message)
	err := json.Unmarshal([]byte(goMsg), message)
	if err != nil {
		//Log("unmarshal error:", err)
		return
	}

	action,params := dispatch(message.Url)
	ctx := &Context{message: message,params: params}
	if action != nil {
		action(ctx)
	} else {
		ctx.Error("Function not found ")
	}

}

