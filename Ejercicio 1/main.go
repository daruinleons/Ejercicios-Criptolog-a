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

var EnglishTopSymbol = []string{
	"e", "t",
}

var MatchingSymbols = []string{
	"‡", "*", "†", "9", ":", ".",
}

var EnglishSymbolRanking = []SymbolFrequency{
	{"e", 12.02},
	{"t", 9.10},
	{"a", 8.12},
	{"o", 7.68},
	{"i", 7.31},
	{"n", 6.95},
	{"s", 6.28},
	{"r", 6.02},
	{"h", 5.92},
	{"l", 4.32},
	{"d", 3.98},
	{"u", 2.88},
	{"c", 2.71},
	{"m", 2.61},
	{"f", 2.30},
	{"y", 2.11},
	{"w", 2.09},
	{"g", 2.03},
	{"p", 1.82},
	{"b", 1.49},
	{"v", 1.11},
	{"k", 0.69},
	{"x", 0.17},
	{"q", 0.11},
	{"j", 0.10},
	{"z", 0.07},
}

func main() {
	symbolFrequencyMap := getSymbolRepetitionRate()
	symbolFrequencyRanking := orderSymbolFrequencyList(symbolFrequencyMap)
	equivalentSymbolList := compareWithLanguageTop(symbolFrequencyRanking)
	textAfterET := replaceSymbolsInText(Message, equivalentSymbolList)
	symbolH := searchTrigramThe(textAfterET)
	textAfterH := strings.ReplaceAll(textAfterET, symbolH, "h")
	equivalentSymbolList = compareWithLanguageRanking(symbolFrequencyRanking)
	textAfterMatchingSymbols := replaceSymbolsInText(textAfterH, equivalentSymbolList)
	fmt.Println(textAfterMatchingSymbols)
}

func getSymbolRepetitionRate() (symbolFrequencyList []SymbolFrequency) {
	splitedMessage := strings.Split(Message, "")
	for _, symbol := range splitedMessage {
		if exist := verificateSymbolFrequencyExist(symbolFrequencyList, symbol); !exist {
			numberRepetitions := strings.Count(Message, symbol)
			symbolFrequencyList = append(symbolFrequencyList, SymbolFrequency{symbol, float64(numberRepetitions) / float64(MessageLength)})
		}
	}
	return
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

func compareWithLanguageTop(symbolFrequencyList []SymbolFrequency) (equivalentSymbolList []EquivalentSymbol) {
	for index := 0; index < len(EnglishTopSymbol); index++ {
		equivalentSymbolList = append(equivalentSymbolList, EquivalentSymbol{symbolFrequencyList[index].Symbol, EnglishTopSymbol[index]})
	}
	return
}

func compareWithLanguageRanking(symbolFrequencyList []SymbolFrequency) (equivalentSymbolList []EquivalentSymbol) {
	for index := 0; index < len(symbolFrequencyList); index++ {
		if verificateMatchingSymbolExist(symbolFrequencyList[index].Symbol){
			equivalentSymbolList = append(equivalentSymbolList, EquivalentSymbol{symbolFrequencyList[index].Symbol, EnglishSymbolRanking[index].Symbol})
		}
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

func searchTrigramThe(text string) string {
	runes := []rune(text)
	var symbolFrequencyMap = make(map[string]float64)
	for index, _ := range runes {
		word := string(runes[index : index+3])
		splitedWord := strings.Split(word, "")
		if splitedWord[0] == "t" && splitedWord[2] == "e" {
			symbolFrequencyMap[splitedWord[1]] += 1
		}
	}
	symbolFrequencyOrderedList := rankBySymbolFrequency(symbolFrequencyMap)
	return symbolFrequencyOrderedList[0].Symbol
}
