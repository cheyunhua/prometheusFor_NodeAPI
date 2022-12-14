package resource

import (
	"awesomeProject/logger"
	model "awesomeProject/mode"
	"errors"
	"fmt"
	jsdata "github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
)

func GetCUrl(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		logger.DefaultLogger.Errorf("get请求失败 error: %+v", err)
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		logger.DefaultLogger.Errorf("读取Body失败 error: %+v", err)

	}

	return body, err1
}
func PastCJson(data []byte) (map[string]interface{}, error) {
	js, err := jsdata.NewJson(data)

	if err != nil {
		return nil, errors.New("data is not meta")
	}
	return js.Map()

}

func GetCCData(flagStr string) []string {
	var list []string
	var IPlist []map[string]interface{}
	//dayTime := time.Now().Add(-time.Hour * time.Duration(7*24)).Unix()
	//hourTime := time.Now().Add(-time.Hour * time.Duration(7*24)).Unix()
	url := fmt.Sprintf("%s/api/v1/query_range?query=%s&start=%v&end=%v&step=15s", model.C.Url.ApiUrl, flagStr, model.StartTime, model.EndTime)
	//url := fmt.Sprintf("%s/api/v1/query_range?query=%s&start=%d&end=%d&step=15s", model.C.Url.ApiUrl, flagStr, model.StartTime, model.EndTime)
	data, _ := GetCUrl(url)
	JSdata, err := PastCJson(data)
	if err != nil {
		log.Println(err)
	}
	for _, v := range JSdata {
		switch vv := v.(type) {
		case map[string]interface{}:
			i := vv["result"]
			switch xxx := i.(type) {
			case []interface{}:
				for _, v := range xxx {
					switch yyy := v.(type) {
					case map[string]interface{}:
						x := yyy["values"]
						y := yyy["metric"]
						switch oa := x.(type) {
						case []interface{}:
							for _, v := range oa {
								switch oc := v.(type) {
								case []interface{}:
									for _, v := range oc {
										if vc, ok := v.(string); ok {
											list = append(list, vc)
										}

									}
								}
							}
						}
						switch ob := y.(type) {
						case map[string]interface{}:
							IPlist = append(IPlist, ob)
							//for _, v := range ob {
							//	if vb, ok := v.(string); ok {

							//	}

							//}

						}

					}
				}

			}

		}

	}

	return list
}
