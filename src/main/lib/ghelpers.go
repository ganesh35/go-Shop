package lib
import(
    "fmt"
    "runtime"
)
func CatchPanic(err *error, functionName string) {
    if r := recover(); r != nil {
        fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

        // Capture the stack trace
        buf := make([]byte, 10000)
        runtime.Stack(buf, false)

        fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))

        if err != nil {
            *err = fmt.Errorf("%v", r)
        }
    } else if err != nil && *err != nil {
        fmt.Printf("%s : ERROR : %v\n", functionName, *err)

        // Capture the stack trace
        buf := make([]byte, 10000)
        runtime.Stack(buf, false)

        fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))
    }
}


func In_array_strings(val string, array []string) (ok bool, i int) {  // Only for string array elements
    for i = range array {
        if ok = array[i] == val; ok {
            return
        }
    }
    return
}

