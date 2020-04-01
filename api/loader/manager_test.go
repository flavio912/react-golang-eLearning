package loader

import (
	"math/rand"
	"testing"

	"github.com/google/uuid"

	"github.com/graph-gophers/dataloader"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type key struct {
	Key string
}

func (k *key) Raw() interface{} {
	return k.Key
}

func (k *key) String() string {
	return k.Key
}

func TestSortManagers(t *testing.T) {
	var (
		managers []gentypes.Manager
		keys     dataloader.Keys
	)

	numToTest := 1000

	for i := 0; i < numToTest; i++ {
		ident := uuid.New().String()
		keys = append(keys, dataloader.StringKey(ident))
		managers = append(managers, gentypes.Manager{
			User: gentypes.User{
				UUID:      ident,
				Email:     "test@test.com",
				FirstName: "Test",
				LastName:  "Person",
				Telephone: "074153232323",
				JobTitle:  "Dev",
				LastLogin: "10/11/2020",
			},
		})
	}

	rand.Seed(1231231235)
	rand.Shuffle(len(managers), func(i, j int) { managers[i], managers[j] = managers[j], managers[i] })
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

	// start := time.Now()
	managers = sortManagers(managers, keys)
	// elapsed := time.Since(start)
	correct := 0
	for i, manager := range managers {
		if manager.UUID == keys[i].String() {
			correct = correct + 1
		}
	}

	if correct < numToTest {
		t.Errorf("Not in correct order, %d out of %d correct", correct, numToTest)
	}
	// fmt.Printf("CORRECT: %d", correct)
	// fmt.Printf("\nTime Taken: %s\n", elapsed)
}
