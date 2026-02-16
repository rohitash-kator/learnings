package file_manager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string `json:"input_file_path"`
	OutputFilePath string `json:"output_file_path"`
}

func (fileManager FileManager) Read() ([]string, error) {
	file, err := os.Open(fileManager.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		// file.Close()
		return nil, errors.New("Failed to read line in file.")
	}

	// file.Close()
	return lines, nil
}

// Alternatively 'any' can be used instead of 'interface{}'
func (fileManager FileManager) Write(fileData interface{}) error {
	file, err := os.Create(fileManager.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}
	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(fileData)

	if err != nil {
		// file.Close()
		return errors.New("Failed to convert data to JSON")
	}

	// file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
