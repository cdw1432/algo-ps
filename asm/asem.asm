section .data
    message db "ASM, Hello World", 0Ah, 00h

section .text
    global _start

_start:
    ; write system call
    mov rax, 0x2000004       ; system call for write (syscall number for macOS)
    mov rdi, 1               ; file descriptor 1 is stdout
    lea rsi, [rel message]   ; load effective address of message
    mov rdx, 18              ; number of bytes, including the newline and null terminator
    syscall                  ; execute syscall (write)

    ; exit system call
    mov rax, 0x2000001       ; system call for exit (syscall number for macOS)
    xor rdi, rdi             ; exit code 0
    syscall                  ; execute syscall (exit)

; to compile(macOS)
;   yasm -f macho64 asem.asm -o asem.o 
;   gcc -o ./bin/program asem.o -e _start
;   ./bin/program

