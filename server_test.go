/*
 *   Copyright 2016 Manik Surtani
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package quotaservice

import (
	"testing"
//	"github.com/maniksurtani/quotaservice/buckets/memory"
//	"io/ioutil"
	"fmt"
	"github.com/maniksurtani/quotaservice/buckets/memory"
)

func TestWithNoCfg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic()")
		} else {
			fmt.Print(r)
		}
	}()
	doWork()
}

func doWork() {
	New("/does/not/exist", &memory.BucketFactory{})
//	_, err := ioutil.ReadFile("Does not exist")
//	if err != nil {
//		panic(fmt.Sprintf("Unable to open config file DNE. Error: %v", err))
//	}
//	panic("x")
}