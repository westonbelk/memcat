package procfs

import (
    "fmt"
    "os"
)

func (p *Process) Bytes(m Mapping) ([]byte, error) {
    if !m.Perms.Read {
        return nil, fmt.Errorf("mapped memory region does not have read permissions")
    }

    bufSize := (m.AddressEnd - m.AddressStart)-1
    buf := make([]byte, bufSize, bufSize)
    
    memPath := fmt.Sprintf("/proc/%s/mem", p.Pid)
    memFD, err := os.Open(memPath)
    if err != nil {
        return nil, err
    }
    defer memFD.Close()
    
    _, err = memFD.ReadAt(buf, int64(m.AddressStart))
    if err != nil {
        return nil, err
    }
    
    return buf,nil
}
