package system
// #include <stdlib.h>
import ( "C"; );

func System(cmd string) int {
    return int(C.system(C.CString(cmd)));
}
