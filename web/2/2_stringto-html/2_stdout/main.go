package main

import (
	"fmt"
)

func main() {
	name := "Romano"

	tpl := `
	<!DOCTYPE html>
	<head><title>Hello world</title></head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
}
