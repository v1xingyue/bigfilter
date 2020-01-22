package main

import (
	"bigfilter/parser"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	log "github.com/sirupsen/logrus"
)

type obj map[string]interface{}

var (
	jsonBytes, rulesBytes []byte
)

func init() {
	jsonBytes, _ = ioutil.ReadFile("test.json")
	rulesBytes, _ = ioutil.ReadFile("rule.txt")
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true})
	s := "xingyue123"
	r := "xingyue\\d+"
	m, _ := regexp.Match(r, []byte(s))
	log.Println(m)

	var rulesString = string(rulesBytes)
	log.Println("rule raw: ", rulesString)
	var info map[string]interface{}
	json.Unmarshal([]byte(jsonBytes), &info)
	ev, err := parser.NewEvaluator(rulesString)
	if err != nil {
		log.Fatal(fmt.Errorf("Error making evaluator from the rule %v, %v", rulesString, err))
	}
	ans, err := ev.Process(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(
		fmt.Sprintf(`
                Rule: %v
                Object: %v
                Answer: %v
        `, rulesString, info, ans),
	)
	if err := ev.LastDebugErr(); err != nil {
		fmt.Println("Last debug error", ev.LastDebugErr())
	}

}
