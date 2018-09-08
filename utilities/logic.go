package utilities

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func PostAPIUriCallEndPoint(urlStr string, requestBody interface{}) interface{} {
	fmt.Println("Url: ", urlStr)
	spew.Dump(requestBody)
	requestJSON, err := json.Marshal(requestBody)
	var outputObj interface{}
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Te", "application/json")
	// req.Header.Add("Authorization", cast.ToString(viper.Get("usertoken")))
	client := &http.Client{}
	//send the request
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &outputObj)
	fmt.Println("Response: ", outputObj)
	return outputObj
}

func GetAPIUriCallEndPoint(urlStr string) interface{} /*map[interface{}]interface{}*/ {

	var outputObj interface{} //make(map[interface{}]interface{})
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Te", "application/json")
	// req.Header.Add("Authorization", cast.ToString(viper.Get("usertoken")))
	client := &http.Client{}
	//send the request
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &outputObj)
	fmt.Println(resp.Body)

	return outputObj
}

//GenerateCSV
func GenerateCSV(csvContent [][]string) []byte {
	b := &bytes.Buffer{}   // creates IO Writer
	wr := csv.NewWriter(b) // creates a csv writer that uses the io buffer.
	for i := 0; i < len(csvContent); i++ {
		wr.Write(csvContent[i]) // converts array of string to comma seperated values for 1 row.
	}
	wr.Flush() // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))

	b.Bytes()

	return b.Bytes()
}

func ImportCsvFunction(req *http.Request) (row [][]string, err error) {
	s := ""
	if req.Method == http.MethodPost {
		f, _, err := req.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		defer f.Close()
		// log.Printf(h)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		s = string(bs)
	}

	NumberOfRow := strings.Split(s, "\n")
	for _, value := range NumberOfRow {
		rowRecord := strings.Split(value, ",")
		if len(rowRecord) > 0 {
			row = append(row, rowRecord)
		}

	}

	return row, nil
}

func MaskLeft(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs)-4; i++ {
		rs[i] = 'X'
	}
	return string(rs)
}
