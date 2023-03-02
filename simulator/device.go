package simulator

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Device struct {
	id         string
	coordinate *Coordinate
	speed      int
	seen       map[string]int
}

func NewDevice(area *Rectangle) *Device {

	return &Device{
		id:         uuid.NewString(),
		coordinate: newDeviceCoordinate(area),
		speed:      rand.Intn(10) + 1,
		seen:       make(map[string]int),
	}
}

func newDeviceCoordinate(area *Rectangle) *Coordinate {

	return &Coordinate{
		Latitute:  area.coordinate.Latitute + float64(rand.Intn(area.breadth))/10000,
		Longitude: area.coordinate.Longitude + float64(rand.Intn(area.length))/10000,
	}
}

func (d *Device) Coordinate() *Coordinate {

	return d.coordinate
}

func (d *Device) Scan(tick int, nearby func(*Coordinate) []*Device) {

	for _, currDevice := range nearby(d.coordinate) {
		if d.id != currDevice.id {
			if lastSeen, ok := d.seen[seenId(currDevice)]; (ok && tick-lastSeen > 3) || !ok {
				log.Infof("[SCAN] device %s coordinate %s near %s coordinate %s", d.id, d.coordinate, currDevice.id, currDevice.coordinate)
				d.seen[seenId(currDevice)] = tick
			}
		}
	}
}

func (d *Device) Move() {

	dest := &Coordinate{
		Latitute:  d.coordinate.Latitute + step(d.speed),
		Longitude: d.coordinate.Longitude + step(d.speed),
	}
	log.Infof("[MOVE] device %s (speed %d) is moving %s -> %s", d.id, d.speed, d.coordinate, dest)
	d.coordinate = dest
}

func step(speed int) float64 {

	return float64(rand.Intn(speed)) / 1000
}

func seenId(device *Device) string {

	return fmt.Sprintf("%s,%s", device.id, device.coordinate)
}
