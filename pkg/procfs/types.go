package procfs

type Process struct {
    Pid Pid
    Maps []Mapping
}


type Pid string

type Mapping struct{
    // Starting address space in the process that the mapping occupies
    AddressStart uint64

    // End of the address space in the process that the mapping occupies
    AddressEnd uint64

    // A set of permissions of the mapping
    Perms Permset

    // The offset into the file/whatever
    Offset uint64

    // The device (major:minor)
    Dev Device

    // The inode on that device
    Inode uint64

    // The file that is backing the mapping.
    // 
    // For ELF files, you can easily coordinate with the offset field by
    // looking at the offset field in the ELF program headers (readelf -l)
    // 
    // There are additional helpful psuedo-paths:
    // [stack]: The initial process's (also know as the main thead's) stack
    // [stack:<tid>] A thread's stack where <tid> is a thread ID. Removed in Linux 4.5
    // [vdso] The virtual dynamically linked shared object.
    // [heap] The process's heap.
    // If the pathname field is blank, this is an anonymous mapping as obtained via mmap(2).
    // There is no easy way to coordinate this back to a process's source,
    // short of running it through gdb, strace, or similar.
    //
    // Pathname is shown unescaped except for newline characters, which are replaced with an
    // octal escape sequence. As a result, it is not possible to determine whether the
    // original pathname contained a newline character or the literal \012 character sequence.
    //
    // If the mapping is file-backed and the file has been deleted, the string " (deleted)" is
    // appended to the pathname. Note that this is ambigous too.
    Pathname string

    Pid* Pid
}

type Permset struct {
    Read bool
    Write bool
    Execute bool // represented by 'x'
    Shared bool
    Private bool // (copy on write)
}

type Device struct {
    Major uint64
    Minor uint64
}

