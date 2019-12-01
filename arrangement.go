package main

import "io/ioutil"
import "encoding/json"

// Arrangement represent a named preconfigured layout containing a []Window.
type Arrangement struct {
	Windows []Window
	Name    string
}

// Arrangements is a list of arrangements
type Arrangements []Arrangement

// Window represents a singular application window and it's desired position keyed by a Name Regex.  An exclusion name regex can be used for more fine-grained filtering
type Window struct {
	NameRegex         string
	NameExlusionRegex string
	X                 int32
	Y                 int32
	Cx                int32
	Cy                int32
}

// ToJSONFile writes a *[]WindowPreferences to the input file name, creating it if it doesn't exist.
func (a *Arrangements) ToJSONFile(filename string) error {
	data, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// NewFromFile reads the input file and attempts to unmarshal the contents into a *[]WindowPreferences.
func NewFromFile(filename string) (*Arrangements, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var arrangements Arrangements
	err = json.Unmarshal(data, &arrangements)
	if err != nil {
		return nil, err
	}

	return &arrangements, nil
}
