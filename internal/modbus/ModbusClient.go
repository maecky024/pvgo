package modbus

import (
	"github.com/gin-gonic/gin"
	"github.com/simonvetter/modbus"
	"log"
	"net/http"
	"time"
)

type JsonResponseKostal struct {
	Timestamp time.Time `json:"timestamp"`
	State     byte      `json:"stage"`
	TotalDC   float32   `json:"totaldc"`
	EastDCC   float32   `json:"eastdcc"`
	EastDCP   float32   `json:"eastdcp"`
	EastDCV   float32   `json:"eastdcv"`
	WestDCC   float32   `json:"westdcc"`
	WestDCP   float32   `json:"westdcp"`
	WestDCV   float32   `json:"westdcv"`
	TotalAC   float32   `json:"totalac"`
	ACP1C     float32   `json:"acp1c"`
	ACP1P     float32   `json:"acp1p"`
	ACP1V     float32   `json:"acp1v"`
	ACP2C     float32   `json:"acp2c"`
	ACP2P     float32   `json:"acp2p"`
	ACP2V     float32   `json:"acp2v"`
	ACP3C     float32   `json:"acp3c"`
	ACP3P     float32   `json:"acp3p"`
	ACP3V     float32   `json:"acp3v"`
}

var statusCodes = map[byte]string{
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

func ReadModbusStatus(c *gin.Context) {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://10.0.0.155:1502",
		Timeout: 1 * time.Second,
	})

	var retval = JsonResponseKostal{}

	err = client.SetUnitId(71)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")

		return
	}
	err = client.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, "")

		return
	}
	err = client.SetEncoding(modbus.BIG_ENDIAN, modbus.LOW_WORD_FIRST)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	kState, err1 := client.ReadBytes(56, 2, modbus.HOLDING_REGISTER)
	retval.State = kState[1]
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fTotalDC, err1 := client.ReadFloat32(100, modbus.HOLDING_REGISTER)
	retval.TotalDC = fTotalDC
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fDCEastC, err1 := client.ReadFloat32(268, modbus.HOLDING_REGISTER)
	retval.EastDCC = fDCEastC
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCEastP, err1 := client.ReadFloat32(270, modbus.HOLDING_REGISTER)
	retval.EastDCP = fDCEastP
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCEastV, err1 := client.ReadFloat32(276, modbus.HOLDING_REGISTER)
	retval.EastDCV = fDCEastV
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCWestC, err1 := client.ReadFloat32(258, modbus.HOLDING_REGISTER)
	retval.WestDCC = fDCWestC
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCWestP, err1 := client.ReadFloat32(260, modbus.HOLDING_REGISTER)
	retval.WestDCP = fDCWestP
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fDCWestV, err1 := client.ReadFloat32(266, modbus.HOLDING_REGISTER)
	retval.WestDCV = fDCWestV
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fTotalAC, err1 := client.ReadFloat32(172, modbus.HOLDING_REGISTER)
	retval.TotalAC = fTotalAC
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fACP1C, err1 := client.ReadFloat32(154, modbus.HOLDING_REGISTER)
	retval.ACP1C = fACP1C
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP1P, err1 := client.ReadFloat32(156, modbus.HOLDING_REGISTER)
	retval.ACP1P = fACP1P
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP1V, err1 := client.ReadFloat32(158, modbus.HOLDING_REGISTER)
	retval.ACP1V = fACP1V
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fACP2C, err1 := client.ReadFloat32(160, modbus.HOLDING_REGISTER)
	retval.ACP2C = fACP2C
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP2P, err1 := client.ReadFloat32(162, modbus.HOLDING_REGISTER)
	retval.ACP2P = fACP2P
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP2V, err1 := client.ReadFloat32(164, modbus.HOLDING_REGISTER)
	retval.ACP2V = fACP2V
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	fACP3C, err1 := client.ReadFloat32(166, modbus.HOLDING_REGISTER)
	retval.ACP3C = fACP3C
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP3P, err1 := client.ReadFloat32(168, modbus.HOLDING_REGISTER)
	retval.ACP3P = fACP3P
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}
	fACP3V, err1 := client.ReadFloat32(170, modbus.HOLDING_REGISTER)
	retval.ACP3V = fACP3V
	if err1 != nil {
		log.Fatal("ERR1: ", err1)
	}

	retval.Timestamp = time.Now()

	defer func(client *modbus.ModbusClient) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	c.JSON(http.StatusOK, retval)
}
