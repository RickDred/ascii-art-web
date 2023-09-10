package art

import (
	"github.com/RickDred/ascii-art-web/pkg/files"
	"log"
	"strings"
)

func Ascii(text, pathToBanner string) string {
	ans := ""
	text = strings.ReplaceAll(text, "\\n", "\n")
	lines := strings.Split(text, "\n")

	temp := ""
	for _, line := range lines {
		if len(line) != 0 {
			break
		}
		temp += "\n"
	}
	if len(temp) == len(lines) {
		return temp[:len(temp)-1]
	}

	bannerData, err := files.Contents(pathToBanner)
	if err != nil {
		log.Fatal(err)
	}

	const charHeight = 9
	font := strings.Split(bannerData, "\n")
	for _, line := range lines {
		if line == "" {
			ans += "\n"
			continue
		}

		for i := 0; i < 8; i++ {
			for _, ch := range line {
				line := (int(ch)-int(' '))*charHeight + i + 1
				ans += font[line]
			}
			ans += "\n"
		}
	}
	return ans
}
