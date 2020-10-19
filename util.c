#include "util.h"

extern void goLog(const char *s);
void WULog(const char *format, ...) {
    const int bufSize = 512;
    char buf[bufSize];
    va_list args;
    va_start(args,format);
    vsnprintf(buf,bufSize, format,args);
    goLog(buf);
    va_end(args);
}