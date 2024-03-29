package main

// GOMAXPROCS

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go a()
	go b()
}
