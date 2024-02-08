package helper

import "github.com/joho/godotenv"

func initHelper() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}