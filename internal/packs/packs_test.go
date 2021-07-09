package packs

import (
	"sort"
	"testing"
)

func TestPacksSort(t *testing.T) {
	unorderedPacks := Packs{
		Pack{Size: 5000, Count: 1},
		Pack{Size: 250, Count: 1},
		Pack{Size: 1000, Count: 1},
		Pack{Size: 500, Count: 1},
	}

	expectedPacks := Packs{
		Pack{Size: 250, Count: 1},
		Pack{Size: 500, Count: 1},
		Pack{Size: 1000, Count: 1},
		Pack{Size: 5000, Count: 1},
	}

	sort.Sort(unorderedPacks)

	for i, pack := range unorderedPacks {
		actual := pack.Size
		expected := expectedPacks[i].Size
		if actual != expected {
			t.Errorf("not in order, expected - %v, actual - %v\n", expected, actual)
		}
	}
}
