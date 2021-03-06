# Papilo
![Go](https://github.com/thealamu/papilo/workflows/Go/badge.svg)

Stream data processing micro-framework; Read, clean, process and store data using pre-defined and custom pipelines.

Papilo packages common and simple data processing functions for reuse, allows definition of custom components/functions and allows permutations of them in powerful ways using the *Pipes and Filters Architecture*.

## Features
- **Pipeline stages**: Read - Process - Store
- **Extensibility**: Extend by adding custom components
- **Pre-defined pipelines**: Papilo offers pre-defined pipelines in its command line tool
- **Custom pipelines**: Organize components to create a custom pipeline flow
- **Concurrency**: Run multiple pipelines concurrently
- **Network source**: Papilo exposes REST and WebSocket APIs for data ingress
- **Custom source**: Define a custom source for data ingress
- **Multiple formats**: Transform input and output data using transformation components

### Architecture
![Architecture](./images/architecture.svg)

The Pipes and Filters architectural pattern provides a structure for systems that process a stream of data.
In this architecture, data is born from a **data source**, passes through **pipes** to intermediate stages called **filter components** and ends up in a **data sink**. Filter Components are the processing units of the pipeline, a filter can enrich (add information to) data, refine (remove information from) data and transform data by delivering data in some other representation. Any two components are connected by pipes; Pipes are the carriers of data into adjacent components. Although this can be implemented in any language, Go lends itself well to this architecture through the use of channels as pipes.

### Defaults
Papilo offers default sources, sinks and components:

- Sources:
    - File: Read lines from a file
    - Stdin: Read lines from standard input (default)
    - Network: A REST endpoint is exposed on a port
    - WebSocket: Full duplex communication, exposed on a port

- Sinks:
    - File: Write sink data to file
    - Stdout: Write sink data to standard output (default)

- Components:
    - Sum: Continuously push the sum of all previous numbers to the sink


## Examples
Read from stdin, write to stdout:
```go
package main

import "github.com/thealamu/papilo/pkg/papilo"

func main() {
    p := papilo.New()
    p.Run() // Default data source is stdin, default data sink is stdout
}
```
Make every character in stream lowercase:
```go
func lowerCmpt(p *papilo.Pipe) {
	for !p.IsClosed { // read for as long as the pipe is open
		// p.Next returns the next data in the pipe
		d, _ := p.Next()
		byteData, ok := d.([]byte)
		if !ok {
			// we did not receive a []byte, we can be resilient and move on
			continue
		}
		// Write to next pipe
		p.Write(bytes.ToLower(byteData))
	}
}

func main() {
	p := papilo.New()
	p.AddComponent(lowerCmpt)
	p.Run()
}
```