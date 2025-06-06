//Package dog provides an exported func "Years" which takes human years and and turns them into dog years
package dog

//Years takes in a value of type string and type int and multiplys the int by seven and returns a value of type string and type int
func Years(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}