package simulator

import "fmt"

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func (l *Coordinate) String() string {

	return fmt.Sprintf("%.4f,%.4f", l.Latitude, l.Longitude)
}
