package service

import "fmt"

type Mock struct {
}

type MockModel struct {
	Source string `json:"source"`
	Data   string `json:"data"`
}

func (m Mock) InsertMock(model MockModel) {
	fmt.Println(model)
	fmt.Println(model.Data)
	fmt.Println(model.Source)
}
