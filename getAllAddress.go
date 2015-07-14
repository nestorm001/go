//package main
//
//import (
//	"bufio"
//	"io"
//	"os"
//	"strconv"
//	"strings"
//)
//
//const outputFileName = "address.xml"
//const outputCodeFileName = "javaCode.txt"
//const openFileName = "GB2260.txt"
//var province string
//var isDirectCity bool = false
//var isFirstCity bool = false
//
//func ReadLine(filePth string, hook func([]byte)) error {
//	f, err := os.Open(filePth)
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//	bfRd := bufio.NewReader(f)
//	for {
//		line, err := bfRd.ReadBytes('\n')
//		hook(line)      //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
//		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
//			if err == io.EOF {
//				return nil
//			}
//			return err
//		}
//	}
//	return nil
//}
//
//func getCodeAndName(line []byte) (code int, name string) {
//	result := strings.Split(strings.Split(string(line), "\n")[0], "	")
//	code, _ = strconv.Atoi(result[0])
//	name = result[1]
//	return code, name
//}
//
//func getProvince() {
//	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
//	defer fCodeOut.Close()
//	fCodeOut.WriteString("private int[] city = {")
//
//	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
//	defer fOut.Close()
//	fOut.WriteString("\n<!-- 省级 -->" + "\r")
//	fOut.WriteString("<string-array name=\"省\">" + "\r")
//
//	ReadLine(openFileName, saveProvince)
//
//	fOut.WriteString("</string-array>" + "\n\r")
//
//	fCodeOut.WriteString("\r};")
//}
//func saveProvince(line []byte) {
//	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
//	defer fOut.Close()
//	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
//	defer fCodeOut.Close()
//	code, name := getCodeAndName(line)
//	if code % 10000 == 0 {
//		fOut.WriteString("<item>" + name + "</item>" + "\r")
//		if code != 110000 {
//			fCodeOut.WriteString(", \rR.array." + name)
//		} else {
//			fCodeOut.WriteString("\rR.array." + name)
//		}
//	}
//}
//
//func getCity() {
//	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
//	defer fOut.Close()
//	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
//	defer fCodeOut.Close()
//
//	fOut.WriteString("\n<!-- 市级 -->" + "\r")
//
//	ReadLine(openFileName, saveCity)
//
//	fOut.WriteString("</string-array>" + "\n\r")
//
//	fCodeOut.WriteString("\r};")
//}
//func saveCity(line []byte) {
//	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
//	defer fOut.Close()
//	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
//	defer fCodeOut.Close()
//	code, name := getCodeAndName(line)
//	if code % 10000 == 0 {
//		if code != 110000 {
//			fOut.WriteString("</string-array>" + "\n\r")
//			fCodeOut.WriteString("\r};")
//		}
//		fOut.WriteString("<string-array name=\"" + name + "\">" + "\r")
//		isDirectCity = false
//		province = name
//		isFirstCity = true
//		fCodeOut.WriteString("\rprivate int[] " + name + " = {")
//	}
//	isCity := code % 10000 != 0 && code % 100 == 0
//	if isCity || isDirectCity {
//		if name == "市辖区" || name == "县" ||
//		name == "省直辖县级行政区划" || name == "自治区直辖县级行政区划" {
//			if name == "市辖区" {
//				fOut.WriteString("<item>" + province + "</item>" + "\r")
//				fCodeOut.WriteString("\rR.array." + province + name)
//			}
//			if name == "省直辖县级行政区划" || name == "自治区直辖县级行政区划" {
//				isDirectCity = true
//			}
//		} else {
//			fOut.WriteString("<item>" + name + "</item>" + "\r")
//			if code % 100 == 0 {
//				if !isFirstCity {
//					fCodeOut.WriteString(", \rR.array." + province + name)
//				} else {
//					fCodeOut.WriteString("\rR.array." + province + name)
//					isFirstCity = false
//				}
//			}
//		}
//	}
//}
//
//func getCounty() {
//	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
//	defer fOut.Close()
//	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
//	defer fCodeOut.Close()
//	fOut.WriteString("\n<!-- 区县级 -->" + "\r")
//	ReadLine(openFileName, saveCounty)
//	fOut.WriteString("</string-array>" + "\n\r")
//}
//func saveCounty(line []byte) {
//	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
//	defer fOut.Close()
//	fCodeOut, _ := os.OpenFile(outputCodeFileName, os.O_APPEND, 0)
//	defer fCodeOut.Close()
//	code, name := getCodeAndName(line)
//	if code % 10000 == 0 {
//		province = name
//		isDirectCity = false
//	}
//	isCity := code % 10000 != 0 && code % 100 == 0
//	if isCity || isDirectCity {
//		if code != 110100 && name != "县" && name != "市辖区" && name != "省直辖县级行政区划" && name != "自治区直辖县级行政区划" {
//			fOut.WriteString("</string-array>" + "\n\r")
//		}
//		if name == "市辖区" || name == "县" ||
//		name == "省直辖县级行政区划" || name == "自治区直辖县级行政区划" {
//			if name == "市辖区" {
//				fOut.WriteString("<string-array name=\"" + province + name + "\">" + "\r")
//			}
//			if name == "省直辖县级行政区划" || name == "自治区直辖县级行政区划" {
//				isDirectCity = true
//			}
//		} else {
//			fOut.WriteString("<string-array name=\"" + province + name + "\">" + "\r")
//		}
//	}
//	if code % 100 != 0 {
//		if name != "市辖区" {
//			fOut.WriteString("<item>" + name + "</item>" + "\r")
//		}
//	}
//}
//
//func main() {
//	os.Remove(outputFileName)
//	os.Create(outputFileName)
//	os.Remove(outputCodeFileName)
//	os.Create(outputCodeFileName)
//	fOut, _ := os.OpenFile(outputFileName, os.O_APPEND, 0)
//	defer fOut.Close()
//	fOut.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>" + "\n")
//	fOut.WriteString("<resources" + "\n")
//	fOut.WriteString("xmlns:tools=\"http://schemas.android.com/tools\"" + "\n")
//	fOut.WriteString("tools:ignore=\"MissingTranslation\">" + "\n")
//	getProvince()
//	getCity()
//	getCounty()
//	fOut.WriteString("</resources>" + "\n")
//}
