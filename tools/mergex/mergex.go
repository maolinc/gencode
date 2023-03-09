package mergex

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Merge() {

}

// mergeFiles 合并两个文件的内容并将结果写入新文件
func MergeFiles(fileP1, fileP2, output string) error {

	// 打开第一个文件
	file1, err := os.Open(fileP1)
	if err != nil {
		return err
	}
	defer file1.Close()

	// 打开第二个文件
	file2, err := os.Open(fileP2)
	if err != nil {
		return err
	}
	defer file2.Close()

	// 创建一个 map，用于保存所有的行
	allLines := make(map[string]bool)

	// 读取第一个文件的每一行并添加到 map 中
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		allLines[scanner.Text()] = true
	}

	// 读取第二个文件的每一行并添加到 map 中
	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		allLines[scanner.Text()] = true
	}

	// 将 map 中的所有键按字典序排序，并保存到字符串切片中
	sortedLines := make([]string, 0, len(allLines))
	for line := range allLines {
		sortedLines = append(sortedLines, line)
	}
	sort.Strings(sortedLines)

	// 打开输出文件
	outputFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 将排序后的行写入输出文件
	writer := bufio.NewWriter(outputFile)
	for _, line := range sortedLines {
		fmt.Fprintln(writer, line)
	}

	// 刷新缓存并检查是否有错误
	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

// splitLines 将字符串按行拆分成字符串切片
func splitLines(content string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// sortLines 将 map 中的所有键按字典序排序，并返回排序后的字符串切片
func sortLines(lines map[string]bool) []string {
	var sorted []string
	for line := range lines {
		sorted = append(sorted, line)
	}
	sort.Strings(sorted)
	return sorted
}

// joinLines 将字符串切片合并成一个字符串，每行以换行符分隔
func joinLines(lines []string) string {
	return strings.Join(lines, "\n")
}
