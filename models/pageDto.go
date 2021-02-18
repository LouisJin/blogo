package models

type PageDto struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func GetDefaultPage() PageDto {
	return PageDto{
		Page: 0,
		Size: 20,
	}
}

func GetLimit(page, size int) (limit, offset int) {
	if size == 0 {
		size = GetDefaultPage().Size
	}
	limit = size
	offset = page * size
	return
}
