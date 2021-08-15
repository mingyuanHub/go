package main

type VAST3 struct {
	Error   *CDATA `xml:",omitempty"`
	Ad      []*Ad
	Version string `xml:"version,attr"`
}

type Ad struct {
	InLine   *InLine
	Id       string `xml:"id,attr,omitempty"`
	Sequence string `xml:"sequence,attr,omitempty"`
}

type InLine struct {
	AdSystem    *AdSystem
	AdTitle     string
	Description string   `xml:",omitempty"`
	Advertiser  string   `xml:",omitempty"`
	Pricing     *Pricing `xml:",omitempty"`
	Survey      *CDATA   `xml:",omitempty"`
	Error       *CDATA   `xml:",omitempty"`
	Impression  *Impression
	Creatives   *Creatives
	Extensions  *Extensions `xml:",omitempty"`
}

type AdSystem struct {
	Value   string `xml:",chardata"`
	Version string `xml:"version,attr,omitempty"`
}

type Pricing struct {
	Value    string `xml:",chardata"`
	Model    string `xml:"model,attr,omitempty"`
	Currency string `xml:"currency,attr,omitempty"`
}

type Impression struct {
	CDATA
	Id string `xml:"id,attr,omitempty"`
}

type Creatives struct {
	Creative []*Creative
}

type Creative struct {
	CreativeExtensions *CreativeExtensions `xml:",omitempty"`
	Linear             *Linear
	CompanionAds       *CompanionAds
	NonLinearAds       *NonLinearAds
	Id                 string `xml:"id,attr,omitempty"`
	Sequence           string `xml:"sequence,attr,omitempty"`
	AdID               string `xml:"adID,attr,omitempty"`
	ApiFramework       string `xml:"apiFramework,attr,omitempty"`
}

type CreativeExtensions struct {
	CreativeExtension *CreativeExtension `xml:",omitempty"`
}

type CreativeExtension struct {
	Value string `xml:",chardata"`
}

type Linear struct {
	AdParameters   *AdParameters `xml:",omitempty"`
	Duration       string
	MediaFiles     *MediaFiles
	TrackingEvents *TrackingEvents `xml:",omitempty"`
	VideoClicks    *VideoClicks    `xml:",omitempty"`
	Icons          *Icons          `xml:",omitempty"`
	SkipOffset     string          `xml:"skipoffset,attr,omitempty"`
}

type AdParameters struct {
	Value      string `xml:",chardata"`
	XmlEncoded string `xml:"xmlEncoded,attr,omitempty"`
}

type MediaFiles struct {
	MediaFile []*MediaFile
}

type MediaFile struct {
	CDATA
	Id                 string `xml:"id,attr,omitempty"`
	Delivery           string `xml:"delivery,attr,omitempty"`
	Type               string `xml:"type,attr,omitempty"`
	Bitrate            string `xml:"bitrate,attr,omitempty"`
	MinBitrate         string `xml:"minBitrate,attr,omitempty"`
	MaxBitrate         string `xml:"maxBitrate,attr,omitempty"`
	Width              string `xml:"width,attr,omitempty"`
	Height             string `xml:"height,attr,omitempty"`
	Scalable           string `xml:"scalable,attr,omitempty"`
	MantainAspectRatio string `xml:"mantainAspectRatio,attr,omitempty"`
	Codec              string `xml:"codec,attr,omitempty"`
	ApiFramework       string `xml:"apiFramework,attr,omitempty"`
}

type TrackingEvents struct {
	Tracking []*Tracking
}

type Tracking struct {
	CDATA
	Event string `xml:"event,attr,omitempty"`
}

type VideoClicks struct {
	ClickThrough  *ClickId `xml:",omitempty"`
	ClickTracking *ClickId `xml:",omitempty"`
	CustomClick   *ClickId `xml:",omitempty"`
}

type ClickId struct {
	CDATA
	Id string `xml:"id,attr,omitempty"`
}

type Icons struct {
	Icon []*Icon
}

type Icon struct {
	StaticResource   *StaticResource `xml:",omitempty"`
	IFrameResource   string          `xml:",omitempty"`
	HTMLResource     string          `xml:",omitempty"`
	IconClicks       *IconClicks     `xml:",omitempty"`
	IconViewTracking *CDATA          `xml:",omitempty"`
	Program          string          `xml:"program,attr,omitempty"`
	Width            string          `xml:"width,attr,omitempty"`
	Height           string          `xml:"height,attr,omitempty"`
	XPosition        string          `xml:"xPosition,attr,omitempty"`
	YPosition        string          `xml:"yPosition,attr,omitempty"`
	Duration         string          `xml:"duration,attr,omitempty"`
	Offset           string          `xml:"offset,attr,omitempty"`
	ApiFramework     string          `xml:"apiFramework,attr,omitempty"`
}

type IconClicks struct {
	IconClickThrough  *CDATA   `xml:",omitempty"`
	IconClickTracking *ClickId `xml:",omitempty"`
}

type CompanionAds struct {
	Companion []*Companion
}

type Companion struct {
	StaticResource         *StaticResource `xml:",omitempty"`
	IFrameResource         string          `xml:",omitempty"`
	HTMLResource           string          `xml:",omitempty"`
	AdParameters           *AdParameters   `xml:",omitempty"`
	AltText                string          `xml:",omitempty"`
	CompanionClickThrough  *CDATA          `xml:",omitempty"`
	CompanionClickTracking *ClickId        `xml:",omitempty"`
	TrackingEvents         *TrackingEvents `xml:",omitempty"`
	Id                     string          `xml:"id,attr,omitempty"`
	Width                  string          `xml:"width,attr,omitempty"`
	Height                 string          `xml:"height,attr,omitempty"`
	AssetWidth             string          `xml:"assetWidth,attr,omitempty"`
	AssetHeight            string          `xml:"assetHeight,attr,omitempty"`
	ExpandedWidth          string          `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight         string          `xml:"expandedHeight,attr,omitempty"`
	ApiFramework           string          `xml:"apiFramework,attr,omitempty"`
	AdSlotID               string          `xml:"adSlotID,attr,omitempty"`
}

type StaticResource struct {
	CDATA
	CreativeType string `xml:"creativeType,attr,omitempty"`
}

type NonLinearAds struct {
	NonLinear *NonLinear
}

type NonLinear struct {
	StaticResource         *StaticResource `xml:",omitempty"`
	IFrameResource         CDATA           `xml:",omitempty"`
	HTMLResource           CDATA           `xml:",omitempty"`
	NonLinearClickThrough  CDATA           `xml:",omitempty"`
	NonLinearClickTracking *ClickId        `xml:",omitempty"`
	AdParameters           string          `xml:",omitempty"`
	TrackingEvents         *TrackingEvents `xml:",omitempty"`
	Id                     string          `xml:"id,attr,omitempty"`
	Width                  string          `xml:"width,attr,omitempty"`
	Height                 string          `xml:"height,attr,omitempty"`
	ExpandedWidth          string          `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight         string          `xml:"expandedHeight,attr,omitempty"`
	Scalable               string          `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio    string          `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration   string          `xml:"minSuggestedDuration,attr,omitempty"`
	ApiFrameworkstring     string          `xml:"apiFramework,attr,omitempty"`
}

type Extensions struct {
	Extension *Extension
}

type Extension struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr,omitempty"`
}

type CDATA struct {
	Text string `xml:",cdata"`
}