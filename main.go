package main

import (
	"flag"
	"fmt"
	"github.com/Dedmustdie/sleeper/config"
	"github.com/Dedmustdie/sleeper/constants"
	"github.com/Dedmustdie/sleeper/utils"
	"strconv"
	"time"
)

func main() {
	cfg := config.Config{
		DefaultMode:           "3",
		GeneralDocumentation:  "Sleeper suspends the execution of the current thread until the time-out interval elapses",
		ModeDocumentation:     "Duration value type: 1-Microsecond, 2-Millisecond, 3-Second, 4-Minute, 5-Hour",
		HelpDocumentation:     "Show documentation",
		ArgumentDocumentation: "Sleeper duration (uint64)",
		ArgumentName:          "<argument>",
		ModeSignature:         "m",
		HelpSignature:         "help",
	}

	help := flag.Bool(cfg.HelpSignature, constants.DefaultHelp,
		cfg.HelpDocumentation)

	var modeCodeString string
	flag.StringVar(&modeCodeString, cfg.ModeSignature, cfg.DefaultMode,
		cfg.ModeDocumentation)
	flag.Parse()

	if *help {
		fmt.Print(utils.GetDocumentation(cfg.GeneralDocumentation,
			[]string{
				utils.FormatFlagDocumentation(cfg.HelpSignature, cfg.HelpDocumentation),
				utils.FormatFlagDocumentation(cfg.ModeSignature, cfg.ModeDocumentation),
				utils.FormatArgDocumentation(cfg.ArgumentName, cfg.ArgumentDocumentation),
			}))
		return
	}

	modeCode, err := strconv.ParseUint(modeCodeString, constants.UintBase, constants.UintBitSize)
	if err != nil || modeCode > constants.ModeMaxValue || modeCode < constants.ModeMinValue {
		fmt.Println("Wrong mode flag argument")
		return
	}

	durationValue, err := strconv.ParseUint(flag.Arg(0), constants.UintBase, constants.UintBitSize)
	if err != nil {
		fmt.Println("Wrong duration argument")
		return
	}

	duration, err := utils.CalculateDuration(durationValue, modeCode)

	time.Sleep(duration)
}
