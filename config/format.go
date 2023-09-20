package config

type Data struct {
	Type     string      `json:"type"`
	X        interface{} `json:"x"`
	Y        interface{} `json:"y"`
	Angle    interface{} `json:"angle"`
	Momentum interface{} `json:"momentum"`
	Sex      interface{} `json:"sex"`
	Icon     string      `json:"icon"`
	Name     string      `json:"name"`
}
