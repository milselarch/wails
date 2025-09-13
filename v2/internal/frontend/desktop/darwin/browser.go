//go:build darwin
// +build darwin

package darwin

import (
	"fmt"
	"github.com/milselarch/wails/v2/internal/frontend/utils"
	"github.com/pkg/browser"
)

// BrowserOpenURL Use the default browser to open the url
func (f *Frontend) BrowserOpenURL(rawURL string) {
	url, err := utils.ValidateAndSanitizeURL(rawURL)
	if err != nil {
		f.logger.Error(fmt.Sprintf("Invalid URL %s", err.Error()))
		return
	}

	// Specific method implementation
	if err := browser.OpenURL(url); err != nil {
		f.logger.Error("Unable to open default system browser")
	}
}
