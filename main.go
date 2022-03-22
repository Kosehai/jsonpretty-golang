package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

//Colors
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func main() {
	reader := bufio.NewReader(os.Stdin)
	var output []rune
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	jsondata := ""

	for i := 0; i < len(output); i++ {
		jsondata += string(output[i])
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(jsondata), &result)
	fmt.Println(printJsonLayer(result, 0))
}

func printJsonLayer(jsondata map[string]interface{}, depth int) string {
	output := ""
	for k, v := range jsondata {
		outstr := ""
		tabs := ""
		for i := 1; i <= depth; i++ {
			tabs += "\t"
		}
		outstr += fmt.Sprintf("%s%s\"%s\"%s: ", tabs, Cyan, k, Reset)
		switch e := v.(type) {
		case string:
			outstr += fmt.Sprintf("%s\"%s\"%s", Green, e, Reset)
		case float64:
			outstr += fmt.Sprintf("%s%.0f%s", Purple, e, Reset)
		case bool:
			outstr += fmt.Sprintf("%s%t%s", Yellow, e, Reset)
		case []interface{}:
			outstr += "["
			for _, x := range e {
				elmtype := ""
				switch x.(type) {
				case string:
					elmtype = Green + "\"%s\", " + Reset
				case float64:
					elmtype = Purple + "%0.f, " + Reset
				case bool:
					elmtype = Yellow + "%t, " + Reset
				default:
					elmtype = "%s, "
				}
				outstr += fmt.Sprintf(elmtype, x)
			}
			outstr = strings.TrimSuffix(outstr, ", "+Reset)
			outstr += Reset + "]"
		case map[string]interface{}:
			outstr += "{\n"
			layer := jsondata[k].(map[string]interface{})
			outstr += printJsonLayer(layer, depth+1)
			outstr += tabs + "}"
		}

		output += outstr + tabs + "\n"
	}
	return output
}
