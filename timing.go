package osu_parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Timing points describe a number of properties regarding beats per minute and hit sounds.
//
// Example of a TimingPoint:
//  66,315.789473684211,4,2,0,45,1,0
//
// Example of an inherited TimingPoint:
//  10171,-100,4,2,0,60,0,1
type TimingPoint struct {
	Offset int // Define when the timing point starts

	// Defines the duration of one beat. It affect the scrolling speed in osu!taiko or
	// osu!mania, and the slider speed in osu!standard.
	MillisecondsPerBeat float64

	Meter       int       // Defines the number of beats in a measure
	SampleSet   SampleSet // Defines the default sample set for hit objects
	SampleIndex int       // The default custom index
	Volume      int       // The default hitsound volume, ranges from 0 to 100 (percent)

	// Tells if the timing point can be inherited from.
	// A positive milliseconds per beat implies inherited is true (1), and a negative one implies it is false (0).
	// Note that false (0) means green line, true (1) means red line.
	Inherited bool

	Kiai bool // Defines whether or not Kiai Time effects are active
}

func (tp TimingPoint) String() string {
	var mpb string

	if tp.Inherited {
		mpb = fmt.Sprintf("%g", tp.MillisecondsPerBeat)
	} else {
		mpb = strconv.Itoa(int(tp.MillisecondsPerBeat))
	}

	return strings.Join([]string{
		strconv.Itoa(tp.Offset),
		mpb,
		strconv.Itoa(tp.Meter),
		strconv.Itoa(int(tp.SampleSet)),
		strconv.Itoa(tp.SampleIndex),
		strconv.Itoa(tp.Volume),
		bool2int2string(tp.MillisecondsPerBeat > 0),
		bool2int2string(tp.Kiai),
	}, ",")
}

// Fills TimingPoint fields with data parsed from string.
func (tp *TimingPoint) FromString(str string) (err error) {
	attrs := strings.Split(str, ",")

	tp.Offset, err = strconv.Atoi(attrs[0])
	if err != nil {
		return err
	}

	tp.MillisecondsPerBeat, err = strconv.ParseFloat(attrs[1], 64)
	if err != nil {
		return err
	}

	tp.Meter, err = strconv.Atoi(attrs[2])
	if err != nil {
		return err
	}

	ss, err := strconv.Atoi(attrs[3])
	if err != nil {
		return err
	}
	tp.SampleSet = SampleSet(ss)

	tp.SampleIndex, err = strconv.Atoi(attrs[4])
	if err != nil {
		return err
	}

	tp.Volume, err = strconv.Atoi(attrs[5])
	if err != nil {
		return err
	}

	tp.Inherited = tp.MillisecondsPerBeat > 0

	tp.Kiai, err = strconv.ParseBool(attrs[7])
	return err
}
