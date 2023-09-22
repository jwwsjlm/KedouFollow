package main

import (
	"github.com/gorilla/websocket"
	"kedou/config"
	"kedou/utils"
	kedouws "kedou/ws"
	"math"
	"math/rand"
	"time"
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
			X:        "0",
			Y:        "0",
			Icon:     "/images/default.png",
		}
		kedou.SendMessage(smg)

	case "update":

		if kedou.WsData.Name == kedou.FollowName {
			kedou.TargetX = kedou.WsData.X.(float64)
			kedou.TargetY = kedou.WsData.Y.(float64)
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
func CircleAround(keDou *kedouws.KeDou) {

	radius := 150.0       // 圆的半径
	angle := 0.0          //旋转角度
	angleIncrement := 3.0 // 每次更新的角度增量
	for {
		if keDou.WsData != nil {
			radian := angle * (math.Pi / 180.0)

			x := keDou.TargetX + radius*math.Cos(radian)
			y := keDou.TargetY + radius*math.Sin(radian)
			randomNumber := rand.Intn(2) //随机性别
			smg := &config.Data{
				Type:     "update",
				Angle:    "5",
				Momentum: "0",
				Name:     keDou.LocalName,
				Sex:      randomNumber,
				X:        x,
				Y:        y,
				Icon:     "/images/default.png",
			}
			keDou.SendMessage(smg)

			//fmt.Printf("Point: (%f, %f)\n", x, y)

			angle += angleIncrement

			time.Sleep(time.Millisecond * 10) // 控制每次更新的时间间隔
		}

	}
}
func main() {

	name := utils.GenerateRandomName()

	keDou, err := kedouws.NewKeDou("guanren", float64(10), float64(5), 3, name)
	if err == nil {
		go CircleAround(keDou)
		keDou.MessageCallback(handleMessage)

	}

	select {}

}
