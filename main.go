package rainbowgo

import "fmt"

type rgb struct {
	r, g, b float64
}

func (c *rgb) clamp() {
	if c.r > 255 {
		c.r = 255
	} else if c.r < 0 {
		c.r = 0
	}
	if c.g > 255 {
		c.g = 255
	} else if c.g < 0 {
		c.g = 0
	}
	if c.b > 255 {
		c.b = 255
	} else if c.b < 0 {
		c.b = 0
	}
}

func (c *rgb) set(r, g, b float64) {
	c.r = r
	c.g = g
	c.b = b
}

func (c *rgb) cpy(cc rgb) {
	c.r = cc.r
	c.g = cc.g
	c.b = cc.b
}

func (c *rgb) addTo(i int, v float64) {
	switch i {
	case 0:
		c.r += v
		break
	case 1:
		c.g += v
		break
	case 2:
		c.b += v
		break
	}
	c.clamp()
}

func (c *rgb) n(i int) float64 {
	switch i {
	case 0:
		return c.r
	case 1:
		return c.g
	case 2:
		return c.b
	default:
		return 0
	}
}

func (c *rgb) print(s string) {
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", int(c.r), int(c.g), int(c.b), s)
}

func (c *rgb) println(s string) {
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", int(c.r), int(c.g), int(c.b), s)
}

type Rainbow struct {
	c rgb
	s float64
}

func (r *Rainbow) Init(len int) {
	if len == 0 {
		panic(`Division by zero.
	Rainbow.Init(len int)
	len must be greater than 0`)
	}
	r.c.set(255, 0, 0)
	r.s = 5.0 * 255.0 / float64(len)
}

func (r *Rainbow) Next() {
	if r.s == 0 {
		panic(`Rainbow.Next() called before Rainbow.Init(len int)`)
	}
	for i := 0; i < 3; i++ {
		if r.c.n((i+1)%3) == 0 && r.c.n((i+2)%3) == 255 {
			r.c.addTo(i, r.s)
		} else if r.c.n((i+1)%3) == 255 && r.c.n((i+2)%3) == 0 {
			r.c.addTo(i, -r.s)
		}
	}
}

func (r *Rainbow) Print(s string) {
	r.c.print(s)
}

func (r *Rainbow) Println(s string) {
	r.c.println(s)
}

func (r *Rainbow) Print2d(s string, len int) {
	var r2 Rainbow
	r2.Init(len)
	r2.c.cpy(r.c)
	for _, v := range s {
		r2.Print(string(v))
		r2.Next()
	}
}

func (r *Rainbow) Print2dln(s string, len int) {
	r.Print2d(s, len)
	fmt.Println()
}
