package library

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}
	if mm.len() != 0 {
		t.Error("NewMusicManager failed.")
	}
	m0 := &MusicEntity{
		"1",
		"My Heart Will Go On",
		"Celion",
		"http://qbox.me/24501234",
		"MP3"
	}
	mm.Add(m0)
	if mm.Len(mm) != 1{
		t.Error("MusicManager.Add() failed. ")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed .")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name ||
	m.Genre != m0.Genre || m.Soure != m0.Soure || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed . Found item mismatch.")
	}
	m, err := mm.Get(0)
	if m == nil{
		t.Error("MusicManager.Get() failed. ")
	}
	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed. ")
	}
}