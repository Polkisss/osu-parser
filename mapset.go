package pcircle

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

func NewMapset(directory string) *Mapset {
	return &Mapset{
		DirectoryPath: directory,
	}
}

// Mapset stores information about beatmaps, its location and other.
type Mapset struct {
	DirectoryPath string // The location of mapset directory, where located beatmaps.

	Beatmaps     []*Beatmap // Unordered list of beatmaps
	BeatmapSetID int        // The web ID of the beatmap set
}

// Scans provided directory and loads .osu files into Mapset.
// Not implemented yet.
func (m *Mapset) Load() (err error) {
	beatmaps, err := filepath.Glob(filepath.Join(m.DirectoryPath, "*.osu"))
	if err != nil {
		return err
	}
	_ = beatmaps

	return nil
}

// Compresses Mapset into .osz and returns its buffer.
func (m Mapset) ToOSZ() (buf *bytes.Buffer, err error) {
	w := zip.NewWriter(buf)

	err = filepath.Walk(m.DirectoryPath, func(path string, info os.FileInfo, err error) error {
		fp, err := filepath.Rel(m.DirectoryPath, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			_, err := w.Create(fp)
			if err != nil {
				return err
			}
			return nil
		}

		realFileData, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		f, err := w.Create(fp)
		if err != nil {
			return err
		}

		_, err = f.Write(realFileData)
		return err
	})

	if err != nil {
		return new(bytes.Buffer), err
	}

	if err := w.Close(); err != nil {
		return new(bytes.Buffer), err
	}

	return buf, nil
}
