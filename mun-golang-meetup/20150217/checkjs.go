// contextdetect - checkjs.go
import (
	"fmt"
	"regexp"
)

// checkJS verifies if the passed Javascript snippet is parsable
// returns nil if JS is OK, err from the otto parser otherwise
func checkJS(js string, verbose bool) (err error) {
	// check for <!-- and replace if present - imitate browser
	re := regexp.MustCompile(`(\n*\s*)<!--`)
	jsmod := js
	if res := re.MatchString(js); res == true {
		jsmod = re.ReplaceAllString(js, "$1//  ")
	}
	_, err = ottoparser.ParseFile(nil, "checkjs", jsmod, ottoparser.IgnoreRegExpErrors)
	if err != nil {
		fmt.Printf("JS ERROR: %s, Js Orig: %sJos Mod: %s\n", err, js, jsmod)
		return err
	}
	return nil
}
