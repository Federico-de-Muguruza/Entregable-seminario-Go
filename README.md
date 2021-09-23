# Parser

### El método Parser(entitiy string) (Result, error) recibe una cadena para crear una entidad formateada de forma correcta y devolverla con o sin error.
- Se sacan los espacios de la cadena para evitar esa clase de problemas.
- Guardo en variables los caracteres correspondientes a cada valor de la estructura final a través de índices para ordenarlo y parsearlos a string para poder concatenar fácil. Para estos casos también utilizo dos métodos para filtrar el primer 0 de la longitud, ya que queda redundante y otro método para devolver los caracteres restante que conforman el Valor y la longitud para luego poder chequear si es correcto.
- Verifico que el Tipo sea "NN" o "TX" y si lo cumple, que el Valor no contenga algún caracter indeseado.
- Parseo la longitud a un entero ya que anteriormente era un string.
- Por último hago chequeos de errType y errLength para saber si hay algún tipo de error y devolver el resultado con o sin errores.

### El método verifyChar(entityLength string) (string) recibe la longitud, que para este punto sería un caracter, para poder filtrarlo si este es un 0.
- Verifico que ese caracter sea un 0, si es así, lo elimino y devuelvo el resultado.

### El método remainingValues(entity string) (string, int) recibe la cadena para devolver los caracteres restantes pertenecientes al Valor y la longitud.
- Devuelvo los caracteres restantes a partir de la posición 4 y la longitud para luego hacer un chequeo.

### El método verifyType(entityType string, entityValue string) (error) verifica que exista el tipo "NN" o "TX" y si existen se fija si hay caracteres indeseados para luego devolver un nil si está todo bien o un error si algo falla.

### El método containsNumber(entityValue string) (error) analiza dada una cadena string si este contiene algún número para poder devolver nil o un error.

### El método containsChar(entity string) (error) analiza dada una cadena string si este contiene algún caracter para poder devolver un nil o un error.
