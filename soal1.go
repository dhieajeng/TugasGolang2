package main

import (
	"fmt"
)

func calculateAverage(scores []int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total) / float64(len(scores))
}

func main() {
	// Data uji
	data := [][]int{
		{96, 108, 89},
		{88, 91, 110},
	}

	bonus1 := [][]int{
		{97, 112, 101},
		{109, 95, 123},
	}

	bonus2 := [][]int{
		{97, 112, 101},
		{109, 95, 106},
	}

	// Skor rata-rata untuk setiap tim
	lumbaLumba := calculateAverage(data[0])
	koala := calculateAverage(data[1])

	fmt.Println("Data 1:")
	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", lumbaLumba)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", koala)

	if lumbaLumba > koala {
		fmt.Println("Tim Lumba-lumba memenangkan trofi!")
	} else if koala > lumbaLumba {
		fmt.Println("Tim Koala memenangkan trofi!")
	} else {
		fmt.Println("Kompetisi berakhir seri!")
	}

	// Bonus 1: Skor minimum 100
	fmt.Println("\nBonus 1:")
	lumbaLumbaBonus1 := calculateAverage(bonus1[0])
	koalaBonus1 := calculateAverage(bonus1[1])

	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", lumbaLumbaBonus1)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", koalaBonus1)

	if lumbaLumbaBonus1 >= 100 && koalaBonus1 >= 100 {
		if lumbaLumbaBonus1 > koalaBonus1 {
			fmt.Println("Tim Lumba-lumba memenangkan trofi!")
		} else if koalaBonus1 > lumbaLumbaBonus1 {
			fmt.Println("Tim Koala memenangkan trofi!")
		} else {
			fmt.Println("Kompetisi berakhir seri!")
		}
	} else {
		fmt.Println("Tidak ada pemenang, skor di bawah minimum!")
	}

	// Bonus 2: Skor minimum untuk seri
	fmt.Println("\nBonus 2:")
	lumbaLumbaBonus2 := calculateAverage(bonus2[0])
	koalaBonus2 := calculateAverage(bonus2[1])

	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", lumbaLumbaBonus2)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", koalaBonus2)

	if lumbaLumbaBonus2 >= 100 && koalaBonus2 >= 100 {
		if lumbaLumbaBonus2 > koalaBonus2 {
			fmt.Println("Tim Lumba-lumba memenangkan trofi!")
		} else if koalaBonus2 > lumbaLumbaBonus2 {
			fmt.Println("Tim Koala memenangkan trofi!")
		} else {
			fmt.Println("Kompetisi berakhir seri!")
		}
	} else {
		fmt.Println("Tidak ada pemenang, skor di bawah minimum!")
	}
}
package main

import (
	"fmt"
)

func calculateAverage(scores []int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total) / float64(len(scores))
}

func main() {
	// Data uji
	data := [][]int{
		{96, 108, 89},
		{88, 91, 110},
	}

	bonus1 := [][]int{
		{97, 112, 101},
		{109, 95, 123},
	}

	bonus2 := [][]int{
		{97, 112, 101},
		{109, 95, 106},
	}

	// Skor rata-rata untuk setiap tim
	lumbaLumba := calculateAverage(data[0])
	koala := calculateAverage(data[1])

	fmt.Println("Data 1:")
	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", lumbaLumba)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", koala)

	if lumbaLumba > koala {
		fmt.Println("Tim Lumba-lumba memenangkan trofi!")
	} else if koala > lumbaLumba {
		fmt.Println("Tim Koala memenangkan trofi!")
	} else {
		fmt.Println("Kompetisi berakhir seri!")
	}

	// Bonus 1: Skor minimum 100
	fmt.Println("\nBonus 1:")
	lumbaLumbaBonus1 := calculateAverage(bonus1[0])
	koalaBonus1 := calculateAverage(bonus1[1])

	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", lumbaLumbaBonus1)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", koalaBonus1)

	if lumbaLumbaBonus1 >= 100 && koalaBonus1 >= 100 {
		if lumbaLumbaBonus1 > koalaBonus1 {
			fmt.Println("Tim Lumba-lumba memenangkan trofi!")
		} else if koalaBonus1 > lumbaLumbaBonus1 {
			fmt.Println("Tim Koala memenangkan trofi!")
		} else {
			fmt.Println("Kompetisi berakhir seri!")
		}
	} else {
		fmt.Println("Tidak ada pemenang, skor di bawah minimum!")
	}

	// Bonus 2: Skor minimum untuk seri
	fmt.Println("\nBonus 2:")
	lumbaLumbaBonus2 := calculateAverage(bonus2[0])
	koalaBonus2 := calculateAverage(bonus2[1])

	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", lumbaLumbaBonus2)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", koalaBonus2)

	if lumbaLumbaBonus2 >= 100 && koalaBonus2 >= 100 {
		if lumbaLumbaBonus2 > koalaBonus2 {
			fmt.Println("Tim Lumba-lumba memenangkan trofi!")
		} else if koalaBonus2 > lumbaLumbaBonus2 {
			fmt.Println("Tim Koala memenangkan trofi!")
		} else {
			fmt.Println("Kompetisi berakhir seri!")
		}
	} else {
		fmt.Println("Tidak ada pemenang, skor di bawah minimum!")
	}
}
