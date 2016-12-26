package variant

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	Release = "release"
	Alpha   = "alpha"
	Beta    = "beta"
)

type Versions struct {
	Current *Version   `json:"current"`
	History []*Version `json:"versions"`
}

func (v *Versions) Len() int {
	return len(v.History)
}

func (v *Versions) Append(vers *Version) {
	v.History = append(v.History, vers)
}

type Version struct {
	Major       int       `json:"major"`
	Minor       int       `json:"minor"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	ReleaseType string    `json:"release_type"`
}

func (v *Version) BumpMajor() {
	v.Major += 1
}

func (v *Version) BumpMinor() {
	v.Minor += 1
}

func NewVersion(description, releaseType string) *Version {
	return &Version{Description: description, ReleaseType: releaseType, Date: time.Now()}
}

func (v *Version) VersionString() string {
	return fmt.Sprintf("%d.%d_%s", v.Major, v.Minor, v.ReleaseType)
}

func Load(path string) (versions *Versions, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &versions)
	return
}

func (v *Versions) Save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	return enc.Encode(v)
}

func (v *Versions) JSON() ([]byte, error) {
	return json.Marshal(v)
}

func (v *Versions) NewMajor(description, releaseType string) {
	vers := NewVersion(description, releaseType)
	if v.Len() > 0 {
		vers.Major = v.Current.Major + 1
	} else {
		vers.Major = 0
	}
	v.Append(vers)
	v.Current = vers
}

func (v *Versions) NewMinor(description, releaseType string) {
	vers := NewVersion(description, releaseType)
	if v.Len() > 0 {
		vers.Minor = v.Current.Minor + 1
	} else {
		vers.Minor = 1
	}
	v.Append(vers)
	v.Current = vers
}