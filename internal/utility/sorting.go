package utility

// IMPORTS {{{
import (
	"slices"

	"github.com/Pairadux/gotm/internal/models"
	"github.com/facette/natsort"
) // }}}

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

func CreatedAsc(ts []models.Task) {
	slices.SortFunc(ts, func(a, b models.Task) int {
		return a.Created.Compare(b.Created)
	})
}
