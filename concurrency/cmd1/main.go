package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"

	"log"
	"net/http"
	"time"
)

func myfun() {
	if r := recover(); r != nil {
		fmt.Println("raj defer " ,r)
	}
}
func main() {
	go func() {
		muxer := mux.NewRouter()
		muxer.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
			typeStruct := struct {
				Raj string `json:"raj"`
			}{}
			//time.Sleep(10 * time.Second)
			d := json.NewDecoder(request.Body).Decode(&typeStruct)
			if d != nil {
				fmt.Println("error deconding the reques")
			}
			fmt.Fprintf(writer, "hello %s", "raj")
		}).Methods(http.MethodGet)
		http.ListenAndServe(":8081", muxer)
	}()
	//ctx := context.Background()
	//ctx = context.WithValue(ctx, "raj", "kumar")
	//ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, 2*time.Second)
	//defer cancelFunc()
	//send(ctxWithTimeout)
	//ctxWithTimeoutSec, cancelFuncSec := context.WithTimeout(ctx, 3*time.Second)
	//defer cancelFuncSec()
	//sendSecond(ctxWithTimeoutSec)
	//fmt.Println("hello raj")

	typeStruct := struct {
		Raj string `json:"raj"`
	}{
		Raj: "hello raj ",
	}
	by, err := json.Marshal(typeStruct)
	defer myfun()
	if err != nil {
		panic("error marshalling the value ")
	}
	//ct := context.Background()
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8081/get", bytes.NewReader(by))
	if err != nil {
		panic("err in requeist formation")
	}
	//ct, cf := context.WithTimeout(ct, 5*time.Second)
	//defer cf()
	//request = request.WithContext(ct)
	fmt.Println(getResponse(request))
	//time.Sleep(20*time.Second)
}

func getResponse(request *http.Request) interface{} {
	ch := make(chan interface{})
	go returnRes(request,ch)
	for {
		select {
		case responseface := <-ch:
			return responseface
		case <-time.After(3 * time.Second):
			return "timedout"
		default:

		}
	}
}

func returnRes(request *http.Request,ch chan interface{}) {
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic("error calling the client" + err.Error())
	}
	defer res.Body.Close()
	var x interface{}
	//fmt.Println(res)
	if res.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		x = string(bodyBytes)
	} else {
		log.Println("erro in getting response bod ", res.Body, res.StatusCode)
		x = "error in resposne"
	}
	//fmt.Println("hello raj " ,x )
	ch <- x
}

func send(ctx context.Context) {
	fmt.Println(ctx.Value("raj"))
	select {
	case <-time.NewTimer(5 * time.Second).C:
		fmt.Println("hello raj")
	case <-ctx.Done():
		fmt.Println("context done")
	}
}

func sendSecond(ctx context.Context) {
	for {
		x := time.Now()
		time.Sleep(9 * time.Second)
		fmt.Println(time.Since(x))
		break
	}
}
