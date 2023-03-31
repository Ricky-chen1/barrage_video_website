package serializer

import "barrage_video_website/model"

type Collection struct {
	Uid    uint   `json:"uid"`
	Name   string `json:"name"`
	IsOpen bool   `json:"is_open"`
}

func BuildCollection(collection model.Collection) Collection {
	return Collection{
		Uid:    collection.Uid,
		Name:   collection.Name,
		IsOpen: collection.IsOpen,
	}
}
