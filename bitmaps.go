package osu_parser

import (
	"errors"
)

// Specifies the speed of the countdown which occurs before the first hit object appears.
type Countdown int

const (
	NO_COUNTDOWN Countdown = iota
	NORMAL_COUNTDOWN
	HALF_COUNTDOWN
	DOUBLE_COUNTDOWN
)

// Specifies which set of hit sounds will be used.
type SampleSet int

const (
	AUTO_SAMPLESET SampleSet = iota
	NORMAL_SAMPLESET
	SOFT_SAMPLESET
	DRUM_SAMPLESET
)

// Set the sample set from string.
func (ss *SampleSet) FromString(sample string) error {
	switch sample {
	case "Auto":
		*ss = AUTO_SAMPLESET
	case "Normal":
		*ss = NORMAL_SAMPLESET
	case "Soft":
		*ss = SOFT_SAMPLESET
	case "Drum":
		*ss = DRUM_SAMPLESET
	default:
		return errors.New("invalid sample set identifier: " + sample)
	}
	return nil
}

func (ss SampleSet) String() string {
	return map[SampleSet]string{
		AUTO_SAMPLESET:   "Auto",
		NORMAL_SAMPLESET: "Normal",
		SOFT_SAMPLESET:   "Soft",
		DRUM_SAMPLESET:   "Drum",
	}[ss]
}

// Defines the game mode of the beatmap.
type GameMode int

const (
	OSU_GAMEMODE GameMode = iota
	TAIKO_GAMEMODE
	CTB_GAMEMODE
	MANIA_GAMEMODE
)

// Specifies a hit sounds to play when the hit object is successfully hit.
type HitSound int

const (
	NORMAL_HITSOUND HitSound = 1 << iota
	WHISTLE_HITSOUND
	FINISH_HITSOUND
	CLAP_HITSOUND
)
