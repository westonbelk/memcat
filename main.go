package main

import (
    "fmt"
    "flag"
    "strconv"
    "os"
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
    maps, err := procfs.ReadMap(process.Pid)
    if err != nil {
        panic(err)
    }
    process.Maps = maps

    for i, mapsEntry := range process.Maps {
        if mapsEntry.Perms.Read {
            b, err := process.Bytes(mapsEntry)
            if err != nil {
                fmt.Println(err)
                continue
            }
            outFD, err := os.Create(fmt.Sprintf("out/%d", i))
            if err != nil {
                panic(err)
            }
            defer outFD.Close()
            outFD.Write(b)
        }
    }
}
