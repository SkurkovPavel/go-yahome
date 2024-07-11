package iot

import "fmt"

func (d *Device) Validate() error {

	if d.Id == "" {
		return fmt.Errorf("deviceID not provided")
	}

	if len(d.Actions) == 0 {
		return fmt.Errorf("device actions not provided")
	}

	if d.Type == "" {
		return fmt.Errorf("device type not provided")
	}
	if d.State == "" {
		return fmt.Errorf("device state not provided")
	}
	return nil
}
func (g *Group) Validate() error {
	if g.Id == "" {
		return fmt.Errorf("group not provided")
	}

	if len(g.Actions) == 0 {
		return fmt.Errorf("group actions not provided")
	}

	return nil
}

func ErrorReturn(method string, err error) error {
	return fmt.Errorf("%s error: %s", method, err.Error())
}
