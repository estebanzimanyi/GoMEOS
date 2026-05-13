package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"

func MeosInitialize(tz string) {
	C.meos_initialize()
	if tz != "" {
		C.meos_initialize_timezone(C.CString(tz))
	}
}

func MeosFinalize() {
	C.meos_finalize()
}
