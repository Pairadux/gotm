/* LICENSE {{{
Copyright © 2025 Austin Gause <a.gause@outlook.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/ // }}}

package utility

// IMPORTS {{{
import (
	"slices"

	"github.com/Pairadux/gotm/internal/models"
	"github.com/facette/natsort"
) // }}}

// NaturalSort sorts tasks (ts) naturally
func NaturalSort(ts []models.Task) {
	slices.SortFunc(ts, func(a, b models.Task) int {
		if natsort.Compare(a.Description, b.Description) {
			return -1
		}
		if natsort.Compare(b.Description, a.Description) {
			return 1
		}
		return 0
	})
}

// CreatedAsc sorts tasks (ts) and sorts them based on the created field in ascending order
func CreatedAsc(ts []models.Task) {
	slices.SortFunc(ts, func(a, b models.Task) int {
		return a.Created.Compare(b.Created)
	})
}
