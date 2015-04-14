package main

import (
	"flag"
	"fmt"
	"os"
	"./timesync"
)

func main() {
	user := flag.String("user", "", "the username")
	flag.Parse();
	if *user == "" {
		*user = os.Getenv("USER")
	}
	fmt.Printf("hello %s", *user)

	editor := os.Getenv("EDITOR");
	if editor == "" {
		editor = "vi"
	}

	fmt.Printf(" your favorite editor is %s\n", editor)
	config := make(map[string]string)
	config["domain"] = "localhost"
	config["port"] = "3000"
	config["user"] = *user
	err := timesync.CheckIn(config, "", "docs", "ww", "", 0)
	if err != nil {
		fmt.Println(err)
	}


}
