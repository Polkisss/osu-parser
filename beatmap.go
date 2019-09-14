package pcircle

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// NewBeatmap returns a new empty Beatmap.
func NewBeatmap() *Beatmap {
	return new(Beatmap)
}

// Beatmap stores information about single beatmap.
type Beatmap struct {
	FileFormatVersion int    // Specifies version of beatmap file
	FilePath          string // The location of beatmap .osu file

	// General
	//
	// Various properties about the beatmap's gameplay.
	AudioFilename        string    // The location of the audio file relative to the current folder
	AudioLeadIn          int       // The amount of time added before the audio file begins playing
	PreviewTime          int       // Defines when the audio file should begin playing when selected in the song selection menu
	Countdown            int       // The speed of the countdown which occurs before the first hit object appears
	SampleSet            SampleSet // Specifies which set of hit sounds will be used throughout the beatmap
	StackLeniency        float64   // How often closely placed hit objects will be stacked together
	GameMode             int       // Defines the game mode of the beatmap (0=osu!, 1=Taiko, 2=Catch the Beat, 3=osu!mania)
	LetterboxInBreaks    bool      // Whether the letterbox appears during breaks
	StoryFireInFront     bool      // Whether or not display the storyboard in front of combo fire
	SkinPreference       string    // The preferred skin to use during gameplay
	EpilepsyWarning      bool      // Whether or not show a 'This beatmap contains scenes with rapidly flashing colours...' warning at the beginning of the beatmap
	CountdownOffset      int       // How many beats earlier the countdown starts
	WidescreenStoryboard bool      // Whether or not the storyboard should be widescreen
	SpecialStyle         bool      // Whether or not use the special N+1 style for osu!mania
	UseSkinSprites       bool      // Whether or not the storyboard can use user's skin resources

	// Editor
	//
	// Saved settings for mappers while editing beatmaps.
	Bookmarks       []int   // A list of times of editor bookmarks
	DistanceSpacing float64 // A multiplier for the "Distance Snap" feature
	BeatDivisor     int     // The beat division for placing objects
	GridSize        int     // The size of the grid for the "Grid Snap" feature
	TimelineZoom    float64 // The zoom in the editor timeline

	// Metadata
	//
	// Descriptive information about the song and beatmap.
	Title         string   // The title of the song limited to ASCII characters
	TitleUnicode  string   // The title of the song with unicode support
	Artist        string   // The name of the song's artist limited to ASCII characters
	ArtistUnicode string   // The name of the song's artist with unicode support
	Creator       string   // The username of the mapper
	Version       string   // The name of the beatmap's difficulty
	Source        string   // Describes the origin of the song
	Tags          []string // A collection of words describing the song
	BeatmapID     int      // The web ID of the single beatmap
	BeatmapSetID  int      // The web ID of the beatmap set (Mapset)

	// Difficulty
	//
	// Values defining the difficulty of the beatmap.
	HPDrainRate       float64 // How fast the health decreases
	CircleSize        float64 // The size of the hit objects in the osu!standard mode. In osu!mania mode, CircleSize is the number of columns
	OverallDifficulty float64 // The harshness of the hit window and the difficulty of spinners
	ApproachRate      float64 // Defines when hit objects start to fade in relatively to when they should be hit
	SliderMultiplier  float64 // Specifies the multiplier of the slider velocity
	SliderTickRate    float64 // The number of ticks per beat

	// Events
	//
	// A list of storyboard events.
	Background *Background // The location of the background image relative to the beatmap directory
	Breaks     []*Break    // Break times through the beatmap
	// todo: storyboards
	// todo: video

	// Timing Points
	//
	// A list of the beatmap's timing points and hitsounds.
	// Describes a number of properties regarding beats per minute and
	// hit sounds. Sorted by offset in the timing points section.
	TimingPoints []*TimingPoint

	// Colours
	//
	// RGB values of the combo colours used.
	ComboColours []*RGB // Defines the colours of combos

	SliderBody          *RGB
	SliderTrackOverride *RGB
	SliderBorder        *RGB
	// Extra colours for sliders

	// Hit Objects
	//
	// A list of the beatmap's hit objects.
	HitObjects []interface{}
}

// FromFile parses specified file and fills Beatmap with data.
func (b *Beatmap) FromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	b.FilePath = path

	scanner := bufio.NewScanner(f)
	var section string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) <= 2 || strings.HasPrefix(line, "//") || strings.HasPrefix(line, ";") {
			continue
		}

		if b.FileFormatVersion == 0 && strings.HasPrefix(line, "osu file format v") {
			b.FileFormatVersion, err = strconv.Atoi(line[17:])
			if err != nil {
				return err
			}
			continue
		}

		if strings.HasPrefix(line, "[") {
			section = strings.TrimRight(strings.TrimLeft(line, "["), "]")
			continue
		}

		// parsing line
		switch section {
		case "General":
			head, data := tokenize(line)
			switch head {
			case "AudioFilename":
				b.AudioFilename = data
			case "AudioLeadIn":
				b.AudioLeadIn, err = strconv.Atoi(data)
				if err != nil {
					return err
				}
			case "PreviewTime":
				b.PreviewTime, err = strconv.Atoi(data)
				if err != nil {
					return err
				}
			case "Countdown":
				b.Countdown, err = strconv.Atoi(data)
				if err != nil {
					return err
				}
			case "SampleSet":
				err = b.SampleSet.FromString(data)
				if err != nil {
					return err
				}
			case "StackLeniency":
				b.StackLeniency, err = strconv.ParseFloat(data, 64)
				if err != nil {
					return err
				}
			case "Mode":
				b.GameMode, err = strconv.Atoi(data)
				if err != nil {
					return err
				}
			case "LetterboxInBreaks":
				b.LetterboxInBreaks, err = string2int2bool(data)
				if err != nil {
					return err
				}
			case "StoryFireInFront":
				b.StoryFireInFront, err = string2int2bool(data)
				if err != nil {
					return err
				}
			case "SkinPreference":
				b.SkinPreference = data
			case "EpilepsyWarning":
				b.EpilepsyWarning, err = string2int2bool(data)
				if err != nil {
					return err
				}
			case "CountdownOffset":
				b.CountdownOffset, err = strconv.Atoi(data)
				if err != nil {
					return err
				}
			case "WidescreenStoryboard":
				b.WidescreenStoryboard, err = string2int2bool(data)
				if err != nil {
					return err
				}
			case "SpecialStyle":
				b.SpecialStyle, err = string2int2bool(data)
				if err != nil {
					return err
				}
			case "UseSkinSprites":
				b.UseSkinSprites, err = string2int2bool(data)
				if err != nil {
					return err
				}
			}

		case "Editor":
			head, data := tokenize(line)
			switch head {
			case "Bookmarks":
				bookmarks := strings.Split(data, ",")
				b.Bookmarks = make([]int, len(bookmarks))
				for i := range bookmarks {
					b.Bookmarks[i], err = strconv.Atoi(bookmarks[i])
					if err != nil {
						return err
					}
				}
			case "DistanceSpacing":
				b.DistanceSpacing, err = strconv.ParseFloat(data, 64)
				if err != nil {
					return err
				}
			case "BeatDivisor":
				b.BeatDivisor, err = strconv.Atoi(data)
				if err != nil {
					return err
				}
			case "GridSize":
				b.GridSize, err = strconv.Atoi(data)
				if err != nil {
					return err
				}
			case "TimelineZoom":
				b.TimelineZoom, err = strconv.ParseFloat(data, 64)
				if err != nil {
					return err
				}
			}

		case "Metadata":
			head, data := tokenize(line)
			switch head {
			case "Title":
				b.Title = data
			case "TitleUnicode":
				b.TitleUnicode = data
			case "Artist":
				b.Artist = data
			case "ArtistUnicode":
				b.ArtistUnicode = data
			case "Creator":
				b.Creator = data
			case "Version":
				b.Version = data
			case "Source":
				b.Source = data
			case "Tags":
				b.Tags = strings.Split(data, ",")
			case "BeatmapID":
				b.BeatmapID, err = strconv.Atoi(data)
				if err != nil {
					return err
				}
			case "BeatmapSetID":
				b.BeatmapSetID, err = strconv.Atoi(data)
				if err != nil {
					return err
				}
			}

		case "Difficulty":
			head, data := tokenize(line)
			switch head {
			case "HPDrainRate":
				b.HPDrainRate, err = strconv.ParseFloat(data, 64)
				if err != nil {
					return err
				}
			case "CircleSize":
				b.CircleSize, err = strconv.ParseFloat(data, 64)
				if err != nil {
					return err
				}
			case "OverallDifficulty":
				b.OverallDifficulty, err = strconv.ParseFloat(data, 64)
				if err != nil {
					return err
				}
			case "ApproachRate":
				b.ApproachRate, err = strconv.ParseFloat(data, 64)
				if err != nil {
					return err
				}
			case "SliderMultiplier":
				b.SliderMultiplier, err = strconv.ParseFloat(data, 64)
				if err != nil {
					return err
				}
			case "SliderTickRate":
				b.SliderTickRate, err = strconv.ParseFloat(data, 64)
				if err != nil {
					return err
				}
			}

		case "Events":
			// todo: storyboards
			// todo: video

			if strings.HasPrefix(line, "2,") {
				// Breaks
				br := &Break{}
				err = br.FromString(line)
				if err != nil {
					return err
				}
				b.Breaks = append(b.Breaks, br)
				continue

			}
			if strings.HasPrefix(line, "0,0,") {
				// Background
				bg := &Background{}
				err = bg.FromString(line)
				if err != nil {
					return err
				}
				b.Background = bg
				continue
			}

		case "TimingPoints":
			tp := new(TimingPoint)
			err = tp.FromString(line)
			if err != nil {
				return err
			}
			b.TimingPoints = append(b.TimingPoints, tp)

		case "Colours":
			head, data := tokenize(line)
			if strings.HasPrefix(head, "Combo") {
				colour := &RGB{}
				err = colour.FromString(data)
				if err != nil {
					return err
				}
				b.ComboColours = append(b.ComboColours, colour)
				continue
			}

			switch head {
			case "SliderBody":
				colour := &RGB{}
				err = colour.FromString(data)
				if err != nil {
					return err
				}
				b.SliderBody = colour
			case "SliderTrackOverride":
				colour := &RGB{}
				err = colour.FromString(data)
				if err != nil {
					return err
				}
				b.SliderTrackOverride = colour
			case "SliderBorder":
				colour := &RGB{}
				err = colour.FromString(data)
				if err != nil {
					return err
				}
				b.SliderBorder = colour
			}

		case "HitObjects":
			objectType, err := strconv.Atoi(strings.Split(line, ",")[3])
			if err != nil {
				return err
			}

			if (CIRCLE & objectType) == 1 {
				hitObject := &Circle{}
				err = hitObject.FromString(line)
				if err != nil {
					return err
				}
				b.HitObjects = append(b.HitObjects, hitObject)

			} else if (SLIDER & objectType) > 0 {
				hitObject := &Slider{}
				err = hitObject.FromString(line)
				if err != nil {
					return err
				}
				b.HitObjects = append(b.HitObjects, hitObject)
				
			} else if (SPINNER & objectType) > 0 {
				hitObject := &Spinner{}
				err = hitObject.FromString(line)
				if err != nil {
					return err
				}
				b.HitObjects = append(b.HitObjects, hitObject)
				
			} else if (MANIA_HOLD_NOTE & objectType) > 0 {
				hitObject := &ManiaHoldNote{}
				err = hitObject.FromString(line)
				if err != nil {
					return err
				}
				b.HitObjects = append(b.HitObjects, hitObject)
			}
		default:
			return errors.New("invalid section in beatmap file: '" + section + "'")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
