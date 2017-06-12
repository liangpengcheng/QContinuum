package network

import (
	"github.com/liangpengcheng/qcontinuum/base"
)

// MsgCallback 消息处理函数
type MsgCallback func(msg *Message)

// EventCallback 时间处理
type EventCallback func(event *Event)

// ProcFunction 回掉的程序
type ProcFunction func()

// Message 消息链
type Message struct {
	Peer *ClientPeer
	Head MessageHead
	Body []byte
}

// Event 自定义事件
type Event struct {
	ID    int32       //eventid
	Param string      //param
	Peer  *ClientPeer //事件中的peer,可以为nil
}

var (
	// ExitEvent 退出
	ExitEvent int32 = 1
	// AddEvent 加入玩家
	AddEvent int32 = 2
	// RemoveEvent 删除玩家
	RemoveEvent int32 = 3
)

// Processor 消息处理器
type Processor struct {
	MessageChan   chan *Message
	SendChan      chan *Message
	EventChan     chan *Event
	FuncChan      chan ProcFunction
	CallbackMap   map[int32]MsgCallback
	EventCallback map[int32]EventCallback
	// ImmediateMode 立即回调消息，如果想要线程安全，必须设置为false，默认为false
	ImmediateMode bool
}

// NewProcessor 新建处理器，包含初始化操作
func NewProcessor() *Processor {
	p := &Processor{
		MessageChan:   make(chan *Message, 1024),
		SendChan:      make(chan *Message, 1024),
		EventChan:     make(chan *Event, 1024),
		FuncChan:      make(chan ProcFunction, 64),
		EventCallback: make(map[int32]EventCallback),
		CallbackMap:   make(map[int32]MsgCallback),
		ImmediateMode: false,
	}
	return p
}

// AddCallback 设置回调
func (p *Processor) addCallback(id int32, callback MsgCallback) {
	p.CallbackMap[id] = callback
}

// AddCallback 设置回调
func (p *Processor) AddCallback(id int32, callback MsgCallback) {
	if id != 0 {
		p.addCallback(id, callback)
	}
}

// RemoveCallback 删除回调
func (p *Processor) RemoveCallback(id int32) {
	delete(p.CallbackMap, id)
}

// AddEventCallback 事件处理函数注册
func (p *Processor) AddEventCallback(id int32, callback EventCallback) {
	p.EventCallback[id] = callback
}

// RemoveEventCallback 删除事件处理回调
func (p *Processor) RemoveEventCallback(id int32) {
	delete(p.EventCallback, id)
}
func (p *Processor) send() {
	for {
		select {
		case msg := <-p.SendChan:
			if msg.Peer.Connection != nil {
				msg.Peer.Connection.Write(msg.Body)
			}
		}
	}
}

// StartProcess 开始处理信息
// 只有调用了这个借口，处理器才会处理实际的信息，以及实际发送消息
func (p *Processor) StartProcess() {
	go p.send()
	base.LogInfo("processor is starting ")
	for {
		select {
		case msg := <-p.MessageChan:
			if cb, ok := p.CallbackMap[msg.Head.ID]; ok {
				cb(msg)
			} else {
				base.LogWarn("can't find callback(%d)", msg.Head.ID)
			}
		case event := <-p.EventChan:
			if event.ID == ExitEvent {
				base.LogInfo("Processor exit : %s", event.Param)
				return
			}
			if cb, ok := p.EventCallback[event.ID]; ok {
				cb(event)
			}
		case f := <-p.FuncChan:
			f()
		}

	}
}
