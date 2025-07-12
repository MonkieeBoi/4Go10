package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

var rootURI fyne.URI

func initStorage(rURI fyne.URI) {
	rootURI = rURI
}

func getHighScoreURI() fyne.URI {
	uri, _ := storage.Child(rootURI, "highscore.txt")
	return uri
}

func readHighScore() int {
	hsURI := getHighScoreURI()
	if ok, err := storage.CanRead(hsURI); !ok || err != nil {
		return 0
	}
	reader, err := storage.Reader(hsURI)
	if err != nil {
		return 0
	}
	hsBytes := make([]byte, 50)
	n, err := reader.Read(hsBytes)
	if n == 0 || err != nil {
		return 0
	}
	hs, err := strconv.Atoi(string(hsBytes[:n]))
	if err != nil {
		return 0
	}
	return hs
}

func updateHighScore(score int) {
	hs := readHighScore()
	if score <= hs {
		return
	}
	hsURI := getHighScoreURI()
	if ok, err := storage.CanWrite(hsURI); !ok || err != nil {
		return
	}
	writer, err := storage.Writer(hsURI)
	if err != nil {
		return
	}
	writer.Write([]byte(strconv.Itoa(score)))
}
