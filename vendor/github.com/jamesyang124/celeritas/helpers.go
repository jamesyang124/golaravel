package celeritas

import "os"

const DIR_MODE = 755

func (c *Celeritas) CreateDirIfNotExisted(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, DIR_MODE); err != nil {
			return err
		}
	}

	return nil
}
