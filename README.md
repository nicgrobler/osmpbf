# osmpbf
forked from github.com/qedus/osmpbf to enable a few small changes for me use case

Package osmpbf is used to decode OpenStreetMap pbf files.

## Installation

```bash
$ go get github.com/nicgrobler/osmpbf
```

## Usage

Usage is similar to `json.Decoder`.

```Go
	f, err := os.Open("greater-london-140324.osm.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := osmpbf.NewDecoder(f)

	// use more memory from the start, it is faster
	d.SetBufferSize(osmpbf.MaxBlobSize)
```
note that we need to supply a list of keys that we want data for. so if looking only for a few tags, add them here, as shown below.
nothing else will be returned
``
	// start decoding with several goroutines, it is faster
	err = d.Start(runtime.GOMAXPROCS(-1), map[string]bool{"building": true, "highway": true, "footpath": true, "power": true})
	if err != nil {
		log.Fatal(err)
	}

	var nc, wc, rc uint64
	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			case *osmpbf.Node:
				// Process Node v.
				nc++
			case *osmpbf.Way:
				// Process Way v.
				wc++
			case *osmpbf.Relation:
				// Process Relation v.
				rc++
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)
```

## Documentation

http://godoc.org/github.com/qedus/osmpbf

## To Do

The parseNodes code has not been tested as I can only find PBF files with DenseNode format.

An Encoder still needs to be created to reverse the process.
