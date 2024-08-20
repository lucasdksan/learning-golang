## Bibliotecas Nativas

**Pacotes Básicos**

* builtin – Pacote contendo definições de funções, tipos e constantes pré-definidos.


```go
package main

import "fmt"

func main() {
    var x int = 10
    fmt.Println(x)
}
```

* errors – Pacote para manipulação de erros.

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("an error occurred")
    fmt.Println(err)
}
```

**Manipulação de Strings e Texto**

* strings – Funções para manipulação de strings.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, World!"
    fmt.Println(strings.ToUpper(str))
}
```

<details close="close">
<summary>Strings Lib</summary>

<strong>type Builder</strong>

```go
type Builder struct {
	// contains filtered or unexported fields
}
```

Um Builder é usado para construir eficientemente uma string usando métodos Builder.Write . Ele minimiza a cópia de memória. O valor zero está pronto para uso. Não copie um Builder diferente de zero.

<strong>func (*Builder) Cap</strong>

Cap retorna a capacidade da fatia de byte subjacente do construtor. É o espaço total alocado para a string que está sendo construída e inclui quaisquer bytes já escritos.

<strong>func (*Builder) Len</strong>

Len retorna o número de bytes acumulados; b.Len() == len(b.String()).

<strong>func (*Builder) Reset</strong>

Redefinir redefine o Builder para ficar vazio.

<strong>func (*Builder) String</strong>

String retorna a string acumulada.

<strong>func (*Builder) Write</strong>

Write acrescenta o conteúdo de p ao buffer de b. Write sempre retorna len(p), nil.

<strong>func (*Builder) WriteByte</strong>

WriteByte acrescenta o byte c ao buffer de b. O erro retornado é sempre nil.

*A diferença entre Write e WriteByte está na quantidade e no tipo de dados que cada um escreve no strings.Builder*

<strong>func (*Builder) WriteRune</strong>

WriteRune anexa a codificação UTF-8 do ponto de código Unicode r ao buffer de b. Ele retorna o comprimento de r e um erro nil.

<strong>func (*Builder) WriteString</strong>

WriteString acrescenta o conteúdo de s ao buffer de b. Ele retorna o comprimento de s e um erro nil.

<strong>type Reader</strong>

```go
type Reader struct {
	// contains filtered or unexported fields
}
```

Um Reader implementa as interfaces io.Reader , io.ReaderAt , io.ByteReader , io.ByteScanner , io.RuneReader , io.RuneScanner , io.Seeker e io.WriterTo lendo de uma string. O valor zero para Reader opera como um Reader de uma string vazia.

<strong>func NewReader</strong>

NewReader retorna um novo Reader lendo de s. É semelhante a bytes.NewBufferString, mas mais eficiente e não gravável.

<strong>func (*Reader) Len</strong>

Len retorna o número de bytes da parte não lida da string.

<strong>func (*Reader) Read</strong>

Read implementa a interface io.Reader.

<strong>func (*Reader) ReadAt</strong>

ReadAt implementa a interface io.ReaderAt.

<strong>func (*Reader) ReadByte</strong>

ReadByte implementa a interface io.ByteReader.

<strong>func (*Reader) ReadRune</strong>

ReadRune implementa a interface io.RuneReader.

<strong>func (*Reader) Reset</strong>

Redefinir redefine o Leitor para que ele leia de s.

<strong>func (*Reader) Seek</strong>

Seek implementa a interface io.Seeker.

<strong>func (*Reader) Size</strong>

Size retorna o comprimento original da string subjacente. Size é o número de bytes disponíveis para leitura via Reader.ReadAt. O valor retornado é sempre o mesmo e não é afetado por chamadas para nenhum outro método.

<strong>func (*Reader) UnreadByte</strong>

UnreadByte implementa a interface io.ByteScanner.

<strong>func (*Reader) UnreadRune</strong>

UnreadRune implementa a interface io.RuneScanner.

<strong>func (*Reader) WriteTo</strong>

WriteTo implementa a interface io.WriterTo.

<strong>type Replacer</strong>

```go
type Replacer struct {
	// contains filtered or unexported fields
}
```

Replacer substitui uma lista de strings por replacements. É seguro para uso concorrente por várias goroutines.

<strong>func NewReplacer</strong>

NewReplacer retorna um novo Replacer de uma lista de pares de strings antigos e novos. As substituições são realizadas na ordem em que aparecem na string de destino, sem correspondências sobrepostas. As comparações de strings antigas são feitas na ordem dos argumentos.

O NewReplacer entra em pânico se receber um número ímpar de argumentos.

<strong>func (*Replacer) Replace</strong>

Replace retorna uma cópia de s com todas as substituições realizadas.

<strong>func (*Replacer) WriteString</strong>

WriteString escreve s em w com todas as substituições realizadas.

</details>

* strconv – Converte strings para tipos numéricos e vice-versa.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    i, err := strconv.Atoi("123")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(i)
}
```

* regexp – Expressões regulares.

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`\w+`)
    fmt.Println(re.FindAllString("Hello, World!", -1))
}
```

* unicode – Pacote para suporte a Unicode.

```go
package main

import (
    "fmt"
    "unicode"
)

func main() {
    r := 'A'
    fmt.Println(unicode.IsLetter(r))
}
```

* unicode/utf8 – Suporte para codificação UTF-8.

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    str := "Hello, 世界"
    fmt.Println(utf8.RuneCountInString(str))
}
```

**Entrada e Saída (I/O)**

* fmt – Funções de formatação e impressão.

```go
package main

import "fmt"

func main() {
    fmt.Printf("Hello, %s!\n", "World")
}
```

* io – Funções básicas de entrada/saída.

O pacote io fornece interfaces básicas para primitivas de E/S. Sua principal tarefa é encapsular implementações existentes de tais primitivas, como aquelas no pacote os, em interfaces públicas compartilhadas que abstraem a funcionalidade, além de algumas outras primitivas relacionadas.

Como essas interfaces e primitivas envolvem operações de nível inferior com várias implementações, a menos que sejam informados de outra forma, os clientes não devem presumir que elas são seguras para execução paralela.

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    io.WriteString(os.Stdout, "Hello, World!\n")
}
```

<strong>func Copy</strong>

Copy cópias de src para dst até que EOF seja alcançado em src ou ocorra um erro. Ele retorna o número de bytes copiados e o primeiro erro encontrado durante a cópia, se houver.

Uma cópia bem-sucedida retorna err == nil, não err == EOF. Como Copy é definido para ler de src até EOF, ele não trata um EOF de Read como um erro a ser relatado.

Se src implementar WriterTo , a cópia será implementada chamando src.WriteTo(dst). Caso contrário, se dst implementar ReaderFrom , a cópia será implementada chamando dst.ReadFrom(src).

<strong>func CopyBuffer</strong>

func CopyBuffer é idêntico a Copy, exceto que ele passa pelo buffer fornecido (se um for necessário) em vez de alocar um temporário. Se buf for nil, um é alocado; caso contrário, se ele tiver comprimento zero, CopyBuffer entra em pânico.

Se src implementar WriterTo ou dst implementar ReaderFrom , buf não será usado para executar a cópia.

<strong>func Pipe</strong>

func Pipe cria um pipe síncrono na memória. Ele pode ser usado para conectar código esperando um io.Reader com código esperando um io.Writer .

Leituras e Escritas no pipe são correspondidas uma a uma, exceto quando múltiplas Leituras são necessárias para consumir uma única Escrita. Ou seja, cada Escrita no PipeWriter bloqueia até que tenha satisfeito uma ou mais Leituras do PipeReader que consumam completamente os dados escritos. Os dados são copiados diretamente da Escrita para a Leitura correspondente (ou Leituras); não há buffer interno.

É seguro chamar Read e Write em paralelo um com o outro ou com Close. Chamadas paralelas para Read e chamadas paralelas para Write também são seguras: as chamadas individuais serão gated sequencialmente.

* io/ioutil – Utilitários I/O.

```go
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("example.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(data))
}
```

* bufio – I/O bufferizado.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
}
```

**Sistema de Arquivos**

* os – Interface com funcionalidades do sistema operacional.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    file.WriteString("Hello, World!")
}

```

* path/filepath – Manipulação de caminhos de arquivo.

```go
package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    path := filepath.Join("dir", "file.txt")
    fmt.Println(path)
}

```

**Rede e Conectividade**

* net – Redes de baixo nível.

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "google.com:80")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    buf := make([]byte, 4096)
    n, _ := conn.Read(buf)
    fmt.Println(string(buf[:n]))
}

```

* net/http – Cliente e servidor HTTP.

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

```

* net/url – Análise e construção de URLs.

```go
package main

import (
    "fmt"
    "net/url"
)

func main() {
    u, err := url.Parse("https://example.com/path?foo=bar")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(u.Host)
}

```

* net/rpc – Implementação de chamadas de procedimento remoto (RPC).

```go
package main

import (
    "net"
    "net/http"
    "net/rpc"
    "fmt"
)

type Args struct {
    A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func main() {
    arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        fmt.Println(e)
    }
    http.Serve(l, nil)
}

```

**Criptografia**

* crypto – Criptografia básica.

```go
package main

import (
    "crypto/md5"
    "fmt"
    "io"
)

func main() {
    h := md5.New()
    io.WriteString(h, "Hello, World!")
    fmt.Printf("%x", h.Sum(nil))
}

```

* crypto/aes – Implementação do algoritmo AES.

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
)

func main() {
    key := []byte("a very very very very secret key") // 32 bytes
    plaintext := []byte("some really really really long plaintext")

    block, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
    }

    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        fmt.Println(err)
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

    fmt.Printf("%x\n", ciphertext)
}
```

* crypto/cipher – Interfaces para algoritmos de cifra.

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
)

func main() {
    key := []byte("a very very very very secret key") // 32 bytes
    block, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
    }

    plaintext := []byte("exampleplaintext")
    ciphertext := make([]byte, len(plaintext))

    iv := []byte("a very secret iv!") // 16 bytes
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext, plaintext)

    fmt.Printf("%x\n", ciphertext)
}

```

* crypto/sha256 – Implementação do SHA256.

```go
package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    h := sha256.New()
    h.Write([]byte("Hello, World!"))
    fmt.Printf("%x", h.Sum(nil))
}

```

* crypto/rand – Geração de números aleatórios criptograficamente seguros.

```go
package main

import (
    "crypto/rand"
    "fmt"
)

func main() {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(b)
}

```

**Compressão**

* compress/bzip2 – Leitura de arquivos bzip2.

```go
package main

import (
    "compress/bzip2"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("file.bz2")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    reader := bzip2.NewReader(file)
    buf := make([]byte, 1024)
    n, _ := reader.Read(buf)
    fmt.Println(string(buf[:n]))
}

```

* compress/gzip – Leitura e escrita de arquivos gzip.

```go
package main

import (
    "compress/gzip"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("file.gz")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    reader, err := gzip.NewReader(file)
    if err != nil {
        fmt.Println(err)
    }
    defer reader.Close()

    buf := make([]byte, 1024)
    n, _ := reader.Read(buf)
    fmt.Println(string(buf[:n]))
}

```

* compress/zlib – Leitura e escrita de streams zlib.

```go
package main

import (
    "compress/zlib"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("file.zlib")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    reader, err := zlib.NewReader(file)
    if err != nil {
        fmt.Println(err)
    }
    defer reader.Close()

    buf := make([]byte, 1024)
    n, _ := reader.Read(buf)
    fmt.Println(string(buf[:n]))
}

```

**Codificação**

* encoding – Interfaces comuns para pacotes de codificação.

```go
package main

import (
    "encoding"
    "fmt"
)

type MyType struct {
    Value string
}

func (m MyType) MarshalText() (text []byte, err error) {
    return []byte(m.Value), nil
}

func (m *MyType) UnmarshalText(text []byte) error {
    m.Value = string(text)
    return nil
}

func main() {
    var v encoding.TextMarshaler = MyType{"hello"}
    text, _ := v.MarshalText()
    fmt.Println(string(text))
}

```

* encoding/json – Codificação e decodificação JSON.

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Alice", 30}
    data, _ := json.Marshal(p)
    fmt.Println(string(data))

    var p2 Person
    json.Unmarshal(data, &p2)
    fmt.Println(p2)
}

```

* encoding/xml – Codificação e decodificação XML.

```go
package main

import (
    "encoding/xml"
    "fmt"
)

type Person struct {
    Name string `xml:"name"`
    Age  int    `xml:"age"`
}

func main() {
    p := Person{"Alice", 30}
    data, _ := xml.Marshal(p)
    fmt.Println(string(data))

    var p2 Person
    xml.Unmarshal(data, &p2)
    fmt.Println(p2)
}

```

* encoding/base64 – Codificação e decodificação base64.

```go
package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    data := "Hello, World!"
    encoded := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(encoded)

    decoded, _ := base64.StdEncoding.DecodeString(encoded)
    fmt.Println(string(decoded))
}

```

**Sincronização e Concorrência**

* sync – Primitivas de sincronização.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        defer wg.Done()
        fmt.Println("Hello, World!")
    }()

    wg.Wait()
}

```

* sync/atomic – Operações atômicas primitivas.

```go
package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var counter int64 = 0
    atomic.AddInt64(&counter, 1)
    fmt.Println(counter)
}

```

**Tempo e Datas**

* time – Funções para manipulação de tempo e datas.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now.Format(time.RFC3339))

    t := time.Date(2024, time.June, 29, 0, 0, 0, 0, time.UTC)
    fmt.Println(t)
}

```

**Log e Debugging**

* log – Registro básico de logs.

```go
package main

import (
    "log"
)

func main() {
    log.Println("Hello, World!")
}

```

* log/syslog – Registro de logs via syslog.

```go
package main

import (
    "log"
    "log/syslog"
)

func main() {
    logwriter, e := syslog.New(syslog.LOG_NOTICE, "myprogram")
    if e == nil {
        log.SetOutput(logwriter)
    }

    log.Println("Hello, World!")
}

```

**Reflexão**

* reflect – Inspeção de tipos e valores em tempo de execução.

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(x))
    fmt.Println("value:", reflect.ValueOf(x))
}

```

**Outras Utilidades**

* math – Funções matemáticas básicas.

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Sqrt(16))
}

```

* math/rand – Gerador de números aleatórios.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println(rand.Intn(100))
}

```

* sort – Ordenação de coleções.

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println(strs)
}

```

* regexp – Expressões regulares.

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`\w+`)
    fmt.Println(re.FindAllString("Hello, World!", -1))
}

```

* testing – Suporte para testes unitários.

```go
package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(1, 2)
    if result != 3 {
        t.Errorf("Expected 3 but got %d", result)
    }
}

func Add(a, b int) int {
    return a + b
}

```
