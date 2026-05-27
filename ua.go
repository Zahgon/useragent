package useragent

import (
	"regexp"
)

// UserAgent struct containing all data extracted from parsed user-agent string
type UserAgent struct {
	VersionNo   VersionNo
	OSVersionNo VersionNo
	URL         string
	String      string
	Name        string
	Version     string
	OS          string
	OSVersion   string
	Device      string
	Mobile      bool
	Tablet      bool
	Desktop     bool
	Bot         bool
}

// Constants for browsers and operating systems for easier comparison
const (
	Windows        = "Windows"
	WindowsPhone   = "Windows Phone"
	WindowsNT      = "Windows NT"
	WindowsPhoneOS = "Windows Phone OS"
	Android        = "Android"
	MacOS          = "macOS"
	IOS            = "iOS"
	Linux          = "Linux"
	FreeBSD        = "FreeBSD"
	ChromeOS       = "ChromeOS"
	BlackBerry     = "BlackBerry"
	CrOS           = "CrOS"
	Harmony        = "Harmony"

	Opera            = "Opera"
	OperaMini        = "Opera Mini"
	OperaTouch       = "Opera Touch"
	Chrome           = "Chrome"
	HeadlessChrome   = "Headless Chrome"
	Firefox          = "Firefox"
	InternetExplorer = "Internet Explorer"
	Safari           = "Safari"
	Edge             = "Edge"
	Vivaldi          = "Vivaldi"
	MobileSafari     = "Mobile Safari"
	NetFront         = "NetFront"
	Mozilla          = "Mozilla"
	Msie             = "MSIE"
	SamsungBrowser   = "Samsung Browser"

	GoogleAdsBot        = "Google Ads Bot"
	Googlebot           = "Googlebot"
	Twitterbot          = "Twitterbot"
	FacebookExternalHit = "facebookexternalhit"
	Applebot            = "Applebot"
	Bingbot             = "Bingbot"
	YandexBot           = "YandexBot"
	YandexAdNet         = "YandexAdNet"

	FacebookApp  = "Facebook App"
	InstagramApp = "Instagram App"
	TiktokApp    = "TikTok App"

	Version = "Version"
	Mobile  = "Mobile"
	Tablet  = "Tablet"

	tablet = "tablet"
)

// Parse user agent string returning UserAgent struct
func Parse(userAgent string) UserAgent { _ = "STUB: not implemented"; return *new(UserAgent) }

// OS lookup

// Opera on iOS

// Chrome on iOS

// Firefox on iOS

// if Chrome and Safari defined, find any other token sent descr

// If mobile flag has already been set, don't override it.

// if tablet, switch mobile to off

// if not already bot, check some popular bots and whether URL is set

// var buffPool = sync.Pool{New: func() interface{} {
// 	return bytes.NewBuffer(make([]byte, 0, 30))
// }}

func parse(userAgent []byte) properties { _ = "STUB: not implemented"; return *new(properties) }

// buff := buffPool.Get().(*bytes.Buffer)
// val := buffPool.Get().(*bytes.Buffer)
// buff.Reset()
// val.Reset()

// if value don't exists, try to get version from the token

// )

// ;

// ;

// (

// [

// ]

// :

// If we are part of a URL just write the character.

// If the following character is not a space, change to a space.

// Otherwise don't write as it's probably a badly formatted key value separator.

//   /

// buffPool.Put(buff)
// buffPool.Put(val)

func checkVer(s string) property { _ = "STUB: not implemented"; return *new(property) }

// ignore returns true if token should be ignored
func ignore(s string) bool { _ = "STUB: not implemented"; return false }

type property struct {
	Key   string
	Value string
}
type properties struct {
	list []property
	url  string
}

func (p properties) get(key string) string { _ = "STUB: not implemented"; return "" }

func (p properties) getIndexValue(key string) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (p properties) exists(key string) bool { _ = "STUB: not implemented"; return false }

// func (p properties) existsIgnoreCase(key string) bool {
// 	for _, prop := range p.list {
// 		if strings.EqualFold(prop.Key, key) {
// 			return true
// 		}
// 	}
// 	return false
// }

func (p properties) existsAny(keys ...string) bool { _ = "STUB: not implemented"; return false }

func (p properties) getAny(keys ...string) (key, value string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (p properties) findMacOSVersion() string { _ = "STUB: not implemented"; return "" }

func (p properties) startsWith(value string) bool { _ = "STUB: not implemented"; return false }

func (p properties) findInstagramVersion() string { _ = "STUB: not implemented"; return "" }

// findBestMatch from the rest of the bunch
// in first cycle only return key with version value
// if withVerValue is false, do another cycle and return any token
func (p properties) findBestMatch(withVerOnly bool) string { _ = "STUB: not implemented"; return "" }

// don't pick if starts with number

// in first check, only return keys with value

var rxMacOSVer = regexp.MustCompile(`[_\d\.]+`)

func findVersion(s string) string { _ = "STUB: not implemented"; return "" }

// findAndroidDevice in tokens
func (p *properties) findAndroidDevice(startIndex int) string { _ = "STUB: not implemented"; return "" }

// probably language tag (en-us etc..), ignore and continue loop

// ignore these tokens, not device names

// leave Tablet tag for later table detection
