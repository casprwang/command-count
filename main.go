package cmc

import (
	"bufio"
	"log"
	"os"
	"path"
	"runtime"
	"sort"
)

func UserHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

func GetSlice() {
	dir := UserHomeDir()

	filePath := path.Join(dir, ".zsh_history")

	src, err := os.Open(filePath)

	if err != nil {
		log.Fatalln(err)
	}

	defer src.Close()
	count := map[string]int{}

	scanner := bufio.NewScanner(src)

	counter := 0
	for scanner.Scan() {
		line := string(scanner.Text())

		counter++
		count[line]++
		// fmt.Println(len(count))
		// fmt.Println(">>>", string(line))
	}

	type kv struct {
		Key string
		Val int
	}

	ss := []kv{}

	for k, v := range count {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i int, j int) bool {
		return ss[i].Val < ss[j].Val
	})

	return ss

}

func GetPercentage(a int, total int) float32 {
	return (float32(a) / float32(total)) * 100
}
