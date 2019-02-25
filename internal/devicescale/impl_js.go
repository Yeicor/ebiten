// Copyright 2018 The Ebiten Authors
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

// +build js

package devicescale

import (
	"github.com/gopherjs/gopherwasm/js"
)

func impl(x, y int) float64 {
	ratio := js.Global().Get("window").Get("devicePixelRatio").Float()
	if ratio == 0 {
		ratio = 1
	}
	return ratio
}