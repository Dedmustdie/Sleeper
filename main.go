package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	cfg := Config{
		defaultMode:           "3",
		generalDocumentation:  "Sleeper suspends the execution of the current thread until the time-out interval elapses",
		modeDocumentation:     "Duration value type: 1-Microsecond, 2-Millisecond, 3-Second, 4-Minute, 5-Hour",
		helpDocumentation:     "Show documentation",
		argumentDocumentation: "Sleeper duration (uint64)",
		argumentName:          "<argument>",
		modeSignature:         "m",
		helpSignature:         "help",
	}

	help := flag.Bool(cfg.helpSignature, defaultHelp,
		cfg.helpDocumentation)

	var modeCodeString string
	flag.StringVar(&modeCodeString, cfg.modeSignature, cfg.defaultMode,
		cfg.modeDocumentation)
	flag.Parse()

	if *help {
		fmt.Print(getDocumentation(cfg.generalDocumentation,
			[]string{
				formatFlagDocumentation(cfg.helpSignature, cfg.helpDocumentation),
				formatFlagDocumentation(cfg.modeSignature, cfg.modeDocumentation),
				formatArgDocumentation(cfg.argumentName, cfg.argumentDocumentation),
			}))
		return
	}

	modeCode, err := strconv.ParseUint(modeCodeString, uintBase, uintBitSize)
	if err != nil || modeCode > modeMaxValue || modeCode < modeMinValue {
		fmt.Println("Wrong mode flag argument")
		return
	}

	durationValue, err := strconv.ParseUint(flag.Arg(0), uintBase, uintBitSize)
	if err != nil {
		fmt.Println("Wrong duration argument")
		return
	}

	duration, err := calculateDuration(durationValue, modeCode)

	time.Sleep(duration)
}
