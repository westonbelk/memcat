package procfs

import (
    "os"
    "io"
)

func (p *Process) PipeBytes(m Mapping) (error) {
    r, w := io.Pipe()
    defer w.Close()

    go func() {
        defer r.Close()
        if _, err := io.Copy(os.Stdout, r); err != nil {
            panic(err)
        }
    }()

    mapSize := int64((m.AddressEnd - m.AddressStart)-1)
    memFD, err := os.Open(p.Directory()+"/mem")
    if err != nil {
        panic(err)
    }
    defer memFD.Close()

    memFD.Seek(int64(m.AddressStart),0)
    io.CopyN(w, memFD, mapSize)
   
    return nil
}
