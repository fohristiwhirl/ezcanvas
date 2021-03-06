package ezcanvas

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func (c *Canvas) Line(r, g, b uint8, mode int, x1, y1, x2, y2 int) {

	if x1 == x2 {
		c.lineVertical(r, g, b, mode, x1, y1, y2)
	} else if y1 == y2 {
		c.lineHorizontal(r, g, b, mode, x1, y1, x2)
	} else {
		c.lineBresenham(r, g, b, mode, x1, y1, x2, y2)
	}
}

func (c *Canvas) lineHorizontal(r, g, b uint8, mode int, x1, y, x2 int) {

	var sx int

	if x1 < x2 {
		sx = 1
	} else {
		sx = -1
	}

	for {
		c.SetByMode(r, g, b, mode, x1, y)
		if x1 == x2 {
			break
		}
		x1 += sx
	}
}

func (c *Canvas) lineVertical(r, g, b uint8, mode int, x, y1, y2 int) {

	var sy int

	if y1 < y2 {
		sy = 1
	} else {
		sy = -1
	}

	for {
		c.SetByMode(r, g, b, mode, x, y1)
		if y1 == y2 {
			break
		}
		y1 += sy
	}
}

func (c *Canvas) lineBresenham(r, g, b uint8, mode int, x1, y1, x2, y2 int) {

	// http://members.chello.at/~easyfilter/bresenham.html

	dx := abs(x2 - x1)
	dy := abs(y2 - y1) * -1

	var sx int
	var sy int

	if x1 < x2 { sx = 1 } else { sx = -1 }
	if y1 < y2 { sy = 1 } else { sy = -1 }

	err := dx + dy

	for {
		c.SetByMode(r, g, b, mode, x1, y1)
		if (x1 == x2 && y1 == y2) { break }
		e2 := 2 * err
		if (e2 >= dy) {
			err += dy
			x1 += sx
		}
		if (e2 <= dx) {
			err += dx
			y1 += sy
		}
	}
}

/*

func (c *Canvas) lineGentle(r, g, b uint8, mode int, x1, y1, x2, y2 int) {

	// Based on an algorithm I read on the web 15 years ago;
	// The webpage has long since vanished.

	var additive int

	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}

	if (y1 < y2) {
		additive = 1;
	} else {
		additive = -1;
	}

	dy_times_two := (y2 - y1) * 2
	if dy_times_two < 0 { dy_times_two *= -1 }

	dx_times_two := (x2 - x1) * 2       // We know we're going right, no need to check for < 0

	the_error := x1 - x2

	for n := x1 ; n <= x2 ; n++ {

		c.SetByMode(r, g, b, mode, n, y1)

		the_error += dy_times_two;
		if the_error > 0 {
			y1 += additive
			the_error -= dx_times_two
		}
	}
}

func (c *Canvas) lineSteep(r, g, b uint8, mode int, x1, y1, x2, y2 int) {

	var additive int

	if y1 > y2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}

	if (x1 < x2) {
		additive = 1;
	} else {
		additive = -1;
	}

	dy_times_two := (y2 - y1) * 2       // We know we're going down, no need to check for < 0

	dx_times_two := (x2 - x1) * 2
	if dx_times_two < 0 { dx_times_two *= -1 }

	the_error := y1 - y2;

	for n := y1 ; n <= y2 ; n++ {

		c.SetByMode(r, g, b, mode, x1, n)

		the_error += dx_times_two
		if the_error > 0 {
			x1 += additive
			the_error -= dy_times_two
		}
	}
}

*/
