package handler

import (
	"github.com/FishZe/go-bili-chat/v2/events"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

// FuncTable
// 空结构体不占空间，作为hashSet使用
// 函数指针内存唯一
type FuncTable = map[*events.BLiveEventHandler]struct{}

type RoomTable = map[int]FuncTable

type CmdTable = map[string]RoomTable

type Handler struct {
	CmdChan chan *events.BLiveEvent
	DoFunc  CmdTable
	//函数反查表
	FuncPath map[*events.BLiveEventHandler]int
}

func (handler *Handler) AddOption(RoomId int, Do events.BLiveEventHandler) *events.BLiveEventHandler {
	cmd := Do.Cmd()
	if _, ok := handler.DoFunc[cmd]; !ok {
		handler.DoFunc[cmd] = make(RoomTable)
	}
	if _, ok := handler.DoFunc[cmd][RoomId]; !ok {
		handler.DoFunc[cmd][RoomId] = make(FuncTable)
	}
	addr := &Do
	//将函数添加进Set
	handler.DoFunc[cmd][RoomId][addr] = struct{}{}
	//函数反查表
	handler.FuncPath[addr] = RoomId
	log.Debug("Add Option: ", cmd, RoomId)
	return addr
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

func (handler *Handler) DelOption(p *events.BLiveEventHandler) {
	if p != nil {
		roomId := handler.FuncPath[p]
		delete(handler.DoFunc[(*p).Cmd()][roomId], p)
		delete(handler.FuncPath, p)
	}
}

func (handler *Handler) doHandler(msg *events.BLiveEvent) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("do %v Handler panic: %v, please issue to me!", msg.Cmd, err)
		}
	}()
	// 执行函数
	if msg.Cmd != "" && msg.RoomId != 0 {
		//房间号分发
		for t := range handler.DoFunc[msg.Cmd][msg.RoomId] {
			log.Debugf("distribute %v cmd", msg.Cmd)
			go func(t *events.BLiveEventHandler, msgEvent *events.BLiveEvent) {
				defer func() {
					if err := recover(); err != nil {
						log.Errorf("do Handler panic: %v, please issue to me!", err)
						debug.PrintStack()
					}
				}()
				(*t).On(msgEvent)
			}(t, msg)
		}
		//全局分发
		for t := range handler.DoFunc[msg.Cmd][0] {
			log.Debugf("distribute %v cmd", msg.Cmd)
			go func(t *events.BLiveEventHandler, msgEvent *events.BLiveEvent) {
				defer func() {
					if err := recover(); err != nil {
						log.Errorf("do Handler panic: %v, please issue to me!", err)
						debug.PrintStack()
					}
				}()
				(*t).On(msgEvent)
			}(t, msg)
		}
	}
}

func (handler *Handler) CmdHandler() {
	for {
		select {
		case msg, ok := <-handler.CmdChan:
			if ok {
				cmd := msg.Cmd
				//命令是字符串
				// 处理命令存在
				if dict, ok := handler.DoFunc[cmd]; ok {
					// 处理房间存在
					_, ok1 := dict[msg.RoomId]
					// 0 为所有房间
					_, ok2 := dict[0]
					if ok1 || ok2 {
						handler.doHandler(msg)
					}
				}
			}
		}
	}
}
