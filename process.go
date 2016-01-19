package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/seibert-media/foswiki-uwc-prepare/foswiki"
)

// processDataContent reads the Foswiki data files and saves the replaced content.
func processDataContent(path string) error {
	fmt.Printf("processing %v\n", path)

	dataFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	dataFile = replaceAllVerbatimTags(&dataFile)
	dataFile = replaceAllBase64Images(&dataFile, path)

	err = ioutil.WriteFile(path, dataFile, 0644)
	if err != nil {
		return err
	}
	return nil
}

// processDirContent calls processDataContent with the path to *.txt Foswiki data files.
func processDirContent(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if strings.ToLower(filepath.Ext(path)) == ".txt" {
		if err := processDataContent(path); err != nil {
			return err
		}
	}
	return nil
}

// startProcessing checks if the needed subdirectories exists and walk through the Foswiki data directory.
func startProcessing(homeDir string) error {
	foswikiDir := foswiki.Dir{homeDir}
	if _, err := foswikiDir.SubdirsExists(); err != nil {
		return err
	}

	if err := filepath.Walk(foswikiDir.DataPath(), processDirContent); err != nil {
		return err
	}
	return nil
}
