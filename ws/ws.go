package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"kedou/config"
	"sync"
	"time"
)

type KeDou struct {
	WsData            *config.Data
	Ws                *websocket.Conn
	LocalName         string
	FollowName        string
	TargetX           float64
	TargetY           float64
	XCoordinateOffset float64
	YCoordinateOffset float64
	HeartbeatTimer    int
	mux               sync.Mutex
}
type MessageHandlerFunc func(w *KeDou)

func NewKeDou(Follow string, X float64, Y float64, timer int, Local string) (*KeDou, error) {
	//websocket.DefaultDialer.Dial("ws://kedou.workerman.net:8280", nil)
	conn, _, err := websocket.DefaultDialer.Dial("ws://kedou.workerman.net:8280", nil)

	if err != nil {
		fmt.Println("无法链接到服务器", err)
		return nil, err
	}

	return &KeDou{
		Ws:                conn,
		FollowName:        Follow,
		XCoordinateOffset: X,
		YCoordinateOffset: Y,
		HeartbeatTimer:    timer,
		LocalName:         Local,
	}, nil
}

func (k *KeDou) MessageCallback(handler MessageHandlerFunc) (config.Data, error) {
	//var message config.Data
	for {
		err := k.Ws.ReadJSON(&k.WsData)
		if err != nil {
			fmt.Println("读取消息失败：", err)

			//return k.WsData, err
		}

		if handler != nil {
			go handler(k)

		}
	}

	//return k.WsData, nil
}

func (k *KeDou) Close() {
	k.Ws.Close()

}
func (k *KeDou) HeartbeatPacket() {
	timer := time.NewTicker(time.Second * time.Duration(k.HeartbeatTimer))
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			// 定时器触发时执行的任务
			k.SendMessage(k.WsData)
			//fmt.Println("定时任务执行了")
		}
	}

}
func (k *KeDou) SendMessage(t *config.Data) error {
	k.mux.Lock()
	defer k.mux.Unlock()
	return k.Ws.WriteJSON(t)
}
