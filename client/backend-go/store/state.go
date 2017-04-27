package store

import (
	"time"
	"io/ioutil"
	"log"
	"path/filepath"
	"crypto/md5"
	"encoding/json"
	"encoding/hex"
)

type State struct {
	FileStates   []FileState
	LastModified time.Time
}

type FileState struct {
	Hash         string
	Path         string
	LastModified time.Time
}

/**
 * Compare the two states and return a new state which
 * only contains the new elements from the parameter state.
 */
func (this *State) GetDiffState(state *State) State {
	if !state.IsNewer(this) {
		return State{}
	}
	diffFileStates := []FileState{}
	fsMap := this.getFileStatesMap()
	for _, fileState := range state.FileStates {
		exFileState, exists := fsMap[fileState.Path]
		if !exists || exFileState.isModified(&fileState) {
			diffFileStates = append(diffFileStates, fileState)
		}
	}
	return State{
		FileStates: diffFileStates,
		LastModified: state.LastModified,
	}
}

func (this *State) IsNewer(state *State) bool {
	return this.LastModified.After(state.LastModified)
}

/**
 * Convenience method for faster access
 */
func (this *State) getFileStatesMap() map[string]FileState {
	fsMap := map[string]FileState{}
	for _, fileState := range this.FileStates {
		fsMap[fileState.Path] = fileState
	}
	return fsMap
}

func (this *FileState) isModified(fileState *FileState) bool {
	return fileState.Hash != this.Hash &&
			fileState.LastModified.After(this.LastModified)
}

func GetState(jsonString string) State {
	var state State
	err := json.Unmarshal([]byte(jsonString), &state)
	if err != nil {
		log.Fatal(err)
	}
	return state
}

func GetLocalState(projectPath string) State {
	fileStates, lastModified := getFileStates(projectPath, "")
	return State{
		FileStates: fileStates,
		LastModified: lastModified,
	}
}

func getFileStates(projectPath, dirPath string) ([]FileState, time.Time) {
	fileStates := []FileState{}
	var lastModified time.Time

	fullPath := filepath.Join(projectPath, dirPath)
	fileInfos, err := ioutil.ReadDir(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range fileInfos {
		path := filepath.Join(dirPath, fileInfo.Name())
		iterLastModified := fileInfo.ModTime().UTC()

		if fileInfo.IsDir() {
			subFileStates, subLastModified := getFileStates(projectPath, path)
			iterLastModified = subLastModified
			fileStates = append(fileStates, subFileStates...)
		} else {
			fileState := FileState{
				Hash: getFileHash(filepath.Join(projectPath, path)),
				Path: path,
				LastModified: iterLastModified,
			}
			fileStates = append(fileStates, fileState)
		}

		if iterLastModified.After(lastModified) {
			lastModified = iterLastModified
		}
	}
	return fileStates, lastModified
}

func getFileHash(path string) string {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return getHash(fileContent)
}

func getHash(fileContent []byte) string {
	hasher := md5.New()
	hasher.Write(fileContent)
	return hex.EncodeToString(hasher.Sum(nil))
}

