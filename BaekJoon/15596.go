package main

func sum(a []int) int {
	var r int
    
    for _, i := range a {
        r += i
    }
    
	return r
}
