package main

func main() {
	// generate index.htm by template
	// add time stamp
	// scan folder structure
	// ask for auther and change detail
	sc, err := NewScanner()
	if err != nil {
		panic(err)
	}

	err = sc.Scan()
	if err != nil {
		panic(err)
	}

	err = sc.GenIndex()
	if err != nil {
		panic(err)
	}

}
