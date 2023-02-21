package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bid, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("context.txt")
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", bid))
	if err != nil {
		panic(err)
	}
}
