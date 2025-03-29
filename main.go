package main

import (
	"fmt"
	"github.com/simonvetter/modbus"
	"log"
	"time"
)

func main() {
	statusCodes := map[byte]string{
		0:  "Aus",
		1:  "Initialisiere",
		2:  "IsoMeas",
		3:  "GridCheck",
		4:  "StartUp",
		5:  "-",
		6:  "Produziert",
		7:  "Throttled",
		8:  "ExtSwitchOff",
		9:  "Update",
		10: "StandBy",
		11: "GridSync",
		12: "GridPreCheck",
		13: "GridSwitchOff",
		14: "Overheating",
		15: "Aus",
		16: "ImproperDcVoltage",
		17: "ESB",
		18: "Unknown",
	}

	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://10.0.0.155:1502",
		Timeout: 1 * time.Second,
	})
	err = client.SetUnitId(71)
	if err != nil {
		return
	}
	err = client.Open()
	if err != nil {
		return
	}
	err = client.SetEncoding(modbus.BIG_ENDIAN, modbus.LOW_WORD_FIRST)
	if err != nil {
		return
	}

	kState, err1 := client.ReadBytes(56, 2, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fTotalDC, err1 := client.ReadFloat32(100, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fDCEastC, err1 := client.ReadFloat32(268, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCEastP, err1 := client.ReadFloat32(270, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCEastV, err1 := client.ReadFloat32(276, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCWestC, err1 := client.ReadFloat32(258, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCWestP, err1 := client.ReadFloat32(260, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCWestV, err1 := client.ReadFloat32(266, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fTotalAC, err1 := client.ReadFloat32(172, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fACP1C, err1 := client.ReadFloat32(154, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP1P, err1 := client.ReadFloat32(156, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP1V, err1 := client.ReadFloat32(158, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fACP2C, err1 := client.ReadFloat32(160, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP2P, err1 := client.ReadFloat32(162, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP2V, err1 := client.ReadFloat32(164, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fACP3C, err1 := client.ReadFloat32(166, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP3P, err1 := client.ReadFloat32(168, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP3V, err1 := client.ReadFloat32(170, modbus.HOLDING_REGISTER)
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fmt.Printf("Status: %s\n", statusCodes[kState[1]])
	fmt.Printf("fTotalDC: %f\n", fTotalDC)
	fmt.Printf("fTotalDC: %f\n", fTotalAC)

	fmt.Printf("fDCEastC: %f\n", fDCEastC)
	fmt.Printf("fDCEastP: %f\n", fDCEastP)
	fmt.Printf("fDCEastV: %f\n", fDCEastV)

	fmt.Printf("fDCWestC: %f\n", fDCWestC)
	fmt.Printf("fDCWestP: %f\n", fDCWestP)
	fmt.Printf("fDCWestV: %f\n", fDCWestV)

	fmt.Printf("fACP1C: %f\n", fACP1C)
	fmt.Printf("fACP1P: %f\n", fACP1P)
	fmt.Printf("fACP1V: %f\n", fACP1V)

	fmt.Printf("fACP2C: %f\n", fACP2C)
	fmt.Printf("fACP2P: %f\n", fACP2P)
	fmt.Printf("fACP2V: %f\n", fACP2V)

	fmt.Printf("fACP3C: %f\n", fACP3C)
	fmt.Printf("fACP3P: %f\n", fACP3P)
	fmt.Printf("fACP3V: %f\n", fACP3V)

	defer client.Close()
}
