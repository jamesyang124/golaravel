package celeritas

import (
	"log"
	"os"
)

const DIR_MODE = 0755

func (c *Celeritas) CreateDirIfNotExisted(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, DIR_MODE); err != nil {
			return err
		}
	}

	return nil
}

func (c *Celeritas) CreateFileIfNotExisted(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Printf(".env file create failed %v", err)
			return err
		}

		// https://stackoverflow.com/questions/75640684/pass-pointer-to-go-defer-function-doesnt-work
		defer func(file *os.File) {
			log.Printf("defer file %v", file)
			_ = file.Close()
		}(file)
	}

	return nil
}
