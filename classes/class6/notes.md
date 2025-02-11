## Memory layout for slices

```go
type StructWithSlice struct {
    A []
}
```

| **Field**      | **Type** | **Size** | **Offset** |
| -------------- | -------- | -------- | ---------- |
| A.ptr          | uintptr  | 8 B      | 0          |
| A.len          | int      | 8 B      | 8          |
| A.cap          | int      | 8 B      | 16         |
| **Total Size** | -        | **24 B** | -          |


## Memory layout for arrays

```go
type UnOptimizedStruct struct {
	a int32
    b [5]int32
	c int16
}
```

| **Field**   | **Type** | **Size** | **Offset** |
| ----------- | -------- | -------- | ---------- |
| a           | int32    | 4 B      | 0          |
| b[0]        | int32    | 4 B      | 4          |
| b[1]        | int32    | 4 B      | 8          |
| b[2]        | int32    | 4 B      | 12         |
| b[3]        | int32    | 4 B      | 16         |
| b[4]        | int32    | 4 B      | 20         |
| c           | int16    | 2 B      | 24         |