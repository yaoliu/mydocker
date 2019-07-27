//__author__ = "YaoYao"
//Date: 2019-07-25
package namespace

/*
#include <errno.h>
#include <stdio.h>
#include <sched.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>

__attribute__((constructor)) void enter_namespace(void){


}

*/
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
