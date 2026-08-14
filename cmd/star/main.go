// Command star previews the sidebar animation without building or deploying the
// SSH server. It renders the exact same ui.RenderStar the TUI uses.
//
//	go run ./cmd/star                        animate at the TUI's real size
//	go run ./cmd/star -rows 26 -cols 40      bigger, to see the shape clearly
//	go run ./cmd/star -still -frame 12       one frame, no animation
//	go run ./cmd/star -sharp 0.6 -spin 0.02  sweep the tunables
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/nourinawadd/ssh-portfolio/ui"
)

func main() {
	rows := flag.Int("rows", 20, "grid rows")
	cols := flag.Int("cols", 22, "grid columns (the TUI sidebar uses 22)")
	fps := flag.Int("fps", 12, "frames per second")
	frame := flag.Int("frame", 0, "starting frame number")
	frames := flag.Int("frames", 0, "stop after N frames (0 = until Ctrl-C)")
	still := flag.Bool("still", false, "print one frame and exit")
	grid := flag.Bool("grid", false, "draw a border showing the cell bounds")

	flag.Float64Var(&ui.StarSharp, "sharp", ui.StarSharp, "arm sharpness, below 1")
	flag.Float64Var(&ui.StarGamma, "gamma", ui.StarGamma, "arm brightness falloff")
	flag.Float64Var(&ui.StarCore, "core", ui.StarCore, "overall gain")
	flag.Float64Var(&ui.StarStretch, "stretch", ui.StarStretch, "vertical elongation")
	flag.Float64Var(&ui.StarSpin, "spin", ui.StarSpin, "rotation per frame")
	flag.Float64Var(&ui.StarPulse, "pulse", ui.StarPulse, "breathing amplitude")
	flag.Float64Var(&ui.StarGrain, "grain", ui.StarGrain, "body texture strength")
	flag.Float64Var(&ui.StarGlow, "glow", ui.StarGlow, "speckle reach past the arms")
	flag.Parse()

	draw := func(f int) string {
		art := ui.RenderStar(f, *rows, *cols)
		if !*grid {
			return art
		}
		return box(art, *cols)
	}

	if *still {
		fmt.Println(draw(*frame))
		return
	}

	if *fps < 1 {
		*fps = 1
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	fmt.Print("\x1b[?25l")
	defer fmt.Print("\x1b[?25h\n")

	tick := time.NewTicker(time.Second / time.Duration(*fps))
	defer tick.Stop()

	f := *frame
	for {
		select {
		case <-sig:
			return
		case <-tick.C:
			fmt.Print("\x1b[H\x1b[2J")
			fmt.Println(draw(f))
			fmt.Printf("\nframe %d   %dx%d   sharp %.2f  spin %.3f   ctrl-c to quit\n",
				f, *cols, *rows, ui.StarSharp, ui.StarSpin)
			f++
			if *frames > 0 && f-*frame >= *frames {
				return
			}
		}
	}
}

func box(art string, cols int) string {
	top := "+" + repeat('-', cols) + "+"
	out := top + "\n"
	for _, line := range splitLines(art) {
		for len(line) < cols {
			line += " "
		}
		out += "|" + line + "|\n"
	}
	return out + top
}

func splitLines(s string) []string {
	var out []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	return append(out, s[start:])
}

func repeat(c byte, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = c
	}
	return string(b)
}
