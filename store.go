package main

import (
	"fmt"
	"math/rand"
	"os"
)

func SaveData1(dir string, path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)

	if err != nil {
		return err
	}

	defer fp.Close()

	_, err = fp.Write(data)

	if err != nil {
		return err
	}

	return fp.Sync()
}

func SaveData2(path string, data []byte) error {

	temp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	fp, err := os.OpenFile(temp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)

	if err != nil {
		return err
	}

	defer func() {
		fp.Close()
		if err != nil {
			os.Remove(temp)
		}
	}()

	_, err = fp.Write(data)

	if err != nil {
		return err
	}

	if err = fp.Sync(); err != nil {
		return err
	}

	err = os.Rename(temp, path)
	return err

}

func main() {

	var path string
	var data []byte

	err := SaveData1(path, data)

	if err != nil {
		fmt.Println("Error occured!")
	}

}
