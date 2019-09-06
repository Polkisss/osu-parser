package pcircle

import (
	"testing"
)

func TestBreak_String(t *testing.T) {
	type fields struct {
		StartTime int
		EndTime   int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Break",
			fields: fields{
				StartTime: 4627,
				EndTime:   5743,
			},
			want: "2,4627,5743",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Break{
				StartTime: tt.fields.StartTime,
				EndTime:   tt.fields.EndTime,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("Break.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBreak_FromString(t *testing.T) {
	type fields struct {
		StartTime int
		EndTime   int
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
			name: "Break",
			fields: fields{
				StartTime: 4627,
				EndTime:   5743,
			},
			args: args{
				str: "2,4627,5743",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Break{
				StartTime: tt.fields.StartTime,
				EndTime:   tt.fields.EndTime,
			}
			if err := b.FromString(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("Break.FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRGB_String(t *testing.T) {
	type fields struct {
		R int
		G int
		B int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "RGB",
			fields: fields{123, 321, 43},
			want:   "123,321,43",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := RGB{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("RGB.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRGB_FromString(t *testing.T) {
	type fields struct {
		R int
		G int
		B int
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
			name:   "RGB",
			fields: fields{123, 321, 43},
			args: args{
				str: "123,321,43",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &RGB{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			if err := c.FromString(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("RGB.FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExtras_String(t *testing.T) {
	type fields struct {
		SampleSet     SampleSet
		AdditionalSet SampleSet
		CustomIndex   int
		SampleVolume  int
		Filename      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty Extras",
			fields: fields{
				SampleSet:     SampleSet(0),
				AdditionalSet: SampleSet(0),
				CustomIndex:   0,
				SampleVolume:  0,
				Filename:      "",
			},
			want: "0:0:0:0:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Extras{
				SampleSet:     tt.fields.SampleSet,
				AdditionalSet: tt.fields.AdditionalSet,
				CustomIndex:   tt.fields.CustomIndex,
				SampleVolume:  tt.fields.SampleVolume,
				Filename:      tt.fields.Filename,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("Extras.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtras_FromString(t *testing.T) {
	type fields struct {
		SampleSet     SampleSet
		AdditionalSet SampleSet
		CustomIndex   int
		SampleVolume  int
		Filename      string
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
			name: "empty Extras",
			fields: fields{
				SampleSet:     SampleSet(0),
				AdditionalSet: SampleSet(0),
				CustomIndex:   0,
				SampleVolume:  0,
				Filename:      "",
			},
			args: args{
				str: "0:0:0:0:",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Extras{
				SampleSet:     tt.fields.SampleSet,
				AdditionalSet: tt.fields.AdditionalSet,
				CustomIndex:   tt.fields.CustomIndex,
				SampleVolume:  tt.fields.SampleVolume,
				Filename:      tt.fields.Filename,
			}
			if err := e.FromString(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("Extras.FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBackground_String(t *testing.T) {
	type fields struct {
		FileName string
		XOffset  int
		YOffset  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Background String",
			fields: fields{
				FileName: "bg.jpg",
				XOffset:  50,
				YOffset:  100,
			},
			want: `0,0,"bg.jpg",50,100`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Background{
				FileName: tt.fields.FileName,
				XOffset:  tt.fields.XOffset,
				YOffset:  tt.fields.YOffset,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("Background.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBackground_FromString(t *testing.T) {
	type fields struct {
		FileName string
		XOffset  int
		YOffset  int
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
			name: "Background FromString with double quotes",
			fields: fields{
				FileName: "bg.jpg",
				XOffset:  50,
				YOffset:  100,
			},
			args:    args{str: `0,0,"bg.jpg",50,100`},
			wantErr: false,
		},
		{
			name: "Background FromString without double quotes",
			fields: fields{
				FileName: "bg.jpg",
				XOffset:  50,
				YOffset:  100,
			},
			args:    args{str: `0,0,bg.jpg,50,100`},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Background{
				FileName: tt.fields.FileName,
				XOffset:  tt.fields.XOffset,
				YOffset:  tt.fields.YOffset,
			}
			if err := b.FromString(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("Background.FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
