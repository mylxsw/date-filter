package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var timeLayout = flag.String("layout", "2006/01/02 15:04:05", "日期模板，遵循golang规范")
var validTime = flag.Duration("valid-time", 10*time.Minute, "有效时间，当前时间前 valid_time 分钟之内的日志将会输出")
var timeOffsetPos = flag.Int("offset", 0, "时间截取位置")
var debug = flag.Bool("debug", false, "是否开启调试模式")

func main() {

	flag.Parse()

	location, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatalf("Parse Error: %v", err)
	}

	startPos := *timeOffsetPos
	endPos := len(*timeLayout) + *timeOffsetPos

	// 根据该标识判断下一行无法解析的时候是否继续输出
	continueOutput := false

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		// 如果当前行小于时间位置，则直接根据continueOutput判断是否输出
		if len(line) < endPos {
			if continueOutput {
				fmt.Print(line)
			}
			continue
		}

		// 解析时间，如果时间解析出错，则认为当前行不包含时间，根据continueOutput判断是否输出
		logTimeStr := line[startPos:endPos]
		logTime, err := time.ParseInLocation(*timeLayout, logTimeStr, location)
		if *debug && err != nil {
			log.Printf("Parse Date [%s] Error: %v", logTimeStr, err)
		}

		if err != nil {
			if continueOutput {
				fmt.Print(line)
			}
			continue
		}

		diff := time.Now().Sub(logTime)
		if diff < *validTime {
			fmt.Print(line)
			continueOutput = true
		} else {
			continueOutput = false
		}
	}
}
