package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	// ID        int    `json:"id" gorm:"column:id;"`
	ID        string `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"column:cloud_name;"`
	Extension string `json:"extension,omitempty" gorm:"column:extension;"`
}

func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal JSONB value:", value))
	}

	var image Image
	if err := json.Unmarshal(bytes, &image); err != nil {
		return err
	}

	*i = image
	return nil
}

func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}

type Images []Image

func (i *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal JSONB value:", value))
	}

	var images Images
	if err := json.Unmarshal(bytes, &images); err != nil {
		return err
	}

	*i = images
	return nil
}

func (i *Images) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}
