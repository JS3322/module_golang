- compress/bzip2 : func NewReader(r io.Reader) io.Reader

- compress/zlib
+ func NweReader(r io.Reader) (io.ReadCloser, error)
+ func NewWriter(w io.Writer) *Writer

- ompress/flate
+ func NewReader(r io.Reader) io.ReadCloser
+ func NewWriter(w io.Writer, level int) (*Writer, error)

- compress/lzw
+ unc NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
+ func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser
