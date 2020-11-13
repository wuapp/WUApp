package wua

/*
#include <stdlib.h>
*/
import "C"

type Message struct {
	Id   string `json:"id"` //for receive request from desktop app
	Url  string `json:"url"`
	Data string `json:"data"`
	//Success string `json:"success"`
	//Error   string `json:"error"`
}
