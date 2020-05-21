package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	Message string = "53‡‡†305))6*;4826)4‡.)4‡);806*;48†8¶60))85;1‡(;:‡*8†83(88)5*†;46(;" +
		"88*96*?;8)*‡(;485);5*†2:*‡(;4956*2(5*_4)8¶8*;4069285);)6†8)4‡‡;1(‡9;48081;8:8‡1;" +
		"48†85;4)485†528806*81(‡9;48;(88;4(‡?34;48)4‡;161;:188;‡?;"
	MessageLength int = 203
)

type SymbolFrequency struct {
	Symbol    string
	Frequency float64
}

type EquivalentSymbol struct {
	Original string
	English  string
}

var MatchingSymbols = []string{
	"‡", "*", "†", "9", ":", ".",
}

var EnglishSymbolRanking = []SymbolFrequency{
	{"e", 12.70},
	{"t", 9.05},
	{"a", 8.16},
	{"o", 7.50},
	{"i", 6.96},
	{"n", 6.74},
	{"s", 6.32},
	{"h", 6.09},
	{"r", 5.98},
	{"d", 4.25},
	{"l", 4.02},
	{"c", 2.78},
	{"u", 2.75},
	{"m", 2.40},
	{"w", 2.36},
	{"f", 2.22},
	{"g", 2.01},
	{"y", 1.97},
	{"p", 1.92},
	{"b", 1.49},
	{"v", 0.97},
	{"k", 0.77},
	{"j", 0.15},
	{"x", 0.14},
	{"q", 0.09},
	{"z", 0.07},
}

func main() {
	symbolFrequencyMap := getSymbolRepetitionRate()
	symbolFrequencyRanking := orderSymbolFrequencyList(symbolFrequencyMap)
	equivalentSymbolList := compareWithLanguageRanking(symbolFrequencyRanking)
	textAfterMatchingSymbols := replaceSymbolsInText(Message, equivalentSymbolList)

	symbolH := searchTrigramThe(textAfterMatchingSymbols, "the", []int{1,0,2})
	textAfterH := strings.ReplaceAll(textAfterMatchingSymbols, symbolH, "h")

	symbolA := searchTrigramThe(textAfterH, "and", []int{0,1,2})
	textAfterA := strings.ReplaceAll(textAfterH, symbolA, "a")

	symbolI :=  searchWordWithFiveLetters(textAfterA, "inthe")
	textAfterI := strings.ReplaceAll(textAfterA, symbolI + "nthe", "inthe")
	fmt.Println(textAfterI)
}

func getSymbolRepetitionRate() (symbolFrequencyList []SymbolFrequency) {
	splitedMessage := strings.Split(Message, "")
	for _, symbol := range splitedMessage {
		if exist := verificateSymbolFrequencyExist(symbolFrequencyList, symbol); !exist {
			numberRepetitions := float64(strings.Count(Message, symbol))
			if verificateFrequencyExist(symbolFrequencyList, numberRepetitions){
				numberRepetitions -= 0.001
			}
			symbolFrequencyList = append(symbolFrequencyList, SymbolFrequency{symbol, numberRepetitions})
		}
	}
	return
}

func verificateFrequencyExist(symbolFrequencyList []SymbolFrequency, frequency float64) bool {
	for _, v := range symbolFrequencyList {
		if v.Frequency == frequency {
			return true
		}
	}
	return false
}

func verificateSymbolFrequencyExist(symbolFrequencyList []SymbolFrequency, symbol string) bool {
	for _, v := range symbolFrequencyList {
		if v.Symbol == symbol {
			return true
		}
	}
	return false
}

func rankBySymbolFrequency(frequencyBySymbolMap map[string]float64) (symbolFrequencyList []SymbolFrequency) {
	for symbol, frequency := range frequencyBySymbolMap {
		symbolFrequencyList = append(symbolFrequencyList, SymbolFrequency{symbol, frequency})
	}
	return orderSymbolFrequencyList(symbolFrequencyList)
}

func orderSymbolFrequencyList(symbolFrequencyList []SymbolFrequency) []SymbolFrequency{
	sort.Slice(symbolFrequencyList, func(i, j int) bool {
		return symbolFrequencyList[i].Frequency > symbolFrequencyList[j].Frequency
	})
	return symbolFrequencyList
}

func compareWithLanguageRanking(symbolFrequencyList []SymbolFrequency) (equivalentSymbolList []EquivalentSymbol) {
	for index := 0; index < len(symbolFrequencyList); index++ {
		//if verificateMatchingSymbolExist(symbolFrequencyList[index].Symbol){
			equivalentSymbolList = append(equivalentSymbolList, EquivalentSymbol{symbolFrequencyList[index].Symbol, EnglishSymbolRanking[index].Symbol})
		//}
	}
	return
}

func verificateMatchingSymbolExist(symbol string) bool {
	for _, v := range MatchingSymbols {
		if v == symbol {
			return true
		}
	}
	return false
}

func replaceSymbolsInText(text string, equivalentSymbolList []EquivalentSymbol) string {
	for _, equivalentSymbol := range equivalentSymbolList {
		text = strings.ReplaceAll(text, equivalentSymbol.Original, equivalentSymbol.English)
	}
	return text
}

func searchTrigramThe(text string, trigram string, positions []int) string {
	runes := []rune(text)
	splitedTrigram := strings.Split(trigram, "")
	var symbolFrequencyMap = make(map[string]float64)
	for index, _ := range runes {
		word := string(runes[index : index+3])
		splitedWord := strings.Split(word, "")

		if splitedWord[positions[1]] == splitedTrigram[positions[1]] && splitedWord[positions[2]] == splitedTrigram[positions[2]] {
			symbolFrequencyMap[splitedWord[positions[0]]] += 1
		}
	}
	symbolFrequencyOrderedList := rankBySymbolFrequency(symbolFrequencyMap)
	return symbolFrequencyOrderedList[0].Symbol
}


func searchWordWithFiveLetters(text string, trigram string) string {
	runes := []rune(text)
	runesTrigram :=   []rune(trigram)
	var symbolFrequencyMap = make(map[string]float64)
	for index, _ := range runes {
		word := string(runes[index : index+5])
		splitedWord := strings.Split(word, "")
		runesWord :=   []rune(word)
		if string(runesTrigram[1: 5]) == string(runesWord[1: 5]){
			symbolFrequencyMap[splitedWord[0]] += 1
		}
	}
	symbolFrequencyOrderedList := rankBySymbolFrequency(symbolFrequencyMap)
	if len(symbolFrequencyOrderedList) > 0 {
		return symbolFrequencyOrderedList[0].Symbol
	} else {
		return ""
	}

}