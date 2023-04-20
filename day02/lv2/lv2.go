package main

import (
	"fmt"
	"reflect"
)

// 定义发声接口，所有实现了playMusic方法的结构体均实现了该接口
type sounder interface {
	playMusic()
}

// 定义真无线耳机结构体
type tws struct {
	Manufacturers        string
	Price                int
	NoiseCancelSupported bool
}

func (t *tws) playMusic() {
	fmt.Printf("类型：%v\t值：%v\n", reflect.TypeOf(*t), *t)
}

// 定义音响结构体
type audio struct {
	Manufacturers     string
	Price             int
	WirelessSupported bool
}

func (a *audio) playMusic() {
	fmt.Printf("类型：%v\t值：%v\n", reflect.TypeOf(*a), *a)
}

func main() {
	t1 := tws{
		"苹果",
		1799,
		true,
	}
	a1 := audio{
		"BOSE",
		2499,
		true,
	}
	t1.playMusic()
	a1.playMusic()

}
