package geo

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

type Coordinate struct {
	latitude  float64
	longitude float64
}

// Returns the latitude of the coordinate.
//
// # Returns
//
//	latitude float64
//
//	The latitude of the coordinate.
func (this Coordinate) Latitude() (latitude float64) {
	return this.latitude
}

// Returns the longitude of the coordinate.
//
// # Returns
//
//	longitude float64
//
//	The longitude of the coordinate.
func (this Coordinate) Longitude() (longitude float64) {
	return this.longitude
}

// Create a new Coordinate instance with the given latitude and longitude.
//
// # Parameters
//
//	latitude float64
//
// The latitude of the coordinate.
//
//	longitude float64
//
// The longitude of the coordinate.
//
// # Returns
//
//	coordinate Coordinate
//
//	The created Coordinate instance if no error occurred.
//
//	err error
//
// Error will be non-nil if the provided values are invalid.
func NewCoordinate(latitude, longitude float64) (coordinate Coordinate, err error) {
	if err := validateCoordinate(latitude, longitude); err != nil {
		return Coordinate{}, err
	}
	coordinate.latitude = latitude
	coordinate.longitude = longitude
	return coordinate, nil
}

// Creates a new Coordinate instance and panics if the values are invalid.
//
// # Parameters
//
//	latitude float64
//
// The latitude of the coordinate.
//
//	longitude float64
//
// The longitude of the coordinate.
//
// # Returns
//
//	coordinate Coordinate
//
//	The created Coordinate instance.
func MustCoordinate(latitude, longitude float64) (coordinate Coordinate) {
	coordinate, err := NewCoordinate(latitude, longitude)
	if err != nil {
		panic(err)
	}
	return coordinate
}

func validateCoordinate(lat, lon float64) error {
	if math.IsNaN(lat) || lat < -90 || lat > 90 {
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if math.IsNaN(lon) || lon < -180 || lon > 180 {
		return fmt.Errorf("invalid longitude: %f", lon)
	}
	return nil
}

// Calculate the distance to another coordinate using the Haversine formula.
//
// # Parameters
//
//	other Coordinate
//
// The other coordinate to calculate the distance to.
//
// # Returns
//
//	distance float64
//
//	The distance between the two coordinates in kilometers.
//
// # Remarks
//
// The distance is calculated using the Haversine formula, which assumes a spherical Earth.
// This provides an approximation and may not be accurate for very long distances.
// Altitude differences are not considered in this calculation.
func (this Coordinate) Distance(other Coordinate) (distance float64) {
	const earthDiameter = 2 * 6371
	rad := math.Pi / 180
	lat1 := this.latitude * rad
	lat2 := other.latitude * rad
	sinDLat := math.Sin((other.latitude - this.latitude) * rad / 2)
	sinDLon := math.Sin((other.longitude - this.longitude) * rad / 2)
	a := (sinDLat * sinDLat) + math.Cos(lat1)*math.Cos(lat2)*(sinDLon*sinDLon)
	if a > 1 {
		a = 1
	}
	if a < 0 {
		a = 0
	}
	return earthDiameter * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

// Calculate the distance between two coordinates using the Haversine formula.
//
// # Parameters
//
//	coordinate1 Coordinate
//
// The first coordinate to calculate the distance from.
//
//	coordinate2 Coordinate
//
// The second coordinate to calculate the distance to.
//
// # Returns
//
//	distance float64
//
//	The distance between the two coordinates in kilometers.
//
// # Remarks
//
// The distance is calculated using the Haversine formula, which assumes a spherical Earth.
// This provides an approximation and may not be accurate for very long distances.
// Altitude differences are not considered in this calculation.
func Distance(coordinate1, coordinate2 Coordinate) float64 {
	return coordinate1.Distance(coordinate2)
}

// Returns a string that contains the latitude and longitude.
//
// # Returns
//
//	string
//
//	A string that contains the latitude and longitude, separated by a comma.
//
// # Remarks
//
// The string representation returned by this method is intended only for debugging use.
// This method does not return any locale-specific formatting for latitude and longitude.
func (this Coordinate) String() string {
	return fmt.Sprintf("%s,%s", strconv.FormatFloat(this.latitude, 'f', -1, 64), strconv.FormatFloat(this.longitude, 'f', -1, 64))
}

type coordinateJSON struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}

func (this coordinateJSON) Validate() error {
	if this.Latitude == nil {
		return fmt.Errorf("latitude is required")
	}
	if this.Longitude == nil {
		return fmt.Errorf("longitude is required")
	}
	return validateCoordinate(*this.Latitude, *this.Longitude)
}

// MarshalJSON implements the json.Marshaler interface for the Coordinate type.
// It converts the Coordinate into a JSON representation.
func (this Coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(coordinateJSON{
		Latitude:  &this.latitude,
		Longitude: &this.longitude,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface for the Coordinate type.
// It parses the JSON representation and populates the Coordinate fields.
func (this *Coordinate) UnmarshalJSON(data []byte) error {
	var coordinate coordinateJSON
	if err := json.Unmarshal(data, &coordinate); err != nil {
		return err
	}
	if err := coordinate.Validate(); err != nil {
		return err
	}
	this.latitude = *coordinate.Latitude
	this.longitude = *coordinate.Longitude
	return nil
}

// Compares the current Coordinate with another Coordinate for equality.
//
// # Parameters
//
//	other Coordinate
//
// The Coordinate to compare with.
//
// # Returns
//
//	bool
//
//	True if the coordinates are equal, false otherwise.
func (this Coordinate) Equals(other Coordinate) bool {
	return this.latitude == other.latitude &&
		this.longitude == other.longitude
}
