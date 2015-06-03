package main

func main() {
	// ask for auther and change detail
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Document Index Generator 1.0 by Test OP4")
	// fmt.Println("Please enter `User` and `Reason for change` to re-scan")
	// fmt.Print("User             : ")
	// user, err := reader.ReadString('\n')
	// if err != nil {
	// 	panic(err)
	// }
	// user = strings.Trim(user, "\n")

	// fmt.Print("Reason for change: ")
	// reason, err := reader.ReadString('\n')
	// if err != nil {
	// 	panic(err)
	// }
	// reason = strings.Trim(reason, "\n")

	// scan for changes
	// start := time.Now()
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

	// save history
	// if _, err = os.Stat("history.txt"); os.IsNotExist(err) {
	// 	os.Create("history.txt")
	// }

	// history, err := os.OpenFile("history.txt", os.O_RDWR|os.O_APPEND, 0660)
	// if err != nil {
	// 	panic(err)
	// }

	// defer history.Close()
	// histLog := fmt.Sprintf("DATE:%s USER:%s REASON:%s\n", start.Format(time.RFC850), user, reason)
	// _, err = history.WriteAt([]byte(histLog), 0)
	// if err != nil {
	// 	panic(err)
	// }

	// stop := time.Since(start)
	// fmt.Print(histLog)
	// fmt.Printf("Done in %s. Please Enter to exit", stop)
	// reader.ReadString('\n')
}
