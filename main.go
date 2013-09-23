package main

// import ()
import "fmt"
import "time"
import "os/exec"

type point struct {
	X, Y int
}
type cuska []point

func addPoints(c cuska) cuska {
	c = append(c, point{5, 5})
	c = append(c, point{5, 6})
	c = append(c, point{5, 7})
	c = append(c, point{5, 8})
	c = append(c, point{6, 8})
	c = append(c, point{7, 8})
	return c
}

func checkIfInPoint(c cuska, p point) bool {
	for _, o := range c {
		if p.X == o.X && p.Y == o.Y {
			return true
		}
	}
	return false
}

func main() {
	c := addPoints(make(cuska, 0, 100))
	w := 10
	h := 10

	o, _ := exec.Command("clear").Output()

	for {
		fmt.Printf("%s", o)
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				if checkIfInPoint(c, point{i, j}) {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Printf("\n")
		}
		time.Sleep(time.Second)
	}
}
