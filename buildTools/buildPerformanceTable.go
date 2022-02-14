package buildtools

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

const ROOT_DIR = "./tasks"

var PREFIXES_CPU []string
var PREFIXES_MEMORY []string

// Task performance result struct
type TaskPerformanceResult struct {
	// leetcode task id
	taskId int
	// language
	lang string
	// leetcode task name
	taskName string
	// leetcode task url name
	taskUrlName string
	// Link to file
	taskFileUrl string
	// Result milliseconds
	resultCpuMs int
	// "better than" cpu result
	resultCpuPercent float32
	// Is CPU results are unstable
	resultCpuIsUnstable bool
	// result MBytes
	resultMem float32
	// "better than" cpu memory
	resultMemoryPercent float32
}

// Build markdown for performance results
// #number | name | cpu scrore                           | memory score
// ---     | ---  | ---                                  | ---
// 1       | 2    | <span style="color:green">95%</span> | <span style="color:orange">60%</span>
func BuildPerformanceTable() string {
	PREFIXES_CPU = []string{"//   Runtime", "--   Runtime"}
	PREFIXES_MEMORY = []string{"//   Memory"}

	fmt.Println("[buildtools] Make performance table...")
	// scan for tasks
	taskResults := readDir(ROOT_DIR)
	fmt.Printf("[buildtools] All tasks were read (%v)\n", len(taskResults))

	// Build markdown
	fmt.Println("[buildtools] Build markdown")
	result := buildMarkdownTable(taskResults)
	saveMarkdown(result)
	fmt.Println("[buildtools] Make performance markdown... Done")
	return result
}

// Scan directory and subdirectories for "taskXXXX" folders and extract task info from those folders
func readDir(parentDirectory string) []*TaskPerformanceResult {
	files, err := ioutil.ReadDir(parentDirectory)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	taskResults := []*TaskPerformanceResult{}
	for _, file := range files {
		if file.IsDir() {
			dirName := file.Name()
			if strings.HasPrefix(dirName, "task") {
				if performanceResult := processTaskDirectory(parentDirectory, dirName); performanceResult != nil {
					taskResults = append(taskResults, performanceResult)
				}
			} else {
				taskResults = append(taskResults, readDir(path.Join(parentDirectory, dirName))...)
			}
		}
	}
	return taskResults
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Save markdown
func saveMarkdown(markdown string) {
	f, err := os.Create("./README_PERFORMANCE.md")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err2 := w.WriteString(markdown + "\n")
	check(err2)
	w.Flush()
}

// Generate markdown for task results
func buildMarkdownTable(taskResults []*TaskPerformanceResult) string {

	percentMarkdown := func(val float32) string {
		color := "#000"
		if val < 50 {
			color = "#f00"
		} else if val >= 50 && val < 80 {
			color = "orange"
		} else if val >= 80 && val < 90 {
			color = "yellow"
		} else {
			color = "green"
		}
		return fmt.Sprintf("<span style=\"color:%v\">> %v%%</span>", color, val)
	}

	builder := new(strings.Builder)
	builder.WriteString("#task id | lang | task name | leetcode | cpu (ms) | cpu (better than) | memory (mBytes) | memory (better than)\n")
	builder.WriteString("--- | --- | --- | --- | --- | --- | --- | ---\n")
	for _, taskResult := range taskResults {
		unstableSign := ""
		if taskResult.resultCpuIsUnstable {
			unstableSign = " [*](#performanceFootnote)"
		}
		builder.WriteString(fmt.Sprintf(
			"%v | %v | [%v](%v) | [leetcode](https://leetcode.com/problems/%v/) | %vms | %v%v | %vmB | %v\n",
			taskResult.taskId,
			taskResult.lang,
			taskResult.taskName,
			taskResult.taskFileUrl,
			taskResult.taskUrlName,
			taskResult.resultCpuMs,
			percentMarkdown(taskResult.resultCpuPercent),
			unstableSign,
			taskResult.resultMem,
			percentMarkdown(taskResult.resultMemoryPercent),
		))
	}
	return builder.String()
}

// Process task directory: find source files and parse the cpu and memory results
func processTaskDirectory(rootDir, taskDir string) *TaskPerformanceResult {
	dir := rootDir + "/" + taskDir
	taskId, err := strconv.Atoi(taskDir[4:])
	if err != nil {
		log.Fatal(err)
		return nil
	}
	goFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for _, goFile := range goFiles {
		goFileName := goFile.Name()
		if strings.HasSuffix(goFileName, "_test.go") {
			continue
		}
		result := readTaskFile(dir+"/"+goFile.Name(), goFile.Name())
		if result == nil {
			return nil
		}
		result.taskId = taskId
		result.taskFileUrl = dir + "/" + goFileName
		printPerformanceResult(result)
		return result
	}
	return nil
}

func printPerformanceResult(p *TaskPerformanceResult) {
	fmt.Printf("id='%v', lang='%v', url='%v', name='%v', cpu=(%vms, >%v%%, u=%v), mem=(%vMb, >%v%%) \n", p.taskId, p.lang, p.taskUrlName, p.taskName, p.resultCpuMs, p.resultCpuPercent, p.resultCpuIsUnstable, p.resultMem, p.resultMemoryPercent)
}

// Read single task file
func readTaskFile(goFileName, goOnlyFileName string) *TaskPerformanceResult {

	urlName := goOnlyFileName[:len(goOnlyFileName)-3]

	// Read file
	file, err := os.Open(goFileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	// Scan lines
	result := &TaskPerformanceResult{
		taskId:              0,
		taskName:            "",
		taskUrlName:         urlName,
		resultCpuMs:         0,
		resultCpuPercent:    0,
		resultCpuIsUnstable: false,
		resultMem:           0,
		resultMemoryPercent: 0,
	}
	scanner := bufio.NewScanner(file)
	doneCpu := false
	doneMem := false
	for (!doneCpu || !doneMem) && scanner.Scan() {
		text := scanner.Text()

		// Read CPU results
		hasCpuPrefix := false
		for _, prefix := range PREFIXES_CPU {
			if strings.HasPrefix(text, prefix) {
				hasCpuPrefix = true
				break
			}
		}
		if hasCpuPrefix {
			doneCpu = true
			milliseconds, percent, lang, isUnstable, taskName, errCpu := parseCpuString(text)
			if errCpu != nil {
				log.Fatal(errCpu)
				continue
			}
			result.taskName = taskName
			result.lang = lang
			result.resultCpuMs = milliseconds
			result.resultCpuPercent = percent
			result.resultCpuIsUnstable = isUnstable
		}

		// Read Memory results
		hasMemPrefix := false
		for _, prefix := range PREFIXES_MEMORY {
			if strings.HasPrefix(text, prefix) {
				hasMemPrefix = true
				break
			}
		}
		if hasMemPrefix {
			doneMem = true
			mBytes, percent, errMem := parseMemString(text)
			if errMem != nil {
				log.Fatal(errMem)
				continue
			}
			result.resultMem = mBytes
			result.resultMemoryPercent = percent
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}
	return result
}

// Parse string "//   Runtime: 4 ms, faster than 98.33% of Go online submissions for Add Two Numbers."
func parseCpuString(s string) (milliseconds int, percent float32, lang string, isUnstable bool, taskName string, err error) {
	re := regexp.MustCompile(`^[/\-][/\-]\s*Runtime:\s*(?P<milliseconds>\d+)\s*ms, faster than (?P<percent>\d+\.*\d*)% of (?P<lang>.*?) online submissions for (?P<name>.*)\..*?(?P<unstable>\*Unstable\*|)$`)
	subNames := re.FindSubmatch([]byte(s))
	msResult, err := strconv.Atoi(string(subNames[1]))
	if err != nil {
		return 0, 0.0, "", false, "", err
	}
	percentResult, err2 := strconv.ParseFloat(string(subNames[2]), 32)
	if err2 != nil {
		return 0, 0.0, "", false, "", err2
	}
	unstableResult := subNames[5]
	return msResult, float32(percentResult), string(subNames[3]), len(unstableResult) > 0, string(subNames[4]), nil
}

// Parse string ""//   Memory Usage: 4.2 MB, less than 61.84% of Go online submissions for Two Sum.""
func parseMemString(s string) (mBytes float32, percent float32, err error) {
	re := regexp.MustCompile(`[/\-][/\-]\s*Memory Usage:\s*(?P<mbytes>\d+\.{0,1}\d*)\s*MB, less than (?P<percent>\d+\.{0,1}\d*)%.*`)
	subNames := re.FindSubmatch([]byte(s))
	mBytesResult, err := strconv.ParseFloat(string(subNames[1]), 32)
	if err != nil {
		return 0.0, 0.0, err
	}
	percentResult, err2 := strconv.ParseFloat(string(subNames[2]), 32)
	if err2 != nil {
		return 0.0, 0.0, err
	}
	return float32(mBytesResult), float32(percentResult), nil
}
