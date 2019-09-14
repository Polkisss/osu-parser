package pcircle

import (
	"archive/zip"
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func NewMapset() *Mapset {
	return new(Mapset)
}

// Mapset stores information about beatmaps, its location and other.
type Mapset struct {
	DirectoryPath string // The location of mapset directory, where located beatmaps.

	Beatmaps     []*Beatmap // Unordered list of beatmaps
	BeatmapSetID int        // The web ID of the beatmap set
}

// FromDirectory scans provided (from structure) directory and loads .osu files into Mapset.
func (m *Mapset) FromDirectory(path string) (err error) {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New("invalid directory: " + path)
	}

	m.DirectoryPath = path

	bfiles, err := filepath.Glob(filepath.Join(path, "*.osu"))
	if err != nil {
		return err
	}

	m.Beatmaps = make([]*Beatmap, len(bfiles))
	for i, bfile := range bfiles {
		beatmap := NewBeatmap()
		err := beatmap.FromFile(bfile)
		if err != nil {
			return err
		}
		m.Beatmaps[i] = beatmap
	}

	if len(m.Beatmaps) > 0 {
		m.BeatmapSetID = m.Beatmaps[0].BeatmapSetID
	}

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
