package tag

import (
	"github.com/goEventListingAPI/entity"
)

//TagServices ... services related to tags
type TagServices interface{
	Tags()([]entity.Tag,[]error)
	Tag(id uint)(*entity.Tag,[]error)
	AddTag(tag *entity.Tag)(*entity.Tag,[]error)
	UpdateTag(tag *entity.Tag)(*entity.Tag,[]error)
	RemoveTag(id uint)(*entity.Tag,[]error)
}
