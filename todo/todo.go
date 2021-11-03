package todo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	Position int
}

type ByPri []Item

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	} else if i.Priority == 3 {
		return "(3)"
	}

	return " "
}

func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].Position < s[j].Position
	}

	return s[i].Priority < s[j].Priority
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err = json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i := range items {
		items[i].Position = i + 1
	}

	return items, nil
}
