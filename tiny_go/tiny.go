package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"bytes"
	"os"
	"image/png"
	"io"
	"bufio"
	"path/filepath"
	"encoding/hex"
	"crypto/md5"
	"strconv"
)

const targetUrl = "https://api.tinify.com/shrink"
const configFile = "tiny_config"
const fileInfoFile = "compress_info"
const stopCompressRatio = 0.9

var client http.Client
var config Configs
var files map[string]fileInfo

type Configs struct {
	ApiKey         string `json:"api_key"`
	CompressFolder []string`json:"compress_folder"`
	IgnoreFile     []string`json:"ignore_file"`
}

type TinyResult struct {
	Input  Input `json:"input"`
	Output Output `json:"output"`
}

type Input struct {
	Size int64 `json:"size"`
	Type string `json:"type"`
}

type Output struct {
	Size   int64 `json:"size"`
	Type   string `json:"type"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
	Ratio  float64 `json:"ratio"`
	Url    string `json:"url"`
}

type fileInfo struct {
	name  string
	md5   string
	ratio float64
	exist bool
}

func initConfigs() {
	jsonBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonBytes, &config)
	if err != nil {
		panic(err)
	}

	fmt.Println("api key: " + config.ApiKey)
	fmt.Println()
	fmt.Println("folders to compress:")
	fmt.Println(config.CompressFolder)
	fmt.Println()
	fmt.Println("ignore files:")
	for i := range config.IgnoreFile {
		config.IgnoreFile[i] = strings.Replace(config.IgnoreFile[i], "/", string(filepath.Separator), -1)
	}
	fmt.Println(config.IgnoreFile)
}

func getAllPaths() []string {
	var paths []string
	for _, path := range config.CompressFolder {
		if !contains(config.IgnoreFile, path) {
			paths = append(paths, getFileList(path)...)
		}
	}
	return paths
}

func compressAllFile(paths []string) {
	for _, path := range paths {
		var fileInfo fileInfo
		var exists bool
		fileInfo, exists = files[path]

		md5Value := calculateMd5(path)

		fmt.Print(path)
		fmt.Print("======>")

		if exists {
			if fileInfo.ratio > stopCompressRatio && fileInfo.md5 == md5Value {
				fmt.Println("pass")
				fileInfo.exist = true
				files[path] = fileInfo
				continue
			}
		} else {
			fileInfo.name = path
		}
		fmt.Println("compress")
		tinyResult := compressFile(path)
		url := tinyResult.Output.Url
		fmt.Println("url: " + url)
		fmt.Printf("ratio: %f\n", tinyResult.Output.Ratio)
		saveImage(url, path)
		ratio := tinyResult.Output.Ratio

		fileInfo.exist = true
		fileInfo.ratio = ratio
		fileInfo.md5 = calculateMd5(path)

		files[path] = fileInfo
	}
}

func saveFileInfo() {
	os.Remove(fileInfoFile)
	os.Create(fileInfoFile)
	out, _ := os.OpenFile(fileInfoFile, os.O_APPEND, 0)
	defer out.Close()
	for path, fileInfo := range files {
		if !fileInfo.exist {
			delete(files, path)
		} else {
			ratioString := strconv.FormatFloat(fileInfo.ratio, 'f', 6, 64)
			info := fileInfo.name + "," + fileInfo.md5 + "," + ratioString + "\n"
			fmt.Print(info)
			out.WriteString(info)
		}
	}
}

func readFileInfo() {
	files = make(map[string]fileInfo)
	f, err := os.Open(fileInfoFile)
	defer f.Close()
	if err != nil {
		return
	}

	bfRd := bufio.NewReader(f)
	lineNum := 0
	for {
		line, err := bfRd.ReadBytes('\n')
		lineNum++
		result := string(line)
		result = strings.Split(strings.Split(result, "\n")[0], "\r")[0]
		if len(strings.TrimSpace(result)) == 0 {
			return
		}
		info := strings.Split(result, ",")
		var fileInfo fileInfo
		fileInfo.name = info[0]
		fileInfo.md5 = info[1]
		fileInfo.ratio, _ = strconv.ParseFloat(info[2], 64)
		fileInfo.exist = false
		files[fileInfo.name] = fileInfo

		if err != nil {
			if err == io.EOF {
				return
			}
		}
	}
}

func getFileList(path string) []string {
	var paths []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if contains(config.IgnoreFile, path) {
			return nil
		}
		if f == nil {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		if !checkIsNormalPng(path) {
			return nil
		}
		paths = append(paths, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return paths
}

func compressFile(path string) TinyResult {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewBuffer(body))
	req.SetBasicAuth("api", config.ApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var tinyResult TinyResult
	json.Unmarshal(b, &tinyResult)
	return tinyResult
}

func saveImage(url string, path string) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	file, _ := os.Create(path)
	defer file.Close()
	io.Copy(file, res.Body)
}

func contains(paths []string, path string) bool {
	for _, s := range paths {
		match, _ := filepath.Match(s, filepath.Dir(path))
		sameFile := s == path
		if match || sameFile {
			return true
		}
	}
	return false
}

func calculateMd5(path string) string {
	file, err := os.Open(path)
	defer file.Close()
	md5h := md5.New()
	if err == nil {
		io.Copy(md5h, file)
	}
	result := hex.EncodeToString(md5h.Sum([]byte("")))
	return result
}

func checkIsNormalPng(file string) bool {
	reader, _ := os.Open(file)
	defer reader.Close()
	_, err := png.DecodeConfig(reader)
	return err == nil && !checkNinePatch(file)
}

func checkNinePatch(file string) bool {
	return strings.HasSuffix(file, ".9.png")
}

func main() {
	client = http.Client{}
	initConfigs()
	readFileInfo()
	compressAllFile(getAllPaths())
	saveFileInfo()
}
