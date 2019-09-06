package pcircle

import (
	"errors"
)

// Countdown specifies the speed of the countdown which occurs
// before the first hit object appears.

// Countdown speed and length.
const (
	NO_COUNTDOWN = iota
	NORMAL_COUNTDOWN
	HALF_COUNTDOWN
	DOUBLE_COUNTDOWN
)

// SampleSet specifies which set of hit sounds will be used.
type SampleSet int

// All possible sample sets.
const (
	AUTO_SAMPLESET SampleSet = iota
	NORMAL_SAMPLESET
	SOFT_SAMPLESET
	DRUM_SAMPLESET
)

// FromString allows you to set sample sets with strings.
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

// String returns string of SampleSet in readable format.
func (ss SampleSet) String() string {
	return map[SampleSet]string{
		AUTO_SAMPLESET:   "Auto",
		NORMAL_SAMPLESET: "Normal",
		SOFT_SAMPLESET:   "Soft",
		DRUM_SAMPLESET:   "Drum",
	}[ss]
}

// GameMode defines the game mode of the beatmap.
type GameMode int

// All possible gamemodes
const (
	OSU_GAMEMODE GameMode = iota
	TAIKO_GAMEMODE
	CTB_GAMEMODE
	MANIA_GAMEMODE
)

// HitSound specifies a hit sounds to play when the hit object is successfully hit.
type HitSound int

// All possible hit sounds
const (
	NO_HITSOUND     HitSound = 0
	NORMAL_HITSOUND HitSound = 1 << iota
	WHISTLE_HITSOUND
	FINISH_HITSOUND
	CLAP_HITSOUND
)
