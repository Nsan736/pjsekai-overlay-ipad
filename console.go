package main

import (
	"fmt"
	"strings"

	"github.com/lithammer/dedent"
	"github.com/sevenc-nanashi/pjsekai-overlay/pkg/pjsekaioverlay"
)

func Title() {
	fmt.Printf(
		strings.TrimSpace(dedent.Dedent(`
      %s== pjsekai-overlay-ipad -------------------------------------------------------%s
        %spjsekai-overlay-ipad / プロセカ風動画作成補助ツール%s
        Version: %s%s%s
        編集 %sNsan%s
        https://github.com/Nsan736/pjsekai-overlay-ipad/
      %s-------------------------------------------------------------------------------%s
    `))+"\n\n",
		RgbColorEscape(0x00afc7), ResetEscape(),
		RgbColorEscape(0x00afc7), ResetEscape(),
		RgbColorEscape(0x0f6ea3), pjsekaioverlay.Version, ResetEscape(),
		RgbColorEscape(0x48b0d5), ResetEscape(),
		RgbColorEscape(0xff5a91), ResetEscape(),
	)

}

func RgbColorEscape(rgb int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", (rgb>>16)&0xff, (rgb>>8)&0xff, rgb&0xff)
}

func AnsiColorEscape(color int) string {
	return fmt.Sprintf("\033[38;5;%dm", color)
}

func ResetEscape() string {
	return "\033[0m"
}
