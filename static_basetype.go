package vk

func translatePublic_Bool32(val Bool32) bool {
	return val != Bool32(FALSE)
}
func translateInternal_Bool32(val bool) Bool32 {
	if val {
		return Bool32(TRUE)
	} else {
		return Bool32(FALSE)
	}
}
