package iim42652

import (
	"encoding/hex"
	"fmt"
)

const ShortMax = 32767
const PowerOnSleep = 250
const ReadMask byte = 0x80

type Bank byte

const (
	Bank0 Bank = 0x00
	Bank1 Bank = 0x01
	Bank2 Bank = 0x02
	Bank4 Bank = 0x04
)

func (b Bank) ToBytes() []byte {
	return []byte{byte(b)}
}

func (b Bank) String() string {
	return hex.EncodeToString(b.ToBytes())
}

type Address byte

func (r Address) ToBytes() []byte {
	return []byte{byte(r)}
}

func (r Address) String() string {
	return hex.EncodeToString(r.ToBytes())
}

type Register struct {
	Bank    Bank
	Address Address
}

func (r *Register) String() string {
	return fmt.Sprintf("bank %s addr:%s", r.Bank, r.Address)
}

var (
	RegisterBankSel            = &Register{Bank0, 0x76}
	RegisterIntfConfig1        = &Register{Bank0, 0x4D}
	RegisterAccelDataX1        = &Register{Bank0, 0x1f}
	RegisterPwrMgmt0           = &Register{Bank0, 0x4e}
	RegisterAccelWomXThreshold = &Register{Bank4, 0x4a}
	RegisterAccelWomYThreshold = &Register{Bank4, 0x4b}
	RegisterAccelWomZThreshold = &Register{Bank4, 0x4c}

	RegisterIntSource1 = &Register{Bank0, 0x66}
	RegisterIntSource4 = &Register{Bank0, 0x69}

	RegisterSdmConfig0  = &Register{Bank0, 0x57}
	RegisterIntStatus2  = &Register{Bank0, 0x37}
	RegisterAccelConfig = &Register{Bank0, 0x50}

	RegisterAntiAliasFilterDelta    = &Register{Bank2, 0x03} // bits 6:1, ACCEL_AAF_DELT: Code from 1 to 63 that allows programming the bandwidth for accelerometer anti-alias filter
	RegisterAntiAliasFilterDeltaSqr = &Register{Bank2, 0x04} //bits 7:0 and Bank 2, register 0x05h, bits 3:0, ACCEL_AAF_DELTSQR: Square of the delt value for accelerometer
	RegisterAntiAliasFilterBitshift = &Register{Bank2, 0x05} //bits 7:4, ACCEL_AAF_BITSHIFT: Bitshift value for accelerometer used in hardware implementation

	// RegisterGyroscopeDataX1 GYRO_DATA_X1 (0x25), GYRO_DATA_X0 (0x26), GYRO_DATA_Y1 (0x27), GYRO_DATA_Y0 (0x28), GYRO_DATA_Z1 (0x29), GYRO_DATA_Z0 (0x2A)
	RegisterGyroscopeDataX1 = &Register{Bank0, 0x25}

	RegisterTemperatureData = &Register{Bank0, 0x1D}
)

const (
	GyroModeLowNoise          byte = 0x0c
	AccelerometerModeLowPower byte = 0x02
	AccelerometerModeLowNoise byte = 0x03
)

const (
	ConfigScaleMask  byte  = 0x07
	ConfigScaleShift uint8 = 5
	ConfigRateMask   byte  = 0x0f
)

//const (
//	TempDataSize  uint8 = 2
//	AccelDataSize uint8 = 6
//	GyroDataSize  uint8 = 6
//)

const (
	Dps2000 uint16 = iota
	Dps1000
	Dps500
	Dps250
	Dps125
	Dps62_5
	Dps31_25
	Dps15_625
)
