package pcircle

import (
	"testing"
)

func TestCircle_String(t *testing.T) {
	type fields struct {
		baseHitObject BaseHitObject
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Circle String",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:        164,
					Y:        260,
					Time:     2434,
					Type:     1,
					HitSound: HitSound(0),
					Extras:   &Extras{0, 0, 0, 0, ""},
				},
			},
			want: "164,260,2434,1,0,0:0:0:0:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				BaseHitObject: tt.fields.baseHitObject,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("Circle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_FromString(t *testing.T) {
	type fields struct {
		baseHitObject BaseHitObject
	}
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Circle From String",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:        164,
					Y:        260,
					Time:     2434,
					Type:     1,
					HitSound: HitSound(0),
					Extras:   &Extras{0, 0, 0, 0, ""},
				},
			},
			args:    args{"164,260,2434,1,0,0:0:0:0:"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				BaseHitObject: tt.fields.baseHitObject,
			}
			if err := c.FromString(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("Circle.FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlider_String(t *testing.T) {
	type fields struct {
		baseHitObject BaseHitObject
		SliderPath    *SliderPath
		Repeat        int
		PixelLength   float64
		HitSound      HitSound
		EdgeHitSounds []HitSound
		EdgeAdditions []*SliderEdgeAddition
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Slider String",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:      424,
					Y:      96,
					Time:   66,
					Type:   2,
					Extras: &Extras{0, 0, 0, 0, ""},
				},
				SliderPath: &SliderPath{
					SliderType: "B",
					CurvePoints: []*SliderCurvePoint{
						{380, 120},
						{332, 96},
						{332, 96},
						{304, 124},
					},
				},
				Repeat:      1,
				PixelLength: 130,
				EdgeHitSounds: []HitSound{
					HitSound(2),
					HitSound(0),
				},
				EdgeAdditions: []*SliderEdgeAddition{
					{
						SampleSet:   SampleSet(0),
						AdditionSet: SampleSet(0),
					},
					{
						SampleSet:   SampleSet(0),
						AdditionSet: SampleSet(0),
					},
				},
			},
			want: "424,96,66,2,0,B|380:120|332:96|332:96|304:124,1,130,2|0,0:0|0:0,0:0:0:0:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Slider{
				BaseHitObject: tt.fields.baseHitObject,
				SliderPath:    tt.fields.SliderPath,
				Repeat:        tt.fields.Repeat,
				PixelLength:   tt.fields.PixelLength,
				HitSound:      tt.fields.HitSound,
				EdgeHitSounds: tt.fields.EdgeHitSounds,
				EdgeAdditions: tt.fields.EdgeAdditions,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Slider.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlider_FromString(t *testing.T) {
	type fields struct {
		baseHitObject BaseHitObject
		SliderPath    *SliderPath
		Repeat        int
		PixelLength   float64
		HitSound      HitSound
		EdgeHitSounds []HitSound
		EdgeAdditions []*SliderEdgeAddition
	}
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Slider FromString",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:      424,
					Y:      96,
					Time:   66,
					Type:   2,
					Extras: &Extras{0, 0, 0, 0, ""},
				},
				SliderPath: &SliderPath{
					SliderType: "B",
					CurvePoints: []*SliderCurvePoint{
						{380, 120},
						{332, 96},
						{332, 96},
						{304, 124},
					},
				},
				Repeat:      1,
				PixelLength: 130,
				EdgeHitSounds: []HitSound{
					HitSound(2),
					HitSound(0),
				},
				EdgeAdditions: []*SliderEdgeAddition{
					{
						SampleSet:   SampleSet(0),
						AdditionSet: SampleSet(0),
					},
					{
						SampleSet:   SampleSet(0),
						AdditionSet: SampleSet(0),
					},
				},
			},
			args: args{
				"424,96,66,2,0,B|380:120|332:96|332:96|304:124,1,130,2|0,0:0|0:0,0:0:0:0:",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Slider{
				BaseHitObject: tt.fields.baseHitObject,
				SliderPath:    tt.fields.SliderPath,
				Repeat:        tt.fields.Repeat,
				PixelLength:   tt.fields.PixelLength,
				HitSound:      tt.fields.HitSound,
				EdgeHitSounds: tt.fields.EdgeHitSounds,
				EdgeAdditions: tt.fields.EdgeAdditions,
			}
			if err := s.FromString(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("Slider.FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpinner_String(t *testing.T) {
	type fields struct {
		baseHitObject BaseHitObject
		EndTime       int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Spinner String with extras",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:        256,
					Y:        192,
					Time:     730,
					Type:     12,
					HitSound: HitSound(8),
					Extras:   &Extras{0, 0, 0, 0, "test"},
				},
				EndTime: 3983,
			},
			want: "256,192,730,12,8,3983,0:0:0:0:test",
		},
		{
			name: "Spinner String no extras",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:        256,
					Y:        192,
					Time:     730,
					Type:     12,
					HitSound: HitSound(8),
					Extras:   nil,
				},
				EndTime: 3983,
			},
			want: "256,192,730,12,8,3983",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Spinner{
				BaseHitObject: tt.fields.baseHitObject,
				EndTime:       tt.fields.EndTime,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Spinner.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpinner_FromString(t *testing.T) {
	type fields struct {
		baseHitObject BaseHitObject
		EndTime       int
	}
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Spinner FromString no extras",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:        256,
					Y:        192,
					Time:     730,
					Type:     12,
					HitSound: HitSound(8),
					Extras:   nil,
				},
				EndTime: 3983,
			},
			args:    args{"256,192,730,12,8,3983"},
			wantErr: false,
		},
		{
			name: "Spinner FromString with extras",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:        256,
					Y:        192,
					Time:     730,
					Type:     12,
					HitSound: HitSound(8),
					Extras:   &Extras{0, 0, 0, 0, "test"},
				},
				EndTime: 3983,
			},
			args:    args{"256,192,730,12,8,3983,0:0:0:0:test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spinner{
				BaseHitObject: tt.fields.baseHitObject,
				EndTime:       tt.fields.EndTime,
			}
			if err := s.FromString(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("Spinner.FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManiaHoldNote_String(t *testing.T) {
	type fields struct {
		baseHitObject BaseHitObject
		EndTime       int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ManiaHoldNote String",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:        329,
					Y:        192,
					Time:     16504,
					Type:     128,
					HitSound: HitSound(0),
					Extras:   &Extras{0, 0, 0, 0, ""},
				},
				EndTime: 16620,
			},
			want: "329,192,16504,128,0,16620:0:0:0:0:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hn := ManiaHoldNote{
				BaseHitObject: tt.fields.baseHitObject,
				EndTime:       tt.fields.EndTime,
			}
			if got := hn.String(); got != tt.want {
				t.Errorf("ManiaHoldNote.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManiaHoldNote_FromString(t *testing.T) {
	type fields struct {
		baseHitObject BaseHitObject
		EndTime       int
	}
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ManiaHoldNote FromString",
			fields: fields{
				baseHitObject: BaseHitObject{
					X:        329,
					Y:        192,
					Time:     16504,
					Type:     128,
					HitSound: HitSound(0),
					Extras:   &Extras{0, 0, 0, 0, ""},
				},
				EndTime: 16620,
			},
			args:    args{"329,192,16504,128,0,16620:0:0:0:0:"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hn := &ManiaHoldNote{
				BaseHitObject: tt.fields.baseHitObject,
				EndTime:       tt.fields.EndTime,
			}
			if err := hn.FromString(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("ManiaHoldNote.FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
