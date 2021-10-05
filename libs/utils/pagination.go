package utils

func GetPagination(page, pageSize int) (int, int) {
	offset := 0
	if page > 0 {
		offset = (page - 1) * pageSize
	}
	return offset, pageSize
}
