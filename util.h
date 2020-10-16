#ifndef _GOUI_UTIL_
#define _GOUI_UTIL_

#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>
//#include "_cgo_export.h"

//extern void goLog(const char *s);

inline int notEmpty(const char* s) {
    return s!=0 && s[0]!='\0';
}


#endif