---
title: "Profiling"
date: 2023-04-06T16:56:59+13:00
draft: false
---

Program analysis tools are extremely important for understanding program behavior. Computer architects need such tools to evaluate how well programs will perform on new architectures. Software writers need tools to analyze their programs and identify critical sections of code. Compiler writers often use such tools to find out how well their instruction scheduling or branch prediction algorithm is performing...

The output of a profiler may be:

A statistical summary of the events observed (a profile)
Summary profile information is often shown annotated against the source code statements where the events occur, so the size of measurement data is linear to the code size of the program
