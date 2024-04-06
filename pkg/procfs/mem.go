package procfs

import (
    "os"
    "io"
)

func (p *Process) PipeBytes(m Mapping) (error) {
    mapSize := int64((m.AddressEnd - m.AddressStart)-1)
    memFD, err := os.Open(p.Pid.Dir()+"/mem")
    if err != nil {
        return err
    }
    defer memFD.Close()

    memFD.Seek(int64(m.AddressStart),0)
    _, err = io.CopyN(os.Stdout, memFD, mapSize)

    return err
}
