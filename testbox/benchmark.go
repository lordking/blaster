package testbox

import (
	"reflect"
	"sync"
	"time"

	"github.com/lordking/blaster/log"
)

//BenchmarkTestCase 性能测试使用的测试用例的接口框架
type BenchmarkTestCase interface {
	GetTestRun(no, cnt int)        //获取每次测试案例的执行前的状态，每个批次的序号，总计计数
	GetCompletedSynchrony(cnt int) //获取每次并发结束的状态，每个批次的总计数
}

//BenchmarkCall 性能测试时，对某个实例的某个方法进行反复调用
//perLimit 每次请求的最大并发次数
func BenchmarkCall(perLimit, max int, testcase BenchmarkTestCase, methodName string) {

	if testcase == nil {
		log.Fatalf("Not found testcase!")
	}

	if methodName == "" {
		log.Fatalf("Not found method!")
		return
	}

	v := reflect.ValueOf(testcase)
	var wg sync.WaitGroup
	var cnt int
	for {

		for i := 0; i < perLimit; i++ {
			cnt++
			wg.Add(1)

			go func(v *reflect.Value, cnt, i int) {
				defer wg.Done()

				testcase.GetTestRun(i, cnt)
				r := v.MethodByName(methodName)
				r.Call(nil)

			}(&v, cnt, i)
		}

		wg.Wait()
		time.Sleep(1000 * time.Millisecond)
		testcase.GetCompletedSynchrony(cnt)

		if cnt >= max && max != -1 {
			break
		}

	}

}
