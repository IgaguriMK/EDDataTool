package main

import (
	"github.com/IgaguriMK/EDDataTool/model/coord"
)

type SystemCoord struct {
	ID     int         `json:"id"`
	ID64   int64       `json:"id64"`
	Name   string      `json:"name"`
	Coords coord.Coord `json:"coords"`
	Date   string      `json:"date"`
}
