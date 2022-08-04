package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	db "Police-Check-Api/db/sqlc"
	"Police-Check-Api/models"
	"Police-Check-Api/util"

	"github.com/gin-gonic/gin"
)

var queries *db.Queries
var config util.Config

func init() {
	var err error
	config, err = util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
	}
}

func PoliceCheck(c *gin.Context) {
	ctx := context.Background()
	resident := models.Resident{}

	err := c.BindJSON(&resident)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't parse request body to json")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	err2 := conn.Ping()

	if err != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Cannot connect to database")
		return
	}

	queries = db.New(conn)

	var i db.Resident

	i, err = queries.GetResident(ctx, resident.NIC)

	clear := err == nil

	if clear {
		if i.Clearance == "clear" {
			clear = true
		} else {
			clear = false
		}
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"uid":   resident.NIC,
			"clear": clear,
		},
	)
}
