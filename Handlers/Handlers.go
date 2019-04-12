package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/sixlib/go/Models"
)

/**
 * index 输出html模板页面
 */
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("views/index.html")
	if err != nil {
		fmt.Println("File reading error", err)
		// return
	}
	fmt.Fprintln(w, string(data))
}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		println(err)
		return
	}
	postMSG := models.SysInfos{
		models.SysInfo{Key: `GOOS`, Name: `系统类型`, Value: runtime.GOOS},
		models.SysInfo{Key: `GOARCH`, Name: `系统架构`, Value: runtime.GOARCH},
		models.SysInfo{Key: `GOMAXPROCS`, Name: `CPU 核数`, Value: strconv.Itoa(runtime.GOMAXPROCS(0))},
		models.SysInfo{Key: `Hostname`, Name: `电脑名称`, Value: name},
	}
	result(w, `200`, `success`, postMSG)

}

// todo 接口 handler
func TodoHandler(w http.ResponseWriter, r *http.Request) {
	todos := models.Todos{
		models.Todo{ID: 1, NAME: "TOM"},
		models.Todo{ID: 2, NAME: "JERRY"},
	}
	result(w, `200`, `success`, todos)
}
func TodoOfKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result(w, `200`, `success`, `todo key: `+vars["key"])
}

// 定义restful 统一返回数据结构
type Result struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 统一返回数据结构抽象方法
func result(w http.ResponseWriter, code string, msg string, data interface{}) {
	result := Result{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
