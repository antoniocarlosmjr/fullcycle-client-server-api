package repository

import (
	"context"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"

	"github.com/fullcycle-client-server-api/server/models"
	"gorm.io/driver/sqlite"
)

const TimeoutSaveDatabase = time.Millisecond * 10

type QuotationEntity struct {
	Bid float32
	gorm.Model
}

func connect() (*gorm.DB, error) {
	conn, err := gorm.Open(sqlite.Open("./repository/quotation.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = conn.AutoMigrate(&QuotationEntity{})
	if err != nil {
		log.Panicf("Data migration failed: %s", err)
	}

	return conn, nil
}

func SaveQuotation(awesomeData models.AwesomeAPIResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), TimeoutSaveDatabase)
	defer cancel()

	db, err := connect()
	if err != nil {
		log.Panicf("Connection With Database failed: %s", err)
	}

	bid, err := strconv.ParseFloat(awesomeData.USDBRL.Bid, 32)
	if err != nil {
		log.Panicf("Invalid value received: %s", err)
	}

	entity := QuotationEntity{Bid: float32(bid)}

	tx := db.WithContext(ctx).Create(&entity)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
