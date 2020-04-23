package main

import (
	f "fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(f.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	f.Println(err)
	if err != nil {
		f.Println("no value.")
	}

	apk := os.Getenv("API_KEY")

	f.Println("env value is " + apk)
}
