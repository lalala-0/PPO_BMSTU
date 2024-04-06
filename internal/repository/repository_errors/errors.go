package repository_errors

import "errors"

var (
	InsertError  = errors.New("DB ERROR: Insert operation was not successful")
	DeleteError  = errors.New("DB ERROR: Delete operation was not successful")
	SelectError  = errors.New("DB ERROR: Select operation was not successful")
	UpdateError  = errors.New("DB ERROR: Update operation was not successful")
	DoesNotExist = errors.New("GET operation has failed. Such row does not exist")

	ConnectionError    = errors.New("DB ERROR: Connection error")
	ConfigParsingError = errors.New("DB ERROR: Config parsing error")
)
