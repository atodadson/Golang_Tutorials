// func getPlayerNames(message string) (player string) {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println(message)
// 	// var container string
// 	// fmt.Scanf("%s", &container)
// 	player_name, err := reader.ReadString('\n')
// 	player = strings.TrimSpace(player_name)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// package main

// import (
// 	"fmt"

// type HighScore struct {
// 	Rank     string `json:"rank"`
// 	Score    int    `json:"score"`
// 	DateTime string `json:"datetime"`
// }

// type HighScores struct {
// 	HighScores []HighScore `json: "highscore"`
// }

// returns the path (string) where the HighScore Json file is stored
// depending on the OS the program is running on
// func getHighScorePath() (file_path string) {
// 	var base string
// 	if runtime.GOOS == "windows" {
// 		base = os.Getenv("APPDATA")
// 	} else {
// 		base = os.Getenv("HOME")
// 	}
// 	file_path = filepath.Join(base, "highscore.json")
// 	return
// }

// func fileExists(filename string) bool {
// 	// Use os.Stat to get the file information
// 	_, err := os.Stat(filename)

// 	// Check if the error is of type "file does not exist"
// 	if os.IsNotExist(err) {
// 		return false
// 	}

// 	// If no error, the file exists
// 	return true
// }

// Takes the HighScores struct as input
// and Write the content to the Highscore file
// func writeHighScore(highscores HighScores) {
// 	// Get the directory for the highscore json
// 	file_path := getHighScorePath()

// 	// Format the content to a json format
// 	jsonData, _ := json.MarshalIndent(highscores, "", " ")

// 	// Write the jsonData to the file
// 	os.WriteFile(file_path, jsonData, 0644)
// }

// func getScoreInfo3() (scoreData [][]string, scores []int) {
// 	file_path := getHighScorePath()
// 	highScores := HighScores{}

// 	highScoreFile, err := os.ReadFile(file_path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	extract_err := json.Unmarshal([]byte(highScoreFile), &highScores)
// 	if extract_err != nil {
// 		log.Fatal(extract_err)
// 	}

// 	for _, highscore := range highScores.HighScores {
// 		scores = append(scores, highscore.Score)
// 		scoreData = append(scoreData, []string{
// 			strconv.Itoa(highscore.Rank),
// 			strconv.Itoa(highscore.Score),
// 			highscore.DateTime,
// 		})
// 	}

// 	return

// csv_file, err := os.Open("Highscore.csv")
// if err != nil {
// 	fmt.Println(err)
// }
// defer csv_file.Close()

// reader := csv.NewReader(csv_file)
// records, err := reader.ReadAll()
// for _, record := range records {
// 	// fmt.Printf("THis is record: %v", record)
// 	scoreData = append(scoreData, record)
// 	score, _ := strconv.Atoi(record[1])
// 	scores = append(scores, score)
// }
// scoreData, scores = scoreData[1:], scores[1:]
// }

// var scoreList = HighScores{
// 	HighScores: []HighScore{
// 		{Rank: "1", Score: 24, DateTime: "2025-04-08 09:51:09"},
// 		{Rank: "2", Score: 12, DateTime: "2025-04-08 10:51:09"},
// 	},
// }

// func getScoreInfo() (scoreData [][]string, scores []int) {
// 	file_path := getHighScorePath()
// 	highScores := HighScores{}

// 	highScoreFile, err := os.ReadFile(file_path)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	extract_err := json.Unmarshal([]byte(highScoreFile), &highScores)
// 	if extract_err != nil {
// 		fmt.Println(extract_err)
// 	}

// 	for _, highscore := range highScores.HighScores {
// 		scores = append(scores, highscore.Score)
// 		scoreData = append(scoreData, []string{
// 			highscore.Rank,
// 			strconv.Itoa(highscore.Score),
// 			highscore.DateTime,
// 		})
// 	}

// 	return
// }

// func SortData(data [][]string) [][]string {
// 	// Sorting the HighScores slice based on Score
// 	sort.Slice(highScores.HighScores, func(i, j int) bool {
// 		return highScores.HighScores[i].Score > highScores.HighScores[j].Score // Descending order
// 	})

// 	// Custom sorting function based on the age (index 2)
// 	sort.SliceStable(data, func(i, j int) bool {
// 		// Convert the age from string to integer for proper comparison
// 		ageI, errI := strconv.Atoi(data[i][1])
// 		ageJ, errJ := strconv.Atoi(data[j][1])

// 		// If there's an error in conversion, consider them equal for now
// 		if errI != nil || errJ != nil {
// 			return false
// 		}

// 		// Sort in descending order based on age
// 		return ageI > ageJ
// 	})

// 	// Update index 0 (ranking) based on the sorted order
// 	for i := range data {
// 		data[i][0] = strconv.Itoa(i + 1) // Update ranking with 1-based index
// 	}

// 	return data
// }

// // go:embed dictionary5.txt
// var diction string

// func main() {
// 	fmt.Println(diction)
// }

// package main

// import (
// 	_ "embed"
// 	"fmt"
// 	"strings"
// )

// // Embed the text file.
// //
// //go:embed dictionary5.txt
// var content string

// func main() {
// 	lines := strings.Split(content, "\n")
// 	fmt.Printf("%T", lines)
// }

package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed mywords.txt
var mywordsString string

//go:embed new_words.txt
var new_wordsString string

//go:embed dictionary5.txt
var dictionaryString string

func main() {
	mywords := strings.Split(mywordsString, "\n")
	new_words := strings.Split(new_wordsString, "\n")
	dictionary := strings.Split(dictionaryString, "\n")

	all_words := append(mywords, new_words...)
	all_words = append(all_words, dictionary...)

	fmt.Printf("len mywords: %v newwords: %v dictionary: %v all_words: %v", len(mywords), len(new_words), len(dictionary), len(all_words))

	var unique_words []string
	my_map := make(map[string]string)
	for _, word := range all_words {
		if _, exists := my_map[word]; !exists {
			my_map[word] = word
			unique_words = append(unique_words, word)
		} else {
			fmt.Printf("This is duplicate %s\n", word)
		}
	}
	fmt.Printf("This is the len of unique words %v", len(unique_words))

	file, _ := os.OpenFile("new_dictionary.txt", os.O_CREATE|os.O_WRONLY, 0644)
	for _, word := range unique_words {
		file.WriteString(word)
	}
}
