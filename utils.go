package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func checkIfCanDownload(filePath string) (string, error) {
	if _, err := os.Stat(filePath); err != nil {
		return "Error Touching the file, Permissions or doesnt exist", err
	}
	return filePath, nil
}

func wget(filepath string, url string) (string, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "Error while trying to reach url", err
	}
	defer resp.Body.Close()

	if filepath != "stdout" {
		// Create the file
		out, err := os.Create(filepath)
		if err != nil {
			return "Error creating output file", err
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		return "Outputed to file", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error converting byte array of body to stinrg", err
	}

	return string(body), nil

}
