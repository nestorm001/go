package main

func main() {
	var a = [...]int{10, 25, 32, 11, 6, 36, 18, 22, 5, 7}
	//    var a [10]int = [10]int{10, 25, 32, 11, 6, 36, 18, 22, 5, 7}
	//    var a [10]int = [...]int{10, 25, 32, 11, 6, 36, 18, 22, 5, 7}
	for i := 0; i < len(a); i++ {
		for while := i; while < len(a); while++ {
			if a[i] > a[while] {
				a[i], a[while] = a[while], a[i]
			}
		}
	}
	i := 0
	for i < len(a) {
		print(a[i], "\t")
		i++
	}
}
