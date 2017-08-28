package library

import (
	"errors"
)

type MusicManager struct {
	musics []MusicEntity
}

func NewMusicManager() *MusicManager {

	return &MusicManager{make([]MusicEntity, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Find(name string) *MusicEntity {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Get(index int) (music *MusicEntity, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Add(music *MusicEntity) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntity {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removedMusic := &m.musics[index]
	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1])
	}else if index == 0 {
		m.musics = make([]MusicEntity, 0)
	}else {
		m.musics = m.musics[:index-1]
	}
	return removedMusic
}

