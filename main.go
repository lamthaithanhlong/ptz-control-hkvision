package main

import (
	"io/ioutil"
	"log"

	"github.com/use-go/onvif"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/media"
	"github.com/use-go/onvif/ptz"
	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

func run() {
	dev, err := onvif.NewDevice("192.168.100.64")
	if err != nil {
		panic(err)
	}

	dev.Authenticate("admin", "ftidev@123")

	res, err := dev.CallMethod(device.GetCapabilities{Category: "PTZ"})
	bs, _ := ioutil.ReadAll(res.Body)
	log.Printf("output %+v %s", res.StatusCode, bs)
	endpoint := dev.GetEndpoint("ptz")
	log.Print(endpoint)

	res, err = dev.CallMethod(media.GetProfiles{})
	log.Print(err)
	bs, _ = ioutil.ReadAll(res.Body)
	log.Printf("Profiles %+v %s", res.StatusCode, bs)

	res, err = dev.CallMethod(ptz.GetConfigurations{})
	log.Print(err)
	bs, _ = ioutil.ReadAll(res.Body)
	log.Printf("output %+v %s", res.StatusCode, bs)

	// res, err = dev.CallMethod(ptz.AbsoluteMove{
	// 	ProfileToken: "Profile_2",
	// 	Position: onvif2.PTZVector{
	// 		PanTilt: onvif2.Vector2D{
	// 			X:     -0.2,
	// 			Y:     0.012499999999999968,
	// 			Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/PositionGenericSpace",
	// 		},
	// 		Zoom: onvif2.Vector1D{
	// 			X:     0,
	// 			Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/PositionGenericSpace",
	// 		},
	// 	},
	// })
	res, err = dev.CallMethod(ptz.RelativeMove{
		ProfileToken: "Profile_2",
		Translation: onvif2.PTZVector{
			PanTilt: onvif2.Vector2D{
				X:     -0.8124999999999999,  //-0.655611,
				Y:     0.049999999999999864, //1,
				Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/TranslationGenericSpace",
			},
			Zoom: onvif2.Vector1D{
				X:     0,
				Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/TranslationGenericSpace",
			},
		},
	})
	// log.Print(err)
	// bs, _ = ioutil.ReadAll(res.Body)
	// log.Printf("output %+v %s", res.StatusCode, bs)

	// res, err = dev.CallMethod(ptz.GetStatus{
	// 	ProfileToken: "Profile_2",
	// })
	// log.Print(err)
	// bs, _ = ioutil.ReadAll(res.Body)
	// log.Printf("output %+v %s", res.StatusCode, bs)
}
func main() {
	run()
}
