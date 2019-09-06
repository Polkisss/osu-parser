package pcircle

import "testing"

func TestTimingPoint_String(t *testing.T) {
	type fields struct {
		Offset              int
		MillisecondsPerBeat float64
		Meter               int
		SampleSet           SampleSet
		SampleIndex         int
		Volume              int
		Inherited           bool
		Kiai                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Timing Point",
			fields: fields{
				Offset:              66,
				MillisecondsPerBeat: 315.789473684211,
				Meter:               4,
				SampleSet:           SampleSet(2),
				SampleIndex:         0,
				Volume:              45,
				Inherited:           true,
				Kiai:                false,
			},
			want: "66,315.789473684211,4,2,0,45,1,0",
		},
		{
			name: "inherited Timing Point",
			fields: fields{
				Offset:              10171,
				MillisecondsPerBeat: -100,
				Meter:               4,
				SampleSet:           SampleSet(2),
				SampleIndex:         0,
				Volume:              60,
				Inherited:           false,
				Kiai:                true,
			},
			want: "10171,-100,4,2,0,60,0,1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := TimingPoint{
				Offset:              tt.fields.Offset,
				MillisecondsPerBeat: tt.fields.MillisecondsPerBeat,
				Meter:               tt.fields.Meter,
				SampleSet:           tt.fields.SampleSet,
				SampleIndex:         tt.fields.SampleIndex,
				Volume:              tt.fields.Volume,
				Inherited:           tt.fields.Inherited,
				Kiai:                tt.fields.Kiai,
			}
			if got := tp.String(); got != tt.want {
				t.Errorf("TimingPoint.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimingPoint_FromString(t *testing.T) {
	type fields struct {
		Offset              int
		MillisecondsPerBeat float64
		Meter               int
		SampleSet           SampleSet
		SampleIndex         int
		Volume              int
		Inherited           bool
		Kiai                bool
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
			name: "Timing Point",
			fields: fields{
				Offset:              66,
				MillisecondsPerBeat: 315.789473684211,
				Meter:               4,
				SampleSet:           SampleSet(2),
				SampleIndex:         0,
				Volume:              45,
				Inherited:           true,
				Kiai:                false,
			},
			args: args{
				str: "66,315.789473684211,4,2,0,45,1,0",
			},
			wantErr: false,
		},
		{
			name: "inherited Timing Point",
			fields: fields{
				Offset:              10171,
				MillisecondsPerBeat: -100,
				Meter:               4,
				SampleSet:           SampleSet(2),
				SampleIndex:         0,
				Volume:              60,
				Inherited:           false,
				Kiai:                true,
			},
			args: args{
				str: "10171,-100,4,2,0,60,0,1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := &TimingPoint{
				Offset:              tt.fields.Offset,
				MillisecondsPerBeat: tt.fields.MillisecondsPerBeat,
				Meter:               tt.fields.Meter,
				SampleSet:           tt.fields.SampleSet,
				SampleIndex:         tt.fields.SampleIndex,
				Volume:              tt.fields.Volume,
				Inherited:           tt.fields.Inherited,
				Kiai:                tt.fields.Kiai,
			}
			if err := tp.FromString(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("TimingPoint.FromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
