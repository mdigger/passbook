package passbook

import (
	"archive/zip"
	"bytes"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"path"

	"github.com/mdigger/pkcs7sign"
)

var ErrNoPass = errors.New("pass.json missed") // the description of pass.json was not added to the file

// Writer allows you to write files in Apple Passbook format.
type Writer struct {
	zip      *zip.Writer       // Packer
	cert     *x509.Certificate // Certificate used for signature
	priv     *rsa.PrivateKey   // Private key used for signature
	hasPass  bool              // Flag that description of passbook added
	manifest map[string]string // Hash of files
}

// NewWriter creates a new Writer that allows you to create an Apple Passbook file.
// As parameters, a stream is passed to which the given file will be written,
// as well as the certificates that will be used to create the digital signature.
func NewWriter(out io.Writer, cert *x509.Certificate, priv *rsa.PrivateKey) *Writer {
	return &Writer{
		zip:      zip.NewWriter(out), // сжимаем при записи
		cert:     cert,
		priv:     priv,
		manifest: make(map[string]string),
	}
}

// Close finishes writing an Apple Passbook file and adds it automatically
// generated manifest and signature file. At the time of creating a digital signature,
// error, which in this case will also be returned. In addition, the error will return if
// the description of pass.json was not added to the file.
func (w *Writer) Close() (err error) {
	if w.zip == nil {
		return nil
	}
	// Do not forget to close the compression in any case
	defer func() {
		if closeErr := w.zip.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
		w.zip = nil
	}()
	// Check that the main description has been added
	if !w.hasPass {
		return ErrNoPass
	}
	// Translate the manifest into JSON
	manifestData, err := json.MarshalIndent(w.manifest, "", "\t")
	if err != nil {
		return err
	}
	// We record the manifest data
	manifestWriter, err := w.zip.Create("manifest.json")
	if err != nil {
		return err
	}
	if _, err = manifestWriter.Write(manifestData); err != nil {
		return err
	}
	// Create a signature
	signature, err := pkcs7.Sign(bytes.NewReader(manifestData), w.cert, w.priv)
	if err != nil {
		return err
	}
	// Write the signature to a file
	signatureWriter, err := w.zip.Create("signature")
	if err != nil {
		return err
	}
	if _, err = signatureWriter.Write(signature); err != nil {
		return err
	}
	return nil
}

// Add adds a new file to the Passbook. Only files with the extension .png and .strings are added.
// Plus, a file called pass.json is added, which is a direct description.
// All other files are ignored.
func (w *Writer) Add(name string, r io.Reader) error {
	if w.zip == nil {
		return io.ErrClosedPipe // write stream closed
	}
	// Ignore unhandled files
	switch path.Ext(name) {
	case ".json": // From json-files we add only the description directly
		if name != "pass.json" {
			return nil
		}
	case ".png": // picture
	case ".strings": // localization
	default: // Everything else is ignored
		return nil
	}
	zipw, err := w.zip.Create(name) // Create a new file in the archive
	if err != nil {
		return err
	}
	hash := sha1.New() // Initialize hash counting
	// At the same time we write to the archive and consider a hash
	if _, err := io.Copy(io.MultiWriter(zipw, hash), r); err != nil {
		return err
	}
	w.manifest[name] = hex.EncodeToString(hash.Sum(nil)) // Save received hash
	if name == "pass.json" {
		w.hasPass = true // Save the flag that the main description is added
	}
	return nil
}
