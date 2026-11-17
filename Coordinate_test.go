package geo_test

import (
	"reflect"
	"testing"

	"github.com/thereisnoplanb/geo"
)

func TestCoordinate_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		latitude  float64
		longitude float64
		want      []byte
		wantErr   bool
	}{
		{
			name:      "valid coordinate",
			latitude:  37.7749,
			longitude: -122.4194,
			want:      []byte(`{"latitude":37.7749,"longitude":-122.4194}`),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this, err := geo.NewCoordinate(tt.latitude, tt.longitude)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got, gotErr := this.MarshalJSON()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("MarshalJSON() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("MarshalJSON() succeeded unexpectedly")
			}
			if string(got) != string(tt.want) {
				t.Errorf("MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    []byte
		wantErr bool
	}{
		{
			name:    "valid coordinate",
			data:    []byte(`{"latitude":37.7749,"longitude":-122.4194}`),
			wantErr: false,
		},
		{
			name:    "invalid coordinate",
			data:    []byte(`{"latitude":100.0,"longitude":200.0}`),
			wantErr: true,
		},
		{
			name:    "missing latitude",
			data:    []byte(`{"longitude":-122.4194}`),
			wantErr: true,
		},
		{
			name:    "missing longitude",
			data:    []byte(`{"latitude":37.7749}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var this geo.Coordinate
			gotErr := this.UnmarshalJSON(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UnmarshalJSON() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UnmarshalJSON() succeeded unexpectedly")
			}
		})
	}
}

func TestCoordinate_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		latitude  float64
		longitude float64
		want      string
	}{
		{
			name:      "valid coordinate",
			latitude:  37.7749,
			longitude: -122.4194,
			want:      "37.7749,-122.4194",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this, err := geo.NewCoordinate(tt.latitude, tt.longitude)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := this.String()
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		coordinate1 geo.Coordinate
		coordinate2 geo.Coordinate
		want        float64
	}{
		{
			name:        "valid coordinates",
			coordinate1: geo.MustCoordinate(37.7749, -122.4194),
			coordinate2: geo.MustCoordinate(34.0522, -118.2437),
			want:        559.1205770615533,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := geo.Distance(tt.coordinate1, tt.coordinate2)
			if got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_Distance(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		latitude  float64
		longitude float64
		// Named input parameters for target function.
		other geo.Coordinate
		want  float64
	}{
		{
			name:      "valid distance",
			latitude:  37.7749,
			longitude: -122.4194,
			other:     geo.MustCoordinate(34.0522, -118.2437),
			want:      559.1205770615533,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this, err := geo.NewCoordinate(tt.latitude, tt.longitude)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := this.Distance(tt.other)
			if got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCoordinate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		latitude  float64
		longitude float64
		want      geo.Coordinate
		wantErr   bool
	}{
		{
			name:      "valid coordinate",
			latitude:  37.7749,
			longitude: -122.4194,
			want:      geo.MustCoordinate(37.7749, -122.4194),
			wantErr:   false,
		},
		{
			name:      "invalid coordinate",
			latitude:  100.0, // invalid latitude
			longitude: -122.4194,
			want:      geo.Coordinate{},
			wantErr:   true,
		},
		{
			name:      "invalid coordinate",
			latitude:  37.7749,
			longitude: -200.0, // invalid longitude
			want:      geo.Coordinate{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := geo.NewCoordinate(tt.latitude, tt.longitude)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewCoordinate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewCoordinate() succeeded unexpectedly")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_Longitude(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		latitude  float64
		longitude float64
		want      float64
	}{
		{
			name:      "valid longitude",
			latitude:  37.7749,
			longitude: -122.4194,
			want:      -122.4194,
		},
		{
			name:      "valid longitude 2",
			latitude:  34.0522,
			longitude: -118.2437,
			want:      -118.2437,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this, err := geo.NewCoordinate(tt.latitude, tt.longitude)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := this.Longitude()
			if got != tt.want {
				t.Errorf("Longitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_Latitude(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		latitude  float64
		longitude float64
		want      float64
	}{
		{
			name:      "valid latitude",
			latitude:  37.7749,
			longitude: -122.4194,
			want:      37.7749,
		},
		{
			name:      "valid latitude 2",
			latitude:  34.0522,
			longitude: -118.2437,
			want:      34.0522,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this, err := geo.NewCoordinate(tt.latitude, tt.longitude)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := this.Latitude()
			if got != tt.want {
				t.Errorf("Latitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_Equals(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		latitude  float64
		longitude float64
		// Named input parameters for target function.
		other geo.Coordinate
		want  bool
	}{
		{
			name:      "equal coordinates",
			latitude:  37.7749,
			longitude: -122.4194,
			other:     geo.MustCoordinate(37.7749, -122.4194),
			want:      true,
		},
		{
			name:      "different coordinates",
			latitude:  37.7749,
			longitude: -122.4194,
			other:     geo.MustCoordinate(34.0522, -118.2437),
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this, err := geo.NewCoordinate(tt.latitude, tt.longitude)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := this.Equals(tt.other)
			if got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
