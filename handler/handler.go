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

var config util.Config
var conn *sql.DB

func init() {
	var err error

	config, err = util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config")

	}
	conn, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Println("HELP")
		return
	}
}

func PoliceCheck(c *gin.Context) {
	ctx := context.Background()
	resident := models.Resident{}

	err := c.BindJSON(&resident)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't parse request body to json")
	}

	if err := conn.Ping(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Cannot connect to database")
		return
	}

	queries := db.New(conn)

	i, err := queries.GetResident(ctx, resident.NIC)

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
			"nic":   resident.NIC,
			"clear": clear,
		},
	)
}
