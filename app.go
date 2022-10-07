package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type Person struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
// 改变接收值为Person结构
// func (a *App) Greet(p Person) string {
// 	return fmt.Sprintf("Hello %s, Age: %d. It's show time!", p.Name, p.Age)
// }

type RandomImage struct {
	Message string
	Status  string
}

type AllBreads struct {
	Message map[string]map[string][]string
	Status  string
}

type ImageByBread struct {
	Message []string
	Status  string
}

// 获取随机图片
func (a *App) GetRandomInageUrl() string {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data RandomImage
	json.Unmarshal(responseData, &data)
	return data.Message
}

// 获取列表
func (a *App) GetBreadList() []string {
	var breeds []string

	response, err := http.Get("https://dog.ceo/api/breeds/list/all")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data AllBreads
	json.Unmarshal(responseData, &data)
	//循环添加带数组中
	for k := range data.Message {
		breeds = append(breeds, k)
	}
	//排序
	sort.Strings(breeds)

	return breeds
}

// 获取指定图片
func (a *App) GetImageUrlByBreed(breed string) []string {
	url := fmt.Sprintf("%s%s%s%s", "https://dog.ceo/api/", "breed/", breed, "/images")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data ImageByBread
	json.Unmarshal(responseData, &data)

	return data.Message
}

// ===========窗口全屏==========
func (a *App) WindowFull() {
	runtime.WindowFullscreen(a.ctx)
}

type Person1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 触发事件
func (a *App) EventsEmit2() {
	runtime.EventsEmit(a.ctx, "click", Person1{
		"tiry",
		22,
	})
}
