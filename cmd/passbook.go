package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/mdigger/commitfile"
	"github.com/mdigger/passbook"
	"github.com/mdigger/pkcs7sign"
)

func main() {
	log.SetFlags(0)
	// инициализируем параметры для приложения
	var certFilename, privFilename, passwd string
	flag.StringVar(&certFilename, "cert", "cert.cer", "file with x509 Certificate")
	flag.StringVar(&privFilename, "key", "key.pem", "file with Private key")
	flag.StringVar(&passwd, "pass", "", "password for Private key")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage of %s [options] filename [dir with files]:\nOptions:\n",
			os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("Empty filename for passbook file")
	}
	// загружаем сертификат для подписи
	log.Printf("Loading sertificate %q", certFilename)
	cert, err := pkcs7.LoadCertificate(certFilename)
	if err != nil {
		log.Fatalln("Error reading certificate file:", err)
	}
	// загружаем приватный ключ для подписи
	log.Printf("Loading private key %q", privFilename)
	priv, err := pkcs7.LoadPKCS1PrivateKeyPEM(privFilename, passwd)
	if err != nil {
		log.Fatalln("Error reading private key:", err)
	}
	// получаем имя результирующего файла с passbook
	filename := flag.Arg(0)
	if filepath.Ext(filename) != ".pkpass" {
		filename += ".pkpass"
	}
	// создаем временный файл с passbook
	log.Printf("Creating %q", filename)
	passbookFile, err := commitfile.Create(filename)
	if err != nil {
		log.Fatalln("Error creating passbook file:", err)
	}
	// инициализируем создание passbook
	passbookWrite := passbook.NewWriter(passbookFile, cert, priv)
	// инициализируем путь до исходных файлов
	base := flag.Arg(1) // второй аргумент в параметрах
	if base == "" {
		base = "."
	}
	// log.Printf("Working dir is %q", base)
	// перебираем все файлы в указанном каталоге
	if err := filepath.Walk(base, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		name, err := filepath.Rel(base, filename) // вычисляем локальное имя для добавления
		if err != nil {
			return err
		}
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
		log.Printf("Adding %q", name)
		file, err := os.Open(filename) // открываем файл для чтения
		if err != nil {
			return err
		}
		// нормализуем имя и добавляем содержимое
		err = passbookWrite.Add(filepath.ToSlash(name), file)
		file.Close()
		return err
	}); err != nil {
		passbookFile.Close()
		log.Fatalln("Error adding file:", err)
	}
	// завершаем формирование passbook
	if err := passbookWrite.Close(); err != nil {
		passbookFile.Close()
		log.Fatalln("Error signing:", err)
	}
	passbookFile.Commit() // подтверждаем успешное создание
	if err := passbookFile.Close(); err != nil {
		log.Fatalln("Error closing file:", err)
	}
	log.Printf("Passbook file %q created\n", filename)
}
