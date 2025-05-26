package main

import (
	"fmt"
	"os"
)

func main() {
	tmp := os.TempDir()
	file, err := os.Create(tmp + "/myfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	msg := "Save the world with Go!"

	_, err = file.WriteString(msg)
	if err != nil {
		panic(err)
	}

	positions := []int{4, 10, 20}
	for _, i := range positions {
		_, err := file.Seek(int64(i), 0)
		if err != nil {
			panic(err)
		}
		file.Write([]byte("X"))
	}

	file.Seek(0, 0)

	result := make([]byte, len(msg))
	if _, err := file.Read(result); err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result)
}
