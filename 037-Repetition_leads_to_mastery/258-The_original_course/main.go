package main

import "fmt"

func main(){
	s1 := "The original course"
	s2 := "Everything from this point forward is the ORIGINAL COURSE."
	s3 := "I have created this course twice."
	s4 := "The course you just went through is the new course. I created the new course to have more hands-on exercise."
	s5 := "The course that follows here is the original course. "
	s6 := "If you want to go through this original course, you will reinforce what you have learned."
	s7 := "You will also gain new perspectives as I cover some aspects in a slightly different way."
	s8 := "I also cover some things in this original course which were not covered in the course you just completed."
	s9 := "I'm leaving this course here as many students have enjoyed and benefitted from this material."
	xs := []string{s1, s2, s3, s4, s5, s6, s7, s8, s9}
	for _, v := range xs {
		fmt.Println(v)
	}
}