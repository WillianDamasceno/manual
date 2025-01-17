package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"manual/helpers"
	"manual/src"
	"net/http"
	"os"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
}

func Update() {
	helpers.CheckSudo()

	apiUrl := fmt.Sprintf(src.ReleaseUrl)
	client := http.Client{}

	print("Searching latest release...\n")

	// Fetch GitHub releases
	response, err := client.Get(apiUrl)
	helpers.Error(err, "Error while fetching the latest release")

	defer response.Body.Close()

	statusCode := response.StatusCode
	helpers.PanicIf(statusCode != http.StatusOK, "Error while fetching the latest release:", statusCode)

	// Parse the response body
	body, err := io.ReadAll(response.Body)
	helpers.Error(err, "Error while reading the response body")

	// Get response body as JSON
	var releases []GitHubRelease

	err = json.Unmarshal([]byte(body), &releases)
	helpers.Error(err, "Error while parsing the response body")
	
	print("Downloading latest release...\n")

	// Fetch new binary
	binaryUrl := src.GetBinaryUrl(releases[0].TagName)
	response, err = client.Get(binaryUrl)
	helpers.Error(err, "Error while downloading the binary")
	defer response.Body.Close()

	statusCode = response.StatusCode
	helpers.PanicIf(statusCode != 200, "Error while downloading the binary:", statusCode)

	// Parse the response body
	body, err = io.ReadAll(response.Body)
	helpers.Error(err, "Error while reading the response body")
	defer response.Body.Close()

	print("Installing latest release...\n")

	// Save binary into /tmp
	file, err := saveBinaryToTempDir(body)
	helpers.Error(err, "Error while saving the binary to a temporary directory")

	// Move binary
	err = os.Rename(file.Name(), src.BinaryPath)
	helpers.Error(err, "Error while moving the binary to /usr/local/bin")

	// Set binary as executable
	os.Chmod(src.BinaryPath, os.FileMode(0755))

	print("Latest release installed successfully!\n\n")
}

func saveBinaryToTempDir(binaryData []byte) (*os.File, error) {
	file, err := os.CreateTemp("", "manual-")
	if err != nil {
		return nil, fmt.Errorf("error creating file in temp directory: %w", err)
	}
	defer file.Close()

	_, err = file.Write(binaryData)
	if err != nil {
		return nil, fmt.Errorf("error writing binary to file: %w", err)
	}

	return file, nil
}
