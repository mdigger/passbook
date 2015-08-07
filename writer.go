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

var ErrNoPass = errors.New("pass.json missed") // в файл не было добавлено описание pass.json

// Writer позволяет записывать файлы в формате Apple Passbook.
type Writer struct {
	zip      *zip.Writer       // упаковщик
	cert     *x509.Certificate // сертификат, используемый для подписи
	priv     *rsa.PrivateKey   // приватный ключ, используемый для подписи
	hasPass  bool              // флаг, что описание passbok добавлено
	manifest map[string]string // хеш файлов
}

// NewWriter создает новый Writer, позволяющий создать файл в формате Apple Passbook.
// В качестве параметров передается поток, в который будет записываться отдаваемый файл,
// а так же сертификаты, которые будут использоваться при создании цифровой подписи.
func NewWriter(out io.Writer, cert *x509.Certificate, priv *rsa.PrivateKey) *Writer {
	return &Writer{
		zip:      zip.NewWriter(out), // сжимаем при записи
		cert:     cert,
		priv:     priv,
		manifest: make(map[string]string),
	}
}

// Close заканчивает запись файла в формате Apple Passbook и добавляет в него автоматически
// сгенерированный манифест и файл подписи. В момент создания цифровой подписи может возникнуть
// ошибка, которая в данном случае так же будет возвращена. Кроме этого, ошибка вернется, если
// в файл не было добавлено описание pass.json.
func (w *Writer) Close() (err error) {
	if w.zip == nil {
		return nil
	}
	// не забываем в любом случае закрыть сжатие
	defer func() {
		if closeErr := w.zip.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
		w.zip = nil
	}()
	// проверяем, что основное описание было добавлено
	if !w.hasPass {
		return ErrNoPass
	}
	// преобразуем манифест в JSON
	manifestData, err := json.MarshalIndent(w.manifest, "", "\t")
	if err != nil {
		return err
	}
	// записываем данные манифеста
	manifestWriter, err := w.zip.Create("manifest.json")
	if err != nil {
		return err
	}
	if _, err = manifestWriter.Write(manifestData); err != nil {
		return err
	}
	// создаем сигнатуру
	signature, err := pkcs7.Sign(bytes.NewReader(manifestData), w.cert, w.priv)
	if err != nil {
		return err
	}
	// записываем сигнатуру в файл
	signatureWriter, err := w.zip.Create("signature")
	if err != nil {
		return err
	}
	if _, err = signatureWriter.Write(signature); err != nil {
		return err
	}
	return nil
}

// Add добавляет новый файл в Passbook. Добавляются только файлы с расширением .png и .strings.
// Плюс, добавляется файл с именем pass.json, который и является непосредственным описанием.
// Все остальные файлы игнорируются.
func (w *Writer) Add(name string, r io.Reader) error {
	if w.zip == nil {
		return io.ErrClosedPipe // поток для записи закрыт
	}
	// игнорируем необрабатываемые файлы
	switch path.Ext(name) {
	case ".json": // из json-файлов добавляем только непосредственно описание
		if name != "pass.json" {
			return nil
		}
	case ".png": // картинки
	case ".strings": // локализация
	default: // все остальное игнорируем
		return nil
	}
	zipw, err := w.zip.Create(name) // создаем новый файл в архиве
	if err != nil {
		return err
	}
	hash := sha1.New() // инициализируем подсчет хеша
	// одновременно записываем в архив и считаем хеш
	if _, err := io.Copy(io.MultiWriter(zipw, hash), r); err != nil {
		return err
	}
	w.manifest[name] = hex.EncodeToString(hash.Sum(nil)) // сохраняем полученный хеш
	if name == "pass.json" {
		w.hasPass = true // сохраняем флаг, что основное описание добавлено
	}
	return nil
}
