package messageType

const (
	Notice    int64 = 10
	Marketing int64 = 20
	AuthCode  int64 = 30
)

var TypeCodeEn = map[int64]string{
	Notice:    "notice",
	Marketing: "marketing",
	AuthCode:  "auth_code",
}
