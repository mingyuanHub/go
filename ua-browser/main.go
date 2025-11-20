package main

import (
	"fmt"
	"regexp"
	"strings"
)

// BrowserDetector 浏览器检测器
type BrowserDetector struct {
	patterns map[string]*regexp.Regexp
}

// NewBrowserDetector 创建浏览器检测器
func NewBrowserDetector() *BrowserDetector {
	return &BrowserDetector{
		patterns: map[string]*regexp.Regexp{
			// 优先级从高到低排列
			"edge":          regexp.MustCompile(`(?i)(edge|edg|edga|edgios)/(\d+)`),
			"ie":            regexp.MustCompile(`(?i)(msie|trident)(\/|\s)([\d\.]+)`),
			"opera":         regexp.MustCompile(`(?i)(opera|opr)/(\d+)`),
			"chrome":        regexp.MustCompile(`(?i)(chrome|crios)/(\d+)`),
			"firefox":       regexp.MustCompile(`(?i)(firefox|fxios)/(\d+)`),
			"safari":        regexp.MustCompile(`(?i)(safari|version)/(\d+)`),
			"android-browser": regexp.MustCompile(`(?i)android.*version/(\d+).*safari`),
			"webview":       regexp.MustCompile(`(?i)(webview|wv)`),
			"netscape":      regexp.MustCompile(`(?i)netscape`),
			"mozilla":       regexp.MustCompile(`(?i)mozilla`),
		},
	}
}

// DetectBrowser 检测浏览器类型
func (bd *BrowserDetector) DetectBrowser(ua string) string {
	if ua == "" {
		return "unknown"
	}

	// 按优先级检测
	for browser, pattern := range bd.patterns {
		if pattern.MatchString(ua) {
			// 特殊处理：避免将Edge识别为Chrome
			if browser == "chrome" {
				if strings.Contains(strings.ToLower(ua), "edg") {
					continue // 跳过，让edge模式匹配
				}
			}
			// 特殊处理：避免将Chrome识别为Safari
			if browser == "safari" {
				if strings.Contains(strings.ToLower(ua), "chrome") {
					continue // 跳过，让chrome模式匹配
				}
			}
			return browser
		}
	}

	return "unknown"
}

// DetectBrowserWithVersion 检测浏览器类型和版本
func (bd *BrowserDetector) DetectBrowserWithVersion(ua string) (string, string) {
	browser := bd.DetectBrowser(ua)
	version := bd.ExtractVersion(ua, browser)
	return browser, version
}

// ExtractVersion 提取浏览器版本
func (bd *BrowserDetector) ExtractVersion(ua, browser string) string {
	if pattern, exists := bd.patterns[browser]; exists {
		matches := pattern.FindStringSubmatch(ua)
		if len(matches) > 2 {
			return matches[len(matches)-1] // 返回版本号
		}
	}
	return "unknown"
}

// GetBrowserName 获取浏览器显示名称
func (bd *BrowserDetector) GetBrowserName(browserKey string) string {
	browserNames := map[string]string{
		"chrome":        "Chrome",
		"safari":        "Safari",
		"edge":          "Microsoft Edge",
		"ie":            "Internet Explorer",
		"firefox":       "Firefox",
		"opera":         "Opera",
		"mozilla":       "Mozilla",
		"netscape":      "Netscape Navigator",
		"android-browser": "Android Browser",
		"webview":       "In-app Browser",
		"unknown":       "Unknown Browser",
	}

	if name, exists := browserNames[browserKey]; exists {
		return name
	}
	return browserNames["unknown"]
}


func main() {
	detector := NewBrowserDetector()

	// 示例UA字符串
	testUAs := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0",
		"Mozilla/5.0 (Linux; Android 10; SM-G975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.120 Mobile Safari/537.36",
	}

	for _, ua := range testUAs {
		browser, version := detector.DetectBrowserWithVersion(ua)
		name := detector.GetBrowserName(browser)
		fmt.Printf("UA: %s\n", ua)
		fmt.Printf("检测结果: %s (版本: %s)\n\n", name, version)
	}
}