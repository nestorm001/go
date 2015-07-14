package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const outputFileName = "address.xml"
const outputCodeFileName = "javaCode.txt"
const openFileName = "GB2260.txt"
var province string
var isDirectCity bool = false
var isFirstCity bool = false

func ReadLine(filePth string, hook func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hook(line)      //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func getCodeAndName(line []byte) (code int, name string) {
	result := strings.Split(strings.Split(strings.Split(string(line), "\n")[0], "\r")[0], "	")
	code, _ = strconv.Atoi(result[0])
	name = result[1]
	return code, name
}

func getProvince() {
	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
	defer fCodeOut.Close()
	fCodeOut.WriteString("private int[] city = {")

	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
	defer fOut.Close()
	fOut.WriteString("\n\t<!-- 省级 -->" + "\n")
	fOut.WriteString("\t<string-array name=\"省\">" + "\n")

	ReadLine(openFileName, saveProvince)

	fOut.WriteString("\t</string-array>" + "\n\n")

	fCodeOut.WriteString("\n};\n")
}
func saveProvince(line []byte) {
	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
	defer fOut.Close()
	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
	defer fCodeOut.Close()
	code, name := getCodeAndName(line)
	if code % 10000 == 0 {
		fOut.WriteString("\t\t<item>" + name + "</item>" + "\n")
		if code != 110000 {
			fCodeOut.WriteString(", \n\t\tR.array." + name)
		} else {
			fCodeOut.WriteString("\n\t\tR.array." + name)
		}
	}
}

func getCity() {
	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
	defer fOut.Close()
	fOut.WriteString("\n\t<!-- 市级 -->" + "\n")
	ReadLine(openFileName, saveCity)
	fOut.WriteString("\t</string-array>" + "\n\n")
}
func saveCity(line []byte) {
	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
	defer fOut.Close()
	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
	defer fCodeOut.Close()
	code, name := getCodeAndName(line)
	if code % 10000 == 0 {
		if code != 110000 {
			fOut.WriteString("\t</string-array>" + "\n\n")
		}
		fOut.WriteString("\t<string-array name=\"" + name + "\">" + "\n")
		isDirectCity = false
		province = name
		isFirstCity = true
	}
	isCity := code % 10000 != 0 && code % 100 == 0
	if isCity || isDirectCity {
		if name == "市辖区" || name == "县" ||
		name == "省直辖县级行政区划" || name == "自治区直辖县级行政区划" {
			if name == "市辖区" {
				fOut.WriteString("\t\t<item>" + province + "</item>" + "\n")
			}
			if name == "省直辖县级行政区划" || name == "自治区直辖县级行政区划" {
				isDirectCity = true
			}
		} else {
			fOut.WriteString("\t\t<item>" + name + "</item>" + "\n")
		}
	}
}

func getCounty() {
	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
	defer fOut.Close()
	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
	defer fCodeOut.Close()
	fOut.WriteString("\n\t<!-- 区县级 -->" + "\n")
	ReadLine(openFileName, saveCounty)
	fOut.WriteString("\t</string-array>" + "\n\n")
	fCodeOut.WriteString("\n};\n")
}
func saveCounty(line []byte) {
	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
	defer fOut.Close()
	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
	defer fCodeOut.Close()
	code, name := getCodeAndName(line)
	if code % 10000 == 0 {
		province = name
		isDirectCity = false
		if code != 110000 {
			fCodeOut.WriteString("\n};")
		}
		fCodeOut.WriteString("\nprivate int[] " + name + " = {")
	}
	isCity := code % 10000 != 0 && code % 100 == 0
	if isCity || isDirectCity {
		if code != 110100 && name != "县" && name != "省直辖县级行政区划" && name != "自治区直辖县级行政区划" {
			fOut.WriteString("\t</string-array>" + "\n\n")
		}
		if name == "市辖区" || name == "县" ||
		name == "省直辖县级行政区划" || name == "自治区直辖县级行政区划" {
			if name == "市辖区" {
				fOut.WriteString("\t<string-array name=\"" + province + name + "\">" + "\n")
				fCodeOut.WriteString("\n\t\tR.array." + province + name)
			}
			if name == "省直辖县级行政区划" || name == "自治区直辖县级行政区划" {
				isDirectCity = true
			}
		} else {
			fOut.WriteString("\t<string-array name=\"" + province + name + "\">" + "\n")
			fCodeOut.WriteString("\n\t\tR.array." + province + name)
		}
	}
	if code % 100 != 0 {
		if name != "市辖区" {
			fOut.WriteString("\t\t<item>" + name + "</item>" + "\n")
		}
	}
}

func main() {
	os.Remove(outputFileName)
	os.Create(outputFileName)
	os.Remove(outputCodeFileName)
	os.Create(outputCodeFileName)
	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
	defer fOut.Close()
	fOut.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>" + "\n")
	fOut.WriteString("<resources" + "\n")
	fOut.WriteString("\t\txmlns:tools=\"http://schemas.android.com/tools\"" + "\n")
	fOut.WriteString("\t\ttools:ignore=\"MissingTranslation\">" + "\n")
	getProvince()
	getCity()
	getCounty()
	fOut.WriteString("</resources>" + "\n")
}
