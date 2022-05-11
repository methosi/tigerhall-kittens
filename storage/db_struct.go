package storage

import "time"

type Tiger struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	DOB         string    `db:"dob"`
	LastSeenAt  time.Time `db:"last_seen_at"`
	LastSeenLat float64   `db:"last_seen_lat"`
	LastSeenLon float64   `db:"last_seen_long"`
	Created     time.Time `db:"created"`
	Updated     time.Time `db:"updated"`
}

type TigerSighting struct {
	SightingLat float64   `db:"sighting_lat"`
	SightingLon float64   `db:"sighting_lon"`
	Image       string    `db:"sighting_image"`
	Created     time.Time `db:"created"`
	Updated     time.Time `db:"updated"`
}
