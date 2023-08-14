# bufpool

## Introduction


## Usage

The following is a basic example of using bufpool:

Import the required package:

```go
import "github.com/lambertxiao/go-bufpool"
```

1. Create a new memory pool:

```go
pool := bufpool.NewGobufpool(uint, func() interface{}) *Gobufpool // Pass the appropriate size and generation function according to actual needs
```

2. Retrieve a memory block from the memory pool:

```go
item := pool.Get() // If the method is called when the memory pool is empty, it will wait indefinitely until a memory block is returned to the pool.

or

item := pool.GetByTime(time.Second) // If an available memory block is not retrieved within the specified timeout time, the method returns a null value (zero value of the corresponding type).
```

3. Return the memory block to the memory pool:

```go
pool.Put(item) // Return item to the pool for future use
```

4. Check the size or capacity of the memory pool:

```go
capacity := pool.Cap() // Get the capacity of the memory pool
```

5. Destroy the memory pool (empty all memory blocks):

```go
pool.Destory() // Clear all memory blocks in the memory pool
```
