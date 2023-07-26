package umshandlers

import "github.com/gin-gonic/gin"

type GeoPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type InsertStrongholdParam struct {
	Points        []GeoPoint `json:"points,omitempty"`         //it's points of stronghold's outlines.
	Name          string     `json:"name,omitempty"`           //stronghold's name
	DetailAddress string     `json:"detail_address,omitempty"` //it's a detail address info about stronghold.
	HotLevel      int8       `json:"hot_level,omitempty"`      //it's popular level about stronghold.
}

// InsertStronghold
// insert a stronghold into system. Those data will be store into MySQL, and cache in Redis.
func InsertStronghold(c *gin.Context) {

}
