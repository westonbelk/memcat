package main

import (
    "fmt"
    "flag"
    "strconv"
    "github.com/westonbelk/procgrep/pkg/procfs"
)


func main() {
    var pidFlag = flag.Int("pid", -1, "PID to search")
    flag.Parse()
    fmt.Printf("attaching to pid %d\n", *pidFlag)

    pid := "self"
    if *pidFlag > 0 {
        pid = strconv.Itoa(*pidFlag)
    }
    
    process := procfs.Process{
        Pid: procfs.Pid(pid),
        Maps: nil,
    }
    maps, err := procfs.ReadMap(&process)
    if err != nil {
        panic(err)
    }
    process.Maps = maps

    for _, mapsEntry := range process.Maps {
        if mapsEntry.Perms.Read {
            err := process.PipeBytes(mapsEntry)
            if err != nil {
                fmt.Println(err)
                continue
            }
        }
    }
}
