package main

import "fmt"

func main() {
	//belajar map

	fmt.Println("Hai, Ulfa")
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["Januari"] = 50
	chicken["Februari"] = 30

	fmt.Println("Chicken Januari \t: ", chicken["Januari"])
	fmt.Println("Chicken Februari \t: ", chicken["Februari"])

	fmt.Println("Length :", len(chicken))
	fmt.Println(chicken)

	delete(chicken, "Januari")
	fmt.Println("Length :", len(chicken))
	fmt.Println(chicken)

	fmt.Println()

	var fruit = map[string]int{"Melon": 20, "Apel": 40}
	fmt.Println("Fruit Melon \t: ", fruit["Melon"])

	var fruit2 = map[string]int{
		"Anggur": 46,
		"Jagung": 32,
		"Jeruk":  15,
		"Pir":    54,
	}
	fmt.Println("Fruit Jagung \t: ", fruit2["Jagung"])

	fmt.Println()

	for key, val := range fruit2 {
		fmt.Println(key, " \t:", val)
	}
}
