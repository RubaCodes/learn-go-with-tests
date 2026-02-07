package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

// il semplice reader una volta che ha letto tutto il byte[],
// nelle chiamate successive tornera una stringa vuota in quanto ha terminato il suo lavoro
// Con il ReadSeeker posso specificare da dove leggere il byte[] [offset, from]
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil

}

// file_system_store.go
func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins < f.league[j].Wins
	})
	return f.league
}
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	// When you range over a slice you are returned the current index of the loop
	//  (in our case i) and a copy of the element at that index.
	//  Changing the Wins value of a copy won't have any effect on the league slice
	//  that we iterate on. For that reason, we need to get the reference to the actual
	//  value by doing league[i] and then changing that value instead.
	//
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	f.database.Encode(f.league)

}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	if player := f.GetLeague().Find(name); player != nil {
		return player.Wins
	}
	return 0

}
