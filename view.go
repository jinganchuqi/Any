package any

import (
/*"log"
"html/template"*/
)
import "fmt"

type View struct {

}

func (v *View) Render(filePath string) {

	fmt.Println(v)
	//t, err := template.ParseFiles(filePath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//data := struct {
	//	Title string
	//	Content string
	//}{
	//	Title: "golang html template demo",
	//	Content: "Hello,World",
	//}
	//err = t.Execute(v.Response, data)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
