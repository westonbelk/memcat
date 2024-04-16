# Memcat

## Description

Streams the mapped virtual memory space of a process to stdout. Useful for extracting live memory of processes for forensic investigation or finding interesting strings in a running program.

Originally inspired by the ctf challenge [Home on the Range](https://westonbelk.com/writeups/2024-utctf/Home%20on%20the%20Range.html).

## Usage

```
Usage of ./memcat:
  -pid int
        PID to export the memory of. Defaults to self. (default -1)
```
