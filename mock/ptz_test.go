package onvif

import (
	"ptz/onvif"
	"testing"

	"gotest.tools/assert"
)

func case1(t *testing.T) {
	temp := onvif.ControllPTZ(-0.8124999999999999, 0.049999999999999864, 0, "192.168.100.64", "admin", "ftidev@123")
	assert.Equal(t, "tramit", temp, "Expect tramit")
}

func case2(t *testing.T) {
	temp := onvif.ControllPTZ(0, 0, 0, "192.168.100.64", "admin", "ftidev@123")
	assert.Equal(t, "tramit", temp, "Expect tramit")
}
