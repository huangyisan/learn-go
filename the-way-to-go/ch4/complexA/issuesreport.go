package complexA

import (
	"log"
	"os"
	"time"
)
import "html/template"

const templ = `{{.TotalCount}} issues:
{{range .Items}} -------------------
Number: {{.Number}}
User: {{.User.Login}}
Number: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func gen() {
	//report, err := template.New("report").
	//	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	//	Parse(templ)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//
	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	result, err := SearchIssues([]string{"yisan"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
