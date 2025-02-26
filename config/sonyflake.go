//go:build sonyflake

package config

type Sonyflake struct {
	Start     string `json:"Start" yaml:"Start" mapstructure:"Start"`             // 开始时间
	MachineId uint16 `json:"MachineId" yaml:"MachineId" mapstructure:"MachineId"` // 机器ID
}

func (c *Sonyflake) MachineID() func() (uint16, error) {
	return func() (uint16, error) {
		return c.MachineId, nil
	}
}
