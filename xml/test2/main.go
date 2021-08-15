package main

import (
	"encoding/xml"
	"fmt"
)

type VAST struct {
	Ad *Ad2  `xml:"Ad"`
	Version  string  `xml:"version,attr"`
}

type Ad2 struct {
	AdSystem string  `xml:"AdSystem"`
}


func main()  {
	//vast := &VAST{
	//	Ad : &Ad2{
	//		AdSystem: "tradplus",
	//	},
	//	Version: "2.0",
	//}

	//txtxml := "<VAST version=\"2.0\">\n  <Ad>\n   <AdSystem>tradplus</AdSystem>\n  </Ad>\n </VAST>"
	//vast1 := &VAST{}
	//err := xml.Unmarshal([]byte(txtxml), &vast1)
	//
	//if err != nil {
	//	fmt.Printf("unmarshal err : %v\n", err)
	//	return
	//}
	//
	//fmt.Println(vast1)
	//
	//resXML, err := xml.MarshalIndent(vast1, " ", " ")
	//if err != nil {
	//	fmt.Printf("marshal xml err :%v\n", err)
	//	return
	//}
	//
	//fmt.Println(xml.Header)
	//fmt.Println(string(resXML))


	vast3()

	//FCreatives()
}

func vast3()  {
	//txt:= "<VAST version=\"3.0\">\n   <Error>\n                <![CDATA[https://aktrack.pubmatic.com/er=[ERRORCODE]]]>\n            </Error>\n    <Ad id=\"123\">\n        <InLine>\n            <AdSystem>PubMatic</AdSystem>\n            <AdTitle>VAST 2.0 Instream Test</AdTitle>\n            <Description>VAST 2.0 Instream Test</Description>\n            <Error>\n                <![CDATA[https://aktrack.pubmatic.com/er=[ERRORCODE]]]>\n            </Error>\n            <Impression>\n                <![CDATA[https://aktrack.pubmatic.com?e=impression]]>\n            </Impression>\n            <Creatives>\n                <Creative AdID=\"123\">\n                    <Linear>\n                        <Duration>00:00:30</Duration>\n                        <TrackingEvents>\n                            <Tracking event=\"creativeView\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=creativeView]]>\n                            </Tracking>\n                            <Tracking event=\"start\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=start]]>\n                            </Tracking>\n                            <Tracking event=\"midpoint\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=midpoint]]>\n                            </Tracking>\n                            <Tracking event=\"firstQuartile\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=firstQuartile]]>\n                            </Tracking>\n                            <Tracking event=\"thirdQuartile\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=thirdQuartile]]>\n                            </Tracking>\n                            <Tracking event=\"complete\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=complete]]>\n                            </Tracking>\n                        </TrackingEvents>\n                        <VideoClicks>\n                            <ClickThrough>\n                                <![CDATA[https://www.pubmatic.com]]>\n                            </ClickThrough>\n                        </VideoClicks>\n                        <MediaFiles>\n                            <MediaFile bitrate=\"500\" delivery=\"progressive\" height=\"460\" maintainAspectRatio=\"true\" scalable=\"true\" type=\"video/mp4\" width=\"480\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/spinning-logo-480x360_video.mp4]]>\n                            </MediaFile>\n                            <MediaFile bitrate=\"500\" delivery=\"progressive\" height=\"460\" maintainAspectRatio=\"true\" scalable=\"true\" type=\"video/ogg\" width=\"480\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/spinning-logo-480x360_video.ogg]]>\n                            </MediaFile>\n                            <MediaFile bitrate=\"500\" delivery=\"progressive\" height=\"300\" maintainAspectRatio=\"true\" scalable=\"true\" type=\"video/x-flv\" width=\"400\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/PubMatic_test_video.flv]]>\n                            </MediaFile>\n                        </MediaFiles>\n                    </Linear>\n                </Creative>\n                <Creative AdID=\"123\">\n                    <NonLinearAds>\n                        <TrackingEvents/>\n                        <NonLinear height=\"50\" minSuggestedDuration=\"00:00:05\" width=\"300\">\n                            <StaticResource creativeType=\"image/jpeg\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/PubMatic_LetsBeClear_300x50.jpeg]]>\n                            </StaticResource>\n                            <NonLinearClickThrough>\n                                <![CDATA[https://www.pubmatic.com]]>\n                            </NonLinearClickThrough>\n                        </NonLinear>\n                    </NonLinearAds>\n                </Creative>\n                <Creative AdID=\"123\">\n                    <CompanionAds>\n                        <Companion height=\"250\" width=\"300\">\n                            <StaticResource creativeType=\"image/jpeg\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/PubMatic_LetsBeClear_320x250.jpg]]>\n                            </StaticResource>\n                            <CompanionClickThrough>\n                                <![CDATA[https://www.pubmatic.com]]>\n                            </CompanionClickThrough>\n                        </Companion>\n                    </CompanionAds>\n                </Creative>\n            </Creatives>\n        </InLine>\n    </Ad>\n</VAST>"
	txt := "<VAST version=\"2.0\">\n    <Ad id=\"{{InMobi-Creative-Id-numeric}}\">\n        <InLine>\n            <AdSystem>InMobi</AdSystem>\n            <AdTitle>InMobi Ad</AdTitle>//\n            <Impression>\n                <![CDATA[{{RM Custom Beacon - video imp}}]]></Impression>\n            <Error>\n                <![CDATA[{{RM Error Beacon}}]]>\n            </Error>\n            <Creatives>\n                <Creative id=\"1\">\n                    <Linear>\n                        <Duration>{{00:00:15}}</Duration>\n                        <TrackingEvents>\n                            <Tracking event=\"impression\">\n                                <![CDATA[{{RMVideo Impression Beacon}}]]>\n                            </Tracking>\n                            <Tracking event=\"start\">\n                                <![CDATA[{{RM VideoStart Beacon}}]]>\n                            </Tracking>\n                            <Tracking event=\"midpoint\">\n                                <![CDATA[{{RMVideo 50% Beacon}}]]>\n                            </Tracking>\n                            <Tracking event=\"firstQuartile\">\n                                <![CDATA[{{RMVideo 25% Beacon}}]]>\n                            </Tracking>\n                            <Tracking event=\"thirdQuartile\">\n                                <![CDATA[{{RMVideo 75% Beacon}}]]>\n                            </Tracking>\n                            <Tracking event=\"complete\">\n                                <![CDATA[{{RMVideo Complete Beacon}}]]>\n                            </Tracking>\n                        </TrackingEvents>\n                        <VideoClicks>\n                            <ClickThrough>\n                                <![CDATA[{{Click Beacon}}]]>\n                            </ClickThrough>\n                        </VideoClicks>\n                        <MediaFiles>\n                            <MediaFile bitrate=\"{{TO FILL BITRATE}}\" delivery=\"progressive\" height=\"{{TO FILL HEIGHT}}\" maintainAspectRatio=\"true\" scalable=\"false\" type=\"video/mp4\" width=\"{{TO FILLWIDTH}}\">\n                                <![CDATA[{{TOFILL VIDEOURL}}]]>\n                            </MediaFile>\n                            <MediaFile bitrate=\"{{TO FILL}}\" delivery=\"progressive\" height=\"{{TO FILL}}\" maintainAspectRatio=\"true\" scalable=\"false\" type=\"video/mp4\" width=\"{{TO FILL}}\">\n                                <![CDATA[{{TO FILL}}]]>\n                            </MediaFile>\n                            <MediaFile bitrate=\"{{TO FILL}}\" delivery=\"progressive\" height=\"{{TO FILL}}\" maintainAspectRatio=\"true\" scalable=\"false\" type=\"video/mp4\" width=\"{{TO FILL}}\">\n                                <![CDATA[{{TO FILL}}]]>\n                            </MediaFile>\n                        </MediaFiles>\n                    </Linear>\n                </Creative>\n            </Creatives>\n        </InLine>\n    </Ad>\n</VAST>"
	vast := &VAST3{}
	err := xml.Unmarshal([]byte(txt), &vast)

	if err != nil {
		fmt.Printf("unmarshal err : %v\n", err)
		return
	}

	fmt.Println(vast.Error)

	resXML, err := xml.MarshalIndent(vast, " ", " ")
	if err != nil {
		fmt.Printf("marshal xml err :%v\n", err)
		return
	}

	fmt.Println(xml.Header)
	fmt.Println(string(resXML))
}

func FCreatives()  {
	txt:= "<Creatives>\n                <Creative AdID=\"123\">\n                    <Linear>\n                        <Duration>00:00:30</Duration>\n                        <TrackingEvents>\n                            <Tracking event=\"creativeView\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=creativeView]]>\n                            </Tracking>\n                            <Tracking event=\"start\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=start]]>\n                            </Tracking>\n                            <Tracking event=\"midpoint\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=midpoint]]>\n                            </Tracking>\n                            <Tracking event=\"firstQuartile\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=firstQuartile]]>\n                            </Tracking>\n                            <Tracking event=\"thirdQuartile\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=thirdQuartile]]>\n                            </Tracking>\n                            <Tracking event=\"complete\">\n                                <![CDATA[https://aktrack.pubmatic.com?e=complete]]>\n                            </Tracking>\n                        </TrackingEvents>\n                        <VideoClicks>\n                            <ClickThrough>\n                                <![CDATA[https://www.pubmatic.com]]>\n                            </ClickThrough>\n                        </VideoClicks>\n                        <MediaFiles>\n                            <MediaFile bitrate=\"500\" delivery=\"progressive\" height=\"460\" maintainAspectRatio=\"true\" scalable=\"true\" type=\"video/mp4\" width=\"480\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/spinning-logo-480x360_video.mp4]]>\n                            </MediaFile>\n                            <MediaFile bitrate=\"500\" delivery=\"progressive\" height=\"460\" maintainAspectRatio=\"true\" scalable=\"true\" type=\"video/ogg\" width=\"480\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/spinning-logo-480x360_video.ogg]]>\n                            </MediaFile>\n                            <MediaFile bitrate=\"500\" delivery=\"progressive\" height=\"300\" maintainAspectRatio=\"true\" scalable=\"true\" type=\"video/x-flv\" width=\"400\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/PubMatic_test_video.flv]]>\n                            </MediaFile>\n                        </MediaFiles>\n                    </Linear>\n                </Creative>\n                <Creative AdID=\"123\">\n                    <NonLinearAds>\n                        <TrackingEvents/>\n                        <NonLinear height=\"50\" minSuggestedDuration=\"00:00:05\" width=\"300\">\n                            <StaticResource creativeType=\"image/jpeg\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/PubMatic_LetsBeClear_300x50.jpeg]]>\n                            </StaticResource>\n                            <NonLinearClickThrough>\n                                <![CDATA[https://www.pubmatic.com]]>\n                            </NonLinearClickThrough>\n                        </NonLinear>\n                    </NonLinearAds>\n                </Creative>\n                <Creative AdID=\"123\">\n                    <CompanionAds>\n                        <Companion height=\"250\" width=\"300\">\n                            <StaticResource creativeType=\"image/jpeg\">\n                                <![CDATA[https://staging.pubmatic.com:8443/test/PubMatic_LetsBeClear_320x250.jpg]]>\n                            </StaticResource>\n                            <CompanionClickThrough>\n                                <![CDATA[https://www.pubmatic.com]]>\n                            </CompanionClickThrough>\n                        </Companion>\n                    </CompanionAds>\n                </Creative>\n            </Creatives>"

	vast := &Creatives{}
	err := xml.Unmarshal([]byte(txt), &vast)

	if err != nil {
		fmt.Printf("unmarshal err : %v\n", err)
		return
	}

	fmt.Println(vast)

	resXML, err := xml.MarshalIndent(vast, " ", " ")
	if err != nil {
		fmt.Printf("marshal xml err :%v\n", err)
		return
	}

	fmt.Println(xml.Header)
	fmt.Println(string(resXML))
}
