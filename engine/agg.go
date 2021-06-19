// Copyright 2021 Rohit Awate
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

func SumInt(items []int) (sum int) {
	for _, item := range items {
		sum += item
	}
	return
}

func SumDouble(items []float64) (sum float64) {
	for _, item := range items {
		sum += item
	}
	return
}

func AvgInt(items []int) float64 {
	return float64(SumInt(items)) / float64(len(items))
}

func AvgDouble(items []float64) float64 {
	return SumDouble(items) / float64(len(items))
}
