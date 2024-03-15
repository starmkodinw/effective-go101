package test

func For() {
	var prints []func()
	for i := 0; i < 5; i++ {
		prints = append(prints, func() { println(i) })
		i++
	}
	for _, p := range prints {
		p()
	}
	// Go 1.22.1
	// 1
	// 3
	// 5

	// else
	// 6
	// 6
	// 6
}
