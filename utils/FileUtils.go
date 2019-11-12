package utils

import (
	"os"
)

func LocalPathExists(localPath string) (bool, error) {

	_, err := os.Stat(localPath)

	if err == nil {
		return true, nil
	} else {
		LogWarn(err)
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func CreateLocalDir(dir string) (bool, error) {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		LogWarn(err)
		return false, err
	} else {
		return true, nil
	}
}
