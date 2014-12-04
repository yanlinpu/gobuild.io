package main

func main() {
	for x := 0; x < 20; x++ {
		println("iter ", x)
		p := make([]byte, 100<<20)
		for i := range p {
			p[i] = byte(i)
		}
	}
}
