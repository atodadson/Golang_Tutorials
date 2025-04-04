package main

import (
    "fmt"
//     "bufio"
    "os"
    "encoding/csv"
    "strconv"
)

func getScoreInfo() (scoreData [][]string, scores []int){
    csv_file ,err := os.Open("Highscore.csv")
    if err != nil {
        fmt.Println(err)
    }
    defer csv_file.Close()

    reader := csv.NewReader(csv_file)
    records, err := reader.ReadAll()
    for _, record := range records {
        scoreData = append(scoreData, record)
        score, _ := strconv.Atoi(record[2])
        scores = append(scores, score)
    }
    scoreData, scores = scoreData[1:], scores[1:]
    return
}

func isHighScore(newscore int, scores []int) (bool) {
    for _, score := range scores {
        if newscore > score {
            fmt.Printf("newscore: %v greater than score: %v", newscore, score)
            return true
        }
    }
    return false
}

func main() {
    scoreData, scores := getScoreInfo()
    b := isHighScore(4, scores)

    fmt.Printf("%v \nSCORE DATA: \n%v \nSCORES: \n %v\n", b, scoreData, scores)

}