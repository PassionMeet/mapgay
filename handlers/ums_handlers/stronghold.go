package umshandlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	param := InsertStrongholdParam{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("InsertStronghold %s", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if len(param.Points) < 3 {
		// if points less than 3, it can't build a polygon.
		log.Println("InsertStronghold param'Points length is less than 3, the points can't build a polygon")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

}

func GetStrongholds(ctx *gin.Context) {

}

func GetStrongholdDetail(ctx *gin.Context) {

}

func DelStronghold(ctx *gin.Context) {

}

func UpdateStronghold(ctx *gin.Context) {

}
