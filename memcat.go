package main

import (
    "fmt"
    "flag"
    "os"
    "strconv"
    "github.com/westonbelk/memcat/pkg/procfs"
)

var pidFlag int

func initFlags() {
    flag.IntVar(&pidFlag, "pid", -1, "PID to export the memory of. Defaults to self.")
    flag.Parse()
}

func main() {
    initFlags()
    fmt.Fprintf(os.Stderr, "attaching to pid %d\n", pidFlag)
    
    // determine the pid
    pid := procfs.Pid("self")
    if pidFlag > 0 {
        pid = procfs.Pid(strconv.Itoa(pidFlag))
    }

    // read current mapped memory spaces
    maps, err := procfs.ReadMap(pid)
    if err != nil {
        panic(err)
    }

    // create process struct
    process := procfs.Process{
        Pid: pid,
        Maps: maps,
    }

    // read maps and output to stdout
    for _, mapsEntry := range process.Maps {
        if mapsEntry.Perms.Read {
            err := process.PipeBytes(mapsEntry)
            if err != nil {
                fmt.Fprintf(os.Stderr, "%s for pathname %s\n", err, mapsEntry.Pathname)
            }
        }
    }
}
