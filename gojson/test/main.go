package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var jsonStr = "{\r\n\t\"a\": \"1\",\r\n\t\"b\": {\r\n\t\t\"c\": 1,\r\n\t\t\"d\": {\r\n\t\t\t\"e\": 11111111\r\n\t\t}\r\n\t},\r\n\t\"f\": \"https: //api.tradplusad.com/api/v1_2/adconf\"\r\n}"

var jsonData = map[string]interface{}{
	"a":1,
	"b":"2",
	"c":map[string]interface{}{
		"d":111,
	},
}

var structHeader = `package main

type Request struct {
`

var structFooter = "}"


func main() {
	//s := map2struct(jsonData, 0)
	//fmt.Println(fmt.Sprintf("%s%s%s", structHeader, s, structFooter))

	a := "\"aaa\""
	fmt.Println(1111111, a[:1] == "\"")

	s1 := strings.Replace(jsonStr, "\t", "", -1)
	sArr := strings.Split(s1, "\r\n")
	s2 := json2struct(sArr)

	fmt.Println(fmt.Sprintf("%s%s%s", structHeader, s2, structFooter))
}

func json2struct(sArr []string) string {

	s := ""
	n := 0

	sNodeEndMap := map[int]string{}

	for _, v := range sArr {

		jArr := strings.Split(v,"\": ")

		if len(jArr) == 1 {
			jv := jArr[0]

			if jv == "{" {
				n ++
			}

			if jv == "}" || jv == "},"  {
				s += sNodeEndMap[n]
				n --
			}

			//todo: 数组判断 [ ]

			continue
		}

		if len(jArr) == 2 {
			jk := strings.Trim(jArr[0], "\"")
			jv := strings.Trim(jArr[1], ",")

			if jv[:1] == "\"" {
				m := map[string]interface{} {
					jk: strings.Trim(jv, "\""),
				}
				s += string2Struct(m, n)
				continue
			}

			if jv[:1] == "{" {

				n ++

				mapNode := map[string]interface{}{
					jk: struct{}{},
				}

				if strings.Index(mapNode2Strcut(mapNode, n), "${NODE}") != -1 {
					s += strings.Split(mapNode2Strcut(mapNode, n), "${NODE}")[0]

					sNodeEndMap[n] = strings.Split(mapNode2Strcut(mapNode, n), "${NODE}")[1]
				}
				continue
			}

			if jvInt, err := strconv.Atoi(jv[:1]); err == nil {
				m := map[string]interface{} {
					jk: jvInt,
				}
				s += string2Struct(m, n)
				continue
			}

			if jvInt, err := strconv.ParseFloat(jv[:1], 64); err == nil {
				m := map[string]interface{} {
					jk: jvInt,
				}
				s += string2Struct(m, n)
				continue
			}

			if jvInt, err := strconv.ParseBool(jv[:1]); err == nil {
				m := map[string]interface{} {
					jk: jvInt,
				}
				s += string2Struct(m, n)
				continue
			}

			fmt.Println("err")
		}

	}

	return s
}

func map2struct(maps map[string]interface{}, n int) string {

	var s string

	for k, v := range maps{
		switch v.(type) {
		case map[string]interface{}:

			mapNode := map[string]interface{}{
				k: struct{}{},
			}

			node := mapNode2Strcut(mapNode, n+1)

			s1 := map2struct(v.(map[string]interface{}), n+1)

			s += strings.Replace(node, "${NODE}", s1, -1)
		default:
			m := map[string]interface{}{
				k:v,
			}
			s += string2Struct(m, n)
		}
	}

	return s
}


func string2Struct(m map[string]interface{}, n int) string {
	s, _ := json.Marshal(m)

	command := fmt.Sprintf("echo '%s' | gojson -name=Request", s)

	out, err := exec.Command("bash", "-c", command).Output()

	var result string
	if err != nil {
		result = err.Error()
	} else {
		result = string(out)
	}

	result = strings.Trim(result, structHeader)
	result = strings.Trim(result, structFooter)
	result = getTab(n) + result

	return result
}

func mapNode2Strcut(m map[string]interface{}, n int) string {
	s, _ := json.Marshal(m)

	command := fmt.Sprintf("echo '%s' | gojson -name=Request", s)

	out, err := exec.Command("bash", "-c", command).Output()

	var result string
	if err != nil {
		result = err.Error()
	} else {
		result = string(out)
	}

	result = strings.Trim(result, structHeader)
	result = strings.Trim(result, structFooter)
	result = getTab(n-1) + result

	result = strings.Replace(result, "struct{}", fmt.Sprintf(`struct{
${NODE}%s}`, getTab(n)), -1)


	//fmt.Println(333333333, result)

	return result
}

func getTab(n int) string {
	tab := ""
	for i:=0;i<n;i++ {
		tab += "        "
	}
	return tab
}