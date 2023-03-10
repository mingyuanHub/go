package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main()	{
	cp := "AdGroup+Creative+ID+0=449536177436&AdGroup+ID+0=116979598542&App+ID=476943146&Backend+Query+ID=CMbGhqezuv0CFTZBwgUdWHkBMw&Creative+ID+0=507897871636&Customer+ID+0=86317845&Landing+Page+0=https://apps.apple.com/app/id1037457231%%3Fmt%%3D8%%26gclid%%3DEAIaIQobChMIxsaGp7O6_QIVNkHCBR1YeQEzEAEYASAAEgIEIfD_BwE&URL+Clickstring+0=C0UPDlRj_Y4aoKLaCid4P2PKFmAMAnOrG04oNwI23ARABIMLs3CBgnbnQgZAFiAEBoAH71NLXA6kCL2B2OZOUpj6oAwGqBL8BT9CvXT8V54ULu5MBu1OWt7HO-VonmoXCMDA2sButN4HDeJTPamnnned3pRk5cJoesfkUNrt8vxruwK0tKkJrTwzrBRwzAe9gxehj8_OkLIN6CaO_CykYJYIIPwvierkGZ2_7ptkiEIzX0mlq3nYANhTL25OFMEu9ea3xDmqIm3AWTwUZjkJNbIqGCxYdE6oh4xn7em5nnPpLeWcYKIbrIvJ2dWw0Ji1HTGlyS7kry1uP2gZNeKR5gPbotRmC2NLABM6JnOSzA4gF_-7G8C6QBgGgBhrABguAB-2qrSiYBwGoB82bsQKoB5mdsQKoB6a-G6gH1ckbqAehAagH_p6xAqgH89EbqAeW2BuoB6qbsQKoB_-esQKoB9-fsQKoB47cG6gHyZyxAqgH5pqxArgH7rvFxfW_n8ZSwAfW2APAB-jYA9gHAfoHEmNvbS5nb29nbGUuQWRXb3Jkc6AI_IipBLAIArgIAdIIDwiAYRABGF8yAooCOgKAQKoNAkNOuBPhAoIUGBoWbW9iaWxlYXBwOjoxLTQ3Njk0MzE0NrAVAdAVAZgWAcoWOgoKMTAzNzQ1NzIzMRomCJS45qLzlavycxCDvPqFqeLbv-MBGAAgACoJNDc2OTQzMTQ2MAMg7fsDKAH4FgGAFwE"

	cp = strings.Replace(cp, "%%", "%", -1)

	s, e := url.ParseQuery(cp)

	fmt.Println(111, s)
	fmt.Println(222, e)

	for k, v := range s {
		fmt.Println(333, k, v)
	}
}
