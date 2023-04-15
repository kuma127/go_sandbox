package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	client := &http.Client{}
	// まず、Contextを作成
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}
	// ヘッダーの追加
	request.Header.Add("Content-Type", "image/jpeg")
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
