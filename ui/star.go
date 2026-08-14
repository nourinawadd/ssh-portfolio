package ui

import (
	"math"
	"strings"
)

// starRamp maps luminance to glyphs, sparsest first.
const starRamp = " .:-=+*#%@"

// Tunables for RenderStar. Kept as vars so cmd/star can sweep them while iterating.
var (
	StarSharp     = 0.70  // superellipse exponent; below 1 gives the concave ✦ arms
	StarGamma     = 0.65  // brightness falloff along the arms; lower = arms stay bright
	StarCore      = 1.28  // overall gain, lets the middle saturate
	StarStretch   = 1.18  // vertical elongation
	StarSpin      = 0.000 // radians of rotation per frame; 0 keeps the ✦ upright
	StarPulse     = 0.065 // size breathing amplitude
	StarPulseRate = 0.095 // breathing speed
	StarGrain     = 0.19  // texture strength across the body
	StarGlow      = 0.35  // speckle reach beyond the arms
)

// RenderStar draws a slowly turning, twinkling four-pointed star as ASCII
// density art sized to rows x cols. Output is plain text; the caller adds colour.
func RenderStar(frame, rows, cols int) string {
	if rows < 1 || cols < 1 {
		return ""
	}

	cx := float64(cols-1) / 2
	cy := float64(rows-1) / 2

	// Vertical radius in rows. Character cells are about twice as tall as they
	// are wide, so the horizontal radius is double that, less the stretch.
	ry := math.Min(cy*0.96, cx*0.48*StarStretch)
	if ry < 1 {
		ry = 1
	}
	ry *= 1 + StarPulse*math.Sin(float64(frame)*StarPulseRate)
	rx := 2 * ry / StarStretch

	a := float64(frame) * StarSpin
	sinA, cosA := math.Sin(a), math.Cos(a)

	var sb strings.Builder
	sb.Grow(rows * (cols + 1))

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ox := (float64(j) - cx) / rx
			oy := (float64(i) - cy) / ry

			sx := ox*cosA + oy*sinA
			sy := -ox*sinA + oy*cosA

			// A superellipse with an exponent below 1 has concave sides and four
			// sharp points — the ✦ silhouette. f <= 1 is inside the star.
			f := math.Pow(math.Abs(sx), StarSharp) + math.Pow(math.Abs(sy), StarSharp)

			var v float64
			if f <= 1 {
				b := math.Pow(1-f, StarGamma)
				n := starHash(j*3+i*7, i*5-j, frame/4)
				v = b * StarCore * (1 - StarGrain + StarGrain*2*n)
			} else {
				reach := (f - 1) / StarGlow
				if reach >= 1 {
					sb.WriteByte(' ')
					continue
				}
				n := starHash(j*7+i*3, i, frame/3+i+j)
				if n > 0.34*(1-reach) {
					sb.WriteByte(' ')
					continue
				}
				v = 0.06 + 1.02*(1-reach)*n
			}

			sb.WriteByte(starRampByte(v))
		}
		if i < rows-1 {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

func starRampByte(v float64) byte {
	if v <= 0 {
		return ' '
	}
	if v > 1 {
		v = 1
	}
	return starRamp[int(v*float64(len(starRamp)-1)+0.5)]
}

func starHash(a, b, c int) float64 {
	x := math.Sin(float64(a)*127.1+float64(b)*311.7+float64(c)*74.7) * 43758.5453
	return x - math.Floor(x)
}
