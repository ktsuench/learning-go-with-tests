package main

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if val, ok := d[key]; ok {
		return val, nil
	}

	return "", error.Error("")
}
