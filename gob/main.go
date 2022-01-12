package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "yzh", Age: 18}

	buffer := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(&u)
	if err != nil {
		panic(err)
	}

	fmt.Println(buffer.Bytes())

	var u2 User
	reader := bytes.NewReader(buffer.Bytes())
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(&u2)
	if err != nil {
		panic(err)
	}

	fmt.Println(u2)

}
