package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

//var commonUrl string = "http://106.54.84.172:3000/"
//var commonUrl string = "https://gitea.com/"

func httpDo(subUrl string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

//获取 GET 请求的数据
func GetData(url string) string {
	client := &http.Client{}
	resp, err := client.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	return string(body)
}

//通过反射，将结构体打印
func SmartPrint(i interface{}){
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType :=reflect.TypeOf(i)
	for i:=0;i<vValue.NumField();i++{
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("\n获取到的对象数据为：")
	for k,v :=range kv{
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
	}
}

