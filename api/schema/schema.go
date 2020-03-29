//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...
package schema

// FROM: https://github.com/tonyghita/graphql-go-example/blob/b146e697830582faa49e95e906c5a562ffa37aa8/schema/schema.go
import (
	"bytes"
)

// String reads the .graphql schema files from the generated _bindata.go file, concatenating the
// files together into one string.
//
// If this method complains about not finding functions AssetNames() or MustAsset(),
// run `go generate` against this package to generate the functions.
func String() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b, err := Asset(name)
		if err != nil {
			panic("asset: Asset(" + name + "): " + err.Error())
		}
		buf.Write(b)

		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
