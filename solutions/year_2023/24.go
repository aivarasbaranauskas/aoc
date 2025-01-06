package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_matrix"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"math"
	"strconv"
	"strings"
)

func init() {
	Solutions[24] = Day24{}
}

type Day24 struct{}

func (d Day24) Part1(input []byte) string {
	hailstones := d.parseInput(input)
	from, to := 200000000000000., 400000000000000.

	type Line struct {
		a, b float64
	}

	lines := _slice.Map(hailstones, func(h Hailstone) Line {
		// x0 + xd * t = x ; xd * t = x - x0 ; t = (x - x0) / xd
		// t = (y - y0) / yd
		// (y - y0) / yd = (x - x0) / xd
		// y/yd - y0/yd = x/xd - x0/xd
		// y/yd = x/xd - x0/xd + y0/yd
		// y = x*yd/xd - x0*yd/xd + y0
		// y = a*x + b ; a = yd/xd ; b = y0 - x0*yd/xd

		return Line{
			a: h.dir.y / h.dir.x,
			b: h.pos.y - h.pos.x*h.dir.y/h.dir.x,
		}
	})

	ct := 0

	for i, line1 := range lines {
		for j := i + 1; j < len(lines); j++ {
			line2 := lines[j]

			// y = ax+b
			// a1x + b1 = a2x + b2
			// (a1 - a2) * x = b2 - b1
			// x = (b2 - b1) / (a1 - a2)
			x := (line2.b - line1.b) / (line1.a - line2.a)
			y := line1.a*x + line1.b
			if from <= x && x <= to && from <= y && y <= to {
				// check if t for both is positive
				// t = (x - x0) / xd
				t1 := (x - hailstones[i].pos.x) / hailstones[i].dir.x
				t2 := (x - hailstones[j].pos.x) / hailstones[j].dir.x
				if t1 >= 0 && t2 >= 0 {
					ct++
				}
			}
		}
	}
	return strconv.Itoa(ct)
}

func (d Day24) Part2(input []byte) string {
	hailstones := d.parseInput(input)
	for x := -500.; x <= 500; x++ {
		for y := -500.; y <= 500; y++ {
			// check if all intersects in same point on xy plane with [x, y] vector adjustment in positive time
			v := Point3{x: x, y: y}
			p, ts := d.tryFindIntersectionOnXyPlaneWithPositiveT(hailstones, v)
			if p == nil {
				continue
			}

			// try find p_z with which the intersection stays true in 3D and z is same, all still with positive time
		ZLoop:
			for z := -500.; z <= 500; z++ {
				v.z = z

				first := true
				for i, h := range hailstones {
					pT := h.CalcModPos(v, ts[i])
					if first {
						p.z = pT.z
						first = false
						continue
					}

					if math.Abs(pT.z-p.z) > 1 {
						continue ZLoop
					}
				}

				return strconv.Itoa(int(p.x) + int(p.y) + int(p.z))
			}
		}
	}
	return "???"
}

func (Day24) tryFindIntersectionOnXyPlaneWithPositiveT(hailstones []Hailstone, v Point3) (*Point3, []float64) {
	var (
		pInt  *Point3
		allTs []float64
	)

	h1 := hailstones[0]
	h1NewD := h1.dir.Add(v)

	for j := 1; j < len(hailstones); j++ {
		h2 := hailstones[j]
		h2NewD := h2.dir.Add(v)

		a := [][]float64{
			{h1NewD.x, -1 * h2NewD.x},
			{h1NewD.y, -1 * h2NewD.y},
		}
		b := [][]float64{
			{h2.pos.x - h1.pos.x},
			{h2.pos.y - h1.pos.y},
		}

		aInv, ok := _matrix.Inverse(a)
		if !ok {
			return nil, nil
		}

		ts := _matrix.Multiply(aInv, b)
		t1, t2 := ts[0][0], ts[1][0]
		if t1 >= 0 && t2 >= 0 {
			p := h1.CalcModPos(v, t1)

			if pInt == nil {
				pInt = &p
				allTs = append(allTs, t1, t2)
			} else if math.Abs((*pInt).x-p.x) > 1 || math.Abs((*pInt).y-p.y) > 1 {
				return nil, nil
			} else {
				allTs = append(allTs, t2)
			}
		} else {
			return nil, nil
		}
	}

	return pInt, allTs
}

func (Day24) parseInput(input []byte) []Hailstone {
	return _slice.Map(
		strings.Split(string(input), "\n"),
		func(line string) Hailstone {
			spl := strings.Split(line, "@")
			p1 := _slice.Map(strings.Split(spl[0], ","), func(n string) float64 { return float64(optimistic.Atoi(strings.TrimSpace(n))) })
			p2 := _slice.Map(strings.Split(spl[1], ","), func(n string) float64 { return float64(optimistic.Atoi(strings.TrimSpace(n))) })

			return Hailstone{
				pos: Point3{
					x: p1[0],
					y: p1[1],
					z: p1[2],
				},
				dir: Point3{
					x: p2[0],
					y: p2[1],
					z: p2[2],
				},
			}
		},
	)
}

type Point3 struct {
	x, y, z float64
}

func (p Point3) Add(p2 Point3) Point3 {
	return Point3{
		x: p.x + p2.x,
		y: p.y + p2.y,
		z: p.z + p2.z,
	}
}

type Hailstone struct {
	pos, dir Point3
}

func (h Hailstone) CalcPos(t float64) Point3 {
	return Point3{
		x: h.pos.x + h.dir.x*t,
		y: h.pos.y + h.dir.y*t,
		z: h.pos.z + h.dir.z*t,
	}
}

func (h Hailstone) CalcModPos(v Point3, t float64) Point3 {
	return Point3{
		x: h.pos.x + (h.dir.x+v.x)*t,
		y: h.pos.y + (h.dir.y+v.y)*t,
		z: h.pos.z + (h.dir.z+v.z)*t,
	}
}
