package pcircle

import (
	"fmt"
	"strconv"
	"strings"
)

// Bitmap for BaseHitObject.Type.
const (
	CIRCLE = 1 << iota
	SLIDER
	NEW_COMBO
	SPINNER
	COMBO_SKIP_1
	COMBO_SKIP_2
	COMBO_SKIP_3
	MANIA_HOLD_NOTE
)

// BaseHitObject provides common information that is used in all hit objects.
type BaseHitObject struct {
	X, Y     int
	Time     int
	Type     int
	HitSound HitSound
	Extras   *Extras
}

// Circle is a single hit in all osu! game modes.
// Example:
//  164,260,2434,1,0,0:0:0:0:
type Circle struct {
	BaseHitObject
}

// String returns string of Circle as it would be in .osu file.
func (c Circle) String() string {
	return strings.Join([]string{
		strconv.Itoa(c.X),
		strconv.Itoa(c.Y),
		strconv.Itoa(c.Time),
		strconv.Itoa(c.Type),
		strconv.Itoa(int(c.HitSound)),
		c.Extras.String(),
	}, ",")
}

// FromString fills Circle fields with data parsed from string.
func (c *Circle) FromString(str string) (err error) {
	attrs := strings.Split(str, ",")

	c.X, err = strconv.Atoi(attrs[0])
	if err != nil {
		return err
	}

	c.Y, err = strconv.Atoi(attrs[1])
	if err != nil {
		return err
	}

	c.Time, err = strconv.Atoi(attrs[2])
	if err != nil {
		return err
	}

	c.Type, err = strconv.Atoi(attrs[3])
	if err != nil {
		return err
	}

	hs, err := strconv.Atoi(attrs[4])
	if err != nil {
		return err
	}
	c.HitSound = HitSound(hs)

	c.Extras = new(Extras)
	return c.Extras.FromString(attrs[5])
}

// Slider also creates droplets in Catch the Beat, yellow drumrolls in Taiko,
// and does not appear in osu!mania.
// Example:
//  424,96,66,2,0,B|380:120|332:96|332:96|304:124,1,130,2|0,0:0|0:0,0:0:0:0:
type Slider struct {
	BaseHitObject

	// Specifies path of the slider.
	SliderPath *SliderPath

	// The number of times a player will go over the slider.
	// A value of 1 will not repeat, 2 will repeat once, 3 twice, and so on.
	Repeat int

	// The length of the slider along the path of the described curve.
	// It is specified in osu!pixels, i.e. relative to the 512×384 virtual screen.
	PixelLength float64

	// Applies only to the body of the slider. Only normal (0) and
	// whistle (2) are supported. The samples played are named
	// like soft-sliderslide4.wav for normal, and normal-sliderwhistle.wav for whistle.
	// These samples are meant to be looped, and may also be empty WAV files to mute
	// the slider.
	HitSound HitSound

	// List of HitSounds to apply to the circles of the slider.
	// The values are the same as those for regular hit objects.
	// The list must contain exactly repeat + 1 values, where the first value
	// is the hit sound to play when the slider is first clicked, and the last one
	// when the slider is released.
	EdgeHitSounds []HitSound

	// List of samples sets to apply to the circles of the slider.
	// The list contains exactly repeat + 1 elements. SampleSet and AdditionSet
	// are the same as for hit circles' extras fields.
	EdgeAdditions []*SliderEdgeAddition
}

// String returns string of Slider as it would be in .osu file.
func (s Slider) String() string {
	hitSounds := make([]string, len(s.EdgeHitSounds))
	for i := range hitSounds {
		hitSounds[i] = strconv.Itoa(int(s.EdgeHitSounds[i]))
	}

	additions := make([]string, len(s.EdgeAdditions))
	for i := range additions {
		additions[i] = strconv.Itoa(int(s.EdgeAdditions[i].SampleSet)) + ":" + strconv.Itoa(int(s.EdgeAdditions[i].AdditionSet))
	}
	return strings.Join([]string{
		strconv.Itoa(s.X),
		strconv.Itoa(s.Y),
		strconv.Itoa(s.Time),
		strconv.Itoa(s.Type),
		strconv.Itoa(int(s.HitSound)),
		s.SliderPath.String(),
		strconv.Itoa(s.Repeat),
		fmt.Sprintf("%g", s.PixelLength),
		strings.Join(hitSounds, "|"),
		strings.Join(additions, "|"),
		s.Extras.String(),
	}, ",")
}

// FromString fills Slider fields with data parsed from string.
func (s *Slider) FromString(str string) (err error) {
	attrs := strings.Split(str, ",")

	s.X, err = strconv.Atoi(attrs[0])
	if err != nil {
		return err
	}

	s.Y, err = strconv.Atoi(attrs[1])
	if err != nil {
		return err
	}

	s.Time, err = strconv.Atoi(attrs[2])
	if err != nil {
		return err
	}

	s.Type, err = strconv.Atoi(attrs[3])
	if err != nil {
		return err
	}

	hs, err := strconv.Atoi(attrs[4])
	if err != nil {
		return err
	}
	s.HitSound = HitSound(hs)

	s.SliderPath = new(SliderPath)
	err = s.SliderPath.FromString(attrs[5])
	if err != nil {
		return err
	}

	s.Repeat, err = strconv.Atoi(attrs[6])
	if err != nil {
		return err
	}

	s.PixelLength, err = strconv.ParseFloat(attrs[7], 64)
	if err != nil {
		return err
	}

	if len(attrs) <= 8 {
		return nil
	}

	edgeHitSounds := strings.Split(attrs[8], "|")
	s.EdgeHitSounds = make([]HitSound, len(edgeHitSounds))
	for i := 0; i < len(edgeHitSounds); i++ {
		hs, err = strconv.Atoi(edgeHitSounds[i])
		if err != nil {
			return err
		}
		s.EdgeHitSounds[i] = HitSound(hs)
	}

	edgeAdditions := strings.Split(attrs[9], "|")
	s.EdgeAdditions = make([]*SliderEdgeAddition, len(edgeAdditions))
	for i := 0; i < len(edgeAdditions); i++ {
		s.EdgeAdditions[i] = new(SliderEdgeAddition)
		err = s.EdgeAdditions[i].FromString(edgeAdditions[i])
		if err != nil {
			return err
		}
	}

	s.Extras = new(Extras)
	return s.Extras.FromString(attrs[10])
}

// SliderEdgeAddition is a sample sets to apply to the circles of the slider.
// SampleSet and AdditionSet are the same as for hit circles' extras fields.
type SliderEdgeAddition struct {
	SampleSet   SampleSet
	AdditionSet SampleSet
}

// String returns string of SliderEdgeAddition as it would be in .osu file.
func (sea SliderEdgeAddition) String() string {
	return strconv.Itoa(int(sea.SampleSet)) + ":" + strconv.Itoa(int(sea.AdditionSet))
}

// FromString fills SliderEdgeAddition fields with data parsed from string.
func (sea *SliderEdgeAddition) FromString(str string) (err error) {
	sep := strings.Index(str, ":")

	ss, err := strconv.Atoi(str[:sep])
	if err != nil {
		return err
	}
	sea.SampleSet = SampleSet(ss)

	ss, err = strconv.Atoi(str[:sep])
	if err != nil {
		return err
	}
	sea.AdditionSet = SampleSet(ss)

	return nil
}

// SliderPath specifies path of the slider.
type SliderPath struct {
	SliderType  string
	CurvePoints []*SliderCurvePoint
}

// String returns string of SliderPath as it would be in .osu file.
func (sp SliderPath) String() string {
	points := make([]string, len(sp.CurvePoints))
	for i := range points {
		points[i] = sp.CurvePoints[i].String()
	}
	return sp.SliderType + "|" + strings.Join(points, "|")
}

// FromString fills SliderPath fields SliderPath with data parsed from string.
func (sp *SliderPath) FromString(str string) (err error) {
	attrs := strings.Split(str, "|")
	sp.SliderType = attrs[0]
	sp.CurvePoints = make([]*SliderCurvePoint, len(attrs)-1)
	for i := 1; i < len(attrs); i++ {
		point := new(SliderCurvePoint)
		err = point.FromString(attrs[i])
		if err != nil {
			return err
		}
		sp.CurvePoints[i-1] = point
	}
	return nil
}

// SliderCurvePoint describe a single point of the slider.
type SliderCurvePoint struct {
	X, Y int
}

// String returns string of SliderCurvePoint as it would be in .osu file.
func (cp SliderCurvePoint) String() string {
	return strconv.Itoa(cp.X) + ":" + strconv.Itoa(cp.Y)
}

// FromString fills SliderCurvePoint fields with data parsed from string.
func (cp *SliderCurvePoint) FromString(str string) (err error) {
	sep := strings.Index(str, ":")
	cp.X, err = strconv.Atoi(str[:sep])
	if err != nil {
		return err
	}

	cp.Y, err = strconv.Atoi(str[sep+1:])
	return err
}

// Spinner also creates bananas in Catch the Beat, a spinner in osu!taiko,
// and does not appear in osu!mania. Hit sounds play at the end of the spinner.
// Example:
//  256,192,730,12,8,3983
type Spinner struct {
	BaseHitObject
	EndTime int // When the spinner will end, in milliseconds from the beginning of the song
}

// String returns string of Spinner as it would be in .osu file.
func (s Spinner) String() string {
	if s.Extras == nil {
		return strings.Join([]string{
			strconv.Itoa(s.X),
			strconv.Itoa(s.Y),
			strconv.Itoa(s.Time),
			strconv.Itoa(s.Type),
			strconv.Itoa(int(s.HitSound)),
			strconv.Itoa(s.EndTime),
		}, ",")
	}
	return strings.Join([]string{
		strconv.Itoa(s.X),
		strconv.Itoa(s.Y),
		strconv.Itoa(s.Time),
		strconv.Itoa(s.Type),
		strconv.Itoa(int(s.HitSound)),
		strconv.Itoa(s.EndTime),
		s.Extras.String(),
	}, ",")
}

// FromString fills Spinner fields with data parsed from string.
func (s *Spinner) FromString(str string) (err error) {
	attrs := strings.Split(str, ",")

	s.X, err = strconv.Atoi(attrs[0])
	if err != nil {
		return err
	}

	s.Y, err = strconv.Atoi(attrs[1])
	if err != nil {
		return err
	}

	s.Time, err = strconv.Atoi(attrs[2])
	if err != nil {
		return err
	}

	s.Type, err = strconv.Atoi(attrs[3])
	if err != nil {
		return err
	}

	hs, err := strconv.Atoi(attrs[4])
	if err != nil {
		return err
	}
	s.HitSound = HitSound(hs)

	s.EndTime, err = strconv.Atoi(attrs[5])
	if err != nil {
		return err
	}

	if len(attrs) > 6 {
		s.Extras = new(Extras)
		return s.Extras.FromString(attrs[6])
	}
	return nil
}

// ManiaHoldNote is a hold note unique to osu!mania.
// Example:
//  329,192,16504,128,0,16620:0:0:0:0:
type ManiaHoldNote struct {
	BaseHitObject
	EndTime int
}

// String returns string of ManiaHoldNote as it would be in .osu file.
func (hn ManiaHoldNote) String() string {
	return strings.Join([]string{
		strconv.Itoa(hn.X),
		strconv.Itoa(hn.Y),
		strconv.Itoa(hn.Time),
		strconv.Itoa(hn.Type),
		strconv.Itoa(int(hn.HitSound)),
		strconv.Itoa(hn.EndTime),
	}, ",") + ":" + hn.Extras.String()
}

// FromString fills ManiaHoldNote fields with data parsed from string.
func (hn *ManiaHoldNote) FromString(str string) (err error) {
	attrs := strings.Split(str, ",")

	hn.X, err = strconv.Atoi(attrs[0])
	if err != nil {
		return err
	}

	hn.Y, err = strconv.Atoi(attrs[1])
	if err != nil {
		return err
	}

	hn.Time, err = strconv.Atoi(attrs[2])
	if err != nil {
		return err
	}

	hn.Type, err = strconv.Atoi(attrs[3])
	if err != nil {
		return err
	}

	hs, err := strconv.Atoi(attrs[4])
	if err != nil {
		return err
	}
	hn.HitSound = HitSound(hs)

	// last parameter is joined in mania so we need to split
	sep := strings.Index(attrs[5], ":")

	hn.EndTime, err = strconv.Atoi(attrs[5][:sep])
	if err != nil {
		return err
	}

	hn.Extras = new(Extras)
	return hn.Extras.FromString(attrs[5][sep+1:])
}
