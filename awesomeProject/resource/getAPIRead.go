package resource

import (
	model "awesomeProject/mode"
	"errors"
	"fmt"
	jsdata "github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
)

func getApiRead(url string) ([]byte, error) {
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
func PastreadJson(data []byte) (map[string]interface{}, error) {
	js, err := jsdata.NewJson(data)

	if err != nil {
		return nil, errors.New("data is not meta")
	}

	return js.Map()

}

//var (
//	cst *time.Location
//)

// CSTLayout China Standard Time Layout
//const CSTLayout = "2006-01-02 15:04:05"

func GetreadData(flagStr string) []string {
	var list []string
	var IPlist []map[string]interface{}
	//dayTime := time.Now().Add(-time.Hour * time.Duration(7*24)).Unix()
	//hourTime := time.Now().Add(-time.Hour * time.Duration(7*24)).Unix()
	url := fmt.Sprintf("%s/api/v1/query_range?query=%s&start=%v&end=%v&step=15s", model.C.Url.ApiUrl, flagStr, model.StartTime, model.EndTime)
	data, _ := getApiRead(url)
	jsdata, err := PastreadJson(data)
	if err != nil {
		log.Println(err)
	}
	for _, v := range jsdata {
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
						switch ob := y.(type) {
						case map[string]interface{}:
							IPlist = append(IPlist, ob)

						}
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
					}
					{

					}

				}
			}

		}

	}
	return list

}

//CreateXlS(GetData("node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes * 100", "host-node"), "test", []string{"Status"})
//	fmt.Println(GetDData(`max(rate(node_disk_read_bytes_total{job=~"host-node"}[30s]))by(instance)/1024`, "host-node"))
