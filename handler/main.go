package handler

import (
	"reflect"
	"time"
)

type Handler struct {
	CmdChan chan map[string]interface{}
	DoFunc  map[string]map[int][]func(event MsgEvent)
}

func (handler *Handler) AddOption(Cmd string, RoomId int, Do func(event MsgEvent)) {
	if _, ok := handler.DoFunc[Cmd]; !ok {
		handler.DoFunc[Cmd] = make(map[int][]func(event MsgEvent))
	}
	if _, ok := handler.DoFunc[Cmd][RoomId]; !ok {
		handler.DoFunc[Cmd][RoomId] = make([]func(event MsgEvent), 0)
	}
	handler.DoFunc[Cmd][RoomId] = append(handler.DoFunc[Cmd][RoomId], Do)
}

func (handler *Handler) CmdHandler() {
	for {
		select {
		case msg, ok := <-handler.CmdChan:
			if ok {
				setFunc := reflect.ValueOf(&Handler{}).MethodByName("Set" + CmdName[msg["cmd"].(string)])
				if setFunc.IsValid() {
					res := setFunc.Call([]reflect.Value{reflect.ValueOf(msg)})
					msgEvent := res[0].Interface().(MsgEvent)
					if !(msgEvent.Cmd == "" || msgEvent.RoomId == 0) {
						if _, ok := handler.DoFunc[msg["cmd"].(string)]; ok {
							if _, ok := handler.DoFunc[msg["cmd"].(string)][msgEvent.RoomId]; ok {
								for _, v := range handler.DoFunc[msg["cmd"].(string)][msgEvent.RoomId] {
									go v(msgEvent)
								}
							}
						}
					}
				}
			}
		default:
			time.Sleep(10 * time.Microsecond)
			continue
		}

	}
}
