package bitwardenwrapper

import (
	"encoding/json"
	"fmt"
)

// GetFolder finds a bitwarden folder by name and returns it or nil and error if it cannot
func GetFolder(folderName string) (*BwFolder, error) {
	folderjson, err := runBw("list", "folders")
	if err != nil {
		return nil, err
	}
	folders := make([]BwFolder, 0)
	json.Unmarshal([]byte(folderjson), &folders)
	for _, folder := range folders {
		if folder.Name == folderName {
			return &folder, nil
		}
	}
	return nil, fmt.Errorf("Cannot find folder called %s", folderName)
}
