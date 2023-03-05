package simulator

import "math"

type Realm struct {
	area    *Rectangle
	devices []*Device
}

type Rectangle struct {
	coordinate *Coordinate
	length     int
	breadth    int
}

func NewRealm(devices int) *Realm {

	rect := &Rectangle{coordinate: &Coordinate{Latitude: 32.0577, Longitude: 34.7666}, length: 200, breadth: 200}
	return &Realm{
		area:    rect,
		devices: createDevices(devices, rect),
	}
}

func (r *Realm) Run(ticks int) {

	for i := 1; i <= ticks; i++ {
		for _, currDevice := range r.devices {
			currDevice.Scan(i, r.nearby)
			currDevice.Move()
		}
	}
}

func createDevices(count int, area *Rectangle) []*Device {

	res := []*Device{}
	for i := 1; i <= count; i++ {
		res = append(res, NewDevice(area))
	}

	return res
}

func (r *Realm) nearby(coordinate *Coordinate) []*Device {

	res := []*Device{}
	for _, currDevice := range r.devices {
		if distance(coordinate, currDevice.coordinate) < 0.01 {
			res = append(res, currDevice)
		}
	}

	return res
}

func distance(c1 *Coordinate, c2 *Coordinate) float64 {

	radlat1 := float64(math.Pi * c1.Latitude / 180)
	radlat2 := float64(math.Pi * c2.Latitude / 180)

	theta := float64(c1.Longitude - c2.Longitude)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	return dist * 1.609344
}
