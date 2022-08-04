package dup

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// Dup 找出重复行并输出
func Dup() {
	counts := make(map[string]int)
	files := os.Args[1:]
	hasDupLine := false
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
			linesSlice, countSlice := sortCountsDesc(counts)
			tempHasDupLine := false
			for index, count := range countSlice {
				if count > 1 {
					// 练习1.4 出现重复行文件名
					if !tempHasDupLine {
						fmt.Printf("file have duplicate lines, name: %s \n", arg)
					}
					fmt.Printf("%d\t%s\n", count, linesSlice[index])
					hasDupLine = true
					tempHasDupLine = true
				}
			}
			tempHasDupLine = false
		}
	}

	if !hasDupLine {
		fmt.Println("there are no duplicate lines!")
	}
}

// RepeatLineFile 从文件中找出重复行
func RepeatLineFile() {
	counts := make(map[string]int)
	for _, filemame := range os.Args[1:] {
		data, err := ioutil.ReadFile(filemame)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup read file err:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func sortCountsDesc(counts map[string]int) ([]string, []int) {
	countSlice := make([]int, 0, len(counts))
	linesSlice := make([]string, 0, len(counts))

	for line := range counts {
		linesSlice = append(linesSlice, line)
	}
	sort.SliceStable(linesSlice, func(i, j int) bool {
		return counts[linesSlice[i]] < counts[linesSlice[j]]
	})

	for _, line := range linesSlice {
		countSlice = append(countSlice, counts[line])
	}
	return linesSlice, countSlice
}

// Pair 由map构造的结构体
type Pair struct {
	Key   string
	Value int
}

// TestStableUseSlice 测试map以稳定方式排序，目前存在相同value值key的排序结果不一致问题，已解决
func TestStableUseSlice() {
	counts := make(map[string]int)
	f, err := os.Open("/Users/boroughfan/GitDocuments/GoLangPractise/ch01/dup/text_feel_the_light_lyrics.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup:%v\n", err)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	f.Close()
	///////////////////////////////////////////////////////////
	linesSlice := make([]string, 0, len(counts))

	for line := range counts {
		linesSlice = append(linesSlice, line)
	}
	sort.SliceStable(linesSlice, func(i, j int) bool {
		if counts[linesSlice[i]] != counts[linesSlice[j]] {
			return counts[linesSlice[i]] < counts[linesSlice[j]]
		}
		// in this example: if lines have same count, order them alphabetically:
		return linesSlice[i] < linesSlice[j]
	})

	for _, line := range linesSlice {
		fmt.Printf("%d\t%s\n", counts[line], line)
	}
}

// TestStableUsePair 测试map以稳定方式排序，目前存在相同value值key的排序结果不一致问题
func TestStableUsePair() {
	counts := make(map[string]int)
	f, err := os.Open("/Users/boroughfan/GitDocuments/GoLangPractise/ch01/dup/text_feel_the_light_lyrics.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup:%v\n", err)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	f.Close()
	///////////////////////////////////////////////////////////
	pairList := make([]Pair, 0, len(counts))
	for line := range counts {
		pairList = append(pairList, Pair{line, counts[line]})
	}
	sort.SliceStable(pairList, func(i, j int) bool { return pairList[i].Value < pairList[j].Value })
	for _, pairs := range pairList {
		fmt.Printf("%d\t%s\n", pairs.Value, pairs.Key)
	}
}

// TestStandardStable 按照官方用例给出的map稳定排序，相同value值输出的key稳定
func TestStandardStable() {
	people := []struct {
		Name string
		Age  int
	}{
		{"I still remember when time was frozen", 1},
		{"Now we have another chance to fly", 1},
		{"\"Feel The Light\"", 1},
		{"Did you expect me to reason with thunder", 1},
		{"You and I can have it all tonight", 1},
		{"(from \"Home\" soundtrack)", 1},
		{"Do you remember when we fell under", 1},
		{"Hmm", 1},
		{"But put together the cracks we'll close in", 1},
		{"Hmm, hmm", 1},
		{"// 测试数据，来源于feel the light歌词", 1},
		{"I still remember when things were broken", 1},
		{"Another chance to make it right", 1},
		{"What seemed forever was just a moment", 1},
		{"So let's bring it back to life", 1},
		{"// this is the dup test file, contents are from the feel the light lyrics", 1},
		{"We're still worth saving", 2},
		{"Here we go, here we go", 2},
		{"Hurry up, hurry up", 2},
		{"There's no more waiting", 2},
		{"Feel better now, feel better now", 3},
		{"Shining in the dark of night", 3},
		{"Shining like the stars tonight", 3},
		{"It's better now, feel better now", 3},
		{"Here I go, here I go", 4},
		{"But we're bringing it all back", 5},
		{"We're bringing it all back", 5},
		{"Feel the light", 6},
		{"I know it's a long shot", 6},
		{"Remember what we forgot", 6},
	}

	// Sort by age preserving name order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	for _, onePeople := range people {
		fmt.Printf("%d\t%s\n", onePeople.Age, onePeople.Name)
	}

}
