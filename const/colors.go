package _const

import "github.com/fatih/color"

var Colors map[string]*color.Color

func init() {
	Colors = make(map[string]*color.Color)
	Colors["yellow"] = color.New(color.FgYellow)
	Colors["red"] = color.New(color.FgRed)
	Colors["blue"] = color.New(color.FgBlue)
	Colors["cyan"] = color.New(color.FgCyan)
	Colors["magenta"] = color.New(color.FgMagenta)
	Colors["white"] = color.New(color.FgWhite)
	Colors["black"] = color.New(color.FgBlack)
	Colors["red"] = color.New(color.FgRed)
	Colors["green"] = color.New(color.FgGreen)
	Colors["yellowBold"] = color.New(color.FgYellow).Add(color.Bold)
	Colors["redBold"] = color.New(color.FgRed).Add(color.Bold)
	Colors["blueBold"] = color.New(color.FgBlue).Add(color.Bold)
	Colors["cyanBold"] = color.New(color.FgCyan).Add(color.Bold)
	Colors["magentaBold"] = color.New(color.FgMagenta).Add(color.Bold)
	Colors["whiteBold"] = color.New(color.FgWhite).Add(color.Bold)
	Colors["blackBold"] = color.New(color.FgBlack).Add(color.Bold)
	Colors["redBold"] = color.New(color.FgRed).Add(color.Bold)
	Colors["greenBold"] = color.New(color.FgGreen).Add(color.Bold)

}
