package onvif

import (
	"fmt"

	"github.com/use-go/onvif"
	"github.com/use-go/onvif/ptz"
	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

// Checking PTZ credential
// Private function
func connectPTZ(ip, username, pass string) *onvif.Device {
	dev, err := onvif.NewDevice(ip)
	if err != nil {
		panic(err)
	}
	if username != "" || pass != "" {
		fmt.Println(username, pass)
		dev.Authenticate(username, pass)
		return dev
	}
	return nil
}

// Controll PTZ
func ControllPTZ(x float64, y float64, z float64, ip, username, pass string) string {
	fmt.Println(x, y, z, ip, username, pass)
	connectPTZ(ip, username, pass).CallMethod(ptz.RelativeMove{
		ProfileToken: "Profile_2",
		Translation: onvif2.PTZVector{
			PanTilt: onvif2.Vector2D{
				X:     x,
				Y:     y,
				Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/TranslationGenericSpace",
			},
			Zoom: onvif2.Vector1D{
				X:     z,
				Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/TranslationGenericSpace",
			},
		},
	})
	return "trammited"
}
