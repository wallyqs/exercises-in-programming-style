package main

import (
        "fmt"
        "os"
        ospath "path/filepath"
        IO "io/ioutil"
        "bufio"
        "strings"
        "bytes"
        "sort"
)

type wordsSortableByFrequency []*wordWithFrequency

type wordWithFrequency struct {
        word string
        frequency int
}

func (d wordsSortableByFrequency) Len() int {
        return len(d)
}

func (d wordsSortableByFrequency) Swap(i, j int) {
        d[i], d[j] = d[j], d[i]
}

func (d wordsSortableByFrequency) Less(i, j int) bool {
        return d[i].frequency < d[j].frequency
}

func main() {
        // Get the current path
        //
        currentDir, _ := os.Getwd()
        // TODO: Check the error

        srcRoot := ospath.Join(currentDir, "src", "exercises-in-programming-style")

        // Join the path til the repo with the text
        //
        PRIDE_AND_PREJUDICE := ospath.Join(srcRoot, "pride-and-prejudice.txt")
        STOP_WORDS := ospath.Join(srcRoot, "stop_words.txt")

        // Generate the stop words and put them in an array
        // in case of stopwords we can just read the file and put in memory
        //
        stopWordsContents, _ := IO.ReadFile(STOP_WORDS)
        stopWordsContents = stopWordsContents[0:(len(stopWordsContents) - 3)] // remove three line breaks at the end
        // TODO: check the error

        // Split the contents of the file to generate the words to ignore
        //
        stopWords := strings.Split(strings.ToLower(string(stopWordsContents)), ",")

        // Now merge the single letters too...
        // Generate the alphabet in lowercase: a..z (97..123 in ascii)
        //
        for i := 97; i < 123; i++ {
                stopWords = append(stopWords, string(i))
        }

        // Leave this one open (defer closing)
        prideAndPrejudiceTextFile, _ := os.Open(PRIDE_AND_PREJUDICE)
        defer prideAndPrejudiceTextFile.Close()

        // For reading the pride and prejudice text, we use a scanner instead
        //
        prideAndPrejudiceTextReader := bufio.NewReader(prideAndPrejudiceTextFile)
        scanner := bufio.NewScanner(prideAndPrejudiceTextReader)

        // Only capture lowercase alphanumeric characters
        //
        wordFrequency := make(map[string]int)
        var wordBuffer bytes.Buffer
        for scanner.Scan() {
                line := strings.ToLower(scanner.Text())

                for _, c := range line {
                        if (c >= 97 && c <= 123) {       // Filter alphanumeric
                                wordBuffer.WriteRune(c)
                        } else if c == 32 {             // Empty space, meaning that we have a word
                                w := wordBuffer.String()

                                wordFrequency[w]++      // Accumulate the word
                                wordBuffer.Truncate(0)  // reset the buffer
                        }
                }
        }

        // Remove the words that should be ignored
        //
        for _, word := range stopWords {
                delete(wordFrequency, word)
        }

        // Turn the wordFrequency map into a list so that entries are comparable?
        //
        sortableWordsList := make(wordsSortableByFrequency, 0, len(wordFrequency))
        for word, frequency := range wordFrequency {
                w := wordWithFrequency{ word, frequency }
                sortableWordsList = append(sortableWordsList, &w) // pass reference
        }

        // Sort!
        //
        sort.Sort(sort.Reverse(sortableWordsList))
        for i, w := range sortableWordsList {
                fmt.Println(w.word, " - ", w.frequency)
                if i > 25 {
                        break
                }
        }
}
