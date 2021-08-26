package buildtools

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const ROOT_DIR = "./tasks"
const PREFIX_CPU = "//   Runtime"
const PREFIX_MEMORY = "//   Memory"

// Task performance result struct
type TaskPerformanceResult struct {
	// leetcode task id
	taskId int
	// leetcode task name
	taskName string
	// leetcode task url name
	taskUrlName string
	// Result milliseconds
	resultCpuMs int
	// "better than" cpu result
	resultCpuPercent float32
	// result MBytes
	resultMem float32
	// "better than" cpu memory
	resultMemoryPercent float32
}

// Build markdown for performance results
// #number | name | cpu scrore | memory score
// --- | --- | --- | ---
// 1 | 2 | <span style="color:green">95%</span> | <span style="color:orange">60%</span>
func BuildPerformanceTable() string {
	fmt.Println("[buildtools] Make performance table...")
	files, err := ioutil.ReadDir(ROOT_DIR)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	// scan for tasks
	taskResults := []*TaskPerformanceResult{}
	for _, taskDir := range files {
		fmt.Printf("%v...", taskDir.Name())

		// Read eqch task dir and extract performance results
		if taskDir.IsDir() {
			if performanceResult := processTaskDirectory(taskDir); performanceResult != nil {
				taskResults = append(taskResults, performanceResult)
			}
		}
		fmt.Printf("DONE\n")

	}

	// Build markdown
	result := buildMarkdownTable(taskResults)
	saveMarkdown(result)
	fmt.Println("[buildtools] Make performance table... Done")
	return result
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
	builder.WriteString("#task id | task name | cpu (ms) | cpu (better than) | memory (mBytes) | memory (better than)\n")
	builder.WriteString("--- | --- | --- | --- | --- | ---\n")
	for _, taskResult := range taskResults {
		builder.WriteString(fmt.Sprintf(
			"%v | [%v](https://leetcode.com/problems/%v/) | %vms | %v | %vmB | %v\n",
			taskResult.taskId,
			taskResult.taskName,
			taskResult.taskUrlName,
			taskResult.resultCpuMs,
			percentMarkdown(taskResult.resultCpuPercent),
			taskResult.resultMem,
			percentMarkdown(taskResult.resultMemoryPercent),
		))
	}
	return builder.String()
}

// Process task directory: find source files and parse the cpu and memory results
func processTaskDirectory(taskDir fs.FileInfo) *TaskPerformanceResult {
	dir := ROOT_DIR + "/" + taskDir.Name()
	taskId, err := strconv.Atoi(taskDir.Name()[4:])
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
		printPerformanceResult(result)
		return result
	}
	return nil
}

func printPerformanceResult(p *TaskPerformanceResult) {
	fmt.Printf("id='%v', url='%v', name='%v', cpu=(%vms, >%v%%), mem=(%vMb, >%v%%) \n", p.taskId, p.taskUrlName, p.taskName, p.resultCpuMs, p.resultCpuPercent, p.resultMem, p.resultMemoryPercent)
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
		resultMem:           0,
		resultMemoryPercent: 0,
	}
	scanner := bufio.NewScanner(file)
	doneCpu := false
	doneMem := false
	for (!doneCpu || !doneMem) && scanner.Scan() {
		text := scanner.Text()

		// Read CPU results
		if strings.HasPrefix(text, PREFIX_CPU) {
			doneCpu = true
			milliseconds, percent, taskName, errCpu := parseCpuString(text)
			if errCpu != nil {
				log.Fatal(errCpu)
				continue
			}
			result.taskName = taskName
			result.resultCpuMs = milliseconds
			result.resultCpuPercent = percent
		}

		// Read Memory results
		if strings.HasPrefix(text, PREFIX_MEMORY) {
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
func parseCpuString(s string) (milliseconds int, percent float32, taskName string, err error) {
	re := regexp.MustCompile(`//\s*Runtime: (?P<milliseconds>\d+) ms, faster than (?P<percent>\d+\.*\d*)% of Go online submissions for (?P<name>.*).`)
	subNames := re.FindSubmatch([]byte(s))
	msResult, err := strconv.Atoi(string(subNames[1]))
	if err != nil {
		return 0, 0.0, "", err
	}
	percentResult, err2 := strconv.ParseFloat(string(subNames[2]), 32)
	if err2 != nil {
		return 0, 0.0, "", err2
	}
	return msResult, float32(percentResult), string(subNames[3]), nil
}

// Parse string ""//   Memory Usage: 4.2 MB, less than 61.84% of Go online submissions for Two Sum.""
func parseMemString(s string) (mBytes float32, percent float32, err error) {
	re := regexp.MustCompile(`//\s*Memory Usage:\s*(?P<mbytes>\d+\.{0,1}\d*)\s*MB, less than (?P<percent>\d+\.{0,1}\d*)%.*`)
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
