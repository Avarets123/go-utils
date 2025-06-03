package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {

	h := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("/tmp/"))

	h.Handle("/static", http.StripPrefix("/static", fileServer))

	h.HandleFunc("/file", FileHandler)

	http.ListenAndServe(":2001", h)

}

func FileHandler(w http.ResponseWriter, r *http.Request) {

	fileName := "/home/osman/Рабочий стол/dumps/baks_db.sql"

	// f, err := os.Open(fileName)
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// _, err = io.Copy(w, f)
	http.ServeFile(w, r, fileName)
	// if err != nil {
	// 	panic(err)
	// }

}

type MyData struct {
	r *http.Response
	e error
}

func ClientReqCancel() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	t := &http.Transport{}
	c := &http.Client{Transport: t}
	dCH := make(chan MyData, 1)
	url := ""

	req, _ := http.NewRequest("GET", url, nil)

	go func() {
		res, err := c.Do(req)
		if err != nil {
			dCH <- MyData{nil, err}
			return
		} else {
			dCH <- MyData{res, nil}
		}
	}()

	select {
	case data := <-dCH:
		if data.e != nil {
			panic(data.e)
		}
		fmt.Println(data.r)
		defer data.r.Body.Close()

	case <-ctx.Done():
		t.CancelRequest(req)
		<-dCH
		panic(ctx.Err())
	}

}
