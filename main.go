package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

type Route struct {
	Src string `json:"src"`
	Dest string `json:"dest"`
}

type Config struct {
	Port int `json:"port"`
	Path string `json:"path"`
	Routes []*Route `json:"routes"`
}
var config Config
func main() {
	be, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	config = Config{}
	err = json.Unmarshal(be, &config)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", router)
	addr := fmt.Sprintf(":%d", config.Port)
	log.Println(fmt.Sprintf(
		"Listen on http://%s:%s",
		"0.0.0.0",
		addr),
	)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

func router(w http.ResponseWriter, r *http.Request) {
	for _, route := range config.Routes {
		// 编译正则
		rp := regexp.MustCompile(route.Src)
		// 路由匹配成功
		if rp.MatchString(r.RequestURI) {
			// 找到分组
			s := rp.FindStringSubmatch(r.RequestURI)
			s = s[1:]
			path := route.Dest
			for i := 0; i < len(s); i ++ {
				v := s[i]
				// 替换路径
				idx := fmt.Sprintf("$%d", i + 1)
				path = strings.Replace(path, idx, v,-1)
			}
			// 输出文件
			http.ServeFile(w, r, filepath.Join(config.Path, path))
			break
		}
	}
}