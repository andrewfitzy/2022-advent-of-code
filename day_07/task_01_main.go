package day_07

import "bufio"
import "strings"
import "regexp"
import "strconv"
import "github.com/emirpasic/gods/stacks/arraystack"

func solve01(input string) int {
	matcher := regexp.MustCompile(`\d+`)
	pwd := arraystack.New()
	filePlan := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		// check if the line is a command
		if strings.HasPrefix(line, "$") {
			command := strings.Split(line, " ")
			if command[1] == "ls" {
				continue // ignore ls
			}

			if command[2] == ".." {
				pwd.Pop()
				continue
			}

			pwd.Push(command[2])
			continue
		}

		entry := strings.Split(line, " ")
		matches := matcher.FindString(entry[0])
		if len(matches) > 0 { // ignore anything that doesn't start with a filesize
			fileSize, _ := strconv.Atoi(matches)

			pwdParts := pwd.Values()
			tmpPwd := ""
			for i := len(pwdParts) - 1; i >= 0; i-- { // add file size to all entries that make up part of the path
				tmpPwd = tmpPwd + pwdParts[i].(string)
				filePlan[tmpPwd] = filePlan[tmpPwd] + fileSize
			}
		}
	}

	//find and sum entries <= 100000
	result := 0
	for _, sizeOnDisk := range filePlan {
		if sizeOnDisk <= 100000 {
			result = result + sizeOnDisk
		}
	}

	return result
}
