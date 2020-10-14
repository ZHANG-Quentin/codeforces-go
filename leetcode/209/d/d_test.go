// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`0`, 
			`0`,
		},
		{
			`3`, 
			`2`,
		},
		{
			`6`, 
			`4`,
		},
		{
			`9`, 
			`14`,
		},
		{
			`333`, 
			`393`,
		},
		// TODO 测试参数的下界和上界
		{
			`1000000000`,
			`756249599`,
		},
		{
			`56162567`,
			`40849914`,
		},
	}
	targetCaseNum :=2
	if err := testutil.RunLeetCodeFuncWithExamples(t, minimumOneBitOperations, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-209/problems/minimum-one-bit-operations-to-make-integers-zero/