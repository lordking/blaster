package common

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AppendIntoFile(t *testing.T) {

	var err error
	var tmpFile = "tmp.txt"
	var testString = "aabb"
	var errorMessage string

	defer ExecCommand("rm", tmpFile)

	t.Log([]byte(testString[0:2]))
	if err = AppendIntoFile([]byte(testString[0:2]), tmpFile); err != nil {
		t.Error(err)
	}

	t.Log([]byte(testString[2:4]))
	if err = AppendIntoFile([]byte(testString[2:4]), tmpFile); err != nil {
		t.Error(err)
	}

	var data []byte
	if data, err = ioutil.ReadFile(tmpFile); err != nil {
		t.Error(err)
	}

	if err != nil {
		errorMessage = err.Error()
	}
	assert.Nil(t, err, errorMessage)

	result := string(data)
	assert.Equal(t, testString, result, "Result not match expected object.")
}

//测试之前，先在创建tmp目录，并在此下建立多个测试文件。
func Test_ReplaceInDirectory(t *testing.T) {
	err := ReplaceInDirectory("./test", ".go", "github.com/lordking/blaster-seed/http/blog", "test")

	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}
	assert.Nil(t, err, errorMessage)
}
