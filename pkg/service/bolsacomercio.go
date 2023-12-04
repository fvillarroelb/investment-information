package repository

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	entity "fv.io/investment-information/pkg/domain/entities"
)

func GetUTMercados(url string) (string, error) {

	// Marshal the struct to a JSON-encoded byte slice
	jsonData, err := json.Marshal("")
	if err != nil {
		return "", err
	}

	// Create an io.Reader from the JSON-encoded byte slice
	payload := bytes.NewReader(jsonData)

	client := &http.Client{}
	req, err := http.NewRequest(entity.POST, url, payload)
	if err != nil {
		//log.Default()(err)
		return "", err
	}

	req.Header.Add("Cookie", "BIGipCookie=000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000; __uzma=79e6e54d-0428-459d-87dc-38886db00d91; __uzmb=1663001672; __uzme=5773; BIGipCookie=000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000; f5avraaaaaaaaaaaaaaaa_session_=DKLHEFMBKHADHCDKCHGCDLMOPGNNFKGMGOFOBIJJKCKDOHJMODBJKBCCAOPGEDOPGFBDDBMNJDPMOGGGKMFACGLBPGOHFDPHBPDCHNMBPGEONNPHJPLDBNFCLFIAIGEI; BIGipServerPool-Push_HTML5_corporativa=684715681.20480.0000; __ssds=2; __ssuzjsr2=a9be0cd8e; __uzmaj2=7d3cd1fb-fd7a-4538-9e8a-154eb558831c; __uzmbj2=1663001673; gb-wbchtbt-uid=1663001674795; _ga=GA1.2.1805551646.1663001675; _gid=GA1.2.745080428.1663001675; _csrf=SHnFl6DaiW6u0B6u2zIEAKdd; __gads=ID=4d6cf2779c3a94b1:T=1663001687:S=ALNI_MZzoILu4slPkv0O5jiQuyIfdMtfFQ; __gpi=UID=000009644fd6d262:T=1663001687:RT=1663001687:S=ALNI_MbR93GQ2YGRn3kInX8TUEIU5v_QOA; _gat=1; __uzmcj2=400402251886; __uzmdj2=1663002427; BIGipServerPool-www.bolsadesantiago.com-HTML5_corporativa=735047329.20480.0000; __uzmc=1655310040454; __uzmd=1663002432; BIGipCookie=000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000; f5avraaaaaaaaaaaaaaaa_session_=GOPKMMCPGDCGPEOLNBNAPOGEAODPJENBHFEEDINCCFEGPDOBDCKJAKENNHFLFCJNBFCDJMNPNBKNAHFIENJANGDNLPMIIEBIJOPMHMOJBCPAFONDNKHNFMKCEANKCHGD; BIGipServerPool-Push_HTML5_corporativa=684715681.20480.0000; BIGipServerPool-www.bolsadesantiago.com-HTML5_corporativa=684715681.20480.0000; __uzma=fe31de40-675b-48b6-b761-ff4a7a94c4e2; __uzmb=1663002361; __uzmc=7648910324972; __uzmd=1663037965; __uzme=9505; _csrf=q3bJQH8z97ajc55kn1uInjvW")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36 Edg/105.0.1343.27")
	req.Header.Add("x-csrf-token", "9G8bMVxi-wPn9NrW2JM8AvOoJOdmpVpP0_IQ")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "text/plain")

	res, error2 := client.Do(req)
	if error2 != nil {
		//fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	resp, _ := beautifyHTTPResponse(res)
	//fmt.Println(resp)

	return string(resp), err
}

func GetDividends(url string, request entity.RequestMarketAll) (string, error) {

	payload := strings.NewReader(`{"fec_pagoini":"2023-01-01","fec_pagofin":"2023-12-31","nemo":""}`)

	client := &http.Client{}
	req, err := http.NewRequest(entity.POST, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Cookie", "BIGipCookie=000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000; __uzma=79e6e54d-0428-459d-87dc-38886db00d91; __uzmb=1663001672; __uzme=5773; BIGipCookie=000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000; f5avraaaaaaaaaaaaaaaa_session_=DKLHEFMBKHADHCDKCHGCDLMOPGNNFKGMGOFOBIJJKCKDOHJMODBJKBCCAOPGEDOPGFBDDBMNJDPMOGGGKMFACGLBPGOHFDPHBPDCHNMBPGEONNPHJPLDBNFCLFIAIGEI; BIGipServerPool-Push_HTML5_corporativa=684715681.20480.0000; __ssds=2; __ssuzjsr2=a9be0cd8e; __uzmaj2=7d3cd1fb-fd7a-4538-9e8a-154eb558831c; __uzmbj2=1663001673; gb-wbchtbt-uid=1663001674795; _ga=GA1.2.1805551646.1663001675; _gid=GA1.2.745080428.1663001675; _csrf=SHnFl6DaiW6u0B6u2zIEAKdd; __gads=ID=4d6cf2779c3a94b1:T=1663001687:S=ALNI_MZzoILu4slPkv0O5jiQuyIfdMtfFQ; __gpi=UID=000009644fd6d262:T=1663001687:RT=1663001687:S=ALNI_MbR93GQ2YGRn3kInX8TUEIU5v_QOA; _gat=1; __uzmcj2=400402251886; __uzmdj2=1663002427; BIGipServerPool-www.bolsadesantiago.com-HTML5_corporativa=735047329.20480.0000; __uzmc=1655310040454; __uzmd=1663002432; f5avraaaaaaaaaaaaaaaa_session_=CDLAOIEAGMNCNJCCPCHOGLAHPIIPLNPKDDDBGDIFIKMNEPEJJMBGHLJOAPIOADHLACADGKNMOIFGBLPCLCIANLAMICNGHJMGFIJLMBBELNBPMGCFHLGOGANHPFGPJNOK; BIGipServerPool-www.bolsadesantiago.com-HTML5_corporativa=735047329.20480.0000; __uzmc=9081010340630; __uzmd=1701456453; _csrf=q3bJQH8z97ajc55kn1uInjvW")
	req.Header.Add("Origin", "https://www.bolsadesantiago.com")
	req.Header.Add("Referer", "https://www.bolsadesantiago.com/dividendos")
	req.Header.Add("sec-sh-ua", "\"Microsoft Edge\";v=\"105\", \" Not;A Brand\";v=\"99\", \"Chromium\";v=\"105\"")
	req.Header.Add("sec-sh-platform", "Windows")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36 Edg/105.0.1343.27")
	req.Header.Add("x-csrf-token", "9G8bMVxi-wPn9NrW2JM8AvOoJOdmpVpP0_IQ")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "es,es-ES;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("sec-ch-ua-mobile", "?0")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	jsonStringResponse := printHTTPResponseBody(res)
	return jsonStringResponse, nil
}

func beautifyHTTPResponse(response *http.Response) (string, error) {
	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Close the response body to prevent resource leaks
	defer response.Body.Close()

	// Unmarshal the JSON data
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	// Marshal the data back to a JSON string with indentation
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(prettyJSON), nil
}

func imprimirBytes(payload *bytes.Reader) string {
	// Read the content of the *bytes.Reader into a string
	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, payload)
	if err != nil {

		return ("Error reading from *bytes.Reader:")
	}

	// Convert the buffer to a string
	contentString := buffer.String()

	// Print the string
	//fmt.Println("Content as string:", contentString)
	return contentString
}

func printHTTPResponseBody(response *http.Response) string {
	var bodyReader io.Reader = response.Body

	// Check if the response is compressed
	if response.Header.Get("Content-Encoding") == "gzip" {
		// If compressed, use a gzip reader to decompress
		gzipReader, err := gzip.NewReader(response.Body)
		if err != nil {
			fmt.Println("Error creating gzip reader:", err)
			return "error decoding"
		}
		defer gzipReader.Close()

		bodyReader = gzipReader
	}

	// Read the response body
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "Read response error"
	}

	// Print the content of the response body
	//fmt.Println(string(body))
	return string(body)
}
