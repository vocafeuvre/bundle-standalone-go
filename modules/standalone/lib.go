package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func outputFileSystemContents(fs http.FileSystem, inputDir string, outputDir string) error {
	// Walk the file system and output the contents of each file to disk
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		// Only process regular files
		if !info.Mode().IsRegular() {
			return nil
		}

		// Open the file from the file system
		file, err := fs.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Read the file contents
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		// Create the output directory if it doesn't exist
		err = os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return err
		}

		// Write the file contents to disk
		outputPath := filepath.Join(outputDir, path[1:])
		err = ioutil.WriteFile(outputPath, contents, os.ModePerm)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
