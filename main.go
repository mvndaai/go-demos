package main

import (
	"fmt"
	"log"

	"github.com/mvndaai/go-interface-demo/reuseable/shape"
	"github.com/mvndaai/go-interface-demo/reuseable/shape/rectangle"
	"github.com/mvndaai/go-interface-demo/reuseable/shape/triangle"
	"github.com/mvndaai/go-interface-demo/testable/database"
)

func main() {
	if err := mainWithError(); err != nil {
		log.Fatal(err)
	}
}

// mainWithError allows handling all errors in the same way you would in the rest of your application
func mainWithError() error {
	db, err := database.NewDatabase()
	if err != nil {
		return fmt.Errorf("could not create database %w", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("could not close db", err)
		}
	}()

	shapes := []shape.EqualDistant{
		triangle.Triangle{},
		rectangle.Rectangle{},
	}
	for _, lenght := range []float64{10, 20, 30} {
		err := db.Save(fmt.Sprint("length ", lenght))
		if err != nil {
			return fmt.Errorf("could not save lenghth %w", err)
		}

		for _, shp := range shapes {
			perimeter := shp.Perimeter(lenght)
			area := shp.Area(lenght)

			entry := fmt.Sprintf("%T - perimeter:%v - area:%v", shp, perimeter, area)
			err := db.Save(entry)
			if err != nil {
				return fmt.Errorf("could not save %w", err)
			}
		}
	}
	return nil
}
