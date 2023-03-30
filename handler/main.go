package handler

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type Handler struct {
	funcId    int
	CmdChan   chan map[string]interface{}
	DoFunc    map[string]map[int][]func(event MsgEvent)
	funcNames map[string]string
}

type jsonCoder interface {
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
}

var JsonCoder jsonCoder

func (handler *Handler) Init() {
	handler.funcNames = make(map[string]string)
}

func (handler *Handler) AddOption(Cmd string, RoomId int, Do func(event MsgEvent), funcName ...string) {
	if _, ok := handler.DoFunc[Cmd]; !ok {
		handler.DoFunc[Cmd] = make(map[int][]func(event MsgEvent))
	}
	if _, ok := handler.DoFunc[Cmd][RoomId]; !ok {
		handler.DoFunc[Cmd][RoomId] = make([]func(event MsgEvent), 0)
	}
	handler.DoFunc[Cmd][RoomId] = append(handler.DoFunc[Cmd][RoomId], Do)
	log.Debug("Add Option: ", Cmd, RoomId, funcName)
	handler.funcId++
	if len(funcName) > 0 {
		handler.funcNames[fmt.Sprintf("%p", Do)] = funcName[0]
	} else {
		handler.funcNames[fmt.Sprintf("%p", Do)] = fmt.Sprintf("#%d", handler.funcId)
	}
}

func (handler *Handler) DelRoomOption(roomId int) {
	// TODO: 检查这里是否需要用sync.Map
	for k, v := range handler.DoFunc {
		if _, ok := v[roomId]; ok {
			delete(handler.DoFunc[k], roomId)
			log.Debug("Del Option: ", k, roomId)
		}
	}
}

func (handler *Handler) DelOption(name string) {
	for k, v := range handler.DoFunc {
		for k1, v1 := range v {
			for i, v2 := range v1 {
				if name == handler.funcNames[fmt.Sprintf("%p", v2)] {
					handler.DoFunc[k][k1] = append(handler.DoFunc[k][k1][:i], handler.DoFunc[k][k1][i+1:]...)
					if len(handler.DoFunc[k][k1]) == 0 {
						delete(handler.DoFunc[k], k1)
					}
					log.Debug("Del Option: ", k, k1, name)
				}
			}
		}
	}
}

func (handler *Handler) doHandler(f reflect.Value, msg map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("doHandler panic: %v, please make an issue to me!", err)
		}
	}()
	res := f.Call([]reflect.Value{reflect.ValueOf(msg)})
	msgEvent := res[0].Interface().(MsgEvent)
	// 执行函数
	if !(msgEvent.Cmd == "" || msgEvent.RoomId == 0) {
		for _, k := range []int{msgEvent.RoomId, 0} {
			for _, v := range handler.DoFunc[msg["cmd"].(string)][k] {
				if log.GetLevel() == log.DebugLevel {
					if name, ok := handler.funcNames[fmt.Sprintf("%p", v)]; ok {
						log.Debugf("distribute %v cmd to %v", msg["cmd"].(string), name)
					} else {
						log.Debugf("distribute %v cmd to %p", msg["cmd"].(string), v)
					}
				}
				go v(msgEvent)
			}
		}
	}
}

func (handler *Handler) CmdHandler() {
	for {
		select {
		case msg, ok := <-handler.CmdChan:
			if ok {
				// 处理命令存在
				if _, ok = handler.DoFunc[msg["cmd"].(string)]; ok {
					// 处理房间存在
					_, ok1 := handler.DoFunc[msg["cmd"].(string)][msg["RoomId"].(int)]
					// 0 为所以房间
					_, ok2 := handler.DoFunc[msg["cmd"].(string)][0]
					if ok1 || ok2 {
						setFunc := reflect.ValueOf(&Handler{}).MethodByName("Set" + CmdName[msg["cmd"].(string)])
						if setFunc.IsValid() {
							handler.doHandler(setFunc, msg)
						} else {
							log.Debug(CmdName[msg["cmd"].(string)] + " not found, please make an issue to me!")
						}
					}
				}
			}
		}
	}
}
