## JSON

Em Go, trabalhar com JSON é uma tarefa comum, especialmente ao lidar com comunicação entre sistemas ou armazenamento de dados em arquivos. Vamos explorar algumas técnicas avançadas para lidar com JSON em Go.

**Marshal e Unmarshal**

As funções json.Marshal e json.Unmarshal são usadas para codificar e decodificar dados JSON em Go, respectivamente. No entanto, existem casos em que precisamos de mais controle sobre esse processo.

**Tags de Campo JSON**

É possível definir tags de campo em estruturas em Go para personalizar a serialização e desserialização de dados JSON. Isso permite controlar o nome dos campos, ignorar campos, definir opções de omição e muito mais.

```go
type Person struct {
    Name string `json:"nome"`
    Age  int    `json:"idade"`
}
```

**Codificadores e Decodificadores Personalizados**

Em alguns casos, você pode precisar de uma lógica personalizada para codificar ou decodificar dados JSON. Você pode implementar a interface json.Marshaler e json.Unmarshaler em seus tipos para fornecer essa lógica personalizada.

```go
type CustomData struct {
    Value int
}

func (c *CustomData) UnmarshalJSON(data []byte) error {
    // Lógica personalizada para decodificar JSON
}

func (c *CustomData) MarshalJSON() ([]byte, error) {
    // Lógica personalizada para codificar JSON
}
```

**Validação de Estrutura JSON**

O pacote encoding/json em Go não fornece suporte nativo para validação de estrutura JSON. No entanto, você pode usar outras bibliotecas, como github.com/go-playground/validator, para validar estruturas de dados JSON.

**Decodificação de JSON Aninhado**

Quando trabalhamos com JSON aninhado, às vezes precisamos acessar campos dentro de campos. Para isso, podemos usar tipos de estruturas aninhadas ou tipos de mapa (map[string]interface{}) para decodificar o JSON.

**Codificação Condicional**

Às vezes, queremos codificar apenas um subconjunto dos campos de uma estrutura em JSON, dependendo de certas condições. Podemos usar a técnica de criar uma estrutura separada que inclui apenas os campos desejados e, em seguida, codificar essa estrutura em JSON.

**Manipulação de Grandes Dados JSON**

Quando trabalhamos com grandes volumes de dados JSON, devemos considerar o uso de decodificação incremental ou streaming para evitar a sobrecarga de memória. Isso pode ser feito usando json.Decoder para ler o JSON em partes e processá-lo conforme necessário.
