package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"kedou/config"
	"kedou/utils"
	kedouws "kedou/ws"
	"math/rand"
)

func DataDistribution(ws *websocket.Conn, c config.Data) {
	//switch c.Type {
	//case "welcome":
	//	fmt.Println("链接成功:改名:")
	//	smg := config.Data{
	//		Type:     "update",
	//		Angle:    "0",
	//		Momentum: "0.850",
	//		Name:     "你你你",
	//		Sex:      "1",
	//		X:        "1428.2",
	//		Y:        "309.9",
	//		Icon:     "/images/default.png",
	//	}
	//	//jsondata, _ := json.Marshal(smg)
	//	//fmt.Println(string(jsondata))
	//	//err := Ws.WriteMessage(websocket.TextMessage, jsondata)
	//	err := SendMessage(ws, smg)
	//	if err == nil {
	//		fmt.Println("发送成功")
	//	} else {
	//		fmt.Println("发送失败")
	//	}
	//	//SendMessage(smg)
	//case "update":
	//	randomNumber := rand.Intn(2)
	//	smg := config.Data{
	//		Type:     "update",
	//		Angle:    "0",
	//		Momentum: "0.850",
	//		Name:     "你你你",
	//		Sex:      randomNumber,
	//		X:        "856",
	//		Y:        "393.5",
	//		Icon:     "/images/default.png",
	//	}
	//	err := SendMessage(ws, smg)
	//	if err == nil {
	//		fmt.Println("发送成功")
	//	} else {
	//		fmt.Println("发送失败")
	//	}
	//	fmt.Println(c.Type)
	//default:
	//	fmt.Println(c.Type)
	//}
}

func handleMessage(kedou *kedouws.KeDou) {
	//fmt.Println(kedou.WsData.Type, kedou.WsData.Name)
	switch kedou.WsData.Type {

	case "welcome":
		randomNumber := rand.Intn(2)
		smg := &config.Data{
			Type:     "update",
			Angle:    "0",
			Momentum: "0",
			Name:     kedou.LocalName,
			Sex:      randomNumber,
			X:        "856",
			Y:        "393.5",
			Icon:     "/images/default.png",
		}
		kedou.SendMessage(smg)

	case "update":
		if kedou.WsData.Name == kedou.FollowName {

			randomNumber := rand.Intn(2)
			smg := &config.Data{
				Type:     "update",
				Angle:    kedou.WsData.Angle,
				Momentum: "0",
				Name:     kedou.LocalName,
				Sex:      randomNumber,
				X:        kedou.WsData.X.(float64) + kedou.XCoordinateOffset,
				Y:        kedou.WsData.Y.(float64) + kedou.XCoordinateOffset,
				Icon:     "/images/default.png",
			}
			kedou.SendMessage(smg)
			//fmt.Println("我看看你是谁")
		}

	}

	//err := (ws, smg)
}
func main() {
	N := 5 // 创建 N 个 KeDou 实例

	// 创建一个切片来存储 KeDou 实例
	keDouList := make([]*kedouws.KeDou, 0)

	for i := 0; i < N; i++ {

		name := utils.GenerateRandomName()
		fmt.Println("开始循环")
		keDou, _ := kedouws.NewKeDou("guanren", float64(10+i), float64(5+i), 3, name)
		keDouList = append(keDouList, keDou)
		fmt.Println("本次name:", name)

		go keDouList[i].MessageCallback(handleMessage)
	}

	fmt.Println("添加完毕")

	// 等待所有 KeDou 实例执行完毕
	for _, keDou := range keDouList {
		keDou.Close()
	}
}
