package main

import (
	"fmt"
	"github.com/lazureykis/dotenv"
	"os"
)

func main() {
	fmt.Println("Before config loaded:")
	fmt.Println("password:", os.Getenv("PASSWORD"))
	fmt.Println("cache host:", os.Getenv("CACHE_HOST"))
	fmt.Println("cookie secret:", os.Getenv("COOKIE_SECRET"))

	fmt.Println()
	fmt.Println("Loading .env file from current directory...")
	fmt.Println()
	dotenv.Go()

	fmt.Println("password:", os.Getenv("PASSWORD"))
	fmt.Println("cache host:", os.Getenv("CACHE_HOST"))
	fmt.Println("cookie secret:", os.Getenv("COOKIE_SECRET"))
}
