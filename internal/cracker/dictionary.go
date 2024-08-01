package cracker

import (
    "fmt"
    "sync"
    "github.com/muesli/termenv"
    "github.com/Spandan7724/enigma/internal/hash"
    "github.com/Spandan7724/enigma/internal/wordlist"
)

func DictionaryAttack(hashType string, targetHash string, wordlistPath string) {
    p := termenv.ColorProfile()
    
    if hashType == "" {
        var err error
        hashType, err = hash.DetectHashType(targetHash)
        if err != nil {
            fmt.Println(termenv.String("Error detecting hash type:").Foreground(p.Color("1")).String(), err)
            return
        }
        fmt.Println(termenv.String("Detected hash type:").Foreground(p.Color("2")).Bold().String(), hashType)
    }

    words, err := wordlist.LoadWordlist(wordlistPath)
    if err != nil {
        fmt.Println(termenv.String("Error loading wordlist:").Foreground(p.Color("1")).String(), err)
        return
    }

    fmt.Println(termenv.String("Loaded wordlist with").Foreground(p.Color("6")).String(), len(words), termenv.String("words").Foreground(p.Color("6")).String())

    var wg sync.WaitGroup
    wordChan := make(chan string)
    resultChan := make(chan string)
    doneChan := make(chan bool)

    // Worker goroutines
    for i := 0; i < 10; i++ { 
        go worker(hashType, targetHash, wordChan, resultChan, &wg)
    }

    // Result handler
    go func() {
        for result := range resultChan {
            fmt.Println(termenv.String("Password found:").Foreground(p.Color("2")).Bold().String(), result)
            close(doneChan)
            return
        }
    }()

    // Feed the word channel
    go func() {
        for _, word := range words {
            wordChan <- word
        }
        close(wordChan)
    }()

    select {
    case <-doneChan:
        close(resultChan)
    case <-wait(&wg):
        fmt.Println(termenv.String("Password not found in wordlist").Foreground(p.Color("1")).String())
    }
}

func worker(hashType, targetHash string, wordChan <-chan string, resultChan chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    for word := range wordChan {
        variations := wordlist.GenerateVariations(word)
        for _, variation := range variations {
            var hashedWord string
            switch hashType {
            case "md5":
                hashedWord = hash.MD5Hash(variation)
            case "sha1":
                hashedWord = hash.SHA1Hash(variation)
            case "sha256":
                hashedWord = hash.SHA256Hash(variation)
            default:
                return
            }
            if hashedWord == targetHash {
                resultChan <- variation
                return
            }
        }
    }
}

func wait(wg *sync.WaitGroup) <-chan struct{} {
    ch := make(chan struct{})
    go func() {
        wg.Wait()
        close(ch)
    }()
    return ch
}