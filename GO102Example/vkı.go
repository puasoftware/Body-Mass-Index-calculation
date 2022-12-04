package main

//Vücut Kitle İndeksi

//Formül :   VKI = kg/(boy*boy)

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Kilonuzu Giriniz:")
	scanner.Scan()

	kilo, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Boyunuzu Giriniz: ")
	scanner.Scan()

	boy, _ := strconv.ParseFloat(scanner.Text(), 64)

	vki := kilo / (boy * boy)

	fmt.Printf("Vücut Kitle İndeksiniz: %f", vki)

}
