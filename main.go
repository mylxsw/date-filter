package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var timeLayout = flag.String("layout", "2006/01/02 15:04:05", "日期模板，遵循golang规范")
var validTime = flag.Duration("valid-time", 10*time.Minute, "有效时间，当前时间前 valid_time 分钟之内的日志将会输出")
var timeOffsetPos = flag.String("offset", "0:19", "时间截取位置")

func parseTimePos(timeOffsetPos string) (int, int, error) {
	timeOffsets := strings.Split(timeOffsetPos, ":")
	if len(timeOffsets) != 2 {
		return 0, 0, fmt.Errorf("时间截取位置错误")
	}

	timeStartPos, err := strconv.Atoi(timeOffsets[0])
	if err != nil {
		return 0, 0, err
	}

	timeEndpos, err := strconv.Atoi(timeOffsets[1])
	if err != nil {
		return 0, 0, err
	}

	return timeStartPos, timeEndpos, nil
}

func main() {

	flag.Parse()

	timeStartPos, timeEndPos, err := parseTimePos(*timeOffsetPos)
	if err != nil {
		os.Exit(2)
	}

	location, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatalf("Parse Error: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		logTime, err := time.ParseInLocation(*timeLayout, line[timeStartPos:timeEndPos], location)
		if err != nil {
			log.Fatalf("Parse Date Error: %v", err)
		}

		diff := time.Now().Sub(logTime)
		if diff < *validTime {
			fmt.Print(line)
		}
	}
}
