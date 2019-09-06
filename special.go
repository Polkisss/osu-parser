package pcircle

import (
	"strconv"
	"strings"
)

// Background specifies parameters of beatmap background.
// Example of an Background:
//  0,0,"bg.jpg",0,0
type Background struct {
	FileName         string
	XOffset, YOffset int
}

// String returns string of Background as it would be in .osu file
func (b Background) String() string {
	return "0,0," + strings.Join(
		[]string{
			`"` + b.FileName + `"`,
			strconv.Itoa(b.XOffset),
			strconv.Itoa(b.YOffset),
		}, ",")
}

// FromString fills Background fields with data parsed from string.
func (b *Background) FromString(str string) (err error) {
	attrs := strings.Split(strings.TrimLeft(str, "0,0,"), ",")

	b.FileName = strings.TrimRight(strings.TrimLeft(attrs[0], `"`), `"`)

	b.XOffset, err = strconv.Atoi(attrs[1])
	if err != nil {
		return err
	}

	b.YOffset, err = strconv.Atoi(attrs[2])
	return err
}

// Break defines a single break period.
// Example of an break period:
//  2,4627,5743
type Break struct {

	// Both are an number of milliseconds from the beginning of the song defining the start and end point of the break period
	StartTime int
	EndTime   int
}

// String returns string of Break as it would be in .osu file
func (b Break) String() string {
	return "2," + strings.Join(
		[]string{
			strconv.Itoa(b.StartTime),
			strconv.Itoa(b.EndTime),
		}, ",")
}

// FromString fills Break fields with data parsed from string.
func (b *Break) FromString(str string) (err error) {
	str = strings.TrimPrefix(str, "2,")

	sep := strings.Index(str, ",")

	b.StartTime, err = strconv.Atoi(str[:sep])
	if err != nil {
		return err
	}

	b.EndTime, err = strconv.Atoi(str[sep+1:])
	return err
}

// RGB provides color manipulation.
// Example:
//  128,128,0
type RGB struct {
	R, G, B int
}

// String returns string of RGB as it would be in .osu file
func (c RGB) String() string {
	return strings.Join([]string{
		strconv.Itoa(c.R),
		strconv.Itoa(c.G),
		strconv.Itoa(c.B),
	}, ",")
}

// FromString fills RGB triplet with data parsed from string.
func (c *RGB) FromString(str string) (err error) {
	attrs := strings.Split(str, ",")

	c.R, err = strconv.Atoi(attrs[0])
	if err != nil {
		return err
	}

	c.G, err = strconv.Atoi(attrs[1])
	if err != nil {
		return err
	}

	c.B, err = strconv.Atoi(attrs[2])
	return err
}

// Extras define additional parameters related to the hit sound samples.
// The most common example:
//  0:0:0:0:
type Extras struct {
	SampleSet     SampleSet // The sample set of the normal hit sound. When sampleSet is 0, its value should be inherited from the timing point.
	AdditionalSet SampleSet // The sample set for the other hit sounds
	CustomIndex   int       // Custom sample set index
	SampleVolume  int       // Volume of the sample, and ranges from 0 to 100 (percent)
	Filename      string    // Names an audio file in the folder to play instead of sounds from sample sets
}

// String returns string of Extras as it would be in .osu file
func (e Extras) String() string {
	return strings.Join([]string{
		strconv.Itoa(int(e.SampleSet)),
		strconv.Itoa(int(e.AdditionalSet)),
		strconv.Itoa(e.CustomIndex),
		strconv.Itoa(e.SampleVolume),
		e.Filename,
	}, ":")
}

// FromString fills Extras fields with data parsed from string.
func (e *Extras) FromString(str string) (err error) {
	attrs := strings.Split(str, ":")

	ss, err := strconv.Atoi(attrs[0])
	if err != nil {
		return err
	}
	e.SampleSet = SampleSet(ss)

	ss, err = strconv.Atoi(attrs[1])
	if err != nil {
		return err
	}
	e.AdditionalSet = SampleSet(ss)

	e.CustomIndex, err = strconv.Atoi(attrs[2])
	if err != nil {
		return err
	}

	e.SampleVolume, err = strconv.Atoi(attrs[3])
	e.Filename = attrs[4]
	return err
}
