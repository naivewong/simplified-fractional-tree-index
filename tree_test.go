package fti

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// NotOk fails the test if an err is nil.
func NotOk(tb testing.TB, err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: expected error, got nothing \033[39m\n\n", filepath.Base(file), line)
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}, msgAndArgs ...interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:%s\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, formatMessage(msgAndArgs), exp, act)
		tb.FailNow()
	}
}

// NotEquals fails the test if exp is equal to act.
func NotEquals(tb testing.TB, exp, act interface{}) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: Expected different exp and got\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func formatMessage(msgAndArgs []interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}

	if msg, ok := msgAndArgs[0].(string); ok {
		return fmt.Sprintf("\n\nmsg: "+msg, msgAndArgs[1:]...)
	}
	return ""
}

func TestTree(t *testing.T) {
	tree := NewTree()
	type testTuple struct{
		key int
		val string
	}
	exp := []testTuple{}
	for i := 100000; i >= 0; i-- {
		exp = append(exp, testTuple{i, strconv.Itoa(i)})
	}
	for _, tuple := range exp {
		tree.Insert(tuple.key, tuple.val)
	}

	t.Log("Tree height", tree.height)

	for _, tuple := range exp {
		node := tree.Search(tuple.key)
		Assert(t, node != nil, "node is nil " + strconv.Itoa(tuple.key))
		Equals(t, node.Value(), tuple.val)
	}

	temp := tree.RangeSearch(1000, 20000)
	Equals(t, len(temp), 19000)
	k := 1000
	for _, node := range temp {
		Equals(t, k, node.Key())
		Equals(t, strconv.Itoa(k), node.Value())
		k++
	}
}