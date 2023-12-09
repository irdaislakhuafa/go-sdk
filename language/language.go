package language

type Language string

const (
	English    = Language("EN")
	Indonesian = Language("ID")
	Japanese   = Language("JP")
	Duetsch    = Language("DE")
)

var (
	httpStatusLanguages = map[Language](map[int]string){
		English:    statusTextEN,
		Indonesian: statusTextID,
	}
)

func HTTPStatusText(lang Language, code int) string {
	return httpStatusLanguages[lang][code]
}
