package complexA

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year  int `json:"released"`
	// omitempty 如果为零值，则忽略该字段
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casblanca", Year: 1999, Color: false, Actors: []string{"abc", "def"}},
	{Title: "Casblanca01", Year: 2000, Color: true, Actors: []string{"hij", "kql", "xyz"}},
	{Title: "Casblanca02", Year: 2222, Actors: []string{"hij", "kql", "xyz"}},
}

func DataMashal() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json Marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func DataMashalIdent() {
	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("json Marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// Unmarshal
	var titles []struct {
		Title string
		// 需要对齐才可以解析
		Year int `json:"released"`
	}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("json unmarshling failed: %s", err)
	}
	fmt.Printf("%v\n", titles)
}
