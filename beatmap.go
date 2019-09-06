package pcircle

// Beatmap stores information about single beatmap.
type Beatmap struct {
	FilePath string // The location of beatmap .osu file

	// General
	//
	// Various properties about the beatmap's gameplay.
	AudioFilename        string    // The location of the audio file relative to the current folder
	AudioLeadIn          int       // The amount of time added before the audio file begins playing
	PreviewTime          int       // Defines when the audio file should begin playing when selected in the song selection menu
	Countdown            Countdown // The speed of the countdown which occurs before the first hit object appears
	SampleSet            SampleSet // Specifies which set of hit sounds will be used throughout the beatmap
	StackLeniency        float64   // How often closely placed hit objects will be stacked together
	GameMode             GameMode  // Defines the game mode of the beatmap (0=osu!, 1=Taiko, 2=Catch the Beat, 3=osu!mania)
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
	Background string  // The location of the background image relative to the beatmap directory
	Breaks     []Break // Break times through the beatmap
	// todo: storyboards

	// Timing Points
	//
	// A list of the beatmap's timing points and hitsounds.
	// Describes a number of properties regarding beats per minute and
	// hit sounds. Sorted by offset in the timing points section.
	TimingPoints []TimingPoint

	// Colours
	//
	// RGB values of the combo colours used.
	Combos []RGB // Defines the colours of combos

	SliderBody          RGB
	SliderTrackOverride RGB
	SliderBorder        RGB
	// Extra colours for sliders

	// Hit Objects
	//
	// A list of the beatmap's hit objects.
	HitObjects []interface{}
}
