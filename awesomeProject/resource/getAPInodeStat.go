package resource

import (
	model "awesomeProject/mode"
	"errors"
	"fmt"
	jsdata "github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetupUrl(url string) ([]byte, error) {
	var logger = log.Default()
	resp, err := http.Get(url)
	if err != nil {
		logger.Printf("get请求失败 error: %+v", err)
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		logger.Printf("读取Body失败 error: %+v", err)

	}

	return body, err1
}
func PastJson(data []byte) (map[string]interface{}, error) {
	js, err := jsdata.NewJson(data)
	if err != nil {
		return nil, errors.New("data is not meta")
	}
	//js.Get("data")
	return js.Map()

}

var (
	cst *time.Location
)

// CSTLayout China Standard Time Layout
//const CSTLayout = "2006-01-02 15:04:05"

func GetData(flagStr, job string) []string {
	//dayTime := time.Now().Add(-time.Hour * time.Duration(7*24)).Unix()
	//hourTime := time.Now().Add(-time.Hour * time.Duration(7*24)).Unix()
	var UpHost = make([]string, 0)
	//var test string
	url := fmt.Sprintf("%s/api/v1/query_range?query=%s&start=%v&end=%v&step=15s", model.C.Url.ApiUrl, flagStr, model.StartTime, model.EndTime)

	data, _ := GetupUrl(url)
	slics, err := PastJson(data)
	if err != nil {
		log.Println(err)
	}
	switch v := slics["data"].(type) {
	default:
		for _, vx := range v.(map[string]interface{}) {
			switch vv := vx.(type) {
			case []interface{}:
				for _, vvvv := range vv {
					switch t := vvvv.(type) {
					case map[string]interface{}:
						for _, kv := range t {
							switch tt := kv.(type) {
							case map[string]interface{}:
								if tt["job"] == job {
									UpHost = append(UpHost, fmt.Sprintf("[%s]%s:%s", job, tt["instance"], tt["__name__"]))
								}

							}
						}

					}

				}

			}

		}

	}
	return UpHost

}
