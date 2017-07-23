package main

func mac() {
    var defaultName = "sam" //allowed
    type myString string
    var customName myString = "sam" //alowed
    customName = defaultName //not allowed
}
