package main

func main() {
	m := make(map[string]int)
	m["Litecoin"] = 50
	m["Stratus"] = 20
	delete(m, "Litecoin")
}
