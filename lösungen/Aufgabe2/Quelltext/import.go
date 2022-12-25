package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func importCrystals() (field Field) {
	data, err := os.ReadFile("template.txt")
	if err != nil {
		log.Panic(err)
	}
	var lines = strings.Split(strings.Replace(string(data), "\r", "", -1), "\n")
	sizeX, err := strconv.Atoi(strings.Split(lines[0], ";")[0])
	if err != nil {
		log.Panic(err)
	}
	sizeY, err := strconv.Atoi(strings.Split(lines[0], ";")[1])
	if err != nil {
		log.Panic(err)
	}
	field = NewField(Vector{sizeX, sizeY})
	for _, line := range lines[1:] {
		var splitLine = strings.Split(line, ";")
		posX, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Panic(err)
		}
		posY, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Panic(err)
		}
		speedPlusX, err := strconv.Atoi(splitLine[2])
		if err != nil {
			log.Panic(err)
		}
		speedPlusY, err := strconv.Atoi(splitLine[3])
		if err != nil {
			log.Panic(err)
		}
		speedMinusX, err := strconv.Atoi(splitLine[4])
		if err != nil {
			log.Panic(err)
		}
		speedMinusY, err := strconv.Atoi(splitLine[5])
		if err != nil {
			log.Panic(err)
		}
		orientation, err := strconv.Atoi(splitLine[6])
		if err != nil {
			log.Panic(err)
		}
		var seed = Seed{
			orientation: uint8(orientation),
			plusSpeed:   Vector{speedPlusX, speedPlusY},
			minusSpeed:  Vector{speedMinusX, speedMinusY},
		}
		var crystal = NewCrystal(seed)
		field.AddCrystal(crystal, Vector{posX, posY})
	}
	return field
}
