package xcode

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

const (
	FAILING_TEST_MATCHER         = `\s*(.+:\d+):\serror:\s[\+\-]\[(.*)\s(.*)\]\s:(?:\s'.*'\s\[FAILED\],)?\s(.*)`
	UI_FAILING_TEST_MATCHER      = `\s{4}t = \s+\d+\.\d+s\s+Assertion Failure: (.*:\d+): (.*)`
	TEST_CASE_PASSED_MATCHER     = `\s*Test Case\s'-\[(.*)\s(.*)\]'\spassed\s\((\d*\.\d{3})\sseconds\)`
	TEST_CASE_STARTED_MATCHER    = `Test Case '-\[(.*) (.*)\]' started.`
	TEST_CASE_PENDING_MATCHER    = `Test Case\s'-\[(.*)\s(.*)PENDING\]'\spassed`
	TEST_CASE_MEASURED_MATCHER   = `[^:]*:[^:]*:\sTest Case\s'-\[(.*)\s(.*)\]'\smeasured\s\[Time,\sseconds\]\saverage:\s(\d*\.\d{3}),`
	TESTS_RUN_COMPLETION_MATCHER = `\s*Test Suite '(?:.*\/)?(.*[ox]ctest.*)' (finished|passed|failed) at (.*)`
	TEST_SUITE_STARTED_MATCHER   = `\s*Test Suite '(?:.*\/)?(.*[ox]ctest.*)' started at(.*)`
	TEST_SUITE_START_MATCHER     = `\s*Test Suite '(.*)' started at`
	EXECUTED_MATCHER             = `\s*Executed$`
)

type TestParser struct {
	whiteColor              *color.Color
	failingTestMatch        *regexp.Regexp
	uiFailingTestMatch      *regexp.Regexp
	testCasePassedMatch     *regexp.Regexp
	testCaseStartedMatch    *regexp.Regexp
	testCasePendingMatch    *regexp.Regexp
	testCaseMeasuredMatch   *regexp.Regexp
	testsRunCompletionMatch *regexp.Regexp
	testSuiteStartedMatch   *regexp.Regexp
	testSuiteStartMatch     *regexp.Regexp
	executedMatch           *regexp.Regexp
}

func NewTestParser() *TestParser {
	// white foreground
	whiteColor := color.New(color.FgWhite)
	whiteColor.Add(color.Bold)

	return &TestParser{
		whiteColor:              whiteColor,
		failingTestMatch:        regexp.MustCompile(FAILING_TEST_MATCHER),
		uiFailingTestMatch:      regexp.MustCompile(UI_FAILING_TEST_MATCHER),
		testCasePassedMatch:     regexp.MustCompile(TEST_CASE_PASSED_MATCHER),
		testCaseStartedMatch:    regexp.MustCompile(TEST_CASE_STARTED_MATCHER),
		testCasePendingMatch:    regexp.MustCompile(TEST_CASE_PENDING_MATCHER),
		testCaseMeasuredMatch:   regexp.MustCompile(TEST_CASE_MEASURED_MATCHER),
		testsRunCompletionMatch: regexp.MustCompile(TESTS_RUN_COMPLETION_MATCHER),
		testSuiteStartedMatch:   regexp.MustCompile(TEST_SUITE_STARTED_MATCHER),
		testSuiteStartMatch:     regexp.MustCompile(TEST_SUITE_START_MATCHER),
		executedMatch:           regexp.MustCompile(EXECUTED_MATCHER),
	}
}

func (parser *TestParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {

	if match := parser.failingTestMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.uiFailingTestMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.testCasePassedMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.testCaseStartedMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.testCasePendingMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.testCaseMeasuredMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.testsRunCompletionMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.testSuiteStartedMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.testSuiteStartMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}

	if match := parser.executedMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}

	return false
}
