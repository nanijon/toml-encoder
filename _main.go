package main

// generate a TOML type from a go type such as a struct.

// NOTE: In this setup, have a size [1] for our slice fields in our structs.
// In this case that slice field is rendered correctly by the toml encoder.
// Check file myfile-1.toml for the output.

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	toml "toml-master"
)

//About : About the personnel (Header) ---------------------------------------//
type About struct {
	Title       template.HTML `toml:"Title"`
	Description template.HTML `toml:"Description"`
	Show        bool          `toml:"Show"`
	Items       [1]AboutItem
}

//AboutItem : contains detail records ----------------------------------------//
type AboutItem struct {
	PersonName     template.HTML
	PersonPosition template.HTML `toml:"PersonPosition"`
	Description    template.HTML `toml:"Description"`
	Image          string        `toml:"Image"`
	Alt            string        `toml:"Alt"`
	Show           bool          `toml:"Show"`
}

// GalleryStore : contains information about a collection of gallaries -------//
type GalleryStore struct {
	Title       template.HTML `toml:"Title"`
	Description template.HTML `toml:"Description"`
	Show        bool          `toml:"Show"`
	Galleries   [1]Gallery
}

// Gallery : contains information about each individual gallary --------------//
type Gallery struct {
	ID            string        `toml:"ID"`
	Name          template.HTML `toml:"Name"`
	Description   template.HTML `toml:"Description"`
	GallerySource string        `toml:"GallerySource"`
	Show          bool          `toml:"Show"`
	Images        [1]Image
}

// Image : contains information about each individual image in the gallery ---//
type Image struct {
	Source  string        `toml:"Source"`
	Alt     string        `toml:"Alt"`
	Caption template.HTML `toml:"Caption"`
	Width   int           `toml:"Width"`
	Height  int           `toml:"Height"`
}

func main() {
	var ab About
	var gs GalleryStore

	var config = map[string]interface{}{
		"About":        ab,
		"GalleryStore": gs,
	}
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		log.Fatal(err)
	}
	// dump values to the command consule.
	fmt.Println("Struct-to-Toml using a map: \n\n", buf.String())

	// write toml to a file
	err := ioutil.WriteFile("myfile-1.toml", buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created toml file myfile-1.toml")
}
