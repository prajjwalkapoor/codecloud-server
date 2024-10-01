package model

type FileInfo struct {
	Path  string     `json:"path"`
	Name  string     `json:"name"`
	Files []FileInfo `json:"files"`
}
