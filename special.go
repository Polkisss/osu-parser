package osu_parser

import (
	"strconv"
	"strings"
)

// Break defines a single break period.
//
// Example of an break period:
//  2,4627,5743
type Break struct {

	// Both are an number of milliseconds from the beginning of the song defining the start and end point of the break period
	StartTime int
	EndTime   int
}

func (b Break) String() string {
	return "2," + strings.Join(
		[]string{
			strconv.Itoa(b.StartTime),
			strconv.Itoa(b.EndTime),
		}, ",")
}

// Fills Break fields with data parsed from string.
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

type RGB struct {
	R, G, B int
}

func (c RGB) String() string {
	return strings.Join([]string{
		strconv.Itoa(c.R),
		strconv.Itoa(c.G),
		strconv.Itoa(c.B),
	}, ",")
}

// Fills RGB triplet with data parsed from string.
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
type Extras struct {
	SampleSet     SampleSet // Changes the sample set of the normal hit sound. When sampleSet is 0, its value should be inherited from the timing point.
	AdditionalSet SampleSet // Changes the sample set for the other hit sounds
	CustomIndex   int       // Custom sample set index
	SampleVolume  int       // Volume of the sample, and ranges from 0 to 100 (percent)
	Filename      string    // Names an audio file in the folder to play instead of sounds from sample sets
}

func (e Extras) String() string {
	return strings.Join([]string{
		strconv.Itoa(int(e.SampleSet)),
		strconv.Itoa(int(e.AdditionalSet)),
		strconv.Itoa(e.CustomIndex),
		strconv.Itoa(e.SampleVolume),
		e.Filename,
	}, ":")
}

// Fills Extras fields with data parsed from string.
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
