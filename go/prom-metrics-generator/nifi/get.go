package nifi

import (
        "prom-metrics-generator/logger"
        "io/ioutil"
        "encoding/json"
		"net/http"
)

func GetCall(
	url string,        
) []byte{
	response, err := http.Get(url)

	var responseData []byte

	if err != nil {
			clog.Error.Println(err.Error())
	} else {
			responseData, err = ioutil.ReadAll(response.Body)
			if err != nil {
					clog.Error.Println(err)
			}
	}
   
	return responseData
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
